// 自动生成模板BabAboutUs
package babAboutUs

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/uploadImage"
	"mime/multipart"
)

// BabAboutUs 结构体
type BabAboutUs struct {
	global.GVA_MODEL
	Message    string                  `json:"message" form:"message" gorm:"type:text;column:message;comment:;"`
	Images     uploadImage.Images      `json:"image"  gorm:"type:json;serializer:json;column:image;comment:图片"`
	FormImages []*multipart.FileHeader `form:"images[]" gorm:"-"`
}

// TableName BabAboutUs 表名
func (BabAboutUs) TableName() string {
	return "bab_about_us"
}
