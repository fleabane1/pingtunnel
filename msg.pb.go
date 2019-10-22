// Code generated by protoc-gen-go. DO NOT EDIT.
// source: msg.proto

package pingtunnel

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

type MyMsg struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Type                 int32    `protobuf:"varint,2,opt,name=type,proto3" json:"type,omitempty"`
	Target               string   `protobuf:"bytes,3,opt,name=target,proto3" json:"target,omitempty"`
	Data                 []byte   `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	Rproto               int32    `protobuf:"varint,5,opt,name=rproto,proto3" json:"rproto,omitempty"`
	Catch                int32    `protobuf:"varint,6,opt,name=catch,proto3" json:"catch,omitempty"`
	Key                  int32    `protobuf:"varint,7,opt,name=key,proto3" json:"key,omitempty"`
	Tcpmode              int32    `protobuf:"varint,8,opt,name=tcpmode,proto3" json:"tcpmode,omitempty"`
	TcpmodeBuffersize    int32    `protobuf:"varint,9,opt,name=tcpmode_buffersize,json=tcpmodeBuffersize,proto3" json:"tcpmode_buffersize,omitempty"`
	TcpmodeMaxwin        int32    `protobuf:"varint,10,opt,name=tcpmode_maxwin,json=tcpmodeMaxwin,proto3" json:"tcpmode_maxwin,omitempty"`
	TcpmodeResendTimems  int32    `protobuf:"varint,11,opt,name=tcpmode_resend_timems,json=tcpmodeResendTimems,proto3" json:"tcpmode_resend_timems,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MyMsg) Reset()         { *m = MyMsg{} }
func (m *MyMsg) String() string { return proto.CompactTextString(m) }
func (*MyMsg) ProtoMessage()    {}
func (*MyMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{0}
}

func (m *MyMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MyMsg.Unmarshal(m, b)
}
func (m *MyMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MyMsg.Marshal(b, m, deterministic)
}
func (m *MyMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MyMsg.Merge(m, src)
}
func (m *MyMsg) XXX_Size() int {
	return xxx_messageInfo_MyMsg.Size(m)
}
func (m *MyMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_MyMsg.DiscardUnknown(m)
}

var xxx_messageInfo_MyMsg proto.InternalMessageInfo

func (m *MyMsg) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *MyMsg) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *MyMsg) GetTarget() string {
	if m != nil {
		return m.Target
	}
	return ""
}

func (m *MyMsg) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *MyMsg) GetRproto() int32 {
	if m != nil {
		return m.Rproto
	}
	return 0
}

func (m *MyMsg) GetCatch() int32 {
	if m != nil {
		return m.Catch
	}
	return 0
}

func (m *MyMsg) GetKey() int32 {
	if m != nil {
		return m.Key
	}
	return 0
}

func (m *MyMsg) GetTcpmode() int32 {
	if m != nil {
		return m.Tcpmode
	}
	return 0
}

func (m *MyMsg) GetTcpmodeBuffersize() int32 {
	if m != nil {
		return m.TcpmodeBuffersize
	}
	return 0
}

func (m *MyMsg) GetTcpmodeMaxwin() int32 {
	if m != nil {
		return m.TcpmodeMaxwin
	}
	return 0
}

func (m *MyMsg) GetTcpmodeResendTimems() int32 {
	if m != nil {
		return m.TcpmodeResendTimems
	}
	return 0
}

type Frame struct {
	Type                 int32    `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	Resend               bool     `protobuf:"varint,2,opt,name=resend,proto3" json:"resend,omitempty"`
	Sendtime             int64    `protobuf:"varint,3,opt,name=sendtime,proto3" json:"sendtime,omitempty"`
	Id                   int32    `protobuf:"varint,4,opt,name=id,proto3" json:"id,omitempty"`
	Data                 []byte   `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`
	Dataid               []int32  `protobuf:"varint,6,rep,packed,name=dataid,proto3" json:"dataid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Frame) Reset()         { *m = Frame{} }
func (m *Frame) String() string { return proto.CompactTextString(m) }
func (*Frame) ProtoMessage()    {}
func (*Frame) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{1}
}

func (m *Frame) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Frame.Unmarshal(m, b)
}
func (m *Frame) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Frame.Marshal(b, m, deterministic)
}
func (m *Frame) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Frame.Merge(m, src)
}
func (m *Frame) XXX_Size() int {
	return xxx_messageInfo_Frame.Size(m)
}
func (m *Frame) XXX_DiscardUnknown() {
	xxx_messageInfo_Frame.DiscardUnknown(m)
}

var xxx_messageInfo_Frame proto.InternalMessageInfo

func (m *Frame) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *Frame) GetResend() bool {
	if m != nil {
		return m.Resend
	}
	return false
}

func (m *Frame) GetSendtime() int64 {
	if m != nil {
		return m.Sendtime
	}
	return 0
}

func (m *Frame) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Frame) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Frame) GetDataid() []int32 {
	if m != nil {
		return m.Dataid
	}
	return nil
}

func init() {
	proto.RegisterType((*MyMsg)(nil), "MyMsg")
	proto.RegisterType((*Frame)(nil), "Frame")
}

func init() { proto.RegisterFile("msg.proto", fileDescriptor_c06e4cca6c2cc899) }

var fileDescriptor_c06e4cca6c2cc899 = []byte{
	// 302 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x91, 0xcd, 0x4e, 0x02, 0x31,
	0x10, 0xc7, 0xb3, 0xbb, 0x74, 0x81, 0x11, 0x89, 0x8e, 0x1f, 0x99, 0x78, 0xda, 0x90, 0x98, 0x70,
	0xd1, 0x83, 0xbe, 0x01, 0x07, 0x6f, 0x5c, 0x1a, 0x4f, 0x5e, 0x48, 0xa1, 0x65, 0x6d, 0xb4, 0xcb,
	0xa6, 0x5b, 0xa2, 0xf8, 0x0a, 0x3e, 0x92, 0x2f, 0x67, 0x3a, 0x14, 0x12, 0x4e, 0x9d, 0xff, 0x47,
	0xa6, 0xc9, 0x6f, 0x60, 0xe8, 0xba, 0xfa, 0xb1, 0xf5, 0x9b, 0xb0, 0x99, 0xfc, 0xe5, 0x20, 0xe6,
	0xbb, 0x79, 0x57, 0xe3, 0x18, 0x72, 0xab, 0x29, 0xab, 0xb2, 0xe9, 0x50, 0xe6, 0x56, 0x23, 0x42,
	0x2f, 0xec, 0x5a, 0x43, 0x79, 0x95, 0x4d, 0x85, 0xe4, 0x19, 0x6f, 0xa1, 0x0c, 0xca, 0xd7, 0x26,
	0x50, 0xc1, 0xbd, 0xa4, 0x62, 0x57, 0xab, 0xa0, 0xa8, 0x57, 0x65, 0xd3, 0x91, 0xe4, 0x39, 0x76,
	0x3d, 0xff, 0x41, 0x82, 0x37, 0x24, 0x85, 0xd7, 0x20, 0x56, 0x2a, 0xac, 0xde, 0xa9, 0x64, 0x7b,
	0x2f, 0xf0, 0x02, 0x8a, 0x0f, 0xb3, 0xa3, 0x3e, 0x7b, 0x71, 0x44, 0x82, 0x7e, 0x58, 0xb5, 0x6e,
	0xa3, 0x0d, 0x0d, 0xd8, 0x3d, 0x48, 0x7c, 0x00, 0x4c, 0xe3, 0x62, 0xb9, 0x5d, 0xaf, 0x8d, 0xef,
	0xec, 0x8f, 0xa1, 0x21, 0x97, 0x2e, 0x53, 0x32, 0x3b, 0x06, 0x78, 0x0f, 0xe3, 0x43, 0xdd, 0xa9,
	0xef, 0x2f, 0xdb, 0x10, 0x70, 0xf5, 0x3c, 0xb9, 0x73, 0x36, 0xf1, 0x09, 0x6e, 0x0e, 0x35, 0x6f,
	0x3a, 0xd3, 0xe8, 0x45, 0xb0, 0xce, 0xb8, 0x8e, 0xce, 0xb8, 0x7d, 0x95, 0x42, 0xc9, 0xd9, 0x2b,
	0x47, 0x93, 0xdf, 0x0c, 0xc4, 0x8b, 0x57, 0xce, 0x1c, 0x69, 0x65, 0xa7, 0xb4, 0xf6, 0x9b, 0x98,
	0xe1, 0x40, 0x26, 0x85, 0x77, 0x30, 0x88, 0x6f, 0x5c, 0xcf, 0x1c, 0x0b, 0x79, 0xd4, 0xe9, 0x0a,
	0x3d, 0xde, 0x92, 0xae, 0xc0, 0x64, 0xc5, 0x29, 0xd9, 0xf8, 0x5a, 0x4d, 0x65, 0x55, 0x44, 0xb2,
	0x7b, 0x35, 0x1b, 0xbd, 0x41, 0x6b, 0x9b, 0x3a, 0x6c, 0x9b, 0xc6, 0x7c, 0x2e, 0x4b, 0xc6, 0xfd,
	0xfc, 0x1f, 0x00, 0x00, 0xff, 0xff, 0x8f, 0x2a, 0xd3, 0xc5, 0xed, 0x01, 0x00, 0x00,
}
