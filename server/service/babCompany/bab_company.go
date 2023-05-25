package babCompany

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/babCompany"
	bab_companyReq "github.com/flipped-aurora/gin-vue-admin/server/model/babCompany/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"gorm.io/gorm"
)

type BabCompanyService struct {
}

// CreateBabCompany 创建BabCompany记录
// Author [piexlmax](https://github.com/piexlmax)
func (companyInfoService *BabCompanyService) CreateBabCompany(companyInfo *babCompany.BabCompany) (err error) {
	err = global.GVA_DB.Create(companyInfo).Error
	return err
}

// DeleteBabCompany 删除BabCompany记录
// Author [piexlmax](https://github.com/piexlmax)
func (companyInfoService *BabCompanyService) DeleteBabCompany(companyInfo babCompany.BabCompany) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&babCompany.BabCompany{}).Where("id = ?", companyInfo.ID).Update("deleted_by", companyInfo.DeletedBy).Error; err != nil {
			return err
		}
		if err = tx.Delete(&companyInfo).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteBabCompanyByIds 批量删除BabCompany记录
// Author [piexlmax](https://github.com/piexlmax)
func (companyInfoService *BabCompanyService) DeleteBabCompanyByIds(ids request.IdsReq, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&babCompany.BabCompany{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", ids.Ids).Delete(&babCompany.BabCompany{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateBabCompany 更新BabCompany记录
// Author [piexlmax](https://github.com/piexlmax)
func (companyInfoService *BabCompanyService) UpdateBabCompany(companyInfo babCompany.BabCompany) (err error) {
	err = global.GVA_DB.Save(&companyInfo).Error
	return err
}

// GetBabCompany 根据id获取BabCompany记录
// Author [piexlmax](https://github.com/piexlmax)
func (companyInfoService *BabCompanyService) GetBabCompany(id uint) (companyInfo babCompany.BabCompany, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&companyInfo).Error
	return
}

// GetBabCompanyInfoList 分页获取BabCompany记录
// Author [piexlmax](https://github.com/piexlmax)
func (companyInfoService *BabCompanyService) GetBabCompanyInfoList(info bab_companyReq.BabCompanySearch) (list []babCompany.BabCompany, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&babCompany.BabCompany{})
	var companyInfos []babCompany.BabCompany
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&companyInfos).Error
	return companyInfos, total, err
}
