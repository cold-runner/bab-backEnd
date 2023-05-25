package babApartment

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/babApartment"
	bab_apartmentReq "github.com/flipped-aurora/gin-vue-admin/server/model/babApartment/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/babNews"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/upload"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

type BabApartmentService struct {
}

var mysqlErr *mysql.MySQLError

// CreateBabApartment 创建BabApartment记录
// Author [piexlmax](https://github.com/piexlmax)
func (a *BabApartmentService) CreateBabApartment(babInfo *babApartment.BabApartment) (err error) {
	if err = global.GVA_DB.Create(babInfo).Error; err != nil && errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
		oss := upload.NewOss()
		for _, v := range babInfo.Images {
			oss.DeleteFile(v.Name)
		}
	}
	return err
}

// DeleteBabApartment 删除BabApartment记录
// Author [piexlmax](https://github.com/piexlmax)
func (a *BabApartmentService) DeleteBabApartment(babInfo babApartment.BabApartment) (err error) {
	oss := upload.NewOss()
	global.GVA_DB.First(&babInfo, babInfo.ID)
	for _, v := range babInfo.Images {
		oss.DeleteFile(v.Name)
	}
	global.GVA_DB.Model(&babInfo).Update("image", "{}")

	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&babApartment.BabApartment{}).Where("id = ?", babInfo.ID).Update("deleted_by", babInfo.DeletedBy).Error; err != nil {
			return err
		}
		if err = tx.Delete(&babInfo).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteBabApartmentByIds 批量删除BabApartment记录
// Author [piexlmax](https://github.com/piexlmax)
func (a *BabApartmentService) DeleteBabApartmentByIds(ids request.IdsReq, deleted_by uint) (err error) {
	oss := upload.NewOss()
	var tmp []babApartment.BabApartment
	global.GVA_DB.Where("id in ?", ids.Ids).Find(&tmp)
	for i := range tmp {
		for _, img := range tmp[i].Images {
			oss.DeleteFile(img.Name)
		}
	}

	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		tx.Model(&babNews.BabNews{}).Where("id in ?", ids.Ids).Update("image", "{}")
		if err := tx.Model(&babApartment.BabApartment{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", ids.Ids).Delete(&babApartment.BabApartment{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

func (a *BabApartmentService) IsStored(babInfo babApartment.BabApartment) bool {
	tmp := &babApartment.BabApartment{}
	if err := global.GVA_DB.Where("name = ? AND city = ? AND deleted_at IS NULL", babInfo.Name, babInfo.City).First(&tmp).Error; err == nil && tmp.ID != babInfo.ID {
		return true
	}
	return false
}

// UpdateBabApartment 更新BabApartment记录
// Author [piexlmax](https://github.com/piexlmax)
func (a *BabApartmentService) UpdateBabApartment(babInfo babApartment.BabApartment) (err error) {
	err = global.GVA_DB.Save(&babInfo).Error
	return err
}

// GetBabApartment 根据id获取BabApartment记录
// Author [piexlmax](https://github.com/piexlmax)
func (a *BabApartmentService) GetBabApartment(id uint) (babInfo babApartment.BabApartment, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&babInfo).Error
	return
}

// GetBabApartmentInfoList 分页获取BabApartment记录
// Author [piexlmax](https://github.com/piexlmax)
func (a *BabApartmentService) GetBabApartmentInfoList(info bab_apartmentReq.BabApartmentSearch) (list []babApartment.BabApartment, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&babApartment.BabApartment{})
	var babInfos []babApartment.BabApartment
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	if info.City != "" {
		db = db.Where("city = ?", info.City)
	}
	if info.Type != "" {
		db = db.Where("type LIKE ?", "%"+info.Type+"%")
	}
	if info.Company != "" {
		db = db.Where("company LIKE ?", "%"+info.Company+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&babInfos).Error
	return babInfos, total, err
}
