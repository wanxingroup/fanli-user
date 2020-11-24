package address

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"

	address_service "dev-gitlab.wanxingrowth.com/fanli/user/pkg/address"
	"dev-gitlab.wanxingrowth.com/fanli/user/pkg/model"
	"dev-gitlab.wanxingrowth.com/fanli/user/pkg/rpc/protos"
)

func TestController_DeleteAddress(t *testing.T) {

	// todo 永康 完善单测
	tests := []struct {
		initData  func() error
		input     *protos.DeleteAddressRequest
		check     func() error
		wantError *protos.Error
	}{
		{
			initData: func() error {
				return address_service.Create(&model.UserAddress{
					Id:           101,
					UserId:       1024,
					ReceiverName: "张三",
					Gender:       1,
					Country:      "中国",
					Province:     "上海",
					City:         "上海",
					District:     "浦东新区",
					Address:      "张江路 668 弄",
					Tag:          0,
					Mobile:       "13167185909",
					Type:         0,
				})
			},
			check: func() error {
				conditions := map[string]interface{}{
					"userId": 1024,
					"id":     101,
				}
				_, err := address_service.Get(conditions)
				if err != nil {
					return err
				}

				// todo 永康 完善单测
				return nil
			},
			input: &protos.DeleteAddressRequest{
				AddressId: 101,
				UserId:    1024,
			},
		},
	}

	for _, test := range tests {

		if test.initData != nil {
			assert.Nil(t, test.initData(), test.input)
		}

		c := &Controller{}
		reply, err := c.DeleteAddress(context.Background(), test.input)
		assert.Nil(t, err)
		if test.wantError != nil || reply.Err != nil {
			assert.Equal(t, test.wantError, reply.Err)
			continue
		}
		if test.check != nil {
			assert.Nil(t, test.check(), test)
		}
	}
}
