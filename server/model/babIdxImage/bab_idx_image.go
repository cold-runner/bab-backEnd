// 自动生成模板BabIdxImage
package babIdxImage

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/uploadImage"
	"mime/multipart"
)

// BabIdxImage 结构体
type BabIdxImage struct {
	global.GVA_MODEL
	HomePage []*multipart.FileHeader `form:"homePage[]" gorm:"-"`
	Enquiry  []*multipart.FileHeader `form:"enquiry[]" gorm:"-"`
	News     []*multipart.FileHeader `form:"news[]" gorm:"-"`

	Hp  uploadImage.Images `gorm:"column:home_page;comment:主页图片;type:json;serializer:json;"`
	Enq uploadImage.Images `gorm:"column:enquiry;comment:询问页展示图;type:json;serializer:json;"`
	Ne  uploadImage.Images `gorm:"column:news;comment:咨询页展示图;type:json;serializer:json;"`
}

// TableName BabIdxImage 表名
func (BabIdxImage) TableName() string {
	return "bab_idx_image"
}
