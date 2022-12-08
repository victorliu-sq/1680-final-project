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

type HelloMsg struct {
	MsgType              uint32   `protobuf:"varint,1,opt,name=msgType,proto3" json:"msgType,omitempty"`
	UdpPort              uint32   `protobuf:"varint,2,opt,name=udpPort,proto3" json:"udpPort,omitempty"`
	Adderss              string   `protobuf:"bytes,3,opt,name=adderss,proto3" json:"adderss,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloMsg) Reset()         { *m = HelloMsg{} }
func (m *HelloMsg) String() string { return proto.CompactTextString(m) }
func (*HelloMsg) ProtoMessage()    {}
func (*HelloMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4b9afc808729447, []int{0}
}

func (m *HelloMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloMsg.Unmarshal(m, b)
}
func (m *HelloMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloMsg.Marshal(b, m, deterministic)
}
func (m *HelloMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloMsg.Merge(m, src)
}
func (m *HelloMsg) XXX_Size() int {
	return xxx_messageInfo_HelloMsg.Size(m)
}
func (m *HelloMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloMsg.DiscardUnknown(m)
}

var xxx_messageInfo_HelloMsg proto.InternalMessageInfo

func (m *HelloMsg) GetMsgType() uint32 {
	if m != nil {
		return m.MsgType
	}
	return 0
}

func (m *HelloMsg) GetUdpPort() uint32 {
	if m != nil {
		return m.UdpPort
	}
	return 0
}

func (m *HelloMsg) GetAdderss() string {
	if m != nil {
		return m.Adderss
	}
	return ""
}

type WelcomeMsg struct {
	MsgType              uint32   `protobuf:"varint,1,opt,name=msgType,proto3" json:"msgType,omitempty"`
	SongNum              uint32   `protobuf:"varint,2,opt,name=songNum,proto3" json:"songNum,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WelcomeMsg) Reset()         { *m = WelcomeMsg{} }
func (m *WelcomeMsg) String() string { return proto.CompactTextString(m) }
func (*WelcomeMsg) ProtoMessage()    {}
func (*WelcomeMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4b9afc808729447, []int{1}
}

func (m *WelcomeMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WelcomeMsg.Unmarshal(m, b)
}
func (m *WelcomeMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WelcomeMsg.Marshal(b, m, deterministic)
}
func (m *WelcomeMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WelcomeMsg.Merge(m, src)
}
func (m *WelcomeMsg) XXX_Size() int {
	return xxx_messageInfo_WelcomeMsg.Size(m)
}
func (m *WelcomeMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_WelcomeMsg.DiscardUnknown(m)
}

var xxx_messageInfo_WelcomeMsg proto.InternalMessageInfo

func (m *WelcomeMsg) GetMsgType() uint32 {
	if m != nil {
		return m.MsgType
	}
	return 0
}

func (m *WelcomeMsg) GetSongNum() uint32 {
	if m != nil {
		return m.SongNum
	}
	return 0
}

func init() {
	proto.RegisterType((*HelloMsg)(nil), "rpcMsg.HelloMsg")
	proto.RegisterType((*WelcomeMsg)(nil), "rpcMsg.WelcomeMsg")
}

func init() { proto.RegisterFile("rpcMsg.proto", fileDescriptor_a4b9afc808729447) }

var fileDescriptor_a4b9afc808729447 = []byte{
	// 188 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x2a, 0x48, 0xf6,
	0x2d, 0x4e, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x83, 0xf0, 0x94, 0x22, 0xb8, 0x38,
	0x3c, 0x52, 0x73, 0x72, 0xf2, 0x7d, 0x8b, 0xd3, 0x85, 0x24, 0xb8, 0xd8, 0x73, 0x8b, 0xd3, 0x43,
	0x2a, 0x0b, 0x52, 0x25, 0x18, 0x15, 0x18, 0x35, 0x78, 0x83, 0x60, 0x5c, 0x90, 0x4c, 0x69, 0x4a,
	0x41, 0x40, 0x7e, 0x51, 0x89, 0x04, 0x13, 0x44, 0x06, 0xca, 0x05, 0xc9, 0x24, 0xa6, 0xa4, 0xa4,
	0x16, 0x15, 0x17, 0x4b, 0x30, 0x2b, 0x30, 0x6a, 0x70, 0x06, 0xc1, 0xb8, 0x4a, 0x0e, 0x5c, 0x5c,
	0xe1, 0xa9, 0x39, 0xc9, 0xf9, 0xb9, 0xa9, 0x04, 0xcd, 0x2e, 0xce, 0xcf, 0x4b, 0xf7, 0x2b, 0xcd,
	0x85, 0x99, 0x0d, 0xe5, 0x1a, 0xb9, 0x70, 0x71, 0x39, 0xe7, 0xe7, 0x95, 0x14, 0xe5, 0xe7, 0x80,
	0x4c, 0x30, 0xe3, 0xe2, 0x09, 0x4e, 0xcd, 0x4b, 0x81, 0xbb, 0x56, 0x40, 0x0f, 0xea, 0x21, 0x98,
	0x88, 0x94, 0x10, 0x4c, 0x04, 0x61, 0xaf, 0x12, 0x83, 0x93, 0x68, 0x94, 0x70, 0x41, 0x76, 0xba,
	0x3e, 0xd8, 0xdb, 0xc9, 0xf9, 0x39, 0xfa, 0x10, 0x35, 0x49, 0x6c, 0x60, 0x01, 0x63, 0x40, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x96, 0x6a, 0xb7, 0xac, 0x17, 0x01, 0x00, 0x00,
}
