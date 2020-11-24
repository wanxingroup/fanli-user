package user

import (
	"fmt"

	"dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/api/errors"
	rpcLog "dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/api/rpc/utils/log"
	"dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/data/database"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"golang.org/x/net/context"

	"dev-gitlab.wanxingrowth.com/fanli/user/pkg/constant"
	"dev-gitlab.wanxingrowth.com/fanli/user/pkg/model"
	"dev-gitlab.wanxingrowth.com/fanli/user/pkg/rpc/protos"
	"dev-gitlab.wanxingrowth.com/fanli/user/pkg/rpc/user/rules"
	"dev-gitlab.wanxingrowth.com/fanli/user/pkg/utils/log"
)

func (_ Controller) SetInviter(ctx context.Context, req *protos.SetInviterRequest) (*protos.SetInviterReply, error) {

	logger := rpcLog.WithRequestId(ctx, log.GetLogger())
	if req == nil {
		logger.Error("request data is nil")
		return nil, fmt.Errorf("request data is nil")
	}

	err := validateSetInviter(req)
	if err != nil {
		return &protos.SetInviterReply{
			Err: convertError(err, errors.CodeRequestParamError),
		}, nil
	}

	lastInviter := &model.UserInviterHistory{}
	err = database.GetDB(constant.DatabaseConfigKey).Where(model.UserInviterHistory{UserId: req.GetUserId()}).Order("`id` DESC").First(lastInviter).Error
	if err == nil && lastInviter.InviterId > 0 {
		err = database.GetDB(constant.DatabaseConfigKey).
			Model(model.User{}).
			Where(&model.User{UserId: req.GetUserId()}).
			Update(&model.User{
				InviterId:   lastInviter.InviterId,
				InviterType: lastInviter.InviterType,
			}).
			Error

		if err != nil {
			logger.WithError(err).Error("update user inviter data error")
			return &protos.SetInviterReply{
				Err: convertError(err, errors.CodeServerInternalError),
			}, nil
		}

		return &protos.SetInviterReply{
			InviterId:   lastInviter.InviterId,
			InviterType: uint32(lastInviter.InviterType),
		}, nil
	}

	return nil, nil
}

func validateSetInviter(req *protos.SetInviterRequest) error {
	return validation.ValidateStruct(req,
		validation.Field(&req.UserId, rules.UserIdRules...),
	)
}
