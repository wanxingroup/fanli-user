package model

import (
	databases "dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/data/database/models"
)

const TableNameUserShop = "user_shop"

type UserShop struct {
	OpenId   string `gorm:"column:openId;type:varchar(40);primary_key;comment:'平台的openId'"`
	UserId   uint64 `gorm:"column:userId;type:bigint unsigned;index:userId;comment:'用户ID'"`
	ShopId   uint64 `gorm:"column:shopId;type:bigint unsigned;index:userId;comment:'商户Id'"`
	OpenType uint8  `gorm:"column:openType;type:tinyint unsigned;not null;default:'0';comment:'0 小程序 1 公众号 2 支付宝等'"`
	databases.Time
}

func (userShop *UserShop) TableName() string {

	return TableNameUserShop
}
