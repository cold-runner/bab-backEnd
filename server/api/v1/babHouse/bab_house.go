package babHouse

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/babHouse"
	babHouseReq "github.com/flipped-aurora/gin-vue-admin/server/model/babHouse/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/uploadImage"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/upload"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type BabHouseApi struct {
}

var houseInfoService = service.ServiceGroupApp.BabHouseServiceGroup.BabHouseService
var privateService = service.ServiceGroupApp.PrivateServiceGroup.PrivateService

// CreateBabHouse 创建BabHouse
// @Tags BabHouse
// @Summary 创建BabHouse
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body babHouse.BabHouse true "创建BabHouse"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /houseInfo/createBabHouse [post]
func (houseInfoApi *BabHouseApi) CreateBabHouse(c *gin.Context) {
	var houseInfo babHouse.BabHouse
	err := c.ShouldBind(&houseInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	houseInfo.CreatedBy = utils.GetUserID(c)
	verify := utils.Rules{
		"Name":         {utils.NotEmpty()},
		"ChineseName":  {utils.NotEmpty()},
		"City":         {utils.NotEmpty()},
		"Company":      {utils.NotEmpty()},
		"Apartment":    {utils.NotEmpty()},
		"Type":         {utils.NotEmpty()},
		"Price":        {utils.NotEmpty()},
		"Area":         {utils.NotEmpty()},
		"Introduction": {utils.NotEmpty()},
	}
	if err := utils.Verify(houseInfo, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 图片处理
	for i := range houseInfo.FormImages {
		url, name, err := privateService.Upload(houseInfo.FormImages[i])
		if err != nil {
			response.FailWithDetailed(err, "文件"+name+"上传失败！", c)
			break
		}
		houseInfo.Images = append(houseInfo.Images, uploadImage.Image{name, url})
	}
	if err := houseInfoService.CreateBabHouse(&houseInfo); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithDetailed(err, "创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteBabHouse 删除BabHouse
// @Tags BabHouse
// @Summary 删除BabHouse
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body babHouse.BabHouse true "删除BabHouse"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /houseInfo/deleteBabHouse [delete]
func (houseInfoApi *BabHouseApi) DeleteBabHouse(c *gin.Context) {
	var houseInfo babHouse.BabHouse
	err := c.ShouldBindJSON(&houseInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	houseInfo.DeletedBy = utils.GetUserID(c)
	if err := houseInfoService.DeleteBabHouse(houseInfo); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteBabHouseByIds 批量删除BabHouse
// @Tags BabHouse
// @Summary 批量删除BabHouse
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除BabHouse"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /houseInfo/deleteBabHouseByIds [delete]
func (houseInfoApi *BabHouseApi) DeleteBabHouseByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	deletedBy := utils.GetUserID(c)
	if err := houseInfoService.DeleteBabHouseByIds(IDS, deletedBy); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateBabHouse 更新BabHouse
// @Tags BabHouse
// @Summary 更新BabHouse
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body babHouse.BabHouse true "更新BabHouse"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /houseInfo/updateBabHouse [put]
func (houseInfoApi *BabHouseApi) UpdateBabHouse(c *gin.Context) {
	var houseInfo babHouse.BabHouse
	err := c.ShouldBind(&houseInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	houseInfo.UpdatedBy = utils.GetUserID(c)
	verify := utils.Rules{
		"Name":         {utils.NotEmpty()},
		"ChineseName":  {utils.NotEmpty()},
		"City":         {utils.NotEmpty()},
		"Company":      {utils.NotEmpty()},
		"Apartment":    {utils.NotEmpty()},
		"Type":         {utils.NotEmpty()},
		"Price":        {utils.NotEmpty()},
		"Area":         {utils.NotEmpty()},
		"Introduction": {utils.NotEmpty()},
		"Image":        {utils.NotEmpty()},
	}
	if err := utils.Verify(houseInfo, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 查重
	if houseInfoService.IsStored(houseInfo) {
		response.FailWithMessage("已有存在的记录！房屋名、公寓名、城市名重复！", c)
		return
	}
	// 图片处理
	oldHouseInfo, err := houseInfoService.GetBabHouse(houseInfo.ID)
	houseInfo.Images = upload.ProcessUpdateImages(oldHouseInfo.Images, houseInfo.FormImages)
	if err := houseInfoService.UpdateBabHouse(houseInfo); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindBabHouse 用id查询BabHouse
// @Tags BabHouse
// @Summary 用id查询BabHouse
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query babHouse.BabHouse true "用id查询BabHouse"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /houseInfo/findBabHouse [get]
func (houseInfoApi *BabHouseApi) FindBabHouse(c *gin.Context) {
	var houseInfo babHouse.BabHouse
	err := c.ShouldBindQuery(&houseInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rehouseInfo, err := houseInfoService.GetBabHouse(houseInfo.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rehouseInfo": rehouseInfo}, c)
	}
}

// GetBabHouseList 分页获取BabHouse列表
// @Tags BabHouse
// @Summary 分页获取BabHouse列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query babHouseReq.BabHouseSearch true "分页获取BabHouse列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /houseInfo/getBabHouseList [get]
func (houseInfoApi *BabHouseApi) GetBabHouseList(c *gin.Context) {
	var pageInfo babHouseReq.BabHouseSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := houseInfoService.GetBabHouseInfoList(pageInfo); err != nil {
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
