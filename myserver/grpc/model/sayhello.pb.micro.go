// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: sayhello.proto

package protoes

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for HelloServer service

func NewHelloServerEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for HelloServer service

type HelloServerService interface {
	SayHi(ctx context.Context, in *HelloRequest, opts ...client.CallOption) (*HelloReplay, error)
	GetMsg(ctx context.Context, in *HelloRequest, opts ...client.CallOption) (*HelloMessage, error)
}

type helloServerService struct {
	c    client.Client
	name string
}

func NewHelloServerService(name string, c client.Client) HelloServerService {
	return &helloServerService{
		c:    c,
		name: name,
	}
}

func (c *helloServerService) SayHi(ctx context.Context, in *HelloRequest, opts ...client.CallOption) (*HelloReplay, error) {
	req := c.c.NewRequest(c.name, "HelloServer.SayHi", in)
	out := new(HelloReplay)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloServerService) GetMsg(ctx context.Context, in *HelloRequest, opts ...client.CallOption) (*HelloMessage, error) {
	req := c.c.NewRequest(c.name, "HelloServer.GetMsg", in)
	out := new(HelloMessage)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for HelloServer service

type HelloServerHandler interface {
	SayHi(context.Context, *HelloRequest, *HelloReplay) error
	GetMsg(context.Context, *HelloRequest, *HelloMessage) error
}

func RegisterHelloServerHandler(s server.Server, hdlr HelloServerHandler, opts ...server.HandlerOption) error {
	type helloServer interface {
		SayHi(ctx context.Context, in *HelloRequest, out *HelloReplay) error
		GetMsg(ctx context.Context, in *HelloRequest, out *HelloMessage) error
	}
	type HelloServer struct {
		helloServer
	}
	h := &helloServerHandler{hdlr}
	return s.Handle(s.NewHandler(&HelloServer{h}, opts...))
}

type helloServerHandler struct {
	HelloServerHandler
}

func (h *helloServerHandler) SayHi(ctx context.Context, in *HelloRequest, out *HelloReplay) error {
	return h.HelloServerHandler.SayHi(ctx, in, out)
}

func (h *helloServerHandler) GetMsg(ctx context.Context, in *HelloRequest, out *HelloMessage) error {
	return h.HelloServerHandler.GetMsg(ctx, in, out)
}
