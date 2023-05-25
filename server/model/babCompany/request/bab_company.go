package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/babCompany"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type BabCompanySearch struct {
	babCompany.BabCompany
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}
