// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.29.0
// source: example/library/profile.proto

package library

import (
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

type Profile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email     string                 `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	DateBirth *timestamppb.Timestamp `protobuf:"bytes,15,opt,name=date_birth,json=dateBirth,proto3" json:"date_birth,omitempty"`
}

func (x *Profile) Reset() {
	*x = Profile{}
	mi := &file_example_library_profile_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Profile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Profile) ProtoMessage() {}

func (x *Profile) ProtoReflect() protoreflect.Message {
	mi := &file_example_library_profile_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Profile.ProtoReflect.Descriptor instead.
func (*Profile) Descriptor() ([]byte, []int) {
	return file_example_library_profile_proto_rawDescGZIP(), []int{0}
}

func (x *Profile) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Profile) GetDateBirth() *timestamppb.Timestamp {
	if x != nil {
		return x.DateBirth
	}
	return nil
}

var File_example_library_profile_proto protoreflect.FileDescriptor

var file_example_library_profile_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72,
	0x79, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x5a, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x12, 0x39, 0x0a, 0x0a, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x62, 0x69, 0x72, 0x74, 0x68,
	0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x64, 0x61, 0x74, 0x65, 0x42, 0x69, 0x72, 0x74, 0x68, 0x42, 0x38, 0x5a,
	0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x65, 0x73, 0x6f,
	0x6d, 0x6e, 0x75, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x6f, 0x72, 0x6d, 0x2f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f,
	0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_example_library_profile_proto_rawDescOnce sync.Once
	file_example_library_profile_proto_rawDescData = file_example_library_profile_proto_rawDesc
)

func file_example_library_profile_proto_rawDescGZIP() []byte {
	file_example_library_profile_proto_rawDescOnce.Do(func() {
		file_example_library_profile_proto_rawDescData = protoimpl.X.CompressGZIP(file_example_library_profile_proto_rawDescData)
	})
	return file_example_library_profile_proto_rawDescData
}

var file_example_library_profile_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_example_library_profile_proto_goTypes = []any{
	(*Profile)(nil),               // 0: example.library.Profile
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
}
var file_example_library_profile_proto_depIdxs = []int32{
	1, // 0: example.library.Profile.date_birth:type_name -> google.protobuf.Timestamp
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_example_library_profile_proto_init() }
func file_example_library_profile_proto_init() {
	if File_example_library_profile_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_example_library_profile_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_example_library_profile_proto_goTypes,
		DependencyIndexes: file_example_library_profile_proto_depIdxs,
		MessageInfos:      file_example_library_profile_proto_msgTypes,
	}.Build()
	File_example_library_profile_proto = out.File
	file_example_library_profile_proto_rawDesc = nil
	file_example_library_profile_proto_goTypes = nil
	file_example_library_profile_proto_depIdxs = nil
}