// Code generated by protoc-gen-go. DO NOT EDIT.
// source: stream.proto

package proto

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

type StreamRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Age                  int64    `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamRequest) Reset()         { *m = StreamRequest{} }
func (m *StreamRequest) String() string { return proto.CompactTextString(m) }
func (*StreamRequest) ProtoMessage()    {}
func (*StreamRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb17ef3f514bfe54, []int{0}
}

func (m *StreamRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamRequest.Unmarshal(m, b)
}
func (m *StreamRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamRequest.Marshal(b, m, deterministic)
}
func (m *StreamRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamRequest.Merge(m, src)
}
func (m *StreamRequest) XXX_Size() int {
	return xxx_messageInfo_StreamRequest.Size(m)
}
func (m *StreamRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StreamRequest proto.InternalMessageInfo

func (m *StreamRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *StreamRequest) GetAge() int64 {
	if m != nil {
		return m.Age
	}
	return 0
}

type StreamResponse struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Age                  int64    `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamResponse) Reset()         { *m = StreamResponse{} }
func (m *StreamResponse) String() string { return proto.CompactTextString(m) }
func (*StreamResponse) ProtoMessage()    {}
func (*StreamResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb17ef3f514bfe54, []int{1}
}

func (m *StreamResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamResponse.Unmarshal(m, b)
}
func (m *StreamResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamResponse.Marshal(b, m, deterministic)
}
func (m *StreamResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamResponse.Merge(m, src)
}
func (m *StreamResponse) XXX_Size() int {
	return xxx_messageInfo_StreamResponse.Size(m)
}
func (m *StreamResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StreamResponse proto.InternalMessageInfo

func (m *StreamResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *StreamResponse) GetAge() int64 {
	if m != nil {
		return m.Age
	}
	return 0
}

func init() {
	proto.RegisterType((*StreamRequest)(nil), "proto.StreamRequest")
	proto.RegisterType((*StreamResponse)(nil), "proto.StreamResponse")
}

func init() { proto.RegisterFile("stream.proto", fileDescriptor_bb17ef3f514bfe54) }

var fileDescriptor_bb17ef3f514bfe54 = []byte{
	// 176 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x2e, 0x29, 0x4a,
	0x4d, 0xcc, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a, 0xa6, 0x5c, 0xbc,
	0xc1, 0x60, 0xe1, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x21, 0x2e, 0x96, 0xbc, 0xc4,
	0xdc, 0x54, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x30, 0x5b, 0x48, 0x80, 0x8b, 0x39, 0x31,
	0x3d, 0x55, 0x82, 0x49, 0x81, 0x51, 0x83, 0x39, 0x08, 0xc4, 0x54, 0x32, 0xe3, 0xe2, 0x83, 0x69,
	0x2b, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x25, 0x4e, 0x9f, 0xd1, 0x39, 0x46, 0x98, 0x7d, 0xc1, 0xa9,
	0x45, 0x65, 0x99, 0xc9, 0xa9, 0x42, 0xb6, 0x5c, 0x5c, 0x4e, 0x89, 0xc5, 0xa9, 0x20, 0x6e, 0x6a,
	0x91, 0x90, 0x08, 0xc4, 0x75, 0x7a, 0x28, 0x6e, 0x92, 0x12, 0x45, 0x13, 0x85, 0x58, 0xa9, 0xc4,
	0x60, 0xc0, 0x08, 0xd3, 0xee, 0x9c, 0x93, 0x99, 0x9a, 0x57, 0x42, 0xa2, 0x76, 0x0d, 0x46, 0x21,
	0x7b, 0x88, 0x76, 0x97, 0xfc, 0xd2, 0xa4, 0x9c, 0x54, 0x92, 0xb5, 0x1b, 0x30, 0x26, 0xb1, 0x81,
	0xe5, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x89, 0xeb, 0x83, 0x98, 0x5d, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// StreamServiceClient is the client API for StreamService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StreamServiceClient interface {
	BaseServer(ctx context.Context, in *StreamRequest, opts ...grpc.CallOption) (StreamService_BaseServerClient, error)
	BaseClient(ctx context.Context, opts ...grpc.CallOption) (StreamService_BaseClientClient, error)
	BaseDouble(ctx context.Context, opts ...grpc.CallOption) (StreamService_BaseDoubleClient, error)
}

type streamServiceClient struct {
	cc *grpc.ClientConn
}

func NewStreamServiceClient(cc *grpc.ClientConn) StreamServiceClient {
	return &streamServiceClient{cc}
}

func (c *streamServiceClient) BaseServer(ctx context.Context, in *StreamRequest, opts ...grpc.CallOption) (StreamService_BaseServerClient, error) {
	stream, err := c.cc.NewStream(ctx, &_StreamService_serviceDesc.Streams[0], "/proto.StreamService/BaseServer", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamServiceBaseServerClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type StreamService_BaseServerClient interface {
	Recv() (*StreamResponse, error)
	grpc.ClientStream
}

type streamServiceBaseServerClient struct {
	grpc.ClientStream
}

func (x *streamServiceBaseServerClient) Recv() (*StreamResponse, error) {
	m := new(StreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *streamServiceClient) BaseClient(ctx context.Context, opts ...grpc.CallOption) (StreamService_BaseClientClient, error) {
	stream, err := c.cc.NewStream(ctx, &_StreamService_serviceDesc.Streams[1], "/proto.StreamService/BaseClient", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamServiceBaseClientClient{stream}
	return x, nil
}

type StreamService_BaseClientClient interface {
	Send(*StreamRequest) error
	CloseAndRecv() (*StreamResponse, error)
	grpc.ClientStream
}

type streamServiceBaseClientClient struct {
	grpc.ClientStream
}

func (x *streamServiceBaseClientClient) Send(m *StreamRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *streamServiceBaseClientClient) CloseAndRecv() (*StreamResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(StreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *streamServiceClient) BaseDouble(ctx context.Context, opts ...grpc.CallOption) (StreamService_BaseDoubleClient, error) {
	stream, err := c.cc.NewStream(ctx, &_StreamService_serviceDesc.Streams[2], "/proto.StreamService/BaseDouble", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamServiceBaseDoubleClient{stream}
	return x, nil
}

type StreamService_BaseDoubleClient interface {
	Send(*StreamRequest) error
	Recv() (*StreamResponse, error)
	grpc.ClientStream
}

type streamServiceBaseDoubleClient struct {
	grpc.ClientStream
}

func (x *streamServiceBaseDoubleClient) Send(m *StreamRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *streamServiceBaseDoubleClient) Recv() (*StreamResponse, error) {
	m := new(StreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StreamServiceServer is the server API for StreamService service.
type StreamServiceServer interface {
	BaseServer(*StreamRequest, StreamService_BaseServerServer) error
	BaseClient(StreamService_BaseClientServer) error
	BaseDouble(StreamService_BaseDoubleServer) error
}

// UnimplementedStreamServiceServer can be embedded to have forward compatible implementations.
type UnimplementedStreamServiceServer struct {
}

func (*UnimplementedStreamServiceServer) BaseServer(req *StreamRequest, srv StreamService_BaseServerServer) error {
	return status.Errorf(codes.Unimplemented, "method BaseServer not implemented")
}
func (*UnimplementedStreamServiceServer) BaseClient(srv StreamService_BaseClientServer) error {
	return status.Errorf(codes.Unimplemented, "method BaseClient not implemented")
}
func (*UnimplementedStreamServiceServer) BaseDouble(srv StreamService_BaseDoubleServer) error {
	return status.Errorf(codes.Unimplemented, "method BaseDouble not implemented")
}

func RegisterStreamServiceServer(s *grpc.Server, srv StreamServiceServer) {
	s.RegisterService(&_StreamService_serviceDesc, srv)
}

func _StreamService_BaseServer_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StreamServiceServer).BaseServer(m, &streamServiceBaseServerServer{stream})
}

type StreamService_BaseServerServer interface {
	Send(*StreamResponse) error
	grpc.ServerStream
}

type streamServiceBaseServerServer struct {
	grpc.ServerStream
}

func (x *streamServiceBaseServerServer) Send(m *StreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _StreamService_BaseClient_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StreamServiceServer).BaseClient(&streamServiceBaseClientServer{stream})
}

type StreamService_BaseClientServer interface {
	SendAndClose(*StreamResponse) error
	Recv() (*StreamRequest, error)
	grpc.ServerStream
}

type streamServiceBaseClientServer struct {
	grpc.ServerStream
}

func (x *streamServiceBaseClientServer) SendAndClose(m *StreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *streamServiceBaseClientServer) Recv() (*StreamRequest, error) {
	m := new(StreamRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _StreamService_BaseDouble_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StreamServiceServer).BaseDouble(&streamServiceBaseDoubleServer{stream})
}

type StreamService_BaseDoubleServer interface {
	Send(*StreamResponse) error
	Recv() (*StreamRequest, error)
	grpc.ServerStream
}

type streamServiceBaseDoubleServer struct {
	grpc.ServerStream
}

func (x *streamServiceBaseDoubleServer) Send(m *StreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *streamServiceBaseDoubleServer) Recv() (*StreamRequest, error) {
	m := new(StreamRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _StreamService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.StreamService",
	HandlerType: (*StreamServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "BaseServer",
			Handler:       _StreamService_BaseServer_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "BaseClient",
			Handler:       _StreamService_BaseClient_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "BaseDouble",
			Handler:       _StreamService_BaseDouble_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "stream.proto",
}
