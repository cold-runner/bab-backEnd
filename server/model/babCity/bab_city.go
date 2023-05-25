// 自动生成模板BabCity
package babCity

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/uploadImage"
	"mime/multipart"
)

// BabCity 结构体
type BabCity struct {
	global.GVA_MODEL
	Name         string                  `json:"name" form:"name" gorm:"column:name;comment:城市英文名称;"`
	ChineseName  string                  `json:"chineseName" form:"chineseName" gorm:"column:chinese_name;comment:城市中文名称;"`
	Images       uploadImage.Images      `json:"image"  gorm:"type:json;serializer:json;column:image;comment:图片"`
	FormImages   []*multipart.FileHeader `form:"images[]" gorm:"-"`
	Introduction string                  `form:"introduction" gorm:"type:text;column:introduction" json:"introduction"`
	CreatedBy    uint                    `gorm:"column:created_by;comment:创建者"`
	UpdatedBy    uint                    `gorm:"column:updated_by;comment:更新者"`
	DeletedBy    uint                    `gorm:"column:deleted_by;comment:删除者"`
}

// TableName BabCity 表名
func (BabCity) TableName() string {
	return "bab_city"
}
