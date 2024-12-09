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

	n := v.Name()
	e.Fields[n] = v
	e.Attrs[n] = v

	errs := []error{}
	if o.Nullable && !f.Desc.HasOptionalKeyword() {
		err := errors.New("nullable field must be optional")
		errs = append(errs, err)
	}
	if o.Key {
		if e.Key != nil {
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

func (f *Attr) FullName() protoreflect.FullName {
	return f.source.Desc.FullName()
}

func (f *Attr) Name() string {
	return string(f.source.Desc.Name())
}

func (f *Attr) Number() protowire.Number {
	return f.source.Desc.Number()
}

func (f *Attr) Type() orm.Type {
	return f.typ
}

func (f *Attr) ProtoType() string {
	if f.source.Message == nil {
		// Field type is proto scalar.
		// e.g. string, int32, ...
		return f.source.Desc.Kind().String()
	} else {
		// Field type is well known type that can be mapped to scalar.
		// e.g. google.protobuf.Timestamp -> time.Time
		return string(f.source.Message.Desc.FullName())
	}
}

func (f *Attr) GoName() string {
	return f.source.GoName
}

func (f *Attr) Source() *protogen.Field {
	return f.source
}

func (f *Attr) Entity() *Entity {
	return f.entity
}

func (f *Attr) IsBound() bool {
	return f.Bound != nil
}

func (f *Attr) HasDefault() bool {
	return f.Default != nil
}

func (f *Attr) IsIgnored() bool {
	return f.Ignored
}

func (f *Attr) IsOptional() bool {
	return f.source.Desc.HasOptionalKeyword()
}

func (f *Attr) IsList() bool {
	return f.source.Desc.IsList()
}

func (f *Attr) IsUnique() bool {
	return f.Unique
}

func (f *Attr) IsNullable() bool {
	return f.Nullable
}

func (f *Attr) IsImmutable() bool {
	return f.Immutable
}

func (f *Attr) setNullable() {
	f.Nullable = true
}

func (f *Attr) setImmutable() {
	f.Immutable = true
}
