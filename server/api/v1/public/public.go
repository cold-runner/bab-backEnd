package public

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/babEnquiry"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type PublicApi struct {
}

var publicService = service.ServiceGroupApp.PublicServiceGroup.PublicService

func (p *PublicApi) HomeInfo(c *gin.Context) {
	homePageInfo, err := publicService.HomeInfo()
	if err != nil {
		global.GVA_LOG.Error("获取主页信息失败!", zap.Error(err))
		response.FailWithMessage("请求失败", c)
	} else {
		response.OkWithData(homePageInfo, c)
	}
}

func (p *PublicApi) ApartmentsByCity(c *gin.Context) {
	city := c.Query("city")
	apartmentByCity, err := publicService.ApartmentsByCity(city)
	if err != nil {
		response.FailWithMessage("根据城市查询公寓信息失败", c)
	} else {
		response.OkWithData(apartmentByCity, c)
	}
}

func (p *PublicApi) HousesByApartment(c *gin.Context) {
	apartment := c.Query("apartment")
	houses, err := publicService.HousesByApartment(apartment)
	if err != nil {
		response.FailWithMessage("根据公寓获取房间列表失败", c)
	} else {
		response.OkWithData(houses, c)
	}
}

func (p *PublicApi) News(c *gin.Context) {
	news, err := publicService.News()
	if err != nil {
		response.FailWithMessage("查询咨询信息失败！", c)
	} else {
		response.OkWithData(news, c)
	}
}

func (p *PublicApi) SingleHouse(c *gin.Context) {
	id := c.Query("houseId")
	house, err := publicService.SingleHouse(id)
	if err != nil {
		response.FailWithMessage("查询单个房屋信息", c)
	} else {
		response.OkWithData(gin.H{
			"id":           house.ID,
			"name":         house.Name,
			"chineseName":  house.ChineseName,
			"city":         house.City,
			"company":      house.Company,
			"apartment":    house.Apartment,
			"type":         house.Type,
			"price":        house.Price,
			"area":         house.Area,
			"introduction": house.Introduction,
			"facility":     house.Facility,
			"leasePeriod":  house.LeasePeriod,
			"image":        house.Images,
		}, c)
	}
}

func (p *PublicApi) AboutUs(c *gin.Context) {
	abus, err := publicService.AboutUs()
	if err != nil {
		response.FailWithMessage("关于我们查询失败", c)
	} else {
		response.OkWithData(gin.H{
			"message": abus.Message,
			"image":   abus.Images,
		}, c)
	}
}

func (p *PublicApi) HomePageImgUrl(c *gin.Context) {
	url, err := publicService.HomePageUrl()
	if err != nil {
		response.FailWithMessage("查询主页图片失败！", c)
	} else {
		response.OkWithData(url, c)
	}
}

func (p *PublicApi) EnquiryImageUrl(c *gin.Context) {
	url, err := publicService.EnquiryImageUrl()
	if err != nil {
		response.FailWithMessage("查询主页图片失败！", c)
	} else {
		response.OkWithData(url, c)
	}
}

func (p *PublicApi) NewsUrl(c *gin.Context) {
	url, err := publicService.NewsUrl()
	if err != nil {
		response.FailWithMessage("查询主页图片失败！", c)
	} else {
		response.OkWithData(url, c)
	}
}

func (p *PublicApi) UserEnquiry(c *gin.Context) {
	var eq babEnquiry.BabEnquiry
	err := c.ShouldBindJSON(&eq)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	eq.CreatedBy = utils.GetUserID(c)
	verify := utils.Rules{
		"firstName": {utils.NotEmpty()},
		"lastName":  {utils.NotEmpty()},
		"city":      {utils.NotEmpty()},
		"phone":     {utils.NotEmpty()},
		"email":     {utils.NotEmpty()},
	}
	if err := utils.Verify(eq, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := publicService.CreateUserEnquiry(&eq); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}
