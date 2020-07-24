// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/categoryService.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type CategoryServiceBool struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CategoryServiceBool) Reset()         { *m = CategoryServiceBool{} }
func (m *CategoryServiceBool) String() string { return proto.CompactTextString(m) }
func (*CategoryServiceBool) ProtoMessage()    {}
func (*CategoryServiceBool) Descriptor() ([]byte, []int) {
	return fileDescriptor_938cd0f1643c580a, []int{0}
}

func (m *CategoryServiceBool) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CategoryServiceBool.Unmarshal(m, b)
}
func (m *CategoryServiceBool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CategoryServiceBool.Marshal(b, m, deterministic)
}
func (m *CategoryServiceBool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CategoryServiceBool.Merge(m, src)
}
func (m *CategoryServiceBool) XXX_Size() int {
	return xxx_messageInfo_CategoryServiceBool.Size(m)
}
func (m *CategoryServiceBool) XXX_DiscardUnknown() {
	xxx_messageInfo_CategoryServiceBool.DiscardUnknown(m)
}

var xxx_messageInfo_CategoryServiceBool proto.InternalMessageInfo

func (m *CategoryServiceBool) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type CategoryServiceCreateRq struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	ParentId             int64    `protobuf:"varint,2,opt,name=parentId,proto3" json:"parentId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CategoryServiceCreateRq) Reset()         { *m = CategoryServiceCreateRq{} }
func (m *CategoryServiceCreateRq) String() string { return proto.CompactTextString(m) }
func (*CategoryServiceCreateRq) ProtoMessage()    {}
func (*CategoryServiceCreateRq) Descriptor() ([]byte, []int) {
	return fileDescriptor_938cd0f1643c580a, []int{1}
}

func (m *CategoryServiceCreateRq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CategoryServiceCreateRq.Unmarshal(m, b)
}
func (m *CategoryServiceCreateRq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CategoryServiceCreateRq.Marshal(b, m, deterministic)
}
func (m *CategoryServiceCreateRq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CategoryServiceCreateRq.Merge(m, src)
}
func (m *CategoryServiceCreateRq) XXX_Size() int {
	return xxx_messageInfo_CategoryServiceCreateRq.Size(m)
}
func (m *CategoryServiceCreateRq) XXX_DiscardUnknown() {
	xxx_messageInfo_CategoryServiceCreateRq.DiscardUnknown(m)
}

var xxx_messageInfo_CategoryServiceCreateRq proto.InternalMessageInfo

func (m *CategoryServiceCreateRq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CategoryServiceCreateRq) GetParentId() int64 {
	if m != nil {
		return m.ParentId
	}
	return 0
}

type CategoryServiceModel struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ParentId             int64    `protobuf:"varint,3,opt,name=parentId,proto3" json:"parentId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CategoryServiceModel) Reset()         { *m = CategoryServiceModel{} }
func (m *CategoryServiceModel) String() string { return proto.CompactTextString(m) }
func (*CategoryServiceModel) ProtoMessage()    {}
func (*CategoryServiceModel) Descriptor() ([]byte, []int) {
	return fileDescriptor_938cd0f1643c580a, []int{2}
}

func (m *CategoryServiceModel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CategoryServiceModel.Unmarshal(m, b)
}
func (m *CategoryServiceModel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CategoryServiceModel.Marshal(b, m, deterministic)
}
func (m *CategoryServiceModel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CategoryServiceModel.Merge(m, src)
}
func (m *CategoryServiceModel) XXX_Size() int {
	return xxx_messageInfo_CategoryServiceModel.Size(m)
}
func (m *CategoryServiceModel) XXX_DiscardUnknown() {
	xxx_messageInfo_CategoryServiceModel.DiscardUnknown(m)
}

var xxx_messageInfo_CategoryServiceModel proto.InternalMessageInfo

func (m *CategoryServiceModel) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *CategoryServiceModel) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CategoryServiceModel) GetParentId() int64 {
	if m != nil {
		return m.ParentId
	}
	return 0
}

type CategoryServiceUpdateRq struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ParentId             int64    `protobuf:"varint,3,opt,name=parentId,proto3" json:"parentId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CategoryServiceUpdateRq) Reset()         { *m = CategoryServiceUpdateRq{} }
func (m *CategoryServiceUpdateRq) String() string { return proto.CompactTextString(m) }
func (*CategoryServiceUpdateRq) ProtoMessage()    {}
func (*CategoryServiceUpdateRq) Descriptor() ([]byte, []int) {
	return fileDescriptor_938cd0f1643c580a, []int{3}
}

func (m *CategoryServiceUpdateRq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CategoryServiceUpdateRq.Unmarshal(m, b)
}
func (m *CategoryServiceUpdateRq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CategoryServiceUpdateRq.Marshal(b, m, deterministic)
}
func (m *CategoryServiceUpdateRq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CategoryServiceUpdateRq.Merge(m, src)
}
func (m *CategoryServiceUpdateRq) XXX_Size() int {
	return xxx_messageInfo_CategoryServiceUpdateRq.Size(m)
}
func (m *CategoryServiceUpdateRq) XXX_DiscardUnknown() {
	xxx_messageInfo_CategoryServiceUpdateRq.DiscardUnknown(m)
}

var xxx_messageInfo_CategoryServiceUpdateRq proto.InternalMessageInfo

func (m *CategoryServiceUpdateRq) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *CategoryServiceUpdateRq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CategoryServiceUpdateRq) GetParentId() int64 {
	if m != nil {
		return m.ParentId
	}
	return 0
}

type CategoryServiceGetRq struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CategoryServiceGetRq) Reset()         { *m = CategoryServiceGetRq{} }
func (m *CategoryServiceGetRq) String() string { return proto.CompactTextString(m) }
func (*CategoryServiceGetRq) ProtoMessage()    {}
func (*CategoryServiceGetRq) Descriptor() ([]byte, []int) {
	return fileDescriptor_938cd0f1643c580a, []int{4}
}

func (m *CategoryServiceGetRq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CategoryServiceGetRq.Unmarshal(m, b)
}
func (m *CategoryServiceGetRq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CategoryServiceGetRq.Marshal(b, m, deterministic)
}
func (m *CategoryServiceGetRq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CategoryServiceGetRq.Merge(m, src)
}
func (m *CategoryServiceGetRq) XXX_Size() int {
	return xxx_messageInfo_CategoryServiceGetRq.Size(m)
}
func (m *CategoryServiceGetRq) XXX_DiscardUnknown() {
	xxx_messageInfo_CategoryServiceGetRq.DiscardUnknown(m)
}

var xxx_messageInfo_CategoryServiceGetRq proto.InternalMessageInfo

func (m *CategoryServiceGetRq) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type CategoryServiceGetRs struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ParentId             int64    `protobuf:"varint,3,opt,name=parentId,proto3" json:"parentId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CategoryServiceGetRs) Reset()         { *m = CategoryServiceGetRs{} }
func (m *CategoryServiceGetRs) String() string { return proto.CompactTextString(m) }
func (*CategoryServiceGetRs) ProtoMessage()    {}
func (*CategoryServiceGetRs) Descriptor() ([]byte, []int) {
	return fileDescriptor_938cd0f1643c580a, []int{5}
}

func (m *CategoryServiceGetRs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CategoryServiceGetRs.Unmarshal(m, b)
}
func (m *CategoryServiceGetRs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CategoryServiceGetRs.Marshal(b, m, deterministic)
}
func (m *CategoryServiceGetRs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CategoryServiceGetRs.Merge(m, src)
}
func (m *CategoryServiceGetRs) XXX_Size() int {
	return xxx_messageInfo_CategoryServiceGetRs.Size(m)
}
func (m *CategoryServiceGetRs) XXX_DiscardUnknown() {
	xxx_messageInfo_CategoryServiceGetRs.DiscardUnknown(m)
}

var xxx_messageInfo_CategoryServiceGetRs proto.InternalMessageInfo

func (m *CategoryServiceGetRs) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *CategoryServiceGetRs) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CategoryServiceGetRs) GetParentId() int64 {
	if m != nil {
		return m.ParentId
	}
	return 0
}

type CategoryServiceGetAllRq struct {
	Offset               string   `protobuf:"bytes,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit                string   `protobuf:"bytes,2,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CategoryServiceGetAllRq) Reset()         { *m = CategoryServiceGetAllRq{} }
func (m *CategoryServiceGetAllRq) String() string { return proto.CompactTextString(m) }
func (*CategoryServiceGetAllRq) ProtoMessage()    {}
func (*CategoryServiceGetAllRq) Descriptor() ([]byte, []int) {
	return fileDescriptor_938cd0f1643c580a, []int{6}
}

func (m *CategoryServiceGetAllRq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CategoryServiceGetAllRq.Unmarshal(m, b)
}
func (m *CategoryServiceGetAllRq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CategoryServiceGetAllRq.Marshal(b, m, deterministic)
}
func (m *CategoryServiceGetAllRq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CategoryServiceGetAllRq.Merge(m, src)
}
func (m *CategoryServiceGetAllRq) XXX_Size() int {
	return xxx_messageInfo_CategoryServiceGetAllRq.Size(m)
}
func (m *CategoryServiceGetAllRq) XXX_DiscardUnknown() {
	xxx_messageInfo_CategoryServiceGetAllRq.DiscardUnknown(m)
}

var xxx_messageInfo_CategoryServiceGetAllRq proto.InternalMessageInfo

func (m *CategoryServiceGetAllRq) GetOffset() string {
	if m != nil {
		return m.Offset
	}
	return ""
}

func (m *CategoryServiceGetAllRq) GetLimit() string {
	if m != nil {
		return m.Limit
	}
	return ""
}

type CategoryServiceGetAllRs struct {
	Result               []*CategoryServiceModel `protobuf:"bytes,1,rep,name=result,proto3" json:"result,omitempty"`
	Count                int64                   `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *CategoryServiceGetAllRs) Reset()         { *m = CategoryServiceGetAllRs{} }
func (m *CategoryServiceGetAllRs) String() string { return proto.CompactTextString(m) }
func (*CategoryServiceGetAllRs) ProtoMessage()    {}
func (*CategoryServiceGetAllRs) Descriptor() ([]byte, []int) {
	return fileDescriptor_938cd0f1643c580a, []int{7}
}

func (m *CategoryServiceGetAllRs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CategoryServiceGetAllRs.Unmarshal(m, b)
}
func (m *CategoryServiceGetAllRs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CategoryServiceGetAllRs.Marshal(b, m, deterministic)
}
func (m *CategoryServiceGetAllRs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CategoryServiceGetAllRs.Merge(m, src)
}
func (m *CategoryServiceGetAllRs) XXX_Size() int {
	return xxx_messageInfo_CategoryServiceGetAllRs.Size(m)
}
func (m *CategoryServiceGetAllRs) XXX_DiscardUnknown() {
	xxx_messageInfo_CategoryServiceGetAllRs.DiscardUnknown(m)
}

var xxx_messageInfo_CategoryServiceGetAllRs proto.InternalMessageInfo

func (m *CategoryServiceGetAllRs) GetResult() []*CategoryServiceModel {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *CategoryServiceGetAllRs) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type CategoryServiceDeleteRq struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CategoryServiceDeleteRq) Reset()         { *m = CategoryServiceDeleteRq{} }
func (m *CategoryServiceDeleteRq) String() string { return proto.CompactTextString(m) }
func (*CategoryServiceDeleteRq) ProtoMessage()    {}
func (*CategoryServiceDeleteRq) Descriptor() ([]byte, []int) {
	return fileDescriptor_938cd0f1643c580a, []int{8}
}

func (m *CategoryServiceDeleteRq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CategoryServiceDeleteRq.Unmarshal(m, b)
}
func (m *CategoryServiceDeleteRq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CategoryServiceDeleteRq.Marshal(b, m, deterministic)
}
func (m *CategoryServiceDeleteRq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CategoryServiceDeleteRq.Merge(m, src)
}
func (m *CategoryServiceDeleteRq) XXX_Size() int {
	return xxx_messageInfo_CategoryServiceDeleteRq.Size(m)
}
func (m *CategoryServiceDeleteRq) XXX_DiscardUnknown() {
	xxx_messageInfo_CategoryServiceDeleteRq.DiscardUnknown(m)
}

var xxx_messageInfo_CategoryServiceDeleteRq proto.InternalMessageInfo

func (m *CategoryServiceDeleteRq) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func init() {
	proto.RegisterType((*CategoryServiceBool)(nil), "proto.CategoryServiceBool")
	proto.RegisterType((*CategoryServiceCreateRq)(nil), "proto.CategoryServiceCreateRq")
	proto.RegisterType((*CategoryServiceModel)(nil), "proto.CategoryServiceModel")
	proto.RegisterType((*CategoryServiceUpdateRq)(nil), "proto.CategoryServiceUpdateRq")
	proto.RegisterType((*CategoryServiceGetRq)(nil), "proto.CategoryServiceGetRq")
	proto.RegisterType((*CategoryServiceGetRs)(nil), "proto.CategoryServiceGetRs")
	proto.RegisterType((*CategoryServiceGetAllRq)(nil), "proto.CategoryServiceGetAllRq")
	proto.RegisterType((*CategoryServiceGetAllRs)(nil), "proto.CategoryServiceGetAllRs")
	proto.RegisterType((*CategoryServiceDeleteRq)(nil), "proto.CategoryServiceDeleteRq")
}

func init() {
	proto.RegisterFile("proto/categoryService.proto", fileDescriptor_938cd0f1643c580a)
}

var fileDescriptor_938cd0f1643c580a = []byte{
	// 356 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0xcd, 0x4a, 0xc3, 0x40,
	0x10, 0xc7, 0x69, 0x62, 0x63, 0x1d, 0x41, 0x61, 0x2c, 0x5a, 0x52, 0x90, 0x92, 0x83, 0xd4, 0x4b,
	0x0b, 0xed, 0x03, 0x88, 0xb6, 0x10, 0x7b, 0xf0, 0xb2, 0xa2, 0xe0, 0x31, 0x26, 0x53, 0x09, 0x6c,
	0xbb, 0x6d, 0x76, 0x23, 0xf8, 0xd2, 0x3e, 0x83, 0x64, 0x37, 0x29, 0x98, 0xaf, 0x4b, 0x4f, 0xed,
	0x64, 0x66, 0x7f, 0xf3, 0xf1, 0xff, 0xc3, 0x70, 0x97, 0x08, 0x25, 0xa6, 0x61, 0xa0, 0xe8, 0x4b,
	0x24, 0x3f, 0xaf, 0x94, 0x7c, 0xc7, 0x21, 0x4d, 0xf4, 0x57, 0xec, 0xea, 0x1f, 0x6f, 0x0a, 0x57,
	0x8b, 0xff, 0xf9, 0x27, 0x21, 0x38, 0x0e, 0xe0, 0x54, 0xa6, 0x61, 0x48, 0x52, 0x0e, 0x3a, 0xa3,
	0xce, 0xb8, 0xc7, 0x8a, 0xd0, 0x5b, 0xc1, 0x4d, 0xe9, 0xc1, 0x22, 0xa1, 0x40, 0x11, 0xdb, 0x23,
	0xc2, 0xc9, 0x36, 0xd8, 0x90, 0x7e, 0x71, 0xc6, 0xf4, 0x7f, 0x74, 0xa1, 0xb7, 0x0b, 0x12, 0xda,
	0xaa, 0x55, 0x34, 0xb0, 0x46, 0x9d, 0xb1, 0xcd, 0x0e, 0xb1, 0xf7, 0x0e, 0xfd, 0x12, 0xea, 0x45,
	0x44, 0xc4, 0xf1, 0x02, 0xac, 0x38, 0xd2, 0x14, 0x9b, 0x59, 0x71, 0x74, 0xe0, 0x5a, 0x0d, 0x5c,
	0xbb, 0xc4, 0xfd, 0xa8, 0x8c, 0xf8, 0xb6, 0x8b, 0xcc, 0x88, 0xc7, 0xa2, 0xef, 0x2a, 0x23, 0xfb,
	0xa4, 0xaa, 0xdc, 0x9a, 0xd5, 0xb2, 0x3a, 0x79, 0x74, 0x7f, 0xbf, 0xb2, 0x9a, 0x4f, 0xea, 0x91,
	0x73, 0xb6, 0xc7, 0x6b, 0x70, 0xc4, 0x7a, 0x2d, 0x49, 0xe5, 0xf7, 0xcf, 0x23, 0xec, 0x43, 0x97,
	0xc7, 0x9b, 0x58, 0xe5, 0x3d, 0x4c, 0xe0, 0x45, 0x4d, 0x20, 0x89, 0x73, 0x70, 0x12, 0x92, 0x29,
	0xcf, 0x40, 0xf6, 0xf8, 0x7c, 0x36, 0x34, 0x8e, 0x99, 0xd4, 0x69, 0xc5, 0xf2, 0xd2, 0xac, 0x4b,
	0x28, 0xd2, 0xad, 0xca, 0x45, 0x36, 0x81, 0x77, 0x5f, 0xe9, 0xb2, 0x24, 0x4e, 0x75, 0x4a, 0xcc,
	0x7e, 0x2d, 0xb8, 0x2c, 0xd5, 0xe2, 0x12, 0x1c, 0x63, 0x2e, 0xbc, 0xad, 0x9f, 0xa1, 0xb0, 0x9e,
	0xeb, 0xd6, 0xe7, 0xb5, 0x97, 0x97, 0xe0, 0x18, 0xfd, 0x9b, 0x28, 0x85, 0x3b, 0x5a, 0x29, 0x0f,
	0x60, 0xfb, 0xa4, 0xb0, 0xe1, 0x18, 0xda, 0x05, 0x6e, 0x4b, 0x52, 0xe2, 0x33, 0x38, 0xe6, 0xc4,
	0x4d, 0x63, 0x14, 0x4a, 0xba, 0xed, 0x79, 0x99, 0x2d, 0x64, 0xce, 0xd8, 0x44, 0x2a, 0x8e, 0xdc,
	0xb6, 0xd0, 0xa7, 0xa3, 0x53, 0xf3, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xee, 0xfb, 0x69, 0xbf,
	0x26, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CategoryServiceClient is the client API for CategoryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CategoryServiceClient interface {
	Create(ctx context.Context, in *CategoryServiceCreateRq, opts ...grpc.CallOption) (*CategoryServiceBool, error)
	Update(ctx context.Context, in *CategoryServiceUpdateRq, opts ...grpc.CallOption) (*CategoryServiceBool, error)
	Get(ctx context.Context, in *CategoryServiceGetRq, opts ...grpc.CallOption) (*CategoryServiceGetRs, error)
	GetAll(ctx context.Context, in *CategoryServiceGetAllRq, opts ...grpc.CallOption) (*CategoryServiceGetAllRs, error)
	Delete(ctx context.Context, in *CategoryServiceDeleteRq, opts ...grpc.CallOption) (*CategoryServiceBool, error)
}

type categoryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCategoryServiceClient(cc grpc.ClientConnInterface) CategoryServiceClient {
	return &categoryServiceClient{cc}
}

func (c *categoryServiceClient) Create(ctx context.Context, in *CategoryServiceCreateRq, opts ...grpc.CallOption) (*CategoryServiceBool, error) {
	out := new(CategoryServiceBool)
	err := c.cc.Invoke(ctx, "/proto.CategoryService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryServiceClient) Update(ctx context.Context, in *CategoryServiceUpdateRq, opts ...grpc.CallOption) (*CategoryServiceBool, error) {
	out := new(CategoryServiceBool)
	err := c.cc.Invoke(ctx, "/proto.CategoryService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryServiceClient) Get(ctx context.Context, in *CategoryServiceGetRq, opts ...grpc.CallOption) (*CategoryServiceGetRs, error) {
	out := new(CategoryServiceGetRs)
	err := c.cc.Invoke(ctx, "/proto.CategoryService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryServiceClient) GetAll(ctx context.Context, in *CategoryServiceGetAllRq, opts ...grpc.CallOption) (*CategoryServiceGetAllRs, error) {
	out := new(CategoryServiceGetAllRs)
	err := c.cc.Invoke(ctx, "/proto.CategoryService/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryServiceClient) Delete(ctx context.Context, in *CategoryServiceDeleteRq, opts ...grpc.CallOption) (*CategoryServiceBool, error) {
	out := new(CategoryServiceBool)
	err := c.cc.Invoke(ctx, "/proto.CategoryService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CategoryServiceServer is the server API for CategoryService service.
type CategoryServiceServer interface {
	Create(context.Context, *CategoryServiceCreateRq) (*CategoryServiceBool, error)
	Update(context.Context, *CategoryServiceUpdateRq) (*CategoryServiceBool, error)
	Get(context.Context, *CategoryServiceGetRq) (*CategoryServiceGetRs, error)
	GetAll(context.Context, *CategoryServiceGetAllRq) (*CategoryServiceGetAllRs, error)
	Delete(context.Context, *CategoryServiceDeleteRq) (*CategoryServiceBool, error)
}

// UnimplementedCategoryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCategoryServiceServer struct {
}

func (*UnimplementedCategoryServiceServer) Create(ctx context.Context, req *CategoryServiceCreateRq) (*CategoryServiceBool, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedCategoryServiceServer) Update(ctx context.Context, req *CategoryServiceUpdateRq) (*CategoryServiceBool, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedCategoryServiceServer) Get(ctx context.Context, req *CategoryServiceGetRq) (*CategoryServiceGetRs, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedCategoryServiceServer) GetAll(ctx context.Context, req *CategoryServiceGetAllRq) (*CategoryServiceGetAllRs, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (*UnimplementedCategoryServiceServer) Delete(ctx context.Context, req *CategoryServiceDeleteRq) (*CategoryServiceBool, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterCategoryServiceServer(s *grpc.Server, srv CategoryServiceServer) {
	s.RegisterService(&_CategoryService_serviceDesc, srv)
}

func _CategoryService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CategoryServiceCreateRq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CategoryService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryServiceServer).Create(ctx, req.(*CategoryServiceCreateRq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CategoryService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CategoryServiceUpdateRq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CategoryService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryServiceServer).Update(ctx, req.(*CategoryServiceUpdateRq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CategoryService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CategoryServiceGetRq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CategoryService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryServiceServer).Get(ctx, req.(*CategoryServiceGetRq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CategoryService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CategoryServiceGetAllRq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CategoryService/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryServiceServer).GetAll(ctx, req.(*CategoryServiceGetAllRq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CategoryService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CategoryServiceDeleteRq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CategoryService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryServiceServer).Delete(ctx, req.(*CategoryServiceDeleteRq))
	}
	return interceptor(ctx, in, info, handler)
}

var _CategoryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.CategoryService",
	HandlerType: (*CategoryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _CategoryService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _CategoryService_Update_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _CategoryService_Get_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _CategoryService_GetAll_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _CategoryService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/categoryService.proto",
}