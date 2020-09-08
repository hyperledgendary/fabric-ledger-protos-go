// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.1
// source: contract_messages.proto

package contract

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

type GetMetadataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetMetadataRequest) Reset() {
	*x = GetMetadataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_messages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMetadataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMetadataRequest) ProtoMessage() {}

func (x *GetMetadataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_contract_messages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMetadataRequest.ProtoReflect.Descriptor instead.
func (*GetMetadataRequest) Descriptor() ([]byte, []int) {
	return file_contract_messages_proto_rawDescGZIP(), []int{0}
}

type GetMetadataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payload []byte `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *GetMetadataResponse) Reset() {
	*x = GetMetadataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_messages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMetadataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMetadataResponse) ProtoMessage() {}

func (x *GetMetadataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_contract_messages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMetadataResponse.ProtoReflect.Descriptor instead.
func (*GetMetadataResponse) Descriptor() ([]byte, []int) {
	return file_contract_messages_proto_rawDescGZIP(), []int{1}
}

func (x *GetMetadataResponse) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

type InvokeTransactionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Context         *TransactionContext `protobuf:"bytes,1,opt,name=context,proto3" json:"context,omitempty"`
	TransactionName string              `protobuf:"bytes,2,opt,name=transaction_name,json=transactionName,proto3" json:"transaction_name,omitempty"`
	Args            [][]byte            `protobuf:"bytes,3,rep,name=args,proto3" json:"args,omitempty"`
	TransientArgs   map[string][]byte   `protobuf:"bytes,4,rep,name=transient_args,json=transientArgs,proto3" json:"transient_args,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *InvokeTransactionRequest) Reset() {
	*x = InvokeTransactionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_messages_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InvokeTransactionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InvokeTransactionRequest) ProtoMessage() {}

func (x *InvokeTransactionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_contract_messages_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InvokeTransactionRequest.ProtoReflect.Descriptor instead.
func (*InvokeTransactionRequest) Descriptor() ([]byte, []int) {
	return file_contract_messages_proto_rawDescGZIP(), []int{2}
}

func (x *InvokeTransactionRequest) GetContext() *TransactionContext {
	if x != nil {
		return x.Context
	}
	return nil
}

func (x *InvokeTransactionRequest) GetTransactionName() string {
	if x != nil {
		return x.TransactionName
	}
	return ""
}

func (x *InvokeTransactionRequest) GetArgs() [][]byte {
	if x != nil {
		return x.Args
	}
	return nil
}

func (x *InvokeTransactionRequest) GetTransientArgs() map[string][]byte {
	if x != nil {
		return x.TransientArgs
	}
	return nil
}

type InvokeTransactionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payload []byte `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *InvokeTransactionResponse) Reset() {
	*x = InvokeTransactionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_messages_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InvokeTransactionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InvokeTransactionResponse) ProtoMessage() {}

func (x *InvokeTransactionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_contract_messages_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InvokeTransactionResponse.ProtoReflect.Descriptor instead.
func (*InvokeTransactionResponse) Descriptor() ([]byte, []int) {
	return file_contract_messages_proto_rawDescGZIP(), []int{3}
}

func (x *InvokeTransactionResponse) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

type InvokeTransactionError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *InvokeTransactionError) Reset() {
	*x = InvokeTransactionError{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_messages_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InvokeTransactionError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InvokeTransactionError) ProtoMessage() {}

func (x *InvokeTransactionError) ProtoReflect() protoreflect.Message {
	mi := &file_contract_messages_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InvokeTransactionError.ProtoReflect.Descriptor instead.
func (*InvokeTransactionError) Descriptor() ([]byte, []int) {
	return file_contract_messages_proto_rawDescGZIP(), []int{4}
}

func (x *InvokeTransactionError) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *InvokeTransactionError) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// This is going to need TLS settings and client CA stuff...
type RegisterPeerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PeerAddress string `protobuf:"bytes,1,opt,name=peer_address,json=peerAddress,proto3" json:"peer_address,omitempty"`
	PeerPort    string `protobuf:"bytes,2,opt,name=peer_port,json=peerPort,proto3" json:"peer_port,omitempty"`
}

func (x *RegisterPeerRequest) Reset() {
	*x = RegisterPeerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_messages_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterPeerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterPeerRequest) ProtoMessage() {}

func (x *RegisterPeerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_contract_messages_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterPeerRequest.ProtoReflect.Descriptor instead.
func (*RegisterPeerRequest) Descriptor() ([]byte, []int) {
	return file_contract_messages_proto_rawDescGZIP(), []int{5}
}

func (x *RegisterPeerRequest) GetPeerAddress() string {
	if x != nil {
		return x.PeerAddress
	}
	return ""
}

func (x *RegisterPeerRequest) GetPeerPort() string {
	if x != nil {
		return x.PeerPort
	}
	return ""
}

type RegisterPeerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RegisterPeerResponse) Reset() {
	*x = RegisterPeerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_messages_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterPeerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterPeerResponse) ProtoMessage() {}

func (x *RegisterPeerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_contract_messages_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterPeerResponse.ProtoReflect.Descriptor instead.
func (*RegisterPeerResponse) Descriptor() ([]byte, []int) {
	return file_contract_messages_proto_rawDescGZIP(), []int{6}
}

var File_contract_messages_proto protoreflect.FileDescriptor

var file_contract_messages_proto_rawDesc = []byte{
	0x0a, 0x17, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x61, 0x63, 0x74, 0x1a, 0x15, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x14, 0x0a, 0x12, 0x47, 0x65,
	0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x2f, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x22, 0xb1, 0x02, 0x0a, 0x18, 0x49, 0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36,
	0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x52, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x72, 0x67, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0c, 0x52,
	0x04, 0x61, 0x72, 0x67, 0x73, 0x12, 0x5c, 0x0a, 0x0e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x65,
	0x6e, 0x74, 0x5f, 0x61, 0x72, 0x67, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x35, 0x2e,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x49, 0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x65, 0x6e, 0x74, 0x41, 0x72, 0x67, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x65, 0x6e, 0x74, 0x41,
	0x72, 0x67, 0x73, 0x1a, 0x40, 0x0a, 0x12, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x65, 0x6e, 0x74,
	0x41, 0x72, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x35, 0x0a, 0x19, 0x49, 0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x46, 0x0a, 0x16,
	0x49, 0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x22, 0x55, 0x0a, 0x13, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x50, 0x65, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x70,
	0x65, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x70, 0x65, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1b,
	0x0a, 0x09, 0x70, 0x65, 0x65, 0x72, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x65, 0x65, 0x72, 0x50, 0x6f, 0x72, 0x74, 0x22, 0x16, 0x0a, 0x14, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x50, 0x65, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x45, 0x5a, 0x43, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x79, 0x70, 0x65, 0x72, 0x6c,
	0x65, 0x64, 0x67, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x79, 0x2f, 0x66, 0x61, 0x62, 0x72, 0x69, 0x63,
	0x2d, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2d, 0x67,
	0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_contract_messages_proto_rawDescOnce sync.Once
	file_contract_messages_proto_rawDescData = file_contract_messages_proto_rawDesc
)

func file_contract_messages_proto_rawDescGZIP() []byte {
	file_contract_messages_proto_rawDescOnce.Do(func() {
		file_contract_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_contract_messages_proto_rawDescData)
	})
	return file_contract_messages_proto_rawDescData
}

var file_contract_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_contract_messages_proto_goTypes = []interface{}{
	(*GetMetadataRequest)(nil),        // 0: contract.GetMetadataRequest
	(*GetMetadataResponse)(nil),       // 1: contract.GetMetadataResponse
	(*InvokeTransactionRequest)(nil),  // 2: contract.InvokeTransactionRequest
	(*InvokeTransactionResponse)(nil), // 3: contract.InvokeTransactionResponse
	(*InvokeTransactionError)(nil),    // 4: contract.InvokeTransactionError
	(*RegisterPeerRequest)(nil),       // 5: contract.RegisterPeerRequest
	(*RegisterPeerResponse)(nil),      // 6: contract.RegisterPeerResponse
	nil,                               // 7: contract.InvokeTransactionRequest.TransientArgsEntry
	(*TransactionContext)(nil),        // 8: contract.TransactionContext
}
var file_contract_messages_proto_depIdxs = []int32{
	8, // 0: contract.InvokeTransactionRequest.context:type_name -> contract.TransactionContext
	7, // 1: contract.InvokeTransactionRequest.transient_args:type_name -> contract.InvokeTransactionRequest.TransientArgsEntry
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_contract_messages_proto_init() }
func file_contract_messages_proto_init() {
	if File_contract_messages_proto != nil {
		return
	}
	file_common_messages_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_contract_messages_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMetadataRequest); i {
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
		file_contract_messages_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMetadataResponse); i {
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
		file_contract_messages_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InvokeTransactionRequest); i {
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
		file_contract_messages_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InvokeTransactionResponse); i {
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
		file_contract_messages_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InvokeTransactionError); i {
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
		file_contract_messages_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterPeerRequest); i {
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
		file_contract_messages_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterPeerResponse); i {
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
			RawDescriptor: file_contract_messages_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_contract_messages_proto_goTypes,
		DependencyIndexes: file_contract_messages_proto_depIdxs,
		MessageInfos:      file_contract_messages_proto_msgTypes,
	}.Build()
	File_contract_messages_proto = out.File
	file_contract_messages_proto_rawDesc = nil
	file_contract_messages_proto_goTypes = nil
	file_contract_messages_proto_depIdxs = nil
}
