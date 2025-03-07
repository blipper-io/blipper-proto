// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.2
// source: blip/blip.proto

package blip

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	BlipService_CreateBlip_FullMethodName        = "/blipper.blip.v1.BlipService/CreateBlip"
	BlipService_GetBlip_FullMethodName           = "/blipper.blip.v1.BlipService/GetBlip"
	BlipService_ListBlips_FullMethodName         = "/blipper.blip.v1.BlipService/ListBlips"
	BlipService_DeleteBlip_FullMethodName        = "/blipper.blip.v1.BlipService/DeleteBlip"
	BlipService_UpdateBlip_FullMethodName        = "/blipper.blip.v1.BlipService/UpdateBlip"
	BlipService_CreateComment_FullMethodName     = "/blipper.blip.v1.BlipService/CreateComment"
	BlipService_GetComments_FullMethodName       = "/blipper.blip.v1.BlipService/GetComments"
	BlipService_UpdateComment_FullMethodName     = "/blipper.blip.v1.BlipService/UpdateComment"
	BlipService_DeleteComment_FullMethodName     = "/blipper.blip.v1.BlipService/DeleteComment"
	BlipService_Vote_FullMethodName              = "/blipper.blip.v1.BlipService/Vote"
	BlipService_GetVoteStatus_FullMethodName     = "/blipper.blip.v1.BlipService/GetVoteStatus"
	BlipService_GetCircleTimeline_FullMethodName = "/blipper.blip.v1.BlipService/GetCircleTimeline"
	BlipService_GetUserTimeline_FullMethodName   = "/blipper.blip.v1.BlipService/GetUserTimeline"
	BlipService_GetTrendingFeed_FullMethodName   = "/blipper.blip.v1.BlipService/GetTrendingFeed"
	BlipService_GetPopularFeed_FullMethodName    = "/blipper.blip.v1.BlipService/GetPopularFeed"
)

// BlipServiceClient is the client API for BlipService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BlipServiceClient interface {
	CreateBlip(ctx context.Context, in *CreateBlipRequest, opts ...grpc.CallOption) (*CreateBlipResponse, error)
	GetBlip(ctx context.Context, in *GetBlipRequest, opts ...grpc.CallOption) (*GetBlipResponse, error)
	ListBlips(ctx context.Context, in *ListBlipsRequest, opts ...grpc.CallOption) (*ListBlipsResponse, error)
	DeleteBlip(ctx context.Context, in *DeleteBlipRequest, opts ...grpc.CallOption) (*DeleteBlipResponse, error)
	UpdateBlip(ctx context.Context, in *UpdateBlipRequest, opts ...grpc.CallOption) (*UpdateBlipResponse, error)
	// Comment operations
	CreateComment(ctx context.Context, in *CreateCommentRequest, opts ...grpc.CallOption) (*CreateCommentResponse, error)
	GetComments(ctx context.Context, in *GetCommentsRequest, opts ...grpc.CallOption) (*GetCommentsResponse, error)
	UpdateComment(ctx context.Context, in *UpdateCommentRequest, opts ...grpc.CallOption) (*UpdateCommentResponse, error)
	DeleteComment(ctx context.Context, in *DeleteCommentRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Vote operations
	Vote(ctx context.Context, in *VoteRequest, opts ...grpc.CallOption) (*VoteResponse, error)
	GetVoteStatus(ctx context.Context, in *GetVoteStatusRequest, opts ...grpc.CallOption) (*GetVoteStatusResponse, error)
	// Timeline operations
	GetCircleTimeline(ctx context.Context, in *GetCircleTimelineRequest, opts ...grpc.CallOption) (*GetCircleTimelineResponse, error)
	GetUserTimeline(ctx context.Context, in *GetUserTimelineRequest, opts ...grpc.CallOption) (*GetUserTimelineResponse, error)
	// Feed operations
	GetTrendingFeed(ctx context.Context, in *GetTrendingFeedRequest, opts ...grpc.CallOption) (*GetTrendingFeedResponse, error)
	GetPopularFeed(ctx context.Context, in *GetPopularFeedRequest, opts ...grpc.CallOption) (*GetPopularFeedResponse, error)
}

type blipServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBlipServiceClient(cc grpc.ClientConnInterface) BlipServiceClient {
	return &blipServiceClient{cc}
}

func (c *blipServiceClient) CreateBlip(ctx context.Context, in *CreateBlipRequest, opts ...grpc.CallOption) (*CreateBlipResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateBlipResponse)
	err := c.cc.Invoke(ctx, BlipService_CreateBlip_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blipServiceClient) GetBlip(ctx context.Context, in *GetBlipRequest, opts ...grpc.CallOption) (*GetBlipResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetBlipResponse)
	err := c.cc.Invoke(ctx, BlipService_GetBlip_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blipServiceClient) ListBlips(ctx context.Context, in *ListBlipsRequest, opts ...grpc.CallOption) (*ListBlipsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListBlipsResponse)
	err := c.cc.Invoke(ctx, BlipService_ListBlips_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blipServiceClient) DeleteBlip(ctx context.Context, in *DeleteBlipRequest, opts ...grpc.CallOption) (*DeleteBlipResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteBlipResponse)
	err := c.cc.Invoke(ctx, BlipService_DeleteBlip_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blipServiceClient) UpdateBlip(ctx context.Context, in *UpdateBlipRequest, opts ...grpc.CallOption) (*UpdateBlipResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateBlipResponse)
	err := c.cc.Invoke(ctx, BlipService_UpdateBlip_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blipServiceClient) CreateComment(ctx context.Context, in *CreateCommentRequest, opts ...grpc.CallOption) (*CreateCommentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateCommentResponse)
	err := c.cc.Invoke(ctx, BlipService_CreateComment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blipServiceClient) GetComments(ctx context.Context, in *GetCommentsRequest, opts ...grpc.CallOption) (*GetCommentsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCommentsResponse)
	err := c.cc.Invoke(ctx, BlipService_GetComments_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blipServiceClient) UpdateComment(ctx context.Context, in *UpdateCommentRequest, opts ...grpc.CallOption) (*UpdateCommentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateCommentResponse)
	err := c.cc.Invoke(ctx, BlipService_UpdateComment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blipServiceClient) DeleteComment(ctx context.Context, in *DeleteCommentRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, BlipService_DeleteComment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blipServiceClient) Vote(ctx context.Context, in *VoteRequest, opts ...grpc.CallOption) (*VoteResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(VoteResponse)
	err := c.cc.Invoke(ctx, BlipService_Vote_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blipServiceClient) GetVoteStatus(ctx context.Context, in *GetVoteStatusRequest, opts ...grpc.CallOption) (*GetVoteStatusResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetVoteStatusResponse)
	err := c.cc.Invoke(ctx, BlipService_GetVoteStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blipServiceClient) GetCircleTimeline(ctx context.Context, in *GetCircleTimelineRequest, opts ...grpc.CallOption) (*GetCircleTimelineResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCircleTimelineResponse)
	err := c.cc.Invoke(ctx, BlipService_GetCircleTimeline_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blipServiceClient) GetUserTimeline(ctx context.Context, in *GetUserTimelineRequest, opts ...grpc.CallOption) (*GetUserTimelineResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUserTimelineResponse)
	err := c.cc.Invoke(ctx, BlipService_GetUserTimeline_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blipServiceClient) GetTrendingFeed(ctx context.Context, in *GetTrendingFeedRequest, opts ...grpc.CallOption) (*GetTrendingFeedResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetTrendingFeedResponse)
	err := c.cc.Invoke(ctx, BlipService_GetTrendingFeed_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blipServiceClient) GetPopularFeed(ctx context.Context, in *GetPopularFeedRequest, opts ...grpc.CallOption) (*GetPopularFeedResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetPopularFeedResponse)
	err := c.cc.Invoke(ctx, BlipService_GetPopularFeed_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BlipServiceServer is the server API for BlipService service.
// All implementations must embed UnimplementedBlipServiceServer
// for forward compatibility.
type BlipServiceServer interface {
	CreateBlip(context.Context, *CreateBlipRequest) (*CreateBlipResponse, error)
	GetBlip(context.Context, *GetBlipRequest) (*GetBlipResponse, error)
	ListBlips(context.Context, *ListBlipsRequest) (*ListBlipsResponse, error)
	DeleteBlip(context.Context, *DeleteBlipRequest) (*DeleteBlipResponse, error)
	UpdateBlip(context.Context, *UpdateBlipRequest) (*UpdateBlipResponse, error)
	// Comment operations
	CreateComment(context.Context, *CreateCommentRequest) (*CreateCommentResponse, error)
	GetComments(context.Context, *GetCommentsRequest) (*GetCommentsResponse, error)
	UpdateComment(context.Context, *UpdateCommentRequest) (*UpdateCommentResponse, error)
	DeleteComment(context.Context, *DeleteCommentRequest) (*emptypb.Empty, error)
	// Vote operations
	Vote(context.Context, *VoteRequest) (*VoteResponse, error)
	GetVoteStatus(context.Context, *GetVoteStatusRequest) (*GetVoteStatusResponse, error)
	// Timeline operations
	GetCircleTimeline(context.Context, *GetCircleTimelineRequest) (*GetCircleTimelineResponse, error)
	GetUserTimeline(context.Context, *GetUserTimelineRequest) (*GetUserTimelineResponse, error)
	// Feed operations
	GetTrendingFeed(context.Context, *GetTrendingFeedRequest) (*GetTrendingFeedResponse, error)
	GetPopularFeed(context.Context, *GetPopularFeedRequest) (*GetPopularFeedResponse, error)
	mustEmbedUnimplementedBlipServiceServer()
}

// UnimplementedBlipServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedBlipServiceServer struct{}

func (UnimplementedBlipServiceServer) CreateBlip(context.Context, *CreateBlipRequest) (*CreateBlipResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBlip not implemented")
}
func (UnimplementedBlipServiceServer) GetBlip(context.Context, *GetBlipRequest) (*GetBlipResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlip not implemented")
}
func (UnimplementedBlipServiceServer) ListBlips(context.Context, *ListBlipsRequest) (*ListBlipsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBlips not implemented")
}
func (UnimplementedBlipServiceServer) DeleteBlip(context.Context, *DeleteBlipRequest) (*DeleteBlipResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBlip not implemented")
}
func (UnimplementedBlipServiceServer) UpdateBlip(context.Context, *UpdateBlipRequest) (*UpdateBlipResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBlip not implemented")
}
func (UnimplementedBlipServiceServer) CreateComment(context.Context, *CreateCommentRequest) (*CreateCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateComment not implemented")
}
func (UnimplementedBlipServiceServer) GetComments(context.Context, *GetCommentsRequest) (*GetCommentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetComments not implemented")
}
func (UnimplementedBlipServiceServer) UpdateComment(context.Context, *UpdateCommentRequest) (*UpdateCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateComment not implemented")
}
func (UnimplementedBlipServiceServer) DeleteComment(context.Context, *DeleteCommentRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteComment not implemented")
}
func (UnimplementedBlipServiceServer) Vote(context.Context, *VoteRequest) (*VoteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Vote not implemented")
}
func (UnimplementedBlipServiceServer) GetVoteStatus(context.Context, *GetVoteStatusRequest) (*GetVoteStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVoteStatus not implemented")
}
func (UnimplementedBlipServiceServer) GetCircleTimeline(context.Context, *GetCircleTimelineRequest) (*GetCircleTimelineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCircleTimeline not implemented")
}
func (UnimplementedBlipServiceServer) GetUserTimeline(context.Context, *GetUserTimelineRequest) (*GetUserTimelineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserTimeline not implemented")
}
func (UnimplementedBlipServiceServer) GetTrendingFeed(context.Context, *GetTrendingFeedRequest) (*GetTrendingFeedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTrendingFeed not implemented")
}
func (UnimplementedBlipServiceServer) GetPopularFeed(context.Context, *GetPopularFeedRequest) (*GetPopularFeedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPopularFeed not implemented")
}
func (UnimplementedBlipServiceServer) mustEmbedUnimplementedBlipServiceServer() {}
func (UnimplementedBlipServiceServer) testEmbeddedByValue()                     {}

// UnsafeBlipServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BlipServiceServer will
// result in compilation errors.
type UnsafeBlipServiceServer interface {
	mustEmbedUnimplementedBlipServiceServer()
}

func RegisterBlipServiceServer(s grpc.ServiceRegistrar, srv BlipServiceServer) {
	// If the following call pancis, it indicates UnimplementedBlipServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&BlipService_ServiceDesc, srv)
}

func _BlipService_CreateBlip_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBlipRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlipServiceServer).CreateBlip(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BlipService_CreateBlip_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlipServiceServer).CreateBlip(ctx, req.(*CreateBlipRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlipService_GetBlip_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBlipRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlipServiceServer).GetBlip(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BlipService_GetBlip_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlipServiceServer).GetBlip(ctx, req.(*GetBlipRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlipService_ListBlips_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBlipsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlipServiceServer).ListBlips(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BlipService_ListBlips_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlipServiceServer).ListBlips(ctx, req.(*ListBlipsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlipService_DeleteBlip_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBlipRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlipServiceServer).DeleteBlip(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BlipService_DeleteBlip_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlipServiceServer).DeleteBlip(ctx, req.(*DeleteBlipRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlipService_UpdateBlip_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBlipRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlipServiceServer).UpdateBlip(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BlipService_UpdateBlip_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlipServiceServer).UpdateBlip(ctx, req.(*UpdateBlipRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlipService_CreateComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlipServiceServer).CreateComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BlipService_CreateComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlipServiceServer).CreateComment(ctx, req.(*CreateCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlipService_GetComments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCommentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlipServiceServer).GetComments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BlipService_GetComments_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlipServiceServer).GetComments(ctx, req.(*GetCommentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlipService_UpdateComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlipServiceServer).UpdateComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BlipService_UpdateComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlipServiceServer).UpdateComment(ctx, req.(*UpdateCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlipService_DeleteComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlipServiceServer).DeleteComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BlipService_DeleteComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlipServiceServer).DeleteComment(ctx, req.(*DeleteCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlipService_Vote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlipServiceServer).Vote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BlipService_Vote_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlipServiceServer).Vote(ctx, req.(*VoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlipService_GetVoteStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVoteStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlipServiceServer).GetVoteStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BlipService_GetVoteStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlipServiceServer).GetVoteStatus(ctx, req.(*GetVoteStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlipService_GetCircleTimeline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCircleTimelineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlipServiceServer).GetCircleTimeline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BlipService_GetCircleTimeline_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlipServiceServer).GetCircleTimeline(ctx, req.(*GetCircleTimelineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlipService_GetUserTimeline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserTimelineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlipServiceServer).GetUserTimeline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BlipService_GetUserTimeline_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlipServiceServer).GetUserTimeline(ctx, req.(*GetUserTimelineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlipService_GetTrendingFeed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTrendingFeedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlipServiceServer).GetTrendingFeed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BlipService_GetTrendingFeed_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlipServiceServer).GetTrendingFeed(ctx, req.(*GetTrendingFeedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlipService_GetPopularFeed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPopularFeedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlipServiceServer).GetPopularFeed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BlipService_GetPopularFeed_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlipServiceServer).GetPopularFeed(ctx, req.(*GetPopularFeedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BlipService_ServiceDesc is the grpc.ServiceDesc for BlipService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BlipService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "blipper.blip.v1.BlipService",
	HandlerType: (*BlipServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBlip",
			Handler:    _BlipService_CreateBlip_Handler,
		},
		{
			MethodName: "GetBlip",
			Handler:    _BlipService_GetBlip_Handler,
		},
		{
			MethodName: "ListBlips",
			Handler:    _BlipService_ListBlips_Handler,
		},
		{
			MethodName: "DeleteBlip",
			Handler:    _BlipService_DeleteBlip_Handler,
		},
		{
			MethodName: "UpdateBlip",
			Handler:    _BlipService_UpdateBlip_Handler,
		},
		{
			MethodName: "CreateComment",
			Handler:    _BlipService_CreateComment_Handler,
		},
		{
			MethodName: "GetComments",
			Handler:    _BlipService_GetComments_Handler,
		},
		{
			MethodName: "UpdateComment",
			Handler:    _BlipService_UpdateComment_Handler,
		},
		{
			MethodName: "DeleteComment",
			Handler:    _BlipService_DeleteComment_Handler,
		},
		{
			MethodName: "Vote",
			Handler:    _BlipService_Vote_Handler,
		},
		{
			MethodName: "GetVoteStatus",
			Handler:    _BlipService_GetVoteStatus_Handler,
		},
		{
			MethodName: "GetCircleTimeline",
			Handler:    _BlipService_GetCircleTimeline_Handler,
		},
		{
			MethodName: "GetUserTimeline",
			Handler:    _BlipService_GetUserTimeline_Handler,
		},
		{
			MethodName: "GetTrendingFeed",
			Handler:    _BlipService_GetTrendingFeed_Handler,
		},
		{
			MethodName: "GetPopularFeed",
			Handler:    _BlipService_GetPopularFeed_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "blip/blip.proto",
}
