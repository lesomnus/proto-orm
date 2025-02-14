package graph

import (
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/encoding/protowire"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type Field interface {
	ProtoTyped

	Kind() protoreflect.Kind
	FullName() protoreflect.FullName
	Name() string
	Number() protowire.Number
	GoName() string

	Entity() *Entity

	IsIgnored() bool
	IsOptional() bool
	IsList() bool
	IsMap() bool
	IsUnique() bool
	IsImmutable() bool
	IsNullable() bool

	setImmutable()
	setNullable()
}

// Copy of protoreflect.FieldDescriptor.
// Do not implement any method that infer some value from another to keep consistency.
type FieldMeta struct {
	kind     protoreflect.Kind
	fullName protoreflect.FullName

	number protowire.Number
	goName string

	isList bool
	isMap  bool
}

func NewFieldMeta(f *protogen.Field) *FieldMeta {
	d := f.Desc
	return &FieldMeta{
		kind:     d.Kind(),
		fullName: d.FullName(),
		number:   d.Number(),
		goName:   f.GoName,
		isList:   d.IsList(),
		isMap:    d.IsMap(),
	}
}

func (m *FieldMeta) Kind() protoreflect.Kind {
	return m.kind
}

func (m *FieldMeta) FullName() protoreflect.FullName {
	return m.fullName
}

func (m *FieldMeta) Name() string {
	return string(m.fullName.Name())
}

func (m *FieldMeta) Number() protowire.Number {
	return m.number
}

func (m *FieldMeta) GoName() string {
	return m.goName
}

func (m *FieldMeta) IsList() bool {
	return m.isList
}

func (m *FieldMeta) IsMap() bool {
	return m.isMap
}
