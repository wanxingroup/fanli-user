package model

import (
	databases "dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/data/database/models"
)

const TableNameUserPointHistory = "user_point_history"

type UserPointHistory struct {
	Id      uint64 `gorm:"column:id;type:bigint unsigned;primary_key;comment:'Id'"`
	UserId  uint64 `gorm:"column:userId;type:bigint unsigned;comment:'用户ID'"`
	Type    uint8  `gorm:"column:type;default:0;comment:'类型: 0 增长 1 消费' "`
	Point   uint64 `gorm:"column:point;type:bigint unsigned;not null;default:0;comment:'积分'"`
	Remark  string `gorm:"column:remark;type:varchar(255);null;default:'';comment:'备注'"`
	OrderId uint64 `gorm:"column:orderId;type:bigint unsigned;null;comment:'关联的 订单 Id,如果不关联为空'"`
	databases.Time
}

func (user *UserPointHistory) TableName() string {

	return TableNameUserPointHistory
}

const (
	UserIncreasePoint = 0 // 增长积分
	UserDecreasePoint = 1 // 减少积分
)
