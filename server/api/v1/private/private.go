package private

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/uploadImage"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
)

type PrivateApi struct {
}

var privateService = service.ServiceGroupApp.PrivateServiceGroup.PrivateService

func (p *PrivateApi) GetApartmentList(c *gin.Context) {
	apartments, _ := privateService.GetApartmentList()
	necessary := make([]string, len(apartments))
	for i := range apartments {
		necessary[i] = apartments[i].Name
	}
	response.OkWithData(necessary, c)
}

func (p *PrivateApi) GetCityList(c *gin.Context) {
	cities, _ := privateService.GetCityList()
	necessary := make([]string, len(cities))
	for i := range cities {
		necessary[i] = cities[i].Name
	}
	response.OkWithData(necessary, c)
}

func (p *PrivateApi) GetCompaniesList(c *gin.Context) {
	companies, _ := privateService.GetCompaniesList()
	necessary := make([]string, len(companies))
	for i := range companies {
		necessary[i] = companies[i].Name
	}
	response.OkWithData(necessary, c)
}

func (p *PrivateApi) UploadImage(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		response.FailWithDetailed(err, "图片上传失败！", c)
	}
	url, fileKey, err := privateService.Upload(file)
	if err != nil {
		response.FailWithDetailed(err, "图片上传失败！", c)
	} else {
		response.OkWithData(gin.H{
			"name": fileKey,
			"url":  url,
		}, c)
	}
}

func (p *PrivateApi) DeleteImage(c *gin.Context) {
	dImage := &uploadImage.Image{}
	if err := c.ShouldBindJSON(dImage); err != nil {
		response.FailWithDetailed(err, "绑定参数失败", c)
		return
	}

	if err := privateService.Delete(dImage.Name); err != nil {
		response.FailWithDetailed(err, "图片删除失败！", c)
	} else {
		response.OkWithMessage("图片删除成功", c)
	}
}
