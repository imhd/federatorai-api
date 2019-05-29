// Code generated by protoc-gen-go. DO NOT EDIT.
// source: apiserver/rawdata/services.proto

package rawdata // import "github.com/containers-ai/federatorai-api/apiserver/rawdata"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import common "github.com/containers-ai/api/common"
import status "google.golang.org/genproto/googleapis/rpc/status"

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

type WriteRawdataRequest struct {
	Rawdata              []*common.Rawdata `protobuf:"bytes,1,rep,name=rawdata,proto3" json:"rawdata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *WriteRawdataRequest) Reset()         { *m = WriteRawdataRequest{} }
func (m *WriteRawdataRequest) String() string { return proto.CompactTextString(m) }
func (*WriteRawdataRequest) ProtoMessage()    {}
func (*WriteRawdataRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_services_733edc4eac0404a3, []int{0}
}
func (m *WriteRawdataRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WriteRawdataRequest.Unmarshal(m, b)
}
func (m *WriteRawdataRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WriteRawdataRequest.Marshal(b, m, deterministic)
}
func (dst *WriteRawdataRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WriteRawdataRequest.Merge(dst, src)
}
func (m *WriteRawdataRequest) XXX_Size() int {
	return xxx_messageInfo_WriteRawdataRequest.Size(m)
}
func (m *WriteRawdataRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WriteRawdataRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WriteRawdataRequest proto.InternalMessageInfo

func (m *WriteRawdataRequest) GetRawdata() []*common.Rawdata {
	if m != nil {
		return m.Rawdata
	}
	return nil
}

type ReadRawdataRequest struct {
	Queries              []*common.Query `protobuf:"bytes,1,rep,name=queries,proto3" json:"queries,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ReadRawdataRequest) Reset()         { *m = ReadRawdataRequest{} }
func (m *ReadRawdataRequest) String() string { return proto.CompactTextString(m) }
func (*ReadRawdataRequest) ProtoMessage()    {}
func (*ReadRawdataRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_services_733edc4eac0404a3, []int{1}
}
func (m *ReadRawdataRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadRawdataRequest.Unmarshal(m, b)
}
func (m *ReadRawdataRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadRawdataRequest.Marshal(b, m, deterministic)
}
func (dst *ReadRawdataRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadRawdataRequest.Merge(dst, src)
}
func (m *ReadRawdataRequest) XXX_Size() int {
	return xxx_messageInfo_ReadRawdataRequest.Size(m)
}
func (m *ReadRawdataRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadRawdataRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadRawdataRequest proto.InternalMessageInfo

func (m *ReadRawdataRequest) GetQueries() []*common.Query {
	if m != nil {
		return m.Queries
	}
	return nil
}

type ReadRawdataResponse struct {
	Rawdata              []*common.Rawdata `protobuf:"bytes,1,rep,name=rawdata,proto3" json:"rawdata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ReadRawdataResponse) Reset()         { *m = ReadRawdataResponse{} }
func (m *ReadRawdataResponse) String() string { return proto.CompactTextString(m) }
func (*ReadRawdataResponse) ProtoMessage()    {}
func (*ReadRawdataResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_services_733edc4eac0404a3, []int{2}
}
func (m *ReadRawdataResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadRawdataResponse.Unmarshal(m, b)
}
func (m *ReadRawdataResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadRawdataResponse.Marshal(b, m, deterministic)
}
func (dst *ReadRawdataResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadRawdataResponse.Merge(dst, src)
}
func (m *ReadRawdataResponse) XXX_Size() int {
	return xxx_messageInfo_ReadRawdataResponse.Size(m)
}
func (m *ReadRawdataResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadRawdataResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReadRawdataResponse proto.InternalMessageInfo

func (m *ReadRawdataResponse) GetRawdata() []*common.Rawdata {
	if m != nil {
		return m.Rawdata
	}
	return nil
}

func init() {
	proto.RegisterType((*WriteRawdataRequest)(nil), "containersai.apiserver.rawdata.WriteRawdataRequest")
	proto.RegisterType((*ReadRawdataRequest)(nil), "containersai.apiserver.rawdata.ReadRawdataRequest")
	proto.RegisterType((*ReadRawdataResponse)(nil), "containersai.apiserver.rawdata.ReadRawdataResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RawdataServiceClient is the client API for RawdataService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RawdataServiceClient interface {
	WriteRawdata(ctx context.Context, in *WriteRawdataRequest, opts ...grpc.CallOption) (*status.Status, error)
	ReadRawdata(ctx context.Context, in *ReadRawdataRequest, opts ...grpc.CallOption) (*ReadRawdataResponse, error)
}

type rawdataServiceClient struct {
	cc *grpc.ClientConn
}

func NewRawdataServiceClient(cc *grpc.ClientConn) RawdataServiceClient {
	return &rawdataServiceClient{cc}
}

func (c *rawdataServiceClient) WriteRawdata(ctx context.Context, in *WriteRawdataRequest, opts ...grpc.CallOption) (*status.Status, error) {
	out := new(status.Status)
	err := c.cc.Invoke(ctx, "/containersai.apiserver.rawdata.RawdataService/WriteRawdata", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rawdataServiceClient) ReadRawdata(ctx context.Context, in *ReadRawdataRequest, opts ...grpc.CallOption) (*ReadRawdataResponse, error) {
	out := new(ReadRawdataResponse)
	err := c.cc.Invoke(ctx, "/containersai.apiserver.rawdata.RawdataService/ReadRawdata", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RawdataServiceServer is the server API for RawdataService service.
type RawdataServiceServer interface {
	WriteRawdata(context.Context, *WriteRawdataRequest) (*status.Status, error)
	ReadRawdata(context.Context, *ReadRawdataRequest) (*ReadRawdataResponse, error)
}

func RegisterRawdataServiceServer(s *grpc.Server, srv RawdataServiceServer) {
	s.RegisterService(&_RawdataService_serviceDesc, srv)
}

func _RawdataService_WriteRawdata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WriteRawdataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RawdataServiceServer).WriteRawdata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/containersai.apiserver.rawdata.RawdataService/WriteRawdata",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RawdataServiceServer).WriteRawdata(ctx, req.(*WriteRawdataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RawdataService_ReadRawdata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadRawdataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RawdataServiceServer).ReadRawdata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/containersai.apiserver.rawdata.RawdataService/ReadRawdata",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RawdataServiceServer).ReadRawdata(ctx, req.(*ReadRawdataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RawdataService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "containersai.apiserver.rawdata.RawdataService",
	HandlerType: (*RawdataServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "WriteRawdata",
			Handler:    _RawdataService_WriteRawdata_Handler,
		},
		{
			MethodName: "ReadRawdata",
			Handler:    _RawdataService_ReadRawdata_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "apiserver/rawdata/services.proto",
}

func init() {
	proto.RegisterFile("apiserver/rawdata/services.proto", fileDescriptor_services_733edc4eac0404a3)
}

var fileDescriptor_services_733edc4eac0404a3 = []byte{
	// 298 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xcd, 0x4a, 0xc3, 0x40,
	0x14, 0x85, 0x5b, 0x04, 0x0b, 0x53, 0x71, 0x31, 0x2e, 0x94, 0x20, 0x52, 0x02, 0x82, 0x9b, 0xce,
	0x40, 0x23, 0x2e, 0xc4, 0x95, 0x4b, 0xc1, 0x85, 0xe9, 0x42, 0x74, 0x77, 0x9b, 0x5c, 0xe3, 0x05,
	0x9b, 0x99, 0xde, 0x99, 0xf8, 0xf3, 0xc0, 0xbe, 0x87, 0xa4, 0x99, 0x68, 0x63, 0x8b, 0xa2, 0xab,
	0x90, 0xe4, 0xcc, 0x97, 0x93, 0x8f, 0x23, 0x46, 0x60, 0xc9, 0x21, 0x3f, 0x23, 0x6b, 0x86, 0x97,
	0x1c, 0x3c, 0xe8, 0xfa, 0x96, 0x32, 0x74, 0xca, 0xb2, 0xf1, 0x46, 0x1e, 0x65, 0xa6, 0xf4, 0x40,
	0x25, 0xb2, 0x03, 0x52, 0x9f, 0x71, 0x15, 0xe2, 0xd1, 0xf1, 0xd7, 0xfb, 0x31, 0x90, 0x06, 0x4b,
	0x3a, 0x33, 0xf3, 0xb9, 0x29, 0xc3, 0xa5, 0xc1, 0x44, 0xfb, 0x85, 0x31, 0xc5, 0x13, 0x6a, 0xb6,
	0x99, 0x76, 0x1e, 0x7c, 0x15, 0xf8, 0xf1, 0xb5, 0xd8, 0xbb, 0x65, 0xf2, 0x98, 0x36, 0xbc, 0x14,
	0x17, 0x15, 0x3a, 0x2f, 0xcf, 0xc4, 0x20, 0x7c, 0xe1, 0xa0, 0x3f, 0xda, 0x3a, 0x19, 0x4e, 0x0e,
	0x55, 0xa7, 0x48, 0x80, 0xb7, 0xa7, 0xda, 0x70, 0x7c, 0x25, 0x64, 0x8a, 0x90, 0x7f, 0xa3, 0x9d,
	0x8a, 0xc1, 0xa2, 0x42, 0x26, 0x74, 0x81, 0x16, 0x6d, 0xa4, 0xdd, 0x54, 0xc8, 0x6f, 0x69, 0x1b,
	0xad, 0xab, 0x75, 0x58, 0xce, 0x9a, 0xd2, 0xe1, 0x7f, 0xab, 0x4d, 0xde, 0xfb, 0x62, 0x37, 0x3c,
	0x9c, 0x36, 0x8e, 0xe5, 0x9d, 0xd8, 0x59, 0xfd, 0x79, 0x99, 0xa8, 0x9f, 0x6d, 0xab, 0x0d, 0xaa,
	0x22, 0xa9, 0x1a, 0xb7, 0x8a, 0x6d, 0xa6, 0xa6, 0x4b, 0xb7, 0x71, 0x4f, 0xbe, 0x8a, 0xe1, 0x4a,
	0x79, 0x39, 0xf9, 0x8d, 0xbc, 0x6e, 0x2d, 0x4a, 0xfe, 0x74, 0xa6, 0xb1, 0x13, 0xf7, 0x2e, 0x2f,
	0xee, 0xcf, 0x0b, 0xf2, 0x8f, 0xd5, 0xac, 0x36, 0xa1, 0xbb, 0xf3, 0x78, 0xc0, 0x1c, 0x19, 0xbc,
	0x61, 0xa0, 0x71, 0x3d, 0x95, 0xb5, 0xf9, 0xcd, 0xb6, 0x97, 0xb3, 0x48, 0x3e, 0x02, 0x00, 0x00,
	0xff, 0xff, 0x5e, 0xe0, 0xaf, 0x31, 0x9a, 0x02, 0x00, 0x00,
}
