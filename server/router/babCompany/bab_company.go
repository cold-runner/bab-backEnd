package babCompany

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type BabCompanyRouter struct {
}

// InitBabCompanyRouter 初始化 BabCompany 路由信息
func (s *BabCompanyRouter) InitBabCompanyRouter(Router *gin.RouterGroup) {
	companyInfoRouter := Router.Group("companyInfo")
	companyInfoRouterWithoutRecord := Router.Group("companyInfo")
	var companyInfoApi = v1.ApiGroupApp.BabCompanyApiGroup.BabCompanyApi
	{
		companyInfoRouter.POST("createBabCompany", companyInfoApi.CreateBabCompany)             // 新建BabCompany
		companyInfoRouter.DELETE("deleteBabCompany", companyInfoApi.DeleteBabCompany)           // 删除BabCompany
		companyInfoRouter.DELETE("deleteBabCompanyByIds", companyInfoApi.DeleteBabCompanyByIds) // 批量删除BabCompany
		companyInfoRouter.PUT("updateBabCompany", companyInfoApi.UpdateBabCompany)              // 更新BabCompany
	}
	{
		companyInfoRouterWithoutRecord.GET("findBabCompany", companyInfoApi.FindBabCompany)       // 根据ID获取BabCompany
		companyInfoRouterWithoutRecord.GET("getBabCompanyList", companyInfoApi.GetBabCompanyList) // 获取BabCompany列表
	}
}
