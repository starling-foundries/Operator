// Code generated by protoc-gen-go. DO NOT EDIT.
// source: check.proto

package twirp_check

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type StatusResp_Status int32

const (
	StatusResp_UNDEFINED StatusResp_Status = 0
	StatusResp_RECIEVED  StatusResp_Status = 1
	StatusResp_PROCESSED StatusResp_Status = 2
	StatusResp_REJECTED  StatusResp_Status = 3
)

var StatusResp_Status_name = map[int32]string{
	0: "UNDEFINED",
	1: "RECIEVED",
	2: "PROCESSED",
	3: "REJECTED",
}

var StatusResp_Status_value = map[string]int32{
	"UNDEFINED": 0,
	"RECIEVED":  1,
	"PROCESSED": 2,
	"REJECTED":  3,
}

func (x StatusResp_Status) String() string {
	return proto.EnumName(StatusResp_Status_name, int32(x))
}

func (StatusResp_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d8d3c606fb107336, []int{1, 0}
}

type ByteArray struct {
	Data                 []byte   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ByteArray) Reset()         { *m = ByteArray{} }
func (m *ByteArray) String() string { return proto.CompactTextString(m) }
func (*ByteArray) ProtoMessage()    {}
func (*ByteArray) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8d3c606fb107336, []int{0}
}

func (m *ByteArray) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ByteArray.Unmarshal(m, b)
}
func (m *ByteArray) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ByteArray.Marshal(b, m, deterministic)
}
func (m *ByteArray) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ByteArray.Merge(m, src)
}
func (m *ByteArray) XXX_Size() int {
	return xxx_messageInfo_ByteArray.Size(m)
}
func (m *ByteArray) XXX_DiscardUnknown() {
	xxx_messageInfo_ByteArray.DiscardUnknown(m)
}

var xxx_messageInfo_ByteArray proto.InternalMessageInfo

func (m *ByteArray) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type StatusResp struct {
	SenderPkey *ByteArray `protobuf:"bytes,1,opt,name=sender_pkey,json=senderPkey,proto3" json:"sender_pkey,omitempty"`
	//tranid should be the chain-addressable transaction id for use in ex: viewblock.io. it is set to nil if this is unknown.
	Tranid               []byte            `protobuf:"bytes,2,opt,name=tranid,proto3" json:"tranid,omitempty"`
	Status               StatusResp_Status `protobuf:"varint,3,opt,name=status,proto3,enum=twirp.check.StatusResp_Status" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *StatusResp) Reset()         { *m = StatusResp{} }
func (m *StatusResp) String() string { return proto.CompactTextString(m) }
func (*StatusResp) ProtoMessage()    {}
func (*StatusResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8d3c606fb107336, []int{1}
}

func (m *StatusResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatusResp.Unmarshal(m, b)
}
func (m *StatusResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatusResp.Marshal(b, m, deterministic)
}
func (m *StatusResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatusResp.Merge(m, src)
}
func (m *StatusResp) XXX_Size() int {
	return xxx_messageInfo_StatusResp.Size(m)
}
func (m *StatusResp) XXX_DiscardUnknown() {
	xxx_messageInfo_StatusResp.DiscardUnknown(m)
}

var xxx_messageInfo_StatusResp proto.InternalMessageInfo

func (m *StatusResp) GetSenderPkey() *ByteArray {
	if m != nil {
		return m.SenderPkey
	}
	return nil
}

func (m *StatusResp) GetTranid() []byte {
	if m != nil {
		return m.Tranid
	}
	return nil
}

func (m *StatusResp) GetStatus() StatusResp_Status {
	if m != nil {
		return m.Status
	}
	return StatusResp_UNDEFINED
}

//Empty only requires the calling user's pubkey
type Empty struct {
	Pubkey               *ByteArray `protobuf:"bytes,1,opt,name=pubkey,proto3" json:"pubkey,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8d3c606fb107336, []int{2}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

func (m *Empty) GetPubkey() *ByteArray {
	if m != nil {
		return m.Pubkey
	}
	return nil
}

//indexes the sent transaction's status within the context of the operator.
type QueueID struct {
	QueueID              uint32   `protobuf:"varint,1,opt,name=queueID,proto3" json:"queueID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueueID) Reset()         { *m = QueueID{} }
func (m *QueueID) String() string { return proto.CompactTextString(m) }
func (*QueueID) ProtoMessage()    {}
func (*QueueID) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8d3c606fb107336, []int{3}
}

func (m *QueueID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueueID.Unmarshal(m, b)
}
func (m *QueueID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueueID.Marshal(b, m, deterministic)
}
func (m *QueueID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueueID.Merge(m, src)
}
func (m *QueueID) XXX_Size() int {
	return xxx_messageInfo_QueueID.Size(m)
}
func (m *QueueID) XXX_DiscardUnknown() {
	xxx_messageInfo_QueueID.DiscardUnknown(m)
}

var xxx_messageInfo_QueueID proto.InternalMessageInfo

func (m *QueueID) GetQueueID() uint32 {
	if m != nil {
		return m.QueueID
	}
	return 0
}

//this object is injected straight into the code field of the encapsulating transaction
type CheckTransactionCoreInfo struct {
	Senderpubkey *ByteArray `protobuf:"bytes,1,opt,name=senderpubkey,proto3" json:"senderpubkey,omitempty"`
	Toaddr       []byte     `protobuf:"bytes,2,opt,name=toaddr,proto3" json:"toaddr,omitempty"`
	Amount       *ByteArray `protobuf:"bytes,3,opt,name=amount,proto3" json:"amount,omitempty"`
	Fee          *ByteArray `protobuf:"bytes,4,opt,name=fee,proto3" json:"fee,omitempty"`
	OperatorAddr []byte     `protobuf:"bytes,5,opt,name=operatorAddr,proto3" json:"operatorAddr,omitempty"`
	//this nonce is account-specific within the scilla contract. The encapsulating transaction will have a nonce from the operator's wallet. this nonce should always be increasing
	Nonce                uint64     `protobuf:"varint,6,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Signature            *ByteArray `protobuf:"bytes,7,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *CheckTransactionCoreInfo) Reset()         { *m = CheckTransactionCoreInfo{} }
func (m *CheckTransactionCoreInfo) String() string { return proto.CompactTextString(m) }
func (*CheckTransactionCoreInfo) ProtoMessage()    {}
func (*CheckTransactionCoreInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8d3c606fb107336, []int{4}
}

func (m *CheckTransactionCoreInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckTransactionCoreInfo.Unmarshal(m, b)
}
func (m *CheckTransactionCoreInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckTransactionCoreInfo.Marshal(b, m, deterministic)
}
func (m *CheckTransactionCoreInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckTransactionCoreInfo.Merge(m, src)
}
func (m *CheckTransactionCoreInfo) XXX_Size() int {
	return xxx_messageInfo_CheckTransactionCoreInfo.Size(m)
}
func (m *CheckTransactionCoreInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckTransactionCoreInfo.DiscardUnknown(m)
}

var xxx_messageInfo_CheckTransactionCoreInfo proto.InternalMessageInfo

func (m *CheckTransactionCoreInfo) GetSenderpubkey() *ByteArray {
	if m != nil {
		return m.Senderpubkey
	}
	return nil
}

func (m *CheckTransactionCoreInfo) GetToaddr() []byte {
	if m != nil {
		return m.Toaddr
	}
	return nil
}

func (m *CheckTransactionCoreInfo) GetAmount() *ByteArray {
	if m != nil {
		return m.Amount
	}
	return nil
}

func (m *CheckTransactionCoreInfo) GetFee() *ByteArray {
	if m != nil {
		return m.Fee
	}
	return nil
}

func (m *CheckTransactionCoreInfo) GetOperatorAddr() []byte {
	if m != nil {
		return m.OperatorAddr
	}
	return nil
}

func (m *CheckTransactionCoreInfo) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *CheckTransactionCoreInfo) GetSignature() *ByteArray {
	if m != nil {
		return m.Signature
	}
	return nil
}

type LastNonceResp struct {
	Nonce                uint64   `protobuf:"varint,1,opt,name=nonce,proto3" json:"nonce,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LastNonceResp) Reset()         { *m = LastNonceResp{} }
func (m *LastNonceResp) String() string { return proto.CompactTextString(m) }
func (*LastNonceResp) ProtoMessage()    {}
func (*LastNonceResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8d3c606fb107336, []int{5}
}

func (m *LastNonceResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LastNonceResp.Unmarshal(m, b)
}
func (m *LastNonceResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LastNonceResp.Marshal(b, m, deterministic)
}
func (m *LastNonceResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LastNonceResp.Merge(m, src)
}
func (m *LastNonceResp) XXX_Size() int {
	return xxx_messageInfo_LastNonceResp.Size(m)
}
func (m *LastNonceResp) XXX_DiscardUnknown() {
	xxx_messageInfo_LastNonceResp.DiscardUnknown(m)
}

var xxx_messageInfo_LastNonceResp proto.InternalMessageInfo

func (m *LastNonceResp) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

type CheckTransaction struct {
	Version              uint32                    `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	Info                 *CheckTransactionCoreInfo `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *CheckTransaction) Reset()         { *m = CheckTransaction{} }
func (m *CheckTransaction) String() string { return proto.CompactTextString(m) }
func (*CheckTransaction) ProtoMessage()    {}
func (*CheckTransaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8d3c606fb107336, []int{6}
}

func (m *CheckTransaction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckTransaction.Unmarshal(m, b)
}
func (m *CheckTransaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckTransaction.Marshal(b, m, deterministic)
}
func (m *CheckTransaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckTransaction.Merge(m, src)
}
func (m *CheckTransaction) XXX_Size() int {
	return xxx_messageInfo_CheckTransaction.Size(m)
}
func (m *CheckTransaction) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckTransaction.DiscardUnknown(m)
}

var xxx_messageInfo_CheckTransaction proto.InternalMessageInfo

func (m *CheckTransaction) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *CheckTransaction) GetInfo() *CheckTransactionCoreInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

type CheckStatusResp struct {
	Receipt              []byte     `protobuf:"bytes,1,opt,name=receipt,proto3" json:"receipt,omitempty"`
	Fee                  *ByteArray `protobuf:"bytes,2,opt,name=fee,proto3" json:"fee,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *CheckStatusResp) Reset()         { *m = CheckStatusResp{} }
func (m *CheckStatusResp) String() string { return proto.CompactTextString(m) }
func (*CheckStatusResp) ProtoMessage()    {}
func (*CheckStatusResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8d3c606fb107336, []int{7}
}

func (m *CheckStatusResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckStatusResp.Unmarshal(m, b)
}
func (m *CheckStatusResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckStatusResp.Marshal(b, m, deterministic)
}
func (m *CheckStatusResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckStatusResp.Merge(m, src)
}
func (m *CheckStatusResp) XXX_Size() int {
	return xxx_messageInfo_CheckStatusResp.Size(m)
}
func (m *CheckStatusResp) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckStatusResp.DiscardUnknown(m)
}

var xxx_messageInfo_CheckStatusResp proto.InternalMessageInfo

func (m *CheckStatusResp) GetReceipt() []byte {
	if m != nil {
		return m.Receipt
	}
	return nil
}

func (m *CheckStatusResp) GetFee() *ByteArray {
	if m != nil {
		return m.Fee
	}
	return nil
}

//responds with the address of the on-chain contract this server is operating.
type OperatorTargetResp struct {
	ContractAddr         []byte   `protobuf:"bytes,1,opt,name=contractAddr,proto3" json:"contractAddr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OperatorTargetResp) Reset()         { *m = OperatorTargetResp{} }
func (m *OperatorTargetResp) String() string { return proto.CompactTextString(m) }
func (*OperatorTargetResp) ProtoMessage()    {}
func (*OperatorTargetResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8d3c606fb107336, []int{8}
}

func (m *OperatorTargetResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OperatorTargetResp.Unmarshal(m, b)
}
func (m *OperatorTargetResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OperatorTargetResp.Marshal(b, m, deterministic)
}
func (m *OperatorTargetResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OperatorTargetResp.Merge(m, src)
}
func (m *OperatorTargetResp) XXX_Size() int {
	return xxx_messageInfo_OperatorTargetResp.Size(m)
}
func (m *OperatorTargetResp) XXX_DiscardUnknown() {
	xxx_messageInfo_OperatorTargetResp.DiscardUnknown(m)
}

var xxx_messageInfo_OperatorTargetResp proto.InternalMessageInfo

func (m *OperatorTargetResp) GetContractAddr() []byte {
	if m != nil {
		return m.ContractAddr
	}
	return nil
}

func init() {
	proto.RegisterEnum("twirp.check.StatusResp_Status", StatusResp_Status_name, StatusResp_Status_value)
	proto.RegisterType((*ByteArray)(nil), "twirp.check.ByteArray")
	proto.RegisterType((*StatusResp)(nil), "twirp.check.StatusResp")
	proto.RegisterType((*Empty)(nil), "twirp.check.Empty")
	proto.RegisterType((*QueueID)(nil), "twirp.check.queueID")
	proto.RegisterType((*CheckTransactionCoreInfo)(nil), "twirp.check.CheckTransactionCoreInfo")
	proto.RegisterType((*LastNonceResp)(nil), "twirp.check.LastNonceResp")
	proto.RegisterType((*CheckTransaction)(nil), "twirp.check.CheckTransaction")
	proto.RegisterType((*CheckStatusResp)(nil), "twirp.check.CheckStatusResp")
	proto.RegisterType((*OperatorTargetResp)(nil), "twirp.check.OperatorTargetResp")
}

func init() { proto.RegisterFile("check.proto", fileDescriptor_d8d3c606fb107336) }

var fileDescriptor_d8d3c606fb107336 = []byte{
	// 532 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xdb, 0x6e, 0xd3, 0x40,
	0x10, 0xc5, 0xb9, 0x38, 0x64, 0x92, 0xd0, 0x68, 0x54, 0x15, 0xab, 0x12, 0x10, 0x2d, 0xaa, 0x94,
	0xa7, 0x3c, 0x04, 0x44, 0xb8, 0x48, 0x48, 0x6d, 0xbc, 0x48, 0x41, 0x28, 0xad, 0x9c, 0x94, 0x57,
	0xb4, 0xb5, 0x37, 0xc1, 0x8a, 0xba, 0x6b, 0xd6, 0x6b, 0x50, 0xfe, 0x8c, 0x17, 0x3e, 0x84, 0xbf,
	0x41, 0x5e, 0x5f, 0x1a, 0x23, 0x92, 0xf2, 0xb6, 0x67, 0xe6, 0xec, 0x19, 0xcf, 0x99, 0x1d, 0x43,
	0xc7, 0xff, 0xca, 0xfd, 0xcd, 0x28, 0x52, 0x52, 0x4b, 0xec, 0xe8, 0x1f, 0xa1, 0x8a, 0x46, 0x26,
	0x44, 0x9e, 0x41, 0xfb, 0x62, 0xab, 0xf9, 0xb9, 0x52, 0x6c, 0x8b, 0x08, 0x8d, 0x80, 0x69, 0xe6,
	0x58, 0x03, 0x6b, 0xd8, 0xf5, 0xcc, 0x99, 0xfc, 0xb6, 0x00, 0x16, 0x9a, 0xe9, 0x24, 0xf6, 0x78,
	0x1c, 0xe1, 0x04, 0x3a, 0x31, 0x17, 0x01, 0x57, 0x5f, 0xa2, 0x0d, 0xdf, 0x1a, 0x66, 0x67, 0x7c,
	0x32, 0xda, 0x91, 0x1c, 0x95, 0x7a, 0x1e, 0x64, 0xd4, 0xab, 0x0d, 0xdf, 0xe2, 0x09, 0xd8, 0x5a,
	0x31, 0x11, 0x06, 0x4e, 0xcd, 0xa8, 0xe7, 0x08, 0x5f, 0x81, 0x1d, 0x1b, 0x79, 0xa7, 0x3e, 0xb0,
	0x86, 0x8f, 0xc6, 0x4f, 0x2b, 0x5a, 0x77, 0x95, 0x8b, 0x63, 0xce, 0x26, 0x17, 0x60, 0x67, 0x11,
	0xec, 0x41, 0xfb, 0x7a, 0xee, 0xd2, 0x0f, 0xb3, 0x39, 0x75, 0xfb, 0x0f, 0xb0, 0x0b, 0x0f, 0x3d,
	0x3a, 0x9d, 0xd1, 0xcf, 0xd4, 0xed, 0x5b, 0x69, 0xf2, 0xca, 0xbb, 0x9c, 0xd2, 0xc5, 0x82, 0xba,
	0xfd, 0x5a, 0x96, 0xfc, 0x48, 0xa7, 0x4b, 0xea, 0xf6, 0xeb, 0x64, 0x02, 0x4d, 0x7a, 0x1b, 0xe9,
	0x2d, 0x8e, 0xc0, 0x8e, 0x92, 0x9b, 0xfb, 0x1b, 0xca, 0x59, 0xe4, 0x39, 0xb4, 0xbe, 0x25, 0x3c,
	0xe1, 0x33, 0x17, 0x9d, 0xf2, 0x68, 0xee, 0xf6, 0xbc, 0x02, 0x92, 0x9f, 0x35, 0x70, 0xa6, 0xa9,
	0xc0, 0x52, 0x31, 0x11, 0x33, 0x5f, 0x87, 0x52, 0x4c, 0xa5, 0xe2, 0x33, 0xb1, 0x92, 0xf8, 0x16,
	0xba, 0x99, 0x39, 0xff, 0x55, 0xb7, 0xc2, 0x35, 0x56, 0x4a, 0x16, 0x04, 0xaa, 0xb4, 0xd2, 0xa0,
	0xb4, 0x0b, 0x76, 0x2b, 0x13, 0xa1, 0x8d, 0x95, 0x07, 0xba, 0xc8, 0x58, 0x38, 0x84, 0xfa, 0x8a,
	0x73, 0xa7, 0x71, 0x90, 0x9c, 0x52, 0x90, 0x40, 0x57, 0x46, 0x5c, 0x31, 0x2d, 0xd5, 0x79, 0x5a,
	0xb7, 0x69, 0xea, 0x56, 0x62, 0x78, 0x0c, 0x4d, 0x21, 0x85, 0xcf, 0x1d, 0x7b, 0x60, 0x0d, 0x1b,
	0x5e, 0x06, 0xf0, 0x25, 0xb4, 0xe3, 0x70, 0x2d, 0x98, 0x4e, 0x14, 0x77, 0x5a, 0x07, 0x2b, 0xdd,
	0x11, 0xc9, 0x19, 0xf4, 0x3e, 0xb1, 0x58, 0xcf, 0x53, 0x09, 0xf3, 0xec, 0x4a, 0x71, 0x6b, 0x47,
	0x9c, 0xac, 0xa1, 0xff, 0xb7, 0xc1, 0xe9, 0x3c, 0xbe, 0x73, 0x15, 0x87, 0x52, 0x14, 0xf3, 0xc8,
	0x21, 0xbe, 0x81, 0x46, 0x28, 0x56, 0xd2, 0x98, 0xd6, 0x19, 0x9f, 0x55, 0xbe, 0x62, 0xdf, 0x9c,
	0x3c, 0x73, 0x85, 0x5c, 0xc3, 0x91, 0x61, 0xec, 0x2c, 0x82, 0x03, 0x2d, 0xc5, 0x7d, 0x1e, 0x46,
	0x3a, 0x5f, 0x97, 0x02, 0x16, 0xb6, 0xd6, 0xee, 0xb5, 0x95, 0xbc, 0x06, 0xbc, 0xcc, 0x2d, 0x5c,
	0x32, 0xb5, 0xe6, 0xda, 0x28, 0x13, 0xe8, 0xfa, 0x52, 0x68, 0xc5, 0x7c, 0x6d, 0xcc, 0xce, 0xe4,
	0x2b, 0xb1, 0xf1, 0x2f, 0x0b, 0x9a, 0xe6, 0x8b, 0xf0, 0x1d, 0x34, 0x16, 0x5c, 0x04, 0xf8, 0xe4,
	0x60, 0x3f, 0xa7, 0xc7, 0x95, 0x74, 0xf1, 0x78, 0x27, 0xe5, 0x12, 0xfd, 0x33, 0x7f, 0xfa, 0x78,
	0xcf, 0x32, 0xe2, 0x7b, 0x38, 0x4a, 0x07, 0xb4, 0x6b, 0x3c, 0x56, 0xb8, 0x66, 0xaf, 0xf6, 0xde,
	0xbf, 0xb1, 0xcd, 0xaf, 0xe8, 0xc5, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x34, 0x38, 0xd7, 0xfb,
	0x99, 0x04, 0x00, 0x00,
}
