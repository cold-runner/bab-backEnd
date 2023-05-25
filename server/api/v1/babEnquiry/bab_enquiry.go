package babEnquiry

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/babEnquiry"
	bab_enquiryReq "github.com/flipped-aurora/gin-vue-admin/server/model/babEnquiry/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type BabEnquiryApi struct {
}

var enquiryInfoService = service.ServiceGroupApp.BabEnquiryServiceGroup.BabEnquiryService

// CreateBabEnquiry 创建BabEnquiry
// @Tags BabEnquiry
// @Summary 创建BabEnquiry
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body babEnquiry.BabEnquiry true "创建BabEnquiry"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /enquiryInfo/createBabEnquiry [post]
func (enquiryInfoApi *BabEnquiryApi) CreateBabEnquiry(c *gin.Context) {
	var enquiryInfo babEnquiry.BabEnquiry
	err := c.ShouldBindJSON(&enquiryInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	enquiryInfo.CreatedBy = utils.GetUserID(c)
	verify := utils.Rules{
		"FirstName": {utils.NotEmpty()},
		"LastName":  {utils.NotEmpty()},
		"City":      {utils.NotEmpty()},
		"Phone":     {utils.NotEmpty()},
		"Email":     {utils.NotEmpty()},
	}
	if err := utils.Verify(enquiryInfo, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := enquiryInfoService.CreateBabEnquiry(&enquiryInfo); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteBabEnquiry 删除BabEnquiry
// @Tags BabEnquiry
// @Summary 删除BabEnquiry
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body babEnquiry.BabEnquiry true "删除BabEnquiry"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /enquiryInfo/deleteBabEnquiry [delete]
func (enquiryInfoApi *BabEnquiryApi) DeleteBabEnquiry(c *gin.Context) {
	var enquiryInfo babEnquiry.BabEnquiry
	err := c.ShouldBindJSON(&enquiryInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	enquiryInfo.DeletedBy = utils.GetUserID(c)
	if err := enquiryInfoService.DeleteBabEnquiry(enquiryInfo); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteBabEnquiryByIds 批量删除BabEnquiry
// @Tags BabEnquiry
// @Summary 批量删除BabEnquiry
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除BabEnquiry"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /enquiryInfo/deleteBabEnquiryByIds [delete]
func (enquiryInfoApi *BabEnquiryApi) DeleteBabEnquiryByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	deletedBy := utils.GetUserID(c)
	if err := enquiryInfoService.DeleteBabEnquiryByIds(IDS, deletedBy); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateBabEnquiry 更新BabEnquiry
// @Tags BabEnquiry
// @Summary 更新BabEnquiry
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body babEnquiry.BabEnquiry true "更新BabEnquiry"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /enquiryInfo/updateBabEnquiry [put]
func (enquiryInfoApi *BabEnquiryApi) UpdateBabEnquiry(c *gin.Context) {
	var enquiryInfo babEnquiry.BabEnquiry
	err := c.ShouldBindJSON(&enquiryInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	enquiryInfo.UpdatedBy = utils.GetUserID(c)
	verify := utils.Rules{
		"FirstName": {utils.NotEmpty()},
		"LastName":  {utils.NotEmpty()},
		"City":      {utils.NotEmpty()},
		"Phone":     {utils.NotEmpty()},
		"Email":     {utils.NotEmpty()},
	}
	if err := utils.Verify(enquiryInfo, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := enquiryInfoService.UpdateBabEnquiry(enquiryInfo); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindBabEnquiry 用id查询BabEnquiry
// @Tags BabEnquiry
// @Summary 用id查询BabEnquiry
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query babEnquiry.BabEnquiry true "用id查询BabEnquiry"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /enquiryInfo/findBabEnquiry [get]
func (enquiryInfoApi *BabEnquiryApi) FindBabEnquiry(c *gin.Context) {
	var enquiryInfo babEnquiry.BabEnquiry
	err := c.ShouldBindQuery(&enquiryInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reenquiryInfo, err := enquiryInfoService.GetBabEnquiry(enquiryInfo.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reenquiryInfo": reenquiryInfo}, c)
	}
}

// GetBabEnquiryList 分页获取BabEnquiry列表
// @Tags BabEnquiry
// @Summary 分页获取BabEnquiry列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bab_enquiryReq.BabEnquirySearch true "分页获取BabEnquiry列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /enquiryInfo/getBabEnquiryList [get]
func (enquiryInfoApi *BabEnquiryApi) GetBabEnquiryList(c *gin.Context) {
	var pageInfo bab_enquiryReq.BabEnquirySearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := enquiryInfoService.GetBabEnquiryInfoList(pageInfo); err != nil {
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
