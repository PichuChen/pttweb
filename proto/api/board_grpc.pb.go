// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package api

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

// BoardServiceClient is the client API for BoardService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BoardServiceClient interface {
	Board(ctx context.Context, in *BoardRequest, opts ...grpc.CallOption) (*BoardReply, error)
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListReply, error)
	Hotboard(ctx context.Context, in *HotboardRequest, opts ...grpc.CallOption) (*HotboardReply, error)
	Content(ctx context.Context, in *ContentRequest, opts ...grpc.CallOption) (*ContentReply, error)
	Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchReply, error)
}

type boardServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBoardServiceClient(cc grpc.ClientConnInterface) BoardServiceClient {
	return &boardServiceClient{cc}
}

func (c *boardServiceClient) Board(ctx context.Context, in *BoardRequest, opts ...grpc.CallOption) (*BoardReply, error) {
	out := new(BoardReply)
	err := c.cc.Invoke(ctx, "/pttbbs.api.BoardService/Board", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *boardServiceClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListReply, error) {
	out := new(ListReply)
	err := c.cc.Invoke(ctx, "/pttbbs.api.BoardService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *boardServiceClient) Hotboard(ctx context.Context, in *HotboardRequest, opts ...grpc.CallOption) (*HotboardReply, error) {
	out := new(HotboardReply)
	err := c.cc.Invoke(ctx, "/pttbbs.api.BoardService/Hotboard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *boardServiceClient) Content(ctx context.Context, in *ContentRequest, opts ...grpc.CallOption) (*ContentReply, error) {
	out := new(ContentReply)
	err := c.cc.Invoke(ctx, "/pttbbs.api.BoardService/Content", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *boardServiceClient) Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchReply, error) {
	out := new(SearchReply)
	err := c.cc.Invoke(ctx, "/pttbbs.api.BoardService/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BoardServiceServer is the server API for BoardService service.
// All implementations must embed UnimplementedBoardServiceServer
// for forward compatibility
type BoardServiceServer interface {
	Board(context.Context, *BoardRequest) (*BoardReply, error)
	List(context.Context, *ListRequest) (*ListReply, error)
	Hotboard(context.Context, *HotboardRequest) (*HotboardReply, error)
	Content(context.Context, *ContentRequest) (*ContentReply, error)
	Search(context.Context, *SearchRequest) (*SearchReply, error)
	mustEmbedUnimplementedBoardServiceServer()
}

// UnimplementedBoardServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBoardServiceServer struct {
}

func (UnimplementedBoardServiceServer) Board(context.Context, *BoardRequest) (*BoardReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Board not implemented")
}
func (UnimplementedBoardServiceServer) List(context.Context, *ListRequest) (*ListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedBoardServiceServer) Hotboard(context.Context, *HotboardRequest) (*HotboardReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Hotboard not implemented")
}
func (UnimplementedBoardServiceServer) Content(context.Context, *ContentRequest) (*ContentReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Content not implemented")
}
func (UnimplementedBoardServiceServer) Search(context.Context, *SearchRequest) (*SearchReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}
func (UnimplementedBoardServiceServer) mustEmbedUnimplementedBoardServiceServer() {}

// UnsafeBoardServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BoardServiceServer will
// result in compilation errors.
type UnsafeBoardServiceServer interface {
	mustEmbedUnimplementedBoardServiceServer()
}

func RegisterBoardServiceServer(s grpc.ServiceRegistrar, srv BoardServiceServer) {
	s.RegisterService(&BoardService_ServiceDesc, srv)
}

func _BoardService_Board_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BoardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BoardServiceServer).Board(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pttbbs.api.BoardService/Board",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BoardServiceServer).Board(ctx, req.(*BoardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BoardService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BoardServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pttbbs.api.BoardService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BoardServiceServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BoardService_Hotboard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HotboardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BoardServiceServer).Hotboard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pttbbs.api.BoardService/Hotboard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BoardServiceServer).Hotboard(ctx, req.(*HotboardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BoardService_Content_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ContentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BoardServiceServer).Content(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pttbbs.api.BoardService/Content",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BoardServiceServer).Content(ctx, req.(*ContentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BoardService_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BoardServiceServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pttbbs.api.BoardService/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BoardServiceServer).Search(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BoardService_ServiceDesc is the grpc.ServiceDesc for BoardService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BoardService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pttbbs.api.BoardService",
	HandlerType: (*BoardServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Board",
			Handler:    _BoardService_Board_Handler,
		},
		{
			MethodName: "List",
			Handler:    _BoardService_List_Handler,
		},
		{
			MethodName: "Hotboard",
			Handler:    _BoardService_Hotboard_Handler,
		},
		{
			MethodName: "Content",
			Handler:    _BoardService_Content_Handler,
		},
		{
			MethodName: "Search",
			Handler:    _BoardService_Search_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "board.proto",
}
