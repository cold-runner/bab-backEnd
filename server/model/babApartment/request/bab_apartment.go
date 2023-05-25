package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/babApartment"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type BabApartmentSearch struct {
	babApartment.BabApartment
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}
