package babCity

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/babCity"
	bab_cityReq "github.com/flipped-aurora/gin-vue-admin/server/model/babCity/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/upload"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

type BabCityService struct {
}

var mysqlErr *mysql.MySQLError

// CreateBabCity 创建BabCity记录
// Author [piexlmax](https://github.com/piexlmax)
func (cityInfoService *BabCityService) CreateBabCity(cityInfo *babCity.BabCity) (err error) {
	if err = global.GVA_DB.Create(cityInfo).Error; err != nil && errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
		oss := upload.NewOss()
		for _, v := range cityInfo.Images {
			oss.DeleteFile(v.Name)
		}
	}
	return err
}

// DeleteBabCity 删除BabCity记录
// Author [piexlmax](https://github.com/piexlmax)
func (cityInfoService *BabCityService) DeleteBabCity(cityInfo babCity.BabCity) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		oss := upload.NewOss()
		global.GVA_DB.First(&cityInfo, cityInfo.ID)
		for _, v := range cityInfo.Images {
			oss.DeleteFile(v.Name)
		}
		global.GVA_DB.Model(&cityInfo).Update("image", "{}")
		if err := tx.Model(&babCity.BabCity{}).Where("id = ?", cityInfo.ID).Update("deleted_by", cityInfo.DeletedBy).Error; err != nil {
			return err
		}
		if err = tx.Delete(&cityInfo).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteBabCityByIds 批量删除BabCity记录
// Author [piexlmax](https://github.com/piexlmax)
func (cityInfoService *BabCityService) DeleteBabCityByIds(ids request.IdsReq, deleted_by uint) (err error) {
	oss := upload.NewOss()
	var city []babCity.BabCity
	global.GVA_DB.Where("id in ?", ids.Ids).Find(&city)
	for i := range city {
		for _, img := range city[i].Images {
			oss.DeleteFile(img.Name)
		}
	}

	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		tx.Model(&babCity.BabCity{}).Where("id in ?", ids.Ids).Update("image", "{}")
		if err := tx.Model(&babCity.BabCity{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", ids.Ids).Delete(&babCity.BabCity{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateBabCity 更新BabCity记录
// Author [piexlmax](https://github.com/piexlmax)
func (cityInfoService *BabCityService) UpdateBabCity(cityInfo babCity.BabCity) (err error) {
	if err = global.GVA_DB.Save(&cityInfo).Error; err != nil && errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
		oss := upload.NewOss()
		for _, v := range cityInfo.Images {
			oss.DeleteFile(v.Name)
		}
	}
	return err
}

// GetBabCity 根据id获取BabCity记录
// Author [piexlmax](https://github.com/piexlmax)
func (cityInfoService *BabCityService) GetBabCity(id uint) (cityInfo babCity.BabCity, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&cityInfo).Error
	return
}

// GetBabCityInfoList 分页获取BabCity记录
// Author [piexlmax](https://github.com/piexlmax)
func (cityInfoService *BabCityService) GetBabCityInfoList(info bab_cityReq.BabCitySearch) (list []babCity.BabCity, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&babCity.BabCity{})
	var cityInfos []babCity.BabCity
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Name != "" {
		db = db.Where("name = ?", info.Name)
	}
	if info.ChineseName != "" {
		db = db.Where("chinese_name like ?", "%"+info.ChineseName+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&cityInfos).Error
	return cityInfos, total, err
}
