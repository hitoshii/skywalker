// Code generated by protoc-gen-go. DO NOT EDIT.
// source: src/skywalker/agent/walker/package.proto

package walker

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

type AType int32

const (
	AType_IPV4       AType = 0
	AType_IPV6       AType = 1
	AType_DOMAINNAME AType = 2
)

var AType_name = map[int32]string{
	0: "IPV4",
	1: "IPV6",
	2: "DOMAINNAME",
}
var AType_value = map[string]int32{
	"IPV4":       0,
	"IPV6":       1,
	"DOMAINNAME": 2,
}

func (x AType) String() string {
	return proto.EnumName(AType_name, int32(x))
}
func (AType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_package_c7a9dd13999fd66a, []int{0}
}

type Request struct {
	Version              int32    `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	Atype                AType    `protobuf:"varint,2,opt,name=atype,proto3,enum=walker.AType" json:"atype,omitempty"`
	Addr                 string   `protobuf:"bytes,3,opt,name=addr,proto3" json:"addr,omitempty"`
	Port                 int32    `protobuf:"varint,4,opt,name=port,proto3" json:"port,omitempty"`
	Key                  []byte   `protobuf:"bytes,5,opt,name=key,proto3" json:"key,omitempty"`
	Iv                   []byte   `protobuf:"bytes,6,opt,name=iv,proto3" json:"iv,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_package_c7a9dd13999fd66a, []int{0}
}
func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (dst *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(dst, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Request) GetAtype() AType {
	if m != nil {
		return m.Atype
	}
	return AType_IPV4
}

func (m *Request) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *Request) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *Request) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *Request) GetIv() []byte {
	if m != nil {
		return m.Iv
	}
	return nil
}

type Response struct {
	Version              int32    `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	Result               string   `protobuf:"bytes,2,opt,name=result,proto3" json:"result,omitempty"`
	Key                  []byte   `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	Iv                   []byte   `protobuf:"bytes,4,opt,name=iv,proto3" json:"iv,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_package_c7a9dd13999fd66a, []int{1}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (dst *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(dst, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Response) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func (m *Response) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *Response) GetIv() []byte {
	if m != nil {
		return m.Iv
	}
	return nil
}

func init() {
	proto.RegisterType((*Request)(nil), "walker.Request")
	proto.RegisterType((*Response)(nil), "walker.Response")
	proto.RegisterEnum("walker.AType", AType_name, AType_value)
}

func init() {
	proto.RegisterFile("src/skywalker/agent/walker/package.proto", fileDescriptor_package_c7a9dd13999fd66a)
}

var fileDescriptor_package_c7a9dd13999fd66a = []byte{
	// 254 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x4f, 0x4b, 0xc3, 0x30,
	0x18, 0xc6, 0x4d, 0xff, 0x6d, 0x7b, 0xd1, 0x52, 0x72, 0x90, 0x1c, 0xcb, 0xbc, 0x04, 0x85, 0x16,
	0x54, 0xbc, 0x17, 0xf4, 0xb0, 0xc3, 0xa6, 0x04, 0xf1, 0x28, 0xc4, 0xed, 0x65, 0x94, 0x8e, 0x26,
	0x26, 0x59, 0xa5, 0x5f, 0xc3, 0x4f, 0x2c, 0x4b, 0xdb, 0x8b, 0x87, 0xdd, 0x7e, 0xcf, 0x13, 0xf2,
	0xe3, 0xe1, 0x05, 0x6e, 0xcd, 0xb6, 0xb4, 0x4d, 0xff, 0x23, 0x0f, 0x0d, 0x9a, 0x52, 0xee, 0xb1,
	0x75, 0xe5, 0x18, 0xb4, 0xdc, 0x36, 0x72, 0x8f, 0x85, 0x36, 0xca, 0x29, 0x9a, 0x0c, 0xed, 0xf2,
	0x97, 0xc0, 0x4c, 0xe0, 0xf7, 0x11, 0xad, 0xa3, 0x0c, 0x66, 0x1d, 0x1a, 0x5b, 0xab, 0x96, 0x91,
	0x9c, 0xf0, 0x58, 0x4c, 0x91, 0xde, 0x40, 0x2c, 0x5d, 0xaf, 0x91, 0x05, 0x39, 0xe1, 0xe9, 0xfd,
	0x55, 0x31, 0xfc, 0x2e, 0xaa, 0xf7, 0x5e, 0xa3, 0x18, 0xde, 0x28, 0x85, 0x48, 0xee, 0x76, 0x86,
	0x85, 0x39, 0xe1, 0x0b, 0xe1, 0xf9, 0xd4, 0x69, 0x65, 0x1c, 0x8b, 0xbc, 0xcf, 0x33, 0xcd, 0x20,
	0x6c, 0xb0, 0x67, 0x71, 0x4e, 0xf8, 0xa5, 0x38, 0x21, 0x4d, 0x21, 0xa8, 0x3b, 0x96, 0xf8, 0x22,
	0xa8, 0xbb, 0xe5, 0x27, 0xcc, 0x05, 0x5a, 0xad, 0x5a, 0x8b, 0x67, 0x46, 0x5d, 0x43, 0x62, 0xd0,
	0x1e, 0x0f, 0xce, 0xaf, 0x5a, 0x88, 0x31, 0x4d, 0xfe, 0xf0, 0xbf, 0x3f, 0x9a, 0xfc, 0xb7, 0x77,
	0x10, 0xfb, 0xe5, 0x74, 0x0e, 0xd1, 0xea, 0xed, 0xe3, 0x31, 0xbb, 0x18, 0xe9, 0x29, 0x23, 0x34,
	0x05, 0x78, 0x7e, 0x5d, 0x57, 0xab, 0xcd, 0xa6, 0x5a, 0xbf, 0x64, 0xc1, 0x57, 0xe2, 0x0f, 0xf6,
	0xf0, 0x17, 0x00, 0x00, 0xff, 0xff, 0x5c, 0xbc, 0x68, 0xf9, 0x5c, 0x01, 0x00, 0x00,
}
