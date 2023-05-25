package babAboutUs

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/babAboutUs"
	babAboutUsReq "github.com/flipped-aurora/gin-vue-admin/server/model/babAboutUs/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/uploadImage"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/upload"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type BabAboutUsApi struct {
}

var abusService = service.ServiceGroupApp.BabAboutUsServiceGroup.BabAboutUsService
var privateService = service.ServiceGroupApp.PrivateServiceGroup.PrivateService

// CreateBabAboutUs 创建BabAboutUs
// @Tags BabAboutUs
// @Summary 创建BabAboutUs
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body babAboutUs.BabAboutUs true "创建BabAboutUs"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /abus/createBabAboutUs [post]
func (abusApi *BabAboutUsApi) CreateBabAboutUs(c *gin.Context) {
	var abus babAboutUs.BabAboutUs
	err := c.ShouldBind(&abus)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Message": {utils.NotEmpty()},
	}
	if err := utils.Verify(abus, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 图片处理
	for i := range abus.FormImages {
		url, name, err := privateService.Upload(abus.FormImages[i])
		if err != nil {
			response.FailWithDetailed(err, "文件"+name+"上传失败！", c)
			break
		}
		abus.Images = append(abus.Images, uploadImage.Image{name, url})
	}
	if err := abusService.CreateBabAboutUs(&abus); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteBabAboutUs 删除BabAboutUs
// @Tags BabAboutUs
// @Summary 删除BabAboutUs
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body babAboutUs.BabAboutUs true "删除BabAboutUs"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /abus/deleteBabAboutUs [delete]
func (abusApi *BabAboutUsApi) DeleteBabAboutUs(c *gin.Context) {
	var abus babAboutUs.BabAboutUs
	err := c.ShouldBindJSON(&abus)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := abusService.DeleteBabAboutUs(abus); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteBabAboutUsByIds 批量删除BabAboutUs
// @Tags BabAboutUs
// @Summary 批量删除BabAboutUs
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除BabAboutUs"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /abus/deleteBabAboutUsByIds [delete]
func (abusApi *BabAboutUsApi) DeleteBabAboutUsByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := abusService.DeleteBabAboutUsByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateBabAboutUs 更新BabAboutUs
// @Tags BabAboutUs
// @Summary 更新BabAboutUs
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body babAboutUs.BabAboutUs true "更新BabAboutUs"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /abus/updateBabAboutUs [put]
func (abusApi *BabAboutUsApi) UpdateBabAboutUs(c *gin.Context) {
	var abus babAboutUs.BabAboutUs
	err := c.ShouldBind(&abus)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Message": {utils.NotEmpty()},
	}
	if err := utils.Verify(abus, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 图片处理
	oldNewsInfo, err := abusService.GetBabAboutUs(abus.ID)
	abus.Images = upload.ProcessUpdateImages(oldNewsInfo.Images, abus.FormImages)

	if err := abusService.UpdateBabAboutUs(abus); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindBabAboutUs 用id查询BabAboutUs
// @Tags BabAboutUs
// @Summary 用id查询BabAboutUs
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query babAboutUs.BabAboutUs true "用id查询BabAboutUs"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /abus/findBabAboutUs [get]
func (abusApi *BabAboutUsApi) FindBabAboutUs(c *gin.Context) {
	var abus babAboutUs.BabAboutUs
	err := c.ShouldBindQuery(&abus)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reabus, err := abusService.GetBabAboutUs(abus.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reabus": reabus}, c)
	}
}

// GetBabAboutUsList 分页获取BabAboutUs列表
// @Tags BabAboutUs
// @Summary 分页获取BabAboutUs列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query babAboutUsReq.BabAboutUsSearch true "分页获取BabAboutUs列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /abus/getBabAboutUsList [get]
func (abusApi *BabAboutUsApi) GetBabAboutUsList(c *gin.Context) {
	var pageInfo babAboutUsReq.BabAboutUsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := abusService.GetBabAboutUsInfoList(pageInfo); err != nil {
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
