package public

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type PublicRouter struct {
}

// InitPublicRouter 初始化 PublicRouter 路由信息
func (p *PublicRouter) InitPublicRouter(Router *gin.RouterGroup) {
	publicRouter := Router.Group("public")

	var publicApi = v1.ApiGroupApp.PublicApiGroup.PublicApi
	{
		publicRouter.GET("homeInfo", publicApi.HomeInfo)                   // 获取主页信息
		publicRouter.GET("apartmentsByCity", publicApi.ApartmentsByCity)   // 根据城市获取公寓信息
		publicRouter.GET("housesByApartment", publicApi.HousesByApartment) // 获取目标公寓中所有的房屋信息
		publicRouter.GET("singleHouse", publicApi.SingleHouse)             // 根据id获取房屋信息
		publicRouter.GET("aboutUs", publicApi.AboutUs)                     // 获取关于我们页信息
		publicRouter.GET("news", publicApi.News)                           // 获取咨询列表
		publicRouter.GET("homePageImgUrl", publicApi.HomePageImgUrl)       // 获取主页图片
		publicRouter.GET("newsImgUrl", publicApi.NewsUrl)                  // 获取咨询页图片
		publicRouter.GET("enquiryImgUrl", publicApi.EnquiryImageUrl)       // 获取询问页图片
	}
	{
		publicRouter.POST("userEnquiry", publicApi.UserEnquiry) // 用户上交询问表单
	}
}
