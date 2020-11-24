package rules

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"

	"dev-gitlab.wanxingrowth.com/fanli/user/pkg/constant"
)

var IdRules = []validation.Rule{
	validation.Required.ErrorObject(validation.NewError(constant.ErrorCodeAddressIdRequired, constant.ErrorMessageAddressIdRequired)),
}

var UserIdRules = []validation.Rule{
	validation.Required.ErrorObject(validation.NewError(constant.ErrorCodeUserIdRequired, constant.ErrorMessageUserIdRequired)),
}

var ReceiverName = []validation.Rule{
	validation.Required.ErrorObject(validation.NewError(constant.ErrorCodeReceiverNameRequired, constant.ErrorMessageReceiverNameRequired)),
}

var Gender = []validation.Rule{
	validation.Required.ErrorObject(validation.NewError(constant.ErrorCodeGenderRequired, constant.ErrorMessageGenderRequired)),
}

var Country = []validation.Rule{
	validation.Required.ErrorObject(validation.NewError(constant.ErrorCodeCountryRequired, constant.ErrorMessageCountryRequired)),
}

var Province = []validation.Rule{
	validation.Required.ErrorObject(validation.NewError(constant.ErrorCodeProvinceRequired, constant.ErrorMessageProvinceRequired)),
}

var City = []validation.Rule{
	validation.Required.ErrorObject(validation.NewError(constant.ErrorCodeCityRequired, constant.ErrorMessageCityRequired)),
}

var District = []validation.Rule{
	validation.Required.ErrorObject(validation.NewError(constant.ErrorCodeDistrictRequired, constant.ErrorMessageDistrictRequired)),
}

var Address = []validation.Rule{
	validation.Required.ErrorObject(validation.NewError(constant.ErrorCodeAddressRequired, constant.ErrorMessageAddressRequired)),
}

var Mobile = []validation.Rule{
	validation.Required.ErrorObject(validation.NewError(constant.ErrorCodeMobileRequired, constant.ErrorMessageMobileRequired)),
}
