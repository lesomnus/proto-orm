// Code generated by "protoc-gen-orm-service", DO NOT EDIT.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.0
// source: example/library/press_svc.proto

package library

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PressAddRequest struct {
	state                   protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Id           []byte                 `protobuf:"bytes,1,opt,name=id,proto3,oneof"`
	xxx_hidden_Book         *BookGetRequest        `protobuf:"bytes,2,opt,name=book,proto3"`
	xxx_hidden_SerialNumber string                 `protobuf:"bytes,3,opt,name=serial_number,json=serialNumber,proto3"`
	xxx_hidden_DateCreated  *timestamppb.Timestamp `protobuf:"bytes,15,opt,name=date_created,json=dateCreated,proto3,oneof"`
	XXX_raceDetectHookData  protoimpl.RaceDetectHookData
	XXX_presence            [1]uint32
	unknownFields           protoimpl.UnknownFields
	sizeCache               protoimpl.SizeCache
}

func (x *PressAddRequest) Reset() {
	*x = PressAddRequest{}
	mi := &file_example_library_press_svc_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PressAddRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PressAddRequest) ProtoMessage() {}

func (x *PressAddRequest) ProtoReflect() protoreflect.Message {
	mi := &file_example_library_press_svc_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *PressAddRequest) GetId() []byte {
	if x != nil {
		return x.xxx_hidden_Id
	}
	return nil
}

func (x *PressAddRequest) GetBook() *BookGetRequest {
	if x != nil {
		return x.xxx_hidden_Book
	}
	return nil
}

func (x *PressAddRequest) GetSerialNumber() string {
	if x != nil {
		return x.xxx_hidden_SerialNumber
	}
	return ""
}

func (x *PressAddRequest) GetDateCreated() *timestamppb.Timestamp {
	if x != nil {
		return x.xxx_hidden_DateCreated
	}
	return nil
}

func (x *PressAddRequest) SetId(v []byte) {
	if v == nil {
		v = []byte{}
	}
	x.xxx_hidden_Id = v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 0, 4)
}

func (x *PressAddRequest) SetBook(v *BookGetRequest) {
	x.xxx_hidden_Book = v
}

func (x *PressAddRequest) SetSerialNumber(v string) {
	x.xxx_hidden_SerialNumber = v
}

func (x *PressAddRequest) SetDateCreated(v *timestamppb.Timestamp) {
	x.xxx_hidden_DateCreated = v
}

func (x *PressAddRequest) HasId() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 0)
}

func (x *PressAddRequest) HasBook() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_Book != nil
}

func (x *PressAddRequest) HasDateCreated() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_DateCreated != nil
}

func (x *PressAddRequest) ClearId() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 0)
	x.xxx_hidden_Id = nil
}

func (x *PressAddRequest) ClearBook() {
	x.xxx_hidden_Book = nil
}

func (x *PressAddRequest) ClearDateCreated() {
	x.xxx_hidden_DateCreated = nil
}

type PressAddRequest_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Id           []byte
	Book         *BookGetRequest
	SerialNumber string
	DateCreated  *timestamppb.Timestamp
}

func (b0 PressAddRequest_builder) Build() *PressAddRequest {
	m0 := &PressAddRequest{}
	b, x := &b0, m0
	_, _ = b, x
	if b.Id != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 0, 4)
		x.xxx_hidden_Id = b.Id
	}
	x.xxx_hidden_Book = b.Book
	x.xxx_hidden_SerialNumber = b.SerialNumber
	x.xxx_hidden_DateCreated = b.DateCreated
	return m0
}

type PressGetRequest struct {
	state             protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Select *PressSelect           `protobuf:"bytes,1,opt,name=select,proto3"`
	xxx_hidden_Key    isPressGetRequest_Key  `protobuf_oneof:"key"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *PressGetRequest) Reset() {
	*x = PressGetRequest{}
	mi := &file_example_library_press_svc_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PressGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PressGetRequest) ProtoMessage() {}

func (x *PressGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_example_library_press_svc_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *PressGetRequest) GetSelect() *PressSelect {
	if x != nil {
		return x.xxx_hidden_Select
	}
	return nil
}

func (x *PressGetRequest) GetId() []byte {
	if x != nil {
		if x, ok := x.xxx_hidden_Key.(*pressGetRequest_Id); ok {
			return x.Id
		}
	}
	return nil
}

func (x *PressGetRequest) GetSerialNumber() *PressGetBySerialNumber {
	if x != nil {
		if x, ok := x.xxx_hidden_Key.(*pressGetRequest_SerialNumber); ok {
			return x.SerialNumber
		}
	}
	return nil
}

func (x *PressGetRequest) SetSelect(v *PressSelect) {
	x.xxx_hidden_Select = v
}

func (x *PressGetRequest) SetId(v []byte) {
	if v == nil {
		v = []byte{}
	}
	x.xxx_hidden_Key = &pressGetRequest_Id{v}
}

func (x *PressGetRequest) SetSerialNumber(v *PressGetBySerialNumber) {
	if v == nil {
		x.xxx_hidden_Key = nil
		return
	}
	x.xxx_hidden_Key = &pressGetRequest_SerialNumber{v}
}

func (x *PressGetRequest) HasSelect() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_Select != nil
}

func (x *PressGetRequest) HasKey() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_Key != nil
}

func (x *PressGetRequest) HasId() bool {
	if x == nil {
		return false
	}
	_, ok := x.xxx_hidden_Key.(*pressGetRequest_Id)
	return ok
}

func (x *PressGetRequest) HasSerialNumber() bool {
	if x == nil {
		return false
	}
	_, ok := x.xxx_hidden_Key.(*pressGetRequest_SerialNumber)
	return ok
}

func (x *PressGetRequest) ClearSelect() {
	x.xxx_hidden_Select = nil
}

func (x *PressGetRequest) ClearKey() {
	x.xxx_hidden_Key = nil
}

func (x *PressGetRequest) ClearId() {
	if _, ok := x.xxx_hidden_Key.(*pressGetRequest_Id); ok {
		x.xxx_hidden_Key = nil
	}
}

func (x *PressGetRequest) ClearSerialNumber() {
	if _, ok := x.xxx_hidden_Key.(*pressGetRequest_SerialNumber); ok {
		x.xxx_hidden_Key = nil
	}
}

const PressGetRequest_Key_not_set_case case_PressGetRequest_Key = 0
const PressGetRequest_Id_case case_PressGetRequest_Key = 2
const PressGetRequest_SerialNumber_case case_PressGetRequest_Key = 3

func (x *PressGetRequest) WhichKey() case_PressGetRequest_Key {
	if x == nil {
		return PressGetRequest_Key_not_set_case
	}
	switch x.xxx_hidden_Key.(type) {
	case *pressGetRequest_Id:
		return PressGetRequest_Id_case
	case *pressGetRequest_SerialNumber:
		return PressGetRequest_SerialNumber_case
	default:
		return PressGetRequest_Key_not_set_case
	}
}

type PressGetRequest_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Select *PressSelect
	// Fields of oneof xxx_hidden_Key:
	Id           []byte
	SerialNumber *PressGetBySerialNumber
	// -- end of xxx_hidden_Key
}

func (b0 PressGetRequest_builder) Build() *PressGetRequest {
	m0 := &PressGetRequest{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_Select = b.Select
	if b.Id != nil {
		x.xxx_hidden_Key = &pressGetRequest_Id{b.Id}
	}
	if b.SerialNumber != nil {
		x.xxx_hidden_Key = &pressGetRequest_SerialNumber{b.SerialNumber}
	}
	return m0
}

type case_PressGetRequest_Key protoreflect.FieldNumber

func (x case_PressGetRequest_Key) String() string {
	md := file_example_library_press_svc_proto_msgTypes[1].Descriptor()
	if x == 0 {
		return "not set"
	}
	return protoimpl.X.MessageFieldStringOf(md, protoreflect.FieldNumber(x))
}

type isPressGetRequest_Key interface {
	isPressGetRequest_Key()
}

type pressGetRequest_Id struct {
	Id []byte `protobuf:"bytes,2,opt,name=id,proto3,oneof"`
}

type pressGetRequest_SerialNumber struct {
	SerialNumber *PressGetBySerialNumber `protobuf:"bytes,3,opt,name=serial_number,json=serialNumber,proto3,oneof"`
}

func (*pressGetRequest_Id) isPressGetRequest_Key() {}

func (*pressGetRequest_SerialNumber) isPressGetRequest_Key() {}

type PressSelect struct {
	state                   protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_All          bool                   `protobuf:"varint,1,opt,name=all,proto3,oneof"`
	xxx_hidden_Book         *BookSelect            `protobuf:"bytes,2,opt,name=book,proto3,oneof"`
	xxx_hidden_SerialNumber bool                   `protobuf:"varint,3,opt,name=serial_number,json=serialNumber,proto3,oneof"`
	xxx_hidden_DateCreated  bool                   `protobuf:"varint,15,opt,name=date_created,json=dateCreated,proto3,oneof"`
	XXX_raceDetectHookData  protoimpl.RaceDetectHookData
	XXX_presence            [1]uint32
	unknownFields           protoimpl.UnknownFields
	sizeCache               protoimpl.SizeCache
}

func (x *PressSelect) Reset() {
	*x = PressSelect{}
	mi := &file_example_library_press_svc_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PressSelect) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PressSelect) ProtoMessage() {}

func (x *PressSelect) ProtoReflect() protoreflect.Message {
	mi := &file_example_library_press_svc_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *PressSelect) GetAll() bool {
	if x != nil {
		return x.xxx_hidden_All
	}
	return false
}

func (x *PressSelect) GetBook() *BookSelect {
	if x != nil {
		return x.xxx_hidden_Book
	}
	return nil
}

func (x *PressSelect) GetSerialNumber() bool {
	if x != nil {
		return x.xxx_hidden_SerialNumber
	}
	return false
}

func (x *PressSelect) GetDateCreated() bool {
	if x != nil {
		return x.xxx_hidden_DateCreated
	}
	return false
}

func (x *PressSelect) SetAll(v bool) {
	x.xxx_hidden_All = v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 0, 4)
}

func (x *PressSelect) SetBook(v *BookSelect) {
	x.xxx_hidden_Book = v
}

func (x *PressSelect) SetSerialNumber(v bool) {
	x.xxx_hidden_SerialNumber = v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 2, 4)
}

func (x *PressSelect) SetDateCreated(v bool) {
	x.xxx_hidden_DateCreated = v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 3, 4)
}

func (x *PressSelect) HasAll() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 0)
}

func (x *PressSelect) HasBook() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_Book != nil
}

func (x *PressSelect) HasSerialNumber() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 2)
}

func (x *PressSelect) HasDateCreated() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 3)
}

func (x *PressSelect) ClearAll() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 0)
	x.xxx_hidden_All = false
}

func (x *PressSelect) ClearBook() {
	x.xxx_hidden_Book = nil
}

func (x *PressSelect) ClearSerialNumber() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 2)
	x.xxx_hidden_SerialNumber = false
}

func (x *PressSelect) ClearDateCreated() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 3)
	x.xxx_hidden_DateCreated = false
}

type PressSelect_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	All          *bool
	Book         *BookSelect
	SerialNumber *bool
	DateCreated  *bool
}

func (b0 PressSelect_builder) Build() *PressSelect {
	m0 := &PressSelect{}
	b, x := &b0, m0
	_, _ = b, x
	if b.All != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 0, 4)
		x.xxx_hidden_All = *b.All
	}
	x.xxx_hidden_Book = b.Book
	if b.SerialNumber != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 2, 4)
		x.xxx_hidden_SerialNumber = *b.SerialNumber
	}
	if b.DateCreated != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 3, 4)
		x.xxx_hidden_DateCreated = *b.DateCreated
	}
	return m0
}

type PressGetBySerialNumber struct {
	state                   protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Book         *BookGetRequest        `protobuf:"bytes,2,opt,name=book,proto3"`
	xxx_hidden_SerialNumber string                 `protobuf:"bytes,3,opt,name=serial_number,json=serialNumber,proto3"`
	unknownFields           protoimpl.UnknownFields
	sizeCache               protoimpl.SizeCache
}

func (x *PressGetBySerialNumber) Reset() {
	*x = PressGetBySerialNumber{}
	mi := &file_example_library_press_svc_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PressGetBySerialNumber) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PressGetBySerialNumber) ProtoMessage() {}

func (x *PressGetBySerialNumber) ProtoReflect() protoreflect.Message {
	mi := &file_example_library_press_svc_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *PressGetBySerialNumber) GetBook() *BookGetRequest {
	if x != nil {
		return x.xxx_hidden_Book
	}
	return nil
}

func (x *PressGetBySerialNumber) GetSerialNumber() string {
	if x != nil {
		return x.xxx_hidden_SerialNumber
	}
	return ""
}

func (x *PressGetBySerialNumber) SetBook(v *BookGetRequest) {
	x.xxx_hidden_Book = v
}

func (x *PressGetBySerialNumber) SetSerialNumber(v string) {
	x.xxx_hidden_SerialNumber = v
}

func (x *PressGetBySerialNumber) HasBook() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_Book != nil
}

func (x *PressGetBySerialNumber) ClearBook() {
	x.xxx_hidden_Book = nil
}

type PressGetBySerialNumber_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Book         *BookGetRequest
	SerialNumber string
}

func (b0 PressGetBySerialNumber_builder) Build() *PressGetBySerialNumber {
	m0 := &PressGetBySerialNumber{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_Book = b.Book
	x.xxx_hidden_SerialNumber = b.SerialNumber
	return m0
}

type PressPatchRequest struct {
	state          protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Key *PressGetRequest       `protobuf:"bytes,1,opt,name=key,proto3"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *PressPatchRequest) Reset() {
	*x = PressPatchRequest{}
	mi := &file_example_library_press_svc_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PressPatchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PressPatchRequest) ProtoMessage() {}

func (x *PressPatchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_example_library_press_svc_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *PressPatchRequest) GetKey() *PressGetRequest {
	if x != nil {
		return x.xxx_hidden_Key
	}
	return nil
}

func (x *PressPatchRequest) SetKey(v *PressGetRequest) {
	x.xxx_hidden_Key = v
}

func (x *PressPatchRequest) HasKey() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_Key != nil
}

func (x *PressPatchRequest) ClearKey() {
	x.xxx_hidden_Key = nil
}

type PressPatchRequest_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Key *PressGetRequest
}

func (b0 PressPatchRequest_builder) Build() *PressPatchRequest {
	m0 := &PressPatchRequest{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_Key = b.Key
	return m0
}

var File_example_library_press_svc_proto protoreflect.FileDescriptor

const file_example_library_press_svc_proto_rawDesc = "" +
	"\n" +
	"\x1fexample/library/press_svc.proto\x12\x0fexample.library\x1a\x1eexample/library/book_svc.proto\x1a\x1bexample/library/press.proto\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1fgoogle/protobuf/timestamp.proto\"\xdc\x01\n" +
	"\x0fPressAddRequest\x12\x13\n" +
	"\x02id\x18\x01 \x01(\fH\x00R\x02id\x88\x01\x01\x123\n" +
	"\x04book\x18\x02 \x01(\v2\x1f.example.library.BookGetRequestR\x04book\x12#\n" +
	"\rserial_number\x18\x03 \x01(\tR\fserialNumber\x12B\n" +
	"\fdate_created\x18\x0f \x01(\v2\x1a.google.protobuf.TimestampH\x01R\vdateCreated\x88\x01\x01B\x05\n" +
	"\x03_idB\x0f\n" +
	"\r_date_created\"\xb0\x01\n" +
	"\x0fPressGetRequest\x124\n" +
	"\x06select\x18\x01 \x01(\v2\x1c.example.library.PressSelectR\x06select\x12\x10\n" +
	"\x02id\x18\x02 \x01(\fH\x00R\x02id\x12N\n" +
	"\rserial_number\x18\x03 \x01(\v2'.example.library.PressGetBySerialNumberH\x00R\fserialNumberB\x05\n" +
	"\x03key\"\xe0\x01\n" +
	"\vPressSelect\x12\x15\n" +
	"\x03all\x18\x01 \x01(\bH\x00R\x03all\x88\x01\x01\x124\n" +
	"\x04book\x18\x02 \x01(\v2\x1b.example.library.BookSelectH\x01R\x04book\x88\x01\x01\x12(\n" +
	"\rserial_number\x18\x03 \x01(\bH\x02R\fserialNumber\x88\x01\x01\x12&\n" +
	"\fdate_created\x18\x0f \x01(\bH\x03R\vdateCreated\x88\x01\x01B\x06\n" +
	"\x04_allB\a\n" +
	"\x05_bookB\x10\n" +
	"\x0e_serial_numberB\x0f\n" +
	"\r_date_created\"r\n" +
	"\x16PressGetBySerialNumber\x123\n" +
	"\x04book\x18\x02 \x01(\v2\x1f.example.library.BookGetRequestR\x04book\x12#\n" +
	"\rserial_number\x18\x03 \x01(\tR\fserialNumber\"G\n" +
	"\x11PressPatchRequest\x122\n" +
	"\x03key\x18\x01 \x01(\v2 .example.library.PressGetRequestR\x03key2\x98\x02\n" +
	"\fPressService\x12?\n" +
	"\x03Add\x12 .example.library.PressAddRequest\x1a\x16.example.library.Press\x12?\n" +
	"\x03Get\x12 .example.library.PressGetRequest\x1a\x16.example.library.Press\x12C\n" +
	"\x05Patch\x12\".example.library.PressPatchRequest\x1a\x16.google.protobuf.Empty\x12A\n" +
	"\x05Erase\x12 .example.library.PressGetRequest\x1a\x16.google.protobuf.EmptyB8Z6github.com/lesomnus/proto-orm/internal/example/libraryb\x06proto3"

var file_example_library_press_svc_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_example_library_press_svc_proto_goTypes = []any{
	(*PressAddRequest)(nil),        // 0: example.library.PressAddRequest
	(*PressGetRequest)(nil),        // 1: example.library.PressGetRequest
	(*PressSelect)(nil),            // 2: example.library.PressSelect
	(*PressGetBySerialNumber)(nil), // 3: example.library.PressGetBySerialNumber
	(*PressPatchRequest)(nil),      // 4: example.library.PressPatchRequest
	(*BookGetRequest)(nil),         // 5: example.library.BookGetRequest
	(*timestamppb.Timestamp)(nil),  // 6: google.protobuf.Timestamp
	(*BookSelect)(nil),             // 7: example.library.BookSelect
	(*Press)(nil),                  // 8: example.library.Press
	(*emptypb.Empty)(nil),          // 9: google.protobuf.Empty
}
var file_example_library_press_svc_proto_depIdxs = []int32{
	5,  // 0: example.library.PressAddRequest.book:type_name -> example.library.BookGetRequest
	6,  // 1: example.library.PressAddRequest.date_created:type_name -> google.protobuf.Timestamp
	2,  // 2: example.library.PressGetRequest.select:type_name -> example.library.PressSelect
	3,  // 3: example.library.PressGetRequest.serial_number:type_name -> example.library.PressGetBySerialNumber
	7,  // 4: example.library.PressSelect.book:type_name -> example.library.BookSelect
	5,  // 5: example.library.PressGetBySerialNumber.book:type_name -> example.library.BookGetRequest
	1,  // 6: example.library.PressPatchRequest.key:type_name -> example.library.PressGetRequest
	0,  // 7: example.library.PressService.Add:input_type -> example.library.PressAddRequest
	1,  // 8: example.library.PressService.Get:input_type -> example.library.PressGetRequest
	4,  // 9: example.library.PressService.Patch:input_type -> example.library.PressPatchRequest
	1,  // 10: example.library.PressService.Erase:input_type -> example.library.PressGetRequest
	8,  // 11: example.library.PressService.Add:output_type -> example.library.Press
	8,  // 12: example.library.PressService.Get:output_type -> example.library.Press
	9,  // 13: example.library.PressService.Patch:output_type -> google.protobuf.Empty
	9,  // 14: example.library.PressService.Erase:output_type -> google.protobuf.Empty
	11, // [11:15] is the sub-list for method output_type
	7,  // [7:11] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_example_library_press_svc_proto_init() }
func file_example_library_press_svc_proto_init() {
	if File_example_library_press_svc_proto != nil {
		return
	}
	file_example_library_book_svc_proto_init()
	file_example_library_press_proto_init()
	file_example_library_press_svc_proto_msgTypes[0].OneofWrappers = []any{}
	file_example_library_press_svc_proto_msgTypes[1].OneofWrappers = []any{
		(*pressGetRequest_Id)(nil),
		(*pressGetRequest_SerialNumber)(nil),
	}
	file_example_library_press_svc_proto_msgTypes[2].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_example_library_press_svc_proto_rawDesc), len(file_example_library_press_svc_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_example_library_press_svc_proto_goTypes,
		DependencyIndexes: file_example_library_press_svc_proto_depIdxs,
		MessageInfos:      file_example_library_press_svc_proto_msgTypes,
	}.Build()
	File_example_library_press_svc_proto = out.File
	file_example_library_press_svc_proto_goTypes = nil
	file_example_library_press_svc_proto_depIdxs = nil
}
