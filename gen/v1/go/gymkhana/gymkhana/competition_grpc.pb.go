// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.1
// source: gymkhana/competition.proto

package gymkhanav1

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

// CompetitionServiceClient is the client API for CompetitionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CompetitionServiceClient interface {
	Competition(ctx context.Context, in *CompetitionRequest, opts ...grpc.CallOption) (*CompetitionResponse, error)
	CompetitionStages(ctx context.Context, in *CompetitionRequest, opts ...grpc.CallOption) (CompetitionService_CompetitionStagesClient, error)
	ListCompetitions(ctx context.Context, in *CompetitionListRequest, opts ...grpc.CallOption) (CompetitionService_ListCompetitionsClient, error)
	CreateCompetition(ctx context.Context, in *CompetitionCreateRequest, opts ...grpc.CallOption) (*CompetitionResponse, error)
	UpdateCompetition(ctx context.Context, in *CompetitionUpdateRequest, opts ...grpc.CallOption) (*CompetitionResponse, error)
	DeleteCompetition(ctx context.Context, in *CompetitionRequest, opts ...grpc.CallOption) (*CompetitionResponse, error)
}

type competitionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCompetitionServiceClient(cc grpc.ClientConnInterface) CompetitionServiceClient {
	return &competitionServiceClient{cc}
}

func (c *competitionServiceClient) Competition(ctx context.Context, in *CompetitionRequest, opts ...grpc.CallOption) (*CompetitionResponse, error) {
	out := new(CompetitionResponse)
	err := c.cc.Invoke(ctx, "/gymkhana.CompetitionService/Competition", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *competitionServiceClient) CompetitionStages(ctx context.Context, in *CompetitionRequest, opts ...grpc.CallOption) (CompetitionService_CompetitionStagesClient, error) {
	stream, err := c.cc.NewStream(ctx, &CompetitionService_ServiceDesc.Streams[0], "/gymkhana.CompetitionService/CompetitionStages", opts...)
	if err != nil {
		return nil, err
	}
	x := &competitionServiceCompetitionStagesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CompetitionService_CompetitionStagesClient interface {
	Recv() (*CompetitionStage, error)
	grpc.ClientStream
}

type competitionServiceCompetitionStagesClient struct {
	grpc.ClientStream
}

func (x *competitionServiceCompetitionStagesClient) Recv() (*CompetitionStage, error) {
	m := new(CompetitionStage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *competitionServiceClient) ListCompetitions(ctx context.Context, in *CompetitionListRequest, opts ...grpc.CallOption) (CompetitionService_ListCompetitionsClient, error) {
	stream, err := c.cc.NewStream(ctx, &CompetitionService_ServiceDesc.Streams[1], "/gymkhana.CompetitionService/ListCompetitions", opts...)
	if err != nil {
		return nil, err
	}
	x := &competitionServiceListCompetitionsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CompetitionService_ListCompetitionsClient interface {
	Recv() (*CompetitionResponse, error)
	grpc.ClientStream
}

type competitionServiceListCompetitionsClient struct {
	grpc.ClientStream
}

func (x *competitionServiceListCompetitionsClient) Recv() (*CompetitionResponse, error) {
	m := new(CompetitionResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *competitionServiceClient) CreateCompetition(ctx context.Context, in *CompetitionCreateRequest, opts ...grpc.CallOption) (*CompetitionResponse, error) {
	out := new(CompetitionResponse)
	err := c.cc.Invoke(ctx, "/gymkhana.CompetitionService/CreateCompetition", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *competitionServiceClient) UpdateCompetition(ctx context.Context, in *CompetitionUpdateRequest, opts ...grpc.CallOption) (*CompetitionResponse, error) {
	out := new(CompetitionResponse)
	err := c.cc.Invoke(ctx, "/gymkhana.CompetitionService/UpdateCompetition", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *competitionServiceClient) DeleteCompetition(ctx context.Context, in *CompetitionRequest, opts ...grpc.CallOption) (*CompetitionResponse, error) {
	out := new(CompetitionResponse)
	err := c.cc.Invoke(ctx, "/gymkhana.CompetitionService/DeleteCompetition", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CompetitionServiceServer is the server API for CompetitionService service.
// All implementations must embed UnimplementedCompetitionServiceServer
// for forward compatibility
type CompetitionServiceServer interface {
	Competition(context.Context, *CompetitionRequest) (*CompetitionResponse, error)
	CompetitionStages(*CompetitionRequest, CompetitionService_CompetitionStagesServer) error
	ListCompetitions(*CompetitionListRequest, CompetitionService_ListCompetitionsServer) error
	CreateCompetition(context.Context, *CompetitionCreateRequest) (*CompetitionResponse, error)
	UpdateCompetition(context.Context, *CompetitionUpdateRequest) (*CompetitionResponse, error)
	DeleteCompetition(context.Context, *CompetitionRequest) (*CompetitionResponse, error)
	mustEmbedUnimplementedCompetitionServiceServer()
}

// UnimplementedCompetitionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCompetitionServiceServer struct {
}

func (UnimplementedCompetitionServiceServer) Competition(context.Context, *CompetitionRequest) (*CompetitionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Competition not implemented")
}
func (UnimplementedCompetitionServiceServer) CompetitionStages(*CompetitionRequest, CompetitionService_CompetitionStagesServer) error {
	return status.Errorf(codes.Unimplemented, "method CompetitionStages not implemented")
}
func (UnimplementedCompetitionServiceServer) ListCompetitions(*CompetitionListRequest, CompetitionService_ListCompetitionsServer) error {
	return status.Errorf(codes.Unimplemented, "method ListCompetitions not implemented")
}
func (UnimplementedCompetitionServiceServer) CreateCompetition(context.Context, *CompetitionCreateRequest) (*CompetitionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCompetition not implemented")
}
func (UnimplementedCompetitionServiceServer) UpdateCompetition(context.Context, *CompetitionUpdateRequest) (*CompetitionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCompetition not implemented")
}
func (UnimplementedCompetitionServiceServer) DeleteCompetition(context.Context, *CompetitionRequest) (*CompetitionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCompetition not implemented")
}
func (UnimplementedCompetitionServiceServer) mustEmbedUnimplementedCompetitionServiceServer() {}

// UnsafeCompetitionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CompetitionServiceServer will
// result in compilation errors.
type UnsafeCompetitionServiceServer interface {
	mustEmbedUnimplementedCompetitionServiceServer()
}

func RegisterCompetitionServiceServer(s grpc.ServiceRegistrar, srv CompetitionServiceServer) {
	s.RegisterService(&CompetitionService_ServiceDesc, srv)
}

func _CompetitionService_Competition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompetitionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompetitionServiceServer).Competition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gymkhana.CompetitionService/Competition",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompetitionServiceServer).Competition(ctx, req.(*CompetitionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompetitionService_CompetitionStages_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(CompetitionRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CompetitionServiceServer).CompetitionStages(m, &competitionServiceCompetitionStagesServer{stream})
}

type CompetitionService_CompetitionStagesServer interface {
	Send(*CompetitionStage) error
	grpc.ServerStream
}

type competitionServiceCompetitionStagesServer struct {
	grpc.ServerStream
}

func (x *competitionServiceCompetitionStagesServer) Send(m *CompetitionStage) error {
	return x.ServerStream.SendMsg(m)
}

func _CompetitionService_ListCompetitions_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(CompetitionListRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CompetitionServiceServer).ListCompetitions(m, &competitionServiceListCompetitionsServer{stream})
}

type CompetitionService_ListCompetitionsServer interface {
	Send(*CompetitionResponse) error
	grpc.ServerStream
}

type competitionServiceListCompetitionsServer struct {
	grpc.ServerStream
}

func (x *competitionServiceListCompetitionsServer) Send(m *CompetitionResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _CompetitionService_CreateCompetition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompetitionCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompetitionServiceServer).CreateCompetition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gymkhana.CompetitionService/CreateCompetition",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompetitionServiceServer).CreateCompetition(ctx, req.(*CompetitionCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompetitionService_UpdateCompetition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompetitionUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompetitionServiceServer).UpdateCompetition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gymkhana.CompetitionService/UpdateCompetition",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompetitionServiceServer).UpdateCompetition(ctx, req.(*CompetitionUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompetitionService_DeleteCompetition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompetitionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompetitionServiceServer).DeleteCompetition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gymkhana.CompetitionService/DeleteCompetition",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompetitionServiceServer).DeleteCompetition(ctx, req.(*CompetitionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CompetitionService_ServiceDesc is the grpc.ServiceDesc for CompetitionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CompetitionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gymkhana.CompetitionService",
	HandlerType: (*CompetitionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Competition",
			Handler:    _CompetitionService_Competition_Handler,
		},
		{
			MethodName: "CreateCompetition",
			Handler:    _CompetitionService_CreateCompetition_Handler,
		},
		{
			MethodName: "UpdateCompetition",
			Handler:    _CompetitionService_UpdateCompetition_Handler,
		},
		{
			MethodName: "DeleteCompetition",
			Handler:    _CompetitionService_DeleteCompetition_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "CompetitionStages",
			Handler:       _CompetitionService_CompetitionStages_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ListCompetitions",
			Handler:       _CompetitionService_ListCompetitions_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "gymkhana/competition.proto",
}
