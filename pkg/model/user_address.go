package model

import (
	databases "dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/data/database/models"
)

const TableNameUserAddress = "user_address"

type UserAddress struct {
	Id           uint64 `gorm:"column:id;type:bigint unsigned;primary_key;comment:'主键ID'"`
	UserId       uint64 `gorm:"column:userId;index:userId;comment:'用户ID'"`
	ReceiverName string `gorm:"column:receiverName;type:varchar(32);default:'';comment:'名称'"`
	Gender       uint8  `gorm:"column:gender;type:tinyint(1) unsigned;not null;default:'0';comment:'男1女2 默认0'"`
	Country      string `gorm:"column:country;type:varchar(32);default:'';comment:'国家'"`
	Province     string `gorm:"column:province;type:varchar(32);default:'';comment:'省份'"`
	City         string `gorm:"column:city;type:varchar(32);default:'';comment:'城市'"`
	District     string `gorm:"column:district;type:varchar(32);default:'';comment:'区县'"`
	Address      string `gorm:"column:address;type:varchar(128);default:'';comment:'详细地址'"`
	Tag          uint8  `gorm:"column:tag;default:'0';comment:'标签0 默认(无标签) 1家，2公司，3学校'"`
	Mobile       string `gorm:"column:mobile;type:char(11);not null;default:'';index:mobile;comment:'手机号'"`
	Type         uint8  `gorm:"column:type;default:1;comment:'是否默认地址1：普通，2默认送货'"`
	databases.Time
}

func (userShop *UserAddress) TableName() string {

	return TableNameUserAddress
}

const (
	UserAddressCommon  = 1 // 普通地址
	UserAddressDefault = 2 // 默认送货地址
)
