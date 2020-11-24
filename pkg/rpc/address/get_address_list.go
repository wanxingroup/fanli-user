package address

import (
	"golang.org/x/net/context"

	"dev-gitlab.wanxingrowth.com/fanli/user/pkg/address"
	"dev-gitlab.wanxingrowth.com/fanli/user/pkg/rpc/errorcode"
	"dev-gitlab.wanxingrowth.com/fanli/user/pkg/rpc/protos"
	"dev-gitlab.wanxingrowth.com/fanli/user/pkg/utils/log"
)

func (_ Controller) GetAddressList(ctx context.Context, req *protos.GetAddressListRequest) (*protos.GetAddressListReply, error) {

	logger := log.GetLogger().WithField("requestData", req)

	var page = uint64(DefaultPage)
	var pageSize = uint64(DefaultPageSize)
	var lastId = uint64(0)

	if req.GetPage() > 0 {
		page = req.GetPage()
	}

	if req.GetPageSize() > 0 {
		pageSize = req.GetPageSize()
	}

	if pageSize > MaxPageSize {
		pageSize = MaxPageSize
	}

	var pageData map[string]uint64

	// 通过 Id 分页
	if req.GetLastId() > 0 {
		lastId = req.GetLastId()
		pageData = map[string]uint64{
			"pageSize": pageSize,
			"lastId":   lastId,
		}
	} else {
		pageData = map[string]uint64{
			"pageSize": pageSize,
			"pageNum":  page,
		}
	}

	if req.GetUserId() == 0 {
		return &protos.GetAddressListReply{Err: &protos.Error{
			Code:    errorcode.RequestParameterError,
			Message: errorcode.RequestParameterErrorMessage,
		}}, nil
	}

	conditions := map[string]interface{}{
		"userId": req.GetUserId(),
	}

	list, count, err := address.GetList(conditions, pageData)

	if err != nil {
		logger.WithError(err).Error("get address list failed")
		return &protos.GetAddressListReply{
			Err: &protos.Error{
				Code:    errorcode.InternalError,
				Message: err.Error(),
			},
		}, nil
	}

	returnList := make([]*protos.Address, len(list))
	for key, addressInfo := range list {

		returnList[key] = convertAddressModelToProtobuf(addressInfo)
	}

	return &protos.GetAddressListReply{
		List:  returnList,
		Count: count,
	}, nil
}
