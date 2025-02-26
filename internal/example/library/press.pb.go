// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        v5.29.0
// source: example/library/press.proto

package library

import (
	_ "github.com/lesomnus/proto-orm"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Press struct {
	state                   protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Id           []byte                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	xxx_hidden_Book         *Book                  `protobuf:"bytes,2,opt,name=book,proto3" json:"book,omitempty"`
	xxx_hidden_SerialNumber string                 `protobuf:"bytes,3,opt,name=serial_number,json=serialNumber,proto3" json:"serial_number,omitempty"`
	xxx_hidden_DateCreated  *timestamppb.Timestamp `protobuf:"bytes,15,opt,name=date_created,json=dateCreated,proto3" json:"date_created,omitempty"`
	unknownFields           protoimpl.UnknownFields
	sizeCache               protoimpl.SizeCache
}

func (x *Press) Reset() {
	*x = Press{}
	mi := &file_example_library_press_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Press) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Press) ProtoMessage() {}

func (x *Press) ProtoReflect() protoreflect.Message {
	mi := &file_example_library_press_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *Press) GetId() []byte {
	if x != nil {
		return x.xxx_hidden_Id
	}
	return nil
}

func (x *Press) GetBook() *Book {
	if x != nil {
		return x.xxx_hidden_Book
	}
	return nil
}

func (x *Press) GetSerialNumber() string {
	if x != nil {
		return x.xxx_hidden_SerialNumber
	}
	return ""
}

func (x *Press) GetDateCreated() *timestamppb.Timestamp {
	if x != nil {
		return x.xxx_hidden_DateCreated
	}
	return nil
}

func (x *Press) SetId(v []byte) {
	if v == nil {
		v = []byte{}
	}
	x.xxx_hidden_Id = v
}

func (x *Press) SetBook(v *Book) {
	x.xxx_hidden_Book = v
}

func (x *Press) SetSerialNumber(v string) {
	x.xxx_hidden_SerialNumber = v
}

func (x *Press) SetDateCreated(v *timestamppb.Timestamp) {
	x.xxx_hidden_DateCreated = v
}

func (x *Press) HasBook() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_Book != nil
}

func (x *Press) HasDateCreated() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_DateCreated != nil
}

func (x *Press) ClearBook() {
	x.xxx_hidden_Book = nil
}

func (x *Press) ClearDateCreated() {
	x.xxx_hidden_DateCreated = nil
}

type Press_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Id           []byte
	Book         *Book
	SerialNumber string
	DateCreated  *timestamppb.Timestamp
}

func (b0 Press_builder) Build() *Press {
	m0 := &Press{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_Id = b.Id
	x.xxx_hidden_Book = b.Book
	x.xxx_hidden_SerialNumber = b.SerialNumber
	x.xxx_hidden_DateCreated = b.DateCreated
	return m0
}

var File_example_library_press_proto protoreflect.FileDescriptor

var file_example_library_press_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72,
	0x79, 0x2f, 0x70, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x65,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x1a, 0x1a,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2f,
	0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x09, 0x6f, 0x72, 0x6d,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xeb, 0x01, 0x0a, 0x05, 0x50, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x1a, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x0a, 0xda, 0xfc,
	0x15, 0x06, 0x10, 0x40, 0x28, 0x01, 0x4a, 0x00, 0x52, 0x02, 0x69, 0x64, 0x12, 0x31, 0x0a, 0x04,
	0x62, 0x6f, 0x6f, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2e, 0x42, 0x6f, 0x6f,
	0x6b, 0x42, 0x06, 0xe2, 0xfc, 0x15, 0x02, 0x40, 0x01, 0x52, 0x04, 0x62, 0x6f, 0x6f, 0x6b, 0x12,
	0x2b, 0x0a, 0x0d, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xda, 0xfc, 0x15, 0x02, 0x40, 0x01, 0x52, 0x0c,
	0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x47, 0x0a, 0x0c,
	0x64, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x0f, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x08,
	0xda, 0xfc, 0x15, 0x04, 0x40, 0x01, 0x4a, 0x00, 0x52, 0x0b, 0x64, 0x61, 0x74, 0x65, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x3a, 0x1d, 0xca, 0xfc, 0x15, 0x19, 0x12, 0x17, 0x12, 0x0d, 0x73,
	0x65, 0x72, 0x69, 0x61, 0x6c, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x1a, 0x02, 0x02, 0x03,
	0x30, 0x01, 0x40, 0x01, 0x42, 0x42, 0xc2, 0xfc, 0x15, 0x06, 0x0a, 0x00, 0x12, 0x02, 0x10, 0x01,
	0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x65, 0x73,
	0x6f, 0x6d, 0x6e, 0x75, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x6f, 0x72, 0x6d, 0x2f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x2f, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_example_library_press_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_example_library_press_proto_goTypes = []any{
	(*Press)(nil),                 // 0: example.library.Press
	(*Book)(nil),                  // 1: example.library.Book
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_example_library_press_proto_depIdxs = []int32{
	1, // 0: example.library.Press.book:type_name -> example.library.Book
	2, // 1: example.library.Press.date_created:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_example_library_press_proto_init() }
func file_example_library_press_proto_init() {
	if File_example_library_press_proto != nil {
		return
	}
	file_example_library_book_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_example_library_press_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_example_library_press_proto_goTypes,
		DependencyIndexes: file_example_library_press_proto_depIdxs,
		MessageInfos:      file_example_library_press_proto_msgTypes,
	}.Build()
	File_example_library_press_proto = out.File
	file_example_library_press_proto_rawDesc = nil
	file_example_library_press_proto_goTypes = nil
	file_example_library_press_proto_depIdxs = nil
}
