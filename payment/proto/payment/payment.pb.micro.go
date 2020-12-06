// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/payment/payment.proto

package micro_super_lightning_service_payment

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

// Api Endpoints for Payment service

func NewPaymentEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Payment service

type PaymentService interface {
	Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (Payment_StreamService, error)
	PingPong(ctx context.Context, opts ...client.CallOption) (Payment_PingPongService, error)
}

type paymentService struct {
	c    client.Client
	name string
}

func NewPaymentService(name string, c client.Client) PaymentService {
	return &paymentService{
		c:    c,
		name: name,
	}
}

func (c *paymentService) Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Payment.Call", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentService) Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (Payment_StreamService, error) {
	req := c.c.NewRequest(c.name, "Payment.Stream", &StreamingRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &paymentServiceStream{stream}, nil
}

type Payment_StreamService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*StreamingResponse, error)
}

type paymentServiceStream struct {
	stream client.Stream
}

func (x *paymentServiceStream) Close() error {
	return x.stream.Close()
}

func (x *paymentServiceStream) Context() context.Context {
	return x.stream.Context()
}

func (x *paymentServiceStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *paymentServiceStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *paymentServiceStream) Recv() (*StreamingResponse, error) {
	m := new(StreamingResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *paymentService) PingPong(ctx context.Context, opts ...client.CallOption) (Payment_PingPongService, error) {
	req := c.c.NewRequest(c.name, "Payment.PingPong", &Ping{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &paymentServicePingPong{stream}, nil
}

type Payment_PingPongService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Ping) error
	Recv() (*Pong, error)
}

type paymentServicePingPong struct {
	stream client.Stream
}

func (x *paymentServicePingPong) Close() error {
	return x.stream.Close()
}

func (x *paymentServicePingPong) Context() context.Context {
	return x.stream.Context()
}

func (x *paymentServicePingPong) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *paymentServicePingPong) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *paymentServicePingPong) Send(m *Ping) error {
	return x.stream.Send(m)
}

func (x *paymentServicePingPong) Recv() (*Pong, error) {
	m := new(Pong)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Payment service

type PaymentHandler interface {
	Call(context.Context, *Request, *Response) error
	Stream(context.Context, *StreamingRequest, Payment_StreamStream) error
	PingPong(context.Context, Payment_PingPongStream) error
}

func RegisterPaymentHandler(s server.Server, hdlr PaymentHandler, opts ...server.HandlerOption) error {
	type payment interface {
		Call(ctx context.Context, in *Request, out *Response) error
		Stream(ctx context.Context, stream server.Stream) error
		PingPong(ctx context.Context, stream server.Stream) error
	}
	type Payment struct {
		payment
	}
	h := &paymentHandler{hdlr}
	return s.Handle(s.NewHandler(&Payment{h}, opts...))
}

type paymentHandler struct {
	PaymentHandler
}

func (h *paymentHandler) Call(ctx context.Context, in *Request, out *Response) error {
	return h.PaymentHandler.Call(ctx, in, out)
}

func (h *paymentHandler) Stream(ctx context.Context, stream server.Stream) error {
	m := new(StreamingRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.PaymentHandler.Stream(ctx, m, &paymentStreamStream{stream})
}

type Payment_StreamStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*StreamingResponse) error
}

type paymentStreamStream struct {
	stream server.Stream
}

func (x *paymentStreamStream) Close() error {
	return x.stream.Close()
}

func (x *paymentStreamStream) Context() context.Context {
	return x.stream.Context()
}

func (x *paymentStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *paymentStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *paymentStreamStream) Send(m *StreamingResponse) error {
	return x.stream.Send(m)
}

func (h *paymentHandler) PingPong(ctx context.Context, stream server.Stream) error {
	return h.PaymentHandler.PingPong(ctx, &paymentPingPongStream{stream})
}

type Payment_PingPongStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Pong) error
	Recv() (*Ping, error)
}

type paymentPingPongStream struct {
	stream server.Stream
}

func (x *paymentPingPongStream) Close() error {
	return x.stream.Close()
}

func (x *paymentPingPongStream) Context() context.Context {
	return x.stream.Context()
}

func (x *paymentPingPongStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *paymentPingPongStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *paymentPingPongStream) Send(m *Pong) error {
	return x.stream.Send(m)
}

func (x *paymentPingPongStream) Recv() (*Ping, error) {
	m := new(Ping)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}
