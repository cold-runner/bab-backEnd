package global

import (
	"time"

	"gorm.io/gorm"
)

type GVA_MODEL struct {
	ID        uint           `form:"ID" gorm:"primarykey"`            // 主键ID
	CreatedAt time.Time      `form:"CreatedAt"`                       // 创建时间
	UpdatedAt time.Time      `form:"UpdatedAt"`                       // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-" form:"DeletedAt"` // 删除时间
}
