// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: publicacao.proto

package publicacao

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

// PublicacaoServiceClient is the client API for PublicacaoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PublicacaoServiceClient interface {
	UpvotePublicacao(ctx context.Context, in *PublicacaoRequest, opts ...grpc.CallOption) (*PublicacaoResponse, error)
}

type publicacaoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPublicacaoServiceClient(cc grpc.ClientConnInterface) PublicacaoServiceClient {
	return &publicacaoServiceClient{cc}
}

func (c *publicacaoServiceClient) UpvotePublicacao(ctx context.Context, in *PublicacaoRequest, opts ...grpc.CallOption) (*PublicacaoResponse, error) {
	out := new(PublicacaoResponse)
	err := c.cc.Invoke(ctx, "/publicacao.PublicacaoService/UpvotePublicacao", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PublicacaoServiceServer is the server API for PublicacaoService service.
// All implementations must embed UnimplementedPublicacaoServiceServer
// for forward compatibility
type PublicacaoServiceServer interface {
	UpvotePublicacao(context.Context, *PublicacaoRequest) (*PublicacaoResponse, error)
	mustEmbedUnimplementedPublicacaoServiceServer()
}

// UnimplementedPublicacaoServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPublicacaoServiceServer struct {
}

func (UnimplementedPublicacaoServiceServer) UpvotePublicacao(context.Context, *PublicacaoRequest) (*PublicacaoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpvotePublicacao not implemented")
}
func (UnimplementedPublicacaoServiceServer) mustEmbedUnimplementedPublicacaoServiceServer() {}

// UnsafePublicacaoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PublicacaoServiceServer will
// result in compilation errors.
type UnsafePublicacaoServiceServer interface {
	mustEmbedUnimplementedPublicacaoServiceServer()
}

func RegisterPublicacaoServiceServer(s grpc.ServiceRegistrar, srv PublicacaoServiceServer) {
	s.RegisterService(&PublicacaoService_ServiceDesc, srv)
}

func _PublicacaoService_UpvotePublicacao_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublicacaoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublicacaoServiceServer).UpvotePublicacao(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/publicacao.PublicacaoService/UpvotePublicacao",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublicacaoServiceServer).UpvotePublicacao(ctx, req.(*PublicacaoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PublicacaoService_ServiceDesc is the grpc.ServiceDesc for PublicacaoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PublicacaoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "publicacao.PublicacaoService",
	HandlerType: (*PublicacaoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpvotePublicacao",
			Handler:    _PublicacaoService_UpvotePublicacao_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "publicacao.proto",
}
