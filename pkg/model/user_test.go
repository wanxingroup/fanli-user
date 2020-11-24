package model

import (
	"testing"
	"time"

	"dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/data/database"
	"github.com/stretchr/testify/assert"

	"dev-gitlab.wanxingrowth.com/fanli/user/pkg/constant"
)

func TestUser_TableName(t *testing.T) {

	user := User{}
	assert.Equal(t, TableNameUser, user.TableName())
}

func TestUser_Save(t *testing.T) {

	tests := []struct {
		input *User
	}{
		{
			input: &User{
				UserId:   100,
				Mobile:   "",
				Name:     "",
				Gender:   0,
				Birthday: time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			},
		},
		{
			input: &User{
				UserId:   101,
				Mobile:   "",
				Name:     "",
				Gender:   0,
				Birthday: time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
				ShopRelated: []*UserShop{
					{
						OpenId:   "abcdefg",
						UserId:   101,
						ShopId:   29299,
						OpenType: 0,
					},
				},
			},
		},
	}

	for _, test := range tests {

		assert.Nil(t, database.GetDB(constant.DatabaseConfigKey).Save(test.input).Error, test)
	}
}
