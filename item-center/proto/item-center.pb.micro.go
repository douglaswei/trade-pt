// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/item-center.proto

package itemcenter

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/asim/go-micro/v3/api"
	client "github.com/asim/go-micro/v3/client"
	server "github.com/asim/go-micro/v3/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for ItemCenter service

func NewItemCenterEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for ItemCenter service

type ItemCenterService interface {
	Call(ctx context.Context, in *CallRequest, opts ...client.CallOption) (*CallResponse, error)
	ClientStream(ctx context.Context, opts ...client.CallOption) (ItemCenter_ClientStreamService, error)
	ServerStream(ctx context.Context, in *ServerStreamRequest, opts ...client.CallOption) (ItemCenter_ServerStreamService, error)
	BidiStream(ctx context.Context, opts ...client.CallOption) (ItemCenter_BidiStreamService, error)
}

type itemCenterService struct {
	c    client.Client
	name string
}

func NewItemCenterService(name string, c client.Client) ItemCenterService {
	return &itemCenterService{
		c:    c,
		name: name,
	}
}

func (c *itemCenterService) Call(ctx context.Context, in *CallRequest, opts ...client.CallOption) (*CallResponse, error) {
	req := c.c.NewRequest(c.name, "ItemCenter.Call", in)
	out := new(CallResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemCenterService) ClientStream(ctx context.Context, opts ...client.CallOption) (ItemCenter_ClientStreamService, error) {
	req := c.c.NewRequest(c.name, "ItemCenter.ClientStream", &ClientStreamRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &itemCenterServiceClientStream{stream}, nil
}

type ItemCenter_ClientStreamService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*ClientStreamRequest) error
}

type itemCenterServiceClientStream struct {
	stream client.Stream
}

func (x *itemCenterServiceClientStream) Close() error {
	return x.stream.Close()
}

func (x *itemCenterServiceClientStream) Context() context.Context {
	return x.stream.Context()
}

func (x *itemCenterServiceClientStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *itemCenterServiceClientStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *itemCenterServiceClientStream) Send(m *ClientStreamRequest) error {
	return x.stream.Send(m)
}

func (c *itemCenterService) ServerStream(ctx context.Context, in *ServerStreamRequest, opts ...client.CallOption) (ItemCenter_ServerStreamService, error) {
	req := c.c.NewRequest(c.name, "ItemCenter.ServerStream", &ServerStreamRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &itemCenterServiceServerStream{stream}, nil
}

type ItemCenter_ServerStreamService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*ServerStreamResponse, error)
}

type itemCenterServiceServerStream struct {
	stream client.Stream
}

func (x *itemCenterServiceServerStream) Close() error {
	return x.stream.Close()
}

func (x *itemCenterServiceServerStream) Context() context.Context {
	return x.stream.Context()
}

func (x *itemCenterServiceServerStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *itemCenterServiceServerStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *itemCenterServiceServerStream) Recv() (*ServerStreamResponse, error) {
	m := new(ServerStreamResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *itemCenterService) BidiStream(ctx context.Context, opts ...client.CallOption) (ItemCenter_BidiStreamService, error) {
	req := c.c.NewRequest(c.name, "ItemCenter.BidiStream", &BidiStreamRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &itemCenterServiceBidiStream{stream}, nil
}

type ItemCenter_BidiStreamService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*BidiStreamRequest) error
	Recv() (*BidiStreamResponse, error)
}

type itemCenterServiceBidiStream struct {
	stream client.Stream
}

func (x *itemCenterServiceBidiStream) Close() error {
	return x.stream.Close()
}

func (x *itemCenterServiceBidiStream) Context() context.Context {
	return x.stream.Context()
}

func (x *itemCenterServiceBidiStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *itemCenterServiceBidiStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *itemCenterServiceBidiStream) Send(m *BidiStreamRequest) error {
	return x.stream.Send(m)
}

func (x *itemCenterServiceBidiStream) Recv() (*BidiStreamResponse, error) {
	m := new(BidiStreamResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for ItemCenter service

type ItemCenterHandler interface {
	Call(context.Context, *CallRequest, *CallResponse) error
	ClientStream(context.Context, ItemCenter_ClientStreamStream) error
	ServerStream(context.Context, *ServerStreamRequest, ItemCenter_ServerStreamStream) error
	BidiStream(context.Context, ItemCenter_BidiStreamStream) error
}

func RegisterItemCenterHandler(s server.Server, hdlr ItemCenterHandler, opts ...server.HandlerOption) error {
	type itemCenter interface {
		Call(ctx context.Context, in *CallRequest, out *CallResponse) error
		ClientStream(ctx context.Context, stream server.Stream) error
		ServerStream(ctx context.Context, stream server.Stream) error
		BidiStream(ctx context.Context, stream server.Stream) error
	}
	type ItemCenter struct {
		itemCenter
	}
	h := &itemCenterHandler{hdlr}
	return s.Handle(s.NewHandler(&ItemCenter{h}, opts...))
}

type itemCenterHandler struct {
	ItemCenterHandler
}

func (h *itemCenterHandler) Call(ctx context.Context, in *CallRequest, out *CallResponse) error {
	return h.ItemCenterHandler.Call(ctx, in, out)
}

func (h *itemCenterHandler) ClientStream(ctx context.Context, stream server.Stream) error {
	return h.ItemCenterHandler.ClientStream(ctx, &itemCenterClientStreamStream{stream})
}

type ItemCenter_ClientStreamStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*ClientStreamRequest, error)
}

type itemCenterClientStreamStream struct {
	stream server.Stream
}

func (x *itemCenterClientStreamStream) Close() error {
	return x.stream.Close()
}

func (x *itemCenterClientStreamStream) Context() context.Context {
	return x.stream.Context()
}

func (x *itemCenterClientStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *itemCenterClientStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *itemCenterClientStreamStream) Recv() (*ClientStreamRequest, error) {
	m := new(ClientStreamRequest)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (h *itemCenterHandler) ServerStream(ctx context.Context, stream server.Stream) error {
	m := new(ServerStreamRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.ItemCenterHandler.ServerStream(ctx, m, &itemCenterServerStreamStream{stream})
}

type ItemCenter_ServerStreamStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*ServerStreamResponse) error
}

type itemCenterServerStreamStream struct {
	stream server.Stream
}

func (x *itemCenterServerStreamStream) Close() error {
	return x.stream.Close()
}

func (x *itemCenterServerStreamStream) Context() context.Context {
	return x.stream.Context()
}

func (x *itemCenterServerStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *itemCenterServerStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *itemCenterServerStreamStream) Send(m *ServerStreamResponse) error {
	return x.stream.Send(m)
}

func (h *itemCenterHandler) BidiStream(ctx context.Context, stream server.Stream) error {
	return h.ItemCenterHandler.BidiStream(ctx, &itemCenterBidiStreamStream{stream})
}

type ItemCenter_BidiStreamStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*BidiStreamResponse) error
	Recv() (*BidiStreamRequest, error)
}

type itemCenterBidiStreamStream struct {
	stream server.Stream
}

func (x *itemCenterBidiStreamStream) Close() error {
	return x.stream.Close()
}

func (x *itemCenterBidiStreamStream) Context() context.Context {
	return x.stream.Context()
}

func (x *itemCenterBidiStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *itemCenterBidiStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *itemCenterBidiStreamStream) Send(m *BidiStreamResponse) error {
	return x.stream.Send(m)
}

func (x *itemCenterBidiStreamStream) Recv() (*BidiStreamRequest, error) {
	m := new(BidiStreamRequest)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}