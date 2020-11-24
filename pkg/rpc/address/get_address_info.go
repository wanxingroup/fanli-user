package address

import (
	"golang.org/x/net/context"

	"dev-gitlab.wanxingrowth.com/fanli/user/pkg/address"
	"dev-gitlab.wanxingrowth.com/fanli/user/pkg/rpc/errorcode"
	"dev-gitlab.wanxingrowth.com/fanli/user/pkg/rpc/protos"
	"dev-gitlab.wanxingrowth.com/fanli/user/pkg/utils/log"
)

func (_ Controller) GetAddressInfo(ctx context.Context, req *protos.GetAddressInfoRequest) (*protos.GetAddressInfoReply, error) {

	logger := log.GetLogger().WithField("requestData", req)

	if req.GetUserId() == 0 || req.GetAddressId() == 0 {
		return &protos.GetAddressInfoReply{Err: &protos.Error{
			Code:    errorcode.RequestParameterError,
			Message: errorcode.RequestParameterErrorMessage,
		}}, nil
	}

	conditions := map[string]interface{}{
		"userId":    req.GetUserId(),
		"addressId": req.GetAddressId(),
	}

	info, err := address.Get(conditions)

	if err != nil {
		logger.WithError(err).Error("get address info failed")
		return &protos.GetAddressInfoReply{
			Err: &protos.Error{
				Code:    errorcode.InternalError,
				Message: err.Error(),
			},
		}, nil
	}

	return &protos.GetAddressInfoReply{
		Address: convertAddressModelToProtobuf(info),
	}, nil
}
