// 自动生成模板BabNews
package babNews

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/uploadImage"
	"mime/multipart"
)

// BabNews 结构体
type BabNews struct {
	global.GVA_MODEL
	Title      string                  `json:"title" form:"title" gorm:"column:title;comment:标题;"`
	Content    string                  `json:"content" form:"content" gorm:"type:text;column:content;comment:资讯内容;"`
	Images     uploadImage.Images      `json:"image"  gorm:"type:json;serializer:json;column:image;comment:图片"`
	FormImages []*multipart.FileHeader `form:"images[]" gorm:"-"`
	CreatedBy  uint                    `gorm:"column:created_by;comment:创建者"`
	UpdatedBy  uint                    `gorm:"column:updated_by;comment:更新者"`
	DeletedBy  uint                    `gorm:"column:deleted_by;comment:删除者"`
}

// TableName BabNews 表名
func (BabNews) TableName() string {
	return "bab_news"
}
