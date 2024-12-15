// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.29.0
// source: example/library/author.proto

package library

import (
	_ "github.com/lesomnus/proto-orm"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Author struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          []byte                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Alias       string                 `protobuf:"bytes,2,opt,name=alias,proto3" json:"alias,omitempty"`
	Name        string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Labels      map[string]string      `protobuf:"bytes,4,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Profile     *Profile               `protobuf:"bytes,5,opt,name=profile,proto3" json:"profile,omitempty"`
	Books       []*Book                `protobuf:"bytes,7,rep,name=books,proto3" json:"books,omitempty"`
	DateCreated *timestamppb.Timestamp `protobuf:"bytes,15,opt,name=date_created,json=dateCreated,proto3" json:"date_created,omitempty"`
}

func (x *Author) Reset() {
	*x = Author{}
	mi := &file_example_library_author_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Author) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Author) ProtoMessage() {}

func (x *Author) ProtoReflect() protoreflect.Message {
	mi := &file_example_library_author_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Author.ProtoReflect.Descriptor instead.
func (*Author) Descriptor() ([]byte, []int) {
	return file_example_library_author_proto_rawDescGZIP(), []int{0}
}

func (x *Author) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Author) GetAlias() string {
	if x != nil {
		return x.Alias
	}
	return ""
}

func (x *Author) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Author) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *Author) GetProfile() *Profile {
	if x != nil {
		return x.Profile
	}
	return nil
}

func (x *Author) GetBooks() []*Book {
	if x != nil {
		return x.Books
	}
	return nil
}

func (x *Author) GetDateCreated() *timestamppb.Timestamp {
	if x != nil {
		return x.DateCreated
	}
	return nil
}

type Book struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          []byte                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title       string                 `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Authors     []*Author              `protobuf:"bytes,4,rep,name=authors,proto3" json:"authors,omitempty"`
	Index       []*Book_Index          `protobuf:"bytes,5,rep,name=index,proto3" json:"index,omitempty"`
	Genres      []string               `protobuf:"bytes,6,rep,name=genres,proto3" json:"genres,omitempty"`
	DateCreated *timestamppb.Timestamp `protobuf:"bytes,15,opt,name=date_created,json=dateCreated,proto3" json:"date_created,omitempty"`
}

func (x *Book) Reset() {
	*x = Book{}
	mi := &file_example_library_author_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Book) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Book) ProtoMessage() {}

func (x *Book) ProtoReflect() protoreflect.Message {
	mi := &file_example_library_author_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Book.ProtoReflect.Descriptor instead.
func (*Book) Descriptor() ([]byte, []int) {
	return file_example_library_author_proto_rawDescGZIP(), []int{1}
}

func (x *Book) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Book) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Book) GetAuthors() []*Author {
	if x != nil {
		return x.Authors
	}
	return nil
}

func (x *Book) GetIndex() []*Book_Index {
	if x != nil {
		return x.Index
	}
	return nil
}

func (x *Book) GetGenres() []string {
	if x != nil {
		return x.Genres
	}
	return nil
}

func (x *Book) GetDateCreated() *timestamppb.Timestamp {
	if x != nil {
		return x.DateCreated
	}
	return nil
}

type Book_Index struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Chapter int32             `protobuf:"varint,1,opt,name=chapter,proto3" json:"chapter,omitempty"`
	Title   string            `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Range   *Book_Index_Range `protobuf:"bytes,3,opt,name=range,proto3" json:"range,omitempty"`
}

func (x *Book_Index) Reset() {
	*x = Book_Index{}
	mi := &file_example_library_author_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Book_Index) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Book_Index) ProtoMessage() {}

func (x *Book_Index) ProtoReflect() protoreflect.Message {
	mi := &file_example_library_author_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Book_Index.ProtoReflect.Descriptor instead.
func (*Book_Index) Descriptor() ([]byte, []int) {
	return file_example_library_author_proto_rawDescGZIP(), []int{1, 0}
}

func (x *Book_Index) GetChapter() int32 {
	if x != nil {
		return x.Chapter
	}
	return 0
}

func (x *Book_Index) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Book_Index) GetRange() *Book_Index_Range {
	if x != nil {
		return x.Range
	}
	return nil
}

type Book_Index_Range struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Begin int32 `protobuf:"varint,1,opt,name=begin,proto3" json:"begin,omitempty"`
	End   int32 `protobuf:"varint,2,opt,name=end,proto3" json:"end,omitempty"`
}

func (x *Book_Index_Range) Reset() {
	*x = Book_Index_Range{}
	mi := &file_example_library_author_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Book_Index_Range) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Book_Index_Range) ProtoMessage() {}

func (x *Book_Index_Range) ProtoReflect() protoreflect.Message {
	mi := &file_example_library_author_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Book_Index_Range.ProtoReflect.Descriptor instead.
func (*Book_Index_Range) Descriptor() ([]byte, []int) {
	return file_example_library_author_proto_rawDescGZIP(), []int{1, 0, 0}
}

func (x *Book_Index_Range) GetBegin() int32 {
	if x != nil {
		return x.Begin
	}
	return 0
}

func (x *Book_Index_Range) GetEnd() int32 {
	if x != nil {
		return x.End
	}
	return 0
}

var File_example_library_author_proto protoreflect.FileDescriptor

var file_example_library_author_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72,
	0x79, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x1a,
	0x1d, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79,
	0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x09, 0x6f, 0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x80, 0x03, 0x0a, 0x06, 0x41,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x1a, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x42, 0x0a, 0xda, 0xfc, 0x15, 0x06, 0x10, 0x40, 0x28, 0x01, 0x4a, 0x00, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x1c, 0x0a, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x06, 0xda, 0xfc, 0x15, 0x02, 0x30, 0x01, 0x52, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x6c, 0x69,
	0x62, 0x72, 0x61, 0x72, 0x79, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x2e, 0x4c, 0x61, 0x62,
	0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73,
	0x12, 0x32, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x18, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x6c, 0x69, 0x62, 0x72,
	0x61, 0x72, 0x79, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x07, 0x70, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x12, 0x33, 0x0a, 0x05, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x18, 0x07, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x6c, 0x69,
	0x62, 0x72, 0x61, 0x72, 0x79, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x42, 0x06, 0xe2, 0xfc, 0x15, 0x02,
	0x22, 0x00, 0x52, 0x05, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x12, 0x47, 0x0a, 0x0c, 0x64, 0x61, 0x74,
	0x65, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x08, 0xda, 0xfc, 0x15,
	0x04, 0x40, 0x01, 0x4a, 0x00, 0x52, 0x0b, 0x64, 0x61, 0x74, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xad, 0x03,
	0x0a, 0x04, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x1a, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x42, 0x0a, 0xda, 0xfc, 0x15, 0x06, 0x10, 0x40, 0x28, 0x01, 0x4a, 0x00, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x3b, 0x0a, 0x07, 0x61, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x65, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2e, 0x41, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x42, 0x08, 0xe2, 0xfc, 0x15, 0x04, 0x10, 0x07, 0x22, 0x00, 0x52, 0x07, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x73, 0x12, 0x31, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x6c,
	0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x2e, 0x49, 0x6e, 0x64, 0x65,
	0x78, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x72,
	0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x73,
	0x12, 0x47, 0x0a, 0x0c, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x42, 0x08, 0xda, 0xfc, 0x15, 0x04, 0x40, 0x01, 0x4a, 0x00, 0x52, 0x0b, 0x64, 0x61,
	0x74, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x1a, 0xa1, 0x01, 0x0a, 0x05, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x63, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x12, 0x37, 0x0a, 0x05, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x21, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x6c, 0x69, 0x62,
	0x72, 0x61, 0x72, 0x79, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x2e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x2e,
	0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x05, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x1a, 0x2f, 0x0a, 0x05,
	0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x65,
	0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x65, 0x6e, 0x64, 0x42, 0x3a, 0xc2,
	0xfc, 0x15, 0x06, 0x0a, 0x00, 0x12, 0x02, 0x10, 0x01, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x65, 0x73, 0x6f, 0x6d, 0x6e, 0x75, 0x73, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x6f, 0x72, 0x6d, 0x2f, 0x5f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x2f, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_example_library_author_proto_rawDescOnce sync.Once
	file_example_library_author_proto_rawDescData = file_example_library_author_proto_rawDesc
)

func file_example_library_author_proto_rawDescGZIP() []byte {
	file_example_library_author_proto_rawDescOnce.Do(func() {
		file_example_library_author_proto_rawDescData = protoimpl.X.CompressGZIP(file_example_library_author_proto_rawDescData)
	})
	return file_example_library_author_proto_rawDescData
}

var file_example_library_author_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_example_library_author_proto_goTypes = []any{
	(*Author)(nil),                // 0: example.library.Author
	(*Book)(nil),                  // 1: example.library.Book
	nil,                           // 2: example.library.Author.LabelsEntry
	(*Book_Index)(nil),            // 3: example.library.Book.Index
	(*Book_Index_Range)(nil),      // 4: example.library.Book.Index.Range
	(*Profile)(nil),               // 5: example.library.Profile
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
}
var file_example_library_author_proto_depIdxs = []int32{
	2, // 0: example.library.Author.labels:type_name -> example.library.Author.LabelsEntry
	5, // 1: example.library.Author.profile:type_name -> example.library.Profile
	1, // 2: example.library.Author.books:type_name -> example.library.Book
	6, // 3: example.library.Author.date_created:type_name -> google.protobuf.Timestamp
	0, // 4: example.library.Book.authors:type_name -> example.library.Author
	3, // 5: example.library.Book.index:type_name -> example.library.Book.Index
	6, // 6: example.library.Book.date_created:type_name -> google.protobuf.Timestamp
	4, // 7: example.library.Book.Index.range:type_name -> example.library.Book.Index.Range
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_example_library_author_proto_init() }
func file_example_library_author_proto_init() {
	if File_example_library_author_proto != nil {
		return
	}
	file_example_library_profile_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_example_library_author_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_example_library_author_proto_goTypes,
		DependencyIndexes: file_example_library_author_proto_depIdxs,
		MessageInfos:      file_example_library_author_proto_msgTypes,
	}.Build()
	File_example_library_author_proto = out.File
	file_example_library_author_proto_rawDesc = nil
	file_example_library_author_proto_goTypes = nil
	file_example_library_author_proto_depIdxs = nil
}