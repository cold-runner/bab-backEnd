// 自动生成模板BabEnquiry
package babEnquiry

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// BabEnquiry 结构体
type BabEnquiry struct {
	global.GVA_MODEL
	FirstName string `json:"firstName" form:"firstName" gorm:"column:first_name;comment:用户英文名;"`
	LastName  string `json:"lastName" form:"lastName" gorm:"column:last_name;comment:用户姓氏;"`
	City      string `json:"city" form:"city" gorm:"column:city;comment:用户要咨询的城市;"`
	Phone     string `json:"phone" form:"phone" gorm:"column:phone;comment:用户手机号;"`
	WechatNo  string `json:"wechatNo" form:"wechatNo" gorm:"column:wechat_no;comment:用户微信号;"`
	Email     string `json:"email" form:"email" gorm:"column:email;comment:用户邮箱地址;"`
	Message   string `json:"message" form:"message" gorm:"type:text;column:message;comment:用户留言;"`
	CreatedBy uint   `gorm:"column:created_by;comment:创建者"`
	UpdatedBy uint   `gorm:"column:updated_by;comment:更新者"`
	DeletedBy uint   `gorm:"column:deleted_by;comment:删除者"`
}

// TableName BabEnquiry 表名
func (BabEnquiry) TableName() string {
	return "bab_enquiry"
}
