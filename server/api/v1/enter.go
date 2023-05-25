package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/babAboutUs"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/babApartment"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/babCity"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/babCompany"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/babEnquiry"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/babHouse"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/babIdxImage"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/babNews"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/example"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/private"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/public"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup       system.ApiGroup
	ExampleApiGroup      example.ApiGroup
	BabCityApiGroup      babCity.ApiGroup
	BabCompanyApiGroup   babCompany.ApiGroup
	BabApartmentApiGroup babApartment.ApiGroup
	BabNewsApiGroup      babNews.ApiGroup
	BabEnquiryApiGroup   babEnquiry.ApiGroup
	BabHouseApiGroup     babHouse.ApiGroup
	PublicApiGroup       public.ApiGroup
	PrivateApiGroup      private.ApiGroup
	BabIdxImageApiGroup  babIdxImage.ApiGroup
	BabAboutUsApiGroup   babAboutUs.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
