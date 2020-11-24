package user

import (
	"testing"
	"time"

	"dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/data/database"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"

	"dev-gitlab.wanxingrowth.com/fanli/user/pkg/constant"
	"dev-gitlab.wanxingrowth.com/fanli/user/pkg/model"
	"dev-gitlab.wanxingrowth.com/fanli/user/pkg/rpc/errorcode"
	"dev-gitlab.wanxingrowth.com/fanli/user/pkg/rpc/protos"
)

func TestController_ModifyUserPoint(t *testing.T) {
	tests := []struct {
		initData  func()
		input     *protos.ModifyUserPointRequest
		want      *protos.ModifyUserPointReply
		wantError *protos.Error
	}{
		{
			initData: func() {
				insertData := model.User{
					UserId:   901,
					Mobile:   "13167185909",
					Name:     "张三",
					Gender:   1,
					Birthday: time.Unix(0, 0),
					Point:    0,
				}

				assert.Nil(t, database.GetDB(constant.DatabaseConfigKey).Create(&insertData).Error)
			},
			input: &protos.ModifyUserPointRequest{
				UserId:  901,
				Remark:  "加",
				OrderId: 0,
				Point:   1,
				Type:    model.UserIncreasePoint,
			},
			want: &protos.ModifyUserPointReply{
				Point: 1,
				Err:   nil,
			},
		},
		{
			initData: func() {
				insertData := &model.User{
					UserId:   902,
					Mobile:   "13167185908",
					Name:     "张三2",
					Gender:   1,
					Birthday: time.Unix(0, 0),
					Point:    10,
				}
				assert.Nil(t, database.GetDB(constant.DatabaseConfigKey).Create(&insertData).Error)

			},
			input: &protos.ModifyUserPointRequest{
				UserId:  902,
				Remark:  "减",
				OrderId: 0,
				Point:   1,
				Type:    model.UserDecreasePoint,
			},
			want: &protos.ModifyUserPointReply{
				Point: 9,
				Err:   nil,
			},
		},
		{
			initData: func() {
				insertData := &model.User{
					UserId:   903,
					Mobile:   "13167185911",
					Name:     "张三2",
					Gender:   1,
					Birthday: time.Unix(0, 0),
					Point:    10,
				}
				assert.Nil(t, database.GetDB(constant.DatabaseConfigKey).Create(&insertData).Error)

			},
			input: &protos.ModifyUserPointRequest{
				UserId:  903,
				Remark:  "减",
				OrderId: 0,
				Point:   11,
				Type:    model.UserDecreasePoint,
			},
			wantError: &protos.Error{
				Code:    errorcode.UserPointIsNotEnough,
				Message: errorcode.UserPointIsNotEnoughMessage,
			},
		},
	}

	for _, test := range tests {

		if test.initData != nil {
			test.initData()
		}

		c := &Controller{}
		reply, err := c.ModifyUserPoint(context.Background(), test.input)
		assert.Nil(t, err)

		if test.wantError != nil {
			assert.Equal(t, test.wantError, reply.Err, test.input)
			continue
		} else {
			assert.Nil(t, reply.Err, test.input)
		}
	}
}
