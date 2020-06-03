// Code generated by protoc-gen-go. DO NOT EDIT.
// source: authz-provider.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type AuthzRequest struct {
	Apikey               int32    `protobuf:"varint,1,opt,name=apikey,proto3" json:"apikey,omitempty"`
	UserInfo             string   `protobuf:"bytes,3,opt,name=user_info,json=userInfo,proto3" json:"user_info,omitempty"`
	SrcIp                string   `protobuf:"bytes,4,opt,name=src_ip,json=srcIp,proto3" json:"src_ip,omitempty"`
	DstIp                string   `protobuf:"bytes,5,opt,name=dst_ip,json=dstIp,proto3" json:"dst_ip,omitempty"`
	Topics               string   `protobuf:"bytes,6,opt,name=topics,proto3" json:"topics,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthzRequest) Reset()         { *m = AuthzRequest{} }
func (m *AuthzRequest) String() string { return proto.CompactTextString(m) }
func (*AuthzRequest) ProtoMessage()    {}
func (*AuthzRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dcbfe955c405ab5c, []int{0}
}

func (m *AuthzRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthzRequest.Unmarshal(m, b)
}
func (m *AuthzRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthzRequest.Marshal(b, m, deterministic)
}
func (m *AuthzRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthzRequest.Merge(m, src)
}
func (m *AuthzRequest) XXX_Size() int {
	return xxx_messageInfo_AuthzRequest.Size(m)
}
func (m *AuthzRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthzRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AuthzRequest proto.InternalMessageInfo

func (m *AuthzRequest) GetApikey() int32 {
	if m != nil {
		return m.Apikey
	}
	return 0
}

func (m *AuthzRequest) GetUserInfo() string {
	if m != nil {
		return m.UserInfo
	}
	return ""
}

func (m *AuthzRequest) GetSrcIp() string {
	if m != nil {
		return m.SrcIp
	}
	return ""
}

func (m *AuthzRequest) GetDstIp() string {
	if m != nil {
		return m.DstIp
	}
	return ""
}

func (m *AuthzRequest) GetTopics() string {
	if m != nil {
		return m.Topics
	}
	return ""
}

type AuthzResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Status               int32    `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthzResponse) Reset()         { *m = AuthzResponse{} }
func (m *AuthzResponse) String() string { return proto.CompactTextString(m) }
func (*AuthzResponse) ProtoMessage()    {}
func (*AuthzResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dcbfe955c405ab5c, []int{1}
}

func (m *AuthzResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthzResponse.Unmarshal(m, b)
}
func (m *AuthzResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthzResponse.Marshal(b, m, deterministic)
}
func (m *AuthzResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthzResponse.Merge(m, src)
}
func (m *AuthzResponse) XXX_Size() int {
	return xxx_messageInfo_AuthzResponse.Size(m)
}
func (m *AuthzResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthzResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AuthzResponse proto.InternalMessageInfo

func (m *AuthzResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *AuthzResponse) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func init() {
	proto.RegisterType((*AuthzRequest)(nil), "proto.AuthzRequest")
	proto.RegisterType((*AuthzResponse)(nil), "proto.AuthzResponse")
}

func init() { proto.RegisterFile("authz-provider.proto", fileDescriptor_dcbfe955c405ab5c) }

var fileDescriptor_dcbfe955c405ab5c = []byte{
	// 229 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x8f, 0xbf, 0x4a, 0x03, 0x41,
	0x10, 0xc6, 0x39, 0xf5, 0xce, 0xdc, 0xa2, 0xcd, 0x1a, 0x65, 0xd1, 0x26, 0xa4, 0x4a, 0x63, 0x0a,
	0x05, 0xfb, 0x54, 0x72, 0x9d, 0xec, 0x0b, 0x84, 0x73, 0x6f, 0x82, 0x8b, 0x70, 0x33, 0xee, 0xcc,
	0x0a, 0xe6, 0x0d, 0x7c, 0x6b, 0xd9, 0x3f, 0x82, 0x56, 0xc3, 0xef, 0x37, 0x30, 0xf3, 0x7d, 0x6a,
	0x39, 0x46, 0x79, 0x3b, 0xde, 0x53, 0xc0, 0x4f, 0x3f, 0x41, 0xd8, 0x52, 0x40, 0x41, 0xdd, 0xe6,
	0xb1, 0xfe, 0x6e, 0xd4, 0xc5, 0x2e, 0xed, 0x2d, 0x7c, 0x44, 0x60, 0xd1, 0x37, 0xaa, 0x1b, 0xc9,
	0xbf, 0xc3, 0x97, 0x69, 0x56, 0xcd, 0xa6, 0xb5, 0x95, 0xf4, 0x9d, 0xea, 0x23, 0x43, 0xd8, 0xfb,
	0xf9, 0x80, 0xe6, 0x74, 0xd5, 0x6c, 0x7a, 0xbb, 0x48, 0x62, 0x98, 0x0f, 0xa8, 0xaf, 0x55, 0xc7,
	0xc1, 0xed, 0x3d, 0x99, 0xb3, 0xbc, 0x69, 0x39, 0xb8, 0x81, 0x92, 0x9e, 0x58, 0x92, 0x6e, 0x8b,
	0x9e, 0x58, 0x06, 0x4a, 0x2f, 0x04, 0xc9, 0x3b, 0x36, 0x5d, 0xd6, 0x95, 0xd6, 0x3b, 0x75, 0x59,
	0xa3, 0x30, 0xe1, 0xcc, 0xa0, 0x8d, 0x3a, 0xe7, 0xe8, 0x1c, 0x30, 0xe7, 0x30, 0x0b, 0xfb, 0x8b,
	0xe9, 0x04, 0xcb, 0x28, 0x91, 0xcd, 0x49, 0x49, 0x59, 0xe8, 0xe1, 0xb9, 0x9e, 0x78, 0xa9, 0x65,
	0xf5, 0x93, 0xea, 0x93, 0xc0, 0xe0, 0x8f, 0xa0, 0xaf, 0x4a, 0xf7, 0xed, 0xdf, 0xc2, 0xb7, 0xcb,
	0xff, 0xb2, 0xbc, 0x7e, 0xed, 0xb2, 0x7c, 0xfc, 0x09, 0x00, 0x00, 0xff, 0xff, 0xad, 0x4a, 0x95,
	0xdb, 0x3d, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AuthzProviderClient is the client API for AuthzProvider service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthzProviderClient interface {
	Authorize(ctx context.Context, in *AuthzRequest, opts ...grpc.CallOption) (*AuthzResponse, error)
}

type authzProviderClient struct {
	cc *grpc.ClientConn
}

func NewAuthzProviderClient(cc *grpc.ClientConn) AuthzProviderClient {
	return &authzProviderClient{cc}
}

func (c *authzProviderClient) Authorize(ctx context.Context, in *AuthzRequest, opts ...grpc.CallOption) (*AuthzResponse, error) {
	out := new(AuthzResponse)
	err := c.cc.Invoke(ctx, "/proto.AuthzProvider/Authorize", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthzProviderServer is the server API for AuthzProvider service.
type AuthzProviderServer interface {
	Authorize(context.Context, *AuthzRequest) (*AuthzResponse, error)
}

func RegisterAuthzProviderServer(s *grpc.Server, srv AuthzProviderServer) {
	s.RegisterService(&_AuthzProvider_serviceDesc, srv)
}

func _AuthzProvider_Authorize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthzRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthzProviderServer).Authorize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AuthzProvider/Authorize",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthzProviderServer).Authorize(ctx, req.(*AuthzRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthzProvider_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.AuthzProvider",
	HandlerType: (*AuthzProviderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Authorize",
			Handler:    _AuthzProvider_Authorize_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "authz-provider.proto",
}
