package application

import (
	"dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/data/database"

	"dev-gitlab.wanxingrowth.com/fanli/user/pkg/constant"
	"dev-gitlab.wanxingrowth.com/fanli/user/pkg/model"
)

func autoMigration() {

	db := database.GetDB(constant.DatabaseConfigKey)
	db.AutoMigrate(model.User{})
	db.AutoMigrate(model.UserShop{})
	db.AutoMigrate(model.UserAddress{})
	db.AutoMigrate(model.UserPointHistory{})
}
