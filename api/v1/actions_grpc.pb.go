// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// ActionServiceClient is the client API for ActionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ActionServiceClient interface {
	GetAction(ctx context.Context, in *GetActionRequest, opts ...grpc.CallOption) (*GetActionResponse, error)
	AddAction(ctx context.Context, in *AddActionRequest, opts ...grpc.CallOption) (*AddActionResponse, error)
	UpdateAction(ctx context.Context, in *UpdateActionRequest, opts ...grpc.CallOption) (*UpdateActionResponse, error)
	DeleteAction(ctx context.Context, in *DeleteActionRequest, opts ...grpc.CallOption) (*DeleteActionResponse, error)
}

type actionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewActionServiceClient(cc grpc.ClientConnInterface) ActionServiceClient {
	return &actionServiceClient{cc}
}

func (c *actionServiceClient) GetAction(ctx context.Context, in *GetActionRequest, opts ...grpc.CallOption) (*GetActionResponse, error) {
	out := new(GetActionResponse)
	err := c.cc.Invoke(ctx, "/ActionService/GetAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actionServiceClient) AddAction(ctx context.Context, in *AddActionRequest, opts ...grpc.CallOption) (*AddActionResponse, error) {
	out := new(AddActionResponse)
	err := c.cc.Invoke(ctx, "/ActionService/AddAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actionServiceClient) UpdateAction(ctx context.Context, in *UpdateActionRequest, opts ...grpc.CallOption) (*UpdateActionResponse, error) {
	out := new(UpdateActionResponse)
	err := c.cc.Invoke(ctx, "/ActionService/UpdateAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actionServiceClient) DeleteAction(ctx context.Context, in *DeleteActionRequest, opts ...grpc.CallOption) (*DeleteActionResponse, error) {
	out := new(DeleteActionResponse)
	err := c.cc.Invoke(ctx, "/ActionService/DeleteAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ActionServiceServer is the server API for ActionService service.
// All implementations must embed UnimplementedActionServiceServer
// for forward compatibility
type ActionServiceServer interface {
	GetAction(context.Context, *GetActionRequest) (*GetActionResponse, error)
	AddAction(context.Context, *AddActionRequest) (*AddActionResponse, error)
	UpdateAction(context.Context, *UpdateActionRequest) (*UpdateActionResponse, error)
	DeleteAction(context.Context, *DeleteActionRequest) (*DeleteActionResponse, error)
	mustEmbedUnimplementedActionServiceServer()
}

// UnimplementedActionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedActionServiceServer struct {
}

func (UnimplementedActionServiceServer) GetAction(context.Context, *GetActionRequest) (*GetActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAction not implemented")
}
func (UnimplementedActionServiceServer) AddAction(context.Context, *AddActionRequest) (*AddActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddAction not implemented")
}
func (UnimplementedActionServiceServer) UpdateAction(context.Context, *UpdateActionRequest) (*UpdateActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAction not implemented")
}
func (UnimplementedActionServiceServer) DeleteAction(context.Context, *DeleteActionRequest) (*DeleteActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAction not implemented")
}
func (UnimplementedActionServiceServer) mustEmbedUnimplementedActionServiceServer() {}

// UnsafeActionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ActionServiceServer will
// result in compilation errors.
type UnsafeActionServiceServer interface {
	mustEmbedUnimplementedActionServiceServer()
}

func RegisterActionServiceServer(s *grpc.Server, srv ActionServiceServer) {
	s.RegisterService(&_ActionService_serviceDesc, srv)
}

func _ActionService_GetAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionServiceServer).GetAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ActionService/GetAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionServiceServer).GetAction(ctx, req.(*GetActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActionService_AddAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionServiceServer).AddAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ActionService/AddAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionServiceServer).AddAction(ctx, req.(*AddActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActionService_UpdateAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionServiceServer).UpdateAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ActionService/UpdateAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionServiceServer).UpdateAction(ctx, req.(*UpdateActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActionService_DeleteAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionServiceServer).DeleteAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ActionService/DeleteAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionServiceServer).DeleteAction(ctx, req.(*DeleteActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ActionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ActionService",
	HandlerType: (*ActionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAction",
			Handler:    _ActionService_GetAction_Handler,
		},
		{
			MethodName: "AddAction",
			Handler:    _ActionService_AddAction_Handler,
		},
		{
			MethodName: "UpdateAction",
			Handler:    _ActionService_UpdateAction_Handler,
		},
		{
			MethodName: "DeleteAction",
			Handler:    _ActionService_DeleteAction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1/actions.proto",
}