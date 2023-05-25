package babEnquiry

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/babEnquiry"
	bab_enquiryReq "github.com/flipped-aurora/gin-vue-admin/server/model/babEnquiry/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"gorm.io/gorm"
)

type BabEnquiryService struct {
}

// CreateBabEnquiry 创建BabEnquiry记录
// Author [piexlmax](https://github.com/piexlmax)
func (enquiryInfoService *BabEnquiryService) CreateBabEnquiry(enquiryInfo *babEnquiry.BabEnquiry) (err error) {
	err = global.GVA_DB.Create(enquiryInfo).Error
	return err
}

// DeleteBabEnquiry 删除BabEnquiry记录
// Author [piexlmax](https://github.com/piexlmax)
func (enquiryInfoService *BabEnquiryService) DeleteBabEnquiry(enquiryInfo babEnquiry.BabEnquiry) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&babEnquiry.BabEnquiry{}).Where("id = ?", enquiryInfo.ID).Update("deleted_by", enquiryInfo.DeletedBy).Error; err != nil {
			return err
		}
		if err = tx.Delete(&enquiryInfo).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteBabEnquiryByIds 批量删除BabEnquiry记录
// Author [piexlmax](https://github.com/piexlmax)
func (enquiryInfoService *BabEnquiryService) DeleteBabEnquiryByIds(ids request.IdsReq, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&babEnquiry.BabEnquiry{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", ids.Ids).Delete(&babEnquiry.BabEnquiry{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateBabEnquiry 更新BabEnquiry记录
// Author [piexlmax](https://github.com/piexlmax)
func (enquiryInfoService *BabEnquiryService) UpdateBabEnquiry(enquiryInfo babEnquiry.BabEnquiry) (err error) {
	err = global.GVA_DB.Save(&enquiryInfo).Error
	return err
}

// GetBabEnquiry 根据id获取BabEnquiry记录
// Author [piexlmax](https://github.com/piexlmax)
func (enquiryInfoService *BabEnquiryService) GetBabEnquiry(id uint) (enquiryInfo babEnquiry.BabEnquiry, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&enquiryInfo).Error
	return
}

// GetBabEnquiryInfoList 分页获取BabEnquiry记录
// Author [piexlmax](https://github.com/piexlmax)
func (enquiryInfoService *BabEnquiryService) GetBabEnquiryInfoList(info bab_enquiryReq.BabEnquirySearch) (list []babEnquiry.BabEnquiry, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&babEnquiry.BabEnquiry{})
	var enquiryInfos []babEnquiry.BabEnquiry
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.FirstName != "" {
		db = db.Where("first_name LIKE ?", "%"+info.FirstName+"%")
	}
	if info.LastName != "" {
		db = db.Where("last_name LIKE ?", "%"+info.LastName+"%")
	}
	if info.City != "" {
		db = db.Where("city = ?", info.City)
	}
	if info.Phone != "" {
		db = db.Where("phone LIKE ?", "%"+info.Phone+"%")
	}
	if info.WechatNo != "" {
		db = db.Where("wechat_no LIKE ?", "%"+info.WechatNo+"%")
	}
	if info.Email != "" {
		db = db.Where("email LIKE ?", "%"+info.Email+"%")
	}
	if info.Message != "" {
		db = db.Where("message LIKE ?", "%"+info.Message+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&enquiryInfos).Error
	return enquiryInfos, total, err
}
