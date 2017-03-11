// Code generated by protoc-gen-go.
// source: grbac.proto
// DO NOT EDIT!

package igrpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type UserRole struct {
	UserRoleId int32 `protobuf:"varint,1,opt,name=UserRoleId" json:"UserRoleId,omitempty"`
	UserId     int32 `protobuf:"varint,2,opt,name=UserId" json:"UserId,omitempty"`
	RoleId     int32 `protobuf:"varint,3,opt,name=RoleId" json:"RoleId,omitempty"`
	RegionId   int32 `protobuf:"varint,4,opt,name=RegionId" json:"RegionId,omitempty"`
	Status     int32 `protobuf:"varint,5,opt,name=Status" json:"Status,omitempty"`
}

func (m *UserRole) Reset()                    { *m = UserRole{} }
func (m *UserRole) String() string            { return proto.CompactTextString(m) }
func (*UserRole) ProtoMessage()               {}
func (*UserRole) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *UserRole) GetUserRoleId() int32 {
	if m != nil {
		return m.UserRoleId
	}
	return 0
}

func (m *UserRole) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *UserRole) GetRoleId() int32 {
	if m != nil {
		return m.RoleId
	}
	return 0
}

func (m *UserRole) GetRegionId() int32 {
	if m != nil {
		return m.RegionId
	}
	return 0
}

func (m *UserRole) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

type Role struct {
	RoleId   int32  `protobuf:"varint,1,opt,name=RoleId" json:"RoleId,omitempty"`
	RegionId int32  `protobuf:"varint,2,opt,name=RegionId" json:"RegionId,omitempty"`
	Code     string `protobuf:"bytes,3,opt,name=Code" json:"Code,omitempty"`
	Name     string `protobuf:"bytes,4,opt,name=Name" json:"Name,omitempty"`
	Status   int32  `protobuf:"varint,5,opt,name=Status" json:"Status,omitempty"`
}

func (m *Role) Reset()                    { *m = Role{} }
func (m *Role) String() string            { return proto.CompactTextString(m) }
func (*Role) ProtoMessage()               {}
func (*Role) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{1} }

func (m *Role) GetRoleId() int32 {
	if m != nil {
		return m.RoleId
	}
	return 0
}

func (m *Role) GetRegionId() int32 {
	if m != nil {
		return m.RegionId
	}
	return 0
}

func (m *Role) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *Role) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Role) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func init() {
	proto.RegisterType((*UserRole)(nil), "igrpc.UserRole")
	proto.RegisterType((*Role)(nil), "igrpc.Role")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for RbacGrpc service

type RbacGrpcClient interface {
	AddUserRole(ctx context.Context, in *UserRole, opts ...grpc.CallOption) (*UserRole, error)
	DelUserRole(ctx context.Context, in *UserRole, opts ...grpc.CallOption) (*UserRole, error)
	GetRoleByRoleCode(ctx context.Context, in *String, opts ...grpc.CallOption) (*Role, error)
}

type rbacGrpcClient struct {
	cc *grpc.ClientConn
}

func NewRbacGrpcClient(cc *grpc.ClientConn) RbacGrpcClient {
	return &rbacGrpcClient{cc}
}

func (c *rbacGrpcClient) AddUserRole(ctx context.Context, in *UserRole, opts ...grpc.CallOption) (*UserRole, error) {
	out := new(UserRole)
	err := grpc.Invoke(ctx, "/igrpc.RbacGrpc/AddUserRole", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rbacGrpcClient) DelUserRole(ctx context.Context, in *UserRole, opts ...grpc.CallOption) (*UserRole, error) {
	out := new(UserRole)
	err := grpc.Invoke(ctx, "/igrpc.RbacGrpc/DelUserRole", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rbacGrpcClient) GetRoleByRoleCode(ctx context.Context, in *String, opts ...grpc.CallOption) (*Role, error) {
	out := new(Role)
	err := grpc.Invoke(ctx, "/igrpc.RbacGrpc/GetRoleByRoleCode", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RbacGrpc service

type RbacGrpcServer interface {
	AddUserRole(context.Context, *UserRole) (*UserRole, error)
	DelUserRole(context.Context, *UserRole) (*UserRole, error)
	GetRoleByRoleCode(context.Context, *String) (*Role, error)
}

func RegisterRbacGrpcServer(s *grpc.Server, srv RbacGrpcServer) {
	s.RegisterService(&_RbacGrpc_serviceDesc, srv)
}

func _RbacGrpc_AddUserRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRole)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RbacGrpcServer).AddUserRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/igrpc.RbacGrpc/AddUserRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RbacGrpcServer).AddUserRole(ctx, req.(*UserRole))
	}
	return interceptor(ctx, in, info, handler)
}

func _RbacGrpc_DelUserRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRole)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RbacGrpcServer).DelUserRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/igrpc.RbacGrpc/DelUserRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RbacGrpcServer).DelUserRole(ctx, req.(*UserRole))
	}
	return interceptor(ctx, in, info, handler)
}

func _RbacGrpc_GetRoleByRoleCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(String)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RbacGrpcServer).GetRoleByRoleCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/igrpc.RbacGrpc/GetRoleByRoleCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RbacGrpcServer).GetRoleByRoleCode(ctx, req.(*String))
	}
	return interceptor(ctx, in, info, handler)
}

var _RbacGrpc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "igrpc.RbacGrpc",
	HandlerType: (*RbacGrpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddUserRole",
			Handler:    _RbacGrpc_AddUserRole_Handler,
		},
		{
			MethodName: "DelUserRole",
			Handler:    _RbacGrpc_DelUserRole_Handler,
		},
		{
			MethodName: "GetRoleByRoleCode",
			Handler:    _RbacGrpc_GetRoleByRoleCode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grbac.proto",
}

func init() { proto.RegisterFile("grbac.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 250 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x4e, 0x2f, 0x4a, 0x4a,
	0x4c, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xcd, 0x4c, 0x2f, 0x2a, 0x48, 0x96, 0xe2,
	0x49, 0xce, 0xcf, 0xcd, 0xcd, 0xcf, 0x83, 0x08, 0x2a, 0xf5, 0x31, 0x72, 0x71, 0x84, 0x16, 0xa7,
	0x16, 0x05, 0xe5, 0xe7, 0xa4, 0x0a, 0xc9, 0x71, 0x71, 0xc1, 0xd8, 0x9e, 0x29, 0x12, 0x8c, 0x0a,
	0x8c, 0x1a, 0xac, 0x41, 0x48, 0x22, 0x42, 0x62, 0x5c, 0x6c, 0x20, 0x9e, 0x67, 0x8a, 0x04, 0x13,
	0x58, 0x0e, 0xca, 0x03, 0x89, 0x43, 0xf5, 0x30, 0x43, 0xc4, 0xa1, 0xea, 0xa5, 0xb8, 0x38, 0x82,
	0x52, 0xd3, 0x33, 0xf3, 0xf3, 0x3c, 0x53, 0x24, 0x58, 0xc0, 0x32, 0x70, 0x3e, 0x48, 0x4f, 0x70,
	0x49, 0x62, 0x49, 0x69, 0xb1, 0x04, 0x2b, 0x44, 0x0f, 0x84, 0xa7, 0x54, 0xc5, 0xc5, 0x02, 0x76,
	0x0b, 0xc2, 0x4c, 0x46, 0x9c, 0x66, 0x32, 0xa1, 0x99, 0x29, 0xc4, 0xc5, 0xe2, 0x9c, 0x9f, 0x92,
	0x0a, 0x76, 0x05, 0x67, 0x10, 0x98, 0x0d, 0x12, 0xf3, 0x4b, 0xcc, 0x4d, 0x05, 0xdb, 0xcf, 0x19,
	0x04, 0x66, 0xe3, 0xb2, 0xdb, 0x68, 0x2e, 0x23, 0x17, 0x47, 0x50, 0x52, 0x62, 0xb2, 0x7b, 0x51,
	0x41, 0xb2, 0x90, 0x3e, 0x17, 0xb7, 0x63, 0x4a, 0x0a, 0x3c, 0x6c, 0xf8, 0xf5, 0xc0, 0xc1, 0xa7,
	0x07, 0x13, 0x90, 0x42, 0x17, 0x00, 0x69, 0x70, 0x49, 0xcd, 0x21, 0x49, 0x83, 0xa0, 0x7b, 0x6a,
	0x09, 0x88, 0xe9, 0x54, 0x09, 0x22, 0xc1, 0xee, 0xe5, 0x85, 0xaa, 0x0a, 0x2e, 0x29, 0xca, 0xcc,
	0x4b, 0x97, 0xe2, 0x86, 0x72, 0x41, 0xf2, 0x49, 0x6c, 0xe0, 0x38, 0x33, 0x06, 0x04, 0x00, 0x00,
	0xff, 0xff, 0xad, 0xf2, 0x3c, 0x91, 0xd7, 0x01, 0x00, 0x00,
}
