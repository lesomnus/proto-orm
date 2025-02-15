// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        v5.29.0
// source: example/library/book.proto

package library

import (
	_ "github.com/lesomnus/proto-orm"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type Book struct {
	state                  protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Id          []byte                 `protobuf:"bytes,1,opt,name=id,proto3"`
	xxx_hidden_Title       string                 `protobuf:"bytes,3,opt,name=title,proto3"`
	xxx_hidden_Authors     *[]*Author             `protobuf:"bytes,4,rep,name=authors,proto3"`
	xxx_hidden_Index       *[]*Book_Index         `protobuf:"bytes,5,rep,name=index,proto3"`
	xxx_hidden_Genres      []string               `protobuf:"bytes,6,rep,name=genres,proto3"`
	xxx_hidden_DateCreated *timestamppb.Timestamp `protobuf:"bytes,15,opt,name=date_created,json=dateCreated,proto3"`
	unknownFields          protoimpl.UnknownFields
	sizeCache              protoimpl.SizeCache
}

func (x *Book) Reset() {
	*x = Book{}
	mi := &file_example_library_book_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Book) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Book) ProtoMessage() {}

func (x *Book) ProtoReflect() protoreflect.Message {
	mi := &file_example_library_book_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *Book) GetId() []byte {
	if x != nil {
		return x.xxx_hidden_Id
	}
	return nil
}

func (x *Book) GetTitle() string {
	if x != nil {
		return x.xxx_hidden_Title
	}
	return ""
}

func (x *Book) GetAuthors() []*Author {
	if x != nil {
		if x.xxx_hidden_Authors != nil {
			return *x.xxx_hidden_Authors
		}
	}
	return nil
}

func (x *Book) GetIndex() []*Book_Index {
	if x != nil {
		if x.xxx_hidden_Index != nil {
			return *x.xxx_hidden_Index
		}
	}
	return nil
}

func (x *Book) GetGenres() []string {
	if x != nil {
		return x.xxx_hidden_Genres
	}
	return nil
}

func (x *Book) GetDateCreated() *timestamppb.Timestamp {
	if x != nil {
		return x.xxx_hidden_DateCreated
	}
	return nil
}

func (x *Book) SetId(v []byte) {
	if v == nil {
		v = []byte{}
	}
	x.xxx_hidden_Id = v
}

func (x *Book) SetTitle(v string) {
	x.xxx_hidden_Title = v
}

func (x *Book) SetAuthors(v []*Author) {
	x.xxx_hidden_Authors = &v
}

func (x *Book) SetIndex(v []*Book_Index) {
	x.xxx_hidden_Index = &v
}

func (x *Book) SetGenres(v []string) {
	x.xxx_hidden_Genres = v
}

func (x *Book) SetDateCreated(v *timestamppb.Timestamp) {
	x.xxx_hidden_DateCreated = v
}

func (x *Book) HasDateCreated() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_DateCreated != nil
}

func (x *Book) ClearDateCreated() {
	x.xxx_hidden_DateCreated = nil
}

type Book_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Id          []byte
	Title       string
	Authors     []*Author
	Index       []*Book_Index
	Genres      []string
	DateCreated *timestamppb.Timestamp
}

func (b0 Book_builder) Build() *Book {
	m0 := &Book{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_Id = b.Id
	x.xxx_hidden_Title = b.Title
	x.xxx_hidden_Authors = &b.Authors
	x.xxx_hidden_Index = &b.Index
	x.xxx_hidden_Genres = b.Genres
	x.xxx_hidden_DateCreated = b.DateCreated
	return m0
}

type Book_Index struct {
	state              protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Chapter int32                  `protobuf:"varint,1,opt,name=chapter,proto3"`
	xxx_hidden_Title   string                 `protobuf:"bytes,2,opt,name=title,proto3"`
	xxx_hidden_Range   *Book_Index_Range      `protobuf:"bytes,3,opt,name=range,proto3"`
	unknownFields      protoimpl.UnknownFields
	sizeCache          protoimpl.SizeCache
}

func (x *Book_Index) Reset() {
	*x = Book_Index{}
	mi := &file_example_library_book_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Book_Index) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Book_Index) ProtoMessage() {}

func (x *Book_Index) ProtoReflect() protoreflect.Message {
	mi := &file_example_library_book_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *Book_Index) GetChapter() int32 {
	if x != nil {
		return x.xxx_hidden_Chapter
	}
	return 0
}

func (x *Book_Index) GetTitle() string {
	if x != nil {
		return x.xxx_hidden_Title
	}
	return ""
}

func (x *Book_Index) GetRange() *Book_Index_Range {
	if x != nil {
		return x.xxx_hidden_Range
	}
	return nil
}

func (x *Book_Index) SetChapter(v int32) {
	x.xxx_hidden_Chapter = v
}

func (x *Book_Index) SetTitle(v string) {
	x.xxx_hidden_Title = v
}

func (x *Book_Index) SetRange(v *Book_Index_Range) {
	x.xxx_hidden_Range = v
}

func (x *Book_Index) HasRange() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_Range != nil
}

func (x *Book_Index) ClearRange() {
	x.xxx_hidden_Range = nil
}

type Book_Index_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Chapter int32
	Title   string
	Range   *Book_Index_Range
}

func (b0 Book_Index_builder) Build() *Book_Index {
	m0 := &Book_Index{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_Chapter = b.Chapter
	x.xxx_hidden_Title = b.Title
	x.xxx_hidden_Range = b.Range
	return m0
}

type Book_Index_Range struct {
	state            protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Begin int32                  `protobuf:"varint,1,opt,name=begin,proto3"`
	xxx_hidden_End   int32                  `protobuf:"varint,2,opt,name=end,proto3"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *Book_Index_Range) Reset() {
	*x = Book_Index_Range{}
	mi := &file_example_library_book_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Book_Index_Range) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Book_Index_Range) ProtoMessage() {}

func (x *Book_Index_Range) ProtoReflect() protoreflect.Message {
	mi := &file_example_library_book_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *Book_Index_Range) GetBegin() int32 {
	if x != nil {
		return x.xxx_hidden_Begin
	}
	return 0
}

func (x *Book_Index_Range) GetEnd() int32 {
	if x != nil {
		return x.xxx_hidden_End
	}
	return 0
}

func (x *Book_Index_Range) SetBegin(v int32) {
	x.xxx_hidden_Begin = v
}

func (x *Book_Index_Range) SetEnd(v int32) {
	x.xxx_hidden_End = v
}

type Book_Index_Range_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Begin int32
	End   int32
}

func (b0 Book_Index_Range_builder) Build() *Book_Index_Range {
	m0 := &Book_Index_Range{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_Begin = b.Begin
	x.xxx_hidden_End = b.End
	return m0
}

var File_example_library_book_proto protoreflect.FileDescriptor

var file_example_library_book_proto_rawDesc = string([]byte{
	0x0a, 0x1a, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72,
	0x79, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x09,
	0x6f, 0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x65, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x2f, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb8, 0x03, 0x0a, 0x04, 0x42, 0x6f, 0x6f, 0x6b,
	0x12, 0x1a, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x0a, 0xda, 0xfc,
	0x15, 0x06, 0x10, 0x40, 0x28, 0x01, 0x4a, 0x00, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x46, 0x0a, 0x07, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x6c, 0x69,
	0x62, 0x72, 0x61, 0x72, 0x79, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x42, 0x13, 0xe2, 0xfc,
	0x15, 0x0f, 0x1a, 0x0d, 0x08, 0x07, 0x12, 0x09, 0x0a, 0x05, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x10,
	0x01, 0x52, 0x07, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x12, 0x31, 0x0a, 0x05, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x65, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2e, 0x42, 0x6f, 0x6f, 0x6b,
	0x2e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x16, 0x0a,
	0x06, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x67,
	0x65, 0x6e, 0x72, 0x65, 0x73, 0x12, 0x47, 0x0a, 0x0c, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x08, 0xda, 0xfc, 0x15, 0x04, 0x40, 0x01, 0x4a,
	0x00, 0x52, 0x0b, 0x64, 0x61, 0x74, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x1a, 0xa1,
	0x01, 0x0a, 0x05, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x70,
	0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x63, 0x68, 0x61, 0x70, 0x74,
	0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x37, 0x0a, 0x05, 0x72, 0x61, 0x6e, 0x67,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x2e, 0x49,
	0x6e, 0x64, 0x65, 0x78, 0x2e, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x05, 0x72, 0x61, 0x6e, 0x67,
	0x65, 0x1a, 0x2f, 0x0a, 0x05, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x65,
	0x67, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x62, 0x65, 0x67, 0x69, 0x6e,
	0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x65,
	0x6e, 0x64, 0x42, 0x42, 0xc2, 0xfc, 0x15, 0x06, 0x0a, 0x00, 0x12, 0x02, 0x10, 0x01, 0x5a, 0x36,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x65, 0x73, 0x6f, 0x6d,
	0x6e, 0x75, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x6f, 0x72, 0x6d, 0x2f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x6c,
	0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var file_example_library_book_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_example_library_book_proto_goTypes = []any{
	(*Book)(nil),                  // 0: example.library.Book
	(*Book_Index)(nil),            // 1: example.library.Book.Index
	(*Book_Index_Range)(nil),      // 2: example.library.Book.Index.Range
	(*Author)(nil),                // 3: example.library.Author
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_example_library_book_proto_depIdxs = []int32{
	3, // 0: example.library.Book.authors:type_name -> example.library.Author
	1, // 1: example.library.Book.index:type_name -> example.library.Book.Index
	4, // 2: example.library.Book.date_created:type_name -> google.protobuf.Timestamp
	2, // 3: example.library.Book.Index.range:type_name -> example.library.Book.Index.Range
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_example_library_book_proto_init() }
func file_example_library_book_proto_init() {
	if File_example_library_book_proto != nil {
		return
	}
	file_example_library_author_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_example_library_book_proto_rawDesc), len(file_example_library_book_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_example_library_book_proto_goTypes,
		DependencyIndexes: file_example_library_book_proto_depIdxs,
		MessageInfos:      file_example_library_book_proto_msgTypes,
	}.Build()
	File_example_library_book_proto = out.File
	file_example_library_book_proto_goTypes = nil
	file_example_library_book_proto_depIdxs = nil
}
