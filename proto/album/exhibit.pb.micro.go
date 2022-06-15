// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/album/exhibit.proto

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

// Client API for ExhibitService service

type ExhibitService interface {
	AddOne(ctx context.Context, in *ReqExhibitAdd, opts ...client.CallOption) (*ReplyExhibitInfo, error)
	UpdateBase(ctx context.Context, in *ReqExhibitUpdate, opts ...client.CallOption) (*ReplyInfo, error)
	RemoveOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyInfo, error)
	GetOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyExhibitInfo, error)
	Search(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyExhibitList, error)
	GetListByFilter(ctx context.Context, in *RequestFilter, opts ...client.CallOption) (*ReplyExhibitList, error)
	GetStatistic(ctx context.Context, in *RequestFilter, opts ...client.CallOption) (*ReplyStatistic, error)
	UpdateByFilter(ctx context.Context, in *RequestUpdate, opts ...client.CallOption) (*ReplyInfo, error)
	UpdateStatus(ctx context.Context, in *RequestIntFlag, opts ...client.CallOption) (*ReplyInfo, error)
	AppendAsset(ctx context.Context, in *RequestList, opts ...client.CallOption) (*ReplyList, error)
	SubtractAsset(ctx context.Context, in *RequestList, opts ...client.CallOption) (*ReplyList, error)
}

type exhibitService struct {
	c    client.Client
	name string
}

func NewExhibitService(name string, c client.Client) ExhibitService {
	return &exhibitService{
		c:    c,
		name: name,
	}
}

func (c *exhibitService) AddOne(ctx context.Context, in *ReqExhibitAdd, opts ...client.CallOption) (*ReplyExhibitInfo, error) {
	req := c.c.NewRequest(c.name, "ExhibitService.AddOne", in)
	out := new(ReplyExhibitInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exhibitService) UpdateBase(ctx context.Context, in *ReqExhibitUpdate, opts ...client.CallOption) (*ReplyInfo, error) {
	req := c.c.NewRequest(c.name, "ExhibitService.UpdateBase", in)
	out := new(ReplyInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exhibitService) RemoveOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyInfo, error) {
	req := c.c.NewRequest(c.name, "ExhibitService.RemoveOne", in)
	out := new(ReplyInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exhibitService) GetOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyExhibitInfo, error) {
	req := c.c.NewRequest(c.name, "ExhibitService.GetOne", in)
	out := new(ReplyExhibitInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exhibitService) Search(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyExhibitList, error) {
	req := c.c.NewRequest(c.name, "ExhibitService.Search", in)
	out := new(ReplyExhibitList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exhibitService) GetListByFilter(ctx context.Context, in *RequestFilter, opts ...client.CallOption) (*ReplyExhibitList, error) {
	req := c.c.NewRequest(c.name, "ExhibitService.GetListByFilter", in)
	out := new(ReplyExhibitList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exhibitService) GetStatistic(ctx context.Context, in *RequestFilter, opts ...client.CallOption) (*ReplyStatistic, error) {
	req := c.c.NewRequest(c.name, "ExhibitService.GetStatistic", in)
	out := new(ReplyStatistic)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exhibitService) UpdateByFilter(ctx context.Context, in *RequestUpdate, opts ...client.CallOption) (*ReplyInfo, error) {
	req := c.c.NewRequest(c.name, "ExhibitService.UpdateByFilter", in)
	out := new(ReplyInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exhibitService) UpdateStatus(ctx context.Context, in *RequestIntFlag, opts ...client.CallOption) (*ReplyInfo, error) {
	req := c.c.NewRequest(c.name, "ExhibitService.UpdateStatus", in)
	out := new(ReplyInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exhibitService) AppendAsset(ctx context.Context, in *RequestList, opts ...client.CallOption) (*ReplyList, error) {
	req := c.c.NewRequest(c.name, "ExhibitService.AppendAsset", in)
	out := new(ReplyList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exhibitService) SubtractAsset(ctx context.Context, in *RequestList, opts ...client.CallOption) (*ReplyList, error) {
	req := c.c.NewRequest(c.name, "ExhibitService.SubtractAsset", in)
	out := new(ReplyList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ExhibitService service

type ExhibitServiceHandler interface {
	AddOne(context.Context, *ReqExhibitAdd, *ReplyExhibitInfo) error
	UpdateBase(context.Context, *ReqExhibitUpdate, *ReplyInfo) error
	RemoveOne(context.Context, *RequestInfo, *ReplyInfo) error
	GetOne(context.Context, *RequestInfo, *ReplyExhibitInfo) error
	Search(context.Context, *RequestInfo, *ReplyExhibitList) error
	GetListByFilter(context.Context, *RequestFilter, *ReplyExhibitList) error
	GetStatistic(context.Context, *RequestFilter, *ReplyStatistic) error
	UpdateByFilter(context.Context, *RequestUpdate, *ReplyInfo) error
	UpdateStatus(context.Context, *RequestIntFlag, *ReplyInfo) error
	AppendAsset(context.Context, *RequestList, *ReplyList) error
	SubtractAsset(context.Context, *RequestList, *ReplyList) error
}

func RegisterExhibitServiceHandler(s server.Server, hdlr ExhibitServiceHandler, opts ...server.HandlerOption) error {
	type exhibitService interface {
		AddOne(ctx context.Context, in *ReqExhibitAdd, out *ReplyExhibitInfo) error
		UpdateBase(ctx context.Context, in *ReqExhibitUpdate, out *ReplyInfo) error
		RemoveOne(ctx context.Context, in *RequestInfo, out *ReplyInfo) error
		GetOne(ctx context.Context, in *RequestInfo, out *ReplyExhibitInfo) error
		Search(ctx context.Context, in *RequestInfo, out *ReplyExhibitList) error
		GetListByFilter(ctx context.Context, in *RequestFilter, out *ReplyExhibitList) error
		GetStatistic(ctx context.Context, in *RequestFilter, out *ReplyStatistic) error
		UpdateByFilter(ctx context.Context, in *RequestUpdate, out *ReplyInfo) error
		UpdateStatus(ctx context.Context, in *RequestIntFlag, out *ReplyInfo) error
		AppendAsset(ctx context.Context, in *RequestList, out *ReplyList) error
		SubtractAsset(ctx context.Context, in *RequestList, out *ReplyList) error
	}
	type ExhibitService struct {
		exhibitService
	}
	h := &exhibitServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&ExhibitService{h}, opts...))
}

type exhibitServiceHandler struct {
	ExhibitServiceHandler
}

func (h *exhibitServiceHandler) AddOne(ctx context.Context, in *ReqExhibitAdd, out *ReplyExhibitInfo) error {
	return h.ExhibitServiceHandler.AddOne(ctx, in, out)
}

func (h *exhibitServiceHandler) UpdateBase(ctx context.Context, in *ReqExhibitUpdate, out *ReplyInfo) error {
	return h.ExhibitServiceHandler.UpdateBase(ctx, in, out)
}

func (h *exhibitServiceHandler) RemoveOne(ctx context.Context, in *RequestInfo, out *ReplyInfo) error {
	return h.ExhibitServiceHandler.RemoveOne(ctx, in, out)
}

func (h *exhibitServiceHandler) GetOne(ctx context.Context, in *RequestInfo, out *ReplyExhibitInfo) error {
	return h.ExhibitServiceHandler.GetOne(ctx, in, out)
}

func (h *exhibitServiceHandler) Search(ctx context.Context, in *RequestInfo, out *ReplyExhibitList) error {
	return h.ExhibitServiceHandler.Search(ctx, in, out)
}

func (h *exhibitServiceHandler) GetListByFilter(ctx context.Context, in *RequestFilter, out *ReplyExhibitList) error {
	return h.ExhibitServiceHandler.GetListByFilter(ctx, in, out)
}

func (h *exhibitServiceHandler) GetStatistic(ctx context.Context, in *RequestFilter, out *ReplyStatistic) error {
	return h.ExhibitServiceHandler.GetStatistic(ctx, in, out)
}

func (h *exhibitServiceHandler) UpdateByFilter(ctx context.Context, in *RequestUpdate, out *ReplyInfo) error {
	return h.ExhibitServiceHandler.UpdateByFilter(ctx, in, out)
}

func (h *exhibitServiceHandler) UpdateStatus(ctx context.Context, in *RequestIntFlag, out *ReplyInfo) error {
	return h.ExhibitServiceHandler.UpdateStatus(ctx, in, out)
}

func (h *exhibitServiceHandler) AppendAsset(ctx context.Context, in *RequestList, out *ReplyList) error {
	return h.ExhibitServiceHandler.AppendAsset(ctx, in, out)
}

func (h *exhibitServiceHandler) SubtractAsset(ctx context.Context, in *RequestList, out *ReplyList) error {
	return h.ExhibitServiceHandler.SubtractAsset(ctx, in, out)
}