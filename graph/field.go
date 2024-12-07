package graph

import (
	"errors"

	orm "github.com/lesomnus/proto-orm"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type Field struct {
	*protogen.Field
	Entity *Entity // Message that defines this field.

	GivenType orm.Type

	Edge *Edge
	// Bound is not nil if this field is ID of target entity connected by edge.
	// Therefore, `Bound.IsEdge` always `true`.
	Bound *Field

	Ignored   bool
	Key       bool
	Unique    bool
	Required  bool
	Nullable  bool
	Immutable bool

	Default *string
}

func NewField(m *protogen.Message, f *protogen.Field, o *orm.FieldOptions) (*Field, error) {
	if o == nil {
		o = &orm.FieldOptions{}
	}

	v := &Field{
		Field: f,

		GivenType: o.Type,
		Ignored:   o.Ignored,
		Key:       o.Key,
		Unique:    o.Unique,
		Required:  !f.Desc.HasOptionalKeyword(),
		Nullable:  o.Nullable,
		Immutable: o.Immutable,
		Default:   o.Default,
	}
	if f.Desc.IsList() {
		v.Required = false
		v.Nullable = false
	}
	if v.Nullable {
		v.Required = false
	}
	if v.Key {
		if !v.Required {
			return nil, errors.New("key field cannot be optional or nullable")
		}

		v.Required = true
		v.Unique = true
		// It does not need to be immutable, but it is complex
		// to enable immutable key in the service generator.
		v.Immutable = true
	}

	return v, nil
}

func (f *Field) IsScalar() bool {
	return f.Message == nil
}

func (f *Field) Name() string {
	return string(f.Field.Desc.Name())
}

func (f *Field) IsEdge() bool {
	return f.resolveType() == orm.Type_TYPE_UNSPECIFIED
}

func (f *Field) HasDefault() bool {
	return f.Default != nil
}

func (f *Field) Type() orm.Type {
	return f.resolveType()
}

func (f *Field) IsBound() bool {
	return f.Bound != nil
}

func (f *Field) IsList() bool {
	if f.Desc.IsList() {
		return true
	}
	if f.IsEdge() {
		return f.Unique
	}
	return false
}

var scalar_mapped_known_messages = map[protoreflect.FullName]orm.Type{
	"google.protobuf.Timestamp": orm.Type_TYPE_TIME,
}

func (f *Field) resolveType() orm.Type {
	if f.GivenType != orm.Type_TYPE_UNSPECIFIED {
		return f.GivenType
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
	if v, ok := scalar_mapped_known_messages[name]; ok {
		return v
	}

	return orm.Type_TYPE_UNSPECIFIED
}
