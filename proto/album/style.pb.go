// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/album/style.proto

package album

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
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

type StyleInfo struct {
	Uid                  string         `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Id                   uint64         `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Created              int64          `protobuf:"varint,3,opt,name=created,proto3" json:"created,omitempty"`
	Updated              int64          `protobuf:"varint,4,opt,name=updated,proto3" json:"updated,omitempty"`
	Creator              string         `protobuf:"bytes,5,opt,name=creator,proto3" json:"creator,omitempty"`
	Operator             string         `protobuf:"bytes,6,opt,name=operator,proto3" json:"operator,omitempty"`
	Name                 string         `protobuf:"bytes,7,opt,name=name,proto3" json:"name,omitempty"`
	Remark               string         `protobuf:"bytes,8,opt,name=remark,proto3" json:"remark,omitempty"`
	Type                 uint32         `protobuf:"varint,9,opt,name=type,proto3" json:"type,omitempty"`
	Cover                string         `protobuf:"bytes,10,opt,name=cover,proto3" json:"cover,omitempty"`
	Background           string         `protobuf:"bytes,11,opt,name=background,proto3" json:"background,omitempty"`
	Prefix               string         `protobuf:"bytes,12,opt,name=prefix,proto3" json:"prefix,omitempty"`
	Width                uint32         `protobuf:"varint,13,opt,name=width,proto3" json:"width,omitempty"`
	Height               uint32         `protobuf:"varint,14,opt,name=height,proto3" json:"height,omitempty"`
	Tags                 []string       `protobuf:"bytes,21,rep,name=tags,proto3" json:"tags,omitempty"`
	Scenes               []string       `protobuf:"bytes,22,rep,name=scenes,proto3" json:"scenes,omitempty"`
	Slots                []*StyleSlot   `protobuf:"bytes,31,rep,name=slots,proto3" json:"slots,omitempty"`
	Relates              []*StyleRelate `protobuf:"bytes,32,rep,name=relates,proto3" json:"relates,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *StyleInfo) Reset()         { *m = StyleInfo{} }
func (m *StyleInfo) String() string { return proto.CompactTextString(m) }
func (*StyleInfo) ProtoMessage()    {}
func (*StyleInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_c94d61fa26c563fe, []int{0}
}

func (m *StyleInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StyleInfo.Unmarshal(m, b)
}
func (m *StyleInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StyleInfo.Marshal(b, m, deterministic)
}
func (m *StyleInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StyleInfo.Merge(m, src)
}
func (m *StyleInfo) XXX_Size() int {
	return xxx_messageInfo_StyleInfo.Size(m)
}
func (m *StyleInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_StyleInfo.DiscardUnknown(m)
}

var xxx_messageInfo_StyleInfo proto.InternalMessageInfo

func (m *StyleInfo) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *StyleInfo) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *StyleInfo) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *StyleInfo) GetUpdated() int64 {
	if m != nil {
		return m.Updated
	}
	return 0
}

func (m *StyleInfo) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *StyleInfo) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *StyleInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *StyleInfo) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *StyleInfo) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *StyleInfo) GetCover() string {
	if m != nil {
		return m.Cover
	}
	return ""
}

func (m *StyleInfo) GetBackground() string {
	if m != nil {
		return m.Background
	}
	return ""
}

func (m *StyleInfo) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

func (m *StyleInfo) GetWidth() uint32 {
	if m != nil {
		return m.Width
	}
	return 0
}

func (m *StyleInfo) GetHeight() uint32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *StyleInfo) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *StyleInfo) GetScenes() []string {
	if m != nil {
		return m.Scenes
	}
	return nil
}

func (m *StyleInfo) GetSlots() []*StyleSlot {
	if m != nil {
		return m.Slots
	}
	return nil
}

func (m *StyleInfo) GetRelates() []*StyleRelate {
	if m != nil {
		return m.Relates
	}
	return nil
}

type StyleSlot struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	X                    int32    `protobuf:"varint,2,opt,name=x,proto3" json:"x,omitempty"`
	Y                    int32    `protobuf:"varint,3,opt,name=y,proto3" json:"y,omitempty"`
	Width                uint32   `protobuf:"varint,4,opt,name=width,proto3" json:"width,omitempty"`
	Height               uint32   `protobuf:"varint,5,opt,name=height,proto3" json:"height,omitempty"`
	Bold                 uint32   `protobuf:"varint,6,opt,name=bold,proto3" json:"bold,omitempty"`
	Size                 uint32   `protobuf:"varint,7,opt,name=size,proto3" json:"size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StyleSlot) Reset()         { *m = StyleSlot{} }
func (m *StyleSlot) String() string { return proto.CompactTextString(m) }
func (*StyleSlot) ProtoMessage()    {}
func (*StyleSlot) Descriptor() ([]byte, []int) {
	return fileDescriptor_c94d61fa26c563fe, []int{1}
}

func (m *StyleSlot) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StyleSlot.Unmarshal(m, b)
}
func (m *StyleSlot) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StyleSlot.Marshal(b, m, deterministic)
}
func (m *StyleSlot) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StyleSlot.Merge(m, src)
}
func (m *StyleSlot) XXX_Size() int {
	return xxx_messageInfo_StyleSlot.Size(m)
}
func (m *StyleSlot) XXX_DiscardUnknown() {
	xxx_messageInfo_StyleSlot.DiscardUnknown(m)
}

var xxx_messageInfo_StyleSlot proto.InternalMessageInfo

func (m *StyleSlot) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *StyleSlot) GetX() int32 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *StyleSlot) GetY() int32 {
	if m != nil {
		return m.Y
	}
	return 0
}

func (m *StyleSlot) GetWidth() uint32 {
	if m != nil {
		return m.Width
	}
	return 0
}

func (m *StyleSlot) GetHeight() uint32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *StyleSlot) GetBold() uint32 {
	if m != nil {
		return m.Bold
	}
	return 0
}

func (m *StyleSlot) GetSize() uint32 {
	if m != nil {
		return m.Size
	}
	return 0
}

type StyleRelate struct {
	Entity               string   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Way                  uint32   `protobuf:"varint,2,opt,name=way,proto3" json:"way,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StyleRelate) Reset()         { *m = StyleRelate{} }
func (m *StyleRelate) String() string { return proto.CompactTextString(m) }
func (*StyleRelate) ProtoMessage()    {}
func (*StyleRelate) Descriptor() ([]byte, []int) {
	return fileDescriptor_c94d61fa26c563fe, []int{2}
}

func (m *StyleRelate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StyleRelate.Unmarshal(m, b)
}
func (m *StyleRelate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StyleRelate.Marshal(b, m, deterministic)
}
func (m *StyleRelate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StyleRelate.Merge(m, src)
}
func (m *StyleRelate) XXX_Size() int {
	return xxx_messageInfo_StyleRelate.Size(m)
}
func (m *StyleRelate) XXX_DiscardUnknown() {
	xxx_messageInfo_StyleRelate.DiscardUnknown(m)
}

var xxx_messageInfo_StyleRelate proto.InternalMessageInfo

func (m *StyleRelate) GetEntity() string {
	if m != nil {
		return m.Entity
	}
	return ""
}

func (m *StyleRelate) GetWay() uint32 {
	if m != nil {
		return m.Way
	}
	return 0
}

type ReqStyleAdd struct {
	Name                 string       `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Remark               string       `protobuf:"bytes,2,opt,name=remark,proto3" json:"remark,omitempty"`
	Operator             string       `protobuf:"bytes,3,opt,name=operator,proto3" json:"operator,omitempty"`
	Type                 uint32       `protobuf:"varint,4,opt,name=type,proto3" json:"type,omitempty"`
	Background           string       `protobuf:"bytes,5,opt,name=background,proto3" json:"background,omitempty"`
	Cover                string       `protobuf:"bytes,6,opt,name=cover,proto3" json:"cover,omitempty"`
	Prefix               string       `protobuf:"bytes,7,opt,name=prefix,proto3" json:"prefix,omitempty"`
	Width                uint32       `protobuf:"varint,8,opt,name=width,proto3" json:"width,omitempty"`
	Height               uint32       `protobuf:"varint,9,opt,name=height,proto3" json:"height,omitempty"`
	Tags                 []string     `protobuf:"bytes,15,rep,name=tags,proto3" json:"tags,omitempty"`
	Scenes               []string     `protobuf:"bytes,16,rep,name=scenes,proto3" json:"scenes,omitempty"`
	Slots                []*StyleSlot `protobuf:"bytes,21,rep,name=slots,proto3" json:"slots,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReqStyleAdd) Reset()         { *m = ReqStyleAdd{} }
func (m *ReqStyleAdd) String() string { return proto.CompactTextString(m) }
func (*ReqStyleAdd) ProtoMessage()    {}
func (*ReqStyleAdd) Descriptor() ([]byte, []int) {
	return fileDescriptor_c94d61fa26c563fe, []int{3}
}

func (m *ReqStyleAdd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqStyleAdd.Unmarshal(m, b)
}
func (m *ReqStyleAdd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqStyleAdd.Marshal(b, m, deterministic)
}
func (m *ReqStyleAdd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqStyleAdd.Merge(m, src)
}
func (m *ReqStyleAdd) XXX_Size() int {
	return xxx_messageInfo_ReqStyleAdd.Size(m)
}
func (m *ReqStyleAdd) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqStyleAdd.DiscardUnknown(m)
}

var xxx_messageInfo_ReqStyleAdd proto.InternalMessageInfo

func (m *ReqStyleAdd) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReqStyleAdd) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *ReqStyleAdd) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ReqStyleAdd) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *ReqStyleAdd) GetBackground() string {
	if m != nil {
		return m.Background
	}
	return ""
}

func (m *ReqStyleAdd) GetCover() string {
	if m != nil {
		return m.Cover
	}
	return ""
}

func (m *ReqStyleAdd) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

func (m *ReqStyleAdd) GetWidth() uint32 {
	if m != nil {
		return m.Width
	}
	return 0
}

func (m *ReqStyleAdd) GetHeight() uint32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *ReqStyleAdd) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *ReqStyleAdd) GetScenes() []string {
	if m != nil {
		return m.Scenes
	}
	return nil
}

func (m *ReqStyleAdd) GetSlots() []*StyleSlot {
	if m != nil {
		return m.Slots
	}
	return nil
}

type ReqStyleUpdate struct {
	Uid                  string       `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Name                 string       `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Remark               string       `protobuf:"bytes,3,opt,name=remark,proto3" json:"remark,omitempty"`
	Type                 uint32       `protobuf:"varint,4,opt,name=type,proto3" json:"type,omitempty"`
	Operator             string       `protobuf:"bytes,5,opt,name=operator,proto3" json:"operator,omitempty"`
	Background           string       `protobuf:"bytes,6,opt,name=background,proto3" json:"background,omitempty"`
	Width                uint32       `protobuf:"varint,7,opt,name=width,proto3" json:"width,omitempty"`
	Height               uint32       `protobuf:"varint,8,opt,name=height,proto3" json:"height,omitempty"`
	Tags                 []string     `protobuf:"bytes,15,rep,name=tags,proto3" json:"tags,omitempty"`
	Scenes               []string     `protobuf:"bytes,16,rep,name=scenes,proto3" json:"scenes,omitempty"`
	Slots                []*StyleSlot `protobuf:"bytes,21,rep,name=slots,proto3" json:"slots,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReqStyleUpdate) Reset()         { *m = ReqStyleUpdate{} }
func (m *ReqStyleUpdate) String() string { return proto.CompactTextString(m) }
func (*ReqStyleUpdate) ProtoMessage()    {}
func (*ReqStyleUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_c94d61fa26c563fe, []int{4}
}

func (m *ReqStyleUpdate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqStyleUpdate.Unmarshal(m, b)
}
func (m *ReqStyleUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqStyleUpdate.Marshal(b, m, deterministic)
}
func (m *ReqStyleUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqStyleUpdate.Merge(m, src)
}
func (m *ReqStyleUpdate) XXX_Size() int {
	return xxx_messageInfo_ReqStyleUpdate.Size(m)
}
func (m *ReqStyleUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqStyleUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_ReqStyleUpdate proto.InternalMessageInfo

func (m *ReqStyleUpdate) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ReqStyleUpdate) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReqStyleUpdate) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *ReqStyleUpdate) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *ReqStyleUpdate) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ReqStyleUpdate) GetBackground() string {
	if m != nil {
		return m.Background
	}
	return ""
}

func (m *ReqStyleUpdate) GetWidth() uint32 {
	if m != nil {
		return m.Width
	}
	return 0
}

func (m *ReqStyleUpdate) GetHeight() uint32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *ReqStyleUpdate) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *ReqStyleUpdate) GetScenes() []string {
	if m != nil {
		return m.Scenes
	}
	return nil
}

func (m *ReqStyleUpdate) GetSlots() []*StyleSlot {
	if m != nil {
		return m.Slots
	}
	return nil
}

type ReplyStyleInfo struct {
	Status               *ReplyStatus `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Info                 *StyleInfo   `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReplyStyleInfo) Reset()         { *m = ReplyStyleInfo{} }
func (m *ReplyStyleInfo) String() string { return proto.CompactTextString(m) }
func (*ReplyStyleInfo) ProtoMessage()    {}
func (*ReplyStyleInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_c94d61fa26c563fe, []int{5}
}

func (m *ReplyStyleInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyStyleInfo.Unmarshal(m, b)
}
func (m *ReplyStyleInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyStyleInfo.Marshal(b, m, deterministic)
}
func (m *ReplyStyleInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyStyleInfo.Merge(m, src)
}
func (m *ReplyStyleInfo) XXX_Size() int {
	return xxx_messageInfo_ReplyStyleInfo.Size(m)
}
func (m *ReplyStyleInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyStyleInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyStyleInfo proto.InternalMessageInfo

func (m *ReplyStyleInfo) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ReplyStyleInfo) GetInfo() *StyleInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

type ReplyStyleList struct {
	Status               *ReplyStatus `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Total                uint32       `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	Pages                uint32       `protobuf:"varint,3,opt,name=pages,proto3" json:"pages,omitempty"`
	List                 []*StyleInfo `protobuf:"bytes,11,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReplyStyleList) Reset()         { *m = ReplyStyleList{} }
func (m *ReplyStyleList) String() string { return proto.CompactTextString(m) }
func (*ReplyStyleList) ProtoMessage()    {}
func (*ReplyStyleList) Descriptor() ([]byte, []int) {
	return fileDescriptor_c94d61fa26c563fe, []int{6}
}

func (m *ReplyStyleList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyStyleList.Unmarshal(m, b)
}
func (m *ReplyStyleList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyStyleList.Marshal(b, m, deterministic)
}
func (m *ReplyStyleList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyStyleList.Merge(m, src)
}
func (m *ReplyStyleList) XXX_Size() int {
	return xxx_messageInfo_ReplyStyleList.Size(m)
}
func (m *ReplyStyleList) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyStyleList.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyStyleList proto.InternalMessageInfo

func (m *ReplyStyleList) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ReplyStyleList) GetTotal() uint32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *ReplyStyleList) GetPages() uint32 {
	if m != nil {
		return m.Pages
	}
	return 0
}

func (m *ReplyStyleList) GetList() []*StyleInfo {
	if m != nil {
		return m.List
	}
	return nil
}

func init() {
	proto.RegisterType((*StyleInfo)(nil), "omo.msp.album.StyleInfo")
	proto.RegisterType((*StyleSlot)(nil), "omo.msp.album.StyleSlot")
	proto.RegisterType((*StyleRelate)(nil), "omo.msp.album.StyleRelate")
	proto.RegisterType((*ReqStyleAdd)(nil), "omo.msp.album.ReqStyleAdd")
	proto.RegisterType((*ReqStyleUpdate)(nil), "omo.msp.album.ReqStyleUpdate")
	proto.RegisterType((*ReplyStyleInfo)(nil), "omo.msp.album.ReplyStyleInfo")
	proto.RegisterType((*ReplyStyleList)(nil), "omo.msp.album.ReplyStyleList")
}

func init() {
	proto.RegisterFile("proto/album/style.proto", fileDescriptor_c94d61fa26c563fe)
}

var fileDescriptor_c94d61fa26c563fe = []byte{
	// 778 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x56, 0x41, 0x6e, 0xdb, 0x38,
	0x14, 0x8d, 0x2c, 0x4b, 0xb6, 0x69, 0xcb, 0x09, 0x88, 0x24, 0x43, 0x18, 0x93, 0x19, 0x43, 0x2b,
	0x2f, 0x06, 0x0e, 0x90, 0x19, 0x60, 0xd6, 0xc9, 0x20, 0x63, 0xcc, 0xa0, 0x45, 0x0b, 0x06, 0xdd,
	0x74, 0x27, 0x4b, 0x8c, 0x2d, 0x58, 0x12, 0x55, 0x91, 0x4e, 0xa2, 0x5e, 0xa2, 0x37, 0xe8, 0xba,
	0xd7, 0xe8, 0x45, 0xba, 0xeb, 0x3d, 0x0a, 0x7e, 0xca, 0x8e, 0x1c, 0xd3, 0x75, 0x91, 0x45, 0x77,
	0x7c, 0xfc, 0x8f, 0x8f, 0x5f, 0xef, 0x3f, 0x1a, 0x46, 0xbf, 0xe4, 0x05, 0x97, 0xfc, 0x3c, 0x48,
	0xa6, 0xcb, 0xf4, 0x5c, 0xc8, 0x32, 0x61, 0x63, 0xd8, 0xc1, 0x1e, 0x4f, 0xf9, 0x38, 0x15, 0xf9,
	0x18, 0x4a, 0x03, 0x52, 0xe7, 0x85, 0x3c, 0x4d, 0x79, 0xa6, 0x89, 0xfe, 0x57, 0x1b, 0x75, 0x6e,
	0xd4, 0xc1, 0xff, 0xb2, 0x5b, 0x8e, 0x8f, 0x90, 0xbd, 0x8c, 0x23, 0x62, 0x0d, 0xad, 0x51, 0x87,
	0xaa, 0x25, 0xee, 0xa3, 0x46, 0x1c, 0x91, 0xc6, 0xd0, 0x1a, 0x35, 0x69, 0x23, 0x8e, 0x30, 0x41,
	0xad, 0xb0, 0x60, 0x81, 0x64, 0x11, 0xb1, 0x87, 0xd6, 0xc8, 0xa6, 0x2b, 0xa8, 0x2a, 0xcb, 0x3c,
	0x82, 0x4a, 0x53, 0x57, 0x2a, 0xb8, 0x3e, 0xc3, 0x0b, 0xe2, 0x80, 0xf2, 0x0a, 0xe2, 0x01, 0x6a,
	0xf3, 0x9c, 0x15, 0x50, 0x72, 0xa1, 0xb4, 0xc6, 0x18, 0xa3, 0x66, 0x16, 0xa4, 0x8c, 0xb4, 0x60,
	0x1f, 0xd6, 0xf8, 0x14, 0xb9, 0x05, 0x4b, 0x83, 0x62, 0x41, 0xda, 0xb0, 0x5b, 0x21, 0xc5, 0x95,
	0x65, 0xce, 0x48, 0x67, 0x68, 0x8d, 0x3c, 0x0a, 0x6b, 0x7c, 0x8c, 0x9c, 0x90, 0xdf, 0xb1, 0x82,
	0x20, 0xa0, 0x6a, 0x80, 0x7f, 0x43, 0x68, 0x1a, 0x84, 0x8b, 0x59, 0xc1, 0x97, 0x59, 0x44, 0xba,
	0x50, 0xaa, 0xed, 0xa8, 0x1b, 0xf2, 0x82, 0xdd, 0xc6, 0x0f, 0xa4, 0xa7, 0x6f, 0xd0, 0x48, 0xa9,
	0xdd, 0xc7, 0x91, 0x9c, 0x13, 0x0f, 0xae, 0xd0, 0x40, 0xb1, 0xe7, 0x2c, 0x9e, 0xcd, 0x25, 0xe9,
	0xc3, 0x76, 0x85, 0xa0, 0x9f, 0x60, 0x26, 0xc8, 0xc9, 0xd0, 0x56, 0xbd, 0xab, 0xb5, 0xe2, 0x8a,
	0x90, 0x65, 0x4c, 0x90, 0x53, 0xd8, 0xad, 0x10, 0x1e, 0x23, 0x47, 0x24, 0x5c, 0x0a, 0xf2, 0xfb,
	0xd0, 0x1e, 0x75, 0x2f, 0xc8, 0x78, 0x63, 0x74, 0x63, 0x18, 0xce, 0x4d, 0xc2, 0x25, 0xd5, 0x34,
	0xfc, 0x17, 0x6a, 0x15, 0x2c, 0x09, 0x24, 0x13, 0x64, 0x08, 0x27, 0x06, 0xa6, 0x13, 0x14, 0x28,
	0x74, 0x45, 0xf5, 0x3f, 0x58, 0xd5, 0x9c, 0x95, 0x94, 0x9a, 0xf3, 0x82, 0x95, 0xab, 0x39, 0x2f,
	0x58, 0x89, 0x7b, 0xc8, 0x7a, 0x80, 0x31, 0x3b, 0xd4, 0x7a, 0x50, 0xa8, 0x84, 0xf9, 0x3a, 0xd4,
	0x2a, 0x1f, 0xbf, 0xbd, 0x69, 0xfe, 0x76, 0xe7, 0xe9, 0xb7, 0x4f, 0x79, 0x12, 0xc1, 0x3c, 0x3d,
	0x0a, 0x6b, 0xb5, 0x27, 0xe2, 0xf7, 0x7a, 0x96, 0x1e, 0x85, 0xb5, 0xff, 0x37, 0xea, 0xd6, 0x3a,
	0x55, 0x72, 0x2c, 0x93, 0xb1, 0x5c, 0x75, 0x55, 0x21, 0xd5, 0xea, 0x7d, 0x50, 0x42, 0x6b, 0x1e,
	0x55, 0x4b, 0xff, 0x73, 0x03, 0x75, 0x29, 0x7b, 0x07, 0x87, 0x2f, 0xa3, 0x68, 0x1d, 0x14, 0xcb,
	0x18, 0x94, 0xc6, 0x46, 0x50, 0xea, 0x81, 0xb3, 0xb7, 0x03, 0x07, 0x21, 0x6a, 0xd6, 0x42, 0xb4,
	0x19, 0x17, 0x67, 0x2b, 0x2e, 0xeb, 0x90, 0xb9, 0xf5, 0x90, 0x3d, 0x86, 0xa8, 0x65, 0x0e, 0x51,
	0xdb, 0x6c, 0x64, 0xc7, 0x18, 0xa2, 0x43, 0x63, 0x88, 0x8e, 0xcc, 0x21, 0x3a, 0xf9, 0xa1, 0x10,
	0xf9, 0x9f, 0x1a, 0xa8, 0xbf, 0xf2, 0xf0, 0x0d, 0x3c, 0x53, 0xc3, 0xdb, 0x5f, 0x19, 0xdb, 0x30,
	0x1a, 0x6b, 0x1b, 0x5f, 0x60, 0xdd, 0xbc, 0xba, 0xd9, 0xce, 0x13, 0xb3, 0x37, 0x8d, 0x75, 0x4d,
	0xc6, 0x6a, 0xab, 0x5a, 0x66, 0xab, 0xda, 0x3f, 0xcd, 0xaa, 0x42, 0x39, 0x95, 0x27, 0xe5, 0xe3,
	0xaf, 0xe4, 0x05, 0x72, 0x85, 0x0c, 0xe4, 0x52, 0x80, 0x59, 0xdb, 0x0f, 0xb0, 0xa2, 0x2b, 0x06,
	0xad, 0x98, 0xf8, 0x0f, 0xd4, 0x8c, 0xb3, 0x5b, 0x0e, 0x5e, 0xee, 0xb8, 0x54, 0x69, 0x53, 0x60,
	0xf9, 0x1f, 0xad, 0xfa, 0xa5, 0x2f, 0x62, 0x21, 0x9f, 0x75, 0xe9, 0x31, 0x72, 0x24, 0x97, 0x41,
	0x52, 0xbd, 0x1e, 0x0d, 0xd4, 0x6e, 0x1e, 0xcc, 0x98, 0x80, 0x09, 0x7a, 0x54, 0x03, 0xd5, 0x60,
	0x12, 0x0b, 0x49, 0xba, 0xbb, 0x5d, 0xd1, 0x0d, 0x2a, 0xd6, 0xc5, 0x97, 0x26, 0xea, 0x69, 0xa7,
	0x58, 0x71, 0x17, 0x87, 0x0c, 0x5f, 0x23, 0xf7, 0x32, 0x8a, 0x5e, 0x65, 0x0c, 0x6f, 0x37, 0xb6,
	0x7e, 0xaa, 0x83, 0x33, 0x73, 0xd3, 0x95, 0xb6, 0x7f, 0x80, 0x27, 0x08, 0xe9, 0x38, 0x5e, 0x05,
	0x82, 0xe1, 0xb3, 0x1d, 0x52, 0x9a, 0x32, 0x20, 0x26, 0xb5, 0x4a, 0xe8, 0x1f, 0xd4, 0xa1, 0x2c,
	0xe5, 0x77, 0x6c, 0x47, 0x4b, 0x4b, 0x26, 0xa4, 0xa2, 0x7e, 0x57, 0xe4, 0x1a, 0xb9, 0x13, 0x26,
	0xf7, 0x29, 0xec, 0xfd, 0xa8, 0x6b, 0xe4, 0xde, 0xb0, 0xa0, 0x08, 0xe7, 0xcf, 0x94, 0x51, 0xf3,
	0xf7, 0x0f, 0xf0, 0x6b, 0x74, 0x38, 0x61, 0x52, 0x81, 0xab, 0xf2, 0xdf, 0x38, 0x91, 0xac, 0xc0,
	0xbf, 0x9a, 0xf5, 0x74, 0x75, 0xbf, 0xe2, 0x4b, 0xd4, 0x9b, 0x30, 0xa9, 0x42, 0x13, 0x0b, 0x19,
	0x87, 0xcf, 0x93, 0xab, 0x0e, 0xfb, 0x07, 0xf8, 0x7f, 0xd4, 0xaf, 0x86, 0xb7, 0xa7, 0xbf, 0xfd,
	0xf3, 0xbb, 0xf2, 0xde, 0x76, 0x6b, 0xff, 0x59, 0xa6, 0x2e, 0x80, 0x3f, 0xbf, 0x05, 0x00, 0x00,
	0xff, 0xff, 0xcc, 0xb2, 0xe9, 0xb2, 0xf1, 0x08, 0x00, 0x00,
}
