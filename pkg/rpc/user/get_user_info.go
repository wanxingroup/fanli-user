package user

import (
	"fmt"

	"dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/api/errors"
	rpcLog "dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/api/rpc/utils/log"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/jinzhu/gorm"
	"golang.org/x/net/context"

	"dev-gitlab.wanxingrowth.com/fanli/user/pkg/rpc/protos"
	"dev-gitlab.wanxingrowth.com/fanli/user/pkg/rpc/user/rules"
	"dev-gitlab.wanxingrowth.com/fanli/user/pkg/utils/log"
)

func (_ Controller) GetUserInfo(ctx context.Context, req *protos.GetUserInfoRequest) (*protos.GetUserInfoReply, error) {

	logger := rpcLog.WithRequestId(ctx, log.GetLogger())

	if req == nil {

		logger.Error("request data is nil")
		return nil, fmt.Errorf("request data is nil")
	}

	logger = logger.WithField("requestData", req)

	err := validateGetUserInfo(req)
	if err != nil {

		return &protos.GetUserInfoReply{
			Err: convertError(err, errors.CodeRequestParamError),
		}, nil
	}

	user, err := getUser(req.GetUserId())
	if err != nil {
		if !gorm.IsRecordNotFoundError(err) {
			logger.WithError(err).Error("get user data error")
			return &protos.GetUserInfoReply{
				Err: convertError(err, errors.CodeServerInternalError),
			}, nil
		}

		logger.WithError(err).Info("user not exist")
		return &protos.GetUserInfoReply{
			Err: convertError(err, errors.CodeRequestParamError),
		}, nil
	}

	return &protos.GetUserInfoReply{
		UserInfo: convertUserModelToProtobuf(user, req.GetShopId()),
	}, nil
}

func validateGetUserInfo(req *protos.GetUserInfoRequest) error {

	return validation.ValidateStruct(req,
		validation.Field(&req.UserId, rules.UserIdRules...),
		validation.Field(&req.ShopId, rules.ShopIdRules...),
	)
}
