// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/album/panorama.proto

package album

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
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
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for PanoramaService service

type PanoramaService interface {
	AddOne(ctx context.Context, in *ReqPanoramaAdd, opts ...client.CallOption) (*ReplyPanoramaInfo, error)
	UpdateBase(ctx context.Context, in *ReqPanoramaUpdate, opts ...client.CallOption) (*ReplyInfo, error)
	RemoveOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyInfo, error)
	GetOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyPanoramaInfo, error)
	Search(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyPanoramaList, error)
	GetListByFilter(ctx context.Context, in *RequestFilter, opts ...client.CallOption) (*ReplyPanoramaList, error)
	GetStatistic(ctx context.Context, in *RequestFilter, opts ...client.CallOption) (*ReplyStatistic, error)
	UpdateByFilter(ctx context.Context, in *RequestUpdate, opts ...client.CallOption) (*ReplyInfo, error)
}

type panoramaService struct {
	c    client.Client
	name string
}

func NewPanoramaService(name string, c client.Client) PanoramaService {
	return &panoramaService{
		c:    c,
		name: name,
	}
}

func (c *panoramaService) AddOne(ctx context.Context, in *ReqPanoramaAdd, opts ...client.CallOption) (*ReplyPanoramaInfo, error) {
	req := c.c.NewRequest(c.name, "PanoramaService.AddOne", in)
	out := new(ReplyPanoramaInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *panoramaService) UpdateBase(ctx context.Context, in *ReqPanoramaUpdate, opts ...client.CallOption) (*ReplyInfo, error) {
	req := c.c.NewRequest(c.name, "PanoramaService.UpdateBase", in)
	out := new(ReplyInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *panoramaService) RemoveOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyInfo, error) {
	req := c.c.NewRequest(c.name, "PanoramaService.RemoveOne", in)
	out := new(ReplyInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *panoramaService) GetOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyPanoramaInfo, error) {
	req := c.c.NewRequest(c.name, "PanoramaService.GetOne", in)
	out := new(ReplyPanoramaInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *panoramaService) Search(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyPanoramaList, error) {
	req := c.c.NewRequest(c.name, "PanoramaService.Search", in)
	out := new(ReplyPanoramaList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *panoramaService) GetListByFilter(ctx context.Context, in *RequestFilter, opts ...client.CallOption) (*ReplyPanoramaList, error) {
	req := c.c.NewRequest(c.name, "PanoramaService.GetListByFilter", in)
	out := new(ReplyPanoramaList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *panoramaService) GetStatistic(ctx context.Context, in *RequestFilter, opts ...client.CallOption) (*ReplyStatistic, error) {
	req := c.c.NewRequest(c.name, "PanoramaService.GetStatistic", in)
	out := new(ReplyStatistic)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *panoramaService) UpdateByFilter(ctx context.Context, in *RequestUpdate, opts ...client.CallOption) (*ReplyInfo, error) {
	req := c.c.NewRequest(c.name, "PanoramaService.UpdateByFilter", in)
	out := new(ReplyInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PanoramaService service

type PanoramaServiceHandler interface {
	AddOne(context.Context, *ReqPanoramaAdd, *ReplyPanoramaInfo) error
	UpdateBase(context.Context, *ReqPanoramaUpdate, *ReplyInfo) error
	RemoveOne(context.Context, *RequestInfo, *ReplyInfo) error
	GetOne(context.Context, *RequestInfo, *ReplyPanoramaInfo) error
	Search(context.Context, *RequestInfo, *ReplyPanoramaList) error
	GetListByFilter(context.Context, *RequestFilter, *ReplyPanoramaList) error
	GetStatistic(context.Context, *RequestFilter, *ReplyStatistic) error
	UpdateByFilter(context.Context, *RequestUpdate, *ReplyInfo) error
}

func RegisterPanoramaServiceHandler(s server.Server, hdlr PanoramaServiceHandler, opts ...server.HandlerOption) error {
	type panoramaService interface {
		AddOne(ctx context.Context, in *ReqPanoramaAdd, out *ReplyPanoramaInfo) error
		UpdateBase(ctx context.Context, in *ReqPanoramaUpdate, out *ReplyInfo) error
		RemoveOne(ctx context.Context, in *RequestInfo, out *ReplyInfo) error
		GetOne(ctx context.Context, in *RequestInfo, out *ReplyPanoramaInfo) error
		Search(ctx context.Context, in *RequestInfo, out *ReplyPanoramaList) error
		GetListByFilter(ctx context.Context, in *RequestFilter, out *ReplyPanoramaList) error
		GetStatistic(ctx context.Context, in *RequestFilter, out *ReplyStatistic) error
		UpdateByFilter(ctx context.Context, in *RequestUpdate, out *ReplyInfo) error
	}
	type PanoramaService struct {
		panoramaService
	}
	h := &panoramaServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&PanoramaService{h}, opts...))
}

type panoramaServiceHandler struct {
	PanoramaServiceHandler
}

func (h *panoramaServiceHandler) AddOne(ctx context.Context, in *ReqPanoramaAdd, out *ReplyPanoramaInfo) error {
	return h.PanoramaServiceHandler.AddOne(ctx, in, out)
}

func (h *panoramaServiceHandler) UpdateBase(ctx context.Context, in *ReqPanoramaUpdate, out *ReplyInfo) error {
	return h.PanoramaServiceHandler.UpdateBase(ctx, in, out)
}

func (h *panoramaServiceHandler) RemoveOne(ctx context.Context, in *RequestInfo, out *ReplyInfo) error {
	return h.PanoramaServiceHandler.RemoveOne(ctx, in, out)
}

func (h *panoramaServiceHandler) GetOne(ctx context.Context, in *RequestInfo, out *ReplyPanoramaInfo) error {
	return h.PanoramaServiceHandler.GetOne(ctx, in, out)
}

func (h *panoramaServiceHandler) Search(ctx context.Context, in *RequestInfo, out *ReplyPanoramaList) error {
	return h.PanoramaServiceHandler.Search(ctx, in, out)
}

func (h *panoramaServiceHandler) GetListByFilter(ctx context.Context, in *RequestFilter, out *ReplyPanoramaList) error {
	return h.PanoramaServiceHandler.GetListByFilter(ctx, in, out)
}

func (h *panoramaServiceHandler) GetStatistic(ctx context.Context, in *RequestFilter, out *ReplyStatistic) error {
	return h.PanoramaServiceHandler.GetStatistic(ctx, in, out)
}

func (h *panoramaServiceHandler) UpdateByFilter(ctx context.Context, in *RequestUpdate, out *ReplyInfo) error {
	return h.PanoramaServiceHandler.UpdateByFilter(ctx, in, out)
}