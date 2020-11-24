package user

import (
	"fmt"

	"dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/api/errors"
	rpcLog "dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/api/rpc/utils/log"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"golang.org/x/net/context"

	"dev-gitlab.wanxingrowth.com/fanli/user/pkg/model"
	"dev-gitlab.wanxingrowth.com/fanli/user/pkg/rpc/errorcode"
	"dev-gitlab.wanxingrowth.com/fanli/user/pkg/rpc/protos"
	"dev-gitlab.wanxingrowth.com/fanli/user/pkg/rpc/user/rules"
	"dev-gitlab.wanxingrowth.com/fanli/user/pkg/utils/log"
)

func (_ Controller) ModifyUserPoint(ctx context.Context, req *protos.ModifyUserPointRequest) (*protos.ModifyUserPointReply, error) {

	logger := rpcLog.WithRequestId(ctx, log.GetLogger())

	if req == nil {

		logger.Error("request data is nil")
		return nil, fmt.Errorf("request data is nil")
	}

	err := validateModifyUserPoint(req)

	if err != nil {

		return &protos.ModifyUserPointReply{
			Err: convertError(err, errors.CodeRequestParamError),
		}, nil
	}

	if req.Type != model.UserIncreasePoint && req.Type != model.UserDecreasePoint {
		return &protos.ModifyUserPointReply{
			Err: convertError(err, errors.CodeRequestParamError),
		}, nil
	}

	reqData := &model.ModifyUserPoint{
		UserId:  req.UserId,
		Remark:  req.Remark,
		OrderId: req.OrderId,
		Point:   req.Point,
		Type:    req.Type,
	}

	// 判断用户 的积分是否满足
	userInfo, _ := getUser(req.UserId)
	if reqData.Type == model.UserDecreasePoint && userInfo.Point < reqData.Point {
		return &protos.ModifyUserPointReply{
			Err: &protos.Error{
				Code:    errorcode.UserPointIsNotEnough,
				Message: errorcode.UserPointIsNotEnoughMessage,
			}}, nil
	}

	err = ModifyUserPoint(reqData)
	if err != nil {
		logger.WithError(err).Error("update User point error")
		return &protos.ModifyUserPointReply{
			Err: convertError(err, errors.CodeServerInternalError),
		}, nil
	}

	userInfo, _ = getUser(req.UserId)

	return &protos.ModifyUserPointReply{
		Point: userInfo.Point,
		Err:   nil,
	}, nil
}

func validateModifyUserPoint(req *protos.ModifyUserPointRequest) error {

	return validation.ValidateStruct(req,
		validation.Field(&req.Point, rules.PointRules...),
		validation.Field(&req.UserId, rules.UserIdRules...),
	)
}
