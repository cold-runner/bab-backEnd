package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/babAboutUs"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type BabAboutUsSearch struct {
	babAboutUs.BabAboutUs
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}
