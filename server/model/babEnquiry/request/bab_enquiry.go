package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/babEnquiry"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type BabEnquirySearch struct {
	babEnquiry.BabEnquiry
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}
