// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        v5.29.0
// source: example/library/member.proto

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

type Member_Level int32

const (
	Member_LEVEL_UNSPECIFIED Member_Level = 0
	Member_LEVEL_BRONZE      Member_Level = 1
	Member_LEVEL_SILVER      Member_Level = 2
	Member_LEVEL_GOLD        Member_Level = 3
)

// Enum value maps for Member_Level.
var (
	Member_Level_name = map[int32]string{
		0: "LEVEL_UNSPECIFIED",
		1: "LEVEL_BRONZE",
		2: "LEVEL_SILVER",
		3: "LEVEL_GOLD",
	}
	Member_Level_value = map[string]int32{
		"LEVEL_UNSPECIFIED": 0,
		"LEVEL_BRONZE":      1,
		"LEVEL_SILVER":      2,
		"LEVEL_GOLD":        3,
	}
)

func (x Member_Level) Enum() *Member_Level {
	p := new(Member_Level)
	*p = x
	return p
}

func (x Member_Level) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Member_Level) Descriptor() protoreflect.EnumDescriptor {
	return file_example_library_member_proto_enumTypes[0].Descriptor()
}

func (Member_Level) Type() protoreflect.EnumType {
	return &file_example_library_member_proto_enumTypes[0]
}

func (x Member_Level) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

type Member struct {
	state                  protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Id          []byte                 `protobuf:"bytes,1,opt,name=id,proto3"`
	xxx_hidden_Name        string                 `protobuf:"bytes,3,opt,name=name,proto3"`
	xxx_hidden_Labels      map[string]string      `protobuf:"bytes,4,rep,name=labels,proto3" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	xxx_hidden_Profile     *Profile               `protobuf:"bytes,5,opt,name=profile,proto3"`
	xxx_hidden_Level       Member_Level           `protobuf:"varint,6,opt,name=level,proto3,enum=example.library.Member_Level"`
	xxx_hidden_Parent      *Member                `protobuf:"bytes,8,opt,name=parent,proto3,oneof"`
	xxx_hidden_Children    *[]*Member             `protobuf:"bytes,9,rep,name=children,proto3"`
	xxx_hidden_DateCreated *timestamppb.Timestamp `protobuf:"bytes,15,opt,name=date_created,json=dateCreated,proto3"`
	unknownFields          protoimpl.UnknownFields
	sizeCache              protoimpl.SizeCache
}

func (x *Member) Reset() {
	*x = Member{}
	mi := &file_example_library_member_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Member) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Member) ProtoMessage() {}

func (x *Member) ProtoReflect() protoreflect.Message {
	mi := &file_example_library_member_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *Member) GetId() []byte {
	if x != nil {
		return x.xxx_hidden_Id
	}
	return nil
}

func (x *Member) GetName() string {
	if x != nil {
		return x.xxx_hidden_Name
	}
	return ""
}

func (x *Member) GetLabels() map[string]string {
	if x != nil {
		return x.xxx_hidden_Labels
	}
	return nil
}

func (x *Member) GetProfile() *Profile {
	if x != nil {
		return x.xxx_hidden_Profile
	}
	return nil
}

func (x *Member) GetLevel() Member_Level {
	if x != nil {
		return x.xxx_hidden_Level
	}
	return Member_LEVEL_UNSPECIFIED
}

func (x *Member) GetParent() *Member {
	if x != nil {
		return x.xxx_hidden_Parent
	}
	return nil
}

func (x *Member) GetChildren() []*Member {
	if x != nil {
		if x.xxx_hidden_Children != nil {
			return *x.xxx_hidden_Children
		}
	}
	return nil
}

func (x *Member) GetDateCreated() *timestamppb.Timestamp {
	if x != nil {
		return x.xxx_hidden_DateCreated
	}
	return nil
}

func (x *Member) SetId(v []byte) {
	if v == nil {
		v = []byte{}
	}
	x.xxx_hidden_Id = v
}

func (x *Member) SetName(v string) {
	x.xxx_hidden_Name = v
}

func (x *Member) SetLabels(v map[string]string) {
	x.xxx_hidden_Labels = v
}

func (x *Member) SetProfile(v *Profile) {
	x.xxx_hidden_Profile = v
}

func (x *Member) SetLevel(v Member_Level) {
	x.xxx_hidden_Level = v
}

func (x *Member) SetParent(v *Member) {
	x.xxx_hidden_Parent = v
}

func (x *Member) SetChildren(v []*Member) {
	x.xxx_hidden_Children = &v
}

func (x *Member) SetDateCreated(v *timestamppb.Timestamp) {
	x.xxx_hidden_DateCreated = v
}

func (x *Member) HasProfile() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_Profile != nil
}

func (x *Member) HasParent() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_Parent != nil
}

func (x *Member) HasDateCreated() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_DateCreated != nil
}

func (x *Member) ClearProfile() {
	x.xxx_hidden_Profile = nil
}

func (x *Member) ClearParent() {
	x.xxx_hidden_Parent = nil
}

func (x *Member) ClearDateCreated() {
	x.xxx_hidden_DateCreated = nil
}

type Member_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Id          []byte
	Name        string
	Labels      map[string]string
	Profile     *Profile
	Level       Member_Level
	Parent      *Member
	Children    []*Member
	DateCreated *timestamppb.Timestamp
}

func (b0 Member_builder) Build() *Member {
	m0 := &Member{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_Id = b.Id
	x.xxx_hidden_Name = b.Name
	x.xxx_hidden_Labels = b.Labels
	x.xxx_hidden_Profile = b.Profile
	x.xxx_hidden_Level = b.Level
	x.xxx_hidden_Parent = b.Parent
	x.xxx_hidden_Children = &b.Children
	x.xxx_hidden_DateCreated = b.DateCreated
	return m0
}

var File_example_library_member_proto protoreflect.FileDescriptor

var file_example_library_member_proto_rawDesc = string([]byte{
	0x0a, 0x1c, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72,
	0x79, 0x2f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x09, 0x6f, 0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2f, 0x70, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe2, 0x04, 0x0a, 0x06, 0x4d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x42, 0x0a, 0xda, 0xfc, 0x15, 0x06, 0x10, 0x40, 0x28, 0x01, 0x4a, 0x00, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x1a, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x06, 0xda, 0xfc, 0x15, 0x02, 0x4a, 0x00, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x43, 0x0a,
	0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2e,
	0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x42, 0x06, 0xda, 0xfc, 0x15, 0x02, 0x4a, 0x00, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x73, 0x12, 0x32, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x6c, 0x69,
	0x62, 0x72, 0x61, 0x72, 0x79, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x07, 0x70,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x33, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e,
	0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x4c,
	0x65, 0x76, 0x65, 0x6c, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x40, 0x0a, 0x06, 0x70,
	0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2e, 0x4d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x42, 0x0a, 0xe2, 0xfc, 0x15, 0x06, 0x1a, 0x02, 0x08, 0x09, 0x38, 0x01,
	0x48, 0x00, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x12, 0x39, 0x0a,
	0x08, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72,
	0x79, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x42, 0x04, 0xe2, 0xfc, 0x15, 0x00, 0x52, 0x08,
	0x63, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x12, 0x47, 0x0a, 0x0c, 0x64, 0x61, 0x74, 0x65,
	0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x08, 0xda, 0xfc, 0x15, 0x04,
	0x40, 0x01, 0x4a, 0x00, 0x52, 0x0b, 0x64, 0x61, 0x74, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x52, 0x0a, 0x05,
	0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x15, 0x0a, 0x11, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c,
	0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x42, 0x52, 0x4f, 0x4e, 0x5a, 0x45, 0x10, 0x01, 0x12, 0x10,
	0x0a, 0x0c, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x53, 0x49, 0x4c, 0x56, 0x45, 0x52, 0x10, 0x02,
	0x12, 0x0e, 0x0a, 0x0a, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x47, 0x4f, 0x4c, 0x44, 0x10, 0x03,
	0x3a, 0x12, 0xca, 0xfc, 0x15, 0x0e, 0x12, 0x0c, 0x12, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x1a, 0x02,
	0x08, 0x03, 0x30, 0x01, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x42,
	0x42, 0xc2, 0xfc, 0x15, 0x06, 0x0a, 0x00, 0x12, 0x02, 0x10, 0x01, 0x5a, 0x36, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x65, 0x73, 0x6f, 0x6d, 0x6e, 0x75, 0x73,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x6f, 0x72, 0x6d, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x6c, 0x69, 0x62, 0x72,
	0x61, 0x72, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var file_example_library_member_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_example_library_member_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_example_library_member_proto_goTypes = []any{
	(Member_Level)(0),             // 0: example.library.Member.Level
	(*Member)(nil),                // 1: example.library.Member
	nil,                           // 2: example.library.Member.LabelsEntry
	(*Profile)(nil),               // 3: example.library.Profile
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_example_library_member_proto_depIdxs = []int32{
	2, // 0: example.library.Member.labels:type_name -> example.library.Member.LabelsEntry
	3, // 1: example.library.Member.profile:type_name -> example.library.Profile
	0, // 2: example.library.Member.level:type_name -> example.library.Member.Level
	1, // 3: example.library.Member.parent:type_name -> example.library.Member
	1, // 4: example.library.Member.children:type_name -> example.library.Member
	4, // 5: example.library.Member.date_created:type_name -> google.protobuf.Timestamp
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_example_library_member_proto_init() }
func file_example_library_member_proto_init() {
	if File_example_library_member_proto != nil {
		return
	}
	file_example_library_profile_proto_init()
	file_example_library_member_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_example_library_member_proto_rawDesc), len(file_example_library_member_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_example_library_member_proto_goTypes,
		DependencyIndexes: file_example_library_member_proto_depIdxs,
		EnumInfos:         file_example_library_member_proto_enumTypes,
		MessageInfos:      file_example_library_member_proto_msgTypes,
	}.Build()
	File_example_library_member_proto = out.File
	file_example_library_member_proto_goTypes = nil
	file_example_library_member_proto_depIdxs = nil
}
