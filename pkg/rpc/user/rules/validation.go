package rules

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"

	"dev-gitlab.wanxingrowth.com/fanli/user/pkg/constant"
	"dev-gitlab.wanxingrowth.com/fanli/user/pkg/rpc/protos"
)

var ShopIdRules = []validation.Rule{
	validation.Required.ErrorObject(validation.NewError(constant.ErrorCodeShopIdRequired, constant.ErrorMessageShopIdRequired)),
}

var UserIdRules = []validation.Rule{
	validation.Required.ErrorObject(validation.NewError(constant.ErrorCodeUserIdRequired, constant.ErrorMessageUserIdRequired)),
}

var OpenIdRules = []validation.Rule{
	validation.Required.ErrorObject(validation.NewError(constant.ErrorCodeOpenIdRequired, constant.ErrorMessageOpenIdRequired)),
}

var MobileRules = []validation.Rule{
	validation.Required.ErrorObject(validation.NewError(constant.ErrorCodeMobileRequired, constant.ErrorMessageMobileRequired)),
}

var PointRules = []validation.Rule{
	validation.Required.ErrorObject(validation.NewError(constant.ErrorCodePointRequired, constant.ErrorMessagePointRequired)),
}

var TypeRules = []validation.Rule{
	validation.Required.ErrorObject(validation.NewError(constant.ErrorCodeTypeRequired, constant.ErrorMessageTypeRequired)),
}

var UserInfoRules = []validation.Rule{
	validation.Required.ErrorObject(validation.NewError(constant.ErrorCodeUserInfoRequired, constant.ErrormessageUserInfoRequired)),
	validation.By(func(value interface{}) error {
		userInfo, ok := value.(*protos.UserInfo)
		if !ok {
			return validation.NewError(constant.ErrorCodeUserInfoInvalid, constant.ErrormessageUserInfoInvalid)
		}

		return validateUserInfo(userInfo)
	}),
}

func validateUserInfo(req *protos.UserInfo) error {
	return validation.ValidateStruct(req,
		validation.Field(req.UserId, UserIdRules...),
	)
}
