// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v4.25.3
// source: protobuf/wallet.proto

package wallet

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ReturnCode int32

const (
	ReturnCode_ERROR   ReturnCode = 0
	ReturnCode_SUCCESS ReturnCode = 1
)

// Enum value maps for ReturnCode.
var (
	ReturnCode_name = map[int32]string{
		0: "ERROR",
		1: "SUCCESS",
	}
	ReturnCode_value = map[string]int32{
		"ERROR":   0,
		"SUCCESS": 1,
	}
)

func (x ReturnCode) Enum() *ReturnCode {
	p := new(ReturnCode)
	*p = x
	return p
}

func (x ReturnCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ReturnCode) Descriptor() protoreflect.EnumDescriptor {
	return file_protobuf_wallet_proto_enumTypes[0].Descriptor()
}

func (ReturnCode) Type() protoreflect.EnumType {
	return &file_protobuf_wallet_proto_enumTypes[0]
}

func (x ReturnCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ReturnCode.Descriptor instead.
func (ReturnCode) EnumDescriptor() ([]byte, []int) {
	return file_protobuf_wallet_proto_rawDescGZIP(), []int{0}
}

type PublicKey struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	CompressPubkey string                 `protobuf:"bytes,1,opt,name=compress_pubkey,json=compressPubkey,proto3" json:"compress_pubkey,omitempty"`
	Pubkey         string                 `protobuf:"bytes,2,opt,name=pubkey,proto3" json:"pubkey,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *PublicKey) Reset() {
	*x = PublicKey{}
	mi := &file_protobuf_wallet_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PublicKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublicKey) ProtoMessage() {}

func (x *PublicKey) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_wallet_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublicKey.ProtoReflect.Descriptor instead.
func (*PublicKey) Descriptor() ([]byte, []int) {
	return file_protobuf_wallet_proto_rawDescGZIP(), []int{0}
}

func (x *PublicKey) GetCompressPubkey() string {
	if x != nil {
		return x.CompressPubkey
	}
	return ""
}

func (x *PublicKey) GetPubkey() string {
	if x != nil {
		return x.Pubkey
	}
	return ""
}

type SupportSignWayRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ConsumerToken string                 `protobuf:"bytes,1,opt,name=consumer_token,json=consumerToken,proto3" json:"consumer_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SupportSignWayRequest) Reset() {
	*x = SupportSignWayRequest{}
	mi := &file_protobuf_wallet_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SupportSignWayRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SupportSignWayRequest) ProtoMessage() {}

func (x *SupportSignWayRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_wallet_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SupportSignWayRequest.ProtoReflect.Descriptor instead.
func (*SupportSignWayRequest) Descriptor() ([]byte, []int) {
	return file_protobuf_wallet_proto_rawDescGZIP(), []int{1}
}

func (x *SupportSignWayRequest) GetConsumerToken() string {
	if x != nil {
		return x.ConsumerToken
	}
	return ""
}

type SignWay struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Schema        string                 `protobuf:"bytes,1,opt,name=schema,proto3" json:"schema,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SignWay) Reset() {
	*x = SignWay{}
	mi := &file_protobuf_wallet_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SignWay) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignWay) ProtoMessage() {}

func (x *SignWay) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_wallet_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignWay.ProtoReflect.Descriptor instead.
func (*SignWay) Descriptor() ([]byte, []int) {
	return file_protobuf_wallet_proto_rawDescGZIP(), []int{2}
}

func (x *SignWay) GetSchema() string {
	if x != nil {
		return x.Schema
	}
	return ""
}

type SupportSignWayResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Code          ReturnCode             `protobuf:"varint,1,opt,name=Code,proto3,enum=dapplink.wallet.ReturnCode" json:"Code,omitempty"`
	Msg           string                 `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	SignWay       []*SignWay             `protobuf:"bytes,3,rep,name=sign_way,json=signWay,proto3" json:"sign_way,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SupportSignWayResponse) Reset() {
	*x = SupportSignWayResponse{}
	mi := &file_protobuf_wallet_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SupportSignWayResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SupportSignWayResponse) ProtoMessage() {}

func (x *SupportSignWayResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_wallet_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SupportSignWayResponse.ProtoReflect.Descriptor instead.
func (*SupportSignWayResponse) Descriptor() ([]byte, []int) {
	return file_protobuf_wallet_proto_rawDescGZIP(), []int{3}
}

func (x *SupportSignWayResponse) GetCode() ReturnCode {
	if x != nil {
		return x.Code
	}
	return ReturnCode_ERROR
}

func (x *SupportSignWayResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *SupportSignWayResponse) GetSignWay() []*SignWay {
	if x != nil {
		return x.SignWay
	}
	return nil
}

type ExportPublicKeyRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ConsumerToken string                 `protobuf:"bytes,1,opt,name=consumer_token,json=consumerToken,proto3" json:"consumer_token,omitempty"`
	// CryptoType
	Type          string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Number        uint64 `protobuf:"varint,3,opt,name=number,proto3" json:"number,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ExportPublicKeyRequest) Reset() {
	*x = ExportPublicKeyRequest{}
	mi := &file_protobuf_wallet_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExportPublicKeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExportPublicKeyRequest) ProtoMessage() {}

func (x *ExportPublicKeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_wallet_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExportPublicKeyRequest.ProtoReflect.Descriptor instead.
func (*ExportPublicKeyRequest) Descriptor() ([]byte, []int) {
	return file_protobuf_wallet_proto_rawDescGZIP(), []int{4}
}

func (x *ExportPublicKeyRequest) GetConsumerToken() string {
	if x != nil {
		return x.ConsumerToken
	}
	return ""
}

func (x *ExportPublicKeyRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *ExportPublicKeyRequest) GetNumber() uint64 {
	if x != nil {
		return x.Number
	}
	return 0
}

type ExportPublicKeyResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Code          ReturnCode             `protobuf:"varint,1,opt,name=Code,proto3,enum=dapplink.wallet.ReturnCode" json:"Code,omitempty"`
	Msg           string                 `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	PublicKey     []*PublicKey           `protobuf:"bytes,3,rep,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ExportPublicKeyResponse) Reset() {
	*x = ExportPublicKeyResponse{}
	mi := &file_protobuf_wallet_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExportPublicKeyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExportPublicKeyResponse) ProtoMessage() {}

func (x *ExportPublicKeyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_wallet_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExportPublicKeyResponse.ProtoReflect.Descriptor instead.
func (*ExportPublicKeyResponse) Descriptor() ([]byte, []int) {
	return file_protobuf_wallet_proto_rawDescGZIP(), []int{5}
}

func (x *ExportPublicKeyResponse) GetCode() ReturnCode {
	if x != nil {
		return x.Code
	}
	return ReturnCode_ERROR
}

func (x *ExportPublicKeyResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *ExportPublicKeyResponse) GetPublicKey() []*PublicKey {
	if x != nil {
		return x.PublicKey
	}
	return nil
}

type SignTxMessageRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ConsumerToken string                 `protobuf:"bytes,1,opt,name=consumer_token,json=consumerToken,proto3" json:"consumer_token,omitempty"`
	// CryptoType
	Type          string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	PublicKey     string `protobuf:"bytes,3,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	MessageHash   string `protobuf:"bytes,4,opt,name=message_hash,json=messageHash,proto3" json:"message_hash,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SignTxMessageRequest) Reset() {
	*x = SignTxMessageRequest{}
	mi := &file_protobuf_wallet_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SignTxMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignTxMessageRequest) ProtoMessage() {}

func (x *SignTxMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_wallet_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignTxMessageRequest.ProtoReflect.Descriptor instead.
func (*SignTxMessageRequest) Descriptor() ([]byte, []int) {
	return file_protobuf_wallet_proto_rawDescGZIP(), []int{6}
}

func (x *SignTxMessageRequest) GetConsumerToken() string {
	if x != nil {
		return x.ConsumerToken
	}
	return ""
}

func (x *SignTxMessageRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *SignTxMessageRequest) GetPublicKey() string {
	if x != nil {
		return x.PublicKey
	}
	return ""
}

func (x *SignTxMessageRequest) GetMessageHash() string {
	if x != nil {
		return x.MessageHash
	}
	return ""
}

type SignTxMessageResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Code          ReturnCode             `protobuf:"varint,1,opt,name=Code,proto3,enum=dapplink.wallet.ReturnCode" json:"Code,omitempty"`
	Msg           string                 `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Signature     string                 `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SignTxMessageResponse) Reset() {
	*x = SignTxMessageResponse{}
	mi := &file_protobuf_wallet_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SignTxMessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignTxMessageResponse) ProtoMessage() {}

func (x *SignTxMessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_wallet_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignTxMessageResponse.ProtoReflect.Descriptor instead.
func (*SignTxMessageResponse) Descriptor() ([]byte, []int) {
	return file_protobuf_wallet_proto_rawDescGZIP(), []int{7}
}

func (x *SignTxMessageResponse) GetCode() ReturnCode {
	if x != nil {
		return x.Code
	}
	return ReturnCode_ERROR
}

func (x *SignTxMessageResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *SignTxMessageResponse) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

var File_protobuf_wallet_proto protoreflect.FileDescriptor

var file_protobuf_wallet_proto_rawDesc = string([]byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x61, 0x6c, 0x6c, 0x65,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x64, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x6e,
	0x6b, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x22, 0x4c, 0x0a, 0x09, 0x50, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73,
	0x73, 0x5f, 0x70, 0x75, 0x62, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x50, 0x75, 0x62, 0x6b, 0x65, 0x79, 0x12, 0x16,
	0x0a, 0x06, 0x70, 0x75, 0x62, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x70, 0x75, 0x62, 0x6b, 0x65, 0x79, 0x22, 0x3e, 0x0a, 0x15, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72,
	0x74, 0x53, 0x69, 0x67, 0x6e, 0x57, 0x61, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x5f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65,
	0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x21, 0x0a, 0x07, 0x53, 0x69, 0x67, 0x6e, 0x57, 0x61,
	0x79, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x22, 0x90, 0x01, 0x0a, 0x16, 0x53, 0x75,
	0x70, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x69, 0x67, 0x6e, 0x57, 0x61, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x64, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x77, 0x61,
	0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x52,
	0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x33, 0x0a, 0x08, 0x73, 0x69, 0x67, 0x6e, 0x5f,
	0x77, 0x61, 0x79, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x64, 0x61, 0x70, 0x70,
	0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x53, 0x69, 0x67, 0x6e,
	0x57, 0x61, 0x79, 0x52, 0x07, 0x73, 0x69, 0x67, 0x6e, 0x57, 0x61, 0x79, 0x22, 0x6b, 0x0a, 0x16,
	0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d,
	0x65, 0x72, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x97, 0x01, 0x0a, 0x17, 0x45, 0x78,
	0x70, 0x6f, 0x72, 0x74, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x64, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x77,
	0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x43, 0x6f, 0x64, 0x65,
	0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x39, 0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x64,
	0x61, 0x70, 0x70, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x50,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x4b, 0x65, 0x79, 0x22, 0x93, 0x01, 0x0a, 0x14, 0x53, 0x69, 0x67, 0x6e, 0x54, 0x78, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x0e,
	0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x61, 0x73, 0x68, 0x22, 0x78, 0x0a, 0x15, 0x53, 0x69, 0x67,
	0x6e, 0x54, 0x78, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2f, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x1b, 0x2e, 0x64, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x77, 0x61, 0x6c, 0x6c,
	0x65, 0x74, 0x2e, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x2a, 0x24, 0x0a, 0x0a, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07,
	0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x01, 0x32, 0xc5, 0x02, 0x0a, 0x0d, 0x57, 0x61,
	0x6c, 0x6c, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x66, 0x0a, 0x11, 0x67,
	0x65, 0x74, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x69, 0x67, 0x6e, 0x57, 0x61, 0x79,
	0x12, 0x26, 0x2e, 0x64, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x77, 0x61, 0x6c, 0x6c,
	0x65, 0x74, 0x2e, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x69, 0x67, 0x6e, 0x57, 0x61,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x64, 0x61, 0x70, 0x70, 0x6c,
	0x69, 0x6e, 0x6b, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x53, 0x75, 0x70, 0x70, 0x6f,
	0x72, 0x74, 0x53, 0x69, 0x67, 0x6e, 0x57, 0x61, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x6a, 0x0a, 0x13, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x50, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x27, 0x2e, 0x64, 0x61, 0x70,
	0x70, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x45, 0x78, 0x70,
	0x6f, 0x72, 0x74, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x64, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x77,
	0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x50, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x60, 0x0a, 0x0d, 0x73, 0x69, 0x67, 0x6e, 0x54, 0x78, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x25, 0x2e, 0x64, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x77, 0x61, 0x6c, 0x6c,
	0x65, 0x74, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x54, 0x78, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x64, 0x61, 0x70, 0x70, 0x6c, 0x69,
	0x6e, 0x6b, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x54, 0x78,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x13, 0x5a, 0x11, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_protobuf_wallet_proto_rawDescOnce sync.Once
	file_protobuf_wallet_proto_rawDescData []byte
)

func file_protobuf_wallet_proto_rawDescGZIP() []byte {
	file_protobuf_wallet_proto_rawDescOnce.Do(func() {
		file_protobuf_wallet_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_protobuf_wallet_proto_rawDesc), len(file_protobuf_wallet_proto_rawDesc)))
	})
	return file_protobuf_wallet_proto_rawDescData
}

var file_protobuf_wallet_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_protobuf_wallet_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_protobuf_wallet_proto_goTypes = []any{
	(ReturnCode)(0),                 // 0: dapplink.wallet.ReturnCode
	(*PublicKey)(nil),               // 1: dapplink.wallet.PublicKey
	(*SupportSignWayRequest)(nil),   // 2: dapplink.wallet.SupportSignWayRequest
	(*SignWay)(nil),                 // 3: dapplink.wallet.SignWay
	(*SupportSignWayResponse)(nil),  // 4: dapplink.wallet.SupportSignWayResponse
	(*ExportPublicKeyRequest)(nil),  // 5: dapplink.wallet.ExportPublicKeyRequest
	(*ExportPublicKeyResponse)(nil), // 6: dapplink.wallet.ExportPublicKeyResponse
	(*SignTxMessageRequest)(nil),    // 7: dapplink.wallet.SignTxMessageRequest
	(*SignTxMessageResponse)(nil),   // 8: dapplink.wallet.SignTxMessageResponse
}
var file_protobuf_wallet_proto_depIdxs = []int32{
	0, // 0: dapplink.wallet.SupportSignWayResponse.Code:type_name -> dapplink.wallet.ReturnCode
	3, // 1: dapplink.wallet.SupportSignWayResponse.sign_way:type_name -> dapplink.wallet.SignWay
	0, // 2: dapplink.wallet.ExportPublicKeyResponse.Code:type_name -> dapplink.wallet.ReturnCode
	1, // 3: dapplink.wallet.ExportPublicKeyResponse.public_key:type_name -> dapplink.wallet.PublicKey
	0, // 4: dapplink.wallet.SignTxMessageResponse.Code:type_name -> dapplink.wallet.ReturnCode
	2, // 5: dapplink.wallet.WalletService.getSupportSignWay:input_type -> dapplink.wallet.SupportSignWayRequest
	5, // 6: dapplink.wallet.WalletService.exportPublicKeyList:input_type -> dapplink.wallet.ExportPublicKeyRequest
	7, // 7: dapplink.wallet.WalletService.signTxMessage:input_type -> dapplink.wallet.SignTxMessageRequest
	4, // 8: dapplink.wallet.WalletService.getSupportSignWay:output_type -> dapplink.wallet.SupportSignWayResponse
	6, // 9: dapplink.wallet.WalletService.exportPublicKeyList:output_type -> dapplink.wallet.ExportPublicKeyResponse
	8, // 10: dapplink.wallet.WalletService.signTxMessage:output_type -> dapplink.wallet.SignTxMessageResponse
	8, // [8:11] is the sub-list for method output_type
	5, // [5:8] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_protobuf_wallet_proto_init() }
func file_protobuf_wallet_proto_init() {
	if File_protobuf_wallet_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_protobuf_wallet_proto_rawDesc), len(file_protobuf_wallet_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protobuf_wallet_proto_goTypes,
		DependencyIndexes: file_protobuf_wallet_proto_depIdxs,
		EnumInfos:         file_protobuf_wallet_proto_enumTypes,
		MessageInfos:      file_protobuf_wallet_proto_msgTypes,
	}.Build()
	File_protobuf_wallet_proto = out.File
	file_protobuf_wallet_proto_goTypes = nil
	file_protobuf_wallet_proto_depIdxs = nil
}
