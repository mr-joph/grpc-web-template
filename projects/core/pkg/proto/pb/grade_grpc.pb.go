// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: proto/grade.proto

package pb

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// GradeServiceClient is the client API for GradeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GradeServiceClient interface {
	GetEducationLevels(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetEducationLevelsOut, error)
}

type gradeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGradeServiceClient(cc grpc.ClientConnInterface) GradeServiceClient {
	return &gradeServiceClient{cc}
}

func (c *gradeServiceClient) GetEducationLevels(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetEducationLevelsOut, error) {
	out := new(GetEducationLevelsOut)
	err := c.cc.Invoke(ctx, "/core.GradeService/GetEducationLevels", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GradeServiceServer is the server API for GradeService service.
// All implementations must embed UnimplementedGradeServiceServer
// for forward compatibility
type GradeServiceServer interface {
	GetEducationLevels(context.Context, *empty.Empty) (*GetEducationLevelsOut, error)
	mustEmbedUnimplementedGradeServiceServer()
}

// UnimplementedGradeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGradeServiceServer struct {
}

func (UnimplementedGradeServiceServer) GetEducationLevels(context.Context, *empty.Empty) (*GetEducationLevelsOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEducationLevels not implemented")
}
func (UnimplementedGradeServiceServer) mustEmbedUnimplementedGradeServiceServer() {}

// UnsafeGradeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GradeServiceServer will
// result in compilation errors.
type UnsafeGradeServiceServer interface {
	mustEmbedUnimplementedGradeServiceServer()
}

func RegisterGradeServiceServer(s grpc.ServiceRegistrar, srv GradeServiceServer) {
	s.RegisterService(&GradeService_ServiceDesc, srv)
}

func _GradeService_GetEducationLevels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GradeServiceServer).GetEducationLevels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.GradeService/GetEducationLevels",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GradeServiceServer).GetEducationLevels(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// GradeService_ServiceDesc is the grpc.ServiceDesc for GradeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GradeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "core.GradeService",
	HandlerType: (*GradeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetEducationLevels",
			Handler:    _GradeService_GetEducationLevels_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/grade.proto",
}
