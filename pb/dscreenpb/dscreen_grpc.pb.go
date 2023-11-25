// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: pb/dscreenpb/dscreen.proto

package dscreenpb

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
	DScreenService_Import_FullMethodName                  = "/dcmp.pb.dscreenpb.DScreenService/Import"
	DScreenService_Platforms_FullMethodName               = "/dcmp.pb.dscreenpb.DScreenService/Platforms"
	DScreenService_GetNFTByPhone_FullMethodName           = "/dcmp.pb.dscreenpb.DScreenService/GetNFTByPhone"
	DScreenService_Display_FullMethodName                 = "/dcmp.pb.dscreenpb.DScreenService/Display"
	DScreenService_QueryProgress_FullMethodName           = "/dcmp.pb.dscreenpb.DScreenService/QueryProgress"
	DScreenService_QueryDisplayLog_FullMethodName         = "/dcmp.pb.dscreenpb.DScreenService/QueryDisplayLog"
	DScreenService_ApplyEnter_FullMethodName              = "/dcmp.pb.dscreenpb.DScreenService/ApplyEnter"
	DScreenService_GetScreencastByDeviceId_FullMethodName = "/dcmp.pb.dscreenpb.DScreenService/GetScreencastByDeviceId"
	DScreenService_BatchNotify_FullMethodName             = "/dcmp.pb.dscreenpb.DScreenService/BatchNotify"
	DScreenService_GetNFTByOwner_FullMethodName           = "/dcmp.pb.dscreenpb.DScreenService/GetNFTByOwner"
	DScreenService_SetAppVersion_FullMethodName           = "/dcmp.pb.dscreenpb.DScreenService/SetAppVersion"
	DScreenService_GetAppVersion_FullMethodName           = "/dcmp.pb.dscreenpb.DScreenService/getAppVersion"
	DScreenService_DelAppVersion_FullMethodName           = "/dcmp.pb.dscreenpb.DScreenService/DelAppVersion"
	DScreenService_SetDeviceVersion_FullMethodName        = "/dcmp.pb.dscreenpb.DScreenService/SetDeviceVersion"
	DScreenService_GetDeviceVersion_FullMethodName        = "/dcmp.pb.dscreenpb.DScreenService/getDeviceVersion"
)

// DScreenServiceClient is the client API for DScreenService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DScreenServiceClient interface {
	Import(ctx context.Context, in *ImportKeyReq, opts ...grpc.CallOption) (*ImportKeyResp, error)
	Platforms(ctx context.Context, in *PlatformsReq, opts ...grpc.CallOption) (*PlatformsResp, error)
	GetNFTByPhone(ctx context.Context, in *GetNFTByPhoneReq, opts ...grpc.CallOption) (*GetNFTByPhoneResp, error)
	Display(ctx context.Context, in *DisplayReq, opts ...grpc.CallOption) (*DisplayResp, error)
	QueryProgress(ctx context.Context, in *QueryProgressReq, opts ...grpc.CallOption) (*QueryProgressResp, error)
	QueryDisplayLog(ctx context.Context, in *QueryDisplayLogReq, opts ...grpc.CallOption) (*QueryDisplayLogResp, error)
	ApplyEnter(ctx context.Context, in *ApplyEnterReq, opts ...grpc.CallOption) (*ApplyEnterResp, error)
	GetScreencastByDeviceId(ctx context.Context, in *GetScreenCastByDeviceIdReq, opts ...grpc.CallOption) (*GetScreenCastByDeviceIdResp, error)
	BatchNotify(ctx context.Context, in *BatchNotifyReq, opts ...grpc.CallOption) (*BassResp, error)
	GetNFTByOwner(ctx context.Context, in *GetNFTByOwnerReq, opts ...grpc.CallOption) (*GetNFTByOwnerResp, error)
	SetAppVersion(ctx context.Context, in *SetAppVersionReq, opts ...grpc.CallOption) (*SetAppVersionResp, error)
	GetAppVersion(ctx context.Context, in *GetAppVersionReq, opts ...grpc.CallOption) (*GetAppVersionResp, error)
	DelAppVersion(ctx context.Context, in *DelAppVersionReq, opts ...grpc.CallOption) (*DelAppVersionResp, error)
	SetDeviceVersion(ctx context.Context, in *SetDeviceVersionReq, opts ...grpc.CallOption) (*SetDeviceVersionResp, error)
	GetDeviceVersion(ctx context.Context, in *GetDeviceVersionReq, opts ...grpc.CallOption) (*GetDeviceVersionResp, error)
}

type dScreenServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDScreenServiceClient(cc grpc.ClientConnInterface) DScreenServiceClient {
	return &dScreenServiceClient{cc}
}

func (c *dScreenServiceClient) Import(ctx context.Context, in *ImportKeyReq, opts ...grpc.CallOption) (*ImportKeyResp, error) {
	out := new(ImportKeyResp)
	err := c.cc.Invoke(ctx, DScreenService_Import_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dScreenServiceClient) Platforms(ctx context.Context, in *PlatformsReq, opts ...grpc.CallOption) (*PlatformsResp, error) {
	out := new(PlatformsResp)
	err := c.cc.Invoke(ctx, DScreenService_Platforms_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dScreenServiceClient) GetNFTByPhone(ctx context.Context, in *GetNFTByPhoneReq, opts ...grpc.CallOption) (*GetNFTByPhoneResp, error) {
	out := new(GetNFTByPhoneResp)
	err := c.cc.Invoke(ctx, DScreenService_GetNFTByPhone_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dScreenServiceClient) Display(ctx context.Context, in *DisplayReq, opts ...grpc.CallOption) (*DisplayResp, error) {
	out := new(DisplayResp)
	err := c.cc.Invoke(ctx, DScreenService_Display_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dScreenServiceClient) QueryProgress(ctx context.Context, in *QueryProgressReq, opts ...grpc.CallOption) (*QueryProgressResp, error) {
	out := new(QueryProgressResp)
	err := c.cc.Invoke(ctx, DScreenService_QueryProgress_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dScreenServiceClient) QueryDisplayLog(ctx context.Context, in *QueryDisplayLogReq, opts ...grpc.CallOption) (*QueryDisplayLogResp, error) {
	out := new(QueryDisplayLogResp)
	err := c.cc.Invoke(ctx, DScreenService_QueryDisplayLog_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dScreenServiceClient) ApplyEnter(ctx context.Context, in *ApplyEnterReq, opts ...grpc.CallOption) (*ApplyEnterResp, error) {
	out := new(ApplyEnterResp)
	err := c.cc.Invoke(ctx, DScreenService_ApplyEnter_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dScreenServiceClient) GetScreencastByDeviceId(ctx context.Context, in *GetScreenCastByDeviceIdReq, opts ...grpc.CallOption) (*GetScreenCastByDeviceIdResp, error) {
	out := new(GetScreenCastByDeviceIdResp)
	err := c.cc.Invoke(ctx, DScreenService_GetScreencastByDeviceId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dScreenServiceClient) BatchNotify(ctx context.Context, in *BatchNotifyReq, opts ...grpc.CallOption) (*BassResp, error) {
	out := new(BassResp)
	err := c.cc.Invoke(ctx, DScreenService_BatchNotify_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dScreenServiceClient) GetNFTByOwner(ctx context.Context, in *GetNFTByOwnerReq, opts ...grpc.CallOption) (*GetNFTByOwnerResp, error) {
	out := new(GetNFTByOwnerResp)
	err := c.cc.Invoke(ctx, DScreenService_GetNFTByOwner_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dScreenServiceClient) SetAppVersion(ctx context.Context, in *SetAppVersionReq, opts ...grpc.CallOption) (*SetAppVersionResp, error) {
	out := new(SetAppVersionResp)
	err := c.cc.Invoke(ctx, DScreenService_SetAppVersion_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dScreenServiceClient) GetAppVersion(ctx context.Context, in *GetAppVersionReq, opts ...grpc.CallOption) (*GetAppVersionResp, error) {
	out := new(GetAppVersionResp)
	err := c.cc.Invoke(ctx, DScreenService_GetAppVersion_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dScreenServiceClient) DelAppVersion(ctx context.Context, in *DelAppVersionReq, opts ...grpc.CallOption) (*DelAppVersionResp, error) {
	out := new(DelAppVersionResp)
	err := c.cc.Invoke(ctx, DScreenService_DelAppVersion_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dScreenServiceClient) SetDeviceVersion(ctx context.Context, in *SetDeviceVersionReq, opts ...grpc.CallOption) (*SetDeviceVersionResp, error) {
	out := new(SetDeviceVersionResp)
	err := c.cc.Invoke(ctx, DScreenService_SetDeviceVersion_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dScreenServiceClient) GetDeviceVersion(ctx context.Context, in *GetDeviceVersionReq, opts ...grpc.CallOption) (*GetDeviceVersionResp, error) {
	out := new(GetDeviceVersionResp)
	err := c.cc.Invoke(ctx, DScreenService_GetDeviceVersion_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DScreenServiceServer is the server API for DScreenService service.
// All implementations should embed UnimplementedDScreenServiceServer
// for forward compatibility
type DScreenServiceServer interface {
	Import(context.Context, *ImportKeyReq) (*ImportKeyResp, error)
	Platforms(context.Context, *PlatformsReq) (*PlatformsResp, error)
	GetNFTByPhone(context.Context, *GetNFTByPhoneReq) (*GetNFTByPhoneResp, error)
	Display(context.Context, *DisplayReq) (*DisplayResp, error)
	QueryProgress(context.Context, *QueryProgressReq) (*QueryProgressResp, error)
	QueryDisplayLog(context.Context, *QueryDisplayLogReq) (*QueryDisplayLogResp, error)
	ApplyEnter(context.Context, *ApplyEnterReq) (*ApplyEnterResp, error)
	GetScreencastByDeviceId(context.Context, *GetScreenCastByDeviceIdReq) (*GetScreenCastByDeviceIdResp, error)
	BatchNotify(context.Context, *BatchNotifyReq) (*BassResp, error)
	GetNFTByOwner(context.Context, *GetNFTByOwnerReq) (*GetNFTByOwnerResp, error)
	SetAppVersion(context.Context, *SetAppVersionReq) (*SetAppVersionResp, error)
	GetAppVersion(context.Context, *GetAppVersionReq) (*GetAppVersionResp, error)
	DelAppVersion(context.Context, *DelAppVersionReq) (*DelAppVersionResp, error)
	SetDeviceVersion(context.Context, *SetDeviceVersionReq) (*SetDeviceVersionResp, error)
	GetDeviceVersion(context.Context, *GetDeviceVersionReq) (*GetDeviceVersionResp, error)
}

// UnimplementedDScreenServiceServer should be embedded to have forward compatible implementations.
type UnimplementedDScreenServiceServer struct {
}

func (UnimplementedDScreenServiceServer) Import(context.Context, *ImportKeyReq) (*ImportKeyResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Import not implemented")
}
func (UnimplementedDScreenServiceServer) Platforms(context.Context, *PlatformsReq) (*PlatformsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Platforms not implemented")
}
func (UnimplementedDScreenServiceServer) GetNFTByPhone(context.Context, *GetNFTByPhoneReq) (*GetNFTByPhoneResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNFTByPhone not implemented")
}
func (UnimplementedDScreenServiceServer) Display(context.Context, *DisplayReq) (*DisplayResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Display not implemented")
}
func (UnimplementedDScreenServiceServer) QueryProgress(context.Context, *QueryProgressReq) (*QueryProgressResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryProgress not implemented")
}
func (UnimplementedDScreenServiceServer) QueryDisplayLog(context.Context, *QueryDisplayLogReq) (*QueryDisplayLogResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryDisplayLog not implemented")
}
func (UnimplementedDScreenServiceServer) ApplyEnter(context.Context, *ApplyEnterReq) (*ApplyEnterResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApplyEnter not implemented")
}
func (UnimplementedDScreenServiceServer) GetScreencastByDeviceId(context.Context, *GetScreenCastByDeviceIdReq) (*GetScreenCastByDeviceIdResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetScreencastByDeviceId not implemented")
}
func (UnimplementedDScreenServiceServer) BatchNotify(context.Context, *BatchNotifyReq) (*BassResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchNotify not implemented")
}
func (UnimplementedDScreenServiceServer) GetNFTByOwner(context.Context, *GetNFTByOwnerReq) (*GetNFTByOwnerResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNFTByOwner not implemented")
}
func (UnimplementedDScreenServiceServer) SetAppVersion(context.Context, *SetAppVersionReq) (*SetAppVersionResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetAppVersion not implemented")
}
func (UnimplementedDScreenServiceServer) GetAppVersion(context.Context, *GetAppVersionReq) (*GetAppVersionResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppVersion not implemented")
}
func (UnimplementedDScreenServiceServer) DelAppVersion(context.Context, *DelAppVersionReq) (*DelAppVersionResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelAppVersion not implemented")
}
func (UnimplementedDScreenServiceServer) SetDeviceVersion(context.Context, *SetDeviceVersionReq) (*SetDeviceVersionResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetDeviceVersion not implemented")
}
func (UnimplementedDScreenServiceServer) GetDeviceVersion(context.Context, *GetDeviceVersionReq) (*GetDeviceVersionResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDeviceVersion not implemented")
}

// UnsafeDScreenServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DScreenServiceServer will
// result in compilation errors.
type UnsafeDScreenServiceServer interface {
	mustEmbedUnimplementedDScreenServiceServer()
}

func RegisterDScreenServiceServer(s grpc.ServiceRegistrar, srv DScreenServiceServer) {
	s.RegisterService(&DScreenService_ServiceDesc, srv)
}

func _DScreenService_Import_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImportKeyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DScreenServiceServer).Import(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DScreenService_Import_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DScreenServiceServer).Import(ctx, req.(*ImportKeyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DScreenService_Platforms_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlatformsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DScreenServiceServer).Platforms(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DScreenService_Platforms_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DScreenServiceServer).Platforms(ctx, req.(*PlatformsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DScreenService_GetNFTByPhone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNFTByPhoneReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DScreenServiceServer).GetNFTByPhone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DScreenService_GetNFTByPhone_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DScreenServiceServer).GetNFTByPhone(ctx, req.(*GetNFTByPhoneReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DScreenService_Display_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DisplayReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DScreenServiceServer).Display(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DScreenService_Display_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DScreenServiceServer).Display(ctx, req.(*DisplayReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DScreenService_QueryProgress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryProgressReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DScreenServiceServer).QueryProgress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DScreenService_QueryProgress_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DScreenServiceServer).QueryProgress(ctx, req.(*QueryProgressReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DScreenService_QueryDisplayLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDisplayLogReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DScreenServiceServer).QueryDisplayLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DScreenService_QueryDisplayLog_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DScreenServiceServer).QueryDisplayLog(ctx, req.(*QueryDisplayLogReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DScreenService_ApplyEnter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplyEnterReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DScreenServiceServer).ApplyEnter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DScreenService_ApplyEnter_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DScreenServiceServer).ApplyEnter(ctx, req.(*ApplyEnterReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DScreenService_GetScreencastByDeviceId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetScreenCastByDeviceIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DScreenServiceServer).GetScreencastByDeviceId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DScreenService_GetScreencastByDeviceId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DScreenServiceServer).GetScreencastByDeviceId(ctx, req.(*GetScreenCastByDeviceIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DScreenService_BatchNotify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchNotifyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DScreenServiceServer).BatchNotify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DScreenService_BatchNotify_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DScreenServiceServer).BatchNotify(ctx, req.(*BatchNotifyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DScreenService_GetNFTByOwner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNFTByOwnerReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DScreenServiceServer).GetNFTByOwner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DScreenService_GetNFTByOwner_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DScreenServiceServer).GetNFTByOwner(ctx, req.(*GetNFTByOwnerReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DScreenService_SetAppVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetAppVersionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DScreenServiceServer).SetAppVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DScreenService_SetAppVersion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DScreenServiceServer).SetAppVersion(ctx, req.(*SetAppVersionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DScreenService_GetAppVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAppVersionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DScreenServiceServer).GetAppVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DScreenService_GetAppVersion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DScreenServiceServer).GetAppVersion(ctx, req.(*GetAppVersionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DScreenService_DelAppVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelAppVersionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DScreenServiceServer).DelAppVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DScreenService_DelAppVersion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DScreenServiceServer).DelAppVersion(ctx, req.(*DelAppVersionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DScreenService_SetDeviceVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetDeviceVersionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DScreenServiceServer).SetDeviceVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DScreenService_SetDeviceVersion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DScreenServiceServer).SetDeviceVersion(ctx, req.(*SetDeviceVersionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DScreenService_GetDeviceVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDeviceVersionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DScreenServiceServer).GetDeviceVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DScreenService_GetDeviceVersion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DScreenServiceServer).GetDeviceVersion(ctx, req.(*GetDeviceVersionReq))
	}
	return interceptor(ctx, in, info, handler)
}

// DScreenService_ServiceDesc is the grpc.ServiceDesc for DScreenService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DScreenService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dcmp.pb.dscreenpb.DScreenService",
	HandlerType: (*DScreenServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Import",
			Handler:    _DScreenService_Import_Handler,
		},
		{
			MethodName: "Platforms",
			Handler:    _DScreenService_Platforms_Handler,
		},
		{
			MethodName: "GetNFTByPhone",
			Handler:    _DScreenService_GetNFTByPhone_Handler,
		},
		{
			MethodName: "Display",
			Handler:    _DScreenService_Display_Handler,
		},
		{
			MethodName: "QueryProgress",
			Handler:    _DScreenService_QueryProgress_Handler,
		},
		{
			MethodName: "QueryDisplayLog",
			Handler:    _DScreenService_QueryDisplayLog_Handler,
		},
		{
			MethodName: "ApplyEnter",
			Handler:    _DScreenService_ApplyEnter_Handler,
		},
		{
			MethodName: "GetScreencastByDeviceId",
			Handler:    _DScreenService_GetScreencastByDeviceId_Handler,
		},
		{
			MethodName: "BatchNotify",
			Handler:    _DScreenService_BatchNotify_Handler,
		},
		{
			MethodName: "GetNFTByOwner",
			Handler:    _DScreenService_GetNFTByOwner_Handler,
		},
		{
			MethodName: "SetAppVersion",
			Handler:    _DScreenService_SetAppVersion_Handler,
		},
		{
			MethodName: "getAppVersion",
			Handler:    _DScreenService_GetAppVersion_Handler,
		},
		{
			MethodName: "DelAppVersion",
			Handler:    _DScreenService_DelAppVersion_Handler,
		},
		{
			MethodName: "SetDeviceVersion",
			Handler:    _DScreenService_SetDeviceVersion_Handler,
		},
		{
			MethodName: "getDeviceVersion",
			Handler:    _DScreenService_GetDeviceVersion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/dscreenpb/dscreen.proto",
}