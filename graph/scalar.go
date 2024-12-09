package graph

import (
	"errors"
	"fmt"

	orm "github.com/lesomnus/proto-orm"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/encoding/protowire"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type ScalarField struct {
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

func (e *Entity) parseScalarField(f *protogen.Field, o *orm.FieldOptions) (*ScalarField, error) {
	if o == nil {
		o = &orm.FieldOptions{}
	}

	v := &ScalarField{
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
	e.Scalars[n] = v

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

func (f *ScalarField) FullName() protoreflect.FullName {
	return f.source.Desc.FullName()
}

func (f *ScalarField) Name() string {
	return string(f.source.Desc.Name())
}

func (f *ScalarField) Number() protowire.Number {
	return f.source.Desc.Number()
}

func (f *ScalarField) Type() orm.Type {
	return f.typ
}

func (f *ScalarField) ProtoType() string {
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

func (f *ScalarField) GoName() string {
	return f.source.GoName
}

func (f *ScalarField) Source() *protogen.Field {
	return f.source
}

func (f *ScalarField) Entity() *Entity {
	return f.entity
}

func (f *ScalarField) IsBound() bool {
	return f.Bound != nil
}

func (f *ScalarField) HasDefault() bool {
	return f.Default != nil
}

func (f *ScalarField) IsIgnored() bool {
	return f.Ignored
}

func (f *ScalarField) IsOptional() bool {
	return f.source.Desc.HasOptionalKeyword()
}

func (f *ScalarField) IsList() bool {
	return f.source.Desc.IsList()
}

func (f *ScalarField) IsUnique() bool {
	return f.Unique
}

func (f *ScalarField) IsNullable() bool {
	return f.Nullable
}

func (f *ScalarField) IsImmutable() bool {
	return f.Immutable
}

func (f *ScalarField) setNullable() {
	f.Nullable = true
}

func (f *ScalarField) setImmutable() {
	f.Immutable = true
}
