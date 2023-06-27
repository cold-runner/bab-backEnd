package babAboutUs

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/babAboutUs"
	babAboutUsReq "github.com/flipped-aurora/gin-vue-admin/server/model/babAboutUs/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type BabAboutUsService struct {
}

// CreateBabAboutUs 创建BabAboutUs记录
// Author [piexlmax](https://github.com/piexlmax)
func (abusService *BabAboutUsService) CreateBabAboutUs(abus *babAboutUs.BabAboutUs) (err error) {
	err = global.GVA_DB.Create(abus).Error
	return err
}

// DeleteBabAboutUs 删除BabAboutUs记录
// Author [piexlmax](https://github.com/piexlmax)
func (abusService *BabAboutUsService) DeleteBabAboutUs(abus babAboutUs.BabAboutUs) (err error) {
	err = global.GVA_DB.Unscoped().Delete(&abus).Error
	return err
}

// DeleteBabAboutUsByIds 批量删除BabAboutUs记录
// Author [piexlmax](https://github.com/piexlmax)
func (abusService *BabAboutUsService) DeleteBabAboutUsByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Unscoped().Delete(&[]babAboutUs.BabAboutUs{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateBabAboutUs 更新BabAboutUs记录
// Author [piexlmax](https://github.com/piexlmax)
func (abusService *BabAboutUsService) UpdateBabAboutUs(abus babAboutUs.BabAboutUs) (err error) {
	err = global.GVA_DB.Save(&abus).Error
	return err
}

// GetBabAboutUs 根据id获取BabAboutUs记录
// Author [piexlmax](https://github.com/piexlmax)
func (abusService *BabAboutUsService) GetBabAboutUs(id uint) (abus babAboutUs.BabAboutUs, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&abus).Error
	return
}

// GetBabAboutUsInfoList 分页获取BabAboutUs记录
// Author [piexlmax](https://github.com/piexlmax)
func (abusService *BabAboutUsService) GetBabAboutUsInfoList(info babAboutUsReq.BabAboutUsSearch) (list []babAboutUs.BabAboutUs, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&babAboutUs.BabAboutUs{})
	var abuss []babAboutUs.BabAboutUs
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&abuss).Error
	return abuss, total, err
}
