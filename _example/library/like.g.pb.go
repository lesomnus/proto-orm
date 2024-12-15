// Code generated by "proto-orm-gen-ent". DO NOT EDIT

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.29.0
// source: example/library/like.g.proto

package library

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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

type LikeAddRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          []byte                 `protobuf:"bytes,1,opt,name=id,proto3,oneof" json:"id,omitempty"`
	Subject     *BookGetRequest        `protobuf:"bytes,3,opt,name=subject,proto3" json:"subject,omitempty"`
	Actor       *MemberGetRequest      `protobuf:"bytes,5,opt,name=actor,proto3" json:"actor,omitempty"`
	DateCreated *timestamppb.Timestamp `protobuf:"bytes,15,opt,name=date_created,json=dateCreated,proto3,oneof" json:"date_created,omitempty"`
}

func (x *LikeAddRequest) Reset() {
	*x = LikeAddRequest{}
	mi := &file_example_library_like_g_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LikeAddRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LikeAddRequest) ProtoMessage() {}

func (x *LikeAddRequest) ProtoReflect() protoreflect.Message {
	mi := &file_example_library_like_g_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LikeAddRequest.ProtoReflect.Descriptor instead.
func (*LikeAddRequest) Descriptor() ([]byte, []int) {
	return file_example_library_like_g_proto_rawDescGZIP(), []int{0}
}

func (x *LikeAddRequest) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *LikeAddRequest) GetSubject() *BookGetRequest {
	if x != nil {
		return x.Subject
	}
	return nil
}

func (x *LikeAddRequest) GetActor() *MemberGetRequest {
	if x != nil {
		return x.Actor
	}
	return nil
}

func (x *LikeAddRequest) GetDateCreated() *timestamppb.Timestamp {
	if x != nil {
		return x.DateCreated
	}
	return nil
}

type LikeGetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Key:
	//
	//	*LikeGetRequest_Id
	//	*LikeGetRequest_Holders
	Key isLikeGetRequest_Key `protobuf_oneof:"key"`
}

func (x *LikeGetRequest) Reset() {
	*x = LikeGetRequest{}
	mi := &file_example_library_like_g_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LikeGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LikeGetRequest) ProtoMessage() {}

func (x *LikeGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_example_library_like_g_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LikeGetRequest.ProtoReflect.Descriptor instead.
func (*LikeGetRequest) Descriptor() ([]byte, []int) {
	return file_example_library_like_g_proto_rawDescGZIP(), []int{1}
}

func (m *LikeGetRequest) GetKey() isLikeGetRequest_Key {
	if m != nil {
		return m.Key
	}
	return nil
}

func (x *LikeGetRequest) GetId() []byte {
	if x, ok := x.GetKey().(*LikeGetRequest_Id); ok {
		return x.Id
	}
	return nil
}

func (x *LikeGetRequest) GetHolders() *LikeGetByHolders {
	if x, ok := x.GetKey().(*LikeGetRequest_Holders); ok {
		return x.Holders
	}
	return nil
}

type isLikeGetRequest_Key interface {
	isLikeGetRequest_Key()
}

type LikeGetRequest_Id struct {
	Id []byte `protobuf:"bytes,1,opt,name=id,proto3,oneof"`
}

type LikeGetRequest_Holders struct {
	Holders *LikeGetByHolders `protobuf:"bytes,2,opt,name=holders,proto3,oneof"`
}

func (*LikeGetRequest_Id) isLikeGetRequest_Key() {}

func (*LikeGetRequest_Holders) isLikeGetRequest_Key() {}

type LikeGetByHolders struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SubjectId []byte            `protobuf:"bytes,2,opt,name=subject_id,json=subjectId,proto3" json:"subject_id,omitempty"`
	Actor     *MemberGetRequest `protobuf:"bytes,5,opt,name=actor,proto3" json:"actor,omitempty"`
}

func (x *LikeGetByHolders) Reset() {
	*x = LikeGetByHolders{}
	mi := &file_example_library_like_g_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LikeGetByHolders) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LikeGetByHolders) ProtoMessage() {}

func (x *LikeGetByHolders) ProtoReflect() protoreflect.Message {
	mi := &file_example_library_like_g_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LikeGetByHolders.ProtoReflect.Descriptor instead.
func (*LikeGetByHolders) Descriptor() ([]byte, []int) {
	return file_example_library_like_g_proto_rawDescGZIP(), []int{2}
}

func (x *LikeGetByHolders) GetSubjectId() []byte {
	if x != nil {
		return x.SubjectId
	}
	return nil
}

func (x *LikeGetByHolders) GetActor() *MemberGetRequest {
	if x != nil {
		return x.Actor
	}
	return nil
}

type LikePatchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key *LikeGetRequest `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *LikePatchRequest) Reset() {
	*x = LikePatchRequest{}
	mi := &file_example_library_like_g_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LikePatchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LikePatchRequest) ProtoMessage() {}

func (x *LikePatchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_example_library_like_g_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LikePatchRequest.ProtoReflect.Descriptor instead.
func (*LikePatchRequest) Descriptor() ([]byte, []int) {
	return file_example_library_like_g_proto_rawDescGZIP(), []int{3}
}

func (x *LikePatchRequest) GetKey() *LikeGetRequest {
	if x != nil {
		return x.Key
	}
	return nil
}

var File_example_library_like_g_proto protoreflect.FileDescriptor

var file_example_library_like_g_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72,
	0x79, 0x2f, 0x6c, 0x69, 0x6b, 0x65, 0x2e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x1a,
	0x1e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79,
	0x2f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x2e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1a, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79,
	0x2f, 0x6c, 0x69, 0x6b, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2f, 0x6d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x2e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70,
	0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf5, 0x01, 0x0a, 0x0e, 0x4c, 0x69,
	0x6b, 0x65, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x13, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x02, 0x69, 0x64, 0x88, 0x01,
	0x01, 0x12, 0x39, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x6c, 0x69, 0x62,
	0x72, 0x61, 0x72, 0x79, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x37, 0x0a, 0x05,
	0x61, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2e, 0x4d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x05,
	0x61, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x42, 0x0a, 0x0c, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x01, 0x52, 0x0b, 0x64, 0x61, 0x74, 0x65, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x88, 0x01, 0x01, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x69, 0x64,
	0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x22, 0x68, 0x0a, 0x0e, 0x4c, 0x69, 0x6b, 0x65, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x48,
	0x00, 0x52, 0x02, 0x69, 0x64, 0x12, 0x3d, 0x0a, 0x07, 0x68, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2e, 0x4c, 0x69, 0x6b, 0x65, 0x47, 0x65, 0x74,
	0x42, 0x79, 0x48, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x73, 0x48, 0x00, 0x52, 0x07, 0x68, 0x6f, 0x6c,
	0x64, 0x65, 0x72, 0x73, 0x42, 0x05, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x6a, 0x0a, 0x10, 0x4c,
	0x69, 0x6b, 0x65, 0x47, 0x65, 0x74, 0x42, 0x79, 0x48, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x73, 0x12,
	0x1d, 0x0a, 0x0a, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x37,
	0x0a, 0x05, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2e,
	0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x52, 0x05, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x22, 0x45, 0x0a, 0x10, 0x4c, 0x69, 0x6b, 0x65, 0x50,
	0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2e, 0x4c, 0x69, 0x6b, 0x65, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x32, 0x91,
	0x02, 0x0a, 0x0b, 0x4c, 0x69, 0x6b, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3d,
	0x0a, 0x03, 0x41, 0x64, 0x64, 0x12, 0x1f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e,
	0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2e, 0x4c, 0x69, 0x6b, 0x65, 0x41, 0x64, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2e, 0x4c, 0x69, 0x6b, 0x65, 0x12, 0x3d, 0x0a,
	0x03, 0x47, 0x65, 0x74, 0x12, 0x1f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x6c,
	0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2e, 0x4c, 0x69, 0x6b, 0x65, 0x47, 0x65, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e,
	0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2e, 0x4c, 0x69, 0x6b, 0x65, 0x12, 0x42, 0x0a, 0x05,
	0x50, 0x61, 0x74, 0x63, 0x68, 0x12, 0x21, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e,
	0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2e, 0x4c, 0x69, 0x6b, 0x65, 0x50, 0x61, 0x74, 0x63,
	0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x12, 0x40, 0x0a, 0x05, 0x45, 0x72, 0x61, 0x73, 0x65, 0x12, 0x1f, 0x2e, 0x65, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2e, 0x4c, 0x69, 0x6b, 0x65,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x42, 0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6c, 0x65, 0x73, 0x6f, 0x6d, 0x6e, 0x75, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d,
	0x6f, 0x72, 0x6d, 0x2f, 0x5f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x6c, 0x69, 0x62,
	0x72, 0x61, 0x72, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_example_library_like_g_proto_rawDescOnce sync.Once
	file_example_library_like_g_proto_rawDescData = file_example_library_like_g_proto_rawDesc
)

func file_example_library_like_g_proto_rawDescGZIP() []byte {
	file_example_library_like_g_proto_rawDescOnce.Do(func() {
		file_example_library_like_g_proto_rawDescData = protoimpl.X.CompressGZIP(file_example_library_like_g_proto_rawDescData)
	})
	return file_example_library_like_g_proto_rawDescData
}

var file_example_library_like_g_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_example_library_like_g_proto_goTypes = []any{
	(*LikeAddRequest)(nil),        // 0: example.library.LikeAddRequest
	(*LikeGetRequest)(nil),        // 1: example.library.LikeGetRequest
	(*LikeGetByHolders)(nil),      // 2: example.library.LikeGetByHolders
	(*LikePatchRequest)(nil),      // 3: example.library.LikePatchRequest
	(*BookGetRequest)(nil),        // 4: example.library.BookGetRequest
	(*MemberGetRequest)(nil),      // 5: example.library.MemberGetRequest
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
	(*Like)(nil),                  // 7: example.library.Like
	(*emptypb.Empty)(nil),         // 8: google.protobuf.Empty
}
var file_example_library_like_g_proto_depIdxs = []int32{
	4,  // 0: example.library.LikeAddRequest.subject:type_name -> example.library.BookGetRequest
	5,  // 1: example.library.LikeAddRequest.actor:type_name -> example.library.MemberGetRequest
	6,  // 2: example.library.LikeAddRequest.date_created:type_name -> google.protobuf.Timestamp
	2,  // 3: example.library.LikeGetRequest.holders:type_name -> example.library.LikeGetByHolders
	5,  // 4: example.library.LikeGetByHolders.actor:type_name -> example.library.MemberGetRequest
	1,  // 5: example.library.LikePatchRequest.key:type_name -> example.library.LikeGetRequest
	0,  // 6: example.library.LikeService.Add:input_type -> example.library.LikeAddRequest
	1,  // 7: example.library.LikeService.Get:input_type -> example.library.LikeGetRequest
	3,  // 8: example.library.LikeService.Patch:input_type -> example.library.LikePatchRequest
	1,  // 9: example.library.LikeService.Erase:input_type -> example.library.LikeGetRequest
	7,  // 10: example.library.LikeService.Add:output_type -> example.library.Like
	7,  // 11: example.library.LikeService.Get:output_type -> example.library.Like
	8,  // 12: example.library.LikeService.Patch:output_type -> google.protobuf.Empty
	8,  // 13: example.library.LikeService.Erase:output_type -> google.protobuf.Empty
	10, // [10:14] is the sub-list for method output_type
	6,  // [6:10] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_example_library_like_g_proto_init() }
func file_example_library_like_g_proto_init() {
	if File_example_library_like_g_proto != nil {
		return
	}
	file_example_library_author_g_proto_init()
	file_example_library_like_proto_init()
	file_example_library_member_g_proto_init()
	file_example_library_like_g_proto_msgTypes[0].OneofWrappers = []any{}
	file_example_library_like_g_proto_msgTypes[1].OneofWrappers = []any{
		(*LikeGetRequest_Id)(nil),
		(*LikeGetRequest_Holders)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_example_library_like_g_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_example_library_like_g_proto_goTypes,
		DependencyIndexes: file_example_library_like_g_proto_depIdxs,
		MessageInfos:      file_example_library_like_g_proto_msgTypes,
	}.Build()
	File_example_library_like_g_proto = out.File
	file_example_library_like_g_proto_rawDesc = nil
	file_example_library_like_g_proto_goTypes = nil
	file_example_library_like_g_proto_depIdxs = nil
}
