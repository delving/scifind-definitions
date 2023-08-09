// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: vhu/v1/services.proto

package vhu

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
	TimelineService_GetTimeline_FullMethodName   = "/vhu.v1.TimelineService/GetTimeline"
	TimelineService_ListTimelines_FullMethodName = "/vhu.v1.TimelineService/ListTimelines"
)

// TimelineServiceClient is the client API for TimelineService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TimelineServiceClient interface {
	GetTimeline(ctx context.Context, in *GetTimelineRequest, opts ...grpc.CallOption) (*GetTimelineResponse, error)
	ListTimelines(ctx context.Context, in *ListTimelinesRequest, opts ...grpc.CallOption) (*ListTimelinesResponse, error)
}

type timelineServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTimelineServiceClient(cc grpc.ClientConnInterface) TimelineServiceClient {
	return &timelineServiceClient{cc}
}

func (c *timelineServiceClient) GetTimeline(ctx context.Context, in *GetTimelineRequest, opts ...grpc.CallOption) (*GetTimelineResponse, error) {
	out := new(GetTimelineResponse)
	err := c.cc.Invoke(ctx, TimelineService_GetTimeline_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *timelineServiceClient) ListTimelines(ctx context.Context, in *ListTimelinesRequest, opts ...grpc.CallOption) (*ListTimelinesResponse, error) {
	out := new(ListTimelinesResponse)
	err := c.cc.Invoke(ctx, TimelineService_ListTimelines_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TimelineServiceServer is the server API for TimelineService service.
// All implementations must embed UnimplementedTimelineServiceServer
// for forward compatibility
type TimelineServiceServer interface {
	GetTimeline(context.Context, *GetTimelineRequest) (*GetTimelineResponse, error)
	ListTimelines(context.Context, *ListTimelinesRequest) (*ListTimelinesResponse, error)
	mustEmbedUnimplementedTimelineServiceServer()
}

// UnimplementedTimelineServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTimelineServiceServer struct {
}

func (UnimplementedTimelineServiceServer) GetTimeline(context.Context, *GetTimelineRequest) (*GetTimelineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTimeline not implemented")
}
func (UnimplementedTimelineServiceServer) ListTimelines(context.Context, *ListTimelinesRequest) (*ListTimelinesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTimelines not implemented")
}
func (UnimplementedTimelineServiceServer) mustEmbedUnimplementedTimelineServiceServer() {}

// UnsafeTimelineServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TimelineServiceServer will
// result in compilation errors.
type UnsafeTimelineServiceServer interface {
	mustEmbedUnimplementedTimelineServiceServer()
}

func RegisterTimelineServiceServer(s grpc.ServiceRegistrar, srv TimelineServiceServer) {
	s.RegisterService(&TimelineService_ServiceDesc, srv)
}

func _TimelineService_GetTimeline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTimelineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimelineServiceServer).GetTimeline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TimelineService_GetTimeline_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimelineServiceServer).GetTimeline(ctx, req.(*GetTimelineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TimelineService_ListTimelines_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTimelinesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimelineServiceServer).ListTimelines(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TimelineService_ListTimelines_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimelineServiceServer).ListTimelines(ctx, req.(*ListTimelinesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TimelineService_ServiceDesc is the grpc.ServiceDesc for TimelineService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TimelineService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "vhu.v1.TimelineService",
	HandlerType: (*TimelineServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTimeline",
			Handler:    _TimelineService_GetTimeline_Handler,
		},
		{
			MethodName: "ListTimelines",
			Handler:    _TimelineService_ListTimelines_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "vhu/v1/services.proto",
}

const (
	MuseumObjectService_GetMuseumObject_FullMethodName   = "/vhu.v1.MuseumObjectService/GetMuseumObject"
	MuseumObjectService_ListMuseumObjects_FullMethodName = "/vhu.v1.MuseumObjectService/ListMuseumObjects"
)

// MuseumObjectServiceClient is the client API for MuseumObjectService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MuseumObjectServiceClient interface {
	GetMuseumObject(ctx context.Context, in *GetMuseumObjectRequest, opts ...grpc.CallOption) (*GetMuseumObjectResponse, error)
	ListMuseumObjects(ctx context.Context, in *ListMuseumObjectsRequest, opts ...grpc.CallOption) (*ListMuseumObjectsResponse, error)
}

type museumObjectServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMuseumObjectServiceClient(cc grpc.ClientConnInterface) MuseumObjectServiceClient {
	return &museumObjectServiceClient{cc}
}

func (c *museumObjectServiceClient) GetMuseumObject(ctx context.Context, in *GetMuseumObjectRequest, opts ...grpc.CallOption) (*GetMuseumObjectResponse, error) {
	out := new(GetMuseumObjectResponse)
	err := c.cc.Invoke(ctx, MuseumObjectService_GetMuseumObject_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *museumObjectServiceClient) ListMuseumObjects(ctx context.Context, in *ListMuseumObjectsRequest, opts ...grpc.CallOption) (*ListMuseumObjectsResponse, error) {
	out := new(ListMuseumObjectsResponse)
	err := c.cc.Invoke(ctx, MuseumObjectService_ListMuseumObjects_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MuseumObjectServiceServer is the server API for MuseumObjectService service.
// All implementations must embed UnimplementedMuseumObjectServiceServer
// for forward compatibility
type MuseumObjectServiceServer interface {
	GetMuseumObject(context.Context, *GetMuseumObjectRequest) (*GetMuseumObjectResponse, error)
	ListMuseumObjects(context.Context, *ListMuseumObjectsRequest) (*ListMuseumObjectsResponse, error)
	mustEmbedUnimplementedMuseumObjectServiceServer()
}

// UnimplementedMuseumObjectServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMuseumObjectServiceServer struct {
}

func (UnimplementedMuseumObjectServiceServer) GetMuseumObject(context.Context, *GetMuseumObjectRequest) (*GetMuseumObjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMuseumObject not implemented")
}
func (UnimplementedMuseumObjectServiceServer) ListMuseumObjects(context.Context, *ListMuseumObjectsRequest) (*ListMuseumObjectsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMuseumObjects not implemented")
}
func (UnimplementedMuseumObjectServiceServer) mustEmbedUnimplementedMuseumObjectServiceServer() {}

// UnsafeMuseumObjectServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MuseumObjectServiceServer will
// result in compilation errors.
type UnsafeMuseumObjectServiceServer interface {
	mustEmbedUnimplementedMuseumObjectServiceServer()
}

func RegisterMuseumObjectServiceServer(s grpc.ServiceRegistrar, srv MuseumObjectServiceServer) {
	s.RegisterService(&MuseumObjectService_ServiceDesc, srv)
}

func _MuseumObjectService_GetMuseumObject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMuseumObjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MuseumObjectServiceServer).GetMuseumObject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MuseumObjectService_GetMuseumObject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MuseumObjectServiceServer).GetMuseumObject(ctx, req.(*GetMuseumObjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MuseumObjectService_ListMuseumObjects_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMuseumObjectsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MuseumObjectServiceServer).ListMuseumObjects(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MuseumObjectService_ListMuseumObjects_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MuseumObjectServiceServer).ListMuseumObjects(ctx, req.(*ListMuseumObjectsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MuseumObjectService_ServiceDesc is the grpc.ServiceDesc for MuseumObjectService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MuseumObjectService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "vhu.v1.MuseumObjectService",
	HandlerType: (*MuseumObjectServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMuseumObject",
			Handler:    _MuseumObjectService_GetMuseumObject_Handler,
		},
		{
			MethodName: "ListMuseumObjects",
			Handler:    _MuseumObjectService_ListMuseumObjects_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "vhu/v1/services.proto",
}
