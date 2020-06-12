// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0-devel
// 	protoc        v3.12.0
// source: carrier.certificate.v1.proto

package protos

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type ClaimOpt int32

const (
	ClaimOpt_INVALID    ClaimOpt = 0
	ClaimOpt_DELEGATION ClaimOpt = 1
)

// Enum value maps for ClaimOpt.
var (
	ClaimOpt_name = map[int32]string{
		0: "INVALID",
		1: "DELEGATION",
	}
	ClaimOpt_value = map[string]int32{
		"INVALID":    0,
		"DELEGATION": 1,
	}
)

func (x ClaimOpt) Enum() *ClaimOpt {
	p := new(ClaimOpt)
	*p = x
	return p
}

func (x ClaimOpt) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ClaimOpt) Descriptor() protoreflect.EnumDescriptor {
	return file_carrier_certificate_v1_proto_enumTypes[0].Descriptor()
}

func (ClaimOpt) Type() protoreflect.EnumType {
	return &file_carrier_certificate_v1_proto_enumTypes[0]
}

func (x ClaimOpt) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ClaimOpt.Descriptor instead.
func (ClaimOpt) EnumDescriptor() ([]byte, []int) {
	return file_carrier_certificate_v1_proto_rawDescGZIP(), []int{0}
}

type ClaimOne struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Target    []byte   `protobuf:"bytes,1,opt,name=target,proto3" json:"target,omitempty"`
	Resources []string `protobuf:"bytes,2,rep,name=resources,proto3" json:"resources,omitempty"`
}

func (x *ClaimOne) Reset() {
	*x = ClaimOne{}
	if protoimpl.UnsafeEnabled {
		mi := &file_carrier_certificate_v1_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClaimOne) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClaimOne) ProtoMessage() {}

func (x *ClaimOne) ProtoReflect() protoreflect.Message {
	mi := &file_carrier_certificate_v1_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClaimOne.ProtoReflect.Descriptor instead.
func (*ClaimOne) Descriptor() ([]byte, []int) {
	return file_carrier_certificate_v1_proto_rawDescGZIP(), []int{0}
}

func (x *ClaimOne) GetTarget() []byte {
	if x != nil {
		return x.Target
	}
	return nil
}

func (x *ClaimOne) GetResources() []string {
	if x != nil {
		return x.Resources
	}
	return nil
}

type ClaimAll struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Shadow    []byte   `protobuf:"bytes,1,opt,name=shadow,proto3" json:"shadow,omitempty"`
	Resources []string `protobuf:"bytes,2,rep,name=resources,proto3" json:"resources,omitempty"`
}

func (x *ClaimAll) Reset() {
	*x = ClaimAll{}
	if protoimpl.UnsafeEnabled {
		mi := &file_carrier_certificate_v1_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClaimAll) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClaimAll) ProtoMessage() {}

func (x *ClaimAll) ProtoReflect() protoreflect.Message {
	mi := &file_carrier_certificate_v1_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClaimAll.ProtoReflect.Descriptor instead.
func (*ClaimAll) Descriptor() ([]byte, []int) {
	return file_carrier_certificate_v1_proto_rawDescGZIP(), []int{1}
}

func (x *ClaimAll) GetShadow() []byte {
	if x != nil {
		return x.Shadow
	}
	return nil
}

func (x *ClaimAll) GetResources() []string {
	if x != nil {
		return x.Resources
	}
	return nil
}

type Revoker struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identity []byte `protobuf:"bytes,1,opt,name=identity,proto3" json:"identity,omitempty"`
}

func (x *Revoker) Reset() {
	*x = Revoker{}
	if protoimpl.UnsafeEnabled {
		mi := &file_carrier_certificate_v1_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Revoker) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Revoker) ProtoMessage() {}

func (x *Revoker) ProtoReflect() protoreflect.Message {
	mi := &file_carrier_certificate_v1_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Revoker.ProtoReflect.Descriptor instead.
func (*Revoker) Descriptor() ([]byte, []int) {
	return file_carrier_certificate_v1_proto_rawDescGZIP(), []int{2}
}

func (x *Revoker) GetIdentity() []byte {
	if x != nil {
		return x.Identity
	}
	return nil
}

type Claim struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Claim:
	//	*Claim_Opt
	//	*Claim_One
	//	*Claim_All
	//	*Claim_Revoker
	Claim isClaim_Claim `protobuf_oneof:"claim"`
}

func (x *Claim) Reset() {
	*x = Claim{}
	if protoimpl.UnsafeEnabled {
		mi := &file_carrier_certificate_v1_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Claim) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Claim) ProtoMessage() {}

func (x *Claim) ProtoReflect() protoreflect.Message {
	mi := &file_carrier_certificate_v1_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Claim.ProtoReflect.Descriptor instead.
func (*Claim) Descriptor() ([]byte, []int) {
	return file_carrier_certificate_v1_proto_rawDescGZIP(), []int{3}
}

func (m *Claim) GetClaim() isClaim_Claim {
	if m != nil {
		return m.Claim
	}
	return nil
}

func (x *Claim) GetOpt() ClaimOpt {
	if x, ok := x.GetClaim().(*Claim_Opt); ok {
		return x.Opt
	}
	return ClaimOpt_INVALID
}

func (x *Claim) GetOne() *ClaimOne {
	if x, ok := x.GetClaim().(*Claim_One); ok {
		return x.One
	}
	return nil
}

func (x *Claim) GetAll() *ClaimAll {
	if x, ok := x.GetClaim().(*Claim_All); ok {
		return x.All
	}
	return nil
}

func (x *Claim) GetRevoker() *Revoker {
	if x, ok := x.GetClaim().(*Claim_Revoker); ok {
		return x.Revoker
	}
	return nil
}

type isClaim_Claim interface {
	isClaim_Claim()
}

type Claim_Opt struct {
	Opt ClaimOpt `protobuf:"varint,1,opt,name=opt,proto3,enum=carrier.certificate.v1.ClaimOpt,oneof"`
}

type Claim_One struct {
	One *ClaimOne `protobuf:"bytes,2,opt,name=one,proto3,oneof"`
}

type Claim_All struct {
	All *ClaimAll `protobuf:"bytes,3,opt,name=all,proto3,oneof"`
}

type Claim_Revoker struct {
	Revoker *Revoker `protobuf:"bytes,4,opt,name=revoker,proto3,oneof"`
}

func (*Claim_Opt) isClaim_Claim() {}

func (*Claim_One) isClaim_Claim() {}

func (*Claim_All) isClaim_Claim() {}

func (*Claim_Revoker) isClaim_Claim() {}

type Certificate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LastValidEpoch uint32   `protobuf:"varint,1,opt,name=last_valid_epoch,json=lastValidEpoch,proto3" json:"last_valid_epoch,omitempty"`
	Identity       []byte   `protobuf:"bytes,2,opt,name=identity,proto3" json:"identity,omitempty"`
	Authority      []byte   `protobuf:"bytes,3,opt,name=authority,proto3" json:"authority,omitempty"`
	Serial         uint64   `protobuf:"varint,4,opt,name=serial,proto3" json:"serial,omitempty"`
	Claims         []*Claim `protobuf:"bytes,5,rep,name=claims,proto3" json:"claims,omitempty"`
}

func (x *Certificate) Reset() {
	*x = Certificate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_carrier_certificate_v1_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Certificate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Certificate) ProtoMessage() {}

func (x *Certificate) ProtoReflect() protoreflect.Message {
	mi := &file_carrier_certificate_v1_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Certificate.ProtoReflect.Descriptor instead.
func (*Certificate) Descriptor() ([]byte, []int) {
	return file_carrier_certificate_v1_proto_rawDescGZIP(), []int{4}
}

func (x *Certificate) GetLastValidEpoch() uint32 {
	if x != nil {
		return x.LastValidEpoch
	}
	return 0
}

func (x *Certificate) GetIdentity() []byte {
	if x != nil {
		return x.Identity
	}
	return nil
}

func (x *Certificate) GetAuthority() []byte {
	if x != nil {
		return x.Authority
	}
	return nil
}

func (x *Certificate) GetSerial() uint64 {
	if x != nil {
		return x.Serial
	}
	return 0
}

func (x *Certificate) GetClaims() []*Claim {
	if x != nil {
		return x.Claims
	}
	return nil
}

type CertificateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LastValidEpoch uint32   `protobuf:"varint,1,opt,name=last_valid_epoch,json=lastValidEpoch,proto3" json:"last_valid_epoch,omitempty"`
	Identity       []byte   `protobuf:"bytes,2,opt,name=identity,proto3" json:"identity,omitempty"`
	Claims         []*Claim `protobuf:"bytes,3,rep,name=claims,proto3" json:"claims,omitempty"`
}

func (x *CertificateRequest) Reset() {
	*x = CertificateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_carrier_certificate_v1_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CertificateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CertificateRequest) ProtoMessage() {}

func (x *CertificateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_carrier_certificate_v1_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CertificateRequest.ProtoReflect.Descriptor instead.
func (*CertificateRequest) Descriptor() ([]byte, []int) {
	return file_carrier_certificate_v1_proto_rawDescGZIP(), []int{5}
}

func (x *CertificateRequest) GetLastValidEpoch() uint32 {
	if x != nil {
		return x.LastValidEpoch
	}
	return 0
}

func (x *CertificateRequest) GetIdentity() []byte {
	if x != nil {
		return x.Identity
	}
	return nil
}

func (x *CertificateRequest) GetClaims() []*Claim {
	if x != nil {
		return x.Claims
	}
	return nil
}

type Authorization struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identity string `protobuf:"bytes,1,opt,name=identity,proto3" json:"identity,omitempty"`
	Resource string `protobuf:"bytes,2,opt,name=resource,proto3" json:"resource,omitempty"`
}

func (x *Authorization) Reset() {
	*x = Authorization{}
	if protoimpl.UnsafeEnabled {
		mi := &file_carrier_certificate_v1_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Authorization) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Authorization) ProtoMessage() {}

func (x *Authorization) ProtoReflect() protoreflect.Message {
	mi := &file_carrier_certificate_v1_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Authorization.ProtoReflect.Descriptor instead.
func (*Authorization) Descriptor() ([]byte, []int) {
	return file_carrier_certificate_v1_proto_rawDescGZIP(), []int{6}
}

func (x *Authorization) GetIdentity() string {
	if x != nil {
		return x.Identity
	}
	return ""
}

func (x *Authorization) GetResource() string {
	if x != nil {
		return x.Resource
	}
	return ""
}

type AuthorizationList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	A []*Authorization `protobuf:"bytes,1,rep,name=a,proto3" json:"a,omitempty"`
}

func (x *AuthorizationList) Reset() {
	*x = AuthorizationList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_carrier_certificate_v1_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthorizationList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorizationList) ProtoMessage() {}

func (x *AuthorizationList) ProtoReflect() protoreflect.Message {
	mi := &file_carrier_certificate_v1_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorizationList.ProtoReflect.Descriptor instead.
func (*AuthorizationList) Descriptor() ([]byte, []int) {
	return file_carrier_certificate_v1_proto_rawDescGZIP(), []int{7}
}

func (x *AuthorizationList) GetA() []*Authorization {
	if x != nil {
		return x.A
	}
	return nil
}

var File_carrier_certificate_v1_proto protoreflect.FileDescriptor

var file_carrier_certificate_v1_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x63, 0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0x2e, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16,
	0x63, 0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0x2e, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x22, 0x40, 0x0a, 0x08, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x4f,
	0x6e, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x22, 0x40, 0x0a, 0x08, 0x43, 0x6c, 0x61, 0x69,
	0x6d, 0x41, 0x6c, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x68, 0x61, 0x64, 0x6f, 0x77, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x73, 0x68, 0x61, 0x64, 0x6f, 0x77, 0x12, 0x1c, 0x0a, 0x09,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x22, 0x25, 0x0a, 0x07, 0x52, 0x65,
	0x76, 0x6f, 0x6b, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x22, 0xef, 0x01, 0x0a, 0x05, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x12, 0x34, 0x0a, 0x03, 0x6f,
	0x70, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x63, 0x61, 0x72, 0x72, 0x69,
	0x65, 0x72, 0x2e, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x4f, 0x70, 0x74, 0x48, 0x00, 0x52, 0x03, 0x6f, 0x70,
	0x74, 0x12, 0x34, 0x0a, 0x03, 0x6f, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20,
	0x2e, 0x63, 0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0x2e, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x4f, 0x6e, 0x65,
	0x48, 0x00, 0x52, 0x03, 0x6f, 0x6e, 0x65, 0x12, 0x34, 0x0a, 0x03, 0x61, 0x6c, 0x6c, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x63, 0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0x2e, 0x63,
	0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c,
	0x61, 0x69, 0x6d, 0x41, 0x6c, 0x6c, 0x48, 0x00, 0x52, 0x03, 0x61, 0x6c, 0x6c, 0x12, 0x3b, 0x0a,
	0x07, 0x72, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f,
	0x2e, 0x63, 0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0x2e, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x72, 0x48,
	0x00, 0x52, 0x07, 0x72, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x72, 0x42, 0x07, 0x0a, 0x05, 0x63, 0x6c,
	0x61, 0x69, 0x6d, 0x22, 0xc0, 0x01, 0x0a, 0x0b, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x5f, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x6c,
	0x61, 0x73, 0x74, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x45, 0x70, 0x6f, 0x63, 0x68, 0x12, 0x1a, 0x0a,
	0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x72, 0x69, 0x61,
	0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x12,
	0x35, 0x0a, 0x06, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1d, 0x2e, 0x63, 0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0x2e, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x52, 0x06,
	0x63, 0x6c, 0x61, 0x69, 0x6d, 0x73, 0x22, 0x91, 0x01, 0x0a, 0x12, 0x43, 0x65, 0x72, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a,
	0x10, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x65, 0x70, 0x6f, 0x63,
	0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x6c, 0x61, 0x73, 0x74, 0x56, 0x61, 0x6c,
	0x69, 0x64, 0x45, 0x70, 0x6f, 0x63, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x12, 0x35, 0x0a, 0x06, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x63, 0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0x2e, 0x63, 0x65,
	0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x61,
	0x69, 0x6d, 0x52, 0x06, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x73, 0x22, 0x47, 0x0a, 0x0d, 0x41, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x22, 0x48, 0x0a, 0x11, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x33, 0x0a, 0x01, 0x61, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x63, 0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0x2e, 0x63, 0x65,
	0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x01, 0x61, 0x2a, 0x27, 0x0a,
	0x08, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x4f, 0x70, 0x74, 0x12, 0x0b, 0x0a, 0x07, 0x49, 0x4e, 0x56,
	0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x44, 0x45, 0x4c, 0x45, 0x47, 0x41,
	0x54, 0x49, 0x4f, 0x4e, 0x10, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_carrier_certificate_v1_proto_rawDescOnce sync.Once
	file_carrier_certificate_v1_proto_rawDescData = file_carrier_certificate_v1_proto_rawDesc
)

func file_carrier_certificate_v1_proto_rawDescGZIP() []byte {
	file_carrier_certificate_v1_proto_rawDescOnce.Do(func() {
		file_carrier_certificate_v1_proto_rawDescData = protoimpl.X.CompressGZIP(file_carrier_certificate_v1_proto_rawDescData)
	})
	return file_carrier_certificate_v1_proto_rawDescData
}

var file_carrier_certificate_v1_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_carrier_certificate_v1_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_carrier_certificate_v1_proto_goTypes = []interface{}{
	(ClaimOpt)(0),              // 0: carrier.certificate.v1.ClaimOpt
	(*ClaimOne)(nil),           // 1: carrier.certificate.v1.ClaimOne
	(*ClaimAll)(nil),           // 2: carrier.certificate.v1.ClaimAll
	(*Revoker)(nil),            // 3: carrier.certificate.v1.Revoker
	(*Claim)(nil),              // 4: carrier.certificate.v1.Claim
	(*Certificate)(nil),        // 5: carrier.certificate.v1.Certificate
	(*CertificateRequest)(nil), // 6: carrier.certificate.v1.CertificateRequest
	(*Authorization)(nil),      // 7: carrier.certificate.v1.Authorization
	(*AuthorizationList)(nil),  // 8: carrier.certificate.v1.AuthorizationList
}
var file_carrier_certificate_v1_proto_depIdxs = []int32{
	0, // 0: carrier.certificate.v1.Claim.opt:type_name -> carrier.certificate.v1.ClaimOpt
	1, // 1: carrier.certificate.v1.Claim.one:type_name -> carrier.certificate.v1.ClaimOne
	2, // 2: carrier.certificate.v1.Claim.all:type_name -> carrier.certificate.v1.ClaimAll
	3, // 3: carrier.certificate.v1.Claim.revoker:type_name -> carrier.certificate.v1.Revoker
	4, // 4: carrier.certificate.v1.Certificate.claims:type_name -> carrier.certificate.v1.Claim
	4, // 5: carrier.certificate.v1.CertificateRequest.claims:type_name -> carrier.certificate.v1.Claim
	7, // 6: carrier.certificate.v1.AuthorizationList.a:type_name -> carrier.certificate.v1.Authorization
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_carrier_certificate_v1_proto_init() }
func file_carrier_certificate_v1_proto_init() {
	if File_carrier_certificate_v1_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_carrier_certificate_v1_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClaimOne); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_carrier_certificate_v1_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClaimAll); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_carrier_certificate_v1_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Revoker); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_carrier_certificate_v1_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Claim); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_carrier_certificate_v1_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Certificate); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_carrier_certificate_v1_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CertificateRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_carrier_certificate_v1_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Authorization); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_carrier_certificate_v1_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthorizationList); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_carrier_certificate_v1_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*Claim_Opt)(nil),
		(*Claim_One)(nil),
		(*Claim_All)(nil),
		(*Claim_Revoker)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_carrier_certificate_v1_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_carrier_certificate_v1_proto_goTypes,
		DependencyIndexes: file_carrier_certificate_v1_proto_depIdxs,
		EnumInfos:         file_carrier_certificate_v1_proto_enumTypes,
		MessageInfos:      file_carrier_certificate_v1_proto_msgTypes,
	}.Build()
	File_carrier_certificate_v1_proto = out.File
	file_carrier_certificate_v1_proto_rawDesc = nil
	file_carrier_certificate_v1_proto_goTypes = nil
	file_carrier_certificate_v1_proto_depIdxs = nil
}