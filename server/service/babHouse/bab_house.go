package babHouse

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/babHouse"
	babHouseReq "github.com/flipped-aurora/gin-vue-admin/server/model/babHouse/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/upload"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

type BabHouseService struct {
}

var mysqlErr *mysql.MySQLError

// CreateBabHouse 创建BabHouse记录
// Author [piexlmax](https://github.com/piexlmax)
func (houseInfoService *BabHouseService) CreateBabHouse(houseInfo *babHouse.BabHouse) (err error) {
	if err = global.GVA_DB.Create(houseInfo).Error; err != nil && errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
		oss := upload.NewOss()
		for _, v := range houseInfo.Images {
			oss.DeleteFile(v.Name)
		}
	}
	return err
}

// DeleteBabHouse 删除BabHouse记录
// Author [piexlmax](https://github.com/piexlmax)
func (houseInfoService *BabHouseService) DeleteBabHouse(houseInfo babHouse.BabHouse) (err error) {
	oss := upload.NewOss()
	global.GVA_DB.First(&houseInfo, houseInfo.ID)
	for _, v := range houseInfo.Images {
		oss.DeleteFile(v.Name)
	}
	global.GVA_DB.Model(&houseInfo).Update("image", "{}")

	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&babHouse.BabHouse{}).Where("id = ?", houseInfo.ID).Update("deleted_by", houseInfo.DeletedBy).Error; err != nil {
			return err
		}
		if err = tx.Delete(&houseInfo).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteBabHouseByIds 批量删除BabHouse记录
// Author [piexlmax](https://github.com/piexlmax)
func (houseInfoService *BabHouseService) DeleteBabHouseByIds(ids request.IdsReq, deleted_by uint) (err error) {
	oss := upload.NewOss()
	var houseInfo []babHouse.BabHouse
	global.GVA_DB.Where("id in ?", ids.Ids).Find(&houseInfo)
	for i := range houseInfo {
		for _, img := range houseInfo[i].Images {
			oss.DeleteFile(img.Name)
		}
	}

	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		tx.Model(&babHouse.BabHouse{}).Where("id in ?", ids.Ids).Update("image", "{}")
		if err := tx.Model(&babHouse.BabHouse{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", ids.Ids).Delete(&babHouse.BabHouse{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

func (houseInfoService *BabHouseService) IsStored(houseInfo babHouse.BabHouse) bool {
	tmp := &babHouse.BabHouse{}
	if err := global.GVA_DB.Where("name = ? AND apartment = ? AND city = ? AND deleted_at IS NULL", houseInfo.Name, houseInfo.Apartment, houseInfo.City).First(&tmp).Error; err == nil && tmp.ID != houseInfo.ID {
		return true
	}
	return false
}

// UpdateBabHouse 更新BabHouse记录
// Author [piexlmax](https://github.com/piexlmax)
func (houseInfoService *BabHouseService) UpdateBabHouse(houseInfo babHouse.BabHouse) (err error) {
	err = global.GVA_DB.Save(&houseInfo).Error
	return err
}

// GetBabHouse 根据id获取BabHouse记录
// Author [piexlmax](https://github.com/piexlmax)
func (houseInfoService *BabHouseService) GetBabHouse(id uint) (houseInfo babHouse.BabHouse, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&houseInfo).Error
	return
}

// GetBabHouseInfoList 分页获取BabHouse记录
// Author [piexlmax](https://github.com/piexlmax)
func (houseInfoService *BabHouseService) GetBabHouseInfoList(info babHouseReq.BabHouseSearch) (list []babHouse.BabHouse, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&babHouse.BabHouse{})
	var houseInfos []babHouse.BabHouse
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	if info.ChineseName != "" {
		db = db.Where("chinese_name LIKE ?", "%"+info.ChineseName+"%")
	}
	if info.City != "" {
		db = db.Where("city = ?", info.City)
	}
	if info.Company != "" {
		db = db.Where("company = ?", info.Company)
	}
	if info.Apartment != "" {
		db = db.Where("apartment = ?", info.Apartment)
	}
	if info.Type != "" {
		db = db.Where("type LIKE ?", "%"+info.Type+"%")
	}
	if info.StartPrice != nil && info.EndPrice != nil {
		db = db.Where("price BETWEEN ? AND ? ", info.StartPrice, info.EndPrice)
	}
	if info.StartArea != nil && info.EndArea != nil {
		db = db.Where("area BETWEEN ? AND ? ", info.StartArea, info.EndArea)
	}
	if info.Facility != "" {
		db = db.Where("facility LIKE ?", "%"+info.Facility+"%")
	}
	if info.LeasePeriod != "" {
		db = db.Where("lease_period LIKE ?", "%"+info.LeasePeriod+"%")
	}
	if info.Introduction != "" {
		db = db.Where("introduction LIKE ?", "%"+info.Introduction+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var OrderStr string
	orderMap := make(map[string]bool)
	orderMap["area"] = true
	if orderMap[info.Sort] {
		OrderStr = info.Sort
		if info.Order == "descending" {
			OrderStr = OrderStr + " desc"
		}
		db = db.Order(OrderStr)
	}

	err = db.Limit(limit).Offset(offset).Find(&houseInfos).Error
	return houseInfos, total, err
}
