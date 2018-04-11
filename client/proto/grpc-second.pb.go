// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/grpc-second.proto

package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type HelloRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *HelloRequest) Reset()                    { *m = HelloRequest{} }
func (m *HelloRequest) String() string            { return proto1.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()               {}
func (*HelloRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *HelloRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type HelloReply struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *HelloReply) Reset()                    { *m = HelloReply{} }
func (m *HelloReply) String() string            { return proto1.CompactTextString(m) }
func (*HelloReply) ProtoMessage()               {}
func (*HelloReply) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *HelloReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto1.RegisterType((*HelloRequest)(nil), "proto.HelloRequest")
	proto1.RegisterType((*HelloReply)(nil), "proto.HelloReply")
}

func init() { proto1.RegisterFile("proto/grpc-second.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 110 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2f, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x4f, 0x2f, 0x2a, 0x48, 0xd6, 0x2d, 0x4e, 0x4d, 0xce, 0xcf, 0x4b, 0xd1, 0x03, 0x8b,
	0x08, 0xb1, 0x82, 0x29, 0x25, 0x25, 0x2e, 0x1e, 0x8f, 0xd4, 0x9c, 0x9c, 0xfc, 0xa0, 0xd4, 0xc2,
	0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x21, 0x2e, 0x96, 0xbc, 0xc4, 0xdc, 0x54, 0x09, 0x46, 0x05, 0x46,
	0x0d, 0xce, 0x20, 0x30, 0x5b, 0x49, 0x8d, 0x8b, 0x0b, 0xaa, 0xa6, 0x20, 0xa7, 0x52, 0x48, 0x82,
	0x8b, 0x3d, 0x37, 0xb5, 0xb8, 0x38, 0x31, 0x1d, 0xa6, 0x08, 0xc6, 0x4d, 0x62, 0x03, 0x1b, 0x69,
	0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xa5, 0x41, 0x00, 0xb1, 0x74, 0x00, 0x00, 0x00,
}
