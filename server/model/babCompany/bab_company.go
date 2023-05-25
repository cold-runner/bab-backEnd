// 自动生成模板BabCompany
package babCompany

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// BabCompany 结构体
type BabCompany struct {
	global.GVA_MODEL
	Name         string `json:"name" form:"name" gorm:"column:name;comment:公司名称;"`
	Introduction string `json:"introduction" form:"introduction" gorm:"type:text;column:introduction;comment:;"`
	CreatedBy    uint   `gorm:"column:created_by;comment:创建者"`
	UpdatedBy    uint   `gorm:"column:updated_by;comment:更新者"`
	DeletedBy    uint   `gorm:"column:deleted_by;comment:删除者"`
}

// TableName BabCompany 表名
func (BabCompany) TableName() string {
	return "bab_company"
}
