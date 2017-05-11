// Code generated by protoc-gen-go.
// source: cache.proto
// DO NOT EDIT!

/*
Package cache is a generated protocol buffer package.

It is generated from these files:
	cache.proto

It has these top-level messages:
	AddRequest
	AddReply
	GetRequest
	GetReply
	RemoveRequest
	RemoveReply
	LenRequest
	LenReply
	ClearRequest
	ClearReply
*/
package cache

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type AddRequest struct {
	Key   string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *AddRequest) Reset()                    { *m = AddRequest{} }
func (m *AddRequest) String() string            { return proto.CompactTextString(m) }
func (*AddRequest) ProtoMessage()               {}
func (*AddRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AddRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *AddRequest) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type AddReply struct {
	Key string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
}

func (m *AddReply) Reset()                    { *m = AddReply{} }
func (m *AddReply) String() string            { return proto.CompactTextString(m) }
func (*AddReply) ProtoMessage()               {}
func (*AddReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AddReply) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type GetRequest struct {
	Key string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
}

func (m *GetRequest) Reset()                    { *m = GetRequest{} }
func (m *GetRequest) String() string            { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()               {}
func (*GetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *GetRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type GetReply struct {
	Key   string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Ok    bool   `protobuf:"varint,3,opt,name=ok" json:"ok,omitempty"`
}

func (m *GetReply) Reset()                    { *m = GetReply{} }
func (m *GetReply) String() string            { return proto.CompactTextString(m) }
func (*GetReply) ProtoMessage()               {}
func (*GetReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *GetReply) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *GetReply) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *GetReply) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

type RemoveRequest struct {
	Key string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
}

func (m *RemoveRequest) Reset()                    { *m = RemoveRequest{} }
func (m *RemoveRequest) String() string            { return proto.CompactTextString(m) }
func (*RemoveRequest) ProtoMessage()               {}
func (*RemoveRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *RemoveRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type RemoveReply struct {
	Key string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
}

func (m *RemoveReply) Reset()                    { *m = RemoveReply{} }
func (m *RemoveReply) String() string            { return proto.CompactTextString(m) }
func (*RemoveReply) ProtoMessage()               {}
func (*RemoveReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *RemoveReply) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type LenRequest struct {
}

func (m *LenRequest) Reset()                    { *m = LenRequest{} }
func (m *LenRequest) String() string            { return proto.CompactTextString(m) }
func (*LenRequest) ProtoMessage()               {}
func (*LenRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type LenReply struct {
	Len int32 `protobuf:"varint,1,opt,name=len" json:"len,omitempty"`
}

func (m *LenReply) Reset()                    { *m = LenReply{} }
func (m *LenReply) String() string            { return proto.CompactTextString(m) }
func (*LenReply) ProtoMessage()               {}
func (*LenReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *LenReply) GetLen() int32 {
	if m != nil {
		return m.Len
	}
	return 0
}

type ClearRequest struct {
}

func (m *ClearRequest) Reset()                    { *m = ClearRequest{} }
func (m *ClearRequest) String() string            { return proto.CompactTextString(m) }
func (*ClearRequest) ProtoMessage()               {}
func (*ClearRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

type ClearReply struct {
}

func (m *ClearReply) Reset()                    { *m = ClearReply{} }
func (m *ClearReply) String() string            { return proto.CompactTextString(m) }
func (*ClearReply) ProtoMessage()               {}
func (*ClearReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func init() {
	proto.RegisterType((*AddRequest)(nil), "cache.AddRequest")
	proto.RegisterType((*AddReply)(nil), "cache.AddReply")
	proto.RegisterType((*GetRequest)(nil), "cache.GetRequest")
	proto.RegisterType((*GetReply)(nil), "cache.GetReply")
	proto.RegisterType((*RemoveRequest)(nil), "cache.RemoveRequest")
	proto.RegisterType((*RemoveReply)(nil), "cache.RemoveReply")
	proto.RegisterType((*LenRequest)(nil), "cache.LenRequest")
	proto.RegisterType((*LenReply)(nil), "cache.LenReply")
	proto.RegisterType((*ClearRequest)(nil), "cache.ClearRequest")
	proto.RegisterType((*ClearReply)(nil), "cache.ClearReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Cache service

type CacheClient interface {
	Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddReply, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetReply, error)
	Remove(ctx context.Context, in *RemoveRequest, opts ...grpc.CallOption) (*RemoveReply, error)
	Len(ctx context.Context, in *LenRequest, opts ...grpc.CallOption) (*LenReply, error)
	Clear(ctx context.Context, in *ClearRequest, opts ...grpc.CallOption) (*ClearReply, error)
}

type cacheClient struct {
	cc *grpc.ClientConn
}

func NewCacheClient(cc *grpc.ClientConn) CacheClient {
	return &cacheClient{cc}
}

func (c *cacheClient) Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddReply, error) {
	out := new(AddReply)
	err := grpc.Invoke(ctx, "/cache.Cache/Add", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetReply, error) {
	out := new(GetReply)
	err := grpc.Invoke(ctx, "/cache.Cache/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheClient) Remove(ctx context.Context, in *RemoveRequest, opts ...grpc.CallOption) (*RemoveReply, error) {
	out := new(RemoveReply)
	err := grpc.Invoke(ctx, "/cache.Cache/Remove", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheClient) Len(ctx context.Context, in *LenRequest, opts ...grpc.CallOption) (*LenReply, error) {
	out := new(LenReply)
	err := grpc.Invoke(ctx, "/cache.Cache/Len", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheClient) Clear(ctx context.Context, in *ClearRequest, opts ...grpc.CallOption) (*ClearReply, error) {
	out := new(ClearReply)
	err := grpc.Invoke(ctx, "/cache.Cache/Clear", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Cache service

type CacheServer interface {
	Add(context.Context, *AddRequest) (*AddReply, error)
	Get(context.Context, *GetRequest) (*GetReply, error)
	Remove(context.Context, *RemoveRequest) (*RemoveReply, error)
	Len(context.Context, *LenRequest) (*LenReply, error)
	Clear(context.Context, *ClearRequest) (*ClearReply, error)
}

func RegisterCacheServer(s *grpc.Server, srv CacheServer) {
	s.RegisterService(&_Cache_serviceDesc, srv)
}

func _Cache_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cache.Cache/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServer).Add(ctx, req.(*AddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cache_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cache.Cache/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cache_Remove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServer).Remove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cache.Cache/Remove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServer).Remove(ctx, req.(*RemoveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cache_Len_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServer).Len(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cache.Cache/Len",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServer).Len(ctx, req.(*LenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cache_Clear_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClearRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServer).Clear(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cache.Cache/Clear",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServer).Clear(ctx, req.(*ClearRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Cache_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cache.Cache",
	HandlerType: (*CacheServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _Cache_Add_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Cache_Get_Handler,
		},
		{
			MethodName: "Remove",
			Handler:    _Cache_Remove_Handler,
		},
		{
			MethodName: "Len",
			Handler:    _Cache_Len_Handler,
		},
		{
			MethodName: "Clear",
			Handler:    _Cache_Clear_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cache.proto",
}

func init() { proto.RegisterFile("cache.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 289 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0xcd, 0x4e, 0xb3, 0x50,
	0x14, 0x2c, 0x10, 0x1a, 0xbe, 0x29, 0x5f, 0xb5, 0xc7, 0x2e, 0x08, 0x31, 0x8a, 0x77, 0x45, 0x62,
	0xd2, 0x44, 0xed, 0x0b, 0xd4, 0x2e, 0xba, 0xe9, 0xea, 0xbe, 0x01, 0x96, 0x93, 0x98, 0x70, 0x05,
	0xac, 0xb4, 0x09, 0x4f, 0xee, 0xd6, 0x70, 0xf9, 0x6b, 0x95, 0xee, 0xee, 0x9c, 0x33, 0x73, 0x86,
	0x99, 0x80, 0xc9, 0x2e, 0xda, 0xbd, 0xf3, 0x22, 0xdf, 0x67, 0x45, 0x46, 0xb6, 0x06, 0x62, 0x09,
	0xac, 0xe2, 0x58, 0xf2, 0xe7, 0x81, 0xbf, 0x0a, 0xba, 0x86, 0x95, 0x70, 0xe9, 0x19, 0x81, 0x11,
	0xfe, 0x93, 0xd5, 0x93, 0xe6, 0xb0, 0x8f, 0x91, 0x3a, 0xb0, 0x67, 0x06, 0x46, 0xe8, 0xca, 0x1a,
	0x88, 0x5b, 0x38, 0x5a, 0x95, 0xab, 0xf2, 0xaf, 0x46, 0xdc, 0x01, 0x1b, 0x2e, 0x2e, 0xde, 0x14,
	0xaf, 0x70, 0xf4, 0x7e, 0x50, 0x3d, 0xec, 0x48, 0x53, 0x98, 0x59, 0xe2, 0x59, 0x81, 0x11, 0x3a,
	0xd2, 0xcc, 0x12, 0xf1, 0x80, 0xff, 0x92, 0x3f, 0xb2, 0x23, 0x5f, 0xb6, 0xb9, 0xc7, 0xa4, 0xa5,
	0x0c, 0x7f, 0xa7, 0x0b, 0x6c, 0x39, 0x6d, 0x0e, 0x54, 0x99, 0x34, 0x6a, 0xb8, 0x8a, 0x53, 0xcd,
	0xb5, 0x65, 0xf5, 0x14, 0x53, 0xb8, 0x6b, 0xc5, 0xd1, 0xbe, 0x65, 0xbb, 0x40, 0x83, 0x73, 0x55,
	0x3e, 0x7f, 0x1b, 0xb0, 0xd7, 0x55, 0x9f, 0xf4, 0x08, 0x6b, 0x15, 0xc7, 0x34, 0x5b, 0xd4, 0x5d,
	0xf7, 0xdd, 0xfa, 0x57, 0xa7, 0xa3, 0x5c, 0x95, 0x62, 0x54, 0x91, 0x37, 0x5c, 0x74, 0xe4, 0xbe,
	0xb4, 0x8e, 0xdc, 0xf6, 0x24, 0x46, 0xb4, 0xc4, 0xb8, 0x8e, 0x43, 0xf3, 0x66, 0x79, 0x56, 0x80,
	0x4f, 0xbf, 0xa6, 0x9d, 0xc5, 0x96, 0xd3, 0xce, 0xa2, 0xcf, 0xdb, 0x59, 0xb4, 0xa1, 0xc5, 0x88,
	0x9e, 0x60, 0xeb, 0x50, 0x74, 0xd3, 0xec, 0x4e, 0x23, 0xfb, 0xb3, 0xf3, 0xa1, 0x96, 0xbc, 0x8d,
	0xf5, 0xdf, 0xf4, 0xf2, 0x13, 0x00, 0x00, 0xff, 0xff, 0x15, 0x2b, 0x61, 0x0b, 0x5c, 0x02, 0x00,
	0x00,
}
