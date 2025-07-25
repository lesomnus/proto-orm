// Code generated by "protoc-gen-orm-service", DO NOT EDIT.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        v5.29.0
// source: example/library/loan_svc.proto

package library

import (
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

type LoanAddRequest struct {
	state                  protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Id          []byte                 `protobuf:"bytes,1,opt,name=id,proto3,oneof" json:"id,omitempty"`
	xxx_hidden_Subject     *BookGetRequest        `protobuf:"bytes,3,opt,name=subject,proto3" json:"subject,omitempty"`
	xxx_hidden_Borrower    *MemberGetRequest      `protobuf:"bytes,5,opt,name=borrower,proto3" json:"borrower,omitempty"`
	xxx_hidden_Policy      string                 `protobuf:"bytes,6,opt,name=policy,proto3" json:"policy,omitempty"`
	xxx_hidden_DateDue     *timestamppb.Timestamp `protobuf:"bytes,13,opt,name=date_due,json=dateDue,proto3" json:"date_due,omitempty"`
	xxx_hidden_DateReturn  *timestamppb.Timestamp `protobuf:"bytes,14,opt,name=date_return,json=dateReturn,proto3,oneof" json:"date_return,omitempty"`
	xxx_hidden_DateCreated *timestamppb.Timestamp `protobuf:"bytes,15,opt,name=date_created,json=dateCreated,proto3,oneof" json:"date_created,omitempty"`
	XXX_raceDetectHookData protoimpl.RaceDetectHookData
	XXX_presence           [1]uint32
	unknownFields          protoimpl.UnknownFields
	sizeCache              protoimpl.SizeCache
}

func (x *LoanAddRequest) Reset() {
	*x = LoanAddRequest{}
	mi := &file_example_library_loan_svc_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoanAddRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoanAddRequest) ProtoMessage() {}

func (x *LoanAddRequest) ProtoReflect() protoreflect.Message {
	mi := &file_example_library_loan_svc_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *LoanAddRequest) GetId() []byte {
	if x != nil {
		return x.xxx_hidden_Id
	}
	return nil
}

func (x *LoanAddRequest) GetSubject() *BookGetRequest {
	if x != nil {
		return x.xxx_hidden_Subject
	}
	return nil
}

func (x *LoanAddRequest) GetBorrower() *MemberGetRequest {
	if x != nil {
		return x.xxx_hidden_Borrower
	}
	return nil
}

func (x *LoanAddRequest) GetPolicy() string {
	if x != nil {
		return x.xxx_hidden_Policy
	}
	return ""
}

func (x *LoanAddRequest) GetDateDue() *timestamppb.Timestamp {
	if x != nil {
		return x.xxx_hidden_DateDue
	}
	return nil
}

func (x *LoanAddRequest) GetDateReturn() *timestamppb.Timestamp {
	if x != nil {
		return x.xxx_hidden_DateReturn
	}
	return nil
}

func (x *LoanAddRequest) GetDateCreated() *timestamppb.Timestamp {
	if x != nil {
		return x.xxx_hidden_DateCreated
	}
	return nil
}

func (x *LoanAddRequest) SetId(v []byte) {
	if v == nil {
		v = []byte{}
	}
	x.xxx_hidden_Id = v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 0, 7)
}

func (x *LoanAddRequest) SetSubject(v *BookGetRequest) {
	x.xxx_hidden_Subject = v
}

func (x *LoanAddRequest) SetBorrower(v *MemberGetRequest) {
	x.xxx_hidden_Borrower = v
}

func (x *LoanAddRequest) SetPolicy(v string) {
	x.xxx_hidden_Policy = v
}

func (x *LoanAddRequest) SetDateDue(v *timestamppb.Timestamp) {
	x.xxx_hidden_DateDue = v
}

func (x *LoanAddRequest) SetDateReturn(v *timestamppb.Timestamp) {
	x.xxx_hidden_DateReturn = v
}

func (x *LoanAddRequest) SetDateCreated(v *timestamppb.Timestamp) {
	x.xxx_hidden_DateCreated = v
}

func (x *LoanAddRequest) HasId() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 0)
}

func (x *LoanAddRequest) HasSubject() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_Subject != nil
}

func (x *LoanAddRequest) HasBorrower() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_Borrower != nil
}

func (x *LoanAddRequest) HasDateDue() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_DateDue != nil
}

func (x *LoanAddRequest) HasDateReturn() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_DateReturn != nil
}

func (x *LoanAddRequest) HasDateCreated() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_DateCreated != nil
}

func (x *LoanAddRequest) ClearId() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 0)
	x.xxx_hidden_Id = nil
}

func (x *LoanAddRequest) ClearSubject() {
	x.xxx_hidden_Subject = nil
}

func (x *LoanAddRequest) ClearBorrower() {
	x.xxx_hidden_Borrower = nil
}

func (x *LoanAddRequest) ClearDateDue() {
	x.xxx_hidden_DateDue = nil
}

func (x *LoanAddRequest) ClearDateReturn() {
	x.xxx_hidden_DateReturn = nil
}

func (x *LoanAddRequest) ClearDateCreated() {
	x.xxx_hidden_DateCreated = nil
}

type LoanAddRequest_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Id          []byte
	Subject     *BookGetRequest
	Borrower    *MemberGetRequest
	Policy      string
	DateDue     *timestamppb.Timestamp
	DateReturn  *timestamppb.Timestamp
	DateCreated *timestamppb.Timestamp
}

func (b0 LoanAddRequest_builder) Build() *LoanAddRequest {
	m0 := &LoanAddRequest{}
	b, x := &b0, m0
	_, _ = b, x
	if b.Id != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 0, 7)
		x.xxx_hidden_Id = b.Id
	}
	x.xxx_hidden_Subject = b.Subject
	x.xxx_hidden_Borrower = b.Borrower
	x.xxx_hidden_Policy = b.Policy
	x.xxx_hidden_DateDue = b.DateDue
	x.xxx_hidden_DateReturn = b.DateReturn
	x.xxx_hidden_DateCreated = b.DateCreated
	return m0
}

type LoanGetRequest struct {
	state             protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Select *LoanSelect            `protobuf:"bytes,1,opt,name=select,proto3" json:"select,omitempty"`
	xxx_hidden_Key    isLoanGetRequest_Key   `protobuf_oneof:"key"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *LoanGetRequest) Reset() {
	*x = LoanGetRequest{}
	mi := &file_example_library_loan_svc_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoanGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoanGetRequest) ProtoMessage() {}

func (x *LoanGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_example_library_loan_svc_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *LoanGetRequest) GetSelect() *LoanSelect {
	if x != nil {
		return x.xxx_hidden_Select
	}
	return nil
}

func (x *LoanGetRequest) GetId() []byte {
	if x != nil {
		if x, ok := x.xxx_hidden_Key.(*loanGetRequest_Id); ok {
			return x.Id
		}
	}
	return nil
}

func (x *LoanGetRequest) SetSelect(v *LoanSelect) {
	x.xxx_hidden_Select = v
}

func (x *LoanGetRequest) SetId(v []byte) {
	if v == nil {
		v = []byte{}
	}
	x.xxx_hidden_Key = &loanGetRequest_Id{v}
}

func (x *LoanGetRequest) HasSelect() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_Select != nil
}

func (x *LoanGetRequest) HasKey() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_Key != nil
}

func (x *LoanGetRequest) HasId() bool {
	if x == nil {
		return false
	}
	_, ok := x.xxx_hidden_Key.(*loanGetRequest_Id)
	return ok
}

func (x *LoanGetRequest) ClearSelect() {
	x.xxx_hidden_Select = nil
}

func (x *LoanGetRequest) ClearKey() {
	x.xxx_hidden_Key = nil
}

func (x *LoanGetRequest) ClearId() {
	if _, ok := x.xxx_hidden_Key.(*loanGetRequest_Id); ok {
		x.xxx_hidden_Key = nil
	}
}

const LoanGetRequest_Key_not_set_case case_LoanGetRequest_Key = 0
const LoanGetRequest_Id_case case_LoanGetRequest_Key = 2

func (x *LoanGetRequest) WhichKey() case_LoanGetRequest_Key {
	if x == nil {
		return LoanGetRequest_Key_not_set_case
	}
	switch x.xxx_hidden_Key.(type) {
	case *loanGetRequest_Id:
		return LoanGetRequest_Id_case
	default:
		return LoanGetRequest_Key_not_set_case
	}
}

type LoanGetRequest_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Select *LoanSelect
	// Fields of oneof xxx_hidden_Key:
	Id []byte
	// -- end of xxx_hidden_Key
}

func (b0 LoanGetRequest_builder) Build() *LoanGetRequest {
	m0 := &LoanGetRequest{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_Select = b.Select
	if b.Id != nil {
		x.xxx_hidden_Key = &loanGetRequest_Id{b.Id}
	}
	return m0
}

type case_LoanGetRequest_Key protoreflect.FieldNumber

func (x case_LoanGetRequest_Key) String() string {
	md := file_example_library_loan_svc_proto_msgTypes[1].Descriptor()
	if x == 0 {
		return "not set"
	}
	return protoimpl.X.MessageFieldStringOf(md, protoreflect.FieldNumber(x))
}

type isLoanGetRequest_Key interface {
	isLoanGetRequest_Key()
}

type loanGetRequest_Id struct {
	Id []byte `protobuf:"bytes,2,opt,name=id,proto3,oneof"`
}

func (*loanGetRequest_Id) isLoanGetRequest_Key() {}

type LoanSelect struct {
	state                  protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_All         bool                   `protobuf:"varint,1,opt,name=all,proto3,oneof" json:"all,omitempty"`
	xxx_hidden_SubjectId   bool                   `protobuf:"varint,2,opt,name=subject_id,json=subjectId,proto3,oneof" json:"subject_id,omitempty"`
	xxx_hidden_Subject     *BookSelect            `protobuf:"bytes,3,opt,name=subject,proto3,oneof" json:"subject,omitempty"`
	xxx_hidden_BorrowerId  bool                   `protobuf:"varint,4,opt,name=borrower_id,json=borrowerId,proto3,oneof" json:"borrower_id,omitempty"`
	xxx_hidden_Borrower    *MemberSelect          `protobuf:"bytes,5,opt,name=borrower,proto3,oneof" json:"borrower,omitempty"`
	xxx_hidden_Policy      bool                   `protobuf:"varint,6,opt,name=policy,proto3,oneof" json:"policy,omitempty"`
	xxx_hidden_DateDue     bool                   `protobuf:"varint,13,opt,name=date_due,json=dateDue,proto3,oneof" json:"date_due,omitempty"`
	xxx_hidden_DateReturn  bool                   `protobuf:"varint,14,opt,name=date_return,json=dateReturn,proto3,oneof" json:"date_return,omitempty"`
	xxx_hidden_DateCreated bool                   `protobuf:"varint,15,opt,name=date_created,json=dateCreated,proto3,oneof" json:"date_created,omitempty"`
	XXX_raceDetectHookData protoimpl.RaceDetectHookData
	XXX_presence           [1]uint32
	unknownFields          protoimpl.UnknownFields
	sizeCache              protoimpl.SizeCache
}

func (x *LoanSelect) Reset() {
	*x = LoanSelect{}
	mi := &file_example_library_loan_svc_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoanSelect) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoanSelect) ProtoMessage() {}

func (x *LoanSelect) ProtoReflect() protoreflect.Message {
	mi := &file_example_library_loan_svc_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *LoanSelect) GetAll() bool {
	if x != nil {
		return x.xxx_hidden_All
	}
	return false
}

func (x *LoanSelect) GetSubjectId() bool {
	if x != nil {
		return x.xxx_hidden_SubjectId
	}
	return false
}

func (x *LoanSelect) GetSubject() *BookSelect {
	if x != nil {
		return x.xxx_hidden_Subject
	}
	return nil
}

func (x *LoanSelect) GetBorrowerId() bool {
	if x != nil {
		return x.xxx_hidden_BorrowerId
	}
	return false
}

func (x *LoanSelect) GetBorrower() *MemberSelect {
	if x != nil {
		return x.xxx_hidden_Borrower
	}
	return nil
}

func (x *LoanSelect) GetPolicy() bool {
	if x != nil {
		return x.xxx_hidden_Policy
	}
	return false
}

func (x *LoanSelect) GetDateDue() bool {
	if x != nil {
		return x.xxx_hidden_DateDue
	}
	return false
}

func (x *LoanSelect) GetDateReturn() bool {
	if x != nil {
		return x.xxx_hidden_DateReturn
	}
	return false
}

func (x *LoanSelect) GetDateCreated() bool {
	if x != nil {
		return x.xxx_hidden_DateCreated
	}
	return false
}

func (x *LoanSelect) SetAll(v bool) {
	x.xxx_hidden_All = v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 0, 9)
}

func (x *LoanSelect) SetSubjectId(v bool) {
	x.xxx_hidden_SubjectId = v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 1, 9)
}

func (x *LoanSelect) SetSubject(v *BookSelect) {
	x.xxx_hidden_Subject = v
}

func (x *LoanSelect) SetBorrowerId(v bool) {
	x.xxx_hidden_BorrowerId = v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 3, 9)
}

func (x *LoanSelect) SetBorrower(v *MemberSelect) {
	x.xxx_hidden_Borrower = v
}

func (x *LoanSelect) SetPolicy(v bool) {
	x.xxx_hidden_Policy = v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 5, 9)
}

func (x *LoanSelect) SetDateDue(v bool) {
	x.xxx_hidden_DateDue = v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 6, 9)
}

func (x *LoanSelect) SetDateReturn(v bool) {
	x.xxx_hidden_DateReturn = v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 7, 9)
}

func (x *LoanSelect) SetDateCreated(v bool) {
	x.xxx_hidden_DateCreated = v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 8, 9)
}

func (x *LoanSelect) HasAll() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 0)
}

func (x *LoanSelect) HasSubjectId() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 1)
}

func (x *LoanSelect) HasSubject() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_Subject != nil
}

func (x *LoanSelect) HasBorrowerId() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 3)
}

func (x *LoanSelect) HasBorrower() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_Borrower != nil
}

func (x *LoanSelect) HasPolicy() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 5)
}

func (x *LoanSelect) HasDateDue() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 6)
}

func (x *LoanSelect) HasDateReturn() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 7)
}

func (x *LoanSelect) HasDateCreated() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 8)
}

func (x *LoanSelect) ClearAll() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 0)
	x.xxx_hidden_All = false
}

func (x *LoanSelect) ClearSubjectId() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 1)
	x.xxx_hidden_SubjectId = false
}

func (x *LoanSelect) ClearSubject() {
	x.xxx_hidden_Subject = nil
}

func (x *LoanSelect) ClearBorrowerId() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 3)
	x.xxx_hidden_BorrowerId = false
}

func (x *LoanSelect) ClearBorrower() {
	x.xxx_hidden_Borrower = nil
}

func (x *LoanSelect) ClearPolicy() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 5)
	x.xxx_hidden_Policy = false
}

func (x *LoanSelect) ClearDateDue() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 6)
	x.xxx_hidden_DateDue = false
}

func (x *LoanSelect) ClearDateReturn() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 7)
	x.xxx_hidden_DateReturn = false
}

func (x *LoanSelect) ClearDateCreated() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 8)
	x.xxx_hidden_DateCreated = false
}

type LoanSelect_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	All         *bool
	SubjectId   *bool
	Subject     *BookSelect
	BorrowerId  *bool
	Borrower    *MemberSelect
	Policy      *bool
	DateDue     *bool
	DateReturn  *bool
	DateCreated *bool
}

func (b0 LoanSelect_builder) Build() *LoanSelect {
	m0 := &LoanSelect{}
	b, x := &b0, m0
	_, _ = b, x
	if b.All != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 0, 9)
		x.xxx_hidden_All = *b.All
	}
	if b.SubjectId != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 1, 9)
		x.xxx_hidden_SubjectId = *b.SubjectId
	}
	x.xxx_hidden_Subject = b.Subject
	if b.BorrowerId != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 3, 9)
		x.xxx_hidden_BorrowerId = *b.BorrowerId
	}
	x.xxx_hidden_Borrower = b.Borrower
	if b.Policy != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 5, 9)
		x.xxx_hidden_Policy = *b.Policy
	}
	if b.DateDue != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 6, 9)
		x.xxx_hidden_DateDue = *b.DateDue
	}
	if b.DateReturn != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 7, 9)
		x.xxx_hidden_DateReturn = *b.DateReturn
	}
	if b.DateCreated != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 8, 9)
		x.xxx_hidden_DateCreated = *b.DateCreated
	}
	return m0
}

var File_example_library_loan_svc_proto protoreflect.FileDescriptor

var file_example_library_loan_svc_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72,
	0x79, 0x2f, 0x6c, 0x6f, 0x61, 0x6e, 0x5f, 0x73, 0x76, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72,
	0x79, 0x1a, 0x1e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x6c, 0x69, 0x62, 0x72, 0x61,
	0x72, 0x79, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x73, 0x76, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1a, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x6c, 0x69, 0x62, 0x72, 0x61,
	0x72, 0x79, 0x2f, 0x6c, 0x6f, 0x61, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x65,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2f, 0x6d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x73, 0x76, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x9c, 0x03, 0x0a, 0x0e, 0x4c, 0x6f, 0x61, 0x6e, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x13, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x48,
	0x00, 0x52, 0x02, 0x69, 0x64, 0x88, 0x01, 0x01, 0x12, 0x39, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x65, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2e, 0x42, 0x6f, 0x6f, 0x6b,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x12, 0x3d, 0x0a, 0x08, 0x62, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x65, 0x72, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e,
	0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x08, 0x62, 0x6f, 0x72, 0x72, 0x6f, 0x77,
	0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x35, 0x0a, 0x08, 0x64, 0x61,
	0x74, 0x65, 0x5f, 0x64, 0x75, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x64, 0x61, 0x74, 0x65, 0x44, 0x75,
	0x65, 0x12, 0x40, 0x0a, 0x0b, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e,
	0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x48, 0x01, 0x52, 0x0a, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e,
	0x88, 0x01, 0x01, 0x12, 0x42, 0x0a, 0x0c, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x02, 0x52, 0x0b, 0x64, 0x61, 0x74, 0x65, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x88, 0x01, 0x01, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x69, 0x64, 0x42, 0x0e,
	0x0a, 0x0c, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x42, 0x0f,
	0x0a, 0x0d, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x22,
	0x5e, 0x0a, 0x0e, 0x4c, 0x6f, 0x61, 0x6e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x33, 0x0a, 0x06, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x6c, 0x69, 0x62, 0x72,
	0x61, 0x72, 0x79, 0x2e, 0x4c, 0x6f, 0x61, 0x6e, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x52, 0x06,
	0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x12, 0x10, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0c, 0x48, 0x00, 0x52, 0x02, 0x69, 0x64, 0x42, 0x05, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x22,
	0xed, 0x03, 0x0a, 0x0a, 0x4c, 0x6f, 0x61, 0x6e, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x12, 0x15,
	0x0a, 0x03, 0x61, 0x6c, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x03, 0x61,
	0x6c, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x22, 0x0a, 0x0a, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x48, 0x01, 0x52, 0x09, 0x73, 0x75, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x3a, 0x0a, 0x07, 0x73, 0x75, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2e, 0x42, 0x6f, 0x6f,
	0x6b, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x48, 0x02, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x88, 0x01, 0x01, 0x12, 0x24, 0x0a, 0x0b, 0x62, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x48, 0x03, 0x52, 0x0a, 0x62, 0x6f,
	0x72, 0x72, 0x6f, 0x77, 0x65, 0x72, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x3e, 0x0a, 0x08, 0x62,
	0x6f, 0x72, 0x72, 0x6f, 0x77, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2e,
	0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x48, 0x04, 0x52, 0x08,
	0x62, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x48, 0x05, 0x52, 0x06, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x88, 0x01, 0x01, 0x12, 0x1e, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x65,
	0x5f, 0x64, 0x75, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x48, 0x06, 0x52, 0x07, 0x64, 0x61,
	0x74, 0x65, 0x44, 0x75, 0x65, 0x88, 0x01, 0x01, 0x12, 0x24, 0x0a, 0x0b, 0x64, 0x61, 0x74, 0x65,
	0x5f, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x08, 0x48, 0x07, 0x52,
	0x0a, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x26,
	0x0a, 0x0c, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x0f,
	0x20, 0x01, 0x28, 0x08, 0x48, 0x08, 0x52, 0x0b, 0x64, 0x61, 0x74, 0x65, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x88, 0x01, 0x01, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x61, 0x6c, 0x6c, 0x42, 0x0d,
	0x0a, 0x0b, 0x5f, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x42, 0x0a, 0x0a,
	0x08, 0x5f, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x62, 0x6f,
	0x72, 0x72, 0x6f, 0x77, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x62, 0x6f,
	0x72, 0x72, 0x6f, 0x77, 0x65, 0x72, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x64, 0x75, 0x65, 0x42, 0x0e,
	0x0a, 0x0c, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x42, 0x0f,
	0x0a, 0x0d, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x32,
	0x8b, 0x01, 0x0a, 0x0b, 0x4c, 0x6f, 0x61, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x3d, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12, 0x1f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2e, 0x4c, 0x6f, 0x61, 0x6e, 0x41, 0x64, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2e, 0x4c, 0x6f, 0x61, 0x6e, 0x12, 0x3d,
	0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x1f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e,
	0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2e, 0x4c, 0x6f, 0x61, 0x6e, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2e, 0x4c, 0x6f, 0x61, 0x6e, 0x42, 0x38, 0x5a,
	0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x65, 0x73, 0x6f,
	0x6d, 0x6e, 0x75, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x6f, 0x72, 0x6d, 0x2f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f,
	0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_example_library_loan_svc_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_example_library_loan_svc_proto_goTypes = []any{
	(*LoanAddRequest)(nil),        // 0: example.library.LoanAddRequest
	(*LoanGetRequest)(nil),        // 1: example.library.LoanGetRequest
	(*LoanSelect)(nil),            // 2: example.library.LoanSelect
	(*BookGetRequest)(nil),        // 3: example.library.BookGetRequest
	(*MemberGetRequest)(nil),      // 4: example.library.MemberGetRequest
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
	(*BookSelect)(nil),            // 6: example.library.BookSelect
	(*MemberSelect)(nil),          // 7: example.library.MemberSelect
	(*Loan)(nil),                  // 8: example.library.Loan
}
var file_example_library_loan_svc_proto_depIdxs = []int32{
	3,  // 0: example.library.LoanAddRequest.subject:type_name -> example.library.BookGetRequest
	4,  // 1: example.library.LoanAddRequest.borrower:type_name -> example.library.MemberGetRequest
	5,  // 2: example.library.LoanAddRequest.date_due:type_name -> google.protobuf.Timestamp
	5,  // 3: example.library.LoanAddRequest.date_return:type_name -> google.protobuf.Timestamp
	5,  // 4: example.library.LoanAddRequest.date_created:type_name -> google.protobuf.Timestamp
	2,  // 5: example.library.LoanGetRequest.select:type_name -> example.library.LoanSelect
	6,  // 6: example.library.LoanSelect.subject:type_name -> example.library.BookSelect
	7,  // 7: example.library.LoanSelect.borrower:type_name -> example.library.MemberSelect
	0,  // 8: example.library.LoanService.Add:input_type -> example.library.LoanAddRequest
	1,  // 9: example.library.LoanService.Get:input_type -> example.library.LoanGetRequest
	8,  // 10: example.library.LoanService.Add:output_type -> example.library.Loan
	8,  // 11: example.library.LoanService.Get:output_type -> example.library.Loan
	10, // [10:12] is the sub-list for method output_type
	8,  // [8:10] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_example_library_loan_svc_proto_init() }
func file_example_library_loan_svc_proto_init() {
	if File_example_library_loan_svc_proto != nil {
		return
	}
	file_example_library_book_svc_proto_init()
	file_example_library_loan_proto_init()
	file_example_library_member_svc_proto_init()
	file_example_library_loan_svc_proto_msgTypes[0].OneofWrappers = []any{}
	file_example_library_loan_svc_proto_msgTypes[1].OneofWrappers = []any{
		(*loanGetRequest_Id)(nil),
	}
	file_example_library_loan_svc_proto_msgTypes[2].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_example_library_loan_svc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_example_library_loan_svc_proto_goTypes,
		DependencyIndexes: file_example_library_loan_svc_proto_depIdxs,
		MessageInfos:      file_example_library_loan_svc_proto_msgTypes,
	}.Build()
	File_example_library_loan_svc_proto = out.File
	file_example_library_loan_svc_proto_rawDesc = nil
	file_example_library_loan_svc_proto_goTypes = nil
	file_example_library_loan_svc_proto_depIdxs = nil
}
