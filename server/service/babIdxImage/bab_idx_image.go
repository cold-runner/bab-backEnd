package babIdxImage

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/babIdxImage"
	babIdxImageReq "github.com/flipped-aurora/gin-vue-admin/server/model/babIdxImage/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type BabIdxImageService struct {
}

// CreateBabIdxImage 创建BabIdxImage记录
// Author [piexlmax](https://github.com/piexlmax)
func (imgService *BabIdxImageService) CreateBabIdxImage(img *babIdxImage.BabIdxImage) (err error) {
	err = global.GVA_DB.Create(img).Error
	return err
}

// DeleteBabIdxImage 删除BabIdxImage记录
// Author [piexlmax](https://github.com/piexlmax)
func (imgService *BabIdxImageService) DeleteBabIdxImage(img babIdxImage.BabIdxImage) (err error) {
	err = global.GVA_DB.Delete(&img).Error
	return err
}

// DeleteBabIdxImageByIds 批量删除BabIdxImage记录
// Author [piexlmax](https://github.com/piexlmax)
func (imgService *BabIdxImageService) DeleteBabIdxImageByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]babIdxImage.BabIdxImage{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateBabIdxImage 更新BabIdxImage记录
// Author [piexlmax](https://github.com/piexlmax)
func (imgService *BabIdxImageService) UpdateBabIdxImage(img babIdxImage.BabIdxImage) (err error) {
	err = global.GVA_DB.Save(&img).Error
	return err
}

// GetBabIdxImage 根据id获取BabIdxImage记录
// Author [piexlmax](https://github.com/piexlmax)
func (imgService *BabIdxImageService) GetBabIdxImage(id uint) (img babIdxImage.BabIdxImage, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&img).Error
	return
}

// GetBabIdxImageInfoList 分页获取BabIdxImage记录
// Author [piexlmax](https://github.com/piexlmax)
func (imgService *BabIdxImageService) GetBabIdxImageInfoList(info babIdxImageReq.BabIdxImageSearch) (list []babIdxImage.BabIdxImage, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&babIdxImage.BabIdxImage{})
	var imgs []babIdxImage.BabIdxImage
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&imgs).Error
	return imgs, total, err
}
