// Code generated by "protoc-gen-orm-service", DO NOT EDIT.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        v5.29.0
// source: example/library/book.svc.proto

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

type BookAddRequest struct {
	state                  protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Id          []byte                 `protobuf:"bytes,1,opt,name=id,proto3,oneof"`
	xxx_hidden_Title       string                 `protobuf:"bytes,3,opt,name=title,proto3"`
	xxx_hidden_Index       *[]*Book_Index         `protobuf:"bytes,5,rep,name=index,proto3"`
	xxx_hidden_Genres      []string               `protobuf:"bytes,6,rep,name=genres,proto3"`
	xxx_hidden_DateCreated *timestamppb.Timestamp `protobuf:"bytes,15,opt,name=date_created,json=dateCreated,proto3,oneof"`
	XXX_raceDetectHookData protoimpl.RaceDetectHookData
	XXX_presence           [1]uint32
	unknownFields          protoimpl.UnknownFields
	sizeCache              protoimpl.SizeCache
}

func (x *BookAddRequest) Reset() {
	*x = BookAddRequest{}
	mi := &file_example_library_book_svc_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BookAddRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookAddRequest) ProtoMessage() {}

func (x *BookAddRequest) ProtoReflect() protoreflect.Message {
	mi := &file_example_library_book_svc_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *BookAddRequest) GetId() []byte {
	if x != nil {
		return x.xxx_hidden_Id
	}
	return nil
}

func (x *BookAddRequest) GetTitle() string {
	if x != nil {
		return x.xxx_hidden_Title
	}
	return ""
}

func (x *BookAddRequest) GetIndex() []*Book_Index {
	if x != nil {
		if x.xxx_hidden_Index != nil {
			return *x.xxx_hidden_Index
		}
	}
	return nil
}

func (x *BookAddRequest) GetGenres() []string {
	if x != nil {
		return x.xxx_hidden_Genres
	}
	return nil
}

func (x *BookAddRequest) GetDateCreated() *timestamppb.Timestamp {
	if x != nil {
		return x.xxx_hidden_DateCreated
	}
	return nil
}

func (x *BookAddRequest) SetId(v []byte) {
	if v == nil {
		v = []byte{}
	}
	x.xxx_hidden_Id = v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 0, 5)
}

func (x *BookAddRequest) SetTitle(v string) {
	x.xxx_hidden_Title = v
}

func (x *BookAddRequest) SetIndex(v []*Book_Index) {
	x.xxx_hidden_Index = &v
}

func (x *BookAddRequest) SetGenres(v []string) {
	x.xxx_hidden_Genres = v
}

func (x *BookAddRequest) SetDateCreated(v *timestamppb.Timestamp) {
	x.xxx_hidden_DateCreated = v
}

func (x *BookAddRequest) HasId() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 0)
}

func (x *BookAddRequest) HasDateCreated() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_DateCreated != nil
}

func (x *BookAddRequest) ClearId() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 0)
	x.xxx_hidden_Id = nil
}

func (x *BookAddRequest) ClearDateCreated() {
	x.xxx_hidden_DateCreated = nil
}

type BookAddRequest_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Id          []byte
	Title       string
	Index       []*Book_Index
	Genres      []string
	DateCreated *timestamppb.Timestamp
}

func (b0 BookAddRequest_builder) Build() *BookAddRequest {
	m0 := &BookAddRequest{}
	b, x := &b0, m0
	_, _ = b, x
	if b.Id != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 0, 5)
		x.xxx_hidden_Id = b.Id
	}
	x.xxx_hidden_Title = b.Title
	x.xxx_hidden_Index = &b.Index
	x.xxx_hidden_Genres = b.Genres
	x.xxx_hidden_DateCreated = b.DateCreated
	return m0
}

type BookGetRequest struct {
	state             protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Select *BookSelect            `protobuf:"bytes,1,opt,name=select,proto3"`
	xxx_hidden_Key    isBookGetRequest_Key   `protobuf_oneof:"key"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *BookGetRequest) Reset() {
	*x = BookGetRequest{}
	mi := &file_example_library_book_svc_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BookGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookGetRequest) ProtoMessage() {}

func (x *BookGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_example_library_book_svc_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *BookGetRequest) GetSelect() *BookSelect {
	if x != nil {
		return x.xxx_hidden_Select
	}
	return nil
}

func (x *BookGetRequest) GetId() []byte {
	if x != nil {
		if x, ok := x.xxx_hidden_Key.(*bookGetRequest_Id); ok {
			return x.Id
		}
	}
	return nil
}

func (x *BookGetRequest) SetSelect(v *BookSelect) {
	x.xxx_hidden_Select = v
}

func (x *BookGetRequest) SetId(v []byte) {
	if v == nil {
		v = []byte{}
	}
	x.xxx_hidden_Key = &bookGetRequest_Id{v}
}

func (x *BookGetRequest) HasSelect() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_Select != nil
}

func (x *BookGetRequest) HasKey() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_Key != nil
}

func (x *BookGetRequest) HasId() bool {
	if x == nil {
		return false
	}
	_, ok := x.xxx_hidden_Key.(*bookGetRequest_Id)
	return ok
}

func (x *BookGetRequest) ClearSelect() {
	x.xxx_hidden_Select = nil
}

func (x *BookGetRequest) ClearKey() {
	x.xxx_hidden_Key = nil
}

func (x *BookGetRequest) ClearId() {
	if _, ok := x.xxx_hidden_Key.(*bookGetRequest_Id); ok {
		x.xxx_hidden_Key = nil
	}
}

const BookGetRequest_Key_not_set_case case_BookGetRequest_Key = 0
const BookGetRequest_Id_case case_BookGetRequest_Key = 2

func (x *BookGetRequest) WhichKey() case_BookGetRequest_Key {
	if x == nil {
		return BookGetRequest_Key_not_set_case
	}
	switch x.xxx_hidden_Key.(type) {
	case *bookGetRequest_Id:
		return BookGetRequest_Id_case
	default:
		return BookGetRequest_Key_not_set_case
	}
}

type BookGetRequest_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Select *BookSelect
	// Fields of oneof xxx_hidden_Key:
	Id []byte
	// -- end of xxx_hidden_Key
}

func (b0 BookGetRequest_builder) Build() *BookGetRequest {
	m0 := &BookGetRequest{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_Select = b.Select
	if b.Id != nil {
		x.xxx_hidden_Key = &bookGetRequest_Id{b.Id}
	}
	return m0
}

type case_BookGetRequest_Key protoreflect.FieldNumber

func (x case_BookGetRequest_Key) String() string {
	md := file_example_library_book_svc_proto_msgTypes[1].Descriptor()
	if x == 0 {
		return "not set"
	}
	return protoimpl.X.MessageFieldStringOf(md, protoreflect.FieldNumber(x))
}

type isBookGetRequest_Key interface {
	isBookGetRequest_Key()
}

type bookGetRequest_Id struct {
	Id []byte `protobuf:"bytes,2,opt,name=id,proto3,oneof"`
}

func (*bookGetRequest_Id) isBookGetRequest_Key() {}

type BookSelect struct {
	state                  protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_All         bool                   `protobuf:"varint,1,opt,name=all,proto3,oneof"`
	xxx_hidden_Title       bool                   `protobuf:"varint,3,opt,name=title,proto3,oneof"`
	xxx_hidden_Authors     *AuthorSelect          `protobuf:"bytes,4,opt,name=authors,proto3,oneof"`
	xxx_hidden_Index       bool                   `protobuf:"varint,5,opt,name=index,proto3,oneof"`
	xxx_hidden_Genres      bool                   `protobuf:"varint,6,opt,name=genres,proto3,oneof"`
	xxx_hidden_DateCreated bool                   `protobuf:"varint,15,opt,name=date_created,json=dateCreated,proto3,oneof"`
	XXX_raceDetectHookData protoimpl.RaceDetectHookData
	XXX_presence           [1]uint32
	unknownFields          protoimpl.UnknownFields
	sizeCache              protoimpl.SizeCache
}

func (x *BookSelect) Reset() {
	*x = BookSelect{}
	mi := &file_example_library_book_svc_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BookSelect) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookSelect) ProtoMessage() {}

func (x *BookSelect) ProtoReflect() protoreflect.Message {
	mi := &file_example_library_book_svc_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *BookSelect) GetAll() bool {
	if x != nil {
		return x.xxx_hidden_All
	}
	return false
}

func (x *BookSelect) GetTitle() bool {
	if x != nil {
		return x.xxx_hidden_Title
	}
	return false
}

func (x *BookSelect) GetAuthors() *AuthorSelect {
	if x != nil {
		return x.xxx_hidden_Authors
	}
	return nil
}

func (x *BookSelect) GetIndex() bool {
	if x != nil {
		return x.xxx_hidden_Index
	}
	return false
}

func (x *BookSelect) GetGenres() bool {
	if x != nil {
		return x.xxx_hidden_Genres
	}
	return false
}

func (x *BookSelect) GetDateCreated() bool {
	if x != nil {
		return x.xxx_hidden_DateCreated
	}
	return false
}

func (x *BookSelect) SetAll(v bool) {
	x.xxx_hidden_All = v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 0, 6)
}

func (x *BookSelect) SetTitle(v bool) {
	x.xxx_hidden_Title = v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 1, 6)
}

func (x *BookSelect) SetAuthors(v *AuthorSelect) {
	x.xxx_hidden_Authors = v
}

func (x *BookSelect) SetIndex(v bool) {
	x.xxx_hidden_Index = v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 3, 6)
}

func (x *BookSelect) SetGenres(v bool) {
	x.xxx_hidden_Genres = v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 4, 6)
}

func (x *BookSelect) SetDateCreated(v bool) {
	x.xxx_hidden_DateCreated = v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 5, 6)
}

func (x *BookSelect) HasAll() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 0)
}

func (x *BookSelect) HasTitle() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 1)
}

func (x *BookSelect) HasAuthors() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_Authors != nil
}

func (x *BookSelect) HasIndex() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 3)
}

func (x *BookSelect) HasGenres() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 4)
}

func (x *BookSelect) HasDateCreated() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 5)
}

func (x *BookSelect) ClearAll() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 0)
	x.xxx_hidden_All = false
}

func (x *BookSelect) ClearTitle() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 1)
	x.xxx_hidden_Title = false
}

func (x *BookSelect) ClearAuthors() {
	x.xxx_hidden_Authors = nil
}

func (x *BookSelect) ClearIndex() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 3)
	x.xxx_hidden_Index = false
}

func (x *BookSelect) ClearGenres() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 4)
	x.xxx_hidden_Genres = false
}

func (x *BookSelect) ClearDateCreated() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 5)
	x.xxx_hidden_DateCreated = false
}

type BookSelect_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	All         *bool
	Title       *bool
	Authors     *AuthorSelect
	Index       *bool
	Genres      *bool
	DateCreated *bool
}

func (b0 BookSelect_builder) Build() *BookSelect {
	m0 := &BookSelect{}
	b, x := &b0, m0
	_, _ = b, x
	if b.All != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 0, 6)
		x.xxx_hidden_All = *b.All
	}
	if b.Title != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 1, 6)
		x.xxx_hidden_Title = *b.Title
	}
	x.xxx_hidden_Authors = b.Authors
	if b.Index != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 3, 6)
		x.xxx_hidden_Index = *b.Index
	}
	if b.Genres != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 4, 6)
		x.xxx_hidden_Genres = *b.Genres
	}
	if b.DateCreated != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 5, 6)
		x.xxx_hidden_DateCreated = *b.DateCreated
	}
	return m0
}

type BookPatchRequest struct {
	state                  protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Key         *BookGetRequest        `protobuf:"bytes,1,opt,name=key,proto3"`
	xxx_hidden_Title       *string                `protobuf:"bytes,5,opt,name=title,proto3,oneof"`
	xxx_hidden_Index       *[]*Book_Index         `protobuf:"bytes,9,rep,name=index,proto3"`
	xxx_hidden_Genres      []string               `protobuf:"bytes,11,rep,name=genres,proto3"`
	XXX_raceDetectHookData protoimpl.RaceDetectHookData
	XXX_presence           [1]uint32
	unknownFields          protoimpl.UnknownFields
	sizeCache              protoimpl.SizeCache
}

func (x *BookPatchRequest) Reset() {
	*x = BookPatchRequest{}
	mi := &file_example_library_book_svc_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BookPatchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookPatchRequest) ProtoMessage() {}

func (x *BookPatchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_example_library_book_svc_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *BookPatchRequest) GetKey() *BookGetRequest {
	if x != nil {
		return x.xxx_hidden_Key
	}
	return nil
}

func (x *BookPatchRequest) GetTitle() string {
	if x != nil {
		if x.xxx_hidden_Title != nil {
			return *x.xxx_hidden_Title
		}
		return ""
	}
	return ""
}

func (x *BookPatchRequest) GetIndex() []*Book_Index {
	if x != nil {
		if x.xxx_hidden_Index != nil {
			return *x.xxx_hidden_Index
		}
	}
	return nil
}

func (x *BookPatchRequest) GetGenres() []string {
	if x != nil {
		return x.xxx_hidden_Genres
	}
	return nil
}

func (x *BookPatchRequest) SetKey(v *BookGetRequest) {
	x.xxx_hidden_Key = v
}

func (x *BookPatchRequest) SetTitle(v string) {
	x.xxx_hidden_Title = &v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 1, 4)
}

func (x *BookPatchRequest) SetIndex(v []*Book_Index) {
	x.xxx_hidden_Index = &v
}

func (x *BookPatchRequest) SetGenres(v []string) {
	x.xxx_hidden_Genres = v
}

func (x *BookPatchRequest) HasKey() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_Key != nil
}

func (x *BookPatchRequest) HasTitle() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 1)
}

func (x *BookPatchRequest) ClearKey() {
	x.xxx_hidden_Key = nil
}

func (x *BookPatchRequest) ClearTitle() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 1)
	x.xxx_hidden_Title = nil
}

type BookPatchRequest_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Key    *BookGetRequest
	Title  *string
	Index  []*Book_Index
	Genres []string
}

func (b0 BookPatchRequest_builder) Build() *BookPatchRequest {
	m0 := &BookPatchRequest{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_Key = b.Key
	if b.Title != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 1, 4)
		x.xxx_hidden_Title = b.Title
	}
	x.xxx_hidden_Index = &b.Index
	x.xxx_hidden_Genres = b.Genres
	return m0
}

var File_example_library_book_svc_proto protoreflect.FileDescriptor

var file_example_library_book_svc_proto_rawDesc = string([]byte{
	0x0a, 0x1e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72,
	0x79, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72,
	0x79, 0x1a, 0x20, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x6c, 0x69, 0x62, 0x72, 0x61,
	0x72, 0x79, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x6c, 0x69, 0x62,
	0x72, 0x61, 0x72, 0x79, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe2, 0x01,
	0x0a, 0x0e, 0x42, 0x6f, 0x6f, 0x6b, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x13, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x02,
	0x69, 0x64, 0x88, 0x01, 0x01, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x31, 0x0a, 0x05, 0x69,
	0x6e, 0x64, 0x65, 0x78, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2e, 0x42, 0x6f, 0x6f,
	0x6b, 0x2e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x16,
	0x0a, 0x06, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06,
	0x67, 0x65, 0x6e, 0x72, 0x65, 0x73, 0x12, 0x42, 0x0a, 0x0c, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x01, 0x52, 0x0b, 0x64, 0x61, 0x74, 0x65,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x88, 0x01, 0x01, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x69,
	0x64, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x22, 0x5e, 0x0a, 0x0e, 0x42, 0x6f, 0x6f, 0x6b, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x33, 0x0a, 0x06, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x6c,
	0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x53, 0x65, 0x6c, 0x65, 0x63,
	0x74, 0x52, 0x06, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x12, 0x10, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x02, 0x69, 0x64, 0x42, 0x05, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x22, 0xa0, 0x02, 0x0a, 0x0a, 0x42, 0x6f, 0x6f, 0x6b, 0x53, 0x65, 0x6c, 0x65, 0x63,
	0x74, 0x12, 0x15, 0x0a, 0x03, 0x61, 0x6c, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00,
	0x52, 0x03, 0x61, 0x6c, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x48, 0x01, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x88, 0x01, 0x01, 0x12, 0x3c, 0x0a, 0x07, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x6c,
	0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x53, 0x65, 0x6c,
	0x65, 0x63, 0x74, 0x48, 0x02, 0x52, 0x07, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x88, 0x01,
	0x01, 0x12, 0x19, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08,
	0x48, 0x03, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06,
	0x67, 0x65, 0x6e, 0x72, 0x65, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x48, 0x04, 0x52, 0x06,
	0x67, 0x65, 0x6e, 0x72, 0x65, 0x73, 0x88, 0x01, 0x01, 0x12, 0x26, 0x0a, 0x0c, 0x64, 0x61, 0x74,
	0x65, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x08, 0x48,
	0x05, 0x52, 0x0b, 0x64, 0x61, 0x74, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x88, 0x01,
	0x01, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x61, 0x6c, 0x6c, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x42,
	0x08, 0x0a, 0x06, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x67, 0x65,
	0x6e, 0x72, 0x65, 0x73, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x22, 0xb5, 0x01, 0x0a, 0x10, 0x42, 0x6f, 0x6f, 0x6b, 0x50, 0x61,
	0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x19, 0x0a,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x88, 0x01, 0x01, 0x12, 0x31, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65,
	0x78, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x2e, 0x49,
	0x6e, 0x64, 0x65, 0x78, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x16, 0x0a, 0x06, 0x67,
	0x65, 0x6e, 0x72, 0x65, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x67, 0x65, 0x6e,
	0x72, 0x65, 0x73, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x32, 0x91, 0x02,
	0x0a, 0x0b, 0x42, 0x6f, 0x6f, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3d, 0x0a,
	0x03, 0x41, 0x64, 0x64, 0x12, 0x1f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x6c,
	0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x41, 0x64, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e,
	0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x3d, 0x0a, 0x03,
	0x47, 0x65, 0x74, 0x12, 0x1f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x6c, 0x69,
	0x62, 0x72, 0x61, 0x72, 0x79, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x6c,
	0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x42, 0x0a, 0x05, 0x50,
	0x61, 0x74, 0x63, 0x68, 0x12, 0x21, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x6c,
	0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x50, 0x61, 0x74, 0x63, 0x68,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12,
	0x40, 0x0a, 0x05, 0x45, 0x72, 0x61, 0x73, 0x65, 0x12, 0x1f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6c, 0x65, 0x73, 0x6f, 0x6d, 0x6e, 0x75, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x6f,
	0x72, 0x6d, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x65, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x2f, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
})

var file_example_library_book_svc_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_example_library_book_svc_proto_goTypes = []any{
	(*BookAddRequest)(nil),        // 0: example.library.BookAddRequest
	(*BookGetRequest)(nil),        // 1: example.library.BookGetRequest
	(*BookSelect)(nil),            // 2: example.library.BookSelect
	(*BookPatchRequest)(nil),      // 3: example.library.BookPatchRequest
	(*Book_Index)(nil),            // 4: example.library.Book.Index
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
	(*AuthorSelect)(nil),          // 6: example.library.AuthorSelect
	(*Book)(nil),                  // 7: example.library.Book
	(*emptypb.Empty)(nil),         // 8: google.protobuf.Empty
}
var file_example_library_book_svc_proto_depIdxs = []int32{
	4,  // 0: example.library.BookAddRequest.index:type_name -> example.library.Book.Index
	5,  // 1: example.library.BookAddRequest.date_created:type_name -> google.protobuf.Timestamp
	2,  // 2: example.library.BookGetRequest.select:type_name -> example.library.BookSelect
	6,  // 3: example.library.BookSelect.authors:type_name -> example.library.AuthorSelect
	1,  // 4: example.library.BookPatchRequest.key:type_name -> example.library.BookGetRequest
	4,  // 5: example.library.BookPatchRequest.index:type_name -> example.library.Book.Index
	0,  // 6: example.library.BookService.Add:input_type -> example.library.BookAddRequest
	1,  // 7: example.library.BookService.Get:input_type -> example.library.BookGetRequest
	3,  // 8: example.library.BookService.Patch:input_type -> example.library.BookPatchRequest
	1,  // 9: example.library.BookService.Erase:input_type -> example.library.BookGetRequest
	7,  // 10: example.library.BookService.Add:output_type -> example.library.Book
	7,  // 11: example.library.BookService.Get:output_type -> example.library.Book
	8,  // 12: example.library.BookService.Patch:output_type -> google.protobuf.Empty
	8,  // 13: example.library.BookService.Erase:output_type -> google.protobuf.Empty
	10, // [10:14] is the sub-list for method output_type
	6,  // [6:10] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_example_library_book_svc_proto_init() }
func file_example_library_book_svc_proto_init() {
	if File_example_library_book_svc_proto != nil {
		return
	}
	file_example_library_author_svc_proto_init()
	file_example_library_book_proto_init()
	file_example_library_book_svc_proto_msgTypes[0].OneofWrappers = []any{}
	file_example_library_book_svc_proto_msgTypes[1].OneofWrappers = []any{
		(*bookGetRequest_Id)(nil),
	}
	file_example_library_book_svc_proto_msgTypes[2].OneofWrappers = []any{}
	file_example_library_book_svc_proto_msgTypes[3].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_example_library_book_svc_proto_rawDesc), len(file_example_library_book_svc_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_example_library_book_svc_proto_goTypes,
		DependencyIndexes: file_example_library_book_svc_proto_depIdxs,
		MessageInfos:      file_example_library_book_svc_proto_msgTypes,
	}.Build()
	File_example_library_book_svc_proto = out.File
	file_example_library_book_svc_proto_goTypes = nil
	file_example_library_book_svc_proto_depIdxs = nil
}
