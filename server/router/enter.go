package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/babAboutUs"
	"github.com/flipped-aurora/gin-vue-admin/server/router/babApartment"
	"github.com/flipped-aurora/gin-vue-admin/server/router/babCity"
	"github.com/flipped-aurora/gin-vue-admin/server/router/babCompany"
	"github.com/flipped-aurora/gin-vue-admin/server/router/babEnquiry"
	"github.com/flipped-aurora/gin-vue-admin/server/router/babHouse"
	"github.com/flipped-aurora/gin-vue-admin/server/router/babIdxImage"
	"github.com/flipped-aurora/gin-vue-admin/server/router/babNews"
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/private"
	"github.com/flipped-aurora/gin-vue-admin/server/router/public"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
)

type RouterGroup struct {
	System       system.RouterGroup
	Example      example.RouterGroup
	BabCity      babCity.RouterGroup
	BabCompany   babCompany.RouterGroup
	BabApartment babApartment.RouterGroup
	BabNews      babNews.RouterGroup
	BabEnquiry   babEnquiry.RouterGroup
	BabHouse     babHouse.RouterGroup
	Public       public.RouterGroup
	Private      private.RouterGroup
	BabIdxImage  babIdxImage.RouterGroup
	BabAboutUs   babAboutUs.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
