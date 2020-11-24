package user

import (
	"fmt"

	"dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/api/errors"
	rpcLog "dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/api/rpc/utils/log"
	"dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/data/database"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/jinzhu/gorm"
	"golang.org/x/net/context"

	"dev-gitlab.wanxingrowth.com/fanli/user/pkg/constant"
	"dev-gitlab.wanxingrowth.com/fanli/user/pkg/model"
	"dev-gitlab.wanxingrowth.com/fanli/user/pkg/rpc/protos"
	"dev-gitlab.wanxingrowth.com/fanli/user/pkg/rpc/user/rules"
	"dev-gitlab.wanxingrowth.com/fanli/user/pkg/utils/log"
)

func (_ Controller) ModifyUserInfo(ctx context.Context, req *protos.ModifyUserInfoRequest) (*protos.ModifyUserInfoReply, error) {

	logger := rpcLog.WithRequestId(ctx, log.GetLogger())

	if req == nil {

		logger.Error("request data is nil")
		return nil, fmt.Errorf("request data is nil")
	}

	err := validateModifyUserInfo(req)
	if err != nil {

		return &protos.ModifyUserInfoReply{
			Err: convertError(err, errors.CodeRequestParamError),
		}, nil
	}

	user := &model.User{}
	err = database.GetDB(constant.DatabaseConfigKey).
		Where(model.User{
			UserId: req.GetUserInfo().GetUserId(),
		}).
		First(user).Error
	if err != nil {
		if !gorm.IsRecordNotFoundError(err) {
			logger.WithError(err).Error("get user data error")
			return &protos.ModifyUserInfoReply{
				Err: convertError(err, errors.CodeServerInternalError),
			}, nil
		}

		logger.WithError(err).Info("user not exist")
		return &protos.ModifyUserInfoReply{
			Err: convertError(err, errors.CodeRequestParamError),
		}, nil
	}

	user.Name = req.GetUserInfo().GetName()
	user.Gender = uint8(req.GetUserInfo().GetGender())
	user.HeadImage = req.GetUserInfo().GetHeadImage()

	birthday, err := parseStringTime(req.GetUserInfo().GetBirthday())
	if err == nil {
		user.Birthday = birthday
	}

	err = database.GetDB(constant.DatabaseConfigKey).
		Model(model.User{}).
		Where(model.User{UserId: user.UserId}).
		Update(user).
		Error

	if err != nil {
		logger.WithError(err).Error("update user data error")
		return &protos.ModifyUserInfoReply{
			Err: convertError(err, errors.CodeServerInternalError),
		}, nil
	}

	return &protos.ModifyUserInfoReply{}, nil
}

func validateModifyUserInfo(req *protos.ModifyUserInfoRequest) error {

	return validation.ValidateStruct(req,
		validation.Field(&req.UserInfo, rules.UserInfoRules...),
	)
}
