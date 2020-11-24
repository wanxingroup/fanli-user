package errorcode

const (
	InternalError    = 500000
	InternalErrorMsg = "服务异常"

	EntityNotFound               = 404000
	EntityNotFoundMsg            = "记录不存在"
	RequestParameterError        = 404001
	RequestParameterErrorMessage = "输入参数错误"
	ShopIdError                  = 404002
	ShopIdErrorMessage           = "shopId 传参有误"
	DeleteAddressError           = 402003 // 删除地址失败
	DeleteAddressErrorMessage    = "删除地址失败"
	UserPointIsNotEnough         = 402004
	UserPointIsNotEnoughMessage  = "用户积分不够"
)
