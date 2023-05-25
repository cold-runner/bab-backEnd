// 自动生成模板BabHouse
package babHouse

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/uploadImage"
	"mime/multipart"
)

// BabHouse 结构体
type BabHouse struct {
	global.GVA_MODEL
	Name         string                  `json:"name" form:"name" gorm:"column:name;comment:房屋英文名;"`
	ChineseName  string                  `json:"chineseName" form:"chineseName" gorm:"column:chinese_name;comment:房屋中文名;"`
	City         string                  `json:"city" form:"city" gorm:"column:city;comment:房屋所在的城市;"`
	Company      string                  `json:"company" form:"company" gorm:"column:company;comment:房屋所属的公司;"`
	Apartment    string                  `json:"apartment" form:"apartment" gorm:"column:apartment;comment:房屋所属的公寓;"`
	Type         string                  `json:"type" form:"type" gorm:"column:type;comment:房屋类型;"`
	Price        *float64                `json:"price" form:"price" gorm:"column:price;comment:;"`
	Area         *int                    `json:"area" form:"area" gorm:"column:area;comment:房屋面积;"`
	Introduction string                  `json:"introduction" form:"introduction" gorm:"type:text;column:introduction;comment:介绍;"`
	Facility     string                  `json:"facility" form:"facility" gorm:"column:facility;comment:设施;"`
	LeasePeriod  string                  `json:"leasePeriod" form:"leasePeriod" gorm:"column:lease_period;comment:租期"`
	Images       uploadImage.Images      `json:"image"  gorm:"type:json;serializer:json;column:image;comment:图片"`
	FormImages   []*multipart.FileHeader `form:"images[]" gorm:"-"`
	CreatedBy    uint                    `gorm:"column:created_by;comment:创建者"`
	UpdatedBy    uint                    `gorm:"column:updated_by;comment:更新者"`
	DeletedBy    uint                    `gorm:"column:deleted_by;comment:删除者"`
}

// TableName BabHouse 表名
func (BabHouse) TableName() string {
	return "bab_house"
}
