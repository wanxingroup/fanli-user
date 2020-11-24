package address

import (
	"time"

	"dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/data/database"
	"dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/data/database/models"
	idcreator "dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/utils/idcreator/snowflake"
	"github.com/jinzhu/gorm"

	"dev-gitlab.wanxingrowth.com/fanli/user/pkg/constant"
	"dev-gitlab.wanxingrowth.com/fanli/user/pkg/model"
)

// 创建地址
func Create(address *model.UserAddress) error {

	address.Id = idcreator.NextID()
	if err := database.GetDB(constant.DatabaseConfigKey).Create(address).Error; err != nil {
		return err
	}
	return nil
}

// 更新地址
func Update(address *model.UserAddress) error {

	updateData := &model.UserAddress{
		ReceiverName: address.ReceiverName,
		Gender:       address.Gender,
		Country:      address.Country,
		Province:     address.Province,
		City:         address.City,
		District:     address.District,
		Address:      address.Address,
		Tag:          address.Tag,
		Mobile:       address.Mobile,
		Type:         address.Type,
		Time: databases.Time{
			BasicTimeFields: databases.BasicTimeFields{
				UpdatedAt: time.Time{},
			},
		},
	}

	if err := database.GetDB(constant.DatabaseConfigKey).Where("id = ? AND userId = ?", address.Id, address.UserId).Update(updateData).Error; err != nil {
		return err
	}
	return nil
}

// 地址列表
func GetList(conditions map[string]interface{}, pageData map[string]uint64) ([]*model.UserAddress, uint64, error) {
	db := database.GetDB(constant.DatabaseConfigKey).Model(&model.UserAddress{})

	if userId, has := conditions["userId"]; has {
		db = db.Where("userId = ?", userId)
	} else {
		return nil, 0, nil
	}

	if lastId, has := pageData["lastId"]; has {
		db = db.Where("id < ?", lastId)
	}

	db = db.Order("id desc")

	results := make([]*model.UserAddress, 0, pageData["pageSize"])

	count, err := constant.FindPageWithLastId(db, pageData, &results)

	return results, count, err

}

func Get(conditions map[string]interface{}) (*model.UserAddress, error) {
	address := new(model.UserAddress)
	db := database.GetDB(constant.DatabaseConfigKey)

	if userId, has := conditions["userId"]; has {
		db = db.Where("userId = ?", userId)
	} else {
		return nil, nil
	}

	if addressId, has := conditions["addressId"]; has {
		db = db.Where("id = ?", addressId)
	}

	err := db.First(address).Error

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}

	return address, nil

}

// 软删除 地址
func Delete(userId, addressId uint64) error {

	if addressId == 0 {
		return nil
	}

	if err := database.GetDB(constant.DatabaseConfigKey).Where("id = ? AND userId = ?", addressId, userId).Delete(&model.UserAddress{}).Error; err != nil {
		return err
	}

	return nil

}
