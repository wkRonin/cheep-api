// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: intr/v1/interactive.proto

package intrv1

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
	InteractiveService_IncrReadCnt_FullMethodName            = "/intr.v1.InteractiveService/IncrReadCnt"
	InteractiveService_Like_FullMethodName                   = "/intr.v1.InteractiveService/Like"
	InteractiveService_CancelLike_FullMethodName             = "/intr.v1.InteractiveService/CancelLike"
	InteractiveService_GetInteractive_FullMethodName         = "/intr.v1.InteractiveService/GetInteractive"
	InteractiveService_GetInteractiveByIds_FullMethodName    = "/intr.v1.InteractiveService/GetInteractiveByIds"
	InteractiveService_SaveFavorites_FullMethodName          = "/intr.v1.InteractiveService/SaveFavorites"
	InteractiveService_DeleteFavorites_FullMethodName        = "/intr.v1.InteractiveService/DeleteFavorites"
	InteractiveService_GetByUidWithCollection_FullMethodName = "/intr.v1.InteractiveService/GetByUidWithCollection"
	InteractiveService_Collect_FullMethodName                = "/intr.v1.InteractiveService/Collect"
	InteractiveService_CancelCollectForFid_FullMethodName    = "/intr.v1.InteractiveService/CancelCollectForFid"
	InteractiveService_CancelCollect_FullMethodName          = "/intr.v1.InteractiveService/CancelCollect"
	InteractiveService_Move_FullMethodName                   = "/intr.v1.InteractiveService/Move"
	InteractiveService_IncrComment_FullMethodName            = "/intr.v1.InteractiveService/IncrComment"
	InteractiveService_DecrComment_FullMethodName            = "/intr.v1.InteractiveService/DecrComment"
)

// InteractiveServiceClient is the client API for InteractiveService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InteractiveServiceClient interface {
	IncrReadCnt(ctx context.Context, in *IncrReadCntRequest, opts ...grpc.CallOption) (*IncrReadCntResponse, error)
	// Like 点赞
	Like(ctx context.Context, in *LikeRequest, opts ...grpc.CallOption) (*LikeResponse, error)
	// CancelLike 取消点赞
	CancelLike(ctx context.Context, in *CancelLikeRequest, opts ...grpc.CallOption) (*CancelLikeResponse, error)
	// GetInteractive 获取单个资源的交互数据
	GetInteractive(ctx context.Context, in *GetInteractiveRequest, opts ...grpc.CallOption) (*GetInteractiveResponse, error)
	// GetInteractiveByIds 批量获取资源的交互数据
	GetInteractiveByIds(ctx context.Context, in *GetInteractiveByIdsRequest, opts ...grpc.CallOption) (*GetInteractiveByIdsResponse, error)
	// 收藏夹
	SaveFavorites(ctx context.Context, in *SaveFavoritesRequest, opts ...grpc.CallOption) (*SaveFavoritesResponse, error)
	DeleteFavorites(ctx context.Context, in *DeleteFavoritesRequest, opts ...grpc.CallOption) (*DeleteFavoritesResponse, error)
	GetByUidWithCollection(ctx context.Context, in *GetByUidWithCollectionRequest, opts ...grpc.CallOption) (*GetByUidWithCollectionResponse, error)
	// Collect 收藏
	Collect(ctx context.Context, in *CollectRequest, opts ...grpc.CallOption) (*CollectResponse, error)
	CancelCollectForFid(ctx context.Context, in *CancelCollectForFidRequest, opts ...grpc.CallOption) (*CancelCollectForFidResponse, error)
	CancelCollect(ctx context.Context, in *CancelCollectRequest, opts ...grpc.CallOption) (*CancelCollectResponse, error)
	Move(ctx context.Context, in *MoveRequest, opts ...grpc.CallOption) (*MoveResponse, error)
	// 评论计数
	IncrComment(ctx context.Context, in *IncrCommentRequest, opts ...grpc.CallOption) (*IncrCommentResponse, error)
	DecrComment(ctx context.Context, in *DecrCommentRequest, opts ...grpc.CallOption) (*DecrCommentResponse, error)
}

type interactiveServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewInteractiveServiceClient(cc grpc.ClientConnInterface) InteractiveServiceClient {
	return &interactiveServiceClient{cc}
}

func (c *interactiveServiceClient) IncrReadCnt(ctx context.Context, in *IncrReadCntRequest, opts ...grpc.CallOption) (*IncrReadCntResponse, error) {
	out := new(IncrReadCntResponse)
	err := c.cc.Invoke(ctx, InteractiveService_IncrReadCnt_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *interactiveServiceClient) Like(ctx context.Context, in *LikeRequest, opts ...grpc.CallOption) (*LikeResponse, error) {
	out := new(LikeResponse)
	err := c.cc.Invoke(ctx, InteractiveService_Like_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *interactiveServiceClient) CancelLike(ctx context.Context, in *CancelLikeRequest, opts ...grpc.CallOption) (*CancelLikeResponse, error) {
	out := new(CancelLikeResponse)
	err := c.cc.Invoke(ctx, InteractiveService_CancelLike_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *interactiveServiceClient) GetInteractive(ctx context.Context, in *GetInteractiveRequest, opts ...grpc.CallOption) (*GetInteractiveResponse, error) {
	out := new(GetInteractiveResponse)
	err := c.cc.Invoke(ctx, InteractiveService_GetInteractive_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *interactiveServiceClient) GetInteractiveByIds(ctx context.Context, in *GetInteractiveByIdsRequest, opts ...grpc.CallOption) (*GetInteractiveByIdsResponse, error) {
	out := new(GetInteractiveByIdsResponse)
	err := c.cc.Invoke(ctx, InteractiveService_GetInteractiveByIds_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *interactiveServiceClient) SaveFavorites(ctx context.Context, in *SaveFavoritesRequest, opts ...grpc.CallOption) (*SaveFavoritesResponse, error) {
	out := new(SaveFavoritesResponse)
	err := c.cc.Invoke(ctx, InteractiveService_SaveFavorites_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *interactiveServiceClient) DeleteFavorites(ctx context.Context, in *DeleteFavoritesRequest, opts ...grpc.CallOption) (*DeleteFavoritesResponse, error) {
	out := new(DeleteFavoritesResponse)
	err := c.cc.Invoke(ctx, InteractiveService_DeleteFavorites_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *interactiveServiceClient) GetByUidWithCollection(ctx context.Context, in *GetByUidWithCollectionRequest, opts ...grpc.CallOption) (*GetByUidWithCollectionResponse, error) {
	out := new(GetByUidWithCollectionResponse)
	err := c.cc.Invoke(ctx, InteractiveService_GetByUidWithCollection_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *interactiveServiceClient) Collect(ctx context.Context, in *CollectRequest, opts ...grpc.CallOption) (*CollectResponse, error) {
	out := new(CollectResponse)
	err := c.cc.Invoke(ctx, InteractiveService_Collect_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *interactiveServiceClient) CancelCollectForFid(ctx context.Context, in *CancelCollectForFidRequest, opts ...grpc.CallOption) (*CancelCollectForFidResponse, error) {
	out := new(CancelCollectForFidResponse)
	err := c.cc.Invoke(ctx, InteractiveService_CancelCollectForFid_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *interactiveServiceClient) CancelCollect(ctx context.Context, in *CancelCollectRequest, opts ...grpc.CallOption) (*CancelCollectResponse, error) {
	out := new(CancelCollectResponse)
	err := c.cc.Invoke(ctx, InteractiveService_CancelCollect_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *interactiveServiceClient) Move(ctx context.Context, in *MoveRequest, opts ...grpc.CallOption) (*MoveResponse, error) {
	out := new(MoveResponse)
	err := c.cc.Invoke(ctx, InteractiveService_Move_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *interactiveServiceClient) IncrComment(ctx context.Context, in *IncrCommentRequest, opts ...grpc.CallOption) (*IncrCommentResponse, error) {
	out := new(IncrCommentResponse)
	err := c.cc.Invoke(ctx, InteractiveService_IncrComment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *interactiveServiceClient) DecrComment(ctx context.Context, in *DecrCommentRequest, opts ...grpc.CallOption) (*DecrCommentResponse, error) {
	out := new(DecrCommentResponse)
	err := c.cc.Invoke(ctx, InteractiveService_DecrComment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InteractiveServiceServer is the server API for InteractiveService service.
// All implementations must embed UnimplementedInteractiveServiceServer
// for forward compatibility
type InteractiveServiceServer interface {
	IncrReadCnt(context.Context, *IncrReadCntRequest) (*IncrReadCntResponse, error)
	// Like 点赞
	Like(context.Context, *LikeRequest) (*LikeResponse, error)
	// CancelLike 取消点赞
	CancelLike(context.Context, *CancelLikeRequest) (*CancelLikeResponse, error)
	// GetInteractive 获取单个资源的交互数据
	GetInteractive(context.Context, *GetInteractiveRequest) (*GetInteractiveResponse, error)
	// GetInteractiveByIds 批量获取资源的交互数据
	GetInteractiveByIds(context.Context, *GetInteractiveByIdsRequest) (*GetInteractiveByIdsResponse, error)
	// 收藏夹
	SaveFavorites(context.Context, *SaveFavoritesRequest) (*SaveFavoritesResponse, error)
	DeleteFavorites(context.Context, *DeleteFavoritesRequest) (*DeleteFavoritesResponse, error)
	GetByUidWithCollection(context.Context, *GetByUidWithCollectionRequest) (*GetByUidWithCollectionResponse, error)
	// Collect 收藏
	Collect(context.Context, *CollectRequest) (*CollectResponse, error)
	CancelCollectForFid(context.Context, *CancelCollectForFidRequest) (*CancelCollectForFidResponse, error)
	CancelCollect(context.Context, *CancelCollectRequest) (*CancelCollectResponse, error)
	Move(context.Context, *MoveRequest) (*MoveResponse, error)
	// 评论计数
	IncrComment(context.Context, *IncrCommentRequest) (*IncrCommentResponse, error)
	DecrComment(context.Context, *DecrCommentRequest) (*DecrCommentResponse, error)
	mustEmbedUnimplementedInteractiveServiceServer()
}

// UnimplementedInteractiveServiceServer must be embedded to have forward compatible implementations.
type UnimplementedInteractiveServiceServer struct {
}

func (UnimplementedInteractiveServiceServer) IncrReadCnt(context.Context, *IncrReadCntRequest) (*IncrReadCntResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IncrReadCnt not implemented")
}
func (UnimplementedInteractiveServiceServer) Like(context.Context, *LikeRequest) (*LikeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Like not implemented")
}
func (UnimplementedInteractiveServiceServer) CancelLike(context.Context, *CancelLikeRequest) (*CancelLikeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelLike not implemented")
}
func (UnimplementedInteractiveServiceServer) GetInteractive(context.Context, *GetInteractiveRequest) (*GetInteractiveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInteractive not implemented")
}
func (UnimplementedInteractiveServiceServer) GetInteractiveByIds(context.Context, *GetInteractiveByIdsRequest) (*GetInteractiveByIdsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInteractiveByIds not implemented")
}
func (UnimplementedInteractiveServiceServer) SaveFavorites(context.Context, *SaveFavoritesRequest) (*SaveFavoritesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveFavorites not implemented")
}
func (UnimplementedInteractiveServiceServer) DeleteFavorites(context.Context, *DeleteFavoritesRequest) (*DeleteFavoritesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFavorites not implemented")
}
func (UnimplementedInteractiveServiceServer) GetByUidWithCollection(context.Context, *GetByUidWithCollectionRequest) (*GetByUidWithCollectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByUidWithCollection not implemented")
}
func (UnimplementedInteractiveServiceServer) Collect(context.Context, *CollectRequest) (*CollectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Collect not implemented")
}
func (UnimplementedInteractiveServiceServer) CancelCollectForFid(context.Context, *CancelCollectForFidRequest) (*CancelCollectForFidResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelCollectForFid not implemented")
}
func (UnimplementedInteractiveServiceServer) CancelCollect(context.Context, *CancelCollectRequest) (*CancelCollectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelCollect not implemented")
}
func (UnimplementedInteractiveServiceServer) Move(context.Context, *MoveRequest) (*MoveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Move not implemented")
}
func (UnimplementedInteractiveServiceServer) IncrComment(context.Context, *IncrCommentRequest) (*IncrCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IncrComment not implemented")
}
func (UnimplementedInteractiveServiceServer) DecrComment(context.Context, *DecrCommentRequest) (*DecrCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DecrComment not implemented")
}
func (UnimplementedInteractiveServiceServer) mustEmbedUnimplementedInteractiveServiceServer() {}

// UnsafeInteractiveServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InteractiveServiceServer will
// result in compilation errors.
type UnsafeInteractiveServiceServer interface {
	mustEmbedUnimplementedInteractiveServiceServer()
}

func RegisterInteractiveServiceServer(s grpc.ServiceRegistrar, srv InteractiveServiceServer) {
	s.RegisterService(&InteractiveService_ServiceDesc, srv)
}

func _InteractiveService_IncrReadCnt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IncrReadCntRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InteractiveServiceServer).IncrReadCnt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InteractiveService_IncrReadCnt_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InteractiveServiceServer).IncrReadCnt(ctx, req.(*IncrReadCntRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InteractiveService_Like_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LikeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InteractiveServiceServer).Like(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InteractiveService_Like_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InteractiveServiceServer).Like(ctx, req.(*LikeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InteractiveService_CancelLike_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelLikeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InteractiveServiceServer).CancelLike(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InteractiveService_CancelLike_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InteractiveServiceServer).CancelLike(ctx, req.(*CancelLikeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InteractiveService_GetInteractive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInteractiveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InteractiveServiceServer).GetInteractive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InteractiveService_GetInteractive_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InteractiveServiceServer).GetInteractive(ctx, req.(*GetInteractiveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InteractiveService_GetInteractiveByIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInteractiveByIdsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InteractiveServiceServer).GetInteractiveByIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InteractiveService_GetInteractiveByIds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InteractiveServiceServer).GetInteractiveByIds(ctx, req.(*GetInteractiveByIdsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InteractiveService_SaveFavorites_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveFavoritesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InteractiveServiceServer).SaveFavorites(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InteractiveService_SaveFavorites_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InteractiveServiceServer).SaveFavorites(ctx, req.(*SaveFavoritesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InteractiveService_DeleteFavorites_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteFavoritesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InteractiveServiceServer).DeleteFavorites(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InteractiveService_DeleteFavorites_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InteractiveServiceServer).DeleteFavorites(ctx, req.(*DeleteFavoritesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InteractiveService_GetByUidWithCollection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByUidWithCollectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InteractiveServiceServer).GetByUidWithCollection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InteractiveService_GetByUidWithCollection_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InteractiveServiceServer).GetByUidWithCollection(ctx, req.(*GetByUidWithCollectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InteractiveService_Collect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CollectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InteractiveServiceServer).Collect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InteractiveService_Collect_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InteractiveServiceServer).Collect(ctx, req.(*CollectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InteractiveService_CancelCollectForFid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelCollectForFidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InteractiveServiceServer).CancelCollectForFid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InteractiveService_CancelCollectForFid_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InteractiveServiceServer).CancelCollectForFid(ctx, req.(*CancelCollectForFidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InteractiveService_CancelCollect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelCollectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InteractiveServiceServer).CancelCollect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InteractiveService_CancelCollect_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InteractiveServiceServer).CancelCollect(ctx, req.(*CancelCollectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InteractiveService_Move_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MoveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InteractiveServiceServer).Move(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InteractiveService_Move_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InteractiveServiceServer).Move(ctx, req.(*MoveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InteractiveService_IncrComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IncrCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InteractiveServiceServer).IncrComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InteractiveService_IncrComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InteractiveServiceServer).IncrComment(ctx, req.(*IncrCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InteractiveService_DecrComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DecrCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InteractiveServiceServer).DecrComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InteractiveService_DecrComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InteractiveServiceServer).DecrComment(ctx, req.(*DecrCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// InteractiveService_ServiceDesc is the grpc.ServiceDesc for InteractiveService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var InteractiveService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "intr.v1.InteractiveService",
	HandlerType: (*InteractiveServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IncrReadCnt",
			Handler:    _InteractiveService_IncrReadCnt_Handler,
		},
		{
			MethodName: "Like",
			Handler:    _InteractiveService_Like_Handler,
		},
		{
			MethodName: "CancelLike",
			Handler:    _InteractiveService_CancelLike_Handler,
		},
		{
			MethodName: "GetInteractive",
			Handler:    _InteractiveService_GetInteractive_Handler,
		},
		{
			MethodName: "GetInteractiveByIds",
			Handler:    _InteractiveService_GetInteractiveByIds_Handler,
		},
		{
			MethodName: "SaveFavorites",
			Handler:    _InteractiveService_SaveFavorites_Handler,
		},
		{
			MethodName: "DeleteFavorites",
			Handler:    _InteractiveService_DeleteFavorites_Handler,
		},
		{
			MethodName: "GetByUidWithCollection",
			Handler:    _InteractiveService_GetByUidWithCollection_Handler,
		},
		{
			MethodName: "Collect",
			Handler:    _InteractiveService_Collect_Handler,
		},
		{
			MethodName: "CancelCollectForFid",
			Handler:    _InteractiveService_CancelCollectForFid_Handler,
		},
		{
			MethodName: "CancelCollect",
			Handler:    _InteractiveService_CancelCollect_Handler,
		},
		{
			MethodName: "Move",
			Handler:    _InteractiveService_Move_Handler,
		},
		{
			MethodName: "IncrComment",
			Handler:    _InteractiveService_IncrComment_Handler,
		},
		{
			MethodName: "DecrComment",
			Handler:    _InteractiveService_DecrComment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "intr/v1/interactive.proto",
}
