// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        v5.29.0
// source: example/library/like.proto

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

type Like struct {
	state                  protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Id          []byte                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	xxx_hidden_SubjectId   []byte                 `protobuf:"bytes,2,opt,name=subject_id,json=subjectId,proto3" json:"subject_id,omitempty"`
	xxx_hidden_Subject     *Book                  `protobuf:"bytes,3,opt,name=subject,proto3" json:"subject,omitempty"`
	xxx_hidden_Actor       *Member                `protobuf:"bytes,5,opt,name=actor,proto3" json:"actor,omitempty"`
	xxx_hidden_DateCreated *timestamppb.Timestamp `protobuf:"bytes,15,opt,name=date_created,json=dateCreated,proto3" json:"date_created,omitempty"`
	unknownFields          protoimpl.UnknownFields
	sizeCache              protoimpl.SizeCache
}

func (x *Like) Reset() {
	*x = Like{}
	mi := &file_example_library_like_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Like) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Like) ProtoMessage() {}

func (x *Like) ProtoReflect() protoreflect.Message {
	mi := &file_example_library_like_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *Like) GetId() []byte {
	if x != nil {
		return x.xxx_hidden_Id
	}
	return nil
}

func (x *Like) GetSubjectId() []byte {
	if x != nil {
		return x.xxx_hidden_SubjectId
	}
	return nil
}

func (x *Like) GetSubject() *Book {
	if x != nil {
		return x.xxx_hidden_Subject
	}
	return nil
}

func (x *Like) GetActor() *Member {
	if x != nil {
		return x.xxx_hidden_Actor
	}
	return nil
}

func (x *Like) GetDateCreated() *timestamppb.Timestamp {
	if x != nil {
		return x.xxx_hidden_DateCreated
	}
	return nil
}

func (x *Like) SetId(v []byte) {
	if v == nil {
		v = []byte{}
	}
	x.xxx_hidden_Id = v
}

func (x *Like) SetSubjectId(v []byte) {
	if v == nil {
		v = []byte{}
	}
	x.xxx_hidden_SubjectId = v
}

func (x *Like) SetSubject(v *Book) {
	x.xxx_hidden_Subject = v
}

func (x *Like) SetActor(v *Member) {
	x.xxx_hidden_Actor = v
}

func (x *Like) SetDateCreated(v *timestamppb.Timestamp) {
	x.xxx_hidden_DateCreated = v
}

func (x *Like) HasSubject() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_Subject != nil
}

func (x *Like) HasActor() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_Actor != nil
}

func (x *Like) HasDateCreated() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_DateCreated != nil
}

func (x *Like) ClearSubject() {
	x.xxx_hidden_Subject = nil
}

func (x *Like) ClearActor() {
	x.xxx_hidden_Actor = nil
}

func (x *Like) ClearDateCreated() {
	x.xxx_hidden_DateCreated = nil
}

type Like_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Id          []byte
	SubjectId   []byte
	Subject     *Book
	Actor       *Member
	DateCreated *timestamppb.Timestamp
}

func (b0 Like_builder) Build() *Like {
	m0 := &Like{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_Id = b.Id
	x.xxx_hidden_SubjectId = b.SubjectId
	x.xxx_hidden_Subject = b.Subject
	x.xxx_hidden_Actor = b.Actor
	x.xxx_hidden_DateCreated = b.DateCreated
	return m0
}

var File_example_library_like_proto protoreflect.FileDescriptor

var file_example_library_like_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72,
	0x79, 0x2f, 0x6c, 0x69, 0x6b, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x1a, 0x1a, 0x65,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2f, 0x62,
	0x6f, 0x6f, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x65, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x2f, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2f, 0x6d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x09, 0x6f, 0x72, 0x6d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x9d, 0x02, 0x0a, 0x04, 0x4c, 0x69, 0x6b, 0x65, 0x12, 0x1a, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x0a, 0xda, 0xfc, 0x15, 0x06, 0x10, 0x40,
	0x28, 0x01, 0x4a, 0x00, 0x52, 0x02, 0x69, 0x64, 0x12, 0x25, 0x0a, 0x0a, 0x73, 0x75, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x06, 0xda, 0xfc,
	0x15, 0x02, 0x10, 0x40, 0x52, 0x09, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12,
	0x39, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61,
	0x72, 0x79, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x42, 0x08, 0xe2, 0xfc, 0x15, 0x04, 0x10, 0x02, 0x40,
	0x01, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x35, 0x0a, 0x05, 0x61, 0x63,
	0x74, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x65, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2e, 0x4d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x42, 0x06, 0xe2, 0xfc, 0x15, 0x02, 0x40, 0x01, 0x52, 0x05, 0x61, 0x63, 0x74, 0x6f,
	0x72, 0x12, 0x47, 0x0a, 0x0c, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x42, 0x08, 0xda, 0xfc, 0x15, 0x04, 0x40, 0x01, 0x4a, 0x00, 0x52, 0x0b, 0x64,
	0x61, 0x74, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x3a, 0x17, 0xca, 0xfc, 0x15, 0x13,
	0x12, 0x11, 0x12, 0x07, 0x68, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x73, 0x1a, 0x02, 0x02, 0x05, 0x30,
	0x01, 0x40, 0x01, 0x42, 0x42, 0xc2, 0xfc, 0x15, 0x06, 0x0a, 0x00, 0x12, 0x02, 0x10, 0x01, 0x5a,
	0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x65, 0x73, 0x6f,
	0x6d, 0x6e, 0x75, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x6f, 0x72, 0x6d, 0x2f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f,
	0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_example_library_like_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_example_library_like_proto_goTypes = []any{
	(*Like)(nil),                  // 0: example.library.Like
	(*Book)(nil),                  // 1: example.library.Book
	(*Member)(nil),                // 2: example.library.Member
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_example_library_like_proto_depIdxs = []int32{
	1, // 0: example.library.Like.subject:type_name -> example.library.Book
	2, // 1: example.library.Like.actor:type_name -> example.library.Member
	3, // 2: example.library.Like.date_created:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_example_library_like_proto_init() }
func file_example_library_like_proto_init() {
	if File_example_library_like_proto != nil {
		return
	}
	file_example_library_book_proto_init()
	file_example_library_member_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_example_library_like_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_example_library_like_proto_goTypes,
		DependencyIndexes: file_example_library_like_proto_depIdxs,
		MessageInfos:      file_example_library_like_proto_msgTypes,
	}.Build()
	File_example_library_like_proto = out.File
	file_example_library_like_proto_rawDesc = nil
	file_example_library_like_proto_goTypes = nil
	file_example_library_like_proto_depIdxs = nil
}
