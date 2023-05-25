package babCity

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/babCity"
	bab_cityReq "github.com/flipped-aurora/gin-vue-admin/server/model/babCity/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/uploadImage"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/upload"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type BabCityApi struct {
}

var cityInfoService = service.ServiceGroupApp.BabCityServiceGroup.BabCityService
var privateService = service.ServiceGroupApp.PrivateServiceGroup.PrivateService

// CreateBabCity 创建BabCity
// @Tags BabCity
// @Summary 创建BabCity
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body babCity.BabCity true "创建BabCity"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cityInfo/createBabCity [post]
func (cityInfoApi *BabCityApi) CreateBabCity(c *gin.Context) {
	var cityInfo babCity.BabCity
	err := c.ShouldBind(&cityInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	cityInfo.CreatedBy = utils.GetUserID(c)
	verify := utils.Rules{
		"Name": {utils.NotEmpty()},
	}
	if err := utils.Verify(cityInfo, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 图片处理
	for i := range cityInfo.FormImages {
		url, name, err := privateService.Upload(cityInfo.FormImages[i])
		if err != nil {
			response.FailWithDetailed(err, "文件"+name+"上传失败！", c)
			break
		}
		cityInfo.Images = append(cityInfo.Images, uploadImage.Image{name, url})
	}

	if err := cityInfoService.CreateBabCity(&cityInfo); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteBabCity 删除BabCity
// @Tags BabCity
// @Summary 删除BabCity
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body babCity.BabCity true "删除BabCity"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cityInfo/deleteBabCity [delete]
func (cityInfoApi *BabCityApi) DeleteBabCity(c *gin.Context) {
	var cityInfo babCity.BabCity
	err := c.ShouldBindJSON(&cityInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	cityInfo.DeletedBy = utils.GetUserID(c)
	if err := cityInfoService.DeleteBabCity(cityInfo); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteBabCityByIds 批量删除BabCity
// @Tags BabCity
// @Summary 批量删除BabCity
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除BabCity"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /cityInfo/deleteBabCityByIds [delete]
func (cityInfoApi *BabCityApi) DeleteBabCityByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	deletedBy := utils.GetUserID(c)
	if err := cityInfoService.DeleteBabCityByIds(IDS, deletedBy); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateBabCity 更新BabCity
// @Tags BabCity
// @Summary 更新BabCity
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body babCity.BabCity true "更新BabCity"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /cityInfo/updateBabCity [put]
func (cityInfoApi *BabCityApi) UpdateBabCity(c *gin.Context) {
	var cityInfo babCity.BabCity
	err := c.ShouldBind(&cityInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	cityInfo.UpdatedBy = utils.GetUserID(c)
	verify := utils.Rules{
		"Name":  {utils.NotEmpty()},
		"Image": {utils.NotEmpty()},
	}
	if err := utils.Verify(cityInfo, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 图片处理
	oldCityInfo, err := cityInfoService.GetBabCity(cityInfo.ID)
	cityInfo.Images = upload.ProcessUpdateImages(oldCityInfo.Images, cityInfo.FormImages)

	if err := cityInfoService.UpdateBabCity(cityInfo); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindBabCity 用id查询BabCity
// @Tags BabCity
// @Summary 用id查询BabCity
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query babCity.BabCity true "用id查询BabCity"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /cityInfo/findBabCity [get]
func (cityInfoApi *BabCityApi) FindBabCity(c *gin.Context) {
	var cityInfo babCity.BabCity
	err := c.ShouldBindQuery(&cityInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if recityInfo, err := cityInfoService.GetBabCity(cityInfo.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recityInfo": recityInfo}, c)
	}
}

// GetBabCityList 分页获取BabCity列表
// @Tags BabCity
// @Summary 分页获取BabCity列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bab_cityReq.BabCitySearch true "分页获取BabCity列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cityInfo/getBabCityList [get]
func (cityInfoApi *BabCityApi) GetBabCityList(c *gin.Context) {
	var pageInfo bab_cityReq.BabCitySearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := cityInfoService.GetBabCityInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
