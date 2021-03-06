// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/extensions/common/tap/v3/common.proto

package envoy_extensions_common_tap_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v31 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	v3 "github.com/envoyproxy/go-control-plane/envoy/config/tap/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type CommonExtensionConfig struct {
	// Types that are valid to be assigned to ConfigType:
	//	*CommonExtensionConfig_AdminConfig
	//	*CommonExtensionConfig_StaticConfig
	//	*CommonExtensionConfig_TapdsConfig
	ConfigType           isCommonExtensionConfig_ConfigType `protobuf_oneof:"config_type"`
	XXX_NoUnkeyedLiteral struct{}                           `json:"-"`
	XXX_unrecognized     []byte                             `json:"-"`
	XXX_sizecache        int32                              `json:"-"`
}

func (m *CommonExtensionConfig) Reset()         { *m = CommonExtensionConfig{} }
func (m *CommonExtensionConfig) String() string { return proto.CompactTextString(m) }
func (*CommonExtensionConfig) ProtoMessage()    {}
func (*CommonExtensionConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_d29738d4c8661d41, []int{0}
}

func (m *CommonExtensionConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommonExtensionConfig.Unmarshal(m, b)
}
func (m *CommonExtensionConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommonExtensionConfig.Marshal(b, m, deterministic)
}
func (m *CommonExtensionConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommonExtensionConfig.Merge(m, src)
}
func (m *CommonExtensionConfig) XXX_Size() int {
	return xxx_messageInfo_CommonExtensionConfig.Size(m)
}
func (m *CommonExtensionConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_CommonExtensionConfig.DiscardUnknown(m)
}

var xxx_messageInfo_CommonExtensionConfig proto.InternalMessageInfo

type isCommonExtensionConfig_ConfigType interface {
	isCommonExtensionConfig_ConfigType()
}

type CommonExtensionConfig_AdminConfig struct {
	AdminConfig *AdminConfig `protobuf:"bytes,1,opt,name=admin_config,json=adminConfig,proto3,oneof"`
}

type CommonExtensionConfig_StaticConfig struct {
	StaticConfig *v3.TapConfig `protobuf:"bytes,2,opt,name=static_config,json=staticConfig,proto3,oneof"`
}

type CommonExtensionConfig_TapdsConfig struct {
	TapdsConfig *CommonExtensionConfig_TapDSConfig `protobuf:"bytes,3,opt,name=tapds_config,json=tapdsConfig,proto3,oneof"`
}

func (*CommonExtensionConfig_AdminConfig) isCommonExtensionConfig_ConfigType() {}

func (*CommonExtensionConfig_StaticConfig) isCommonExtensionConfig_ConfigType() {}

func (*CommonExtensionConfig_TapdsConfig) isCommonExtensionConfig_ConfigType() {}

func (m *CommonExtensionConfig) GetConfigType() isCommonExtensionConfig_ConfigType {
	if m != nil {
		return m.ConfigType
	}
	return nil
}

func (m *CommonExtensionConfig) GetAdminConfig() *AdminConfig {
	if x, ok := m.GetConfigType().(*CommonExtensionConfig_AdminConfig); ok {
		return x.AdminConfig
	}
	return nil
}

func (m *CommonExtensionConfig) GetStaticConfig() *v3.TapConfig {
	if x, ok := m.GetConfigType().(*CommonExtensionConfig_StaticConfig); ok {
		return x.StaticConfig
	}
	return nil
}

func (m *CommonExtensionConfig) GetTapdsConfig() *CommonExtensionConfig_TapDSConfig {
	if x, ok := m.GetConfigType().(*CommonExtensionConfig_TapdsConfig); ok {
		return x.TapdsConfig
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*CommonExtensionConfig) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*CommonExtensionConfig_AdminConfig)(nil),
		(*CommonExtensionConfig_StaticConfig)(nil),
		(*CommonExtensionConfig_TapdsConfig)(nil),
	}
}

type CommonExtensionConfig_TapDSConfig struct {
	ConfigSource         *v31.ConfigSource `protobuf:"bytes,1,opt,name=config_source,json=configSource,proto3" json:"config_source,omitempty"`
	Name                 string            `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *CommonExtensionConfig_TapDSConfig) Reset()         { *m = CommonExtensionConfig_TapDSConfig{} }
func (m *CommonExtensionConfig_TapDSConfig) String() string { return proto.CompactTextString(m) }
func (*CommonExtensionConfig_TapDSConfig) ProtoMessage()    {}
func (*CommonExtensionConfig_TapDSConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_d29738d4c8661d41, []int{0, 0}
}

func (m *CommonExtensionConfig_TapDSConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommonExtensionConfig_TapDSConfig.Unmarshal(m, b)
}
func (m *CommonExtensionConfig_TapDSConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommonExtensionConfig_TapDSConfig.Marshal(b, m, deterministic)
}
func (m *CommonExtensionConfig_TapDSConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommonExtensionConfig_TapDSConfig.Merge(m, src)
}
func (m *CommonExtensionConfig_TapDSConfig) XXX_Size() int {
	return xxx_messageInfo_CommonExtensionConfig_TapDSConfig.Size(m)
}
func (m *CommonExtensionConfig_TapDSConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_CommonExtensionConfig_TapDSConfig.DiscardUnknown(m)
}

var xxx_messageInfo_CommonExtensionConfig_TapDSConfig proto.InternalMessageInfo

func (m *CommonExtensionConfig_TapDSConfig) GetConfigSource() *v31.ConfigSource {
	if m != nil {
		return m.ConfigSource
	}
	return nil
}

func (m *CommonExtensionConfig_TapDSConfig) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type AdminConfig struct {
	ConfigId             string   `protobuf:"bytes,1,opt,name=config_id,json=configId,proto3" json:"config_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AdminConfig) Reset()         { *m = AdminConfig{} }
func (m *AdminConfig) String() string { return proto.CompactTextString(m) }
func (*AdminConfig) ProtoMessage()    {}
func (*AdminConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_d29738d4c8661d41, []int{1}
}

func (m *AdminConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdminConfig.Unmarshal(m, b)
}
func (m *AdminConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdminConfig.Marshal(b, m, deterministic)
}
func (m *AdminConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdminConfig.Merge(m, src)
}
func (m *AdminConfig) XXX_Size() int {
	return xxx_messageInfo_AdminConfig.Size(m)
}
func (m *AdminConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_AdminConfig.DiscardUnknown(m)
}

var xxx_messageInfo_AdminConfig proto.InternalMessageInfo

func (m *AdminConfig) GetConfigId() string {
	if m != nil {
		return m.ConfigId
	}
	return ""
}

func init() {
	proto.RegisterType((*CommonExtensionConfig)(nil), "envoy.extensions.common.tap.v3.CommonExtensionConfig")
	proto.RegisterType((*CommonExtensionConfig_TapDSConfig)(nil), "envoy.extensions.common.tap.v3.CommonExtensionConfig.TapDSConfig")
	proto.RegisterType((*AdminConfig)(nil), "envoy.extensions.common.tap.v3.AdminConfig")
}

func init() {
	proto.RegisterFile("envoy/extensions/common/tap/v3/common.proto", fileDescriptor_d29738d4c8661d41)
}

var fileDescriptor_d29738d4c8661d41 = []byte{
	// 433 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0xdf, 0xea, 0x12, 0x41,
	0x14, 0xc7, 0x1b, 0xb5, 0xd2, 0xd9, 0x15, 0x62, 0x20, 0x8a, 0x0d, 0xc4, 0xa4, 0x0b, 0xc1, 0xda,
	0x05, 0x25, 0x02, 0xa9, 0x0b, 0xc7, 0x04, 0xbb, 0x33, 0xed, 0x5e, 0xa6, 0xdd, 0xd1, 0x16, 0xdc,
	0x99, 0x61, 0x77, 0x5c, 0xf4, 0x0d, 0xa2, 0x47, 0xe8, 0x39, 0x7a, 0x85, 0xe8, 0x95, 0xc2, 0xab,
	0x98, 0x3f, 0xdb, 0xae, 0x3f, 0x44, 0xe1, 0x77, 0xb7, 0x67, 0xce, 0x99, 0xcf, 0xf9, 0x7e, 0xcf,
	0x99, 0x85, 0x03, 0xca, 0x72, 0x7e, 0x0c, 0xe8, 0x41, 0x52, 0x96, 0xc5, 0x9c, 0x65, 0x41, 0xc8,
	0x93, 0x84, 0xb3, 0x40, 0x12, 0x11, 0xe4, 0x23, 0x1b, 0xf9, 0x22, 0xe5, 0x92, 0xa3, 0x8e, 0x2e,
	0xf6, 0xcb, 0x62, 0xdf, 0xa6, 0x25, 0x11, 0x7e, 0x3e, 0xf2, 0xfa, 0x06, 0x16, 0x72, 0xb6, 0x89,
	0xb7, 0x41, 0xc8, 0x53, 0x6a, 0x10, 0x2a, 0x5c, 0x67, 0x7c, 0x9f, 0x86, 0xd4, 0x90, 0xbc, 0xee,
	0x59, 0xe5, 0x85, 0x5e, 0xde, 0xcb, 0x7d, 0x24, 0x48, 0x40, 0x18, 0xe3, 0x92, 0x48, 0x2d, 0x2c,
	0xa7, 0xa9, 0x6a, 0x1a, 0xb3, 0xad, 0x2d, 0x79, 0x96, 0x93, 0x5d, 0x1c, 0x11, 0x49, 0x83, 0xe2,
	0xc3, 0x24, 0x7a, 0xbf, 0x1a, 0xf0, 0xe9, 0x54, 0xc3, 0x66, 0x85, 0xd4, 0xa9, 0xee, 0x84, 0x16,
	0xd0, 0x25, 0x51, 0x12, 0xb3, 0xb5, 0xe9, 0xfc, 0x1c, 0x74, 0x41, 0xdf, 0x19, 0x0e, 0xfc, 0xeb,
	0xc6, 0xfc, 0x89, 0xba, 0x63, 0x10, 0xf3, 0x07, 0x4b, 0x87, 0x94, 0x21, 0x9a, 0xc1, 0x76, 0xa6,
	0x14, 0x86, 0x05, 0xb2, 0xa6, 0x91, 0x1d, 0x8b, 0x34, 0x87, 0x05, 0xe7, 0x0b, 0x11, 0xff, 0x29,
	0xae, 0xb9, 0x66, 0x31, 0x1b, 0xe8, 0x4a, 0x22, 0xa2, 0xac, 0xa0, 0xd4, 0x35, 0x65, 0x72, 0x4b,
	0xd8, 0x45, 0x97, 0xaa, 0xcd, 0xc7, 0x55, 0x29, 0x57, 0x83, 0x4d, 0xe8, 0xfd, 0x01, 0xd0, 0xa9,
	0xa4, 0xd1, 0x67, 0xd8, 0x3e, 0xdb, 0x8f, 0x9d, 0x48, 0xef, 0x5c, 0xbe, 0x5a, 0xa5, 0x69, 0xa7,
	0xc2, 0x95, 0xae, 0xc4, 0xcd, 0x13, 0x7e, 0xf8, 0x03, 0xd4, 0x9e, 0x80, 0xa5, 0x1b, 0x56, 0xce,
	0xd1, 0x0b, 0xd8, 0x60, 0x24, 0xa1, 0x7a, 0x10, 0x2d, 0xfc, 0xf8, 0x84, 0x1b, 0x69, 0xad, 0x0b,
	0x96, 0xfa, 0x70, 0x3c, 0xff, 0xf9, 0xfb, 0x7b, 0x67, 0x0a, 0x27, 0x77, 0xf0, 0xa5, 0xa7, 0x21,
	0xd9, 0x89, 0x6f, 0xe4, 0xb6, 0xb1, 0xf1, 0x7b, 0x45, 0x7a, 0x07, 0xdf, 0xde, 0x8b, 0x84, 0x11,
	0x74, 0xac, 0x6f, 0x79, 0x14, 0x14, 0xd5, 0xff, 0x62, 0xd0, 0xdb, 0x42, 0xa7, 0xb2, 0x68, 0xf4,
	0x0a, 0xb6, 0x6c, 0x49, 0x1c, 0xe9, 0xb1, 0x54, 0xcc, 0x34, 0x4d, 0xe6, 0x53, 0x34, 0x1e, 0x2a,
	0x19, 0x6f, 0xec, 0x7f, 0x74, 0x45, 0x46, 0x85, 0x8c, 0x3f, 0xc0, 0xd7, 0x31, 0x37, 0x13, 0x16,
	0x29, 0x3f, 0x1c, 0x6f, 0x6c, 0x19, 0x3b, 0xc6, 0xc3, 0x42, 0x3d, 0xee, 0x05, 0xf8, 0xfa, 0x48,
	0xbf, 0xf2, 0xd1, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x72, 0x0e, 0x82, 0x98, 0xbc, 0x03, 0x00,
	0x00,
}
