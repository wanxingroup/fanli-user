package user

import (
	"context"
	"testing"
	"time"

	"dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/data/database"
	"github.com/stretchr/testify/assert"

	"dev-gitlab.wanxingrowth.com/fanli/user/pkg/constant"
	"dev-gitlab.wanxingrowth.com/fanli/user/pkg/model"
	"dev-gitlab.wanxingrowth.com/fanli/user/pkg/rpc/protos"
)

func TestController_BindMobile(t *testing.T) {
	tests := []struct {
		input    *protos.BindMobileRequest
		wantData *protos.BindMobileReply
		initData func()
	}{
		{
			initData: func() {

				userData := &model.User{
					UserId:   2001,
					Mobile:   "13167587909",
					Name:     "张三",
					Gender:   1,
					Birthday: time.Unix(0, 0),
					Point:    0,
				}
				err := database.GetDB(constant.DatabaseConfigKey).Create(userData).Error
				userData = &model.User{
					UserId:   9987,
					Mobile:   "",
					Name:     "",
					Gender:   0,
					Birthday: time.Unix(0, 0),
					Point:    0,
				}
				err = database.GetDB(constant.DatabaseConfigKey).Create(userData).Error
				assert.Nil(t, err)

				userShop := &model.UserShop{
					OpenId:   "111222333",
					UserId:   2001,
					ShopId:   1009,
					OpenType: 0,
				}
				err = database.GetDB(constant.DatabaseConfigKey).Create(userShop).Error
				userShop = &model.UserShop{
					OpenId:   "333222111",
					UserId:   9987,
					ShopId:   1010,
					OpenType: 0,
				}
				err = database.GetDB(constant.DatabaseConfigKey).Create(userShop).Error

				assert.Nil(t, err)
			},
			input: &protos.BindMobileRequest{
				ShopId: 1010,
				Mobile: "13167587909",
				UserId: 9987,
			},
		},
	}

	for _, test := range tests {

		if test.initData != nil {
			test.initData()
		}

		c := Controller{}
		reply, err := c.BindMobile(context.Background(), test.input)
		assert.Nil(t, err)

		if test.wantData != nil {
			assert.NotNil(t, reply)
		}

	}
}
