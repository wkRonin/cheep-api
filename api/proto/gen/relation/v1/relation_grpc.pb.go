// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: relation/v1/relation.proto

package followv1

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
	FollowService_Follow_FullMethodName       = "/follow.v1.FollowService/Follow"
	FollowService_CancelFollow_FullMethodName = "/follow.v1.FollowService/CancelFollow"
	FollowService_GetFollowing_FullMethodName = "/follow.v1.FollowService/GetFollowing"
	FollowService_GetFans_FullMethodName      = "/follow.v1.FollowService/GetFans"
	FollowService_FollowInfo_FullMethodName   = "/follow.v1.FollowService/FollowInfo"
)

// FollowServiceClient is the client API for FollowService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FollowServiceClient interface {
	// 增删
	Follow(ctx context.Context, in *FollowRequest, opts ...grpc.CallOption) (*FollowResponse, error)
	CancelFollow(ctx context.Context, in *CancelFollowRequest, opts ...grpc.CallOption) (*CancelFollowResponse, error)
	// 获得某个人的关注列表
	GetFollowing(ctx context.Context, in *GetFollowingRequest, opts ...grpc.CallOption) (*GetFollowingResponse, error)
	// 获得粉丝列表
	GetFans(ctx context.Context, in *GetFansRequest, opts ...grpc.CallOption) (*GetFansResponse, error)
	// 获得某个人关注另外一个人的详细信息
	FollowInfo(ctx context.Context, in *FollowInfoRequest, opts ...grpc.CallOption) (*FollowInfoResponse, error)
}

type followServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFollowServiceClient(cc grpc.ClientConnInterface) FollowServiceClient {
	return &followServiceClient{cc}
}

func (c *followServiceClient) Follow(ctx context.Context, in *FollowRequest, opts ...grpc.CallOption) (*FollowResponse, error) {
	out := new(FollowResponse)
	err := c.cc.Invoke(ctx, FollowService_Follow_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *followServiceClient) CancelFollow(ctx context.Context, in *CancelFollowRequest, opts ...grpc.CallOption) (*CancelFollowResponse, error) {
	out := new(CancelFollowResponse)
	err := c.cc.Invoke(ctx, FollowService_CancelFollow_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *followServiceClient) GetFollowing(ctx context.Context, in *GetFollowingRequest, opts ...grpc.CallOption) (*GetFollowingResponse, error) {
	out := new(GetFollowingResponse)
	err := c.cc.Invoke(ctx, FollowService_GetFollowing_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *followServiceClient) GetFans(ctx context.Context, in *GetFansRequest, opts ...grpc.CallOption) (*GetFansResponse, error) {
	out := new(GetFansResponse)
	err := c.cc.Invoke(ctx, FollowService_GetFans_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *followServiceClient) FollowInfo(ctx context.Context, in *FollowInfoRequest, opts ...grpc.CallOption) (*FollowInfoResponse, error) {
	out := new(FollowInfoResponse)
	err := c.cc.Invoke(ctx, FollowService_FollowInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FollowServiceServer is the server API for FollowService service.
// All implementations must embed UnimplementedFollowServiceServer
// for forward compatibility
type FollowServiceServer interface {
	// 增删
	Follow(context.Context, *FollowRequest) (*FollowResponse, error)
	CancelFollow(context.Context, *CancelFollowRequest) (*CancelFollowResponse, error)
	// 获得某个人的关注列表
	GetFollowing(context.Context, *GetFollowingRequest) (*GetFollowingResponse, error)
	// 获得粉丝列表
	GetFans(context.Context, *GetFansRequest) (*GetFansResponse, error)
	// 获得某个人关注另外一个人的详细信息
	FollowInfo(context.Context, *FollowInfoRequest) (*FollowInfoResponse, error)
	mustEmbedUnimplementedFollowServiceServer()
}

// UnimplementedFollowServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFollowServiceServer struct {
}

func (UnimplementedFollowServiceServer) Follow(context.Context, *FollowRequest) (*FollowResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Follow not implemented")
}
func (UnimplementedFollowServiceServer) CancelFollow(context.Context, *CancelFollowRequest) (*CancelFollowResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelFollow not implemented")
}
func (UnimplementedFollowServiceServer) GetFollowing(context.Context, *GetFollowingRequest) (*GetFollowingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFollowing not implemented")
}
func (UnimplementedFollowServiceServer) GetFans(context.Context, *GetFansRequest) (*GetFansResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFans not implemented")
}
func (UnimplementedFollowServiceServer) FollowInfo(context.Context, *FollowInfoRequest) (*FollowInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FollowInfo not implemented")
}
func (UnimplementedFollowServiceServer) mustEmbedUnimplementedFollowServiceServer() {}

// UnsafeFollowServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FollowServiceServer will
// result in compilation errors.
type UnsafeFollowServiceServer interface {
	mustEmbedUnimplementedFollowServiceServer()
}

func RegisterFollowServiceServer(s grpc.ServiceRegistrar, srv FollowServiceServer) {
	s.RegisterService(&FollowService_ServiceDesc, srv)
}

func _FollowService_Follow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FollowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FollowServiceServer).Follow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FollowService_Follow_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FollowServiceServer).Follow(ctx, req.(*FollowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FollowService_CancelFollow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelFollowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FollowServiceServer).CancelFollow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FollowService_CancelFollow_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FollowServiceServer).CancelFollow(ctx, req.(*CancelFollowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FollowService_GetFollowing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFollowingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FollowServiceServer).GetFollowing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FollowService_GetFollowing_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FollowServiceServer).GetFollowing(ctx, req.(*GetFollowingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FollowService_GetFans_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFansRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FollowServiceServer).GetFans(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FollowService_GetFans_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FollowServiceServer).GetFans(ctx, req.(*GetFansRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FollowService_FollowInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FollowInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FollowServiceServer).FollowInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FollowService_FollowInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FollowServiceServer).FollowInfo(ctx, req.(*FollowInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FollowService_ServiceDesc is the grpc.ServiceDesc for FollowService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FollowService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "follow.v1.FollowService",
	HandlerType: (*FollowServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Follow",
			Handler:    _FollowService_Follow_Handler,
		},
		{
			MethodName: "CancelFollow",
			Handler:    _FollowService_CancelFollow_Handler,
		},
		{
			MethodName: "GetFollowing",
			Handler:    _FollowService_GetFollowing_Handler,
		},
		{
			MethodName: "GetFans",
			Handler:    _FollowService_GetFans_Handler,
		},
		{
			MethodName: "FollowInfo",
			Handler:    _FollowService_FollowInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "relation/v1/relation.proto",
}
