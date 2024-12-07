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

type Entity struct {
	File   *protogen.File // Proto file where the message defines this entity is.
	Source *protogen.Message

	Key     *Field
	Fields  map[protoreflect.FullName]*Field
	Indexes []*Index
	Rpcs    map[RpcOp]*Rpc
}

func SetKey(e *Entity, k *Field) error {
	if e.Key != nil {
		return fmt.Errorf("%s: %s: key already exist at %s", k.Name(), "there can be only one key", e.Key.Name())
	}

	e.Key = k
	return nil
}

func NewEntity(f *protogen.File, m *protogen.Message, o *orm.MessageOptions) (*Entity, error) {
	v := &Entity{
		File:    f,
		Source:  m,
		Fields:  map[protoreflect.FullName]*Field{},
		Indexes: []*Index{},
	}

	errs := []error{}
	for _, f := range m.Fields {
		o := proto.GetExtension(f.Desc.Options(), orm.E_Field).(*orm.FieldOptions)
		w, err := NewField(m, f, o)
		if err != nil {
			err := fmt.Errorf("%s: %w", f.Desc.Name(), err)
			errs = append(errs, err)
			continue
		}
		if w.Ignored {
			continue
		}
		if w.Key {
			if err := SetKey(v, w); err != nil {
				errs = append(errs, err)
			}
		}

		w.Entity = v
		name := w.Desc.FullName()
		v.Fields[name] = w
	}
	for _, i := range o.Indexes {
		refs := []*Field{}
		for _, r := range i.Refs {
			f := v.FieldByNumber(protowire.Number(r))
			if f == nil {
				err := fmt.Errorf("index reference not found: %d", r)
				errs = append(errs, err)
				continue
			}

			refs = append(refs, f)
		}

		w := &Index{
			name:   i.Name,
			Entity: v,
			Refs:   refs,

			Key:       i.Key,
			Unique:    i.Unique,
			Immutable: i.Immutable,
		}

		v.Indexes = append(v.Indexes, w)
	}

	if len(errs) > 0 {
		return nil, errors.Join(errs...)
	}

	rpc_opts := orm.ResolveRpcOptions(f, m)
	r := NewRpcParser()
	v.Rpcs = r.Parse(v, rpc_opts)

	return v, nil
}

type Key interface {
	Name() string
}

func (e *Entity) KeyLikes() []Key {
	vs := []Key{e.Key}
	for _, v := range e.Fields {
		if v == e.Key {
			continue
		}
		if !v.Unique {
			continue
		}

		vs = append(vs, v)
	}
	for _, v := range e.Indexes {
		if !v.Unique {
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

func (e *Entity) HasScalarFields() bool {
	for v := range maps.Values(e.Fields) {
		if !v.IsEdge() {
			return true
		}
	}
	return false
}

func (e *Entity) HasEdges() bool {
	for v := range maps.Values(e.Fields) {
		if v.IsEdge() {
			return true
		}
	}
	return false
}

func (e *Entity) HasIndexes() bool {
	return len(e.Indexes) > 0
}

func (e *Entity) FieldByNumber(n protowire.Number) *Field {
	for _, f := range e.Fields {
		if f.Desc.Number() == n {
			return f
		}
	}
	return nil
}

func (e *Entity) FieldsSortByNumber() []*Field {
	fs := slices.Collect(maps.Values(e.Fields))
	slices.SortFunc(fs, func(a *Field, b *Field) int {
		return int(a.Desc.Number()) - int(b.Desc.Number())
	})

	return fs
}
