// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/hashicorp/nomad/plugins/base/base.proto

package base

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import hclspec "github.com/hashicorp/nomad/plugins/shared/hclspec"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// PluginType enumerates the type of plugins Nomad supports
type PluginType int32

const (
	PluginType_UNKNOWN PluginType = 0
	PluginType_DRIVER  PluginType = 1
	PluginType_DEVICE  PluginType = 2
)

var PluginType_name = map[int32]string{
	0: "UNKNOWN",
	1: "DRIVER",
	2: "DEVICE",
}
var PluginType_value = map[string]int32{
	"UNKNOWN": 0,
	"DRIVER":  1,
	"DEVICE":  2,
}

func (x PluginType) String() string {
	return proto.EnumName(PluginType_name, int32(x))
}
func (PluginType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_base_0160727fc8db573d, []int{0}
}

// PluginInfoRequest is used to request the plugins basic information.
type PluginInfoRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PluginInfoRequest) Reset()         { *m = PluginInfoRequest{} }
func (m *PluginInfoRequest) String() string { return proto.CompactTextString(m) }
func (*PluginInfoRequest) ProtoMessage()    {}
func (*PluginInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_base_0160727fc8db573d, []int{0}
}
func (m *PluginInfoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PluginInfoRequest.Unmarshal(m, b)
}
func (m *PluginInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PluginInfoRequest.Marshal(b, m, deterministic)
}
func (dst *PluginInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PluginInfoRequest.Merge(dst, src)
}
func (m *PluginInfoRequest) XXX_Size() int {
	return xxx_messageInfo_PluginInfoRequest.Size(m)
}
func (m *PluginInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PluginInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PluginInfoRequest proto.InternalMessageInfo

// PluginInfoResponse returns basic information about the plugin such
// that Nomad can decide whether to load the plugin or not.
type PluginInfoResponse struct {
	// type indicates what type of plugin this is.
	Type PluginType `protobuf:"varint,1,opt,name=type,proto3,enum=hashicorp.nomad.plugins.base.PluginType" json:"type,omitempty"`
	// plugin_api_version indicates the version of the Nomad Plugin API
	// this plugin is built against.
	PluginApiVersion string `protobuf:"bytes,2,opt,name=plugin_api_version,json=pluginApiVersion,proto3" json:"plugin_api_version,omitempty"`
	// plugin_version is the semver version of this individual plugin.
	// This is divorce from Nomad’s development and versioning.
	PluginVersion string `protobuf:"bytes,3,opt,name=plugin_version,json=pluginVersion,proto3" json:"plugin_version,omitempty"`
	// name is the name of the plugin
	Name                 string   `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PluginInfoResponse) Reset()         { *m = PluginInfoResponse{} }
func (m *PluginInfoResponse) String() string { return proto.CompactTextString(m) }
func (*PluginInfoResponse) ProtoMessage()    {}
func (*PluginInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_base_0160727fc8db573d, []int{1}
}
func (m *PluginInfoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PluginInfoResponse.Unmarshal(m, b)
}
func (m *PluginInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PluginInfoResponse.Marshal(b, m, deterministic)
}
func (dst *PluginInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PluginInfoResponse.Merge(dst, src)
}
func (m *PluginInfoResponse) XXX_Size() int {
	return xxx_messageInfo_PluginInfoResponse.Size(m)
}
func (m *PluginInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PluginInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PluginInfoResponse proto.InternalMessageInfo

func (m *PluginInfoResponse) GetType() PluginType {
	if m != nil {
		return m.Type
	}
	return PluginType_UNKNOWN
}

func (m *PluginInfoResponse) GetPluginApiVersion() string {
	if m != nil {
		return m.PluginApiVersion
	}
	return ""
}

func (m *PluginInfoResponse) GetPluginVersion() string {
	if m != nil {
		return m.PluginVersion
	}
	return ""
}

func (m *PluginInfoResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// ConfigSchemaRequest is used to request the configurations schema.
type ConfigSchemaRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConfigSchemaRequest) Reset()         { *m = ConfigSchemaRequest{} }
func (m *ConfigSchemaRequest) String() string { return proto.CompactTextString(m) }
func (*ConfigSchemaRequest) ProtoMessage()    {}
func (*ConfigSchemaRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_base_0160727fc8db573d, []int{2}
}
func (m *ConfigSchemaRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigSchemaRequest.Unmarshal(m, b)
}
func (m *ConfigSchemaRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigSchemaRequest.Marshal(b, m, deterministic)
}
func (dst *ConfigSchemaRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigSchemaRequest.Merge(dst, src)
}
func (m *ConfigSchemaRequest) XXX_Size() int {
	return xxx_messageInfo_ConfigSchemaRequest.Size(m)
}
func (m *ConfigSchemaRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigSchemaRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigSchemaRequest proto.InternalMessageInfo

// ConfigSchemaResponse returns the plugins configuration schema.
type ConfigSchemaResponse struct {
	// spec is the plugins configuration schema
	Spec                 *hclspec.Spec `protobuf:"bytes,1,opt,name=spec,proto3" json:"spec,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ConfigSchemaResponse) Reset()         { *m = ConfigSchemaResponse{} }
func (m *ConfigSchemaResponse) String() string { return proto.CompactTextString(m) }
func (*ConfigSchemaResponse) ProtoMessage()    {}
func (*ConfigSchemaResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_base_0160727fc8db573d, []int{3}
}
func (m *ConfigSchemaResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigSchemaResponse.Unmarshal(m, b)
}
func (m *ConfigSchemaResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigSchemaResponse.Marshal(b, m, deterministic)
}
func (dst *ConfigSchemaResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigSchemaResponse.Merge(dst, src)
}
func (m *ConfigSchemaResponse) XXX_Size() int {
	return xxx_messageInfo_ConfigSchemaResponse.Size(m)
}
func (m *ConfigSchemaResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigSchemaResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigSchemaResponse proto.InternalMessageInfo

func (m *ConfigSchemaResponse) GetSpec() *hclspec.Spec {
	if m != nil {
		return m.Spec
	}
	return nil
}

// SetConfigRequest is used to set the configuration
type SetConfigRequest struct {
	// msgpack_config is the configuration encoded as MessagePack.
	MsgpackConfig        []byte   `protobuf:"bytes,1,opt,name=msgpack_config,json=msgpackConfig,proto3" json:"msgpack_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetConfigRequest) Reset()         { *m = SetConfigRequest{} }
func (m *SetConfigRequest) String() string { return proto.CompactTextString(m) }
func (*SetConfigRequest) ProtoMessage()    {}
func (*SetConfigRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_base_0160727fc8db573d, []int{4}
}
func (m *SetConfigRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetConfigRequest.Unmarshal(m, b)
}
func (m *SetConfigRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetConfigRequest.Marshal(b, m, deterministic)
}
func (dst *SetConfigRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetConfigRequest.Merge(dst, src)
}
func (m *SetConfigRequest) XXX_Size() int {
	return xxx_messageInfo_SetConfigRequest.Size(m)
}
func (m *SetConfigRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetConfigRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetConfigRequest proto.InternalMessageInfo

func (m *SetConfigRequest) GetMsgpackConfig() []byte {
	if m != nil {
		return m.MsgpackConfig
	}
	return nil
}

// SetConfigResponse is used to respond to setting the configuration
type SetConfigResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetConfigResponse) Reset()         { *m = SetConfigResponse{} }
func (m *SetConfigResponse) String() string { return proto.CompactTextString(m) }
func (*SetConfigResponse) ProtoMessage()    {}
func (*SetConfigResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_base_0160727fc8db573d, []int{5}
}
func (m *SetConfigResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetConfigResponse.Unmarshal(m, b)
}
func (m *SetConfigResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetConfigResponse.Marshal(b, m, deterministic)
}
func (dst *SetConfigResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetConfigResponse.Merge(dst, src)
}
func (m *SetConfigResponse) XXX_Size() int {
	return xxx_messageInfo_SetConfigResponse.Size(m)
}
func (m *SetConfigResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SetConfigResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SetConfigResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*PluginInfoRequest)(nil), "hashicorp.nomad.plugins.base.PluginInfoRequest")
	proto.RegisterType((*PluginInfoResponse)(nil), "hashicorp.nomad.plugins.base.PluginInfoResponse")
	proto.RegisterType((*ConfigSchemaRequest)(nil), "hashicorp.nomad.plugins.base.ConfigSchemaRequest")
	proto.RegisterType((*ConfigSchemaResponse)(nil), "hashicorp.nomad.plugins.base.ConfigSchemaResponse")
	proto.RegisterType((*SetConfigRequest)(nil), "hashicorp.nomad.plugins.base.SetConfigRequest")
	proto.RegisterType((*SetConfigResponse)(nil), "hashicorp.nomad.plugins.base.SetConfigResponse")
	proto.RegisterEnum("hashicorp.nomad.plugins.base.PluginType", PluginType_name, PluginType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BasePluginClient is the client API for BasePlugin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BasePluginClient interface {
	// PluginInfo describes the type and version of a plugin.
	PluginInfo(ctx context.Context, in *PluginInfoRequest, opts ...grpc.CallOption) (*PluginInfoResponse, error)
	// ConfigSchema returns the schema for parsing the plugins configuration.
	ConfigSchema(ctx context.Context, in *ConfigSchemaRequest, opts ...grpc.CallOption) (*ConfigSchemaResponse, error)
	// SetConfig is used to set the configuration.
	SetConfig(ctx context.Context, in *SetConfigRequest, opts ...grpc.CallOption) (*SetConfigResponse, error)
}

type basePluginClient struct {
	cc *grpc.ClientConn
}

func NewBasePluginClient(cc *grpc.ClientConn) BasePluginClient {
	return &basePluginClient{cc}
}

func (c *basePluginClient) PluginInfo(ctx context.Context, in *PluginInfoRequest, opts ...grpc.CallOption) (*PluginInfoResponse, error) {
	out := new(PluginInfoResponse)
	err := c.cc.Invoke(ctx, "/hashicorp.nomad.plugins.base.BasePlugin/PluginInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basePluginClient) ConfigSchema(ctx context.Context, in *ConfigSchemaRequest, opts ...grpc.CallOption) (*ConfigSchemaResponse, error) {
	out := new(ConfigSchemaResponse)
	err := c.cc.Invoke(ctx, "/hashicorp.nomad.plugins.base.BasePlugin/ConfigSchema", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basePluginClient) SetConfig(ctx context.Context, in *SetConfigRequest, opts ...grpc.CallOption) (*SetConfigResponse, error) {
	out := new(SetConfigResponse)
	err := c.cc.Invoke(ctx, "/hashicorp.nomad.plugins.base.BasePlugin/SetConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BasePluginServer is the server API for BasePlugin service.
type BasePluginServer interface {
	// PluginInfo describes the type and version of a plugin.
	PluginInfo(context.Context, *PluginInfoRequest) (*PluginInfoResponse, error)
	// ConfigSchema returns the schema for parsing the plugins configuration.
	ConfigSchema(context.Context, *ConfigSchemaRequest) (*ConfigSchemaResponse, error)
	// SetConfig is used to set the configuration.
	SetConfig(context.Context, *SetConfigRequest) (*SetConfigResponse, error)
}

func RegisterBasePluginServer(s *grpc.Server, srv BasePluginServer) {
	s.RegisterService(&_BasePlugin_serviceDesc, srv)
}

func _BasePlugin_PluginInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PluginInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BasePluginServer).PluginInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hashicorp.nomad.plugins.base.BasePlugin/PluginInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BasePluginServer).PluginInfo(ctx, req.(*PluginInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BasePlugin_ConfigSchema_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfigSchemaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BasePluginServer).ConfigSchema(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hashicorp.nomad.plugins.base.BasePlugin/ConfigSchema",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BasePluginServer).ConfigSchema(ctx, req.(*ConfigSchemaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BasePlugin_SetConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BasePluginServer).SetConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hashicorp.nomad.plugins.base.BasePlugin/SetConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BasePluginServer).SetConfig(ctx, req.(*SetConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BasePlugin_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hashicorp.nomad.plugins.base.BasePlugin",
	HandlerType: (*BasePluginServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PluginInfo",
			Handler:    _BasePlugin_PluginInfo_Handler,
		},
		{
			MethodName: "ConfigSchema",
			Handler:    _BasePlugin_ConfigSchema_Handler,
		},
		{
			MethodName: "SetConfig",
			Handler:    _BasePlugin_SetConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/hashicorp/nomad/plugins/base/base.proto",
}

func init() {
	proto.RegisterFile("github.com/hashicorp/nomad/plugins/base/base.proto", fileDescriptor_base_0160727fc8db573d)
}

var fileDescriptor_base_0160727fc8db573d = []byte{
	// 430 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0xdf, 0x8b, 0xd3, 0x40,
	0x10, 0xbe, 0xd4, 0x50, 0xb9, 0xb9, 0x1f, 0xc4, 0x3d, 0x85, 0x12, 0x7c, 0x38, 0x02, 0x42, 0x91,
	0x63, 0xe3, 0xc5, 0x27, 0xc1, 0x87, 0xb3, 0xb5, 0x0f, 0x45, 0xa8, 0x92, 0x6a, 0x15, 0x5f, 0xc2,
	0x76, 0xbb, 0x4d, 0x82, 0xcd, 0xee, 0x36, 0x9b, 0x2a, 0xfd, 0xdb, 0xfc, 0x87, 0xfc, 0x33, 0x64,
	0x77, 0x93, 0x1a, 0x45, 0x4b, 0x7c, 0x49, 0x86, 0x99, 0xef, 0xfb, 0x66, 0xe6, 0x1b, 0x16, 0xa2,
	0x34, 0xaf, 0xb2, 0xdd, 0x12, 0x53, 0x51, 0x84, 0x19, 0x51, 0x59, 0x4e, 0x45, 0x29, 0x43, 0x2e,
	0x0a, 0xb2, 0x0a, 0xe5, 0x66, 0x97, 0xe6, 0x5c, 0x85, 0x4b, 0xa2, 0x98, 0xf9, 0x60, 0x59, 0x8a,
	0x4a, 0xa0, 0xc7, 0x07, 0x20, 0x36, 0x40, 0x5c, 0x03, 0xb1, 0xc6, 0xf8, 0x77, 0x1d, 0x14, 0x55,
	0x46, 0x4a, 0xb6, 0x0a, 0x33, 0xba, 0x51, 0x92, 0x51, 0xfd, 0x4f, 0x74, 0x60, 0xf5, 0x83, 0x2b,
	0x78, 0xf0, 0xce, 0x00, 0xa7, 0x7c, 0x2d, 0x62, 0xb6, 0xdd, 0x31, 0x55, 0x05, 0xdf, 0x1d, 0x40,
	0xed, 0xac, 0x92, 0x82, 0x2b, 0x86, 0x5e, 0x82, 0x5b, 0xed, 0x25, 0x1b, 0x38, 0xd7, 0xce, 0xf0,
	0x32, 0x1a, 0xe2, 0x63, 0xa3, 0x61, 0xcb, 0x7f, 0xbf, 0x97, 0x2c, 0x36, 0x2c, 0x74, 0x03, 0xc8,
	0x02, 0x12, 0x22, 0xf3, 0xe4, 0x2b, 0x2b, 0x55, 0x2e, 0xf8, 0xa0, 0x77, 0xed, 0x0c, 0x4f, 0x63,
	0xcf, 0x56, 0x5e, 0xc9, 0x7c, 0x61, 0xf3, 0xe8, 0x09, 0x5c, 0xd6, 0xe8, 0x06, 0x79, 0xcf, 0x20,
	0x2f, 0x6c, 0xb6, 0x81, 0x21, 0x70, 0x39, 0x29, 0xd8, 0xc0, 0x35, 0x45, 0x13, 0x07, 0x8f, 0xe0,
	0x6a, 0x2c, 0xf8, 0x3a, 0x4f, 0xe7, 0x34, 0x63, 0x05, 0x69, 0x96, 0xfa, 0x04, 0x0f, 0x7f, 0x4f,
	0xd7, 0x5b, 0xdd, 0x81, 0xab, 0xfd, 0x30, 0x5b, 0x9d, 0x45, 0x37, 0xff, 0xdc, 0xca, 0xfa, 0x88,
	0x6b, 0x1f, 0xf1, 0x5c, 0x32, 0x1a, 0x1b, 0x66, 0xf0, 0x02, 0xbc, 0x39, 0xab, 0xac, 0x78, 0xdd,
	0x4d, 0xcf, 0x5f, 0xa8, 0x54, 0x12, 0xfa, 0x25, 0xa1, 0xa6, 0x60, 0xf4, 0xcf, 0xe3, 0x8b, 0x3a,
	0x6b, 0xd1, 0xda, 0xfe, 0x16, 0xd5, 0x4e, 0xf4, 0xf4, 0x16, 0xe0, 0x97, 0x7b, 0xe8, 0x0c, 0xee,
	0x7f, 0x98, 0xbd, 0x99, 0xbd, 0xfd, 0x38, 0xf3, 0x4e, 0x10, 0x40, 0xff, 0x75, 0x3c, 0x5d, 0x4c,
	0x62, 0xcf, 0x31, 0xf1, 0x64, 0x31, 0x1d, 0x4f, 0xbc, 0x5e, 0xf4, 0xa3, 0x07, 0x30, 0x22, 0x8a,
	0x59, 0x1e, 0xda, 0x36, 0x0a, 0xfa, 0x7e, 0x28, 0xec, 0x72, 0xa9, 0xd6, 0xfd, 0xfd, 0x67, 0xdd,
	0x09, 0x76, 0xe4, 0xe0, 0x04, 0x7d, 0x83, 0xf3, 0xb6, 0xbd, 0xe8, 0xf6, 0xb8, 0xc6, 0x5f, 0x2e,
	0xe4, 0x47, 0xff, 0x43, 0x39, 0x34, 0xe6, 0x70, 0x7a, 0xb0, 0x10, 0xe1, 0xe3, 0x12, 0x7f, 0x9e,
	0xc9, 0x0f, 0x3b, 0xe3, 0x9b, 0x7e, 0xa3, 0xfe, 0x67, 0x57, 0xd7, 0x96, 0x7d, 0xf3, 0x80, 0x9e,
	0xff, 0x0c, 0x00, 0x00, 0xff, 0xff, 0x66, 0xa2, 0xc0, 0xe2, 0xd6, 0x03, 0x00, 0x00,
}
