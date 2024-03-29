// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/album/sheet.proto

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

// Client API for SheetService service

type SheetService interface {
	AddOne(ctx context.Context, in *ReqSheetAdd, opts ...client.CallOption) (*ReplySheetInfo, error)
	UpdateBase(ctx context.Context, in *ReqSheetUpdate, opts ...client.CallOption) (*ReplyInfo, error)
	RemoveOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyInfo, error)
	GetOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplySheetInfo, error)
	Search(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplySheetList, error)
	GetListBy(ctx context.Context, in *RequestFilter, opts ...client.CallOption) (*ReplySheetList, error)
	GetStatistic(ctx context.Context, in *RequestFilter, opts ...client.CallOption) (*ReplyStatistic, error)
	UpdateBy(ctx context.Context, in *RequestUpdate, opts ...client.CallOption) (*ReplyInfo, error)
	UpdatePages(ctx context.Context, in *ReqSheetPages, opts ...client.CallOption) (*ReplySheetPages, error)
}

type sheetService struct {
	c    client.Client
	name string
}

func NewSheetService(name string, c client.Client) SheetService {
	return &sheetService{
		c:    c,
		name: name,
	}
}

func (c *sheetService) AddOne(ctx context.Context, in *ReqSheetAdd, opts ...client.CallOption) (*ReplySheetInfo, error) {
	req := c.c.NewRequest(c.name, "SheetService.AddOne", in)
	out := new(ReplySheetInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sheetService) UpdateBase(ctx context.Context, in *ReqSheetUpdate, opts ...client.CallOption) (*ReplyInfo, error) {
	req := c.c.NewRequest(c.name, "SheetService.UpdateBase", in)
	out := new(ReplyInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sheetService) RemoveOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyInfo, error) {
	req := c.c.NewRequest(c.name, "SheetService.RemoveOne", in)
	out := new(ReplyInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sheetService) GetOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplySheetInfo, error) {
	req := c.c.NewRequest(c.name, "SheetService.GetOne", in)
	out := new(ReplySheetInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sheetService) Search(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplySheetList, error) {
	req := c.c.NewRequest(c.name, "SheetService.Search", in)
	out := new(ReplySheetList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sheetService) GetListBy(ctx context.Context, in *RequestFilter, opts ...client.CallOption) (*ReplySheetList, error) {
	req := c.c.NewRequest(c.name, "SheetService.GetListBy", in)
	out := new(ReplySheetList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sheetService) GetStatistic(ctx context.Context, in *RequestFilter, opts ...client.CallOption) (*ReplyStatistic, error) {
	req := c.c.NewRequest(c.name, "SheetService.GetStatistic", in)
	out := new(ReplyStatistic)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sheetService) UpdateBy(ctx context.Context, in *RequestUpdate, opts ...client.CallOption) (*ReplyInfo, error) {
	req := c.c.NewRequest(c.name, "SheetService.UpdateBy", in)
	out := new(ReplyInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sheetService) UpdatePages(ctx context.Context, in *ReqSheetPages, opts ...client.CallOption) (*ReplySheetPages, error) {
	req := c.c.NewRequest(c.name, "SheetService.UpdatePages", in)
	out := new(ReplySheetPages)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SheetService service

type SheetServiceHandler interface {
	AddOne(context.Context, *ReqSheetAdd, *ReplySheetInfo) error
	UpdateBase(context.Context, *ReqSheetUpdate, *ReplyInfo) error
	RemoveOne(context.Context, *RequestInfo, *ReplyInfo) error
	GetOne(context.Context, *RequestInfo, *ReplySheetInfo) error
	Search(context.Context, *RequestInfo, *ReplySheetList) error
	GetListBy(context.Context, *RequestFilter, *ReplySheetList) error
	GetStatistic(context.Context, *RequestFilter, *ReplyStatistic) error
	UpdateBy(context.Context, *RequestUpdate, *ReplyInfo) error
	UpdatePages(context.Context, *ReqSheetPages, *ReplySheetPages) error
}

func RegisterSheetServiceHandler(s server.Server, hdlr SheetServiceHandler, opts ...server.HandlerOption) error {
	type sheetService interface {
		AddOne(ctx context.Context, in *ReqSheetAdd, out *ReplySheetInfo) error
		UpdateBase(ctx context.Context, in *ReqSheetUpdate, out *ReplyInfo) error
		RemoveOne(ctx context.Context, in *RequestInfo, out *ReplyInfo) error
		GetOne(ctx context.Context, in *RequestInfo, out *ReplySheetInfo) error
		Search(ctx context.Context, in *RequestInfo, out *ReplySheetList) error
		GetListBy(ctx context.Context, in *RequestFilter, out *ReplySheetList) error
		GetStatistic(ctx context.Context, in *RequestFilter, out *ReplyStatistic) error
		UpdateBy(ctx context.Context, in *RequestUpdate, out *ReplyInfo) error
		UpdatePages(ctx context.Context, in *ReqSheetPages, out *ReplySheetPages) error
	}
	type SheetService struct {
		sheetService
	}
	h := &sheetServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&SheetService{h}, opts...))
}

type sheetServiceHandler struct {
	SheetServiceHandler
}

func (h *sheetServiceHandler) AddOne(ctx context.Context, in *ReqSheetAdd, out *ReplySheetInfo) error {
	return h.SheetServiceHandler.AddOne(ctx, in, out)
}

func (h *sheetServiceHandler) UpdateBase(ctx context.Context, in *ReqSheetUpdate, out *ReplyInfo) error {
	return h.SheetServiceHandler.UpdateBase(ctx, in, out)
}

func (h *sheetServiceHandler) RemoveOne(ctx context.Context, in *RequestInfo, out *ReplyInfo) error {
	return h.SheetServiceHandler.RemoveOne(ctx, in, out)
}

func (h *sheetServiceHandler) GetOne(ctx context.Context, in *RequestInfo, out *ReplySheetInfo) error {
	return h.SheetServiceHandler.GetOne(ctx, in, out)
}

func (h *sheetServiceHandler) Search(ctx context.Context, in *RequestInfo, out *ReplySheetList) error {
	return h.SheetServiceHandler.Search(ctx, in, out)
}

func (h *sheetServiceHandler) GetListBy(ctx context.Context, in *RequestFilter, out *ReplySheetList) error {
	return h.SheetServiceHandler.GetListBy(ctx, in, out)
}

func (h *sheetServiceHandler) GetStatistic(ctx context.Context, in *RequestFilter, out *ReplyStatistic) error {
	return h.SheetServiceHandler.GetStatistic(ctx, in, out)
}

func (h *sheetServiceHandler) UpdateBy(ctx context.Context, in *RequestUpdate, out *ReplyInfo) error {
	return h.SheetServiceHandler.UpdateBy(ctx, in, out)
}

func (h *sheetServiceHandler) UpdatePages(ctx context.Context, in *ReqSheetPages, out *ReplySheetPages) error {
	return h.SheetServiceHandler.UpdatePages(ctx, in, out)
}
