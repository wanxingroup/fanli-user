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

func (_ Controller) BindMobile(ctx context.Context, req *protos.BindMobileRequest) (*protos.BindMobileReply, error) {

	logger := rpcLog.WithRequestId(ctx, log.GetLogger())

	if req == nil {

		logger.Error("request data is nil")
		return nil, fmt.Errorf("request data is nil")
	}

	err := validateBindMobileRequestData(req)
	if err != nil {

		return &protos.BindMobileReply{
			Err: convertError(err, errors.CodeRequestParamError),
		}, nil
	}

	// 尝试检测手机号是否已经绑定过用户 ID
	user := &model.User{}
	err = database.GetDB(constant.DatabaseConfigKey).Where(model.User{Mobile: req.GetMobile()}).First(user).Error
	if err != nil {
		if !gorm.IsRecordNotFoundError(err) {
			logger.WithError(err).Error("get user data error")
			return &protos.BindMobileReply{
				Err: convertError(err, errors.CodeServerInternalError),
			}, nil
		}

		// 如果手机号没有绑定过用户，获取当前用户，进行更新绑定
		user, err := getUser(req.GetUserId())
		if err != nil {
			if gorm.IsRecordNotFoundError(err) {
				logger.WithError(err).Info("user not exist")
				return &protos.BindMobileReply{
					Err: convertError(err, errors.CodeRequestParamError),
				}, nil
			}
		}

		user.Mobile = req.GetMobile()
		err = database.GetDB(constant.DatabaseConfigKey).
			Model(model.User{}).
			Where(&model.User{UserId: user.UserId}).
			Update(&model.User{
				Mobile: user.Mobile,
			}).
			Error
		if err != nil {
			logger.WithError(err).Error("update user data error")
			return &protos.BindMobileReply{
				Err: convertError(err, errors.CodeServerInternalError),
			}, nil
		}

		return &protos.BindMobileReply{
			UserInfo: convertUserModelToProtobuf(user, req.GetShopId()),
		}, nil
	}

	// 如果手机号绑定过用户 ID
	// 如果用户 ID 对应的 OpenID 里面是否包含当前请求的 OpenID，则返回当前用户
	if req.GetUserId() == user.UserId {
		return &protos.BindMobileReply{
			UserInfo: convertUserModelToProtobuf(user, req.GetShopId()),
		}, nil
	}

	// 需要执行合并用户操作（对支持合并操作的服务都发送合并数据请求），然后返回当前用户
	err = mergeUser(req.GetUserId(), user.UserId)
	if err != nil {
		logger.WithError(err).Error("merge user error")
	}

	return &protos.BindMobileReply{
		UserInfo: convertUserModelToProtobuf(user, req.GetShopId()),
	}, nil
}

func validateBindMobileRequestData(req *protos.BindMobileRequest) error {

	return validation.ValidateStruct(req,
		validation.Field(&req.ShopId, rules.ShopIdRules...),
		validation.Field(&req.Mobile, rules.MobileRules...),
		validation.Field(&req.UserId, rules.UserIdRules...),
	)
}

func mergeUser(fromUserId, toUserId uint64) error {
	// 合并用户的 UserId 以及 OpenId
	mergeUserShopErr := mergeUserShop(fromUserId, toUserId)
	if mergeUserShopErr != nil {
		log.GetLogger().WithError(mergeUserShopErr)
	}

	// 这一步将来再做，目前协商好，执行流程是所有用户购买会员卡前，都需要绑定手机号，所以需要合并数据的就只有用户信息
	mergeUserErr := mergeUserData(fromUserId, toUserId)
	if mergeUserErr != nil {
		log.GetLogger().WithError(mergeUserErr)
	}

	mergeInviterHistoryErr := mergeInviterHistoryData(fromUserId, toUserId)
	if mergeInviterHistoryErr != nil {
		log.GetLogger().WithError(mergeInviterHistoryErr)
	}

	return nil
}

// 合并所有的邀请记录
func mergeInviterHistoryData(fromUserId, toUserId uint64) (err error) {
	err = database.GetDB(constant.DatabaseConfigKey).Model(&model.UserInviterHistory{}).Where("userId = ?", fromUserId).Update("userId", toUserId).Error
	return
}

func mergeUserData(fromUserId uint64, toUserId uint64) error {
	fromUser, err := getUser(fromUserId)
	if err != nil {
		return err
	}

	toUser, err := getUser(toUserId)
	if err != nil {
		return err
	}

	if len(toUser.Name) == 0 && len(fromUser.Name) > 0 {
		toUser.Name = fromUser.Name
	}

	if toUser.Gender == model.GenderSecrecy && fromUser.Gender != model.GenderSecrecy {
		toUser.Gender = fromUser.Gender
	}

	if toUser.Birthday.IsZero() && !fromUser.Birthday.IsZero() {
		toUser.Birthday = fromUser.Birthday
	}

	return database.GetDB(constant.DatabaseConfigKey).Save(toUser).Error
}

func mergeUserShop(fromUserId uint64, toUserId uint64) (err error) {
	err = database.GetDB(constant.DatabaseConfigKey).Model(&model.UserShop{}).Where("userId = ?", fromUserId).Update("userId", toUserId).Error
	return
}
