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
	ObjectRecordService_GetObject_FullMethodName   = "/vhu.v1.ObjectRecordService/GetObject"
	ObjectRecordService_ListObjects_FullMethodName = "/vhu.v1.ObjectRecordService/ListObjects"
)

// ObjectRecordServiceClient is the client API for ObjectRecordService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ObjectRecordServiceClient interface {
	GetObject(ctx context.Context, in *GetObjectRequest, opts ...grpc.CallOption) (*GetObjectResponse, error)
	ListObjects(ctx context.Context, in *ListObjectsRequest, opts ...grpc.CallOption) (*ListObjectsResponse, error)
}

type objectRecordServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewObjectRecordServiceClient(cc grpc.ClientConnInterface) ObjectRecordServiceClient {
	return &objectRecordServiceClient{cc}
}

func (c *objectRecordServiceClient) GetObject(ctx context.Context, in *GetObjectRequest, opts ...grpc.CallOption) (*GetObjectResponse, error) {
	out := new(GetObjectResponse)
	err := c.cc.Invoke(ctx, ObjectRecordService_GetObject_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *objectRecordServiceClient) ListObjects(ctx context.Context, in *ListObjectsRequest, opts ...grpc.CallOption) (*ListObjectsResponse, error) {
	out := new(ListObjectsResponse)
	err := c.cc.Invoke(ctx, ObjectRecordService_ListObjects_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ObjectRecordServiceServer is the server API for ObjectRecordService service.
// All implementations must embed UnimplementedObjectRecordServiceServer
// for forward compatibility
type ObjectRecordServiceServer interface {
	GetObject(context.Context, *GetObjectRequest) (*GetObjectResponse, error)
	ListObjects(context.Context, *ListObjectsRequest) (*ListObjectsResponse, error)
	mustEmbedUnimplementedObjectRecordServiceServer()
}

// UnimplementedObjectRecordServiceServer must be embedded to have forward compatible implementations.
type UnimplementedObjectRecordServiceServer struct {
}

func (UnimplementedObjectRecordServiceServer) GetObject(context.Context, *GetObjectRequest) (*GetObjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetObject not implemented")
}
func (UnimplementedObjectRecordServiceServer) ListObjects(context.Context, *ListObjectsRequest) (*ListObjectsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListObjects not implemented")
}
func (UnimplementedObjectRecordServiceServer) mustEmbedUnimplementedObjectRecordServiceServer() {}

// UnsafeObjectRecordServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ObjectRecordServiceServer will
// result in compilation errors.
type UnsafeObjectRecordServiceServer interface {
	mustEmbedUnimplementedObjectRecordServiceServer()
}

func RegisterObjectRecordServiceServer(s grpc.ServiceRegistrar, srv ObjectRecordServiceServer) {
	s.RegisterService(&ObjectRecordService_ServiceDesc, srv)
}

func _ObjectRecordService_GetObject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetObjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectRecordServiceServer).GetObject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ObjectRecordService_GetObject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectRecordServiceServer).GetObject(ctx, req.(*GetObjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ObjectRecordService_ListObjects_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListObjectsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectRecordServiceServer).ListObjects(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ObjectRecordService_ListObjects_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectRecordServiceServer).ListObjects(ctx, req.(*ListObjectsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ObjectRecordService_ServiceDesc is the grpc.ServiceDesc for ObjectRecordService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ObjectRecordService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "vhu.v1.ObjectRecordService",
	HandlerType: (*ObjectRecordServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetObject",
			Handler:    _ObjectRecordService_GetObject_Handler,
		},
		{
			MethodName: "ListObjects",
			Handler:    _ObjectRecordService_ListObjects_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "vhu/v1/services.proto",
}

const (
	TimelineService_GetTimeline_FullMethodName = "/vhu.v1.TimelineService/GetTimeline"
)

// TimelineServiceClient is the client API for TimelineService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TimelineServiceClient interface {
	GetTimeline(ctx context.Context, in *GetTimelineRequest, opts ...grpc.CallOption) (*GetTimelineResponse, error)
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

// TimelineServiceServer is the server API for TimelineService service.
// All implementations must embed UnimplementedTimelineServiceServer
// for forward compatibility
type TimelineServiceServer interface {
	GetTimeline(context.Context, *GetTimelineRequest) (*GetTimelineResponse, error)
	mustEmbedUnimplementedTimelineServiceServer()
}

// UnimplementedTimelineServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTimelineServiceServer struct {
}

func (UnimplementedTimelineServiceServer) GetTimeline(context.Context, *GetTimelineRequest) (*GetTimelineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTimeline not implemented")
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
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "vhu/v1/services.proto",
}
