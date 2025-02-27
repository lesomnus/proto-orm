// Code generated by "protoc-gen-orm-service", DO NOT EDIT.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        v5.29.0
// source: example/library/membership_svc.proto

package library

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type MembershipAddRequest struct {
	state                  protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Id          []byte                 `protobuf:"bytes,1,opt,name=id,proto3,oneof"`
	xxx_hidden_Member      *MemberGetRequest      `protobuf:"bytes,2,opt,name=member,proto3"`
	xxx_hidden_Point       int32                  `protobuf:"varint,3,opt,name=point,proto3"`
	XXX_raceDetectHookData protoimpl.RaceDetectHookData
	XXX_presence           [1]uint32
	unknownFields          protoimpl.UnknownFields
	sizeCache              protoimpl.SizeCache
}

func (x *MembershipAddRequest) Reset() {
	*x = MembershipAddRequest{}
	mi := &file_example_library_membership_svc_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MembershipAddRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MembershipAddRequest) ProtoMessage() {}

func (x *MembershipAddRequest) ProtoReflect() protoreflect.Message {
	mi := &file_example_library_membership_svc_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *MembershipAddRequest) GetId() []byte {
	if x != nil {
		return x.xxx_hidden_Id
	}
	return nil
}

func (x *MembershipAddRequest) GetMember() *MemberGetRequest {
	if x != nil {
		return x.xxx_hidden_Member
	}
	return nil
}

func (x *MembershipAddRequest) GetPoint() int32 {
	if x != nil {
		return x.xxx_hidden_Point
	}
	return 0
}

func (x *MembershipAddRequest) SetId(v []byte) {
	if v == nil {
		v = []byte{}
	}
	x.xxx_hidden_Id = v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 0, 3)
}

func (x *MembershipAddRequest) SetMember(v *MemberGetRequest) {
	x.xxx_hidden_Member = v
}

func (x *MembershipAddRequest) SetPoint(v int32) {
	x.xxx_hidden_Point = v
}

func (x *MembershipAddRequest) HasId() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 0)
}

func (x *MembershipAddRequest) HasMember() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_Member != nil
}

func (x *MembershipAddRequest) ClearId() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 0)
	x.xxx_hidden_Id = nil
}

func (x *MembershipAddRequest) ClearMember() {
	x.xxx_hidden_Member = nil
}

type MembershipAddRequest_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Id     []byte
	Member *MemberGetRequest
	Point  int32
}

func (b0 MembershipAddRequest_builder) Build() *MembershipAddRequest {
	m0 := &MembershipAddRequest{}
	b, x := &b0, m0
	_, _ = b, x
	if b.Id != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 0, 3)
		x.xxx_hidden_Id = b.Id
	}
	x.xxx_hidden_Member = b.Member
	x.xxx_hidden_Point = b.Point
	return m0
}

type MembershipGetRequest struct {
	state             protoimpl.MessageState     `protogen:"opaque.v1"`
	xxx_hidden_Select *MembershipSelect          `protobuf:"bytes,1,opt,name=select,proto3"`
	xxx_hidden_Key    isMembershipGetRequest_Key `protobuf_oneof:"key"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *MembershipGetRequest) Reset() {
	*x = MembershipGetRequest{}
	mi := &file_example_library_membership_svc_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MembershipGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MembershipGetRequest) ProtoMessage() {}

func (x *MembershipGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_example_library_membership_svc_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *MembershipGetRequest) GetSelect() *MembershipSelect {
	if x != nil {
		return x.xxx_hidden_Select
	}
	return nil
}

func (x *MembershipGetRequest) GetId() []byte {
	if x != nil {
		if x, ok := x.xxx_hidden_Key.(*membershipGetRequest_Id); ok {
			return x.Id
		}
	}
	return nil
}

func (x *MembershipGetRequest) SetSelect(v *MembershipSelect) {
	x.xxx_hidden_Select = v
}

func (x *MembershipGetRequest) SetId(v []byte) {
	if v == nil {
		v = []byte{}
	}
	x.xxx_hidden_Key = &membershipGetRequest_Id{v}
}

func (x *MembershipGetRequest) HasSelect() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_Select != nil
}

func (x *MembershipGetRequest) HasKey() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_Key != nil
}

func (x *MembershipGetRequest) HasId() bool {
	if x == nil {
		return false
	}
	_, ok := x.xxx_hidden_Key.(*membershipGetRequest_Id)
	return ok
}

func (x *MembershipGetRequest) ClearSelect() {
	x.xxx_hidden_Select = nil
}

func (x *MembershipGetRequest) ClearKey() {
	x.xxx_hidden_Key = nil
}

func (x *MembershipGetRequest) ClearId() {
	if _, ok := x.xxx_hidden_Key.(*membershipGetRequest_Id); ok {
		x.xxx_hidden_Key = nil
	}
}

const MembershipGetRequest_Key_not_set_case case_MembershipGetRequest_Key = 0
const MembershipGetRequest_Id_case case_MembershipGetRequest_Key = 2

func (x *MembershipGetRequest) WhichKey() case_MembershipGetRequest_Key {
	if x == nil {
		return MembershipGetRequest_Key_not_set_case
	}
	switch x.xxx_hidden_Key.(type) {
	case *membershipGetRequest_Id:
		return MembershipGetRequest_Id_case
	default:
		return MembershipGetRequest_Key_not_set_case
	}
}

type MembershipGetRequest_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Select *MembershipSelect
	// Fields of oneof xxx_hidden_Key:
	Id []byte
	// -- end of xxx_hidden_Key
}

func (b0 MembershipGetRequest_builder) Build() *MembershipGetRequest {
	m0 := &MembershipGetRequest{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_Select = b.Select
	if b.Id != nil {
		x.xxx_hidden_Key = &membershipGetRequest_Id{b.Id}
	}
	return m0
}

type case_MembershipGetRequest_Key protoreflect.FieldNumber

func (x case_MembershipGetRequest_Key) String() string {
	md := file_example_library_membership_svc_proto_msgTypes[1].Descriptor()
	if x == 0 {
		return "not set"
	}
	return protoimpl.X.MessageFieldStringOf(md, protoreflect.FieldNumber(x))
}

type isMembershipGetRequest_Key interface {
	isMembershipGetRequest_Key()
}

type membershipGetRequest_Id struct {
	Id []byte `protobuf:"bytes,2,opt,name=id,proto3,oneof"`
}

func (*membershipGetRequest_Id) isMembershipGetRequest_Key() {}

type MembershipSelect struct {
	state                  protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_All         bool                   `protobuf:"varint,1,opt,name=all,proto3,oneof"`
	xxx_hidden_Member      *MemberSelect          `protobuf:"bytes,2,opt,name=member,proto3,oneof"`
	xxx_hidden_Point       bool                   `protobuf:"varint,3,opt,name=point,proto3,oneof"`
	XXX_raceDetectHookData protoimpl.RaceDetectHookData
	XXX_presence           [1]uint32
	unknownFields          protoimpl.UnknownFields
	sizeCache              protoimpl.SizeCache
}

func (x *MembershipSelect) Reset() {
	*x = MembershipSelect{}
	mi := &file_example_library_membership_svc_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MembershipSelect) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MembershipSelect) ProtoMessage() {}

func (x *MembershipSelect) ProtoReflect() protoreflect.Message {
	mi := &file_example_library_membership_svc_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *MembershipSelect) GetAll() bool {
	if x != nil {
		return x.xxx_hidden_All
	}
	return false
}

func (x *MembershipSelect) GetMember() *MemberSelect {
	if x != nil {
		return x.xxx_hidden_Member
	}
	return nil
}

func (x *MembershipSelect) GetPoint() bool {
	if x != nil {
		return x.xxx_hidden_Point
	}
	return false
}

func (x *MembershipSelect) SetAll(v bool) {
	x.xxx_hidden_All = v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 0, 3)
}

func (x *MembershipSelect) SetMember(v *MemberSelect) {
	x.xxx_hidden_Member = v
}

func (x *MembershipSelect) SetPoint(v bool) {
	x.xxx_hidden_Point = v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 2, 3)
}

func (x *MembershipSelect) HasAll() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 0)
}

func (x *MembershipSelect) HasMember() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_Member != nil
}

func (x *MembershipSelect) HasPoint() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 2)
}

func (x *MembershipSelect) ClearAll() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 0)
	x.xxx_hidden_All = false
}

func (x *MembershipSelect) ClearMember() {
	x.xxx_hidden_Member = nil
}

func (x *MembershipSelect) ClearPoint() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 2)
	x.xxx_hidden_Point = false
}

type MembershipSelect_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	All    *bool
	Member *MemberSelect
	Point  *bool
}

func (b0 MembershipSelect_builder) Build() *MembershipSelect {
	m0 := &MembershipSelect{}
	b, x := &b0, m0
	_, _ = b, x
	if b.All != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 0, 3)
		x.xxx_hidden_All = *b.All
	}
	x.xxx_hidden_Member = b.Member
	if b.Point != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 2, 3)
		x.xxx_hidden_Point = *b.Point
	}
	return m0
}

type MembershipPatchRequest struct {
	state                  protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Key         *MembershipGetRequest  `protobuf:"bytes,1,opt,name=key,proto3"`
	xxx_hidden_Member      *MemberGetRequest      `protobuf:"bytes,3,opt,name=member,proto3"`
	xxx_hidden_Point       int32                  `protobuf:"varint,5,opt,name=point,proto3,oneof"`
	XXX_raceDetectHookData protoimpl.RaceDetectHookData
	XXX_presence           [1]uint32
	unknownFields          protoimpl.UnknownFields
	sizeCache              protoimpl.SizeCache
}

func (x *MembershipPatchRequest) Reset() {
	*x = MembershipPatchRequest{}
	mi := &file_example_library_membership_svc_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MembershipPatchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MembershipPatchRequest) ProtoMessage() {}

func (x *MembershipPatchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_example_library_membership_svc_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *MembershipPatchRequest) GetKey() *MembershipGetRequest {
	if x != nil {
		return x.xxx_hidden_Key
	}
	return nil
}

func (x *MembershipPatchRequest) GetMember() *MemberGetRequest {
	if x != nil {
		return x.xxx_hidden_Member
	}
	return nil
}

func (x *MembershipPatchRequest) GetPoint() int32 {
	if x != nil {
		return x.xxx_hidden_Point
	}
	return 0
}

func (x *MembershipPatchRequest) SetKey(v *MembershipGetRequest) {
	x.xxx_hidden_Key = v
}

func (x *MembershipPatchRequest) SetMember(v *MemberGetRequest) {
	x.xxx_hidden_Member = v
}

func (x *MembershipPatchRequest) SetPoint(v int32) {
	x.xxx_hidden_Point = v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 2, 3)
}

func (x *MembershipPatchRequest) HasKey() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_Key != nil
}

func (x *MembershipPatchRequest) HasMember() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_Member != nil
}

func (x *MembershipPatchRequest) HasPoint() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 2)
}

func (x *MembershipPatchRequest) ClearKey() {
	x.xxx_hidden_Key = nil
}

func (x *MembershipPatchRequest) ClearMember() {
	x.xxx_hidden_Member = nil
}

func (x *MembershipPatchRequest) ClearPoint() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 2)
	x.xxx_hidden_Point = 0
}

type MembershipPatchRequest_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Key    *MembershipGetRequest
	Member *MemberGetRequest
	Point  *int32
}

func (b0 MembershipPatchRequest_builder) Build() *MembershipPatchRequest {
	m0 := &MembershipPatchRequest{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_Key = b.Key
	x.xxx_hidden_Member = b.Member
	if b.Point != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 2, 3)
		x.xxx_hidden_Point = *b.Point
	}
	return m0
}

var File_example_library_membership_svc_proto protoreflect.FileDescriptor

var file_example_library_membership_svc_proto_rawDesc = string([]byte{
	0x0a, 0x24, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72,
	0x79, 0x2f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x5f, 0x73, 0x76, 0x63,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e,
	0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x1a, 0x20, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x2f, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f,
	0x73, 0x76, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x65, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x2f, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2f, 0x6d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x73, 0x68, 0x69, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70,
	0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x83, 0x01, 0x0a, 0x14, 0x4d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x13, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52,
	0x02, 0x69, 0x64, 0x88, 0x01, 0x01, 0x12, 0x39, 0x0a, 0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x69, 0x64, 0x22, 0x6a,
	0x0a, 0x14, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x39, 0x0a, 0x06, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73,
	0x68, 0x69, 0x70, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x52, 0x06, 0x73, 0x65, 0x6c, 0x65, 0x63,
	0x74, 0x12, 0x10, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52,
	0x02, 0x69, 0x64, 0x42, 0x05, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x9d, 0x01, 0x0a, 0x10, 0x4d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x12,
	0x15, 0x0a, 0x03, 0x61, 0x6c, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x03,
	0x61, 0x6c, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x3a, 0x0a, 0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x53,
	0x65, 0x6c, 0x65, 0x63, 0x74, 0x48, 0x01, 0x52, 0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x88,
	0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x08, 0x48, 0x02, 0x52, 0x05, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x42, 0x06, 0x0a,
	0x04, 0x5f, 0x61, 0x6c, 0x6c, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x42, 0x08, 0x0a, 0x06, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x22, 0xb1, 0x01, 0x0a, 0x16, 0x4d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x50, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x37, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x25, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x6c, 0x69, 0x62,
	0x72, 0x61, 0x72, 0x79, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x39,
	0x0a, 0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21,
	0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79,
	0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x52, 0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x19, 0x0a, 0x05, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x05, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x32, 0xbb,
	0x02, 0x0a, 0x11, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x49, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12, 0x25, 0x2e, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2e, 0x4d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x6c, 0x69, 0x62,
	0x72, 0x61, 0x72, 0x79, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x12,
	0x49, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x25, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73,
	0x68, 0x69, 0x70, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2e,
	0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x12, 0x48, 0x0a, 0x05, 0x50, 0x61,
	0x74, 0x63, 0x68, 0x12, 0x27, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x6c, 0x69,
	0x62, 0x72, 0x61, 0x72, 0x79, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70,
	0x50, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x12, 0x46, 0x0a, 0x05, 0x45, 0x72, 0x61, 0x73, 0x65, 0x12, 0x25, 0x2e,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2e,
	0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x38, 0x5a, 0x36,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x65, 0x73, 0x6f, 0x6d,
	0x6e, 0x75, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x6f, 0x72, 0x6d, 0x2f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x6c,
	0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var file_example_library_membership_svc_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_example_library_membership_svc_proto_goTypes = []any{
	(*MembershipAddRequest)(nil),   // 0: example.library.MembershipAddRequest
	(*MembershipGetRequest)(nil),   // 1: example.library.MembershipGetRequest
	(*MembershipSelect)(nil),       // 2: example.library.MembershipSelect
	(*MembershipPatchRequest)(nil), // 3: example.library.MembershipPatchRequest
	(*MemberGetRequest)(nil),       // 4: example.library.MemberGetRequest
	(*MemberSelect)(nil),           // 5: example.library.MemberSelect
	(*Membership)(nil),             // 6: example.library.Membership
	(*emptypb.Empty)(nil),          // 7: google.protobuf.Empty
}
var file_example_library_membership_svc_proto_depIdxs = []int32{
	4, // 0: example.library.MembershipAddRequest.member:type_name -> example.library.MemberGetRequest
	2, // 1: example.library.MembershipGetRequest.select:type_name -> example.library.MembershipSelect
	5, // 2: example.library.MembershipSelect.member:type_name -> example.library.MemberSelect
	1, // 3: example.library.MembershipPatchRequest.key:type_name -> example.library.MembershipGetRequest
	4, // 4: example.library.MembershipPatchRequest.member:type_name -> example.library.MemberGetRequest
	0, // 5: example.library.MembershipService.Add:input_type -> example.library.MembershipAddRequest
	1, // 6: example.library.MembershipService.Get:input_type -> example.library.MembershipGetRequest
	3, // 7: example.library.MembershipService.Patch:input_type -> example.library.MembershipPatchRequest
	1, // 8: example.library.MembershipService.Erase:input_type -> example.library.MembershipGetRequest
	6, // 9: example.library.MembershipService.Add:output_type -> example.library.Membership
	6, // 10: example.library.MembershipService.Get:output_type -> example.library.Membership
	7, // 11: example.library.MembershipService.Patch:output_type -> google.protobuf.Empty
	7, // 12: example.library.MembershipService.Erase:output_type -> google.protobuf.Empty
	9, // [9:13] is the sub-list for method output_type
	5, // [5:9] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_example_library_membership_svc_proto_init() }
func file_example_library_membership_svc_proto_init() {
	if File_example_library_membership_svc_proto != nil {
		return
	}
	file_example_library_member_svc_proto_init()
	file_example_library_membership_proto_init()
	file_example_library_membership_svc_proto_msgTypes[0].OneofWrappers = []any{}
	file_example_library_membership_svc_proto_msgTypes[1].OneofWrappers = []any{
		(*membershipGetRequest_Id)(nil),
	}
	file_example_library_membership_svc_proto_msgTypes[2].OneofWrappers = []any{}
	file_example_library_membership_svc_proto_msgTypes[3].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_example_library_membership_svc_proto_rawDesc), len(file_example_library_membership_svc_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_example_library_membership_svc_proto_goTypes,
		DependencyIndexes: file_example_library_membership_svc_proto_depIdxs,
		MessageInfos:      file_example_library_membership_svc_proto_msgTypes,
	}.Build()
	File_example_library_membership_svc_proto = out.File
	file_example_library_membership_svc_proto_goTypes = nil
	file_example_library_membership_svc_proto_depIdxs = nil
}
