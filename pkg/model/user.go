package model

import (
	"time"

	databases "dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/data/database/models"
)

const TableNameUser = "user"

type User struct {
	UserId           uint64              `gorm:"column:userId;type:bigint unsigned;primary_key;comment:'用户ID'"`
	Mobile           string              `gorm:"column:mobile;type:char(11);not null;default:'';index:mobile;comment:'手机号'"`
	Name             string              `gorm:"column:name;type:varchar(40);not null;default:'';comment:'用户姓名'"`
	HeadImage        string              `gorm:"column:headImage;type:varchar(255);not null;default:'';comment:'用户头像'"`
	Gender           uint8               `gorm:"column:gender;type:tinyint(1) unsigned;not null;default:'0';comment:'男1女2 默认0'"`
	Birthday         time.Time           `gorm:"column:birthday;comment:'生日'"`
	Point            uint64              `gorm:"column:point;type:bigint unsigned;not null;default:0;comment:'积分'"`
	ShopRelated      []*UserShop         `gorm:"foreignkey:UserId"`
	UserPointHistory []*UserPointHistory `gorm:"foreignkey:UserId"`
	InviterId        uint64              `gorm:"column:inviterId;type:bigint unsigned;not null;default:0;comment:'邀请者ID'"`
	InviterType      uint8               `gorm:"column:inviterType;type:tinyint(1) unsigned;not null;default:0;comment:'邀请者类型'"`
	databases.Time
}

func (user *User) TableName() string {

	return TableNameUser
}

// 返回数据的结构体
type UserInfo struct {
	UserId      uint64    `json:"userId"`
	Mobile      string    `json:"mobile"`
	Name        string    `json:"name"`
	Gender      uint8     `json:"gender"`
	Birthday    time.Time `json:"birthday"`
	OpenId      string    `json:"openId"`
	InviterId   uint64    `json:"inviterId"`
	InviterType uint8     `json:"inviterType"`
	CreatedAt   time.Time `json:"createdAt"`
	HeadImage   string    `json:"headImage"`
}

// 处理请求的结构体
type ModifyUserPoint struct {
	// 用户 Id
	UserId uint64 `json:"userId"`
	// 备注
	Remark string `json:"remark"`
	// 关联的订单 Id
	OrderId uint64 `json:"orderId"`
	// 修改对应积分
	Point uint64 `json:"point"`
	// 操作类型
	Type uint32 `json:"type"`
}
