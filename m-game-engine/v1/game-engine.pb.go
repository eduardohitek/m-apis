// Code generated by protoc-gen-go. DO NOT EDIT.
// source: game-engine.proto

package gameengine

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type GetSizeRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetSizeRequest) Reset()         { *m = GetSizeRequest{} }
func (m *GetSizeRequest) String() string { return proto.CompactTextString(m) }
func (*GetSizeRequest) ProtoMessage()    {}
func (*GetSizeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_383212ef93ee0b9b, []int{0}
}

func (m *GetSizeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetSizeRequest.Unmarshal(m, b)
}
func (m *GetSizeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetSizeRequest.Marshal(b, m, deterministic)
}
func (m *GetSizeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSizeRequest.Merge(m, src)
}
func (m *GetSizeRequest) XXX_Size() int {
	return xxx_messageInfo_GetSizeRequest.Size(m)
}
func (m *GetSizeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSizeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetSizeRequest proto.InternalMessageInfo

type GetSizeResponse struct {
	Size                 float64  `protobuf:"fixed64,1,opt,name=size,proto3" json:"size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetSizeResponse) Reset()         { *m = GetSizeResponse{} }
func (m *GetSizeResponse) String() string { return proto.CompactTextString(m) }
func (*GetSizeResponse) ProtoMessage()    {}
func (*GetSizeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_383212ef93ee0b9b, []int{1}
}

func (m *GetSizeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetSizeResponse.Unmarshal(m, b)
}
func (m *GetSizeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetSizeResponse.Marshal(b, m, deterministic)
}
func (m *GetSizeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSizeResponse.Merge(m, src)
}
func (m *GetSizeResponse) XXX_Size() int {
	return xxx_messageInfo_GetSizeResponse.Size(m)
}
func (m *GetSizeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSizeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetSizeResponse proto.InternalMessageInfo

func (m *GetSizeResponse) GetSize() float64 {
	if m != nil {
		return m.Size
	}
	return 0
}

type SetScoreRequest struct {
	Score                float64  `protobuf:"fixed64,1,opt,name=score,proto3" json:"score,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetScoreRequest) Reset()         { *m = SetScoreRequest{} }
func (m *SetScoreRequest) String() string { return proto.CompactTextString(m) }
func (*SetScoreRequest) ProtoMessage()    {}
func (*SetScoreRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_383212ef93ee0b9b, []int{2}
}

func (m *SetScoreRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetScoreRequest.Unmarshal(m, b)
}
func (m *SetScoreRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetScoreRequest.Marshal(b, m, deterministic)
}
func (m *SetScoreRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetScoreRequest.Merge(m, src)
}
func (m *SetScoreRequest) XXX_Size() int {
	return xxx_messageInfo_SetScoreRequest.Size(m)
}
func (m *SetScoreRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetScoreRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetScoreRequest proto.InternalMessageInfo

func (m *SetScoreRequest) GetScore() float64 {
	if m != nil {
		return m.Score
	}
	return 0
}

type SetScoreResponse struct {
	Set                  bool     `protobuf:"varint,1,opt,name=set,proto3" json:"set,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetScoreResponse) Reset()         { *m = SetScoreResponse{} }
func (m *SetScoreResponse) String() string { return proto.CompactTextString(m) }
func (*SetScoreResponse) ProtoMessage()    {}
func (*SetScoreResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_383212ef93ee0b9b, []int{3}
}

func (m *SetScoreResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetScoreResponse.Unmarshal(m, b)
}
func (m *SetScoreResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetScoreResponse.Marshal(b, m, deterministic)
}
func (m *SetScoreResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetScoreResponse.Merge(m, src)
}
func (m *SetScoreResponse) XXX_Size() int {
	return xxx_messageInfo_SetScoreResponse.Size(m)
}
func (m *SetScoreResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SetScoreResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SetScoreResponse proto.InternalMessageInfo

func (m *SetScoreResponse) GetSet() bool {
	if m != nil {
		return m.Set
	}
	return false
}

func init() {
	proto.RegisterType((*GetSizeRequest)(nil), "m.gameengine.v1.GetSizeRequest")
	proto.RegisterType((*GetSizeResponse)(nil), "m.gameengine.v1.GetSizeResponse")
	proto.RegisterType((*SetScoreRequest)(nil), "m.gameengine.v1.SetScoreRequest")
	proto.RegisterType((*SetScoreResponse)(nil), "m.gameengine.v1.SetScoreResponse")
}

func init() { proto.RegisterFile("game-engine.proto", fileDescriptor_383212ef93ee0b9b) }

var fileDescriptor_383212ef93ee0b9b = []byte{
	// 205 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4c, 0x4f, 0xcc, 0x4d,
	0xd5, 0x4d, 0xcd, 0x4b, 0xcf, 0xcc, 0x4b, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xcf,
	0xd5, 0x03, 0x09, 0x42, 0xc5, 0xca, 0x0c, 0x95, 0x04, 0xb8, 0xf8, 0xdc, 0x53, 0x4b, 0x82, 0x33,
	0xab, 0x52, 0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x94, 0x54, 0xb9, 0xf8, 0xe1, 0x22, 0xc5,
	0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x42, 0x42, 0x5c, 0x2c, 0xc5, 0x99, 0x55, 0xa9, 0x12, 0x8c, 0x0a,
	0x8c, 0x1a, 0x8c, 0x41, 0x60, 0xb6, 0x92, 0x3a, 0x17, 0x7f, 0x70, 0x6a, 0x49, 0x70, 0x72, 0x7e,
	0x11, 0x4c, 0xa7, 0x90, 0x08, 0x17, 0x6b, 0x31, 0x88, 0x0f, 0x55, 0x07, 0xe1, 0x28, 0xa9, 0x70,
	0x09, 0x20, 0x14, 0x42, 0x0d, 0x14, 0xe0, 0x62, 0x2e, 0x4e, 0x2d, 0x01, 0xab, 0xe3, 0x08, 0x02,
	0x31, 0x8d, 0x56, 0x33, 0x72, 0x71, 0xb9, 0x27, 0xe6, 0xa6, 0xba, 0x82, 0x5d, 0x26, 0xe4, 0xc3,
	0xc5, 0x0e, 0x75, 0x84, 0x90, 0xbc, 0x1e, 0x9a, 0x9b, 0xf5, 0x50, 0x1d, 0x2c, 0xa5, 0x80, 0x5b,
	0x01, 0xd4, 0x3a, 0x7f, 0x2e, 0x0e, 0x98, 0x13, 0x84, 0x30, 0x55, 0xa3, 0x79, 0x43, 0x4a, 0x11,
	0x8f, 0x0a, 0x88, 0x81, 0x4e, 0x3c, 0x51, 0x5c, 0x08, 0x15, 0x49, 0x6c, 0xe0, 0xb0, 0x35, 0x06,
	0x04, 0x00, 0x00, 0xff, 0xff, 0xcd, 0xf1, 0x50, 0xe0, 0x70, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GameEngineClient is the client API for GameEngine service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GameEngineClient interface {
	GetSize(ctx context.Context, in *GetSizeRequest, opts ...grpc.CallOption) (*GetSizeResponse, error)
	SetScore(ctx context.Context, in *SetScoreRequest, opts ...grpc.CallOption) (*SetScoreResponse, error)
}

type gameEngineClient struct {
	cc *grpc.ClientConn
}

func NewGameEngineClient(cc *grpc.ClientConn) GameEngineClient {
	return &gameEngineClient{cc}
}

func (c *gameEngineClient) GetSize(ctx context.Context, in *GetSizeRequest, opts ...grpc.CallOption) (*GetSizeResponse, error) {
	out := new(GetSizeResponse)
	err := c.cc.Invoke(ctx, "/m.gameengine.v1.GameEngine/GetSize", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameEngineClient) SetScore(ctx context.Context, in *SetScoreRequest, opts ...grpc.CallOption) (*SetScoreResponse, error) {
	out := new(SetScoreResponse)
	err := c.cc.Invoke(ctx, "/m.gameengine.v1.GameEngine/SetScore", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GameEngineServer is the server API for GameEngine service.
type GameEngineServer interface {
	GetSize(context.Context, *GetSizeRequest) (*GetSizeResponse, error)
	SetScore(context.Context, *SetScoreRequest) (*SetScoreResponse, error)
}

// UnimplementedGameEngineServer can be embedded to have forward compatible implementations.
type UnimplementedGameEngineServer struct {
}

func (*UnimplementedGameEngineServer) GetSize(ctx context.Context, req *GetSizeRequest) (*GetSizeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSize not implemented")
}
func (*UnimplementedGameEngineServer) SetScore(ctx context.Context, req *SetScoreRequest) (*SetScoreResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetScore not implemented")
}

func RegisterGameEngineServer(s *grpc.Server, srv GameEngineServer) {
	s.RegisterService(&_GameEngine_serviceDesc, srv)
}

func _GameEngine_GetSize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSizeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameEngineServer).GetSize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/m.gameengine.v1.GameEngine/GetSize",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameEngineServer).GetSize(ctx, req.(*GetSizeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameEngine_SetScore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetScoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameEngineServer).SetScore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/m.gameengine.v1.GameEngine/SetScore",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameEngineServer).SetScore(ctx, req.(*SetScoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GameEngine_serviceDesc = grpc.ServiceDesc{
	ServiceName: "m.gameengine.v1.GameEngine",
	HandlerType: (*GameEngineServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSize",
			Handler:    _GameEngine_GetSize_Handler,
		},
		{
			MethodName: "SetScore",
			Handler:    _GameEngine_SetScore_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "game-engine.proto",
}
