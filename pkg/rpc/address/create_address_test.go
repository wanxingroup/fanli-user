package address

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"

	"dev-gitlab.wanxingrowth.com/fanli/user/pkg/rpc/protos"
)

func TestController_CreateAddress(t *testing.T) {
	tests := []struct {
		input     *protos.CreateAddressRequest
		want      *protos.CreateAddressReply
		wantError *protos.Error
	}{
		{
			input: &protos.CreateAddressRequest{Address: &protos.Address{
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
			}},
			want: &protos.CreateAddressReply{
				Err: nil,
				Address: &protos.Address{
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
				},
			},
		},
	}

	for _, test := range tests {

		c := &Controller{}
		reply, err := c.CreateAddress(context.Background(), test.input)
		assert.Nil(t, err)

		if test.wantError != nil {
			assert.Equal(t, test.wantError, reply.Err, test.input)
			continue
		} else {
			assert.Nil(t, reply.Err, test.input)
		}

		assert.NotNil(t, reply.Address, test.input)

		assert.Greater(t, reply.Address.Id, uint64(0), test.input)

	}
}
