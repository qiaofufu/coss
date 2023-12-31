// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: pb/meta/meta.proto

package meta

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Meta_CreateMeta_FullMethodName = "/meta.Meta/CreateMeta"
	Meta_Exist_FullMethodName      = "/meta.Meta/Exist"
	Meta_GetMeta_FullMethodName    = "/meta.Meta/GetMeta"
)

// MetaClient is the client API for Meta service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MetaClient interface {
	CreateMeta(ctx context.Context, in *CreateMetaReq, opts ...grpc.CallOption) (*CreateMetaResp, error)
	Exist(ctx context.Context, in *ExistReq, opts ...grpc.CallOption) (*ExistResp, error)
	GetMeta(ctx context.Context, in *GetMetaReq, opts ...grpc.CallOption) (*GetMetaResp, error)
}

type metaClient struct {
	cc grpc.ClientConnInterface
}

func NewMetaClient(cc grpc.ClientConnInterface) MetaClient {
	return &metaClient{cc}
}

func (c *metaClient) CreateMeta(ctx context.Context, in *CreateMetaReq, opts ...grpc.CallOption) (*CreateMetaResp, error) {
	out := new(CreateMetaResp)
	err := c.cc.Invoke(ctx, Meta_CreateMeta_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metaClient) Exist(ctx context.Context, in *ExistReq, opts ...grpc.CallOption) (*ExistResp, error) {
	out := new(ExistResp)
	err := c.cc.Invoke(ctx, Meta_Exist_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metaClient) GetMeta(ctx context.Context, in *GetMetaReq, opts ...grpc.CallOption) (*GetMetaResp, error) {
	out := new(GetMetaResp)
	err := c.cc.Invoke(ctx, Meta_GetMeta_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MetaServer is the server API for Meta service.
// All implementations must embed UnimplementedMetaServer
// for forward compatibility
type MetaServer interface {
	CreateMeta(context.Context, *CreateMetaReq) (*CreateMetaResp, error)
	Exist(context.Context, *ExistReq) (*ExistResp, error)
	GetMeta(context.Context, *GetMetaReq) (*GetMetaResp, error)
	mustEmbedUnimplementedMetaServer()
}

// UnimplementedMetaServer must be embedded to have forward compatible implementations.
type UnimplementedMetaServer struct {
}

func (UnimplementedMetaServer) CreateMeta(context.Context, *CreateMetaReq) (*CreateMetaResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMeta not implemented")
}
func (UnimplementedMetaServer) Exist(context.Context, *ExistReq) (*ExistResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Exist not implemented")
}
func (UnimplementedMetaServer) GetMeta(context.Context, *GetMetaReq) (*GetMetaResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMeta not implemented")
}
func (UnimplementedMetaServer) mustEmbedUnimplementedMetaServer() {}

// UnsafeMetaServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MetaServer will
// result in compilation errors.
type UnsafeMetaServer interface {
	mustEmbedUnimplementedMetaServer()
}

func RegisterMetaServer(s grpc.ServiceRegistrar, srv MetaServer) {
	s.RegisterService(&Meta_ServiceDesc, srv)
}

func _Meta_CreateMeta_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMetaReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetaServer).CreateMeta(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Meta_CreateMeta_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetaServer).CreateMeta(ctx, req.(*CreateMetaReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Meta_Exist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetaServer).Exist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Meta_Exist_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetaServer).Exist(ctx, req.(*ExistReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Meta_GetMeta_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMetaReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetaServer).GetMeta(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Meta_GetMeta_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetaServer).GetMeta(ctx, req.(*GetMetaReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Meta_ServiceDesc is the grpc.ServiceDesc for Meta service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Meta_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "meta.Meta",
	HandlerType: (*MetaServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateMeta",
			Handler:    _Meta_CreateMeta_Handler,
		},
		{
			MethodName: "Exist",
			Handler:    _Meta_Exist_Handler,
		},
		{
			MethodName: "GetMeta",
			Handler:    _Meta_GetMeta_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/meta/meta.proto",
}
