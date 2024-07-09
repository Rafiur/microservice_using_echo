// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.2
// source: employee_to_salary_service.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	EmployeeToSalary_CreateSalary_FullMethodName = "/proto.EmployeeToSalary/CreateSalary"
	EmployeeToSalary_GetAllSalary_FullMethodName = "/proto.EmployeeToSalary/GetAllSalary"
	EmployeeToSalary_UpdateSalary_FullMethodName = "/proto.EmployeeToSalary/UpdateSalary"
	EmployeeToSalary_DeleteSalary_FullMethodName = "/proto.EmployeeToSalary/DeleteSalary"
)

// EmployeeToSalaryClient is the client API for EmployeeToSalary service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EmployeeToSalaryClient interface {
	CreateSalary(ctx context.Context, in *CreateSalaryRequest, opts ...grpc.CallOption) (*CreateSalaryResponse, error)
	GetAllSalary(ctx context.Context, in *GetAllSalaryRequest, opts ...grpc.CallOption) (*GetAllSalaryResponse, error)
	UpdateSalary(ctx context.Context, in *UpdateSalaryRequest, opts ...grpc.CallOption) (*UpdateSalaryResponse, error)
	DeleteSalary(ctx context.Context, in *DeleteSalaryRequest, opts ...grpc.CallOption) (*DeleteSalaryResposne, error)
}

type employeeToSalaryClient struct {
	cc grpc.ClientConnInterface
}

func NewEmployeeToSalaryClient(cc grpc.ClientConnInterface) EmployeeToSalaryClient {
	return &employeeToSalaryClient{cc}
}

func (c *employeeToSalaryClient) CreateSalary(ctx context.Context, in *CreateSalaryRequest, opts ...grpc.CallOption) (*CreateSalaryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateSalaryResponse)
	err := c.cc.Invoke(ctx, EmployeeToSalary_CreateSalary_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *employeeToSalaryClient) GetAllSalary(ctx context.Context, in *GetAllSalaryRequest, opts ...grpc.CallOption) (*GetAllSalaryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAllSalaryResponse)
	err := c.cc.Invoke(ctx, EmployeeToSalary_GetAllSalary_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *employeeToSalaryClient) UpdateSalary(ctx context.Context, in *UpdateSalaryRequest, opts ...grpc.CallOption) (*UpdateSalaryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateSalaryResponse)
	err := c.cc.Invoke(ctx, EmployeeToSalary_UpdateSalary_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *employeeToSalaryClient) DeleteSalary(ctx context.Context, in *DeleteSalaryRequest, opts ...grpc.CallOption) (*DeleteSalaryResposne, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteSalaryResposne)
	err := c.cc.Invoke(ctx, EmployeeToSalary_DeleteSalary_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EmployeeToSalaryServer is the server API for EmployeeToSalary service.
// All implementations must embed UnimplementedEmployeeToSalaryServer
// for forward compatibility
type EmployeeToSalaryServer interface {
	CreateSalary(context.Context, *CreateSalaryRequest) (*CreateSalaryResponse, error)
	GetAllSalary(context.Context, *GetAllSalaryRequest) (*GetAllSalaryResponse, error)
	UpdateSalary(context.Context, *UpdateSalaryRequest) (*UpdateSalaryResponse, error)
	DeleteSalary(context.Context, *DeleteSalaryRequest) (*DeleteSalaryResposne, error)
	mustEmbedUnimplementedEmployeeToSalaryServer()
}

// UnimplementedEmployeeToSalaryServer must be embedded to have forward compatible implementations.
type UnimplementedEmployeeToSalaryServer struct {
}

func (UnimplementedEmployeeToSalaryServer) CreateSalary(context.Context, *CreateSalaryRequest) (*CreateSalaryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSalary not implemented")
}
func (UnimplementedEmployeeToSalaryServer) GetAllSalary(context.Context, *GetAllSalaryRequest) (*GetAllSalaryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllSalary not implemented")
}
func (UnimplementedEmployeeToSalaryServer) UpdateSalary(context.Context, *UpdateSalaryRequest) (*UpdateSalaryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSalary not implemented")
}
func (UnimplementedEmployeeToSalaryServer) DeleteSalary(context.Context, *DeleteSalaryRequest) (*DeleteSalaryResposne, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSalary not implemented")
}
func (UnimplementedEmployeeToSalaryServer) mustEmbedUnimplementedEmployeeToSalaryServer() {}

// UnsafeEmployeeToSalaryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EmployeeToSalaryServer will
// result in compilation errors.
type UnsafeEmployeeToSalaryServer interface {
	mustEmbedUnimplementedEmployeeToSalaryServer()
}

func RegisterEmployeeToSalaryServer(s grpc.ServiceRegistrar, srv EmployeeToSalaryServer) {
	s.RegisterService(&EmployeeToSalary_ServiceDesc, srv)
}

func _EmployeeToSalary_CreateSalary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSalaryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmployeeToSalaryServer).CreateSalary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EmployeeToSalary_CreateSalary_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmployeeToSalaryServer).CreateSalary(ctx, req.(*CreateSalaryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmployeeToSalary_GetAllSalary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllSalaryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmployeeToSalaryServer).GetAllSalary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EmployeeToSalary_GetAllSalary_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmployeeToSalaryServer).GetAllSalary(ctx, req.(*GetAllSalaryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmployeeToSalary_UpdateSalary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSalaryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmployeeToSalaryServer).UpdateSalary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EmployeeToSalary_UpdateSalary_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmployeeToSalaryServer).UpdateSalary(ctx, req.(*UpdateSalaryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmployeeToSalary_DeleteSalary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSalaryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmployeeToSalaryServer).DeleteSalary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EmployeeToSalary_DeleteSalary_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmployeeToSalaryServer).DeleteSalary(ctx, req.(*DeleteSalaryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// EmployeeToSalary_ServiceDesc is the grpc.ServiceDesc for EmployeeToSalary service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EmployeeToSalary_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.EmployeeToSalary",
	HandlerType: (*EmployeeToSalaryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSalary",
			Handler:    _EmployeeToSalary_CreateSalary_Handler,
		},
		{
			MethodName: "GetAllSalary",
			Handler:    _EmployeeToSalary_GetAllSalary_Handler,
		},
		{
			MethodName: "UpdateSalary",
			Handler:    _EmployeeToSalary_UpdateSalary_Handler,
		},
		{
			MethodName: "DeleteSalary",
			Handler:    _EmployeeToSalary_DeleteSalary_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "employee_to_salary_service.proto",
}
