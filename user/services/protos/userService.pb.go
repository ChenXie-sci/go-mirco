// Code generated by protoc-gen-go. DO NOT EDIT.
// source: userService.proto

package protos

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

type UserRequest struct {
	// @inject_tag: json:"user_name" form:"user_name" uri:"user_name"
	UserName string `protobuf:"bytes,1,opt,name=UserName,proto3" json:"UserName,omitempty"`
	// @inject_tag: json:"password" form:"password" uri:"password"
	Password string `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
	// @inject_tag: json:"password_confirm" form:"password_confirm" uri:"password_confirm"
	PasswordConfirm      string   `protobuf:"bytes,3,opt,name=PasswordConfirm,proto3" json:"PasswordConfirm,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserRequest) Reset()         { *m = UserRequest{} }
func (m *UserRequest) String() string { return proto.CompactTextString(m) }
func (*UserRequest) ProtoMessage()    {}
func (*UserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_874b4b11a4ddc7c4, []int{0}
}

func (m *UserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserRequest.Unmarshal(m, b)
}
func (m *UserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserRequest.Marshal(b, m, deterministic)
}
func (m *UserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserRequest.Merge(m, src)
}
func (m *UserRequest) XXX_Size() int {
	return xxx_messageInfo_UserRequest.Size(m)
}
func (m *UserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserRequest proto.InternalMessageInfo

func (m *UserRequest) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *UserRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *UserRequest) GetPasswordConfirm() string {
	if m != nil {
		return m.PasswordConfirm
	}
	return ""
}

type UserDetailResponse struct {
	UserDetail           *UserModel `protobuf:"bytes,1,opt,name=UserDetail,proto3" json:"UserDetail,omitempty"`
	Code                 uint32     `protobuf:"varint,2,opt,name=Code,proto3" json:"Code,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *UserDetailResponse) Reset()         { *m = UserDetailResponse{} }
func (m *UserDetailResponse) String() string { return proto.CompactTextString(m) }
func (*UserDetailResponse) ProtoMessage()    {}
func (*UserDetailResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_874b4b11a4ddc7c4, []int{1}
}

func (m *UserDetailResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserDetailResponse.Unmarshal(m, b)
}
func (m *UserDetailResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserDetailResponse.Marshal(b, m, deterministic)
}
func (m *UserDetailResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserDetailResponse.Merge(m, src)
}
func (m *UserDetailResponse) XXX_Size() int {
	return xxx_messageInfo_UserDetailResponse.Size(m)
}
func (m *UserDetailResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserDetailResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserDetailResponse proto.InternalMessageInfo

func (m *UserDetailResponse) GetUserDetail() *UserModel {
	if m != nil {
		return m.UserDetail
	}
	return nil
}

func (m *UserDetailResponse) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func init() {
	proto.RegisterType((*UserRequest)(nil), "services.UserRequest")
	proto.RegisterType((*UserDetailResponse)(nil), "services.UserDetailResponse")
}

func init() { proto.RegisterFile("userService.proto", fileDescriptor_874b4b11a4ddc7c4) }

var fileDescriptor_874b4b11a4ddc7c4 = []byte{
	// 242 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x50, 0x4d, 0x4b, 0x03, 0x31,
	0x10, 0x65, 0x55, 0xa4, 0x3b, 0xab, 0xa8, 0x23, 0xc2, 0xb2, 0x78, 0x90, 0x9e, 0x7a, 0x8a, 0xd0,
	0x1e, 0xbd, 0x88, 0xeb, 0x51, 0x45, 0x22, 0xbd, 0x08, 0x1e, 0x56, 0x77, 0x2c, 0x81, 0x76, 0xa7,
	0x66, 0x52, 0xfd, 0x23, 0xfe, 0x60, 0x49, 0xd2, 0xa8, 0xdd, 0x63, 0x4f, 0x99, 0xf7, 0xc1, 0xcb,
	0x9b, 0x81, 0x93, 0x95, 0x90, 0x7d, 0x22, 0xfb, 0x69, 0xde, 0x48, 0x2d, 0x2d, 0x3b, 0xc6, 0x81,
	0x44, 0x28, 0xd5, 0xb1, 0x17, 0xef, 0xb9, 0xa5, 0xb9, 0x44, 0x6d, 0xc8, 0x50, 0x4c, 0x85, 0xac,
	0xa6, 0x8f, 0x15, 0x89, 0xc3, 0x0a, 0x06, 0x1e, 0x3e, 0x34, 0x0b, 0x2a, 0xb3, 0x8b, 0x6c, 0x94,
	0xeb, 0x5f, 0xec, 0xb5, 0xc7, 0x46, 0xe4, 0x8b, 0x6d, 0x5b, 0xee, 0x44, 0x2d, 0x61, 0x1c, 0xc1,
	0x51, 0x9a, 0x6b, 0xee, 0xde, 0x8d, 0x5d, 0x94, 0xbb, 0xc1, 0xd2, 0xa7, 0x87, 0x2f, 0x80, 0x3e,
	0xf1, 0x96, 0x5c, 0x63, 0xe6, 0x9a, 0x64, 0xc9, 0x9d, 0x10, 0x4e, 0x00, 0xfe, 0xd8, 0xf0, 0x73,
	0x31, 0x3e, 0x55, 0xa9, 0xb7, 0x9a, 0xa6, 0xda, 0xfa, 0x9f, 0x0d, 0x11, 0xf6, 0x6a, 0x6e, 0x29,
	0x94, 0x39, 0xd4, 0x61, 0x1e, 0x7f, 0x67, 0x71, 0xa1, 0xf5, 0x05, 0xf0, 0x1a, 0x72, 0x0f, 0xef,
	0x78, 0x66, 0x3a, 0x3c, 0xdb, 0x4c, 0x5c, 0x2f, 0x5d, 0x9d, 0x6f, 0xd2, 0xbd, 0x6a, 0x35, 0x1c,
	0x44, 0xf3, 0xcc, 0x88, 0x23, 0xbb, 0x55, 0xc8, 0x4d, 0xf1, 0x9c, 0xab, 0xcb, 0xab, 0x70, 0x72,
	0x79, 0xdd, 0x0f, 0xef, 0xe4, 0x27, 0x00, 0x00, 0xff, 0xff, 0x96, 0x3f, 0x26, 0xee, 0xab, 0x01,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserServiceClient interface {
	UserLogin(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserDetailResponse, error)
	UserRegister(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserDetailResponse, error)
}

type userServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserServiceClient(cc *grpc.ClientConn) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) UserLogin(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserDetailResponse, error) {
	out := new(UserDetailResponse)
	err := c.cc.Invoke(ctx, "/services.UserService/UserLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UserRegister(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserDetailResponse, error) {
	out := new(UserDetailResponse)
	err := c.cc.Invoke(ctx, "/services.UserService/UserRegister", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	UserLogin(context.Context, *UserRequest) (*UserDetailResponse, error)
	UserRegister(context.Context, *UserRequest) (*UserDetailResponse, error)
}

// UnimplementedUserServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (*UnimplementedUserServiceServer) UserLogin(ctx context.Context, req *UserRequest) (*UserDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserLogin not implemented")
}
func (*UnimplementedUserServiceServer) UserRegister(ctx context.Context, req *UserRequest) (*UserDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserRegister not implemented")
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_UserLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UserLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.UserService/UserLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UserLogin(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UserRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UserRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.UserService/UserRegister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UserRegister(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "services.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UserLogin",
			Handler:    _UserService_UserLogin_Handler,
		},
		{
			MethodName: "UserRegister",
			Handler:    _UserService_UserRegister_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "userService.proto",
}
