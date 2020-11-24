package user

import (
	"context"
	"strconv"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"

	"dev-gitlab.wanxingrowth.com/fanli/user/pkg/rpc/errorcode"
	"dev-gitlab.wanxingrowth.com/fanli/user/pkg/rpc/protos"
	"dev-gitlab.wanxingrowth.com/fanli/user/pkg/user"
	"dev-gitlab.wanxingrowth.com/fanli/user/pkg/utils/log"
)

const defaultPage = 1
const defaultPageSize = 20
const maxPageSize = 100

func (cor Controller) GetUserList(ctx context.Context, req *protos.GetUserListRequest) (reply *protos.GetUserListReply, err error) {

	logger := log.GetLogger().WithField("requestData", req)

	logger.Info("Get User list requestData")

	var page = uint64(defaultPage)
	var pageSize = uint64(defaultPageSize)

	if req.GetPage() > 0 {
		page = req.GetPage()
	}

	if req.GetPageSize() > 0 {
		pageSize = req.GetPageSize()
	}

	if pageSize > maxPageSize {
		pageSize = maxPageSize
	}
	err = cor.ValidateGetUserListRequest(req)
	if err != nil {
		return &protos.GetUserListReply{Err: &protos.Error{
			Code:    errorcode.RequestParameterError,
			Message: err.Error(),
		}}, nil
	}

	conditions := map[string]interface{}{
		"shopId": req.GetShopId(),
	}

	if !validation.IsEmpty(req.GetMobile()) {
		conditions["mobile"] = req.GetMobile()
	}

	if !validation.IsEmpty(req.GetUserIds()) {
		conditions["userIds"] = req.GetUserIds()
	}

	if !validation.IsEmpty(req.GetUserName()) {
		conditions["name"] = req.GetUserName()
	}

	if !validation.IsEmpty(req.GetRegStartTime()) && !validation.IsEmpty(req.GetRegEndTime()) {
		conditions["getRegStartTime"] = req.GetRegStartTime()
		conditions["getRegEndTime"] = req.GetRegEndTime()
	}

	userList, count, err := user.GetUserListByCondition(conditions, page, pageSize)

	if err != nil {
		logger.WithError(err).Error("get user list failed")
		return &protos.GetUserListReply{
			Err: &protos.Error{
				Code:    errorcode.InternalError,
				Message: err.Error(),
			},
		}, nil
	}

	userInfoList := make([]*protos.UserInfo, len(userList))
	for key, userInfo := range userList {
		gender, _ := strconv.ParseUint(string(userInfo.Gender), 10, 64)

		userInfoList[key] = &protos.UserInfo{
			UserId:    userInfo.UserId,
			Mobile:    userInfo.Mobile,
			Name:      userInfo.Name,
			Gender:    gender,
			Birthday:  userInfo.Birthday.Format("2006-01-02 15:04:05"),
			OpenId:    userInfo.OpenId,
			CreatedAt: userInfo.CreatedAt.Format("2006-01-02 15:04:05"),
		}
	}

	return &protos.GetUserListReply{
		UserInfoList: userInfoList,
		Count:        count,
	}, nil
}
func (cor Controller) ValidateGetUserListRequest(req *protos.GetUserListRequest) error {
	return validation.ValidateStruct(req,
		validation.Field(&req.ShopId, validation.Required, validation.Min(uint64(0)).Exclusive().
			ErrorObject(validation.NewError(strconv.Itoa(errorcode.ShopIdError), errorcode.ShopIdErrorMessage)),
		),
		validation.Field(&req.Mobile, validation.Length(0, 13), is.Digit),
	)
}
