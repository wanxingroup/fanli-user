package user

import (
	"strconv"
	"time"

	"dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/data/database"
	idcreator "dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/utils/idcreator/snowflake"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/jinzhu/gorm"

	"dev-gitlab.wanxingrowth.com/fanli/user/pkg/constant"
	"dev-gitlab.wanxingrowth.com/fanli/user/pkg/model"
	"dev-gitlab.wanxingrowth.com/fanli/user/pkg/rpc/protos"
)

func convertError(err error, defaultErrorCode int64) *protos.Error {

	if err == nil {
		return nil
	}

	if validationError, ok := err.(validation.Error); ok {

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
	}

	return &protos.Error{
		Code:    defaultErrorCode,
		Message: err.Error(),
	}
}

func createNewUser() (user *model.User, err error) {

	user = &model.User{
		UserId:   idcreator.NextID(),
		Birthday: time.Unix(0, 0),
	}

	err = database.GetDB(constant.DatabaseConfigKey).Save(user).Error
	return
}

func createUserShopRelated(shopId uint64, userId uint64, openId string) (userShopRelated *model.UserShop, err error) {

	userShopRelated = &model.UserShop{
		OpenId:   openId,
		UserId:   userId,
		ShopId:   shopId,
		OpenType: model.OpenTypeMiniProgram,
	}

	err = database.GetDB(constant.DatabaseConfigKey).Save(userShopRelated).Error
	return
}

func getUser(userId uint64) (user *model.User, err error) {

	user = &model.User{}

	err = database.GetDB(constant.DatabaseConfigKey).Preload("ShopRelated").Where(model.User{UserId: userId}).First(user).Error
	return
}

func getUserByOpenId(openId string) (user *model.User, err error) {

	userShopRelated := &model.UserShop{}
	err = database.GetDB(constant.DatabaseConfigKey).Where(model.UserShop{OpenId: openId}).First(userShopRelated).Error

	if err != nil {
		return
	}

	if userShopRelated.UserId <= 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return getUser(userShopRelated.UserId)
}

func ModifyUserPoint(modifyData *model.ModifyUserPoint) error {
	db := database.GetDB(constant.DatabaseConfigKey)

	// 开启事务
	tx := db.Begin()
	userTx := tx.Model(&model.User{})
	// 判断是加还是减
	expression := ""
	if modifyData.Type == model.UserIncreasePoint {
		// 增
		expression = "point + ?"
	}
	if modifyData.Type == model.UserDecreasePoint {
		// 减
		expression = "point - ?"
	}

	if err := userTx.Where("userId = ?", modifyData.UserId).Update("point", gorm.Expr(expression, modifyData.Point)).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 写入日志表
	pointTx := tx.Model(&model.UserPointHistory{})
	userPointData := &model.UserPointHistory{
		Id:      idcreator.NextID(),
		UserId:  modifyData.UserId,
		Type:    uint8(modifyData.Type),
		Point:   modifyData.Point,
		Remark:  modifyData.Remark,
		OrderId: modifyData.OrderId,
	}

	if err := pointTx.Create(userPointData).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil

}

func convertUserModelToProtobuf(user *model.User, currentShopId uint64) *protos.UserInfo {

	openId := ""
	if len(user.ShopRelated) > 0 {
		for _, relatedRecord := range user.ShopRelated {
			if relatedRecord.ShopId == currentShopId {
				openId = relatedRecord.OpenId
				break
			}
		}
	}

	return &protos.UserInfo{
		UserId:      user.UserId,
		Mobile:      user.Mobile,
		Name:        user.Name,
		Gender:      uint64(user.Gender),
		Birthday:    user.Birthday.Format("2006-01-02 15:04:05"),
		OpenId:      openId,
		CreatedAt:   user.CreatedAt.Format("2006-01-02 15:04:05"),
		Point:       user.Point,
		InviterType: uint32(user.InviterType),
		InviterId:   user.InviterId,
	}
}

func parseStringTime(datetime string) (time.Time, error) {

	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		loc = time.FixedZone("Asia/Shanghai", 8*60*60)
	}

	return time.ParseInLocation("2006-01-02 15:04:05", datetime, loc)
}
