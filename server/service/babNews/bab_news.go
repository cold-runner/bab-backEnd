package babNews

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/babNews"
	bab_newsReq "github.com/flipped-aurora/gin-vue-admin/server/model/babNews/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/upload"
	"gorm.io/gorm"
)

type BabNewsService struct {
}

// CreateBabNews 创建BabNews记录
// Author [piexlmax](https://github.com/piexlmax)
func (newsInfoService *BabNewsService) CreateBabNews(newsInfo *babNews.BabNews) (err error) {
	err = global.GVA_DB.Create(newsInfo).Error
	return err
}

// DeleteBabNews 删除BabNews记录
// Author [piexlmax](https://github.com/piexlmax)
func (newsInfoService *BabNewsService) DeleteBabNews(newsInfo babNews.BabNews) (err error) {
	oss := upload.NewOss()
	global.GVA_DB.First(&newsInfo, newsInfo.ID)
	for _, v := range newsInfo.Images {
		oss.DeleteFile(v.Name)
	}
	global.GVA_DB.Model(&newsInfo).Update("image", "{}")

	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&babNews.BabNews{}).Where("id = ?", newsInfo.ID).Update("deleted_by", newsInfo.DeletedBy).Error; err != nil {
			return err
		}
		if err = tx.Delete(&newsInfo).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteBabNewsByIds 批量删除BabNews记录
// Author [piexlmax](https://github.com/piexlmax)
func (newsInfoService *BabNewsService) DeleteBabNewsByIds(ids request.IdsReq, deleted_by uint) (err error) {
	oss := upload.NewOss()
	var newsInfo []babNews.BabNews
	global.GVA_DB.Where("id in ?", ids.Ids).Find(&newsInfo)
	for i := range newsInfo {
		for _, img := range newsInfo[i].Images {
			oss.DeleteFile(img.Name)
		}
	}

	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		tx.Model(&babNews.BabNews{}).Where("id in ?", ids.Ids).Update("image", "{}")

		if err := tx.Model(&babNews.BabNews{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", ids.Ids).Delete(&babNews.BabNews{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateBabNews 更新BabNews记录
// Author [piexlmax](https://github.com/piexlmax)
func (newsInfoService *BabNewsService) UpdateBabNews(newsInfo babNews.BabNews) (err error) {
	err = global.GVA_DB.Save(&newsInfo).Error
	return err
}

// GetBabNews 根据id获取BabNews记录
// Author [piexlmax](https://github.com/piexlmax)
func (newsInfoService *BabNewsService) GetBabNews(id uint) (newsInfo babNews.BabNews, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&newsInfo).Error
	return
}

// GetBabNewsInfoList 分页获取BabNews记录
// Author [piexlmax](https://github.com/piexlmax)
func (newsInfoService *BabNewsService) GetBabNewsInfoList(info bab_newsReq.BabNewsSearch) (list []babNews.BabNews, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&babNews.BabNews{})
	var newsInfos []babNews.BabNews
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Title != "" {
		db = db.Where("title LIKE ?", "%"+info.Title+"%")
	}
	if info.Content != "" {
		db = db.Where("content LIKE ?", "%"+info.Content+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&newsInfos).Error
	return newsInfos, total, err
}
