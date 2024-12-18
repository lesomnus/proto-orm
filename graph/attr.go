package graph

import (
	"errors"
	"fmt"

	orm "github.com/lesomnus/proto-orm"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/encoding/protowire"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type Attr struct {
	source *protogen.Field
	entity *Entity // Message that defines this field.
	typ    orm.Type

	// Bound is not nil if this field is ID of target entity connected by edge.
	// Therefore, `Bound.IsEdge` always `true`.
	Bound   *Edge
	Default *string

	Ignored   bool
	Unique    bool
	Nullable  bool
	Immutable bool
}

func (e *Entity) parseAttr(f *protogen.Field, o *orm.FieldOptions) (*Attr, error) {
	if o == nil {
		o = &orm.FieldOptions{}
	}

	v := &Attr{
		source: f,
		entity: e,
		typ:    resolveType(f, o),

		Ignored:   o.Ignored,
		Unique:    o.Unique,
		Nullable:  o.Nullable,
		Immutable: o.Immutable,
		Default:   o.Default,
	}
	if v.typ == orm.Type_TYPE_MESSAGE {
		v.typ = orm.Type_TYPE_JSON
	}

	// d := f.Desc
	n := v.Name()
	e.Fields[n] = v
	e.Attrs[n] = v

	errs := []error{}
	if o.Nullable && !f.Desc.HasOptionalKeyword() {
		err := errors.New("nullable field must be optional")
		errs = append(errs, err)
	}
	if o.Key {
		d := f.Desc
		if d.IsList() || d.HasOptionalKeyword() {
			err := fmt.Errorf("key cannot have either repeated or optional cardinality")
			errs = append(errs, err)
		} else if e.Key != nil {
			err := fmt.Errorf("key already exist at %s", e.Key.Name())
			errs = append(errs, err)
		} else {
			e.Key = v
			v.Unique = true
			// It does not need to be immutable, but it is complex
			// to enable immutable key in the service generator.
			v.Immutable = true
		}
	}

	if len(errs) > 0 {
		return nil, errors.Join(errs...)
	}
	return v, nil
}

func (a *Attr) FullName() protoreflect.FullName {
	return a.source.Desc.FullName()
}

func (a *Attr) Name() string {
	return string(a.source.Desc.Name())
}

func (a *Attr) Number() protowire.Number {
	return a.source.Desc.Number()
}

func (a *Attr) Type() orm.Type {
	return a.typ
}

// `q` is qualifier that returns given `ident` as qualified name like: "Printf" -> "fmt.Printf".
// You may use `protogen.GeneratedFile.QualifiedGoIdent` in most of cases.
func (a *Attr) GoType(q func(ident protogen.GoIdent) string) string {
	if !a.IsList() {
		switch a.Type() {
		case orm.Type_TYPE_UUID:
			return q(protogen.GoImportPath("github.com/google/uuid").Ident("UUID"))
		case orm.Type_TYPE_TIME:
			return q(protogen.GoImportPath("time").Ident("Date"))
		}
	}

	return a.goType(a.source, q)
}

func (a *Attr) goType(f *protogen.Field, q func(ident protogen.GoIdent) string) string {
	d := f.Desc
	if d.IsMap() {
		k := f.Message.Fields[0]
		v := f.Message.Fields[1]
		return fmt.Sprintf("map[%s]%s", a.goType(k, q), a.goType(v, q))
	}

	v := ""
	k := d.Kind()
	switch k {
	case protoreflect.BoolKind:
		v = "bool"
	case protoreflect.EnumKind:
		v = q(f.Enum.GoIdent)
	case protoreflect.Uint32Kind,
		protoreflect.Fixed32Kind:
		v = "uint32"
	case protoreflect.Int32Kind,
		protoreflect.Sint32Kind,
		protoreflect.Sfixed32Kind:
		v = "int32"
	case protoreflect.Uint64Kind,
		protoreflect.Fixed64Kind:
		v = "uint64"
	case protoreflect.Int64Kind,
		protoreflect.Sint64Kind,
		protoreflect.Sfixed64Kind:
		v = "int64"
	case protoreflect.FloatKind:
		v = "float32"
	case protoreflect.DoubleKind:
		v = "float64"
	case protoreflect.StringKind:
		v = "string"
	case protoreflect.BytesKind:
		v = "[]byte"
	case protoreflect.MessageKind:
		v = "*" + q(f.Message.GoIdent)
	case protoreflect.GroupKind:
	default:
		panic(fmt.Sprintf("unknown type or type not supported: %s", k.String()))
	}

	if d.IsList() {
		return "[]" + v
	}
	return v
}

func (a *Attr) ProtoType() string {
	return a.protoType(a.source.Desc)
}

func (a *Attr) protoType(d protoreflect.FieldDescriptor) string {
	if d.IsMap() {
		return fmt.Sprintf("map<%s,%s>", a.protoType(d.MapKey()), a.protoType(d.MapValue()))
	}

	switch k := d.Kind(); k {
	case protoreflect.EnumKind:
		return string(d.Enum().FullName())
	case protoreflect.MessageKind:
		// Field type is another message.
		// e.g. google.protobuf.Timestamp, example.library.Book, ...
		return string(d.Message().FullName())
	default:
		// Field type is proto scalar.
		// e.g. string, int32, ...
		return k.String()
	}
}

func (a *Attr) ImportPath() string {
	if a.source.Message == nil {
		return ""
	}
	return a.source.Message.Desc.ParentFile().Path()
}

func (a *Attr) GoName() string {
	return a.source.GoName
}

func (a *Attr) Source() *protogen.Field {
	return a.source
}

func (a *Attr) Entity() *Entity {
	return a.entity
}

func (a *Attr) IsBound() bool {
	return a.Bound != nil
}

func (a *Attr) HasDefault() bool {
	return !a.typ.IsPrimitive() && a.Default != nil
}

func (a *Attr) IsJson() bool {
	return a.typ == orm.Type_TYPE_JSON
}

func (a *Attr) IsIgnored() bool {
	return a.Ignored
}

func (a *Attr) IsOptional() bool {
	d := a.source.Desc
	return d.HasOptionalKeyword() || d.IsList() || d.IsMap()
}

// The type of proto field implies optional.
// TODO: better name?
func (a *Attr) IsSoft() bool {
	return a.IsList() || a.IsMap()
}

func (a *Attr) IsList() bool {
	return a.source.Desc.IsList()
}

func (a *Attr) IsMap() bool {
	return a.source.Desc.IsMap()
}

func (a *Attr) IsUnique() bool {
	return a.Unique
}

func (a *Attr) IsNullable() bool {
	return a.Nullable
}

func (a *Attr) IsImmutable() bool {
	return a.Immutable
}

func (a *Attr) setNullable() {
	a.Nullable = true
}

func (a *Attr) setImmutable() {
	a.Immutable = true
}
