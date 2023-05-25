package babIdxImage

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type BabIdxImageRouter struct {
}

// InitBabIdxImageRouter 初始化 BabIdxImage 路由信息
func (s *BabIdxImageRouter) InitBabIdxImageRouter(Router *gin.RouterGroup) {
	imgRouter := Router.Group("img")
	imgRouterWithoutRecord := Router.Group("img")
	var imgApi = v1.ApiGroupApp.BabIdxImageApiGroup.BabIdxImageApi
	{
		imgRouter.POST("createBabIdxImage", imgApi.CreateBabIdxImage)             // 新建BabIdxImage
		imgRouter.DELETE("deleteBabIdxImage", imgApi.DeleteBabIdxImage)           // 删除BabIdxImage
		imgRouter.DELETE("deleteBabIdxImageByIds", imgApi.DeleteBabIdxImageByIds) // 批量删除BabIdxImage
		imgRouter.PUT("updateBabIdxImage", imgApi.UpdateBabIdxImage)              // 更新BabIdxImage
	}
	{
		imgRouterWithoutRecord.GET("findBabIdxImage", imgApi.FindBabIdxImage)       // 根据ID获取BabIdxImage
		imgRouterWithoutRecord.GET("getBabIdxImageList", imgApi.GetBabIdxImageList) // 获取BabIdxImage列表
	}
}
