package constant

import (
	"github.com/jinzhu/gorm"
)

const DatabaseConfigKey = "common"

const MaxPageSize = 100

type PageStruct struct {
	PageNum, PageSize uint64
}

func FindPage(db *gorm.DB, pageStruct PageStruct, out interface{}) (uint64, error) {

	if pageStruct.PageNum < 0 || pageStruct.PageSize < 0 {
		return 0, nil
	}

	if pageStruct.PageNum > 0 && pageStruct.PageSize > 0 {
		db = db.Offset((pageStruct.PageNum - 1) * pageStruct.PageSize)
	}

	if pageStruct.PageSize > 0 {
		db = db.Limit(pageStruct.PageSize)
	}

	err := db.Find(out).Error

	if err != nil {
		return 0, err
	}

	var count uint64
	err = db.Count(&count).Error

	return count, err

}

func FindPageWithLastId(db *gorm.DB, pageData map[string]uint64, out interface{}) (uint64, error) {

	// 先查询
	pageSize, has := pageData["pageSize"]

	if has {
		if pageSize > MaxPageSize {
			db = db.Limit(MaxPageSize)
		} else {
			db = db.Limit(pageSize)
		}

	} else {
		db = db.Limit(MaxPageSize)
	}

	// 不传页码为通过最后一条 ID 查询
	if pageNum, has := pageData["pageNum"]; has {
		db = db.Offset((pageNum - 1) * pageSize)
	}

	err := db.Find(out).Error

	if err != nil {
		return 0, err
	}

	var count uint64
	err = db.Count(&count).Error
	return count, err
}
