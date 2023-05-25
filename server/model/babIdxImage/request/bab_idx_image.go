package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/babIdxImage"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type BabIdxImageSearch struct {
	babIdxImage.BabIdxImage
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}
