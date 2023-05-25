package babApartment

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type BabApartmentRouter struct {
}

// InitBabApartmentRouter 初始化 BabApartment 路由信息
func (s *BabApartmentRouter) InitBabApartmentRouter(Router *gin.RouterGroup) {
	babInfoRouter := Router.Group("babInfo")
	babInfoRouterWithoutRecord := Router.Group("babInfo")
	var babInfoApi = v1.ApiGroupApp.BabApartmentApiGroup.BabApartmentApi
	{
		babInfoRouter.POST("createBabApartment", babInfoApi.CreateBabApartment)             // 新建BabApartment
		babInfoRouter.DELETE("deleteBabApartment", babInfoApi.DeleteBabApartment)           // 删除BabApartment
		babInfoRouter.DELETE("deleteBabApartmentByIds", babInfoApi.DeleteBabApartmentByIds) // 批量删除BabApartment
		babInfoRouter.PUT("updateBabApartment", babInfoApi.UpdateBabApartment)              // 更新BabApartment
	}
	{
		babInfoRouterWithoutRecord.GET("findBabApartment", babInfoApi.FindBabApartment)       // 根据ID获取BabApartment
		babInfoRouterWithoutRecord.GET("getBabApartmentList", babInfoApi.GetBabApartmentList) // 获取BabApartment列表
	}
}
