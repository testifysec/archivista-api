// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: archivist.proto

package archivist

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ArchivistClient is the client API for Archivist service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ArchivistClient interface {
	GetBySubjectDigest(ctx context.Context, in *GetBySubjectDigestRequest, opts ...grpc.CallOption) (*GetBySubjectDigestResponse, error)
}

type archivistClient struct {
	cc grpc.ClientConnInterface
}

func NewArchivistClient(cc grpc.ClientConnInterface) ArchivistClient {
	return &archivistClient{cc}
}

func (c *archivistClient) GetBySubjectDigest(ctx context.Context, in *GetBySubjectDigestRequest, opts ...grpc.CallOption) (*GetBySubjectDigestResponse, error) {
	out := new(GetBySubjectDigestResponse)
	err := c.cc.Invoke(ctx, "/archivist.Archivist/GetBySubjectDigest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ArchivistServer is the server API for Archivist service.
// All implementations must embed UnimplementedArchivistServer
// for forward compatibility
type ArchivistServer interface {
	GetBySubjectDigest(context.Context, *GetBySubjectDigestRequest) (*GetBySubjectDigestResponse, error)
	mustEmbedUnimplementedArchivistServer()
}

// UnimplementedArchivistServer must be embedded to have forward compatible implementations.
type UnimplementedArchivistServer struct {
}

func (UnimplementedArchivistServer) GetBySubjectDigest(context.Context, *GetBySubjectDigestRequest) (*GetBySubjectDigestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBySubjectDigest not implemented")
}
func (UnimplementedArchivistServer) mustEmbedUnimplementedArchivistServer() {}

// UnsafeArchivistServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ArchivistServer will
// result in compilation errors.
type UnsafeArchivistServer interface {
	mustEmbedUnimplementedArchivistServer()
}

func RegisterArchivistServer(s grpc.ServiceRegistrar, srv ArchivistServer) {
	s.RegisterService(&Archivist_ServiceDesc, srv)
}

func _Archivist_GetBySubjectDigest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBySubjectDigestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArchivistServer).GetBySubjectDigest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/archivist.Archivist/GetBySubjectDigest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArchivistServer).GetBySubjectDigest(ctx, req.(*GetBySubjectDigestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Archivist_ServiceDesc is the grpc.ServiceDesc for Archivist service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Archivist_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "archivist.Archivist",
	HandlerType: (*ArchivistServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBySubjectDigest",
			Handler:    _Archivist_GetBySubjectDigest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "archivist.proto",
}

// CollectorClient is the client API for Collector service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CollectorClient interface {
	Store(ctx context.Context, in *StoreRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type collectorClient struct {
	cc grpc.ClientConnInterface
}

func NewCollectorClient(cc grpc.ClientConnInterface) CollectorClient {
	return &collectorClient{cc}
}

func (c *collectorClient) Store(ctx context.Context, in *StoreRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/archivist.Collector/Store", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CollectorServer is the server API for Collector service.
// All implementations must embed UnimplementedCollectorServer
// for forward compatibility
type CollectorServer interface {
	Store(context.Context, *StoreRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedCollectorServer()
}

// UnimplementedCollectorServer must be embedded to have forward compatible implementations.
type UnimplementedCollectorServer struct {
}

func (UnimplementedCollectorServer) Store(context.Context, *StoreRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Store not implemented")
}
func (UnimplementedCollectorServer) mustEmbedUnimplementedCollectorServer() {}

// UnsafeCollectorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CollectorServer will
// result in compilation errors.
type UnsafeCollectorServer interface {
	mustEmbedUnimplementedCollectorServer()
}

func RegisterCollectorServer(s grpc.ServiceRegistrar, srv CollectorServer) {
	s.RegisterService(&Collector_ServiceDesc, srv)
}

func _Collector_Store_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectorServer).Store(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/archivist.Collector/Store",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectorServer).Store(ctx, req.(*StoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Collector_ServiceDesc is the grpc.ServiceDesc for Collector service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Collector_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "archivist.Collector",
	HandlerType: (*CollectorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Store",
			Handler:    _Collector_Store_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "archivist.proto",
}

// StorageClient is the client API for Storage service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StorageClient interface {
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
}

type storageClient struct {
	cc grpc.ClientConnInterface
}

func NewStorageClient(cc grpc.ClientConnInterface) StorageClient {
	return &storageClient{cc}
}

func (c *storageClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/archivist.Storage/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StorageServer is the server API for Storage service.
// All implementations must embed UnimplementedStorageServer
// for forward compatibility
type StorageServer interface {
	Get(context.Context, *GetRequest) (*GetResponse, error)
	mustEmbedUnimplementedStorageServer()
}

// UnimplementedStorageServer must be embedded to have forward compatible implementations.
type UnimplementedStorageServer struct {
}

func (UnimplementedStorageServer) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedStorageServer) mustEmbedUnimplementedStorageServer() {}

// UnsafeStorageServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StorageServer will
// result in compilation errors.
type UnsafeStorageServer interface {
	mustEmbedUnimplementedStorageServer()
}

func RegisterStorageServer(s grpc.ServiceRegistrar, srv StorageServer) {
	s.RegisterService(&Storage_ServiceDesc, srv)
}

func _Storage_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/archivist.Storage/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Storage_ServiceDesc is the grpc.ServiceDesc for Storage service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Storage_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "archivist.Storage",
	HandlerType: (*StorageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Storage_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "archivist.proto",
}