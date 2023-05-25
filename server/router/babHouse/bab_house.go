package babHouse

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type BabHouseRouter struct {
}

// InitBabHouseRouter 初始化 BabHouse 路由信息
func (s *BabHouseRouter) InitBabHouseRouter(Router *gin.RouterGroup) {
	houseInfoRouter := Router.Group("houseInfo")
	houseInfoRouterWithoutRecord := Router.Group("houseInfo")
	var houseInfoApi = v1.ApiGroupApp.BabHouseApiGroup.BabHouseApi
	{
		houseInfoRouter.POST("createBabHouse", houseInfoApi.CreateBabHouse)             // 新建BabHouse
		houseInfoRouter.DELETE("deleteBabHouse", houseInfoApi.DeleteBabHouse)           // 删除BabHouse
		houseInfoRouter.DELETE("deleteBabHouseByIds", houseInfoApi.DeleteBabHouseByIds) // 批量删除BabHouse
		houseInfoRouter.PUT("updateBabHouse", houseInfoApi.UpdateBabHouse)              // 更新BabHouse
	}
	{
		houseInfoRouterWithoutRecord.GET("findBabHouse", houseInfoApi.FindBabHouse)       // 根据ID获取BabHouse
		houseInfoRouterWithoutRecord.GET("getBabHouseList", houseInfoApi.GetBabHouseList) // 获取BabHouse列表
	}
}
