// Code generated by protoc-gen-go.
// source: latency.proto
// DO NOT EDIT!

/*
Package latency is a generated protocol buffer package.

It is generated from these files:
	latency.proto

It has these top-level messages:
	Request
	Result
*/
package latency

import proto "github.com/golang/protobuf/proto"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal

type Request struct {
	User string `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
	Ip   string `protobuf:"bytes,2,opt,name=ip" json:"ip,omitempty"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}

type Result struct {
	Location string `protobuf:"bytes,1,opt,name=location" json:"location,omitempty"`
	Latency  uint32 `protobuf:"varint,2,opt,name=latency" json:"latency,omitempty"`
	User     string `protobuf:"bytes,3,opt,name=user" json:"user,omitempty"`
	Pinging  bool   `protobuf:"varint,4,opt,name=pinging" json:"pinging,omitempty"`
}

func (m *Result) Reset()         { *m = Result{} }
func (m *Result) String() string { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()    {}

// Client API for LatencyChecker service

type LatencyCheckerClient interface {
	Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Result, error)
}

type latencyCheckerClient struct {
	cc *grpc.ClientConn
}

func NewLatencyCheckerClient(cc *grpc.ClientConn) LatencyCheckerClient {
	return &latencyCheckerClient{cc}
}

func (c *latencyCheckerClient) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := grpc.Invoke(ctx, "/latency.LatencyChecker/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for LatencyChecker service

type LatencyCheckerServer interface {
	Ping(context.Context, *Request) (*Result, error)
}

func RegisterLatencyCheckerServer(s *grpc.Server, srv LatencyCheckerServer) {
	s.RegisterService(&_LatencyChecker_serviceDesc, srv)
}

func _LatencyChecker_Ping_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(Request)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(LatencyCheckerServer).Ping(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _LatencyChecker_serviceDesc = grpc.ServiceDesc{
	ServiceName: "latency.LatencyChecker",
	HandlerType: (*LatencyCheckerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _LatencyChecker_Ping_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}
