// Code generated by "proto-orm-gen-ent". DO NOT EDIT

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.29.0
// source: example/library/loan.g.proto

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

type LoanAddRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          []byte                 `protobuf:"bytes,1,opt,name=id,proto3,oneof" json:"id,omitempty"`
	Subject     *BookGetRequest        `protobuf:"bytes,3,opt,name=subject,proto3" json:"subject,omitempty"`
	Borrower    *MemberGetRequest      `protobuf:"bytes,5,opt,name=borrower,proto3" json:"borrower,omitempty"`
	DateDue     *timestamppb.Timestamp `protobuf:"bytes,13,opt,name=date_due,json=dateDue,proto3" json:"date_due,omitempty"`
	DateReturn  *timestamppb.Timestamp `protobuf:"bytes,14,opt,name=date_return,json=dateReturn,proto3,oneof" json:"date_return,omitempty"`
	DateCreated *timestamppb.Timestamp `protobuf:"bytes,15,opt,name=date_created,json=dateCreated,proto3,oneof" json:"date_created,omitempty"`
}

func (x *LoanAddRequest) Reset() {
	*x = LoanAddRequest{}
	mi := &file_example_library_loan_g_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoanAddRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoanAddRequest) ProtoMessage() {}

func (x *LoanAddRequest) ProtoReflect() protoreflect.Message {
	mi := &file_example_library_loan_g_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoanAddRequest.ProtoReflect.Descriptor instead.
func (*LoanAddRequest) Descriptor() ([]byte, []int) {
	return file_example_library_loan_g_proto_rawDescGZIP(), []int{0}
}

func (x *LoanAddRequest) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *LoanAddRequest) GetSubject() *BookGetRequest {
	if x != nil {
		return x.Subject
	}
	return nil
}

func (x *LoanAddRequest) GetBorrower() *MemberGetRequest {
	if x != nil {
		return x.Borrower
	}
	return nil
}

func (x *LoanAddRequest) GetDateDue() *timestamppb.Timestamp {
	if x != nil {
		return x.DateDue
	}
	return nil
}

func (x *LoanAddRequest) GetDateReturn() *timestamppb.Timestamp {
	if x != nil {
		return x.DateReturn
	}
	return nil
}

func (x *LoanAddRequest) GetDateCreated() *timestamppb.Timestamp {
	if x != nil {
		return x.DateCreated
	}
	return nil
}

type LoanGetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Key:
	//
	//	*LoanGetRequest_Id
	Key isLoanGetRequest_Key `protobuf_oneof:"key"`
}

func (x *LoanGetRequest) Reset() {
	*x = LoanGetRequest{}
	mi := &file_example_library_loan_g_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoanGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoanGetRequest) ProtoMessage() {}

func (x *LoanGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_example_library_loan_g_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoanGetRequest.ProtoReflect.Descriptor instead.
func (*LoanGetRequest) Descriptor() ([]byte, []int) {
	return file_example_library_loan_g_proto_rawDescGZIP(), []int{1}
}

func (m *LoanGetRequest) GetKey() isLoanGetRequest_Key {
	if m != nil {
		return m.Key
	}
	return nil
}

func (x *LoanGetRequest) GetId() []byte {
	if x, ok := x.GetKey().(*LoanGetRequest_Id); ok {
		return x.Id
	}
	return nil
}

type isLoanGetRequest_Key interface {
	isLoanGetRequest_Key()
}

type LoanGetRequest_Id struct {
	Id []byte `protobuf:"bytes,1,opt,name=id,proto3,oneof"`
}

func (*LoanGetRequest_Id) isLoanGetRequest_Key() {}

type LoanPatchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key            *LoanGetRequest        `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	DateDue        *timestamppb.Timestamp `protobuf:"bytes,25,opt,name=date_due,json=dateDue,proto3,oneof" json:"date_due,omitempty"`
	DateReturn     *timestamppb.Timestamp `protobuf:"bytes,27,opt,name=date_return,json=dateReturn,proto3,oneof" json:"date_return,omitempty"`
	DateReturnNull bool                   `protobuf:"varint,28,opt,name=date_return_null,json=dateReturnNull,proto3" json:"date_return_null,omitempty"`
}

func (x *LoanPatchRequest) Reset() {
	*x = LoanPatchRequest{}
	mi := &file_example_library_loan_g_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoanPatchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoanPatchRequest) ProtoMessage() {}

func (x *LoanPatchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_example_library_loan_g_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoanPatchRequest.ProtoReflect.Descriptor instead.
func (*LoanPatchRequest) Descriptor() ([]byte, []int) {
	return file_example_library_loan_g_proto_rawDescGZIP(), []int{2}
}

func (x *LoanPatchRequest) GetKey() *LoanGetRequest {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *LoanPatchRequest) GetDateDue() *timestamppb.Timestamp {
	if x != nil {
		return x.DateDue
	}
	return nil
}

func (x *LoanPatchRequest) GetDateReturn() *timestamppb.Timestamp {
	if x != nil {
		return x.DateReturn
	}
	return nil
}

func (x *LoanPatchRequest) GetDateReturnNull() bool {
	if x != nil {
		return x.DateReturnNull
	}
	return false
}

var File_example_library_loan_g_proto protoreflect.FileDescriptor

var file_example_library_loan_g_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72,
	0x79, 0x2f, 0x6c, 0x6f, 0x61, 0x6e, 0x2e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x1a,
	0x1e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79,
	0x2f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x2e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1a, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79,
	0x2f, 0x6c, 0x6f, 0x61, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2f, 0x6d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x2e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70,
	0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x84, 0x03, 0x0a, 0x0e, 0x4c, 0x6f,
	0x61, 0x6e, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x13, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x02, 0x69, 0x64, 0x88, 0x01,
	0x01, 0x12, 0x39, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x6c, 0x69, 0x62,
	0x72, 0x61, 0x72, 0x79, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x3d, 0x0a, 0x08,
	0x62, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21,
	0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79,
	0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x52, 0x08, 0x62, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x65, 0x72, 0x12, 0x35, 0x0a, 0x08, 0x64,
	0x61, 0x74, 0x65, 0x5f, 0x64, 0x75, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x64, 0x61, 0x74, 0x65, 0x44,
	0x75, 0x65, 0x12, 0x40, 0x0a, 0x0b, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x72, 0x65, 0x74, 0x75, 0x72,
	0x6e, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x48, 0x01, 0x52, 0x0a, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x74, 0x75, 0x72,
	0x6e, 0x88, 0x01, 0x01, 0x12, 0x42, 0x0a, 0x0c, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x02, 0x52, 0x0b, 0x64, 0x61, 0x74, 0x65, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x88, 0x01, 0x01, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x69, 0x64, 0x42,
	0x0e, 0x0a, 0x0c, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x42,
	0x0f, 0x0a, 0x0d, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x22, 0x29, 0x0a, 0x0e, 0x4c, 0x6f, 0x61, 0x6e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x10, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00,
	0x52, 0x02, 0x69, 0x64, 0x42, 0x05, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x8a, 0x02, 0x0a, 0x10,
	0x4c, 0x6f, 0x61, 0x6e, 0x50, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x31, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2e,
	0x4c, 0x6f, 0x61, 0x6e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x3a, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x64, 0x75, 0x65, 0x18,
	0x19, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x48, 0x00, 0x52, 0x07, 0x64, 0x61, 0x74, 0x65, 0x44, 0x75, 0x65, 0x88, 0x01, 0x01, 0x12,
	0x40, 0x0a, 0x0b, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x18, 0x1b,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x48, 0x01, 0x52, 0x0a, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x88, 0x01,
	0x01, 0x12, 0x28, 0x0a, 0x10, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e,
	0x5f, 0x6e, 0x75, 0x6c, 0x6c, 0x18, 0x1c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x64, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x4e, 0x75, 0x6c, 0x6c, 0x42, 0x0b, 0x0a, 0x09, 0x5f,
	0x64, 0x61, 0x74, 0x65, 0x5f, 0x64, 0x75, 0x65, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x64, 0x61, 0x74,
	0x65, 0x5f, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x32, 0x91, 0x02, 0x0a, 0x0b, 0x4c, 0x6f, 0x61,
	0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3d, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12,
	0x1f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72,
	0x79, 0x2e, 0x4c, 0x6f, 0x61, 0x6e, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x15, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61,
	0x72, 0x79, 0x2e, 0x4c, 0x6f, 0x61, 0x6e, 0x12, 0x3d, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x1f,
	0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79,
	0x2e, 0x4c, 0x6f, 0x61, 0x6e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x15, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72,
	0x79, 0x2e, 0x4c, 0x6f, 0x61, 0x6e, 0x12, 0x42, 0x0a, 0x05, 0x50, 0x61, 0x74, 0x63, 0x68, 0x12,
	0x21, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72,
	0x79, 0x2e, 0x4c, 0x6f, 0x61, 0x6e, 0x50, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x40, 0x0a, 0x05, 0x45, 0x72,
	0x61, 0x73, 0x65, 0x12, 0x1f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x6c, 0x69,
	0x62, 0x72, 0x61, 0x72, 0x79, 0x2e, 0x4c, 0x6f, 0x61, 0x6e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x38, 0x5a, 0x36,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x65, 0x73, 0x6f, 0x6d,
	0x6e, 0x75, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x6f, 0x72, 0x6d, 0x2f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x6c,
	0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_example_library_loan_g_proto_rawDescOnce sync.Once
	file_example_library_loan_g_proto_rawDescData = file_example_library_loan_g_proto_rawDesc
)

func file_example_library_loan_g_proto_rawDescGZIP() []byte {
	file_example_library_loan_g_proto_rawDescOnce.Do(func() {
		file_example_library_loan_g_proto_rawDescData = protoimpl.X.CompressGZIP(file_example_library_loan_g_proto_rawDescData)
	})
	return file_example_library_loan_g_proto_rawDescData
}

var file_example_library_loan_g_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_example_library_loan_g_proto_goTypes = []any{
	(*LoanAddRequest)(nil),        // 0: example.library.LoanAddRequest
	(*LoanGetRequest)(nil),        // 1: example.library.LoanGetRequest
	(*LoanPatchRequest)(nil),      // 2: example.library.LoanPatchRequest
	(*BookGetRequest)(nil),        // 3: example.library.BookGetRequest
	(*MemberGetRequest)(nil),      // 4: example.library.MemberGetRequest
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
	(*Loan)(nil),                  // 6: example.library.Loan
	(*emptypb.Empty)(nil),         // 7: google.protobuf.Empty
}
var file_example_library_loan_g_proto_depIdxs = []int32{
	3,  // 0: example.library.LoanAddRequest.subject:type_name -> example.library.BookGetRequest
	4,  // 1: example.library.LoanAddRequest.borrower:type_name -> example.library.MemberGetRequest
	5,  // 2: example.library.LoanAddRequest.date_due:type_name -> google.protobuf.Timestamp
	5,  // 3: example.library.LoanAddRequest.date_return:type_name -> google.protobuf.Timestamp
	5,  // 4: example.library.LoanAddRequest.date_created:type_name -> google.protobuf.Timestamp
	1,  // 5: example.library.LoanPatchRequest.key:type_name -> example.library.LoanGetRequest
	5,  // 6: example.library.LoanPatchRequest.date_due:type_name -> google.protobuf.Timestamp
	5,  // 7: example.library.LoanPatchRequest.date_return:type_name -> google.protobuf.Timestamp
	0,  // 8: example.library.LoanService.Add:input_type -> example.library.LoanAddRequest
	1,  // 9: example.library.LoanService.Get:input_type -> example.library.LoanGetRequest
	2,  // 10: example.library.LoanService.Patch:input_type -> example.library.LoanPatchRequest
	1,  // 11: example.library.LoanService.Erase:input_type -> example.library.LoanGetRequest
	6,  // 12: example.library.LoanService.Add:output_type -> example.library.Loan
	6,  // 13: example.library.LoanService.Get:output_type -> example.library.Loan
	7,  // 14: example.library.LoanService.Patch:output_type -> google.protobuf.Empty
	7,  // 15: example.library.LoanService.Erase:output_type -> google.protobuf.Empty
	12, // [12:16] is the sub-list for method output_type
	8,  // [8:12] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_example_library_loan_g_proto_init() }
func file_example_library_loan_g_proto_init() {
	if File_example_library_loan_g_proto != nil {
		return
	}
	file_example_library_author_g_proto_init()
	file_example_library_loan_proto_init()
	file_example_library_member_g_proto_init()
	file_example_library_loan_g_proto_msgTypes[0].OneofWrappers = []any{}
	file_example_library_loan_g_proto_msgTypes[1].OneofWrappers = []any{
		(*LoanGetRequest_Id)(nil),
	}
	file_example_library_loan_g_proto_msgTypes[2].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_example_library_loan_g_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_example_library_loan_g_proto_goTypes,
		DependencyIndexes: file_example_library_loan_g_proto_depIdxs,
		MessageInfos:      file_example_library_loan_g_proto_msgTypes,
	}.Build()
	File_example_library_loan_g_proto = out.File
	file_example_library_loan_g_proto_rawDesc = nil
	file_example_library_loan_g_proto_goTypes = nil
	file_example_library_loan_g_proto_depIdxs = nil
}