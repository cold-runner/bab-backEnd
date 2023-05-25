package babAboutUs

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type BabAboutUsRouter struct {
}

// InitBabAboutUsRouter 初始化 BabAboutUs 路由信息
func (s *BabAboutUsRouter) InitBabAboutUsRouter(Router *gin.RouterGroup) {
	abusRouter := Router.Group("abus")
	abusRouterWithoutRecord := Router.Group("abus")
	var abusApi = v1.ApiGroupApp.BabAboutUsApiGroup.BabAboutUsApi
	{
		abusRouter.POST("createBabAboutUs", abusApi.CreateBabAboutUs)             // 新建BabAboutUs
		abusRouter.DELETE("deleteBabAboutUs", abusApi.DeleteBabAboutUs)           // 删除BabAboutUs
		abusRouter.DELETE("deleteBabAboutUsByIds", abusApi.DeleteBabAboutUsByIds) // 批量删除BabAboutUs
		abusRouter.PUT("updateBabAboutUs", abusApi.UpdateBabAboutUs)              // 更新BabAboutUs
	}
	{
		abusRouterWithoutRecord.GET("findBabAboutUs", abusApi.FindBabAboutUs)       // 根据ID获取BabAboutUs
		abusRouterWithoutRecord.GET("getBabAboutUsList", abusApi.GetBabAboutUsList) // 获取BabAboutUs列表
	}
}
