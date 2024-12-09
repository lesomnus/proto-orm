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

func (a *Attr) ProtoType() string {
	if a.source.Message == nil {
		// Field type is proto scalar.
		// e.g. string, int32, ...
		return a.source.Desc.Kind().String()
	} else {
		// Field type is well known type that can be mapped to scalar.
		// e.g. google.protobuf.Timestamp -> time.Time
		return string(a.source.Message.Desc.FullName())
	}
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
	return a.Default != nil
}

func (a *Attr) IsJson() bool {
	return a.typ == orm.Type_TYPE_JSON
}

func (a *Attr) IsIgnored() bool {
	return a.Ignored
}

func (a *Attr) IsOptional() bool {
	d := a.source.Desc
	return d.HasOptionalKeyword() || d.IsList()
}

func (a *Attr) IsList() bool {
	return a.source.Desc.IsList()
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
