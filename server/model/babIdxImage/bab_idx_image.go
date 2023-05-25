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
	HomePage *multipart.FileHeader `json:"homePage" form:"homePage" gorm:"-"`
	Enquiry  *multipart.FileHeader `json:"enquiry" form:"enquiry" gorm:"-"`
	News     *multipart.FileHeader `json:"news" form:"news" gorm:"-"`

	Hp  uploadImage.Image `gorm:"column:home_page;comment:主页图片;type:json;serializer:json;"`
	Enq uploadImage.Image `gorm:"column:enquiry;comment:询问页展示图;type:json;serializer:json;"`
	Ne  uploadImage.Image `gorm:"column:news;comment:咨询页展示图;type:json;serializer:json;"`
}

// TableName BabIdxImage 表名
func (BabIdxImage) TableName() string {
	return "bab_idx_image"
}
