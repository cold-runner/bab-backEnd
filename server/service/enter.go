package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/babAboutUs"
	"github.com/flipped-aurora/gin-vue-admin/server/service/babApartment"
	"github.com/flipped-aurora/gin-vue-admin/server/service/babCity"
	"github.com/flipped-aurora/gin-vue-admin/server/service/babCompany"
	"github.com/flipped-aurora/gin-vue-admin/server/service/babEnquiry"
	"github.com/flipped-aurora/gin-vue-admin/server/service/babHouse"
	"github.com/flipped-aurora/gin-vue-admin/server/service/babIdxImage"
	"github.com/flipped-aurora/gin-vue-admin/server/service/babNews"
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/private"
	"github.com/flipped-aurora/gin-vue-admin/server/service/public"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup       system.ServiceGroup
	ExampleServiceGroup      example.ServiceGroup
	BabCityServiceGroup      babCity.ServiceGroup
	BabCompanyServiceGroup   babCompany.ServiceGroup
	BabApartmentServiceGroup babApartment.ServiceGroup
	BabNewsServiceGroup      babNews.ServiceGroup
	BabEnquiryServiceGroup   babEnquiry.ServiceGroup
	BabHouseServiceGroup     babHouse.ServiceGroup
	PublicServiceGroup       public.ServiceGroup
	PrivateServiceGroup      private.ServiceGroup
	BabIdxImageServiceGroup  babIdxImage.ServiceGroup
	BabAboutUsServiceGroup   babAboutUs.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
