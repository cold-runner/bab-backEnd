package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/babCity"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type BabCitySearch struct {
	babCity.BabCity
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}
