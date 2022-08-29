// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: internal/proto/word.proto

package proto

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

// WordServiceClient is the client API for WordService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WordServiceClient interface {
	// 特定の文字列をユーザの単語帳として登録する。
	CreateWord(ctx context.Context, in *CreateWordRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 指定したユーザIDの単語一覧を表示する。
	ListWord(ctx context.Context, in *ListWordRequest, opts ...grpc.CallOption) (*ListWordResponse, error)
	// 指定したIDの単語の説明文を変更する。
	UpdateWord(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 指定したIDの単語を削除する。
	DeleteWord(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type wordServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWordServiceClient(cc grpc.ClientConnInterface) WordServiceClient {
	return &wordServiceClient{cc}
}

func (c *wordServiceClient) CreateWord(ctx context.Context, in *CreateWordRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/word.WordService/CreateWord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wordServiceClient) ListWord(ctx context.Context, in *ListWordRequest, opts ...grpc.CallOption) (*ListWordResponse, error) {
	out := new(ListWordResponse)
	err := c.cc.Invoke(ctx, "/word.WordService/ListWord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wordServiceClient) UpdateWord(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/word.WordService/UpdateWord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wordServiceClient) DeleteWord(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/word.WordService/DeleteWord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WordServiceServer is the server API for WordService service.
// All implementations should embed UnimplementedWordServiceServer
// for forward compatibility
type WordServiceServer interface {
	// 特定の文字列をユーザの単語帳として登録する。
	CreateWord(context.Context, *CreateWordRequest) (*emptypb.Empty, error)
	// 指定したユーザIDの単語一覧を表示する。
	ListWord(context.Context, *ListWordRequest) (*ListWordResponse, error)
	// 指定したIDの単語の説明文を変更する。
	UpdateWord(context.Context, *UpdateRequest) (*emptypb.Empty, error)
	// 指定したIDの単語を削除する。
	DeleteWord(context.Context, *DeleteRequest) (*emptypb.Empty, error)
}

// UnimplementedWordServiceServer should be embedded to have forward compatible implementations.
type UnimplementedWordServiceServer struct {
}

func (UnimplementedWordServiceServer) CreateWord(context.Context, *CreateWordRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateWord not implemented")
}
func (UnimplementedWordServiceServer) ListWord(context.Context, *ListWordRequest) (*ListWordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListWord not implemented")
}
func (UnimplementedWordServiceServer) UpdateWord(context.Context, *UpdateRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateWord not implemented")
}
func (UnimplementedWordServiceServer) DeleteWord(context.Context, *DeleteRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteWord not implemented")
}

// UnsafeWordServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WordServiceServer will
// result in compilation errors.
type UnsafeWordServiceServer interface {
	mustEmbedUnimplementedWordServiceServer()
}

func RegisterWordServiceServer(s grpc.ServiceRegistrar, srv WordServiceServer) {
	s.RegisterService(&WordService_ServiceDesc, srv)
}

func _WordService_CreateWord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateWordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WordServiceServer).CreateWord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/word.WordService/CreateWord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WordServiceServer).CreateWord(ctx, req.(*CreateWordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WordService_ListWord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListWordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WordServiceServer).ListWord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/word.WordService/ListWord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WordServiceServer).ListWord(ctx, req.(*ListWordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WordService_UpdateWord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WordServiceServer).UpdateWord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/word.WordService/UpdateWord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WordServiceServer).UpdateWord(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WordService_DeleteWord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WordServiceServer).DeleteWord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/word.WordService/DeleteWord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WordServiceServer).DeleteWord(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// WordService_ServiceDesc is the grpc.ServiceDesc for WordService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WordService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "word.WordService",
	HandlerType: (*WordServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateWord",
			Handler:    _WordService_CreateWord_Handler,
		},
		{
			MethodName: "ListWord",
			Handler:    _WordService_ListWord_Handler,
		},
		{
			MethodName: "UpdateWord",
			Handler:    _WordService_UpdateWord_Handler,
		},
		{
			MethodName: "DeleteWord",
			Handler:    _WordService_DeleteWord_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/proto/word.proto",
}
