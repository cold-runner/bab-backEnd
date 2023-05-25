package babNews

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type BabNewsRouter struct {
}

// InitBabNewsRouter 初始化 BabNews 路由信息
func (s *BabNewsRouter) InitBabNewsRouter(Router *gin.RouterGroup) {
	newsInfoRouter := Router.Group("newsInfo")
	newsInfoRouterWithoutRecord := Router.Group("newsInfo")
	var newsInfoApi = v1.ApiGroupApp.BabNewsApiGroup.BabNewsApi
	{
		newsInfoRouter.POST("createBabNews", newsInfoApi.CreateBabNews)             // 新建BabNews
		newsInfoRouter.DELETE("deleteBabNews", newsInfoApi.DeleteBabNews)           // 删除BabNews
		newsInfoRouter.DELETE("deleteBabNewsByIds", newsInfoApi.DeleteBabNewsByIds) // 批量删除BabNews
		newsInfoRouter.PUT("updateBabNews", newsInfoApi.UpdateBabNews)              // 更新BabNews
	}
	{
		newsInfoRouterWithoutRecord.GET("findBabNews", newsInfoApi.FindBabNews)       // 根据ID获取BabNews
		newsInfoRouterWithoutRecord.GET("getBabNewsList", newsInfoApi.GetBabNewsList) // 获取BabNews列表
	}
}
