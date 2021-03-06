// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/network/echo/v2/echo.proto

package envoy_config_filter_network_echo_v2

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
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

type Echo struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Echo) Reset()         { *m = Echo{} }
func (m *Echo) String() string { return proto.CompactTextString(m) }
func (*Echo) ProtoMessage()    {}
func (*Echo) Descriptor() ([]byte, []int) {
	return fileDescriptor_91604fcc0a62c54b, []int{0}
}

func (m *Echo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Echo.Unmarshal(m, b)
}
func (m *Echo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Echo.Marshal(b, m, deterministic)
}
func (m *Echo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Echo.Merge(m, src)
}
func (m *Echo) XXX_Size() int {
	return xxx_messageInfo_Echo.Size(m)
}
func (m *Echo) XXX_DiscardUnknown() {
	xxx_messageInfo_Echo.DiscardUnknown(m)
}

var xxx_messageInfo_Echo proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Echo)(nil), "envoy.config.filter.network.echo.v2.Echo")
}

func init() {
	proto.RegisterFile("envoy/config/filter/network/echo/v2/echo.proto", fileDescriptor_91604fcc0a62c54b)
}

var fileDescriptor_91604fcc0a62c54b = []byte{
	// 179 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x4b, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x4f, 0xce, 0xcf, 0x4b, 0xcb, 0x4c, 0xd7, 0x4f, 0xcb, 0xcc, 0x29, 0x49, 0x2d, 0xd2,
	0xcf, 0x4b, 0x2d, 0x29, 0xcf, 0x2f, 0xca, 0xd6, 0x4f, 0x4d, 0xce, 0xc8, 0xd7, 0x2f, 0x33, 0x02,
	0xd3, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0xca, 0x60, 0xf5, 0x7a, 0x10, 0xf5, 0x7a, 0x10,
	0xf5, 0x7a, 0x50, 0xf5, 0x7a, 0x60, 0x75, 0x65, 0x46, 0x52, 0x72, 0xa5, 0x29, 0x05, 0x89, 0xfa,
	0x89, 0x79, 0x79, 0xf9, 0x25, 0x89, 0x25, 0x99, 0xf9, 0x79, 0xc5, 0xfa, 0xb9, 0x99, 0xe9, 0x45,
	0x89, 0x25, 0xa9, 0x10, 0x43, 0x94, 0xd8, 0xb8, 0x58, 0x5c, 0x93, 0x33, 0xf2, 0x9d, 0x0a, 0x3e,
	0xcd, 0xf8, 0xd7, 0xcf, 0xaa, 0x25, 0xa4, 0x01, 0x31, 0x34, 0xb5, 0xa2, 0x24, 0x35, 0xaf, 0x18,
	0xa4, 0x1e, 0x6a, 0x70, 0x31, 0x9a, 0xc9, 0xc6, 0x5c, 0x86, 0x99, 0xf9, 0x10, 0x17, 0x17, 0x14,
	0xe5, 0x57, 0x54, 0xea, 0x11, 0xe1, 0x18, 0x27, 0x4e, 0x90, 0x55, 0x01, 0x20, 0x7b, 0x03, 0x18,
	0x93, 0xd8, 0xc0, 0x0e, 0x30, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xc3, 0x68, 0x16, 0x1d, 0xf7,
	0x00, 0x00, 0x00,
}
