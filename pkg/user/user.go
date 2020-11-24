package user

import (
	"fmt"

	"dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/data/database"
	validation "github.com/go-ozzo/ozzo-validation/v4"

	"dev-gitlab.wanxingrowth.com/fanli/user/pkg/constant"
	"dev-gitlab.wanxingrowth.com/fanli/user/pkg/model"
)

/**
 * todo 先用 关联查询查询 用户和商铺 关系列表
 */
func GetUserListByCondition(conditions map[string]interface{}, pageNum, pageSize uint64) ([]*model.UserInfo, uint64, error) {
	db := database.GetDB(constant.DatabaseConfigKey).
		Table(model.TableNameUserShop + " AS userShop").
		Joins("LEFT JOIN " + model.TableNameUser + " AS user ON userShop.userId = user.userId").
		Select("user.userId as user_id, user.mobile, user.name, user.createdAt as created_at")

	db = db.Where("shopId = ?", conditions["shopId"])

	if !validation.IsEmpty(conditions["userIds"]) {
		db = db.Where("user.userId in (?)", conditions["userIds"])
	}

	if !validation.IsEmpty(conditions["mobile"]) {
		db = db.Where("user.mobile LIKE ?", fmt.Sprintf("%%%s%%", conditions["mobile"]))
	}

	if !validation.IsEmpty(conditions["name"]) {
		db = db.Where("user.name LIKE ?", fmt.Sprintf("%%%s%%", conditions["name"]))
	}

	if !validation.IsEmpty(conditions["getRegStartTime"]) && !validation.IsEmpty(conditions["getRegEndTime"]) {
		db = db.Where("user.createdAt BETWEEN ? AND ?", conditions["getRegStartTime"], conditions["getRegEndTime"])
	}

	results := make([]*model.UserInfo, 0, pageSize)

	var count uint64

	pageStruct := constant.PageStruct{
		PageSize: pageSize,
		PageNum:  pageNum,
	}

	count, err := constant.FindPage(db, pageStruct, &results)
	return results, count, err
}
