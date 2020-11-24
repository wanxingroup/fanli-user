package address

import (
	"strconv"

	validation "github.com/go-ozzo/ozzo-validation/v4"

	"dev-gitlab.wanxingrowth.com/fanli/user/pkg/model"
	"dev-gitlab.wanxingrowth.com/fanli/user/pkg/rpc/protos"
)

const (
	DefaultPage     = 1
	DefaultPageSize = 5
	MaxPageSize     = 100
)

func convertError(err error, defaultErrorCode int64) *protos.Error {

	if err == nil {
		return nil
	}

	switch validationError := err.(type) {
	case validation.Error:

		code, convertError := strconv.ParseInt(validationError.Code(), 10, 64)
		if convertError != nil {
			return &protos.Error{
				Code:    defaultErrorCode,
				Message: validationError.Error(),
			}
		}
		return &protos.Error{
			Code:    code,
			Message: validationError.Error(),
		}

	case validation.Errors:
		if len(validationError) > 1 {
			return &protos.Error{
				Code:    defaultErrorCode,
				Message: validationError.Error(),
			}
		}

		for _, validationErr := range validationError {
			return convertError(validationErr, defaultErrorCode)
		}
	}

	return &protos.Error{
		Code:    defaultErrorCode,
		Message: err.Error(),
	}
}

func convertAddressModelToProtobuf(address *model.UserAddress) *protos.Address {

	return &protos.Address{
		Id:           address.Id,
		UserId:       address.UserId,
		ReceiverName: address.ReceiverName,
		Gender:       uint64(address.Gender),
		Country:      address.Country,
		Province:     address.Province,
		City:         address.City,
		District:     address.District,
		Address:      address.Address,
		Tag:          uint32(address.Tag),
		Mobile:       address.Mobile,
		Type:         uint32(address.Type),
	}
}
