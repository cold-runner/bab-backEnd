package babEnquiry

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type BabEnquiryRouter struct {
}

// InitBabEnquiryRouter 初始化 BabEnquiry 路由信息
func (s *BabEnquiryRouter) InitBabEnquiryRouter(Router *gin.RouterGroup) {
	enquiryInfoRouter := Router.Group("enquiryInfo")
	enquiryInfoRouterWithoutRecord := Router.Group("enquiryInfo")
	var enquiryInfoApi = v1.ApiGroupApp.BabEnquiryApiGroup.BabEnquiryApi
	{
		enquiryInfoRouter.POST("createBabEnquiry", enquiryInfoApi.CreateBabEnquiry)             // 新建BabEnquiry
		enquiryInfoRouter.DELETE("deleteBabEnquiry", enquiryInfoApi.DeleteBabEnquiry)           // 删除BabEnquiry
		enquiryInfoRouter.DELETE("deleteBabEnquiryByIds", enquiryInfoApi.DeleteBabEnquiryByIds) // 批量删除BabEnquiry
		enquiryInfoRouter.PUT("updateBabEnquiry", enquiryInfoApi.UpdateBabEnquiry)              // 更新BabEnquiry
	}
	{
		enquiryInfoRouterWithoutRecord.GET("findBabEnquiry", enquiryInfoApi.FindBabEnquiry)       // 根据ID获取BabEnquiry
		enquiryInfoRouterWithoutRecord.GET("getBabEnquiryList", enquiryInfoApi.GetBabEnquiryList) // 获取BabEnquiry列表
	}
}
