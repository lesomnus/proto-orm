package pbgen

import (
	"fmt"

	"google.golang.org/protobuf/reflect/protoreflect"
)

type Type string

const (
	TypeDouble   Type = "double"
	TypeFloat    Type = "float"
	TypeInt32    Type = "int32"
	TypeInt64    Type = "int64"
	TypeUint32   Type = "uint32"
	TypeUint64   Type = "uint64"
	TypeSint32   Type = "sint32"
	TypeSint64   Type = "sint64"
	TypeFixed32  Type = "fixed32"
	TypeFixed64  Type = "fixed64"
	TypeSfixed32 Type = "sfixed32"
	TypeSfixed64 Type = "sfixed64"
	TypeBool     Type = "bool"
	TypeString   Type = "string"
	TypeBytes    Type = "bytes"

	TypeEmpty Type = "google.protobuf.Empty"
)

type Label string

const (
	LabelRequired Label = "required"
	LabelOptional Label = "optional"
	LabelRepeated Label = "repeated"
)

type Visibility string

const (
	VisibilityWeak   Visibility = "weak"
	VisibilityPublic Visibility = "public"
)

type Edition struct {
	Keyword string
	Value   string
}

var (
	SyntaxProto2 = Edition{"syntax", "proto2"}
	SyntaxProto3 = Edition{"syntax", "proto3"}
	Edition2023  = Edition{"edition", "2023"}
)

type ProtoFile struct {
	Edition Edition
	Package protoreflect.FullName
	Imports []Import
	Options []Option

	TopLevelDefinitions []TopLevelDef
}

func (ProtoFile) TemplateName() string {
	return "proto-file"
}

func (f *ProtoFile) AddImport(v Import) (int, bool) {
	for i, w := range f.Imports {
		if w.Name == v.Name {
			return i, false
		}
	}

	f.Imports = append(f.Imports, v)
	return len(f.Imports) - 1, true
}

type Import struct {
	Name       string
	Visibility Visibility
}

const (
	ImportWellKnownEmpty = "google/protobuf/empty.proto"
)

type Enum struct {
	Name    protoreflect.Name
	Options []Option
	Body    []EnumBody

	topLevelDef_
	messageBody_
}

func (Enum) TemplateName() string {
	return "enum"
}

type EnumBody interface{ enumBody() }
type enumBody_ struct{}

func (enumBody_) enumBody() {}

type EnumField struct {
	Name    string
	Number  int
	Options []Option

	enumBody_
}

func (EnumField) TemplateName() string {
	return "enum-field"
}

type Message struct {
	FullName protoreflect.FullName
	Body     []MessageBody

	topLevelDef_
	messageBody_
}

func (Message) TemplateName() string {
	return "message"
}

type MessageBody interface{ messageBody() }
type messageBody_ struct{}

func (messageBody_) messageBody() {}

type MessageField struct {
	Label   Label
	Type    Type
	Name    protoreflect.Name
	Number  int
	Options []Option

	messageBody_
}

func (MessageField) TemplateName() string {
	return "message-field"
}

type MessageOneof struct {
	Name    protoreflect.Name
	Options []Option
	Body    []MessageOneofBody

	messageBody_
}

func (MessageOneof) TemplateName() string {
	return "message-oneof"
}

type MessageOneofBody interface{ messageOneofBody() }
type messageOneofBody_ struct{}

func (messageOneofBody_) messageOneofBody() {}

type MessageOneofField struct {
	Type    Type
	Name    protoreflect.Name
	Number  int
	Options []Option

	messageOneofBody_
}

func (MessageOneofField) TemplateName() string {
	return "message-oneof-field"
}

type Service struct {
	Name protoreflect.Name
	Body []ServiceBody

	topLevelDef_
}

func (Service) TemplateName() string {
	return "service"
}

type ServiceBody interface{ serviceBody() }
type serviceBody_ struct{}

func (serviceBody_) serviceBody() {}

type Rpc struct {
	Name     protoreflect.Name
	Request  RpcType
	Response RpcType

	Options []Option

	serviceBody_
}

func (Rpc) TemplateName() string {
	return "rpc"
}

type RpcType struct {
	Type
	Stream bool
}

type OptionValue interface{ optionValue() }
type optionValue_ struct{}

func (optionValue_) optionValue() {}

type Option struct {
	Name  protoreflect.FullName
	Value OptionValue

	Standard bool
}

func (o *Option) NameString() string {
	if !o.Standard {
		return fmt.Sprintf("(%s)", o.Name)
	}
	return string(o.Name)
}

type Comment struct {
	Value     string
	Multiline bool

	topLevelDef_
	enumBody_
	messageBody_
	messageOneofBody_
	serviceBody_
}

func (Comment) TemplateName() string {
	return "comment"
}

type TopLevelDef interface{ topLevelDef() }
type topLevelDef_ struct{}

func (topLevelDef_) topLevelDef() {}

type Data struct {
	Fields []DataField

	dataValue_
}

func (Data) TemplateName() string {
	return "data"
}

type DataList struct {
	Values []Data

	dataValue_
}

func (DataList) TemplateName() string {
	return "data-list"
}

type DataField struct {
	Name  string
	Value DataValue
}

type DataValue interface{ dataValue() }
type dataValue_ struct {
	optionValue_
}

func (dataValue_) dataValue() {}

type DataString struct {
	Value string

	dataValue_
}

func (DataString) TemplateName() string {
	return "data-string"
}

type UnsafeLiteral struct {
	Value string

	dataValue_
}

func (UnsafeLiteral) TemplateName() string {
	return "unsafe-literal"
}
