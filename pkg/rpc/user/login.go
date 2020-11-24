package user

import (
	"fmt"
	"strconv"
	"strings"

	"dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/api/errors"
	rpcLog "dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/api/rpc/utils/log"
	"dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/data/database"
	idcreator "dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/utils/idcreator/snowflake"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/jinzhu/gorm"
	"golang.org/x/net/context"

	"dev-gitlab.wanxingrowth.com/fanli/user/pkg/constant"
	"dev-gitlab.wanxingrowth.com/fanli/user/pkg/model"
	"dev-gitlab.wanxingrowth.com/fanli/user/pkg/rpc/protos"
	"dev-gitlab.wanxingrowth.com/fanli/user/pkg/rpc/user/rules"
	"dev-gitlab.wanxingrowth.com/fanli/user/pkg/utils/log"
)

func (_ Controller) Login(ctx context.Context, req *protos.LoginRequest) (*protos.LoginReply, error) {

	logger := rpcLog.WithRequestId(ctx, log.GetLogger())

	if req == nil {

		logger.Error("request data is nil")
		return nil, fmt.Errorf("request data is nil")
	}

	err := validateLoginRequestData(req)
	if err != nil {

		return &protos.LoginReply{
			Err: convertError(err, errors.CodeRequestParamError),
		}, nil
	}

	userShopRelated := &model.UserShop{}
	err = database.GetDB(constant.DatabaseConfigKey).Where(&model.UserShop{
		OpenId: req.GetOpenId(),
		ShopId: req.GetShopId(),
	}).First(userShopRelated).Error

	if err != nil {
		if !gorm.IsRecordNotFoundError(err) {
			logger.WithError(err).Error("get user shop record error")
			return &protos.LoginReply{
				Err: convertError(err, errors.CodeServerInternalError),
			}, nil
		}

		user, err := createNewUser()
		if err != nil {
			logger.WithError(err).Error("create new user error")
			return &protos.LoginReply{
				Err: convertError(err, errors.CodeServerInternalError),
			}, nil
		}

		userShopRelated, err = createUserShopRelated(req.GetShopId(), user.UserId, req.GetOpenId())
		if err != nil {
			logger.WithError(err).Error("create user and shop related error")
			return &protos.LoginReply{
				Err: convertError(err, errors.CodeServerInternalError),
			}, nil
		}
	}

	user, err := getUser(userShopRelated.UserId)
	if err != nil {
		logger.WithError(err).Error("get user error")
		return &protos.LoginReply{
			Err: convertError(err, errors.CodeServerInternalError),
		}, nil
	}

	// 添加邀请记录
	if user.InviterId == 0 && len(req.GetInviterCode()) > 0 {
		inviterCode := strings.Split(req.GetInviterCode(), "-")
		inviterType, _ := strconv.ParseUint(inviterCode[1], 10, 64)
		inviterId, _ := strconv.ParseUint(inviterCode[2], 10, 64)
		if inviterId > 0 && inviterType > 0 {
			inviterHistory := model.UserInviterHistory{
				Id:          idcreator.NextID(),
				UserId:      user.UserId,
				InviterId:   inviterId,
				InviterType: uint8(inviterType),
			}

			err = database.GetDB(constant.DatabaseConfigKey).Create(inviterHistory).Error
			if err != nil {
				logger.WithError(err).Error("create inviter history err")
			}
		}

	}

	return &protos.LoginReply{
		UserInfo: convertUserModelToProtobuf(user, req.GetShopId()),
	}, nil
}

func validateLoginRequestData(req *protos.LoginRequest) error {

	return validation.ValidateStruct(req,
		validation.Field(&req.ShopId, rules.ShopIdRules...),
		validation.Field(&req.OpenId, rules.OpenIdRules...),
	)
}
