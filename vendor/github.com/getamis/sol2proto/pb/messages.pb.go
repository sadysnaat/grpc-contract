// Code generated by protoc-gen-go. DO NOT EDIT.
// source: messages.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	messages.proto

It has these top-level messages:
	Empty
	TransactOpts
	TransactionReq
	TransactionResp
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type TransactOpts struct {
	FromAddress string `protobuf:"bytes,1,opt,name=from_address,json=fromAddress" json:"from_address,omitempty"`
	PrivateKey  string `protobuf:"bytes,2,opt,name=private_key,json=privateKey" json:"private_key,omitempty"`
	Nonce       int64  `protobuf:"varint,3,opt,name=nonce" json:"nonce,omitempty"`
	Value       int64  `protobuf:"varint,4,opt,name=value" json:"value,omitempty"`
	GasPrice    int64  `protobuf:"varint,5,opt,name=gas_price,json=gasPrice" json:"gas_price,omitempty"`
	GasLimit    int64  `protobuf:"varint,6,opt,name=gas_limit,json=gasLimit" json:"gas_limit,omitempty"`
}

func (m *TransactOpts) Reset()                    { *m = TransactOpts{} }
func (m *TransactOpts) String() string            { return proto.CompactTextString(m) }
func (*TransactOpts) ProtoMessage()               {}
func (*TransactOpts) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *TransactOpts) GetFromAddress() string {
	if m != nil {
		return m.FromAddress
	}
	return ""
}

func (m *TransactOpts) GetPrivateKey() string {
	if m != nil {
		return m.PrivateKey
	}
	return ""
}

func (m *TransactOpts) GetNonce() int64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *TransactOpts) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *TransactOpts) GetGasPrice() int64 {
	if m != nil {
		return m.GasPrice
	}
	return 0
}

func (m *TransactOpts) GetGasLimit() int64 {
	if m != nil {
		return m.GasLimit
	}
	return 0
}

type TransactionReq struct {
	Opts *TransactOpts `protobuf:"bytes,1,opt,name=opts" json:"opts,omitempty"`
}

func (m *TransactionReq) Reset()                    { *m = TransactionReq{} }
func (m *TransactionReq) String() string            { return proto.CompactTextString(m) }
func (*TransactionReq) ProtoMessage()               {}
func (*TransactionReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *TransactionReq) GetOpts() *TransactOpts {
	if m != nil {
		return m.Opts
	}
	return nil
}

type TransactionResp struct {
	TxHash string `protobuf:"bytes,1,opt,name=tx_hash,json=txHash" json:"tx_hash,omitempty"`
}

func (m *TransactionResp) Reset()                    { *m = TransactionResp{} }
func (m *TransactionResp) String() string            { return proto.CompactTextString(m) }
func (*TransactionResp) ProtoMessage()               {}
func (*TransactionResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *TransactionResp) GetTxHash() string {
	if m != nil {
		return m.TxHash
	}
	return ""
}

func init() {
	proto.RegisterType((*Empty)(nil), "pb.Empty")
	proto.RegisterType((*TransactOpts)(nil), "pb.TransactOpts")
	proto.RegisterType((*TransactionReq)(nil), "pb.TransactionReq")
	proto.RegisterType((*TransactionResp)(nil), "pb.TransactionResp")
}

func init() { proto.RegisterFile("messages.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 250 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x40, 0x49, 0xdb, 0xa4, 0x76, 0x52, 0xaa, 0x2c, 0x82, 0x0b, 0x1e, 0xac, 0xc1, 0x43, 0xf1,
	0x90, 0x83, 0x82, 0x77, 0x0f, 0x82, 0xa0, 0xa0, 0x04, 0xef, 0x61, 0x92, 0x8e, 0xc9, 0x62, 0x93,
	0x5d, 0x77, 0xd6, 0xd2, 0xfc, 0x99, 0x9f, 0x27, 0xd9, 0xb4, 0xa2, 0xc7, 0xf7, 0xde, 0x2e, 0xcc,
	0x0c, 0x2c, 0x1a, 0x62, 0xc6, 0x8a, 0x38, 0x35, 0x56, 0x3b, 0x2d, 0x46, 0xa6, 0x48, 0xa6, 0x10,
	0x3e, 0x34, 0xc6, 0x75, 0xc9, 0x77, 0x00, 0xf3, 0x37, 0x8b, 0x2d, 0x63, 0xe9, 0x5e, 0x8c, 0x63,
	0x71, 0x09, 0xf3, 0x77, 0xab, 0x9b, 0x1c, 0xd7, 0x6b, 0x4b, 0xcc, 0x32, 0x58, 0x06, 0xab, 0x59,
	0x16, 0xf7, 0xee, 0x7e, 0x50, 0xe2, 0x02, 0x62, 0x63, 0xd5, 0x16, 0x1d, 0xe5, 0x1f, 0xd4, 0xc9,
	0x91, 0x7f, 0x01, 0x7b, 0xf5, 0x44, 0x9d, 0x38, 0x85, 0xb0, 0xd5, 0x6d, 0x49, 0x72, 0xbc, 0x0c,
	0x56, 0xe3, 0x6c, 0x80, 0xde, 0x6e, 0x71, 0xf3, 0x45, 0x72, 0x32, 0x58, 0x0f, 0xe2, 0x1c, 0x66,
	0x15, 0x72, 0x6e, 0xac, 0x2a, 0x49, 0x86, 0xbe, 0x1c, 0x55, 0xc8, 0xaf, 0x3d, 0x1f, 0xe2, 0x46,
	0x35, 0xca, 0xc9, 0xe8, 0x37, 0x3e, 0xf7, 0x9c, 0xdc, 0xc1, 0xe2, 0x30, 0xb9, 0xd2, 0x6d, 0x46,
	0x9f, 0xe2, 0x0a, 0x26, 0xda, 0xb8, 0x61, 0xe6, 0xf8, 0xe6, 0x24, 0x35, 0x45, 0xfa, 0x77, 0xb7,
	0xcc, 0xd7, 0xe4, 0x1a, 0x8e, 0xff, 0xfd, 0x63, 0x23, 0xce, 0x60, 0xea, 0x76, 0x79, 0x8d, 0x5c,
	0xef, 0xf7, 0x8d, 0xdc, 0xee, 0x11, 0xb9, 0x2e, 0x22, 0x7f, 0xb2, 0xdb, 0x9f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x94, 0xa6, 0xba, 0x66, 0x44, 0x01, 0x00, 0x00,
}
