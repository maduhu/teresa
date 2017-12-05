// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pkg/protobuf/resource/resource.proto

/*
Package resource is a generated protocol buffer package.

It is generated from these files:
	pkg/protobuf/resource/resource.proto

It has these top-level messages:
	CreateRequest
	CreateResponse
	DeleteRequest
	Empty
*/
package resource

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

type CreateRequest struct {
	Name     string                   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	TeamName string                   `protobuf:"bytes,2,opt,name=team_name,json=teamName" json:"team_name,omitempty"`
	Settings []*CreateRequest_Setting `protobuf:"bytes,3,rep,name=settings" json:"settings,omitempty"`
}

func (m *CreateRequest) Reset()                    { *m = CreateRequest{} }
func (m *CreateRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()               {}
func (*CreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CreateRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateRequest) GetTeamName() string {
	if m != nil {
		return m.TeamName
	}
	return ""
}

func (m *CreateRequest) GetSettings() []*CreateRequest_Setting {
	if m != nil {
		return m.Settings
	}
	return nil
}

type CreateRequest_Setting struct {
	Key   string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *CreateRequest_Setting) Reset()                    { *m = CreateRequest_Setting{} }
func (m *CreateRequest_Setting) String() string            { return proto.CompactTextString(m) }
func (*CreateRequest_Setting) ProtoMessage()               {}
func (*CreateRequest_Setting) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func (m *CreateRequest_Setting) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *CreateRequest_Setting) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type CreateResponse struct {
	Text string `protobuf:"bytes,1,opt,name=text" json:"text,omitempty"`
}

func (m *CreateResponse) Reset()                    { *m = CreateResponse{} }
func (m *CreateResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()               {}
func (*CreateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CreateResponse) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

type DeleteRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *DeleteRequest) Reset()                    { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()               {}
func (*DeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *DeleteRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func init() {
	proto.RegisterType((*CreateRequest)(nil), "resource.CreateRequest")
	proto.RegisterType((*CreateRequest_Setting)(nil), "resource.CreateRequest.Setting")
	proto.RegisterType((*CreateResponse)(nil), "resource.CreateResponse")
	proto.RegisterType((*DeleteRequest)(nil), "resource.DeleteRequest")
	proto.RegisterType((*Empty)(nil), "resource.Empty")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Resource service

type ResourceClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*Empty, error)
}

type resourceClient struct {
	cc *grpc.ClientConn
}

func NewResourceClient(cc *grpc.ClientConn) ResourceClient {
	return &resourceClient{cc}
}

func (c *resourceClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := grpc.Invoke(ctx, "/resource.Resource/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/resource.Resource/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Resource service

type ResourceServer interface {
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	Delete(context.Context, *DeleteRequest) (*Empty, error)
}

func RegisterResourceServer(s *grpc.Server, srv ResourceServer) {
	s.RegisterService(&_Resource_serviceDesc, srv)
}

func _Resource_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/resource.Resource/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Resource_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/resource.Resource/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Resource_serviceDesc = grpc.ServiceDesc{
	ServiceName: "resource.Resource",
	HandlerType: (*ResourceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Resource_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Resource_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/protobuf/resource/resource.proto",
}

func init() { proto.RegisterFile("pkg/protobuf/resource/resource.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 256 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0xc6, 0x89, 0xb1, 0x69, 0x3a, 0x52, 0x95, 0x41, 0x30, 0xc4, 0x83, 0x25, 0xf6, 0xd0, 0x53,
	0x8a, 0xf1, 0xd8, 0xa3, 0x7a, 0xf5, 0x10, 0x1f, 0x40, 0x52, 0x19, 0x8b, 0xb4, 0xf9, 0xe3, 0xee,
	0xac, 0x58, 0x7c, 0x29, 0x1f, 0x51, 0xf6, 0x5f, 0xb4, 0x88, 0xbd, 0x7d, 0x33, 0xbf, 0x9d, 0xf9,
	0xbe, 0x61, 0x61, 0xda, 0xad, 0x57, 0xf3, 0x4e, 0xb4, 0xdc, 0x2e, 0xd5, 0xcb, 0x5c, 0x90, 0x6c,
	0x95, 0x78, 0xa6, 0x5e, 0xe4, 0x06, 0x61, 0xec, 0xeb, 0xec, 0x2b, 0x80, 0xf1, 0xad, 0xa0, 0x8a,
	0xa9, 0xa4, 0x37, 0x45, 0x92, 0x11, 0xe1, 0xb0, 0xa9, 0x6a, 0x4a, 0x82, 0x49, 0x30, 0x1b, 0x95,
	0x46, 0xe3, 0x05, 0x8c, 0x98, 0xaa, 0xfa, 0xc9, 0x80, 0x03, 0x03, 0x62, 0xdd, 0x78, 0xd0, 0x70,
	0x01, 0xb1, 0x24, 0xe6, 0xd7, 0x66, 0x25, 0x93, 0x70, 0x12, 0xce, 0x8e, 0x8a, 0xcb, 0xbc, 0xf7,
	0xdb, 0xd9, 0x9d, 0x3f, 0xda, 0x77, 0x65, 0x3f, 0x90, 0x5e, 0xc3, 0xd0, 0x35, 0xf1, 0x14, 0xc2,
	0x35, 0x6d, 0x9d, 0xaf, 0x96, 0x78, 0x06, 0x83, 0xf7, 0x6a, 0xa3, 0xbc, 0xa5, 0x2d, 0xb2, 0x29,
	0x1c, 0xfb, 0xad, 0xb2, 0x6b, 0x1b, 0x49, 0x3a, 0x32, 0xd3, 0x07, 0xfb, 0xc8, 0x5a, 0x67, 0x57,
	0x30, 0xbe, 0xa3, 0x0d, 0xed, 0xbd, 0x2b, 0x1b, 0xc2, 0xe0, 0xbe, 0xee, 0x78, 0x5b, 0x7c, 0x42,
	0x5c, 0xba, 0xc8, 0xb8, 0x80, 0xc8, 0xee, 0xc7, 0xf3, 0x7f, 0xee, 0x48, 0x93, 0xbf, 0xc0, 0x45,
	0x29, 0x20, 0xb2, 0xb6, 0xbf, 0x87, 0x77, 0x82, 0xa4, 0x27, 0x3f, 0xc0, 0x98, 0x2f, 0x23, 0xf3,
	0x29, 0x37, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x62, 0x6c, 0xf3, 0x95, 0xbc, 0x01, 0x00, 0x00,
}
