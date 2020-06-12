// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0-devel
// 	protoc        v3.12.0
// source: carrier.config.v1.proto

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

type AuthorizationGet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AuthorizationGet) Reset() {
	*x = AuthorizationGet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_carrier_config_v1_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthorizationGet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorizationGet) ProtoMessage() {}

func (x *AuthorizationGet) ProtoReflect() protoreflect.Message {
	mi := &file_carrier_config_v1_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorizationGet.ProtoReflect.Descriptor instead.
func (*AuthorizationGet) Descriptor() ([]byte, []int) {
	return file_carrier_config_v1_proto_rawDescGZIP(), []int{0}
}

type AuthorizationAdd struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identity string `protobuf:"bytes,1,opt,name=identity,proto3" json:"identity,omitempty"`
	Path     string `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *AuthorizationAdd) Reset() {
	*x = AuthorizationAdd{}
	if protoimpl.UnsafeEnabled {
		mi := &file_carrier_config_v1_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthorizationAdd) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorizationAdd) ProtoMessage() {}

func (x *AuthorizationAdd) ProtoReflect() protoreflect.Message {
	mi := &file_carrier_config_v1_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorizationAdd.ProtoReflect.Descriptor instead.
func (*AuthorizationAdd) Descriptor() ([]byte, []int) {
	return file_carrier_config_v1_proto_rawDescGZIP(), []int{1}
}

func (x *AuthorizationAdd) GetIdentity() string {
	if x != nil {
		return x.Identity
	}
	return ""
}

func (x *AuthorizationAdd) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type AuthorizationDel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identity string `protobuf:"bytes,1,opt,name=identity,proto3" json:"identity,omitempty"`
}

func (x *AuthorizationDel) Reset() {
	*x = AuthorizationDel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_carrier_config_v1_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthorizationDel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorizationDel) ProtoMessage() {}

func (x *AuthorizationDel) ProtoReflect() protoreflect.Message {
	mi := &file_carrier_config_v1_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorizationDel.ProtoReflect.Descriptor instead.
func (*AuthorizationDel) Descriptor() ([]byte, []int) {
	return file_carrier_config_v1_proto_rawDescGZIP(), []int{2}
}

func (x *AuthorizationDel) GetIdentity() string {
	if x != nil {
		return x.Identity
	}
	return ""
}

type InteractiveAuth struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Otp string `protobuf:"bytes,1,opt,name=otp,proto3" json:"otp,omitempty"`
}

func (x *InteractiveAuth) Reset() {
	*x = InteractiveAuth{}
	if protoimpl.UnsafeEnabled {
		mi := &file_carrier_config_v1_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InteractiveAuth) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InteractiveAuth) ProtoMessage() {}

func (x *InteractiveAuth) ProtoReflect() protoreflect.Message {
	mi := &file_carrier_config_v1_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InteractiveAuth.ProtoReflect.Descriptor instead.
func (*InteractiveAuth) Descriptor() ([]byte, []int) {
	return file_carrier_config_v1_proto_rawDescGZIP(), []int{3}
}

func (x *InteractiveAuth) GetOtp() string {
	if x != nil {
		return x.Otp
	}
	return ""
}

type NetworkJoin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Secret string `protobuf:"bytes,1,opt,name=secret,proto3" json:"secret,omitempty"`
}

func (x *NetworkJoin) Reset() {
	*x = NetworkJoin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_carrier_config_v1_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetworkJoin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkJoin) ProtoMessage() {}

func (x *NetworkJoin) ProtoReflect() protoreflect.Message {
	mi := &file_carrier_config_v1_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkJoin.ProtoReflect.Descriptor instead.
func (*NetworkJoin) Descriptor() ([]byte, []int) {
	return file_carrier_config_v1_proto_rawDescGZIP(), []int{4}
}

func (x *NetworkJoin) GetSecret() string {
	if x != nil {
		return x.Secret
	}
	return ""
}

type NetworkGet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NetworkGet) Reset() {
	*x = NetworkGet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_carrier_config_v1_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetworkGet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkGet) ProtoMessage() {}

func (x *NetworkGet) ProtoReflect() protoreflect.Message {
	mi := &file_carrier_config_v1_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkGet.ProtoReflect.Descriptor instead.
func (*NetworkGet) Descriptor() ([]byte, []int) {
	return file_carrier_config_v1_proto_rawDescGZIP(), []int{5}
}

type ConfigResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok    bool   `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	Error string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *ConfigResult) Reset() {
	*x = ConfigResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_carrier_config_v1_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigResult) ProtoMessage() {}

func (x *ConfigResult) ProtoReflect() protoreflect.Message {
	mi := &file_carrier_config_v1_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigResult.ProtoReflect.Descriptor instead.
func (*ConfigResult) Descriptor() ([]byte, []int) {
	return file_carrier_config_v1_proto_rawDescGZIP(), []int{6}
}

func (x *ConfigResult) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

func (x *ConfigResult) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type AuthListResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Auth []*AuthorizationAdd `protobuf:"bytes,1,rep,name=auth,proto3" json:"auth,omitempty"`
}

func (x *AuthListResult) Reset() {
	*x = AuthListResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_carrier_config_v1_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthListResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthListResult) ProtoMessage() {}

func (x *AuthListResult) ProtoReflect() protoreflect.Message {
	mi := &file_carrier_config_v1_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthListResult.ProtoReflect.Descriptor instead.
func (*AuthListResult) Descriptor() ([]byte, []int) {
	return file_carrier_config_v1_proto_rawDescGZIP(), []int{7}
}

func (x *AuthListResult) GetAuth() []*AuthorizationAdd {
	if x != nil {
		return x.Auth
	}
	return nil
}

type NetworkGetResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *NetworkGetResult) Reset() {
	*x = NetworkGetResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_carrier_config_v1_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetworkGetResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkGetResult) ProtoMessage() {}

func (x *NetworkGetResult) ProtoReflect() protoreflect.Message {
	mi := &file_carrier_config_v1_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkGetResult.ProtoReflect.Descriptor instead.
func (*NetworkGetResult) Descriptor() ([]byte, []int) {
	return file_carrier_config_v1_proto_rawDescGZIP(), []int{8}
}

func (x *NetworkGetResult) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

var File_carrier_config_v1_proto protoreflect.FileDescriptor

var file_carrier_config_v1_proto_rawDesc = []byte{
	0x0a, 0x17, 0x63, 0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x63, 0x61, 0x72, 0x72, 0x69,
	0x65, 0x72, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x22, 0x12, 0x0a, 0x10,
	0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x65, 0x74,
	0x22, 0x42, 0x0a, 0x10, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x41, 0x64, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x70, 0x61, 0x74, 0x68, 0x22, 0x2e, 0x0a, 0x10, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x22, 0x23, 0x0a, 0x0f, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74,
	0x69, 0x76, 0x65, 0x41, 0x75, 0x74, 0x68, 0x12, 0x10, 0x0a, 0x03, 0x6f, 0x74, 0x70, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6f, 0x74, 0x70, 0x22, 0x25, 0x0a, 0x0b, 0x4e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x4a, 0x6f, 0x69, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x22, 0x0c, 0x0a, 0x0a, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x47, 0x65, 0x74, 0x22, 0x34,
	0x0a, 0x0c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x22, 0x49, 0x0a, 0x0e, 0x41, 0x75, 0x74, 0x68, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x37, 0x0a, 0x04, 0x61, 0x75, 0x74, 0x68, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x63, 0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x64, 0x64, 0x52, 0x04, 0x61, 0x75, 0x74, 0x68, 0x22,
	0x2c, 0x0a, 0x10, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x32, 0x9d, 0x03,
	0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x50, 0x0a, 0x08, 0x61, 0x75, 0x74, 0x68,
	0x5f, 0x61, 0x64, 0x64, 0x12, 0x23, 0x2e, 0x63, 0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x64, 0x64, 0x1a, 0x1f, 0x2e, 0x63, 0x61, 0x72, 0x72,
	0x69, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x53, 0x0a, 0x09, 0x61, 0x75,
	0x74, 0x68, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x23, 0x2e, 0x63, 0x61, 0x72, 0x72, 0x69, 0x65,
	0x72, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x65, 0x74, 0x1a, 0x21, 0x2e, 0x63,
	0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31,
	0x2e, 0x41, 0x75, 0x74, 0x68, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12,
	0x50, 0x0a, 0x08, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x64, 0x65, 0x6c, 0x12, 0x23, 0x2e, 0x63, 0x61,
	0x72, 0x72, 0x69, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e,
	0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x6c,
	0x1a, 0x1f, 0x2e, 0x63, 0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x12, 0x4d, 0x0a, 0x07, 0x6e, 0x65, 0x74, 0x5f, 0x67, 0x65, 0x74, 0x12, 0x1d, 0x2e, 0x63,
	0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31,
	0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x47, 0x65, 0x74, 0x1a, 0x23, 0x2e, 0x63, 0x61,
	0x72, 0x72, 0x69, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e,
	0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x12, 0x4b, 0x0a, 0x08, 0x6e, 0x65, 0x74, 0x5f, 0x6a, 0x6f, 0x69, 0x6e, 0x12, 0x1e, 0x2e, 0x63,
	0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31,
	0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x4a, 0x6f, 0x69, 0x6e, 0x1a, 0x1f, 0x2e, 0x63,
	0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x42, 0x0a, 0x5a,
	0x08, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_carrier_config_v1_proto_rawDescOnce sync.Once
	file_carrier_config_v1_proto_rawDescData = file_carrier_config_v1_proto_rawDesc
)

func file_carrier_config_v1_proto_rawDescGZIP() []byte {
	file_carrier_config_v1_proto_rawDescOnce.Do(func() {
		file_carrier_config_v1_proto_rawDescData = protoimpl.X.CompressGZIP(file_carrier_config_v1_proto_rawDescData)
	})
	return file_carrier_config_v1_proto_rawDescData
}

var file_carrier_config_v1_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_carrier_config_v1_proto_goTypes = []interface{}{
	(*AuthorizationGet)(nil), // 0: carrier.config.v1.AuthorizationGet
	(*AuthorizationAdd)(nil), // 1: carrier.config.v1.AuthorizationAdd
	(*AuthorizationDel)(nil), // 2: carrier.config.v1.AuthorizationDel
	(*InteractiveAuth)(nil),  // 3: carrier.config.v1.InteractiveAuth
	(*NetworkJoin)(nil),      // 4: carrier.config.v1.NetworkJoin
	(*NetworkGet)(nil),       // 5: carrier.config.v1.NetworkGet
	(*ConfigResult)(nil),     // 6: carrier.config.v1.ConfigResult
	(*AuthListResult)(nil),   // 7: carrier.config.v1.AuthListResult
	(*NetworkGetResult)(nil), // 8: carrier.config.v1.NetworkGetResult
}
var file_carrier_config_v1_proto_depIdxs = []int32{
	1, // 0: carrier.config.v1.AuthListResult.auth:type_name -> carrier.config.v1.AuthorizationAdd
	1, // 1: carrier.config.v1.Config.auth_add:input_type -> carrier.config.v1.AuthorizationAdd
	0, // 2: carrier.config.v1.Config.auth_list:input_type -> carrier.config.v1.AuthorizationGet
	2, // 3: carrier.config.v1.Config.auth_del:input_type -> carrier.config.v1.AuthorizationDel
	5, // 4: carrier.config.v1.Config.net_get:input_type -> carrier.config.v1.NetworkGet
	4, // 5: carrier.config.v1.Config.net_join:input_type -> carrier.config.v1.NetworkJoin
	6, // 6: carrier.config.v1.Config.auth_add:output_type -> carrier.config.v1.ConfigResult
	7, // 7: carrier.config.v1.Config.auth_list:output_type -> carrier.config.v1.AuthListResult
	6, // 8: carrier.config.v1.Config.auth_del:output_type -> carrier.config.v1.ConfigResult
	8, // 9: carrier.config.v1.Config.net_get:output_type -> carrier.config.v1.NetworkGetResult
	6, // 10: carrier.config.v1.Config.net_join:output_type -> carrier.config.v1.ConfigResult
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_carrier_config_v1_proto_init() }
func file_carrier_config_v1_proto_init() {
	if File_carrier_config_v1_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_carrier_config_v1_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthorizationGet); i {
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
		file_carrier_config_v1_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthorizationAdd); i {
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
		file_carrier_config_v1_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthorizationDel); i {
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
		file_carrier_config_v1_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InteractiveAuth); i {
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
		file_carrier_config_v1_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetworkJoin); i {
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
		file_carrier_config_v1_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetworkGet); i {
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
		file_carrier_config_v1_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigResult); i {
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
		file_carrier_config_v1_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthListResult); i {
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
		file_carrier_config_v1_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetworkGetResult); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_carrier_config_v1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_carrier_config_v1_proto_goTypes,
		DependencyIndexes: file_carrier_config_v1_proto_depIdxs,
		MessageInfos:      file_carrier_config_v1_proto_msgTypes,
	}.Build()
	File_carrier_config_v1_proto = out.File
	file_carrier_config_v1_proto_rawDesc = nil
	file_carrier_config_v1_proto_goTypes = nil
	file_carrier_config_v1_proto_depIdxs = nil
}
