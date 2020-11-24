package address

import (
	"dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/api/errors"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"golang.org/x/net/context"

	address_service "dev-gitlab.wanxingrowth.com/fanli/user/pkg/address"
	"dev-gitlab.wanxingrowth.com/fanli/user/pkg/model"
	"dev-gitlab.wanxingrowth.com/fanli/user/pkg/rpc/address/rules"
	"dev-gitlab.wanxingrowth.com/fanli/user/pkg/rpc/protos"
	"dev-gitlab.wanxingrowth.com/fanli/user/pkg/utils/log"
	"dev-gitlab.wanxingrowth.com/fanli/user/pkg/utils/validate"
)

func (_ Controller) SetAddress(ctx context.Context, req *protos.SetAddressRequest) (*protos.SetAddressReply, error) {
	var (
		err    error
		logger = log.GetLogger()
	)

	err = ValidateSetAddress(req)
	if err != nil {
		return &protos.SetAddressReply{Err: convertError(err, errors.CodeRequestParamError)}, nil
	}

	addressData := req.Address

	if !validate.CheckMobile(addressData.Mobile) {

		return &protos.SetAddressReply{
			Err: convertError(err, errors.CodeRequestParamError),
		}, nil
	}

	address := &model.UserAddress{
		Id:           addressData.Id,
		UserId:       addressData.UserId,
		ReceiverName: addressData.ReceiverName,
		Gender:       uint8(addressData.Gender),
		Country:      addressData.Country,
		Province:     addressData.Province,
		City:         addressData.City,
		District:     addressData.District,
		Address:      addressData.Address,
		Tag:          uint8(addressData.Tag),
		Mobile:       addressData.Mobile,
		Type:         uint8(addressData.Type),
	}

	err = address_service.Update(address)
	if err != nil {
		logger.WithError(err).Error("create address failed")
		return &protos.SetAddressReply{
			Err: convertError(err, errors.CodeServerInternalError),
		}, nil
	}

	return &protos.SetAddressReply{
		Address: convertAddressModelToProtobuf(address),
	}, nil
}

func ValidateSetAddress(req *protos.SetAddressRequest) error {

	return validation.ValidateStruct(req.Address,
		validation.Field(&req.Address.ReceiverName, rules.ReceiverName...),
		validation.Field(&req.Address.UserId, rules.UserIdRules...),
		validation.Field(&req.Address.Id, rules.IdRules...),
		validation.Field(&req.Address.Gender, rules.Gender...),
		validation.Field(&req.Address.Country, rules.Country...),
		validation.Field(&req.Address.Province, rules.Province...),
		validation.Field(&req.Address.City, rules.City...),
		validation.Field(&req.Address.District, rules.District...),
		validation.Field(&req.Address.Address, rules.Address...),
		validation.Field(&req.Address.Mobile, rules.Mobile...),
	)
}
