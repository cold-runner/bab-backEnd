package babCity

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type BabCityRouter struct {
}

// InitBabCityRouter 初始化 BabCity 路由信息
func (s *BabCityRouter) InitBabCityRouter(Router *gin.RouterGroup) {
	cityInfoRouter := Router.Group("cityInfo")
	cityInfoRouterWithoutRecord := Router.Group("cityInfo")
	var cityInfoApi = v1.ApiGroupApp.BabCityApiGroup.BabCityApi
	{
		cityInfoRouter.POST("createBabCity", cityInfoApi.CreateBabCity)             // 新建BabCity
		cityInfoRouter.DELETE("deleteBabCity", cityInfoApi.DeleteBabCity)           // 删除BabCity
		cityInfoRouter.DELETE("deleteBabCityByIds", cityInfoApi.DeleteBabCityByIds) // 批量删除BabCity
		cityInfoRouter.PUT("updateBabCity", cityInfoApi.UpdateBabCity)              // 更新BabCity
	}
	{
		cityInfoRouterWithoutRecord.GET("findBabCity", cityInfoApi.FindBabCity)       // 根据ID获取BabCity
		cityInfoRouterWithoutRecord.GET("getBabCityList", cityInfoApi.GetBabCityList) // 获取BabCity列表
	}
}
