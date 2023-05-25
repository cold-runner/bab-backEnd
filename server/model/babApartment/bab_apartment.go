// 自动生成模板BabApartment
package babApartment

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/uploadImage"
	"mime/multipart"
)

// BabApartment 结构体
type BabApartment struct {
	global.GVA_MODEL
	Name         string                  `json:"name" form:"name" gorm:"column:name;comment:公寓英文名称;"`
	City         string                  `json:"city" form:"city" gorm:"column:city;comment:城市英文名称;"`
	Type         string                  `json:"type" form:"type" gorm:"column:type;comment:户型;"`
	Company      string                  `json:"company" form:"company" gorm:"column:company;comment:公寓所属公司;"`
	Introduction string                  `json:"introduction" form:"introduction" gorm:"type:text;column:introduction;comment:公寓介绍;"`
	Images       uploadImage.Images      `json:"image"  gorm:"type:json;serializer:json;column:image;comment:图片"`
	FormImages   []*multipart.FileHeader `form:"images[]" gorm:"-"`
	CreatedBy    uint                    `gorm:"column:created_by;comment:创建者"`
	UpdatedBy    uint                    `gorm:"column:updated_by;comment:更新者"`
	DeletedBy    uint                    `gorm:"column:deleted_by;comment:删除者"`
}

// TableName BabApartment 表名
func (BabApartment) TableName() string {
	return "bab_apartment"
}
