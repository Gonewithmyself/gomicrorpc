// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common.proto

package proto

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

type RespType int32

const (
	RespType_NONE    RespType = 0
	RespType_ASCEND  RespType = 1
	RespType_DESCEND RespType = 2
)

var RespType_name = map[int32]string{
	0: "NONE",
	1: "ASCEND",
	2: "DESCEND",
}

var RespType_value = map[string]int32{
	"NONE":    0,
	"ASCEND":  1,
	"DESCEND": 2,
}

func (x RespType) String() string {
	return proto.EnumName(RespType_name, int32(x))
}

func (RespType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{0}
}

type SayParam struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SayParam) Reset()         { *m = SayParam{} }
func (m *SayParam) String() string { return proto.CompactTextString(m) }
func (*SayParam) ProtoMessage()    {}
func (*SayParam) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{0}
}

func (m *SayParam) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SayParam.Unmarshal(m, b)
}
func (m *SayParam) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SayParam.Marshal(b, m, deterministic)
}
func (m *SayParam) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SayParam.Merge(m, src)
}
func (m *SayParam) XXX_Size() int {
	return xxx_messageInfo_SayParam.Size(m)
}
func (m *SayParam) XXX_DiscardUnknown() {
	xxx_messageInfo_SayParam.DiscardUnknown(m)
}

var xxx_messageInfo_SayParam proto.InternalMessageInfo

func (m *SayParam) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type Pair struct {
	Key                  int32    `protobuf:"varint,1,opt,name=key,proto3" json:"key,omitempty"`
	Values               string   `protobuf:"bytes,2,opt,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pair) Reset()         { *m = Pair{} }
func (m *Pair) String() string { return proto.CompactTextString(m) }
func (*Pair) ProtoMessage()    {}
func (*Pair) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{1}
}

func (m *Pair) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pair.Unmarshal(m, b)
}
func (m *Pair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pair.Marshal(b, m, deterministic)
}
func (m *Pair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pair.Merge(m, src)
}
func (m *Pair) XXX_Size() int {
	return xxx_messageInfo_Pair.Size(m)
}
func (m *Pair) XXX_DiscardUnknown() {
	xxx_messageInfo_Pair.DiscardUnknown(m)
}

var xxx_messageInfo_Pair proto.InternalMessageInfo

func (m *Pair) GetKey() int32 {
	if m != nil {
		return m.Key
	}
	return 0
}

func (m *Pair) GetValues() string {
	if m != nil {
		return m.Values
	}
	return ""
}

type SayResponse struct {
	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	// 数组
	Values []string `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty"`
	// map
	Header               map[string]*Pair `protobuf:"bytes,3,rep,name=header,proto3" json:"header,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Type                 RespType         `protobuf:"varint,4,opt,name=type,proto3,enum=proto.RespType" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *SayResponse) Reset()         { *m = SayResponse{} }
func (m *SayResponse) String() string { return proto.CompactTextString(m) }
func (*SayResponse) ProtoMessage()    {}
func (*SayResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{2}
}

func (m *SayResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SayResponse.Unmarshal(m, b)
}
func (m *SayResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SayResponse.Marshal(b, m, deterministic)
}
func (m *SayResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SayResponse.Merge(m, src)
}
func (m *SayResponse) XXX_Size() int {
	return xxx_messageInfo_SayResponse.Size(m)
}
func (m *SayResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SayResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SayResponse proto.InternalMessageInfo

func (m *SayResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *SayResponse) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

func (m *SayResponse) GetHeader() map[string]*Pair {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *SayResponse) GetType() RespType {
	if m != nil {
		return m.Type
	}
	return RespType_NONE
}

func init() {
	proto.RegisterEnum("proto.RespType", RespType_name, RespType_value)
	proto.RegisterType((*SayParam)(nil), "proto.SayParam")
	proto.RegisterType((*Pair)(nil), "proto.Pair")
	proto.RegisterType((*SayResponse)(nil), "proto.SayResponse")
	proto.RegisterMapType((map[string]*Pair)(nil), "proto.SayResponse.HeaderEntry")
}

func init() { proto.RegisterFile("common.proto", fileDescriptor_555bd8c177793206) }

var fileDescriptor_555bd8c177793206 = []byte{
	// 283 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0xdd, 0x4a, 0x84, 0x40,
	0x14, 0xc7, 0x77, 0xfc, 0xca, 0x3d, 0x13, 0xad, 0x9c, 0x8b, 0x90, 0x25, 0xc2, 0xec, 0x46, 0x82,
	0x24, 0x8c, 0x22, 0xba, 0x8b, 0xd6, 0xd8, 0x2b, 0x5b, 0xc6, 0x5e, 0x60, 0xaa, 0xa1, 0x22, 0x75,
	0x44, 0x2d, 0x98, 0x77, 0xed, 0x61, 0xc2, 0xd1, 0x58, 0x63, 0xaf, 0x3c, 0xc7, 0xff, 0x07, 0xe7,
	0x37, 0xb0, 0xff, 0x22, 0xcb, 0x52, 0x56, 0x71, 0xdd, 0xc8, 0x4e, 0xa2, 0xad, 0x3f, 0xe1, 0x11,
	0xb8, 0x39, 0x57, 0x1b, 0xde, 0xf0, 0x12, 0x3d, 0x30, 0xcb, 0xf6, 0xcd, 0x27, 0x01, 0x89, 0xe6,
	0xac, 0x1f, 0xc3, 0x0b, 0xb0, 0x36, 0xfc, 0xa3, 0xe9, 0x95, 0x4f, 0xa1, 0xb4, 0x62, 0xb3, 0x7e,
	0xc4, 0x43, 0x70, 0xbe, 0x79, 0xf1, 0x25, 0x5a, 0xdf, 0xd0, 0xf6, 0x71, 0x0b, 0x7f, 0x08, 0xd0,
	0x9c, 0x2b, 0x26, 0xda, 0x5a, 0x56, 0xad, 0xd8, 0xed, 0xfc, 0x97, 0x34, 0xb7, 0x49, 0xbc, 0x06,
	0xe7, 0x5d, 0xf0, 0x57, 0xd1, 0xf8, 0x66, 0x60, 0x46, 0x34, 0x39, 0x1e, 0x0e, 0x8d, 0x27, 0x6d,
	0xf1, 0x5a, 0x1b, 0xd2, 0xaa, 0x6b, 0x14, 0x1b, 0xdd, 0x78, 0x0a, 0x56, 0xa7, 0x6a, 0xe1, 0x5b,
	0x01, 0x89, 0x0e, 0x92, 0xc5, 0x98, 0xea, 0x23, 0x4f, 0xaa, 0x16, 0x4c, 0x8b, 0xcb, 0x07, 0xa0,
	0x93, 0xec, 0x94, 0x67, 0x3e, 0xf0, 0x9c, 0x80, 0xad, 0xef, 0xd0, 0x38, 0x34, 0xa1, 0x63, 0x4d,
	0x4f, 0xcf, 0x06, 0xe5, 0xd6, 0xb8, 0x21, 0x67, 0xe7, 0xe0, 0xfe, 0x35, 0xa3, 0x0b, 0x56, 0xf6,
	0x98, 0xa5, 0xde, 0x0c, 0x01, 0x9c, 0xbb, 0xfc, 0x3e, 0xcd, 0x56, 0x1e, 0x41, 0x0a, 0x7b, 0xab,
	0x74, 0x58, 0x8c, 0xe4, 0x0a, 0xcc, 0x9c, 0x2b, 0x8c, 0xc1, 0x5e, 0x8b, 0xa2, 0x90, 0xb8, 0xd8,
	0x32, 0xe9, 0x27, 0x5f, 0xe2, 0x2e, 0x64, 0x38, 0x7b, 0x76, 0xf4, 0xcf, 0xcb, 0xdf, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x88, 0x6f, 0x82, 0xd2, 0xb2, 0x01, 0x00, 0x00,
}
