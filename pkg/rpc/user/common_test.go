package user

import (
	"testing"
	"time"

	"dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/data/database"
	"github.com/stretchr/testify/assert"

	"dev-gitlab.wanxingrowth.com/fanli/user/pkg/constant"
	"dev-gitlab.wanxingrowth.com/fanli/user/pkg/model"
)

func TestGetUserAndUpdate(t *testing.T) {

	tests := []struct {
		initData func()
		userId   uint64
	}{
		{
			initData: func() {
				user := model.User{
					UserId:   10001,
					Mobile:   "",
					Name:     "",
					Gender:   0,
					Birthday: time.Unix(0, 0),
				}
				assert.Nil(t, database.GetDB(constant.DatabaseConfigKey).Save(&user).Error)
			},
			userId: 10001,
		},
		{
			initData: func() {
				userShop := model.UserShop{
					OpenId: "kii2i22",
					UserId: 10001,
					ShopId: 22222,
				}
				assert.Nil(t, database.GetDB(constant.DatabaseConfigKey).Save(&userShop).Error)
			},
			userId: 10001,
		},
	}

	for _, test := range tests {

		if test.initData != nil {
			test.initData()
		}

		user, err := getUser(test.userId)
		assert.Nil(t, err, test)
		if err != nil {
			continue
		}

		err = database.GetDB(constant.DatabaseConfigKey).Model(model.User{}).
			Where(&model.User{UserId: user.UserId}).
			Update(&model.User{
				Mobile: "18800222222",
			}).Error
		assert.Nil(t, err, test)
	}
}

func TestGetUserAndSave(t *testing.T) {

	tests := []struct {
		initData func()
		userId   uint64
	}{
		{
			initData: func() {
				user := model.User{
					UserId:   10001,
					Mobile:   "",
					Name:     "",
					Gender:   0,
					Birthday: time.Unix(0, 0),
				}
				assert.Nil(t, database.GetDB(constant.DatabaseConfigKey).Save(&user).Error)
			},
			userId: 10001,
		},
		{
			initData: func() {
				userShop := model.UserShop{
					OpenId: "kii2i22",
					UserId: 10001,
					ShopId: 22222,
				}
				assert.Nil(t, database.GetDB(constant.DatabaseConfigKey).Save(&userShop).Error)
			},
			userId: 10001,
		},
	}

	for _, test := range tests {

		if test.initData != nil {
			test.initData()
		}

		user, err := getUser(test.userId)
		assert.Nil(t, err, test)
		if err != nil {
			continue
		}

		user.Mobile = "13800138000"

		err = database.GetDB(constant.DatabaseConfigKey).Save(user).Error
		assert.Nil(t, err, test)
	}
}
