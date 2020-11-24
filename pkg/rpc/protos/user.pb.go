// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

/*
Package protos is a generated protocol buffer package.

It is generated from these files:
	user.proto
	user_address.proto
	user_common.proto

It has these top-level messages:
	GetUserListRequest
	UserInfo
	GetUserListReply
	BindMobileRequest
	BindMobileReply
	GetUserInfoRequest
	GetUserInfoReply
	ModifyUserInfoRequest
	ModifyUserInfoReply
	SetInviterRequest
	SetInviterReply
	GetUserInfoByOpenIdRequest
	GetUserInfoByOpenIdReply
	LoginRequest
	LoginReply
	ModifyUserPointRequest
	ModifyUserPointReply
	Address
	GetAddressListRequest
	GetAddressListReply
	GetAddressInfoRequest
	GetAddressInfoReply
	SetAddressRequest
	SetAddressReply
	CreateAddressRequest
	CreateAddressReply
	DeleteAddressRequest
	DeleteAddressReply
	Error
*/
package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GetUserListRequest struct {
	// 页码
	Page uint64 `protobuf:"varint,1,opt,name=page" json:"page,omitempty"`
	// 分页大小
	PageSize uint64 `protobuf:"varint,2,opt,name=PageSize" json:"PageSize,omitempty"`
	// 搜索条件（手机号）
	Mobile string `protobuf:"bytes,3,opt,name=mobile" json:"mobile,omitempty"`
	// 搜索条件（性别）
	Gender uint64 `protobuf:"varint,4,opt,name=gender" json:"gender,omitempty"`
	// 搜索条件（用户注册开始时间）
	RegStartTime string `protobuf:"bytes,5,opt,name=regStartTime" json:"regStartTime,omitempty"`
	// 搜索条件（用户注册结束时间）
	RegEndTime string `protobuf:"bytes,6,opt,name=regEndTime" json:"regEndTime,omitempty"`
	// 搜索条件（店铺 ID）
	ShopId uint64 `protobuf:"varint,7,opt,name=shopId" json:"shopId,omitempty"`
	// 搜索条件 用户 IDs
	UserIds []uint64 `protobuf:"varint,8,rep,packed,name=userIds" json:"userIds,omitempty"`
	// 搜索条件 用户昵称
	UserName string `protobuf:"bytes,9,opt,name=userName" json:"userName,omitempty"`
}

func (m *GetUserListRequest) Reset()                    { *m = GetUserListRequest{} }
func (m *GetUserListRequest) String() string            { return proto.CompactTextString(m) }
func (*GetUserListRequest) ProtoMessage()               {}
func (*GetUserListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GetUserListRequest) GetPage() uint64 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *GetUserListRequest) GetPageSize() uint64 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *GetUserListRequest) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

func (m *GetUserListRequest) GetGender() uint64 {
	if m != nil {
		return m.Gender
	}
	return 0
}

func (m *GetUserListRequest) GetRegStartTime() string {
	if m != nil {
		return m.RegStartTime
	}
	return ""
}

func (m *GetUserListRequest) GetRegEndTime() string {
	if m != nil {
		return m.RegEndTime
	}
	return ""
}

func (m *GetUserListRequest) GetShopId() uint64 {
	if m != nil {
		return m.ShopId
	}
	return 0
}

func (m *GetUserListRequest) GetUserIds() []uint64 {
	if m != nil {
		return m.UserIds
	}
	return nil
}

func (m *GetUserListRequest) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

// 用户信息
type UserInfo struct {
	// 用户 ID
	UserId uint64 `protobuf:"varint,1,opt,name=userId" json:"userId,omitempty"`
	// 手机号
	Mobile string `protobuf:"bytes,2,opt,name=mobile" json:"mobile,omitempty"`
	// 用户名称
	Name string `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	// 用户性别
	Gender uint64 `protobuf:"varint,4,opt,name=gender" json:"gender,omitempty"`
	// 用户生日
	Birthday string `protobuf:"bytes,5,opt,name=birthday" json:"birthday,omitempty"`
	// 微信 OpenID
	OpenId string `protobuf:"bytes,6,opt,name=openId" json:"openId,omitempty"`
	// 创建时间（注册时间）
	CreatedAt string `protobuf:"bytes,7,opt,name=createdAt" json:"createdAt,omitempty"`
	// 积分
	Point uint64 `protobuf:"varint,8,opt,name=point" json:"point,omitempty"`
	// 邀请者类型
	InviterType uint32 `protobuf:"varint,9,opt,name=inviterType" json:"inviterType,omitempty"`
	// 邀请者ID
	InviterId uint64 `protobuf:"varint,10,opt,name=inviterId" json:"inviterId,omitempty"`
	// 用户头像
	HeadImage string `protobuf:"bytes,11,opt,name=headImage" json:"headImage,omitempty"`
}

func (m *UserInfo) Reset()                    { *m = UserInfo{} }
func (m *UserInfo) String() string            { return proto.CompactTextString(m) }
func (*UserInfo) ProtoMessage()               {}
func (*UserInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *UserInfo) GetUserId() uint64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *UserInfo) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

func (m *UserInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserInfo) GetGender() uint64 {
	if m != nil {
		return m.Gender
	}
	return 0
}

func (m *UserInfo) GetBirthday() string {
	if m != nil {
		return m.Birthday
	}
	return ""
}

func (m *UserInfo) GetOpenId() string {
	if m != nil {
		return m.OpenId
	}
	return ""
}

func (m *UserInfo) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *UserInfo) GetPoint() uint64 {
	if m != nil {
		return m.Point
	}
	return 0
}

func (m *UserInfo) GetInviterType() uint32 {
	if m != nil {
		return m.InviterType
	}
	return 0
}

func (m *UserInfo) GetInviterId() uint64 {
	if m != nil {
		return m.InviterId
	}
	return 0
}

func (m *UserInfo) GetHeadImage() string {
	if m != nil {
		return m.HeadImage
	}
	return ""
}

type GetUserListReply struct {
	// 错误信息
	Err *Error `protobuf:"bytes,1,opt,name=err" json:"err,omitempty"`
	// 用户列表
	UserInfoList []*UserInfo `protobuf:"bytes,2,rep,name=UserInfoList" json:"UserInfoList,omitempty"`
	// 总数量
	Count uint64 `protobuf:"varint,3,opt,name=count" json:"count,omitempty"`
}

func (m *GetUserListReply) Reset()                    { *m = GetUserListReply{} }
func (m *GetUserListReply) String() string            { return proto.CompactTextString(m) }
func (*GetUserListReply) ProtoMessage()               {}
func (*GetUserListReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *GetUserListReply) GetErr() *Error {
	if m != nil {
		return m.Err
	}
	return nil
}

func (m *GetUserListReply) GetUserInfoList() []*UserInfo {
	if m != nil {
		return m.UserInfoList
	}
	return nil
}

func (m *GetUserListReply) GetCount() uint64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type BindMobileRequest struct {
	// 店铺 ID
	ShopId uint64 `protobuf:"varint,1,opt,name=shopId" json:"shopId,omitempty"`
	// 手机号
	Mobile string `protobuf:"bytes,2,opt,name=mobile" json:"mobile,omitempty"`
	// 用户 ID
	UserId uint64 `protobuf:"varint,3,opt,name=userId" json:"userId,omitempty"`
}

func (m *BindMobileRequest) Reset()                    { *m = BindMobileRequest{} }
func (m *BindMobileRequest) String() string            { return proto.CompactTextString(m) }
func (*BindMobileRequest) ProtoMessage()               {}
func (*BindMobileRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *BindMobileRequest) GetShopId() uint64 {
	if m != nil {
		return m.ShopId
	}
	return 0
}

func (m *BindMobileRequest) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

func (m *BindMobileRequest) GetUserId() uint64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type BindMobileReply struct {
	// 用户信息
	UserInfo *UserInfo `protobuf:"bytes,1,opt,name=UserInfo" json:"UserInfo,omitempty"`
	// 错误信息
	Err *Error `protobuf:"bytes,2,opt,name=err" json:"err,omitempty"`
}

func (m *BindMobileReply) Reset()                    { *m = BindMobileReply{} }
func (m *BindMobileReply) String() string            { return proto.CompactTextString(m) }
func (*BindMobileReply) ProtoMessage()               {}
func (*BindMobileReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *BindMobileReply) GetUserInfo() *UserInfo {
	if m != nil {
		return m.UserInfo
	}
	return nil
}

func (m *BindMobileReply) GetErr() *Error {
	if m != nil {
		return m.Err
	}
	return nil
}

type GetUserInfoRequest struct {
	// 用户 ID
	UserId uint64 `protobuf:"varint,1,opt,name=userId" json:"userId,omitempty"`
	// 店铺 ID
	ShopId uint64 `protobuf:"varint,2,opt,name=shopId" json:"shopId,omitempty"`
}

func (m *GetUserInfoRequest) Reset()                    { *m = GetUserInfoRequest{} }
func (m *GetUserInfoRequest) String() string            { return proto.CompactTextString(m) }
func (*GetUserInfoRequest) ProtoMessage()               {}
func (*GetUserInfoRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *GetUserInfoRequest) GetUserId() uint64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *GetUserInfoRequest) GetShopId() uint64 {
	if m != nil {
		return m.ShopId
	}
	return 0
}

type GetUserInfoReply struct {
	// 用户信息
	UserInfo *UserInfo `protobuf:"bytes,1,opt,name=UserInfo" json:"UserInfo,omitempty"`
	// 错误信息
	Err *Error `protobuf:"bytes,2,opt,name=err" json:"err,omitempty"`
}

func (m *GetUserInfoReply) Reset()                    { *m = GetUserInfoReply{} }
func (m *GetUserInfoReply) String() string            { return proto.CompactTextString(m) }
func (*GetUserInfoReply) ProtoMessage()               {}
func (*GetUserInfoReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *GetUserInfoReply) GetUserInfo() *UserInfo {
	if m != nil {
		return m.UserInfo
	}
	return nil
}

func (m *GetUserInfoReply) GetErr() *Error {
	if m != nil {
		return m.Err
	}
	return nil
}

type ModifyUserInfoRequest struct {
	// 用户信息
	UserInfo *UserInfo `protobuf:"bytes,1,opt,name=UserInfo" json:"UserInfo,omitempty"`
}

func (m *ModifyUserInfoRequest) Reset()                    { *m = ModifyUserInfoRequest{} }
func (m *ModifyUserInfoRequest) String() string            { return proto.CompactTextString(m) }
func (*ModifyUserInfoRequest) ProtoMessage()               {}
func (*ModifyUserInfoRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ModifyUserInfoRequest) GetUserInfo() *UserInfo {
	if m != nil {
		return m.UserInfo
	}
	return nil
}

type ModifyUserInfoReply struct {
	// 错误信息
	Err *Error `protobuf:"bytes,1,opt,name=err" json:"err,omitempty"`
}

func (m *ModifyUserInfoReply) Reset()                    { *m = ModifyUserInfoReply{} }
func (m *ModifyUserInfoReply) String() string            { return proto.CompactTextString(m) }
func (*ModifyUserInfoReply) ProtoMessage()               {}
func (*ModifyUserInfoReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *ModifyUserInfoReply) GetErr() *Error {
	if m != nil {
		return m.Err
	}
	return nil
}

type SetInviterRequest struct {
	UserId uint64 `protobuf:"varint,1,opt,name=userId" json:"userId,omitempty"`
}

func (m *SetInviterRequest) Reset()                    { *m = SetInviterRequest{} }
func (m *SetInviterRequest) String() string            { return proto.CompactTextString(m) }
func (*SetInviterRequest) ProtoMessage()               {}
func (*SetInviterRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *SetInviterRequest) GetUserId() uint64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type SetInviterReply struct {
	Err         *Error `protobuf:"bytes,1,opt,name=err" json:"err,omitempty"`
	InviterType uint32 `protobuf:"varint,2,opt,name=inviterType" json:"inviterType,omitempty"`
	InviterId   uint64 `protobuf:"varint,3,opt,name=inviterId" json:"inviterId,omitempty"`
}

func (m *SetInviterReply) Reset()                    { *m = SetInviterReply{} }
func (m *SetInviterReply) String() string            { return proto.CompactTextString(m) }
func (*SetInviterReply) ProtoMessage()               {}
func (*SetInviterReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *SetInviterReply) GetErr() *Error {
	if m != nil {
		return m.Err
	}
	return nil
}

func (m *SetInviterReply) GetInviterType() uint32 {
	if m != nil {
		return m.InviterType
	}
	return 0
}

func (m *SetInviterReply) GetInviterId() uint64 {
	if m != nil {
		return m.InviterId
	}
	return 0
}

type GetUserInfoByOpenIdRequest struct {
	// 微信 OpenID
	OpenId string `protobuf:"bytes,1,opt,name=openId" json:"openId,omitempty"`
	// 店铺 ID
	ShopId uint64 `protobuf:"varint,2,opt,name=shopId" json:"shopId,omitempty"`
}

func (m *GetUserInfoByOpenIdRequest) Reset()                    { *m = GetUserInfoByOpenIdRequest{} }
func (m *GetUserInfoByOpenIdRequest) String() string            { return proto.CompactTextString(m) }
func (*GetUserInfoByOpenIdRequest) ProtoMessage()               {}
func (*GetUserInfoByOpenIdRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *GetUserInfoByOpenIdRequest) GetOpenId() string {
	if m != nil {
		return m.OpenId
	}
	return ""
}

func (m *GetUserInfoByOpenIdRequest) GetShopId() uint64 {
	if m != nil {
		return m.ShopId
	}
	return 0
}

type GetUserInfoByOpenIdReply struct {
	// 用户信息
	UserInfo *UserInfo `protobuf:"bytes,1,opt,name=UserInfo" json:"UserInfo,omitempty"`
	// 错误信息
	Err *Error `protobuf:"bytes,2,opt,name=err" json:"err,omitempty"`
}

func (m *GetUserInfoByOpenIdReply) Reset()                    { *m = GetUserInfoByOpenIdReply{} }
func (m *GetUserInfoByOpenIdReply) String() string            { return proto.CompactTextString(m) }
func (*GetUserInfoByOpenIdReply) ProtoMessage()               {}
func (*GetUserInfoByOpenIdReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *GetUserInfoByOpenIdReply) GetUserInfo() *UserInfo {
	if m != nil {
		return m.UserInfo
	}
	return nil
}

func (m *GetUserInfoByOpenIdReply) GetErr() *Error {
	if m != nil {
		return m.Err
	}
	return nil
}

type LoginRequest struct {
	// 微信 OpenID
	OpenId string `protobuf:"bytes,1,opt,name=openId" json:"openId,omitempty"`
	// 店铺 ID
	ShopId uint64 `protobuf:"varint,2,opt,name=shopId" json:"shopId,omitempty"`
	// 邀请码
	InviterCode string `protobuf:"bytes,3,opt,name=inviterCode" json:"inviterCode,omitempty"`
}

func (m *LoginRequest) Reset()                    { *m = LoginRequest{} }
func (m *LoginRequest) String() string            { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()               {}
func (*LoginRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *LoginRequest) GetOpenId() string {
	if m != nil {
		return m.OpenId
	}
	return ""
}

func (m *LoginRequest) GetShopId() uint64 {
	if m != nil {
		return m.ShopId
	}
	return 0
}

func (m *LoginRequest) GetInviterCode() string {
	if m != nil {
		return m.InviterCode
	}
	return ""
}

type LoginReply struct {
	// 用户信息
	UserInfo *UserInfo `protobuf:"bytes,1,opt,name=UserInfo" json:"UserInfo,omitempty"`
	// 错误信息
	Err *Error `protobuf:"bytes,2,opt,name=err" json:"err,omitempty"`
}

func (m *LoginReply) Reset()                    { *m = LoginReply{} }
func (m *LoginReply) String() string            { return proto.CompactTextString(m) }
func (*LoginReply) ProtoMessage()               {}
func (*LoginReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *LoginReply) GetUserInfo() *UserInfo {
	if m != nil {
		return m.UserInfo
	}
	return nil
}

func (m *LoginReply) GetErr() *Error {
	if m != nil {
		return m.Err
	}
	return nil
}

type ModifyUserPointRequest struct {
	// 用户 ID
	UserId uint64 `protobuf:"varint,1,opt,name=userId" json:"userId,omitempty"`
	// 备注
	Remark string `protobuf:"bytes,2,opt,name=remark" json:"remark,omitempty"`
	// 关联的订单 Id
	OrderId uint64 `protobuf:"varint,3,opt,name=orderId" json:"orderId,omitempty"`
	// 修改对应积分
	Point uint64 `protobuf:"varint,4,opt,name=point" json:"point,omitempty"`
	// 操作类型
	Type uint32 `protobuf:"varint,5,opt,name=type" json:"type,omitempty"`
}

func (m *ModifyUserPointRequest) Reset()                    { *m = ModifyUserPointRequest{} }
func (m *ModifyUserPointRequest) String() string            { return proto.CompactTextString(m) }
func (*ModifyUserPointRequest) ProtoMessage()               {}
func (*ModifyUserPointRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *ModifyUserPointRequest) GetUserId() uint64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *ModifyUserPointRequest) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *ModifyUserPointRequest) GetOrderId() uint64 {
	if m != nil {
		return m.OrderId
	}
	return 0
}

func (m *ModifyUserPointRequest) GetPoint() uint64 {
	if m != nil {
		return m.Point
	}
	return 0
}

func (m *ModifyUserPointRequest) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

type ModifyUserPointReply struct {
	// 用户修改后的积分
	Point uint64 `protobuf:"varint,1,opt,name=Point" json:"Point,omitempty"`
	// 错误信息
	Err *Error `protobuf:"bytes,2,opt,name=err" json:"err,omitempty"`
}

func (m *ModifyUserPointReply) Reset()                    { *m = ModifyUserPointReply{} }
func (m *ModifyUserPointReply) String() string            { return proto.CompactTextString(m) }
func (*ModifyUserPointReply) ProtoMessage()               {}
func (*ModifyUserPointReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{16} }

func (m *ModifyUserPointReply) GetPoint() uint64 {
	if m != nil {
		return m.Point
	}
	return 0
}

func (m *ModifyUserPointReply) GetErr() *Error {
	if m != nil {
		return m.Err
	}
	return nil
}

func init() {
	proto.RegisterType((*GetUserListRequest)(nil), "user.pkg.rpc.protos.GetUserListRequest")
	proto.RegisterType((*UserInfo)(nil), "user.pkg.rpc.protos.UserInfo")
	proto.RegisterType((*GetUserListReply)(nil), "user.pkg.rpc.protos.GetUserListReply")
	proto.RegisterType((*BindMobileRequest)(nil), "user.pkg.rpc.protos.BindMobileRequest")
	proto.RegisterType((*BindMobileReply)(nil), "user.pkg.rpc.protos.BindMobileReply")
	proto.RegisterType((*GetUserInfoRequest)(nil), "user.pkg.rpc.protos.GetUserInfoRequest")
	proto.RegisterType((*GetUserInfoReply)(nil), "user.pkg.rpc.protos.GetUserInfoReply")
	proto.RegisterType((*ModifyUserInfoRequest)(nil), "user.pkg.rpc.protos.ModifyUserInfoRequest")
	proto.RegisterType((*ModifyUserInfoReply)(nil), "user.pkg.rpc.protos.ModifyUserInfoReply")
	proto.RegisterType((*SetInviterRequest)(nil), "user.pkg.rpc.protos.SetInviterRequest")
	proto.RegisterType((*SetInviterReply)(nil), "user.pkg.rpc.protos.SetInviterReply")
	proto.RegisterType((*GetUserInfoByOpenIdRequest)(nil), "user.pkg.rpc.protos.GetUserInfoByOpenIdRequest")
	proto.RegisterType((*GetUserInfoByOpenIdReply)(nil), "user.pkg.rpc.protos.GetUserInfoByOpenIdReply")
	proto.RegisterType((*LoginRequest)(nil), "user.pkg.rpc.protos.LoginRequest")
	proto.RegisterType((*LoginReply)(nil), "user.pkg.rpc.protos.LoginReply")
	proto.RegisterType((*ModifyUserPointRequest)(nil), "user.pkg.rpc.protos.ModifyUserPointRequest")
	proto.RegisterType((*ModifyUserPointReply)(nil), "user.pkg.rpc.protos.ModifyUserPointReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for UserController service

type UserControllerClient interface {
	// 用户列表
	GetUserList(ctx context.Context, in *GetUserListRequest, opts ...grpc.CallOption) (*GetUserListReply, error)
	// 绑定用户手机号
	BindMobile(ctx context.Context, in *BindMobileRequest, opts ...grpc.CallOption) (*BindMobileReply, error)
	// 获取用户会员信息
	GetUserInfo(ctx context.Context, in *GetUserInfoRequest, opts ...grpc.CallOption) (*GetUserInfoReply, error)
	// 设置用户会员信息
	ModifyUserInfo(ctx context.Context, in *ModifyUserInfoRequest, opts ...grpc.CallOption) (*ModifyUserInfoReply, error)
	// 根据用户 OpenId 获取用户信息
	GetUserInfoByOpenId(ctx context.Context, in *GetUserInfoByOpenIdRequest, opts ...grpc.CallOption) (*GetUserInfoByOpenIdReply, error)
	// 登录用户操作（跟根据 OpenId 获取用户信息的不同在于，如果用户不存在，就会生成新的用户 ID）
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginReply, error)
	// 加/扣积分
	ModifyUserPoint(ctx context.Context, in *ModifyUserPointRequest, opts ...grpc.CallOption) (*ModifyUserPointReply, error)
	// 设置邀请者（上线）
	SetInviter(ctx context.Context, in *SetInviterRequest, opts ...grpc.CallOption) (*SetInviterReply, error)
}

type userControllerClient struct {
	cc *grpc.ClientConn
}

func NewUserControllerClient(cc *grpc.ClientConn) UserControllerClient {
	return &userControllerClient{cc}
}

func (c *userControllerClient) GetUserList(ctx context.Context, in *GetUserListRequest, opts ...grpc.CallOption) (*GetUserListReply, error) {
	out := new(GetUserListReply)
	err := grpc.Invoke(ctx, "/user.pkg.rpc.protos.UserController/GetUserList", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userControllerClient) BindMobile(ctx context.Context, in *BindMobileRequest, opts ...grpc.CallOption) (*BindMobileReply, error) {
	out := new(BindMobileReply)
	err := grpc.Invoke(ctx, "/user.pkg.rpc.protos.UserController/BindMobile", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userControllerClient) GetUserInfo(ctx context.Context, in *GetUserInfoRequest, opts ...grpc.CallOption) (*GetUserInfoReply, error) {
	out := new(GetUserInfoReply)
	err := grpc.Invoke(ctx, "/user.pkg.rpc.protos.UserController/GetUserInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userControllerClient) ModifyUserInfo(ctx context.Context, in *ModifyUserInfoRequest, opts ...grpc.CallOption) (*ModifyUserInfoReply, error) {
	out := new(ModifyUserInfoReply)
	err := grpc.Invoke(ctx, "/user.pkg.rpc.protos.UserController/ModifyUserInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userControllerClient) GetUserInfoByOpenId(ctx context.Context, in *GetUserInfoByOpenIdRequest, opts ...grpc.CallOption) (*GetUserInfoByOpenIdReply, error) {
	out := new(GetUserInfoByOpenIdReply)
	err := grpc.Invoke(ctx, "/user.pkg.rpc.protos.UserController/GetUserInfoByOpenId", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userControllerClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginReply, error) {
	out := new(LoginReply)
	err := grpc.Invoke(ctx, "/user.pkg.rpc.protos.UserController/Login", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userControllerClient) ModifyUserPoint(ctx context.Context, in *ModifyUserPointRequest, opts ...grpc.CallOption) (*ModifyUserPointReply, error) {
	out := new(ModifyUserPointReply)
	err := grpc.Invoke(ctx, "/user.pkg.rpc.protos.UserController/ModifyUserPoint", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userControllerClient) SetInviter(ctx context.Context, in *SetInviterRequest, opts ...grpc.CallOption) (*SetInviterReply, error) {
	out := new(SetInviterReply)
	err := grpc.Invoke(ctx, "/user.pkg.rpc.protos.UserController/SetInviter", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserController service

type UserControllerServer interface {
	// 用户列表
	GetUserList(context.Context, *GetUserListRequest) (*GetUserListReply, error)
	// 绑定用户手机号
	BindMobile(context.Context, *BindMobileRequest) (*BindMobileReply, error)
	// 获取用户会员信息
	GetUserInfo(context.Context, *GetUserInfoRequest) (*GetUserInfoReply, error)
	// 设置用户会员信息
	ModifyUserInfo(context.Context, *ModifyUserInfoRequest) (*ModifyUserInfoReply, error)
	// 根据用户 OpenId 获取用户信息
	GetUserInfoByOpenId(context.Context, *GetUserInfoByOpenIdRequest) (*GetUserInfoByOpenIdReply, error)
	// 登录用户操作（跟根据 OpenId 获取用户信息的不同在于，如果用户不存在，就会生成新的用户 ID）
	Login(context.Context, *LoginRequest) (*LoginReply, error)
	// 加/扣积分
	ModifyUserPoint(context.Context, *ModifyUserPointRequest) (*ModifyUserPointReply, error)
	// 设置邀请者（上线）
	SetInviter(context.Context, *SetInviterRequest) (*SetInviterReply, error)
}

func RegisterUserControllerServer(s *grpc.Server, srv UserControllerServer) {
	s.RegisterService(&_UserController_serviceDesc, srv)
}

func _UserController_GetUserList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserControllerServer).GetUserList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.pkg.rpc.protos.UserController/GetUserList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserControllerServer).GetUserList(ctx, req.(*GetUserListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserController_BindMobile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BindMobileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserControllerServer).BindMobile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.pkg.rpc.protos.UserController/BindMobile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserControllerServer).BindMobile(ctx, req.(*BindMobileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserController_GetUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserControllerServer).GetUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.pkg.rpc.protos.UserController/GetUserInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserControllerServer).GetUserInfo(ctx, req.(*GetUserInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserController_ModifyUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModifyUserInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserControllerServer).ModifyUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.pkg.rpc.protos.UserController/ModifyUserInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserControllerServer).ModifyUserInfo(ctx, req.(*ModifyUserInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserController_GetUserInfoByOpenId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserInfoByOpenIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserControllerServer).GetUserInfoByOpenId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.pkg.rpc.protos.UserController/GetUserInfoByOpenId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserControllerServer).GetUserInfoByOpenId(ctx, req.(*GetUserInfoByOpenIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserController_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserControllerServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.pkg.rpc.protos.UserController/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserControllerServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserController_ModifyUserPoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModifyUserPointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserControllerServer).ModifyUserPoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.pkg.rpc.protos.UserController/ModifyUserPoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserControllerServer).ModifyUserPoint(ctx, req.(*ModifyUserPointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserController_SetInviter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetInviterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserControllerServer).SetInviter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.pkg.rpc.protos.UserController/SetInviter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserControllerServer).SetInviter(ctx, req.(*SetInviterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserController_serviceDesc = grpc.ServiceDesc{
	ServiceName: "user.pkg.rpc.protos.UserController",
	HandlerType: (*UserControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserList",
			Handler:    _UserController_GetUserList_Handler,
		},
		{
			MethodName: "BindMobile",
			Handler:    _UserController_BindMobile_Handler,
		},
		{
			MethodName: "GetUserInfo",
			Handler:    _UserController_GetUserInfo_Handler,
		},
		{
			MethodName: "ModifyUserInfo",
			Handler:    _UserController_ModifyUserInfo_Handler,
		},
		{
			MethodName: "GetUserInfoByOpenId",
			Handler:    _UserController_GetUserInfoByOpenId_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _UserController_Login_Handler,
		},
		{
			MethodName: "ModifyUserPoint",
			Handler:    _UserController_ModifyUserPoint_Handler,
		},
		{
			MethodName: "SetInviter",
			Handler:    _UserController_SetInviter_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}

func init() { proto.RegisterFile("user.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 836 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x56, 0x4f, 0x6f, 0xd3, 0x4a,
	0x10, 0xaf, 0x1d, 0x27, 0x4d, 0x26, 0x7d, 0xed, 0xeb, 0xb6, 0xaf, 0xb2, 0xac, 0xf7, 0x1e, 0x61,
	0xc5, 0x9f, 0x40, 0x21, 0x48, 0xe5, 0xc4, 0xb1, 0x2d, 0x15, 0x8a, 0xd4, 0x42, 0xe5, 0x96, 0x4b,
	0x41, 0x2a, 0x6e, 0xbc, 0x4d, 0xac, 0x26, 0x5e, 0xb3, 0xd9, 0x80, 0x52, 0x0e, 0x1c, 0x90, 0x38,
	0xf1, 0x05, 0xf8, 0x0e, 0xf0, 0x1d, 0xd1, 0xee, 0x3a, 0xf6, 0x26, 0x71, 0xea, 0xb4, 0x87, 0x9e,
	0xb2, 0x33, 0x9e, 0xf9, 0xed, 0xfe, 0x66, 0x7f, 0x33, 0x1b, 0x80, 0x41, 0x9f, 0xb0, 0x46, 0xc4,
	0x28, 0xa7, 0x68, 0x4d, 0xad, 0x2f, 0xda, 0x0d, 0x16, 0xb5, 0x94, 0xaf, 0xef, 0xac, 0x0a, 0xe7,
	0x69, 0x8b, 0xf6, 0x7a, 0x34, 0x54, 0x3e, 0xfc, 0xdd, 0x04, 0xf4, 0x8a, 0xf0, 0xb7, 0x7d, 0xc2,
	0xf6, 0x83, 0x3e, 0x77, 0xc9, 0xc7, 0x01, 0xe9, 0x73, 0x84, 0xc0, 0x8a, 0xbc, 0x36, 0xb1, 0x8d,
	0x9a, 0x51, 0xb7, 0x5c, 0xb9, 0x46, 0x0e, 0x94, 0x0f, 0xbd, 0x36, 0x39, 0x0a, 0x2e, 0x89, 0x6d,
	0x4a, 0x7f, 0x62, 0xa3, 0x0d, 0x28, 0xf5, 0xe8, 0x59, 0xd0, 0x25, 0x76, 0xa1, 0x66, 0xd4, 0x2b,
	0x6e, 0x6c, 0x09, 0x7f, 0x9b, 0x84, 0x3e, 0x61, 0xb6, 0x25, 0x33, 0x62, 0x0b, 0x61, 0x58, 0x62,
	0xa4, 0x7d, 0xc4, 0x3d, 0xc6, 0x8f, 0x83, 0x1e, 0xb1, 0x8b, 0x32, 0x6b, 0xcc, 0x87, 0xfe, 0x07,
	0x60, 0xa4, 0xbd, 0x17, 0xfa, 0x32, 0xa2, 0x24, 0x23, 0x34, 0x8f, 0xc0, 0xee, 0x77, 0x68, 0xd4,
	0xf4, 0xed, 0x45, 0x85, 0xad, 0x2c, 0x64, 0xc3, 0xa2, 0xe0, 0xd9, 0xf4, 0xfb, 0x76, 0xb9, 0x56,
	0xa8, 0x5b, 0xee, 0xc8, 0x14, 0x0c, 0xc4, 0xf2, 0xb5, 0xd7, 0x23, 0x76, 0x45, 0xe2, 0x25, 0x36,
	0xfe, 0x65, 0x42, 0x59, 0x54, 0xa1, 0x19, 0x9e, 0x53, 0x01, 0xad, 0x72, 0xe2, 0x02, 0xc4, 0x96,
	0x46, 0xd3, 0x1c, 0xa3, 0x89, 0xc0, 0x0a, 0x05, 0xa8, 0x22, 0x2f, 0xd7, 0x33, 0xa9, 0x3b, 0x50,
	0x3e, 0x0b, 0x18, 0xef, 0xf8, 0xde, 0x30, 0xa6, 0x9d, 0xd8, 0x22, 0x87, 0x46, 0x24, 0x6c, 0xfa,
	0x31, 0xdd, 0xd8, 0x42, 0xff, 0x42, 0xa5, 0xc5, 0x88, 0xc7, 0x89, 0xbf, 0xcd, 0x25, 0xdb, 0x8a,
	0x9b, 0x3a, 0xd0, 0x3a, 0x14, 0x23, 0x1a, 0x84, 0xdc, 0x2e, 0xcb, 0x8d, 0x94, 0x81, 0x6a, 0x50,
	0x0d, 0xc2, 0x4f, 0x01, 0x27, 0xec, 0x78, 0x18, 0x29, 0xbe, 0x7f, 0xb9, 0xba, 0x4b, 0xa0, 0xc6,
	0x66, 0xd3, 0xb7, 0x41, 0xe6, 0xa6, 0x0e, 0xf1, 0xb5, 0x43, 0x3c, 0xbf, 0xd9, 0x13, 0x3a, 0xa8,
	0xaa, 0x3d, 0x13, 0x07, 0xfe, 0x69, 0xc0, 0xdf, 0x63, 0xba, 0x89, 0xba, 0x43, 0xf4, 0x04, 0x0a,
	0x84, 0x31, 0x59, 0xb3, 0xea, 0x96, 0xd3, 0xc8, 0x90, 0x60, 0x63, 0x8f, 0x31, 0xca, 0x5c, 0x11,
	0x86, 0xb6, 0x61, 0x69, 0x54, 0x70, 0x01, 0x61, 0x9b, 0xb5, 0x42, 0xbd, 0xba, 0xf5, 0x5f, 0x66,
	0xda, 0x28, 0xd0, 0x1d, 0x4b, 0x11, 0xcc, 0x5b, 0x74, 0x10, 0x72, 0x59, 0x78, 0xcb, 0x55, 0x06,
	0x7e, 0x07, 0xab, 0x3b, 0x41, 0xe8, 0x1f, 0xc8, 0xbb, 0x19, 0x29, 0x3a, 0x55, 0x8b, 0x31, 0xa6,
	0x96, 0x59, 0x57, 0x9a, 0x4a, 0xa0, 0xa0, 0x4b, 0x00, 0x5f, 0xc2, 0x8a, 0x0e, 0x2e, 0x68, 0xbf,
	0x48, 0x95, 0x13, 0x73, 0xcf, 0x21, 0x91, 0x0a, 0x2d, 0xae, 0x98, 0x39, 0x57, 0xc5, 0xf0, 0xcb,
	0xa4, 0x57, 0x25, 0x4c, 0xca, 0x6c, 0x96, 0x58, 0x63, 0xc6, 0xa6, 0xce, 0x18, 0x7f, 0x49, 0x6e,
	0x4e, 0xa1, 0xdc, 0x2a, 0x05, 0x17, 0xfe, 0x39, 0xa0, 0x7e, 0x70, 0x3e, 0x9c, 0x64, 0x71, 0xf3,
	0x13, 0xe0, 0x5d, 0x58, 0x9b, 0xc4, 0xbc, 0xb6, 0x1a, 0xf1, 0x26, 0xac, 0x1e, 0x11, 0xde, 0x54,
	0xf2, 0xcf, 0x29, 0x2d, 0xfe, 0x0a, 0x2b, 0x7a, 0xf0, 0xf5, 0xb5, 0x3f, 0xd1, 0x9c, 0x66, 0x4e,
	0x73, 0x16, 0x26, 0x9a, 0x13, 0xef, 0x83, 0xa3, 0xdd, 0xe1, 0xce, 0xf0, 0x8d, 0x9c, 0x13, 0xda,
	0xb1, 0xe3, 0x31, 0x62, 0x8c, 0x8d, 0x91, 0x59, 0x8a, 0xf8, 0x66, 0x80, 0x9d, 0x09, 0x77, 0xab,
	0xd2, 0xf8, 0x00, 0x4b, 0xfb, 0xb4, 0x1d, 0x84, 0x37, 0x64, 0xa1, 0xd5, 0x74, 0x97, 0xfa, 0xa3,
	0x59, 0xac, 0xbb, 0xf0, 0x00, 0x20, 0xde, 0xe1, 0x56, 0x89, 0xfd, 0x30, 0x60, 0x23, 0x15, 0xe8,
	0xa1, 0x98, 0xce, 0x73, 0xf4, 0x2e, 0x23, 0x3d, 0x8f, 0x5d, 0x8c, 0xa6, 0x92, 0xb2, 0xc4, 0xdb,
	0x46, 0x99, 0xaf, 0x69, 0x62, 0x64, 0xa6, 0x8f, 0x80, 0xa5, 0x3f, 0x02, 0x08, 0x2c, 0x2e, 0x04,
	0x56, 0x94, 0x02, 0x93, 0x6b, 0x7c, 0x02, 0xeb, 0x53, 0xa7, 0x11, 0xf5, 0x58, 0x87, 0xa2, 0xb4,
	0xe2, 0xa3, 0x28, 0xe3, 0x7a, 0x54, 0xb7, 0x7e, 0x97, 0x60, 0x59, 0xc0, 0xee, 0xd2, 0x90, 0x33,
	0xda, 0xed, 0x12, 0x86, 0x4e, 0xa1, 0xaa, 0x3d, 0x14, 0xe8, 0x61, 0x26, 0xc4, 0xf4, 0x5f, 0x10,
	0xe7, 0x7e, 0x7e, 0x60, 0xd4, 0x1d, 0xe2, 0x05, 0xf4, 0x1e, 0x20, 0x9d, 0xc8, 0xe8, 0x41, 0x66,
	0xda, 0xd4, 0x7b, 0xe0, 0xdc, 0xcb, 0x8d, 0x53, 0xe8, 0xe9, 0xf1, 0xe5, 0xcd, 0x5f, 0x79, 0x7c,
	0x6d, 0x9e, 0x5d, 0x7d, 0xfc, 0x64, 0x48, 0xe1, 0x05, 0xd4, 0x81, 0xe5, 0xf1, 0xe9, 0x85, 0x1e,
	0x67, 0xa6, 0x66, 0x8e, 0x4d, 0xa7, 0x3e, 0x57, 0xac, 0xda, 0xe9, 0x33, 0xac, 0x65, 0x74, 0x39,
	0x7a, 0x96, 0x77, 0xd2, 0x89, 0xf1, 0xe2, 0x3c, 0x9d, 0x3f, 0x41, 0x6d, 0x7c, 0x00, 0x45, 0xd9,
	0x77, 0xe8, 0x6e, 0x66, 0xa6, 0xde, 0xf5, 0xce, 0x9d, 0xab, 0x42, 0x14, 0xdc, 0x05, 0xac, 0x4c,
	0x08, 0x18, 0x6d, 0xe6, 0x94, 0x41, 0x6f, 0x3a, 0xe7, 0xd1, 0x7c, 0xc1, 0x89, 0xba, 0xd2, 0x51,
	0x3f, 0x43, 0x5d, 0x53, 0x0f, 0xc7, 0x0c, 0x75, 0x4d, 0xbc, 0x19, 0x78, 0x61, 0xa7, 0x7c, 0x52,
	0x52, 0xdf, 0xce, 0xd4, 0xef, 0xf3, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xbb, 0xd4, 0xf0, 0x7b,
	0xc5, 0x0b, 0x00, 0x00,
}