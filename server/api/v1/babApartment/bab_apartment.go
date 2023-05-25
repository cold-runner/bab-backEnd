package babApartment

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/babApartment"
	bab_apartmentReq "github.com/flipped-aurora/gin-vue-admin/server/model/babApartment/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/uploadImage"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/upload"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type BabApartmentApi struct {
}

var babInfoService = service.ServiceGroupApp.BabApartmentServiceGroup.BabApartmentService
var privateService = service.ServiceGroupApp.PrivateServiceGroup.PrivateService

// CreateBabApartment 创建BabApartment
// @Tags BabApartment
// @Summary 创建BabApartment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body babApartment.BabApartment true "创建BabApartment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /babInfo/createBabApartment [post]
func (babInfoApi *BabApartmentApi) CreateBabApartment(c *gin.Context) {
	var babInfo babApartment.BabApartment
	err := c.ShouldBind(&babInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	babInfo.CreatedBy = utils.GetUserID(c)
	verify := utils.Rules{
		"Name":         {utils.NotEmpty()},
		"City":         {utils.NotEmpty()},
		"Type":         {utils.NotEmpty()},
		"Company":      {utils.NotEmpty()},
		"Introduction": {utils.NotEmpty()},
	}
	if err := utils.Verify(babInfo, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 图片处理
	for i := range babInfo.FormImages {
		url, name, err := privateService.Upload(babInfo.FormImages[i])
		if err != nil {
			response.FailWithDetailed(err, "文件"+name+"上传失败！", c)
			break
		}
		babInfo.Images = append(babInfo.Images, uploadImage.Image{name, url})
	}
	if err := babInfoService.CreateBabApartment(&babInfo); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithDetailed(err, "创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteBabApartment 删除BabApartment
// @Tags BabApartment
// @Summary 删除BabApartment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body babApartment.BabApartment true "删除BabApartment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /babInfo/deleteBabApartment [delete]
func (babInfoApi *BabApartmentApi) DeleteBabApartment(c *gin.Context) {
	var babInfo babApartment.BabApartment
	err := c.ShouldBindJSON(&babInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	babInfo.DeletedBy = utils.GetUserID(c)
	if err := babInfoService.DeleteBabApartment(babInfo); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteBabApartmentByIds 批量删除BabApartment
// @Tags BabApartment
// @Summary 批量删除BabApartment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除BabApartment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /babInfo/deleteBabApartmentByIds [delete]
func (babInfoApi *BabApartmentApi) DeleteBabApartmentByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	deletedBy := utils.GetUserID(c)
	if err := babInfoService.DeleteBabApartmentByIds(IDS, deletedBy); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateBabApartment 更新BabApartment
// @Tags BabApartment
// @Summary 更新BabApartment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body babApartment.BabApartment true "更新BabApartment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /babInfo/updateBabApartment [put]
func (babInfoApi *BabApartmentApi) UpdateBabApartment(c *gin.Context) {
	var babInfo babApartment.BabApartment
	err := c.ShouldBind(&babInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	babInfo.UpdatedBy = utils.GetUserID(c)
	verify := utils.Rules{
		"Name":         {utils.NotEmpty()},
		"City":         {utils.NotEmpty()},
		"Type":         {utils.NotEmpty()},
		"Company":      {utils.NotEmpty()},
		"Introduction": {utils.NotEmpty()},
		"Image":        {utils.NotEmpty()},
	}
	if err = utils.Verify(babInfo, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 查重
	if babInfoService.IsStored(babInfo) {
		response.FailWithMessage("已有存在的记录！公寓名、城市名重复！", c)
		return
	}
	// 图片处理
	oldApart, err := babInfoService.GetBabApartment(babInfo.ID)
	babInfo.Images = upload.ProcessUpdateImages(oldApart.Images, babInfo.FormImages)

	if err = babInfoService.UpdateBabApartment(babInfo); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindBabApartment 用id查询BabApartment
// @Tags BabApartment
// @Summary 用id查询BabApartment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query babApartment.BabApartment true "用id查询BabApartment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /babInfo/findBabApartment [get]
func (babInfoApi *BabApartmentApi) FindBabApartment(c *gin.Context) {
	var babInfo babApartment.BabApartment
	err := c.ShouldBindQuery(&babInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rebabInfo, err := babInfoService.GetBabApartment(babInfo.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rebabInfo": rebabInfo}, c)
	}
}

// GetBabApartmentList 分页获取BabApartment列表
// @Tags BabApartment
// @Summary 分页获取BabApartment列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bab_apartmentReq.BabApartmentSearch true "分页获取BabApartment列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /babInfo/getBabApartmentList [get]
func (babInfoApi *BabApartmentApi) GetBabApartmentList(c *gin.Context) {
	var pageInfo bab_apartmentReq.BabApartmentSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := babInfoService.GetBabApartmentInfoList(pageInfo); err != nil {
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
