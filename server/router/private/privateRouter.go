package private

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type PrivateRouter struct {
}

func (p *PrivateRouter) InitPrivateRouter(Router *gin.RouterGroup) {
	privareRouter := Router.Group("private")
	var privateApi = v1.ApiGroupApp.PrivateApiGroup.PrivateApi
	{
		privareRouter.GET("getApartmentList", privateApi.GetApartmentList)
		privareRouter.GET("getCityList", privateApi.GetCityList)
		privareRouter.GET("getCompaniesList", privateApi.GetCompaniesList)

		privareRouter.POST("uploadImage", privateApi.UploadImage)
		privareRouter.DELETE("deleteImage", privateApi.DeleteImage)
	}
}
