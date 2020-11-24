package model

import (
	"testing"

	"dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/data/database"
	databases "dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/data/database/models"
	"github.com/stretchr/testify/assert"

	"dev-gitlab.wanxingrowth.com/fanli/user/pkg/constant"
)

func TestUserAddress_TableName(t *testing.T) {

	address := UserAddress{}
	assert.Equal(t, TableNameUserAddress, address.TableName())
}

func TestUserAddress_Save(t *testing.T) {
	tests := []struct {
		input *UserAddress
	}{
		{
			input: &UserAddress{
				Id:           1,
				UserId:       2,
				ReceiverName: "张三",
				Gender:       0,
				Country:      "中国",
				Province:     "上海",
				City:         "上海",
				District:     "浦东新区",
				Address:      "张江路 668 弄",
				Tag:          0,
				Mobile:       "13167185909",
				Type:         0,
				Time:         databases.Time{},
			},
		},
	}

	for _, test := range tests {

		assert.Nil(t, database.GetDB(constant.DatabaseConfigKey).Save(test.input).Error, test)
	}
}
