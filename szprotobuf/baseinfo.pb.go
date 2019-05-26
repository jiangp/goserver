// Code generated by protoc-gen-go. DO NOT EDIT.
// source: baseinfo.proto

package szprotobuf

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

// Req userinfo
type ReqUserInfo struct {
	Userid               int64    `protobuf:"varint,1,opt,name=userid,proto3" json:"userid,omitempty"`
	Svrtoken             string   `protobuf:"bytes,2,opt,name=svrtoken,proto3" json:"svrtoken,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqUserInfo) Reset()         { *m = ReqUserInfo{} }
func (m *ReqUserInfo) String() string { return proto.CompactTextString(m) }
func (*ReqUserInfo) ProtoMessage()    {}
func (*ReqUserInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_baseinfo_f325304d735cc892, []int{0}
}
func (m *ReqUserInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqUserInfo.Unmarshal(m, b)
}
func (m *ReqUserInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqUserInfo.Marshal(b, m, deterministic)
}
func (dst *ReqUserInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqUserInfo.Merge(dst, src)
}
func (m *ReqUserInfo) XXX_Size() int {
	return xxx_messageInfo_ReqUserInfo.Size(m)
}
func (m *ReqUserInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqUserInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ReqUserInfo proto.InternalMessageInfo

func (m *ReqUserInfo) GetUserid() int64 {
	if m != nil {
		return m.Userid
	}
	return 0
}

func (m *ReqUserInfo) GetSvrtoken() string {
	if m != nil {
		return m.Svrtoken
	}
	return ""
}

type ResUserInfo struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResUserInfo) Reset()         { *m = ResUserInfo{} }
func (m *ResUserInfo) String() string { return proto.CompactTextString(m) }
func (*ResUserInfo) ProtoMessage()    {}
func (*ResUserInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_baseinfo_f325304d735cc892, []int{1}
}
func (m *ResUserInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResUserInfo.Unmarshal(m, b)
}
func (m *ResUserInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResUserInfo.Marshal(b, m, deterministic)
}
func (dst *ResUserInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResUserInfo.Merge(dst, src)
}
func (m *ResUserInfo) XXX_Size() int {
	return xxx_messageInfo_ResUserInfo.Size(m)
}
func (m *ResUserInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ResUserInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ResUserInfo proto.InternalMessageInfo

func (m *ResUserInfo) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ResUserInfo) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*ReqUserInfo)(nil), "szprotobuf.ReqUserInfo")
	proto.RegisterType((*ResUserInfo)(nil), "szprotobuf.ResUserInfo")
}

func init() { proto.RegisterFile("baseinfo.proto", fileDescriptor_baseinfo_f325304d735cc892) }

var fileDescriptor_baseinfo_f325304d735cc892 = []byte{
	// 139 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4b, 0x4a, 0x2c, 0x4e,
	0xcd, 0xcc, 0x4b, 0xcb, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x2a, 0xae, 0x02, 0x33,
	0x92, 0x4a, 0xd3, 0x94, 0x1c, 0xb9, 0xb8, 0x83, 0x52, 0x0b, 0x43, 0x8b, 0x53, 0x8b, 0x3c, 0xf3,
	0xd2, 0xf2, 0x85, 0xc4, 0xb8, 0xd8, 0x4a, 0x8b, 0x53, 0x8b, 0x32, 0x53, 0x24, 0x18, 0x15, 0x18,
	0x35, 0x98, 0x83, 0xa0, 0x3c, 0x21, 0x29, 0x2e, 0x8e, 0xe2, 0xb2, 0xa2, 0x92, 0xfc, 0xec, 0xd4,
	0x3c, 0x09, 0x26, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x38, 0x5f, 0xc9, 0x18, 0x64, 0x44, 0x31, 0xdc,
	0x08, 0x21, 0x2e, 0x96, 0xe4, 0xfc, 0x94, 0x54, 0xb0, 0x01, 0xac, 0x41, 0x60, 0xb6, 0x90, 0x00,
	0x17, 0x73, 0x6e, 0x71, 0x3a, 0x54, 0x27, 0x88, 0x99, 0xc4, 0x06, 0x76, 0x81, 0x31, 0x20, 0x00,
	0x00, 0xff, 0xff, 0xdf, 0x0a, 0x54, 0x9b, 0x9c, 0x00, 0x00, 0x00,
}