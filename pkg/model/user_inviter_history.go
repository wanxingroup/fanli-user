package model

import (
	databases "dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/data/database/models"
)

const TableNameUserInviterHistory = "user_inviter_history"

type UserInviterHistory struct {
	Id          uint64 `gorm:"column:id;type:bigint unsigned;primary_key;comment:'Id'"`
	UserId      uint64 `gorm:"column:userId;type:bigint unsigned;index:userId_idx;comment:'用户ID'"`
	InviterId   uint64 `gorm:"column:inviterId;type:bigint unsigned;not null;default:0;comment:'邀请者ID'"`
	InviterType uint8  `gorm:"column:inviterType;type:tinyint(1) unsigned;not null;default:0;comment:'邀请者类型'"`
	databases.Time
}

func (*UserInviterHistory) TableName() string {

	return TableNameUserInviterHistory
}

const (
	InviterTypeDefault    = 1 //默认平台
	InviterTypeFranchisee = 2 //加盟商
	InviterTypeShop       = 3 //店铺
	InviterTypeStaff      = 4 //员工
	InviterTypeUser       = 5 //用户
)
