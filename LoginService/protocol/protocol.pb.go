// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protocol.proto

// 生成的程式在 Golang 中將會屬於 `token` 套件。

package protocol

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// data ={Account,Password}
type SignInRequest struct {
	Data                 string   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignInRequest) Reset()         { *m = SignInRequest{} }
func (m *SignInRequest) String() string { return proto.CompactTextString(m) }
func (*SignInRequest) ProtoMessage()    {}
func (*SignInRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2bc2336598a3f7e0, []int{0}
}

func (m *SignInRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignInRequest.Unmarshal(m, b)
}
func (m *SignInRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignInRequest.Marshal(b, m, deterministic)
}
func (m *SignInRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignInRequest.Merge(m, src)
}
func (m *SignInRequest) XXX_Size() int {
	return xxx_messageInfo_SignInRequest.Size(m)
}
func (m *SignInRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SignInRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SignInRequest proto.InternalMessageInfo

func (m *SignInRequest) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

// data ={Token}
type SignInResponse struct {
	Data                 string   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignInResponse) Reset()         { *m = SignInResponse{} }
func (m *SignInResponse) String() string { return proto.CompactTextString(m) }
func (*SignInResponse) ProtoMessage()    {}
func (*SignInResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2bc2336598a3f7e0, []int{1}
}

func (m *SignInResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignInResponse.Unmarshal(m, b)
}
func (m *SignInResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignInResponse.Marshal(b, m, deterministic)
}
func (m *SignInResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignInResponse.Merge(m, src)
}
func (m *SignInResponse) XXX_Size() int {
	return xxx_messageInfo_SignInResponse.Size(m)
}
func (m *SignInResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SignInResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SignInResponse proto.InternalMessageInfo

func (m *SignInResponse) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

// data ={Token}
type SignOutRequest struct {
	Data                 string   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignOutRequest) Reset()         { *m = SignOutRequest{} }
func (m *SignOutRequest) String() string { return proto.CompactTextString(m) }
func (*SignOutRequest) ProtoMessage()    {}
func (*SignOutRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2bc2336598a3f7e0, []int{2}
}

func (m *SignOutRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignOutRequest.Unmarshal(m, b)
}
func (m *SignOutRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignOutRequest.Marshal(b, m, deterministic)
}
func (m *SignOutRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignOutRequest.Merge(m, src)
}
func (m *SignOutRequest) XXX_Size() int {
	return xxx_messageInfo_SignOutRequest.Size(m)
}
func (m *SignOutRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SignOutRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SignOutRequest proto.InternalMessageInfo

func (m *SignOutRequest) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type SignOutResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignOutResponse) Reset()         { *m = SignOutResponse{} }
func (m *SignOutResponse) String() string { return proto.CompactTextString(m) }
func (*SignOutResponse) ProtoMessage()    {}
func (*SignOutResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2bc2336598a3f7e0, []int{3}
}

func (m *SignOutResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignOutResponse.Unmarshal(m, b)
}
func (m *SignOutResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignOutResponse.Marshal(b, m, deterministic)
}
func (m *SignOutResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignOutResponse.Merge(m, src)
}
func (m *SignOutResponse) XXX_Size() int {
	return xxx_messageInfo_SignOutResponse.Size(m)
}
func (m *SignOutResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SignOutResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SignOutResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*SignInRequest)(nil), "protocol.SignInRequest")
	proto.RegisterType((*SignInResponse)(nil), "protocol.SignInResponse")
	proto.RegisterType((*SignOutRequest)(nil), "protocol.SignOutRequest")
	proto.RegisterType((*SignOutResponse)(nil), "protocol.SignOutResponse")
}

func init() { proto.RegisterFile("protocol.proto", fileDescriptor_2bc2336598a3f7e0) }

var fileDescriptor_2bc2336598a3f7e0 = []byte{
	// 166 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x4f, 0xce, 0xcf, 0xd1, 0x03, 0x33, 0x84, 0x38, 0x60, 0x7c, 0x25, 0x65, 0x2e, 0xde, 0xe0,
	0xcc, 0xf4, 0x3c, 0xcf, 0xbc, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x21, 0x2e, 0x96,
	0x94, 0xc4, 0x92, 0x44, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x30, 0x5b, 0x49, 0x85, 0x8b,
	0x0f, 0xa6, 0xa8, 0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x15, 0x9f, 0x2a, 0xff, 0xd2, 0x12, 0x7c, 0x66,
	0x09, 0x72, 0xf1, 0xc3, 0x55, 0x41, 0x0c, 0x33, 0x9a, 0xc4, 0xc8, 0xc5, 0xef, 0x93, 0x9f, 0x9e,
	0x99, 0xe7, 0x9c, 0x9f, 0x57, 0x52, 0x94, 0x9f, 0x93, 0x93, 0x5a, 0x24, 0x64, 0xcb, 0xc5, 0x06,
	0xb1, 0x52, 0x48, 0x5c, 0x0f, 0xee, 0x78, 0x14, 0x97, 0x4a, 0x49, 0x60, 0x4a, 0x40, 0x0c, 0x54,
	0x62, 0x10, 0x72, 0xe0, 0x62, 0x87, 0xda, 0x22, 0x84, 0xa6, 0x0c, 0xe1, 0x3c, 0x29, 0x49, 0x2c,
	0x32, 0x30, 0x13, 0x92, 0xd8, 0xc0, 0x72, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8c, 0x49,
	0x93, 0xcd, 0x3b, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// LoginControllerClient is the client API for LoginController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LoginControllerClient interface {
	// SignIn 會接收 SignInRequest 資料產生token，最終會回傳 SignInResponse
	SignIn(ctx context.Context, in *SignInRequest, opts ...grpc.CallOption) (*SignInResponse, error)
	// SignOut 會接收 SignOutRequest 資料驗證並產生新的token，最終會回傳 SignOutResponse
	SignOut(ctx context.Context, in *SignOutRequest, opts ...grpc.CallOption) (*SignOutResponse, error)
}

type loginControllerClient struct {
	cc *grpc.ClientConn
}

func NewLoginControllerClient(cc *grpc.ClientConn) LoginControllerClient {
	return &loginControllerClient{cc}
}

func (c *loginControllerClient) SignIn(ctx context.Context, in *SignInRequest, opts ...grpc.CallOption) (*SignInResponse, error) {
	out := new(SignInResponse)
	err := c.cc.Invoke(ctx, "/protocol.LoginController/SignIn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loginControllerClient) SignOut(ctx context.Context, in *SignOutRequest, opts ...grpc.CallOption) (*SignOutResponse, error) {
	out := new(SignOutResponse)
	err := c.cc.Invoke(ctx, "/protocol.LoginController/SignOut", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LoginControllerServer is the server API for LoginController service.
type LoginControllerServer interface {
	// SignIn 會接收 SignInRequest 資料產生token，最終會回傳 SignInResponse
	SignIn(context.Context, *SignInRequest) (*SignInResponse, error)
	// SignOut 會接收 SignOutRequest 資料驗證並產生新的token，最終會回傳 SignOutResponse
	SignOut(context.Context, *SignOutRequest) (*SignOutResponse, error)
}

// UnimplementedLoginControllerServer can be embedded to have forward compatible implementations.
type UnimplementedLoginControllerServer struct {
}

func (*UnimplementedLoginControllerServer) SignIn(ctx context.Context, req *SignInRequest) (*SignInResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignIn not implemented")
}
func (*UnimplementedLoginControllerServer) SignOut(ctx context.Context, req *SignOutRequest) (*SignOutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignOut not implemented")
}

func RegisterLoginControllerServer(s *grpc.Server, srv LoginControllerServer) {
	s.RegisterService(&_LoginController_serviceDesc, srv)
}

func _LoginController_SignIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignInRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginControllerServer).SignIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.LoginController/SignIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginControllerServer).SignIn(ctx, req.(*SignInRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LoginController_SignOut_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignOutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginControllerServer).SignOut(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.LoginController/SignOut",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginControllerServer).SignOut(ctx, req.(*SignOutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _LoginController_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protocol.LoginController",
	HandlerType: (*LoginControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SignIn",
			Handler:    _LoginController_SignIn_Handler,
		},
		{
			MethodName: "SignOut",
			Handler:    _LoginController_SignOut_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protocol.proto",
}
