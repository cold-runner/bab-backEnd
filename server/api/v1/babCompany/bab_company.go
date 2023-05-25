package babCompany

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/babCompany"
	bab_companyReq "github.com/flipped-aurora/gin-vue-admin/server/model/babCompany/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type BabCompanyApi struct {
}

var companyInfoService = service.ServiceGroupApp.BabCompanyServiceGroup.BabCompanyService

// CreateBabCompany 创建BabCompany
// @Tags BabCompany
// @Summary 创建BabCompany
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body babCompany.BabCompany true "创建BabCompany"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /companyInfo/createBabCompany [post]
func (companyInfoApi *BabCompanyApi) CreateBabCompany(c *gin.Context) {
	var companyInfo babCompany.BabCompany
	err := c.ShouldBindJSON(&companyInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	companyInfo.CreatedBy = utils.GetUserID(c)
	verify := utils.Rules{
		"Name": {utils.NotEmpty()},
	}
	if err := utils.Verify(companyInfo, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := companyInfoService.CreateBabCompany(&companyInfo); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteBabCompany 删除BabCompany
// @Tags BabCompany
// @Summary 删除BabCompany
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body babCompany.BabCompany true "删除BabCompany"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /companyInfo/deleteBabCompany [delete]
func (companyInfoApi *BabCompanyApi) DeleteBabCompany(c *gin.Context) {
	var companyInfo babCompany.BabCompany
	err := c.ShouldBindJSON(&companyInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	companyInfo.DeletedBy = utils.GetUserID(c)
	if err := companyInfoService.DeleteBabCompany(companyInfo); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteBabCompanyByIds 批量删除BabCompany
// @Tags BabCompany
// @Summary 批量删除BabCompany
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除BabCompany"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /companyInfo/deleteBabCompanyByIds [delete]
func (companyInfoApi *BabCompanyApi) DeleteBabCompanyByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	deletedBy := utils.GetUserID(c)
	if err := companyInfoService.DeleteBabCompanyByIds(IDS, deletedBy); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateBabCompany 更新BabCompany
// @Tags BabCompany
// @Summary 更新BabCompany
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body babCompany.BabCompany true "更新BabCompany"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /companyInfo/updateBabCompany [put]
func (companyInfoApi *BabCompanyApi) UpdateBabCompany(c *gin.Context) {
	var companyInfo babCompany.BabCompany
	err := c.ShouldBindJSON(&companyInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	companyInfo.UpdatedBy = utils.GetUserID(c)
	verify := utils.Rules{
		"Name": {utils.NotEmpty()},
	}
	if err := utils.Verify(companyInfo, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := companyInfoService.UpdateBabCompany(companyInfo); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindBabCompany 用id查询BabCompany
// @Tags BabCompany
// @Summary 用id查询BabCompany
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query babCompany.BabCompany true "用id查询BabCompany"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /companyInfo/findBabCompany [get]
func (companyInfoApi *BabCompanyApi) FindBabCompany(c *gin.Context) {
	var companyInfo babCompany.BabCompany
	err := c.ShouldBindQuery(&companyInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if recompanyInfo, err := companyInfoService.GetBabCompany(companyInfo.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recompanyInfo": recompanyInfo}, c)
	}
}

// GetBabCompanyList 分页获取BabCompany列表
// @Tags BabCompany
// @Summary 分页获取BabCompany列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bab_companyReq.BabCompanySearch true "分页获取BabCompany列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /companyInfo/getBabCompanyList [get]
func (companyInfoApi *BabCompanyApi) GetBabCompanyList(c *gin.Context) {
	var pageInfo bab_companyReq.BabCompanySearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := companyInfoService.GetBabCompanyInfoList(pageInfo); err != nil {
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
