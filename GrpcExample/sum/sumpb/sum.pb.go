// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sum/sumpb/sum.proto

package sumpb

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

type SumRequest struct {
	A                    int32    `protobuf:"varint,1,opt,name=a,proto3" json:"a,omitempty"`
	B                    int32    `protobuf:"varint,2,opt,name=b,proto3" json:"b,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SumRequest) Reset()         { *m = SumRequest{} }
func (m *SumRequest) String() string { return proto.CompactTextString(m) }
func (*SumRequest) ProtoMessage()    {}
func (*SumRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e82cd6e6acf87a3, []int{0}
}

func (m *SumRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SumRequest.Unmarshal(m, b)
}
func (m *SumRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SumRequest.Marshal(b, m, deterministic)
}
func (m *SumRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SumRequest.Merge(m, src)
}
func (m *SumRequest) XXX_Size() int {
	return xxx_messageInfo_SumRequest.Size(m)
}
func (m *SumRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SumRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SumRequest proto.InternalMessageInfo

func (m *SumRequest) GetA() int32 {
	if m != nil {
		return m.A
	}
	return 0
}

func (m *SumRequest) GetB() int32 {
	if m != nil {
		return m.B
	}
	return 0
}

type SumResponse struct {
	Sumresult            int32    `protobuf:"varint,1,opt,name=sumresult,proto3" json:"sumresult,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SumResponse) Reset()         { *m = SumResponse{} }
func (m *SumResponse) String() string { return proto.CompactTextString(m) }
func (*SumResponse) ProtoMessage()    {}
func (*SumResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e82cd6e6acf87a3, []int{1}
}

func (m *SumResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SumResponse.Unmarshal(m, b)
}
func (m *SumResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SumResponse.Marshal(b, m, deterministic)
}
func (m *SumResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SumResponse.Merge(m, src)
}
func (m *SumResponse) XXX_Size() int {
	return xxx_messageInfo_SumResponse.Size(m)
}
func (m *SumResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SumResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SumResponse proto.InternalMessageInfo

func (m *SumResponse) GetSumresult() int32 {
	if m != nil {
		return m.Sumresult
	}
	return 0
}

func init() {
	proto.RegisterType((*SumRequest)(nil), "sum.SumRequest")
	proto.RegisterType((*SumResponse)(nil), "sum.SumResponse")
}

func init() { proto.RegisterFile("sum/sumpb/sum.proto", fileDescriptor_1e82cd6e6acf87a3) }

var fileDescriptor_1e82cd6e6acf87a3 = []byte{
	// 158 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0x2e, 0xcd, 0xd5,
	0x2f, 0x2e, 0xcd, 0x2d, 0x48, 0x02, 0x91, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0xcc, 0xc5,
	0xa5, 0xb9, 0x4a, 0x1a, 0x5c, 0x5c, 0xc1, 0xa5, 0xb9, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25,
	0x42, 0x3c, 0x5c, 0x8c, 0x89, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xac, 0x41, 0x8c, 0x89, 0x20, 0x5e,
	0x92, 0x04, 0x13, 0x84, 0x97, 0xa4, 0xa4, 0xcd, 0xc5, 0x0d, 0x56, 0x59, 0x5c, 0x90, 0x9f, 0x57,
	0x9c, 0x2a, 0x24, 0xc3, 0xc5, 0x59, 0x5c, 0x9a, 0x5b, 0x94, 0x5a, 0x5c, 0x9a, 0x53, 0x02, 0xd5,
	0x82, 0x10, 0x30, 0x32, 0xe3, 0xe2, 0x70, 0x4c, 0x49, 0xc9, 0x2c, 0xc9, 0xcc, 0xcf, 0x13, 0xd2,
	0xe2, 0x62, 0x0e, 0x2e, 0xcd, 0x15, 0xe2, 0xd7, 0x03, 0x59, 0x8d, 0xb0, 0x4c, 0x4a, 0x00, 0x21,
	0x00, 0x31, 0x53, 0x89, 0xc1, 0x89, 0x3d, 0x8a, 0x15, 0xec, 0xcc, 0x24, 0x36, 0xb0, 0x1b, 0x8d,
	0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xe5, 0xa7, 0xa2, 0x7b, 0xba, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AdditionClient is the client API for Addition service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AdditionClient interface {
	Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error)
}

type additionClient struct {
	cc *grpc.ClientConn
}

func NewAdditionClient(cc *grpc.ClientConn) AdditionClient {
	return &additionClient{cc}
}

func (c *additionClient) Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error) {
	out := new(SumResponse)
	err := c.cc.Invoke(ctx, "/sum.Addition/Sum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdditionServer is the server API for Addition service.
type AdditionServer interface {
	Sum(context.Context, *SumRequest) (*SumResponse, error)
}

// UnimplementedAdditionServer can be embedded to have forward compatible implementations.
type UnimplementedAdditionServer struct {
}

func (*UnimplementedAdditionServer) Sum(ctx context.Context, req *SumRequest) (*SumResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sum not implemented")
}

func RegisterAdditionServer(s *grpc.Server, srv AdditionServer) {
	s.RegisterService(&_Addition_serviceDesc, srv)
}

func _Addition_Sum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdditionServer).Sum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sum.Addition/Sum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdditionServer).Sum(ctx, req.(*SumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Addition_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sum.Addition",
	HandlerType: (*AdditionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sum",
			Handler:    _Addition_Sum_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sum/sumpb/sum.proto",
}
