package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/babHouse"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type BabHouseSearch struct {
	babHouse.BabHouse
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	StartPrice     *float64   `json:"startPrice" form:"startPrice"`
	EndPrice       *float64   `json:"endPrice" form:"endPrice"`
	StartArea      *int       `json:"startArea" form:"startArea"`
	EndArea        *int       `json:"endArea" form:"endArea"`
	request.PageInfo
	Sort  string `json:"sort" form:"sort"`
	Order string `json:"order" form:"order"`
}
