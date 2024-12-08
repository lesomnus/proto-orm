package graph

import (
	"errors"
	"fmt"
	"maps"
	"slices"
	"strings"

	orm "github.com/lesomnus/proto-orm"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/encoding/protowire"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type Field interface {
	FullName() protoreflect.FullName
	Name() string
	Number() protowire.Number
	ProtoType() string
	GoName() string

	Source() *protogen.Field
	Entity() *Entity

	IsIgnored() bool
	IsOptional() bool
	IsList() bool
	IsUnique() bool
	IsImmutable() bool
	IsNullable() bool

	setImmutable()
	setNullable()
}

type Entity struct {
	File   *protogen.File // Proto file where the message defines this entity is.
	Source *protogen.Message

	Key     *ScalarField
	Fields  map[string]Field
	Scalars map[string]*ScalarField
	Edges   map[string]*Edge
	Indexes []*Index
	Rpcs    map[RpcOp]*Rpc
}

func newEntity(f *protogen.File, m *protogen.Message) *Entity {
	return &Entity{
		File:   f,
		Source: m,

		Fields:  map[string]Field{},
		Scalars: map[string]*ScalarField{},
		Edges:   map[string]*Edge{},
		Indexes: []*Index{},
		Rpcs:    map[RpcOp]*Rpc{},
	}
}

func (e *Entity) prase(g Graph, o *orm.MessageOptions) error {
	if o == nil {
		o = &orm.MessageOptions{}
	}

	errs := []error{}
	for _, f := range e.Source.Fields {
		o := f.Desc.Options()
		if f.Desc.IsMap() {
			// TODO: warns map is not supported
			continue
		}

		of := proto.GetExtension(o, orm.E_Field).(*orm.FieldOptions)
		oe := proto.GetExtension(o, orm.E_Edge).(*orm.EdgeOptions)

		t := resolveType(f, of)
		if t != orm.Type_TYPE_MESSAGE {
			if f.Desc.IsList() {
				// TODO: warns list field is not supported.
				continue
			}
			if oe != nil {
				// TODO: warns field is not an edge
			}

			if _, err := e.parseScalarField(f, of); err != nil {
				err = fmt.Errorf(`add field "%s": %w`, f.Desc.Name(), err)
				errs = append(errs, err)
				continue
			}
		} else {
			if of != nil {
				// TODO: warns field is an edge
			}

			target_name := f.Message.Desc.FullName()
			target, ok := g[target_name]
			if !ok {
				err := fmt.Errorf(`%s: target "%s" does not exist`, f.Desc.FullName(), target_name)
				errs = append(errs, err)
				continue
			}
			if oe != nil && oe.From > 0 {
				n := protowire.Number(oe.From)
				f := target.FieldByNumber(n)
				if f == nil {
					// Reference may not be parsed yet so parse it first.
					i := slices.IndexFunc(target.Source.Fields, func(v *protogen.Field) bool {
						return v.Desc.Number() == n
					})
					if i < 0 {
						// Back-reference not found.
						// This error will be reported by `e.parseEdge`.
					} else if r := target.Source.Fields[i]; r.Message.Desc.FullName() != e.FullName() {
						// Back-reference does not references this entity.
						// This error will be reported by `e.parseEdge`.
					} else {
						r := target.Source.Fields[i]
						o := proto.GetExtension(r.Desc.Options(), orm.E_Edge).(*orm.EdgeOptions)
						target.parseEdge(r, o, e)
					}
				}
			}
			if _, err := e.parseEdge(f, oe, target); err != nil {
				err = fmt.Errorf(`add edge "%s": %w`, f.Desc.Name(), err)
				errs = append(errs, err)
				continue
			}
		}
	}

	for _, i := range o.Indexes {
		refs := []Field{}
		for _, r := range i.Refs {
			f := e.FieldByNumber(protowire.Number(r))
			if f == nil {
				err := fmt.Errorf("index reference not found: %d", r)
				errs = append(errs, err)
				continue
			}
			refs = append(refs, f)

			if i.Immutable {
				f.setImmutable()
			}

			// Note that uniqueness of the index does not make its members unique.
			// Assume unique index of columns X and Y, there may be (a, b) and (a, c), so X is not unique.
		}

		w := &Index{
			name:   i.Name,
			Entity: e,
			Refs:   refs,

			Unique:    i.Unique,
			Immutable: i.Immutable,
		}

		e.Indexes = append(e.Indexes, w)
	}

	if len(errs) > 0 {
		return errors.Join(errs...)
	}

	rpc_opts := orm.ResolveRpcOptions(e.File, e.Source)
	r := NewRpcParser()
	e.Rpcs = r.Parse(e, rpc_opts)

	return nil
}

type Key interface {
	Name() string
}

func (e *Entity) KeyLikes() []Key {
	vs := []Key{}
	for _, v := range e.FieldsSortByNumber() {
		if !v.IsUnique() {
			continue
		}
		if _, ok := v.(*Edge); ok {
			// TODO: it can be a key if O2O.
			continue
		}

		vs = append(vs, v)
	}
	for _, v := range e.Indexes {
		if !v.IsUnique() {
			continue
		}

		vs = append(vs, v)
	}

	return vs
}

func (e *Entity) FullName() protoreflect.FullName {
	return e.Source.Desc.FullName()
}

func (e *Entity) Name() string {
	return string(e.Source.Desc.Name())
}

func (e *Entity) Filename() string {
	return strings.ToLower(e.Name()) + ".go"
}

func (e *Entity) HasIndexes() bool {
	return len(e.Indexes) > 0
}

func (e *Entity) FieldByNumber(n protowire.Number) Field {
	for _, f := range e.Fields {
		if f.Number() == n {
			return f
		}
	}
	return nil
}

func (e *Entity) FieldsSortByNumber() []Field {
	fs := slices.Collect(maps.Values(e.Fields))
	slices.SortFunc(fs, func(a Field, b Field) int {
		return int(a.Number()) - int(b.Number())
	})

	return fs
}

func (e *Entity) ScalarsSortByNumber() []*ScalarField {
	fs := slices.Collect(maps.Values(e.Scalars))
	slices.SortFunc(fs, func(a *ScalarField, b *ScalarField) int {
		return int(a.Number()) - int(b.Number())
	})

	return fs
}

func (e *Entity) EdgesSortByNumber() []*Edge {
	fs := slices.Collect(maps.Values(e.Edges))
	slices.SortFunc(fs, func(a *Edge, b *Edge) int {
		return int(a.Number()) - int(b.Number())
	})

	return fs
}

var scalar_mapped_known_messages = map[protoreflect.FullName]orm.Type{
	"google.protobuf.Timestamp": orm.Type_TYPE_TIME,
}

func resolveType(f *protogen.Field, o *orm.FieldOptions) orm.Type {
	if o != nil && o.Type != orm.Type_TYPE_UNSPECIFIED {
		return o.Type
	}

	v := orm.Type_TYPE_UNSPECIFIED
	switch f.Desc.Kind() {
	case protoreflect.BoolKind:
		v = orm.Type_TYPE_BOOL
	case protoreflect.EnumKind:
		v = orm.Type_TYPE_ENUM
	case protoreflect.Int32Kind:
		v = orm.Type_TYPE_INT32
	case protoreflect.Sint32Kind:
		v = orm.Type_TYPE_SINT32
	case protoreflect.Sfixed32Kind:
		v = orm.Type_TYPE_SFIXED32
	case protoreflect.Uint32Kind:
		v = orm.Type_TYPE_UINT32
	case protoreflect.Fixed32Kind:
		v = orm.Type_TYPE_FIXED32
	case protoreflect.Int64Kind:
		v = orm.Type_TYPE_INT64
	case protoreflect.Sint64Kind:
		v = orm.Type_TYPE_SINT64
	case protoreflect.Sfixed64Kind:
		v = orm.Type_TYPE_SFIXED64
	case protoreflect.Uint64Kind:
		v = orm.Type_TYPE_UINT64
	case protoreflect.Fixed64Kind:
		v = orm.Type_TYPE_FIXED64
	case protoreflect.FloatKind:
		v = orm.Type_TYPE_FLOAT
	case protoreflect.DoubleKind:
		v = orm.Type_TYPE_DOUBLE
	case protoreflect.StringKind:
		v = orm.Type_TYPE_STRING
	case protoreflect.BytesKind:
		v = orm.Type_TYPE_BYTES
	case protoreflect.MessageKind:
		v = orm.Type_TYPE_MESSAGE
	case protoreflect.GroupKind:
		v = orm.Type_TYPE_GROUP

	default:
		panic("unknown proto type")
	}
	if v != orm.Type_TYPE_MESSAGE {
		return v
	}

	name := f.Message.Desc.FullName()
	if w, ok := scalar_mapped_known_messages[name]; ok {
		return w
	}

	return v
}
