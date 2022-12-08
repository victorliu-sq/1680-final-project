// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpcMsg.proto

package rpcMsg

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

type RPCType int32

const (
	RPCType_REQUEST_HELLO    RPCType = 0
	RPCType_RESPONSE_WELCOME RPCType = 1
)

var RPCType_name = map[int32]string{
	0: "REQUEST_HELLO",
	1: "RESPONSE_WELCOME",
}

var RPCType_value = map[string]int32{
	"REQUEST_HELLO":    0,
	"RESPONSE_WELCOME": 1,
}

func (x RPCType) String() string {
	return proto.EnumName(RPCType_name, int32(x))
}

func (RPCType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a4b9afc808729447, []int{0}
}

type RequestHello struct {
	MsgType              uint32   `protobuf:"varint,1,opt,name=msgType,proto3" json:"msgType,omitempty"`
	UdpPort              uint32   `protobuf:"varint,2,opt,name=udpPort,proto3" json:"udpPort,omitempty"`
	Address              string   `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestHello) Reset()         { *m = RequestHello{} }
func (m *RequestHello) String() string { return proto.CompactTextString(m) }
func (*RequestHello) ProtoMessage()    {}
func (*RequestHello) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4b9afc808729447, []int{0}
}

func (m *RequestHello) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestHello.Unmarshal(m, b)
}
func (m *RequestHello) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestHello.Marshal(b, m, deterministic)
}
func (m *RequestHello) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestHello.Merge(m, src)
}
func (m *RequestHello) XXX_Size() int {
	return xxx_messageInfo_RequestHello.Size(m)
}
func (m *RequestHello) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestHello.DiscardUnknown(m)
}

var xxx_messageInfo_RequestHello proto.InternalMessageInfo

func (m *RequestHello) GetMsgType() uint32 {
	if m != nil {
		return m.MsgType
	}
	return 0
}

func (m *RequestHello) GetUdpPort() uint32 {
	if m != nil {
		return m.UdpPort
	}
	return 0
}

func (m *RequestHello) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type ResponseWelcome struct {
	MsgType              uint32   `protobuf:"varint,1,opt,name=msgType,proto3" json:"msgType,omitempty"`
	SongNum              uint32   `protobuf:"varint,2,opt,name=songNum,proto3" json:"songNum,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResponseWelcome) Reset()         { *m = ResponseWelcome{} }
func (m *ResponseWelcome) String() string { return proto.CompactTextString(m) }
func (*ResponseWelcome) ProtoMessage()    {}
func (*ResponseWelcome) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4b9afc808729447, []int{1}
}

func (m *ResponseWelcome) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseWelcome.Unmarshal(m, b)
}
func (m *ResponseWelcome) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseWelcome.Marshal(b, m, deterministic)
}
func (m *ResponseWelcome) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseWelcome.Merge(m, src)
}
func (m *ResponseWelcome) XXX_Size() int {
	return xxx_messageInfo_ResponseWelcome.Size(m)
}
func (m *ResponseWelcome) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseWelcome.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseWelcome proto.InternalMessageInfo

func (m *ResponseWelcome) GetMsgType() uint32 {
	if m != nil {
		return m.MsgType
	}
	return 0
}

func (m *ResponseWelcome) GetSongNum() uint32 {
	if m != nil {
		return m.SongNum
	}
	return 0
}

func init() {
	proto.RegisterEnum("rpcMsg.RPCType", RPCType_name, RPCType_value)
	proto.RegisterType((*RequestHello)(nil), "rpcMsg.RequestHello")
	proto.RegisterType((*ResponseWelcome)(nil), "rpcMsg.ResponseWelcome")
}

func init() { proto.RegisterFile("rpcMsg.proto", fileDescriptor_a4b9afc808729447) }

var fileDescriptor_a4b9afc808729447 = []byte{
	// 261 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xcf, 0x4a, 0xf3, 0x40,
	0x14, 0xc5, 0x9b, 0xef, 0x83, 0x16, 0x87, 0x56, 0xd3, 0x58, 0x31, 0xb8, 0x2a, 0x59, 0x15, 0x17,
	0x2d, 0xd4, 0x27, 0xd0, 0x30, 0x90, 0x45, 0xd2, 0xc4, 0x49, 0xb5, 0x20, 0x42, 0xa9, 0xc9, 0x65,
	0x10, 0x27, 0xb9, 0xe3, 0xdc, 0x44, 0xf0, 0xed, 0xa5, 0xf9, 0x83, 0xba, 0x71, 0xf9, 0x3b, 0xe7,
	0xc0, 0xb9, 0xf7, 0xb0, 0xb1, 0xd1, 0x59, 0x44, 0x72, 0xa9, 0x0d, 0x56, 0xe8, 0x0c, 0x5b, 0xf2,
	0x9e, 0xd9, 0x58, 0xc0, 0x7b, 0x0d, 0x54, 0x05, 0xa0, 0x14, 0x3a, 0x2e, 0x1b, 0x15, 0x24, 0xb7,
	0x9f, 0x1a, 0x5c, 0x6b, 0x6e, 0x2d, 0x26, 0xa2, 0xc7, 0xa3, 0x53, 0xe7, 0x3a, 0x41, 0x53, 0xb9,
	0xff, 0x5a, 0xa7, 0xc3, 0xa3, 0x73, 0xc8, 0x73, 0x03, 0x44, 0xee, 0xff, 0xb9, 0xb5, 0x38, 0x11,
	0x3d, 0x7a, 0x9c, 0x9d, 0x09, 0x20, 0x8d, 0x25, 0xc1, 0x0e, 0x54, 0x86, 0x05, 0xfc, 0x5d, 0x40,
	0x58, 0xca, 0x4d, 0x5d, 0xf4, 0x05, 0x1d, 0x5e, 0xaf, 0xd9, 0x48, 0x24, 0x7e, 0x13, 0x9a, 0xb2,
	0x89, 0xe0, 0xf7, 0x0f, 0x3c, 0xdd, 0xee, 0x03, 0x1e, 0x86, 0xb1, 0x3d, 0x70, 0x66, 0xcc, 0x16,
	0x3c, 0x4d, 0xe2, 0x4d, 0xca, 0xf7, 0x3b, 0x1e, 0xfa, 0x71, 0xc4, 0x6d, 0x6b, 0xfd, 0xc8, 0xa6,
	0x3e, 0x96, 0x95, 0x41, 0x15, 0x91, 0x4c, 0xc1, 0x7c, 0xbc, 0x66, 0xe0, 0xdc, 0xb2, 0xd3, 0xe0,
	0x50, 0xe6, 0x0a, 0x9a, 0x67, 0x23, 0x92, 0xce, 0x6c, 0xd9, 0xcd, 0xf2, 0x73, 0x85, 0xab, 0xcb,
	0x6f, 0xf5, 0xd7, 0xf5, 0xde, 0xe0, 0xee, 0xe2, 0xe9, 0x5c, 0xbf, 0xc9, 0x55, 0xb3, 0x62, 0x86,
	0x6a, 0xd5, 0x06, 0x5f, 0x86, 0x8d, 0x70, 0xf3, 0x15, 0x00, 0x00, 0xff, 0xff, 0x0e, 0x48, 0xc8,
	0x3d, 0x66, 0x01, 0x00, 0x00,
}