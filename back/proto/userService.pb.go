// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/userService.proto

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

type UserServiceBool struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserServiceBool) Reset()         { *m = UserServiceBool{} }
func (m *UserServiceBool) String() string { return proto.CompactTextString(m) }
func (*UserServiceBool) ProtoMessage()    {}
func (*UserServiceBool) Descriptor() ([]byte, []int) {
	return fileDescriptor_ba03f2599437003f, []int{0}
}

func (m *UserServiceBool) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserServiceBool.Unmarshal(m, b)
}
func (m *UserServiceBool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserServiceBool.Marshal(b, m, deterministic)
}
func (m *UserServiceBool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserServiceBool.Merge(m, src)
}
func (m *UserServiceBool) XXX_Size() int {
	return xxx_messageInfo_UserServiceBool.Size(m)
}
func (m *UserServiceBool) XXX_DiscardUnknown() {
	xxx_messageInfo_UserServiceBool.DiscardUnknown(m)
}

var xxx_messageInfo_UserServiceBool proto.InternalMessageInfo

func (m *UserServiceBool) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type UserSocial struct {
	Icon                 string   `protobuf:"bytes,1,opt,name=icon,proto3" json:"icon,omitempty"`
	Url                  string   `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserSocial) Reset()         { *m = UserSocial{} }
func (m *UserSocial) String() string { return proto.CompactTextString(m) }
func (*UserSocial) ProtoMessage()    {}
func (*UserSocial) Descriptor() ([]byte, []int) {
	return fileDescriptor_ba03f2599437003f, []int{1}
}

func (m *UserSocial) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserSocial.Unmarshal(m, b)
}
func (m *UserSocial) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserSocial.Marshal(b, m, deterministic)
}
func (m *UserSocial) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserSocial.Merge(m, src)
}
func (m *UserSocial) XXX_Size() int {
	return xxx_messageInfo_UserSocial.Size(m)
}
func (m *UserSocial) XXX_DiscardUnknown() {
	xxx_messageInfo_UserSocial.DiscardUnknown(m)
}

var xxx_messageInfo_UserSocial proto.InternalMessageInfo

func (m *UserSocial) GetIcon() string {
	if m != nil {
		return m.Icon
	}
	return ""
}

func (m *UserSocial) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

type UserServiceCreateRq struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Active               bool     `protobuf:"varint,3,opt,name=active,proto3" json:"active,omitempty"`
	Name                 string   `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserServiceCreateRq) Reset()         { *m = UserServiceCreateRq{} }
func (m *UserServiceCreateRq) String() string { return proto.CompactTextString(m) }
func (*UserServiceCreateRq) ProtoMessage()    {}
func (*UserServiceCreateRq) Descriptor() ([]byte, []int) {
	return fileDescriptor_ba03f2599437003f, []int{2}
}

func (m *UserServiceCreateRq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserServiceCreateRq.Unmarshal(m, b)
}
func (m *UserServiceCreateRq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserServiceCreateRq.Marshal(b, m, deterministic)
}
func (m *UserServiceCreateRq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserServiceCreateRq.Merge(m, src)
}
func (m *UserServiceCreateRq) XXX_Size() int {
	return xxx_messageInfo_UserServiceCreateRq.Size(m)
}
func (m *UserServiceCreateRq) XXX_DiscardUnknown() {
	xxx_messageInfo_UserServiceCreateRq.DiscardUnknown(m)
}

var xxx_messageInfo_UserServiceCreateRq proto.InternalMessageInfo

func (m *UserServiceCreateRq) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UserServiceCreateRq) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *UserServiceCreateRq) GetActive() bool {
	if m != nil {
		return m.Active
	}
	return false
}

func (m *UserServiceCreateRq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type UserServiceModel struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Email                string   `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	CreatedAt            string   `protobuf:"bytes,4,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,5,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	Active               bool     `protobuf:"varint,6,opt,name=active,proto3" json:"active,omitempty"`
	Name                 string   `protobuf:"bytes,7,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserServiceModel) Reset()         { *m = UserServiceModel{} }
func (m *UserServiceModel) String() string { return proto.CompactTextString(m) }
func (*UserServiceModel) ProtoMessage()    {}
func (*UserServiceModel) Descriptor() ([]byte, []int) {
	return fileDescriptor_ba03f2599437003f, []int{3}
}

func (m *UserServiceModel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserServiceModel.Unmarshal(m, b)
}
func (m *UserServiceModel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserServiceModel.Marshal(b, m, deterministic)
}
func (m *UserServiceModel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserServiceModel.Merge(m, src)
}
func (m *UserServiceModel) XXX_Size() int {
	return xxx_messageInfo_UserServiceModel.Size(m)
}
func (m *UserServiceModel) XXX_DiscardUnknown() {
	xxx_messageInfo_UserServiceModel.DiscardUnknown(m)
}

var xxx_messageInfo_UserServiceModel proto.InternalMessageInfo

func (m *UserServiceModel) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UserServiceModel) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UserServiceModel) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *UserServiceModel) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *UserServiceModel) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func (m *UserServiceModel) GetActive() bool {
	if m != nil {
		return m.Active
	}
	return false
}

func (m *UserServiceModel) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type UserServiceUpdateRq struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Active               bool     `protobuf:"varint,3,opt,name=active,proto3" json:"active,omitempty"`
	Name                 string   `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Id                   int64    `protobuf:"varint,5,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserServiceUpdateRq) Reset()         { *m = UserServiceUpdateRq{} }
func (m *UserServiceUpdateRq) String() string { return proto.CompactTextString(m) }
func (*UserServiceUpdateRq) ProtoMessage()    {}
func (*UserServiceUpdateRq) Descriptor() ([]byte, []int) {
	return fileDescriptor_ba03f2599437003f, []int{4}
}

func (m *UserServiceUpdateRq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserServiceUpdateRq.Unmarshal(m, b)
}
func (m *UserServiceUpdateRq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserServiceUpdateRq.Marshal(b, m, deterministic)
}
func (m *UserServiceUpdateRq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserServiceUpdateRq.Merge(m, src)
}
func (m *UserServiceUpdateRq) XXX_Size() int {
	return xxx_messageInfo_UserServiceUpdateRq.Size(m)
}
func (m *UserServiceUpdateRq) XXX_DiscardUnknown() {
	xxx_messageInfo_UserServiceUpdateRq.DiscardUnknown(m)
}

var xxx_messageInfo_UserServiceUpdateRq proto.InternalMessageInfo

func (m *UserServiceUpdateRq) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UserServiceUpdateRq) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *UserServiceUpdateRq) GetActive() bool {
	if m != nil {
		return m.Active
	}
	return false
}

func (m *UserServiceUpdateRq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserServiceUpdateRq) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type UserServiceGetRq struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserServiceGetRq) Reset()         { *m = UserServiceGetRq{} }
func (m *UserServiceGetRq) String() string { return proto.CompactTextString(m) }
func (*UserServiceGetRq) ProtoMessage()    {}
func (*UserServiceGetRq) Descriptor() ([]byte, []int) {
	return fileDescriptor_ba03f2599437003f, []int{5}
}

func (m *UserServiceGetRq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserServiceGetRq.Unmarshal(m, b)
}
func (m *UserServiceGetRq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserServiceGetRq.Marshal(b, m, deterministic)
}
func (m *UserServiceGetRq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserServiceGetRq.Merge(m, src)
}
func (m *UserServiceGetRq) XXX_Size() int {
	return xxx_messageInfo_UserServiceGetRq.Size(m)
}
func (m *UserServiceGetRq) XXX_DiscardUnknown() {
	xxx_messageInfo_UserServiceGetRq.DiscardUnknown(m)
}

var xxx_messageInfo_UserServiceGetRq proto.InternalMessageInfo

func (m *UserServiceGetRq) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type UserServiceLoginRq struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserServiceLoginRq) Reset()         { *m = UserServiceLoginRq{} }
func (m *UserServiceLoginRq) String() string { return proto.CompactTextString(m) }
func (*UserServiceLoginRq) ProtoMessage()    {}
func (*UserServiceLoginRq) Descriptor() ([]byte, []int) {
	return fileDescriptor_ba03f2599437003f, []int{6}
}

func (m *UserServiceLoginRq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserServiceLoginRq.Unmarshal(m, b)
}
func (m *UserServiceLoginRq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserServiceLoginRq.Marshal(b, m, deterministic)
}
func (m *UserServiceLoginRq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserServiceLoginRq.Merge(m, src)
}
func (m *UserServiceLoginRq) XXX_Size() int {
	return xxx_messageInfo_UserServiceLoginRq.Size(m)
}
func (m *UserServiceLoginRq) XXX_DiscardUnknown() {
	xxx_messageInfo_UserServiceLoginRq.DiscardUnknown(m)
}

var xxx_messageInfo_UserServiceLoginRq proto.InternalMessageInfo

func (m *UserServiceLoginRq) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UserServiceLoginRq) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type UserServiceLoginRs struct {
	Token                string            `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	User                 *UserServiceModel `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *UserServiceLoginRs) Reset()         { *m = UserServiceLoginRs{} }
func (m *UserServiceLoginRs) String() string { return proto.CompactTextString(m) }
func (*UserServiceLoginRs) ProtoMessage()    {}
func (*UserServiceLoginRs) Descriptor() ([]byte, []int) {
	return fileDescriptor_ba03f2599437003f, []int{7}
}

func (m *UserServiceLoginRs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserServiceLoginRs.Unmarshal(m, b)
}
func (m *UserServiceLoginRs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserServiceLoginRs.Marshal(b, m, deterministic)
}
func (m *UserServiceLoginRs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserServiceLoginRs.Merge(m, src)
}
func (m *UserServiceLoginRs) XXX_Size() int {
	return xxx_messageInfo_UserServiceLoginRs.Size(m)
}
func (m *UserServiceLoginRs) XXX_DiscardUnknown() {
	xxx_messageInfo_UserServiceLoginRs.DiscardUnknown(m)
}

var xxx_messageInfo_UserServiceLoginRs proto.InternalMessageInfo

func (m *UserServiceLoginRs) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *UserServiceLoginRs) GetUser() *UserServiceModel {
	if m != nil {
		return m.User
	}
	return nil
}

type UserServiceGetRs struct {
	Id                   int64         `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Email                string        `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	UpdatedAt            string        `protobuf:"bytes,3,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	Name                 string        `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Description          string        `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	Socials              []*UserSocial `protobuf:"bytes,6,rep,name=socials,proto3" json:"socials,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *UserServiceGetRs) Reset()         { *m = UserServiceGetRs{} }
func (m *UserServiceGetRs) String() string { return proto.CompactTextString(m) }
func (*UserServiceGetRs) ProtoMessage()    {}
func (*UserServiceGetRs) Descriptor() ([]byte, []int) {
	return fileDescriptor_ba03f2599437003f, []int{8}
}

func (m *UserServiceGetRs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserServiceGetRs.Unmarshal(m, b)
}
func (m *UserServiceGetRs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserServiceGetRs.Marshal(b, m, deterministic)
}
func (m *UserServiceGetRs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserServiceGetRs.Merge(m, src)
}
func (m *UserServiceGetRs) XXX_Size() int {
	return xxx_messageInfo_UserServiceGetRs.Size(m)
}
func (m *UserServiceGetRs) XXX_DiscardUnknown() {
	xxx_messageInfo_UserServiceGetRs.DiscardUnknown(m)
}

var xxx_messageInfo_UserServiceGetRs proto.InternalMessageInfo

func (m *UserServiceGetRs) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UserServiceGetRs) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UserServiceGetRs) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func (m *UserServiceGetRs) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserServiceGetRs) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *UserServiceGetRs) GetSocials() []*UserSocial {
	if m != nil {
		return m.Socials
	}
	return nil
}

type UserServiceGetAllRq struct {
	Offset               string   `protobuf:"bytes,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit                string   `protobuf:"bytes,2,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserServiceGetAllRq) Reset()         { *m = UserServiceGetAllRq{} }
func (m *UserServiceGetAllRq) String() string { return proto.CompactTextString(m) }
func (*UserServiceGetAllRq) ProtoMessage()    {}
func (*UserServiceGetAllRq) Descriptor() ([]byte, []int) {
	return fileDescriptor_ba03f2599437003f, []int{9}
}

func (m *UserServiceGetAllRq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserServiceGetAllRq.Unmarshal(m, b)
}
func (m *UserServiceGetAllRq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserServiceGetAllRq.Marshal(b, m, deterministic)
}
func (m *UserServiceGetAllRq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserServiceGetAllRq.Merge(m, src)
}
func (m *UserServiceGetAllRq) XXX_Size() int {
	return xxx_messageInfo_UserServiceGetAllRq.Size(m)
}
func (m *UserServiceGetAllRq) XXX_DiscardUnknown() {
	xxx_messageInfo_UserServiceGetAllRq.DiscardUnknown(m)
}

var xxx_messageInfo_UserServiceGetAllRq proto.InternalMessageInfo

func (m *UserServiceGetAllRq) GetOffset() string {
	if m != nil {
		return m.Offset
	}
	return ""
}

func (m *UserServiceGetAllRq) GetLimit() string {
	if m != nil {
		return m.Limit
	}
	return ""
}

type UserServiceGetAllRs struct {
	Result               []*UserServiceModel `protobuf:"bytes,1,rep,name=result,proto3" json:"result,omitempty"`
	Count                int64               `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *UserServiceGetAllRs) Reset()         { *m = UserServiceGetAllRs{} }
func (m *UserServiceGetAllRs) String() string { return proto.CompactTextString(m) }
func (*UserServiceGetAllRs) ProtoMessage()    {}
func (*UserServiceGetAllRs) Descriptor() ([]byte, []int) {
	return fileDescriptor_ba03f2599437003f, []int{10}
}

func (m *UserServiceGetAllRs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserServiceGetAllRs.Unmarshal(m, b)
}
func (m *UserServiceGetAllRs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserServiceGetAllRs.Marshal(b, m, deterministic)
}
func (m *UserServiceGetAllRs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserServiceGetAllRs.Merge(m, src)
}
func (m *UserServiceGetAllRs) XXX_Size() int {
	return xxx_messageInfo_UserServiceGetAllRs.Size(m)
}
func (m *UserServiceGetAllRs) XXX_DiscardUnknown() {
	xxx_messageInfo_UserServiceGetAllRs.DiscardUnknown(m)
}

var xxx_messageInfo_UserServiceGetAllRs proto.InternalMessageInfo

func (m *UserServiceGetAllRs) GetResult() []*UserServiceModel {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *UserServiceGetAllRs) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type UserServiceDeleteRq struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserServiceDeleteRq) Reset()         { *m = UserServiceDeleteRq{} }
func (m *UserServiceDeleteRq) String() string { return proto.CompactTextString(m) }
func (*UserServiceDeleteRq) ProtoMessage()    {}
func (*UserServiceDeleteRq) Descriptor() ([]byte, []int) {
	return fileDescriptor_ba03f2599437003f, []int{11}
}

func (m *UserServiceDeleteRq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserServiceDeleteRq.Unmarshal(m, b)
}
func (m *UserServiceDeleteRq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserServiceDeleteRq.Marshal(b, m, deterministic)
}
func (m *UserServiceDeleteRq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserServiceDeleteRq.Merge(m, src)
}
func (m *UserServiceDeleteRq) XXX_Size() int {
	return xxx_messageInfo_UserServiceDeleteRq.Size(m)
}
func (m *UserServiceDeleteRq) XXX_DiscardUnknown() {
	xxx_messageInfo_UserServiceDeleteRq.DiscardUnknown(m)
}

var xxx_messageInfo_UserServiceDeleteRq proto.InternalMessageInfo

func (m *UserServiceDeleteRq) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func init() {
	proto.RegisterType((*UserServiceBool)(nil), "proto.UserServiceBool")
	proto.RegisterType((*UserSocial)(nil), "proto.UserSocial")
	proto.RegisterType((*UserServiceCreateRq)(nil), "proto.UserServiceCreateRq")
	proto.RegisterType((*UserServiceModel)(nil), "proto.UserServiceModel")
	proto.RegisterType((*UserServiceUpdateRq)(nil), "proto.UserServiceUpdateRq")
	proto.RegisterType((*UserServiceGetRq)(nil), "proto.UserServiceGetRq")
	proto.RegisterType((*UserServiceLoginRq)(nil), "proto.UserServiceLoginRq")
	proto.RegisterType((*UserServiceLoginRs)(nil), "proto.UserServiceLoginRs")
	proto.RegisterType((*UserServiceGetRs)(nil), "proto.UserServiceGetRs")
	proto.RegisterType((*UserServiceGetAllRq)(nil), "proto.UserServiceGetAllRq")
	proto.RegisterType((*UserServiceGetAllRs)(nil), "proto.UserServiceGetAllRs")
	proto.RegisterType((*UserServiceDeleteRq)(nil), "proto.UserServiceDeleteRq")
}

func init() {
	proto.RegisterFile("proto/userService.proto", fileDescriptor_ba03f2599437003f)
}

var fileDescriptor_ba03f2599437003f = []byte{
	// 539 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0xcf, 0x8b, 0xd4, 0x30,
	0x14, 0x66, 0x26, 0xd3, 0xce, 0xee, 0x1b, 0xd0, 0x35, 0xca, 0x6e, 0x2d, 0x1e, 0x86, 0x80, 0xb0,
	0x30, 0xb0, 0x0b, 0xe3, 0xc1, 0xcb, 0x0a, 0xae, 0x2b, 0xce, 0x45, 0x2f, 0x95, 0xc5, 0x8b, 0x97,
	0xda, 0x66, 0x25, 0x98, 0x69, 0x66, 0x9a, 0x74, 0x3d, 0x8b, 0x7f, 0x95, 0xe0, 0x1f, 0x27, 0x79,
	0x49, 0x77, 0xda, 0x4e, 0x2b, 0x22, 0xec, 0xa9, 0x7d, 0x3f, 0xf2, 0xe5, 0x7b, 0x5f, 0xbe, 0x04,
	0x4e, 0x36, 0xa5, 0x32, 0xea, 0xbc, 0xd2, 0xbc, 0xfc, 0xc8, 0xcb, 0x5b, 0x91, 0xf1, 0x33, 0xcc,
	0xd0, 0x00, 0x3f, 0x6c, 0x01, 0x0f, 0xaf, 0x77, 0xb5, 0x37, 0x4a, 0x49, 0x1a, 0xc1, 0x54, 0x57,
	0x59, 0xc6, 0xb5, 0x8e, 0x46, 0xf3, 0xd1, 0xe9, 0x41, 0x52, 0x87, 0x6c, 0x09, 0x80, 0xcd, 0x2a,
	0x13, 0xa9, 0xa4, 0x14, 0x26, 0x22, 0x53, 0x05, 0x36, 0x1d, 0x26, 0xf8, 0x4f, 0x8f, 0x80, 0x54,
	0xa5, 0x8c, 0xc6, 0x98, 0xb2, 0xbf, 0x4c, 0xc3, 0xe3, 0xc6, 0x06, 0x57, 0x25, 0x4f, 0x0d, 0x4f,
	0xb6, 0xf4, 0x09, 0x04, 0x7c, 0x9d, 0x0a, 0xe9, 0x57, 0xbb, 0x80, 0xc6, 0x70, 0xb0, 0x49, 0xb5,
	0xfe, 0xae, 0xca, 0xdc, 0x63, 0xdc, 0xc5, 0xf4, 0x18, 0xc2, 0x34, 0x33, 0xe2, 0x96, 0x47, 0x04,
	0x59, 0xf9, 0xc8, 0xd2, 0x28, 0xd2, 0x35, 0x8f, 0x26, 0x8e, 0x86, 0xfd, 0x67, 0xbf, 0x47, 0x70,
	0xd4, 0xd8, 0xf5, 0x83, 0xca, 0xb9, 0xa4, 0x0f, 0x60, 0x2c, 0x72, 0xdc, 0x8f, 0x24, 0x63, 0x91,
	0xef, 0x28, 0x8c, 0x87, 0x28, 0x90, 0x0e, 0x85, 0x67, 0x70, 0x98, 0xe1, 0x00, 0xf9, 0xa5, 0xf1,
	0xfb, 0xed, 0x12, 0xb6, 0x5a, 0x6d, 0x72, 0x5f, 0x0d, 0x5c, 0xf5, 0x2e, 0xd1, 0xa0, 0x1f, 0xf6,
	0xd2, 0x9f, 0x36, 0xe8, 0xff, 0x1c, 0xb5, 0x44, 0xbb, 0x46, 0x90, 0xfb, 0x16, 0xcd, 0xeb, 0x13,
	0xd4, 0xfa, 0x30, 0xd6, 0xd2, 0x70, 0xc5, 0x4d, 0xb2, 0xed, 0x6a, 0xc8, 0xde, 0x01, 0x6d, 0xf4,
	0xbc, 0x57, 0x5f, 0x45, 0xf1, 0x3f, 0x3c, 0xd9, 0xa7, 0x1e, 0x1c, 0x6d, 0x71, 0x8c, 0xfa, 0xc6,
	0x6b, 0x8b, 0xb9, 0x80, 0x2e, 0x60, 0x62, 0xed, 0x8c, 0x18, 0xb3, 0xe5, 0x89, 0xf3, 0xf3, 0x59,
	0xf7, 0xb8, 0x13, 0x6c, 0x62, 0xbf, 0x46, 0x7b, 0x53, 0xe8, 0x7f, 0x74, 0x42, 0xeb, 0x3c, 0x49,
	0xf7, 0x3c, 0xfb, 0x14, 0x9c, 0xc3, 0x2c, 0xe7, 0x3a, 0x2b, 0xc5, 0xc6, 0x08, 0x55, 0x78, 0x0f,
	0x34, 0x53, 0x74, 0x01, 0x53, 0x8d, 0xb7, 0x47, 0x47, 0xe1, 0x9c, 0x9c, 0xce, 0x96, 0x8f, 0x9a,
	0xf4, 0xb1, 0x92, 0xd4, 0x1d, 0xec, 0xaa, 0xe5, 0x82, 0x15, 0x37, 0x97, 0x52, 0x26, 0x5b, 0x7b,
	0xa6, 0xea, 0xe6, 0x46, 0x73, 0xe3, 0x65, 0xf1, 0x91, 0x9d, 0x42, 0x8a, 0xb5, 0x30, 0xf5, 0x14,
	0x18, 0xb0, 0xcf, 0x7d, 0x20, 0x9a, 0x9e, 0x43, 0x58, 0x72, 0x5d, 0x49, 0x0b, 0x42, 0xfe, 0x26,
	0xa3, 0x6f, 0xb3, 0xe8, 0x99, 0xaa, 0x0a, 0x87, 0x4e, 0x12, 0x17, 0xb0, 0xe7, 0x2d, 0xf4, 0xb7,
	0x5c, 0x72, 0x34, 0x6a, 0x47, 0xe0, 0xe5, 0x0f, 0x02, 0xb3, 0x46, 0x1f, 0xbd, 0x80, 0xd0, 0xbd,
	0x04, 0x34, 0xde, 0xdf, 0xb7, 0x7e, 0x23, 0xe2, 0xe3, 0xfd, 0x1a, 0x3e, 0x50, 0x17, 0x10, 0xba,
	0x2b, 0xd1, 0xb7, 0xba, 0xbe, 0x2c, 0x83, 0xab, 0x5f, 0x02, 0x59, 0x71, 0x43, 0x7b, 0x06, 0x46,
	0x8b, 0xc7, 0x03, 0x05, 0x4d, 0x5f, 0x41, 0x80, 0xc6, 0xa4, 0x4f, 0xf7, 0x3b, 0xbc, 0xf3, 0xe3,
	0xc1, 0x92, 0xa6, 0xaf, 0x21, 0x74, 0xea, 0xf7, 0xb1, 0xae, 0x0f, 0x37, 0x1e, 0xae, 0x69, 0x3b,
	0xb7, 0x53, 0xb8, 0x0f, 0xa1, 0xd6, 0x7e, 0x68, 0xee, 0x2f, 0x21, 0xa6, 0x5f, 0xfc, 0x09, 0x00,
	0x00, 0xff, 0xff, 0x91, 0xa3, 0x07, 0x4b, 0x12, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserServiceClient interface {
	Create(ctx context.Context, in *UserServiceCreateRq, opts ...grpc.CallOption) (*UserServiceBool, error)
	Update(ctx context.Context, in *UserServiceUpdateRq, opts ...grpc.CallOption) (*UserServiceBool, error)
	Get(ctx context.Context, in *UserServiceGetRq, opts ...grpc.CallOption) (*UserServiceGetRs, error)
	Login(ctx context.Context, in *UserServiceLoginRq, opts ...grpc.CallOption) (*UserServiceLoginRs, error)
	GetAll(ctx context.Context, in *UserServiceGetAllRq, opts ...grpc.CallOption) (*UserServiceGetAllRs, error)
	Delete(ctx context.Context, in *UserServiceDeleteRq, opts ...grpc.CallOption) (*UserServiceBool, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) Create(ctx context.Context, in *UserServiceCreateRq, opts ...grpc.CallOption) (*UserServiceBool, error) {
	out := new(UserServiceBool)
	err := c.cc.Invoke(ctx, "/proto.UserService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Update(ctx context.Context, in *UserServiceUpdateRq, opts ...grpc.CallOption) (*UserServiceBool, error) {
	out := new(UserServiceBool)
	err := c.cc.Invoke(ctx, "/proto.UserService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Get(ctx context.Context, in *UserServiceGetRq, opts ...grpc.CallOption) (*UserServiceGetRs, error) {
	out := new(UserServiceGetRs)
	err := c.cc.Invoke(ctx, "/proto.UserService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Login(ctx context.Context, in *UserServiceLoginRq, opts ...grpc.CallOption) (*UserServiceLoginRs, error) {
	out := new(UserServiceLoginRs)
	err := c.cc.Invoke(ctx, "/proto.UserService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetAll(ctx context.Context, in *UserServiceGetAllRq, opts ...grpc.CallOption) (*UserServiceGetAllRs, error) {
	out := new(UserServiceGetAllRs)
	err := c.cc.Invoke(ctx, "/proto.UserService/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Delete(ctx context.Context, in *UserServiceDeleteRq, opts ...grpc.CallOption) (*UserServiceBool, error) {
	out := new(UserServiceBool)
	err := c.cc.Invoke(ctx, "/proto.UserService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	Create(context.Context, *UserServiceCreateRq) (*UserServiceBool, error)
	Update(context.Context, *UserServiceUpdateRq) (*UserServiceBool, error)
	Get(context.Context, *UserServiceGetRq) (*UserServiceGetRs, error)
	Login(context.Context, *UserServiceLoginRq) (*UserServiceLoginRs, error)
	GetAll(context.Context, *UserServiceGetAllRq) (*UserServiceGetAllRs, error)
	Delete(context.Context, *UserServiceDeleteRq) (*UserServiceBool, error)
}

// UnimplementedUserServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (*UnimplementedUserServiceServer) Create(ctx context.Context, req *UserServiceCreateRq) (*UserServiceBool, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedUserServiceServer) Update(ctx context.Context, req *UserServiceUpdateRq) (*UserServiceBool, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedUserServiceServer) Get(ctx context.Context, req *UserServiceGetRq) (*UserServiceGetRs, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedUserServiceServer) Login(ctx context.Context, req *UserServiceLoginRq) (*UserServiceLoginRs, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (*UnimplementedUserServiceServer) GetAll(ctx context.Context, req *UserServiceGetAllRq) (*UserServiceGetAllRs, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (*UnimplementedUserServiceServer) Delete(ctx context.Context, req *UserServiceDeleteRq) (*UserServiceBool, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserServiceCreateRq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Create(ctx, req.(*UserServiceCreateRq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserServiceUpdateRq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Update(ctx, req.(*UserServiceUpdateRq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserServiceGetRq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Get(ctx, req.(*UserServiceGetRq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserServiceLoginRq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Login(ctx, req.(*UserServiceLoginRq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserServiceGetAllRq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserService/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetAll(ctx, req.(*UserServiceGetAllRq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserServiceDeleteRq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Delete(ctx, req.(*UserServiceDeleteRq))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _UserService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _UserService_Update_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _UserService_Get_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _UserService_Login_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _UserService_GetAll_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _UserService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/userService.proto",
}
