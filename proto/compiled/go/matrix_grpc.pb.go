// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb_matrix

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// MatrixClient is the client API for Matrix service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MatrixClient interface {
	StartSession(ctx context.Context, in *StartSessionRequest, opts ...grpc.CallOption) (Matrix_StartSessionClient, error)
	GetAccessToken(ctx context.Context, in *GetAccessTokenRequest, opts ...grpc.CallOption) (*GetAccessTokenResponse, error)
	JoinRoom(ctx context.Context, in *JoinRoomRequest, opts ...grpc.CallOption) (*JoinRoomResponse, error)
	Leave(ctx context.Context, in *LeaveRequest, opts ...grpc.CallOption) (*LeaveResponse, error)
	SendTextMessage(ctx context.Context, in *SendTextMessageRequest, opts ...grpc.CallOption) (*SendTextMessageResponse, error)
	SendReadReceipt(ctx context.Context, in *SendReadReceiptRequest, opts ...grpc.CallOption) (*SendReadReceiptResponse, error)
}

type matrixClient struct {
	cc grpc.ClientConnInterface
}

func NewMatrixClient(cc grpc.ClientConnInterface) MatrixClient {
	return &matrixClient{cc}
}

func (c *matrixClient) StartSession(ctx context.Context, in *StartSessionRequest, opts ...grpc.CallOption) (Matrix_StartSessionClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Matrix_serviceDesc.Streams[0], "/matrix.Matrix/StartSession", opts...)
	if err != nil {
		return nil, err
	}
	x := &matrixStartSessionClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Matrix_StartSessionClient interface {
	Recv() (*SessionEvent, error)
	grpc.ClientStream
}

type matrixStartSessionClient struct {
	grpc.ClientStream
}

func (x *matrixStartSessionClient) Recv() (*SessionEvent, error) {
	m := new(SessionEvent)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *matrixClient) GetAccessToken(ctx context.Context, in *GetAccessTokenRequest, opts ...grpc.CallOption) (*GetAccessTokenResponse, error) {
	out := new(GetAccessTokenResponse)
	err := c.cc.Invoke(ctx, "/matrix.Matrix/GetAccessToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *matrixClient) JoinRoom(ctx context.Context, in *JoinRoomRequest, opts ...grpc.CallOption) (*JoinRoomResponse, error) {
	out := new(JoinRoomResponse)
	err := c.cc.Invoke(ctx, "/matrix.Matrix/JoinRoom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *matrixClient) Leave(ctx context.Context, in *LeaveRequest, opts ...grpc.CallOption) (*LeaveResponse, error) {
	out := new(LeaveResponse)
	err := c.cc.Invoke(ctx, "/matrix.Matrix/Leave", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *matrixClient) SendTextMessage(ctx context.Context, in *SendTextMessageRequest, opts ...grpc.CallOption) (*SendTextMessageResponse, error) {
	out := new(SendTextMessageResponse)
	err := c.cc.Invoke(ctx, "/matrix.Matrix/SendTextMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *matrixClient) SendReadReceipt(ctx context.Context, in *SendReadReceiptRequest, opts ...grpc.CallOption) (*SendReadReceiptResponse, error) {
	out := new(SendReadReceiptResponse)
	err := c.cc.Invoke(ctx, "/matrix.Matrix/SendReadReceipt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MatrixServer is the server API for Matrix service.
// All implementations must embed UnimplementedMatrixServer
// for forward compatibility
type MatrixServer interface {
	StartSession(*StartSessionRequest, Matrix_StartSessionServer) error
	GetAccessToken(context.Context, *GetAccessTokenRequest) (*GetAccessTokenResponse, error)
	JoinRoom(context.Context, *JoinRoomRequest) (*JoinRoomResponse, error)
	Leave(context.Context, *LeaveRequest) (*LeaveResponse, error)
	SendTextMessage(context.Context, *SendTextMessageRequest) (*SendTextMessageResponse, error)
	SendReadReceipt(context.Context, *SendReadReceiptRequest) (*SendReadReceiptResponse, error)
	mustEmbedUnimplementedMatrixServer()
}

// UnimplementedMatrixServer must be embedded to have forward compatible implementations.
type UnimplementedMatrixServer struct {
}

func (UnimplementedMatrixServer) StartSession(*StartSessionRequest, Matrix_StartSessionServer) error {
	return status.Errorf(codes.Unimplemented, "method StartSession not implemented")
}
func (UnimplementedMatrixServer) GetAccessToken(context.Context, *GetAccessTokenRequest) (*GetAccessTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccessToken not implemented")
}
func (UnimplementedMatrixServer) JoinRoom(context.Context, *JoinRoomRequest) (*JoinRoomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinRoom not implemented")
}
func (UnimplementedMatrixServer) Leave(context.Context, *LeaveRequest) (*LeaveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Leave not implemented")
}
func (UnimplementedMatrixServer) SendTextMessage(context.Context, *SendTextMessageRequest) (*SendTextMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendTextMessage not implemented")
}
func (UnimplementedMatrixServer) SendReadReceipt(context.Context, *SendReadReceiptRequest) (*SendReadReceiptResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendReadReceipt not implemented")
}
func (UnimplementedMatrixServer) mustEmbedUnimplementedMatrixServer() {}

// UnsafeMatrixServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MatrixServer will
// result in compilation errors.
type UnsafeMatrixServer interface {
	mustEmbedUnimplementedMatrixServer()
}

func RegisterMatrixServer(s *grpc.Server, srv MatrixServer) {
	s.RegisterService(&_Matrix_serviceDesc, srv)
}

func _Matrix_StartSession_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StartSessionRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MatrixServer).StartSession(m, &matrixStartSessionServer{stream})
}

type Matrix_StartSessionServer interface {
	Send(*SessionEvent) error
	grpc.ServerStream
}

type matrixStartSessionServer struct {
	grpc.ServerStream
}

func (x *matrixStartSessionServer) Send(m *SessionEvent) error {
	return x.ServerStream.SendMsg(m)
}

func _Matrix_GetAccessToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccessTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MatrixServer).GetAccessToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/matrix.Matrix/GetAccessToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MatrixServer).GetAccessToken(ctx, req.(*GetAccessTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Matrix_JoinRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinRoomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MatrixServer).JoinRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/matrix.Matrix/JoinRoom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MatrixServer).JoinRoom(ctx, req.(*JoinRoomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Matrix_Leave_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LeaveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MatrixServer).Leave(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/matrix.Matrix/Leave",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MatrixServer).Leave(ctx, req.(*LeaveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Matrix_SendTextMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendTextMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MatrixServer).SendTextMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/matrix.Matrix/SendTextMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MatrixServer).SendTextMessage(ctx, req.(*SendTextMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Matrix_SendReadReceipt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendReadReceiptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MatrixServer).SendReadReceipt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/matrix.Matrix/SendReadReceipt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MatrixServer).SendReadReceipt(ctx, req.(*SendReadReceiptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Matrix_serviceDesc = grpc.ServiceDesc{
	ServiceName: "matrix.Matrix",
	HandlerType: (*MatrixServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAccessToken",
			Handler:    _Matrix_GetAccessToken_Handler,
		},
		{
			MethodName: "JoinRoom",
			Handler:    _Matrix_JoinRoom_Handler,
		},
		{
			MethodName: "Leave",
			Handler:    _Matrix_Leave_Handler,
		},
		{
			MethodName: "SendTextMessage",
			Handler:    _Matrix_SendTextMessage_Handler,
		},
		{
			MethodName: "SendReadReceipt",
			Handler:    _Matrix_SendReadReceipt_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StartSession",
			Handler:       _Matrix_StartSession_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "matrix.proto",
}