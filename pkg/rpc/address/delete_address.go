package address

import (
	"golang.org/x/net/context"

	"dev-gitlab.wanxingrowth.com/fanli/user/pkg/address"
	"dev-gitlab.wanxingrowth.com/fanli/user/pkg/rpc/errorcode"
	"dev-gitlab.wanxingrowth.com/fanli/user/pkg/rpc/protos"
	"dev-gitlab.wanxingrowth.com/fanli/user/pkg/utils/log"
)

func (controller Controller) DeleteAddress(ctx context.Context, req *protos.DeleteAddressRequest) (*protos.DeleteAddressReply, error) {

	logger := log.GetLogger().WithField("requestData", req)

	if req.GetUserId() == 0 {
		return &protos.DeleteAddressReply{Err: &protos.Error{
			Code:    errorcode.RequestParameterError,
			Message: errorcode.RequestParameterErrorMessage,
		}}, nil
	}

	if req.GetAddressId() == 0 {
		return &protos.DeleteAddressReply{Err: &protos.Error{
			Code:    errorcode.RequestParameterError,
			Message: errorcode.RequestParameterErrorMessage,
		}}, nil
	}

	if err := address.Delete(req.GetUserId(), req.GetAddressId()); err != nil {
		logger.WithError(err).Error("delete address failed")
		return &protos.DeleteAddressReply{
			Err: &protos.Error{
				Code:    errorcode.DeleteAddressError,
				Message: errorcode.DeleteAddressErrorMessage,
			},
		}, nil
	}

	return &protos.DeleteAddressReply{}, nil

}
