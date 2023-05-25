package babNews

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/babNews"
	bab_newsReq "github.com/flipped-aurora/gin-vue-admin/server/model/babNews/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/uploadImage"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/upload"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type BabNewsApi struct {
}

var newsInfoService = service.ServiceGroupApp.BabNewsServiceGroup.BabNewsService
var privateService = service.ServiceGroupApp.PrivateServiceGroup.PrivateService

// CreateBabNews 创建BabNews
// @Tags BabNews
// @Summary 创建BabNews
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body babNews.BabNews true "创建BabNews"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /newsInfo/createBabNews [post]
func (newsInfoApi *BabNewsApi) CreateBabNews(c *gin.Context) {
	var newsInfo babNews.BabNews
	err := c.ShouldBind(&newsInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	newsInfo.CreatedBy = utils.GetUserID(c)
	verify := utils.Rules{
		"Title":   {utils.NotEmpty()},
		"Content": {utils.NotEmpty()},
	}
	if err := utils.Verify(newsInfo, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 图片处理
	for i := range newsInfo.FormImages {
		url, name, err := privateService.Upload(newsInfo.FormImages[i])
		if err != nil {
			response.FailWithDetailed(err, "文件"+name+"上传失败！", c)
			break
		}
		newsInfo.Images = append(newsInfo.Images, uploadImage.Image{name, url})
	}

	if err := newsInfoService.CreateBabNews(&newsInfo); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteBabNews 删除BabNews
// @Tags BabNews
// @Summary 删除BabNews
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body babNews.BabNews true "删除BabNews"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /newsInfo/deleteBabNews [delete]
func (newsInfoApi *BabNewsApi) DeleteBabNews(c *gin.Context) {
	var newsInfo babNews.BabNews
	err := c.ShouldBindJSON(&newsInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	newsInfo.DeletedBy = utils.GetUserID(c)
	if err := newsInfoService.DeleteBabNews(newsInfo); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteBabNewsByIds 批量删除BabNews
// @Tags BabNews
// @Summary 批量删除BabNews
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除BabNews"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /newsInfo/deleteBabNewsByIds [delete]
func (newsInfoApi *BabNewsApi) DeleteBabNewsByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	deletedBy := utils.GetUserID(c)
	if err := newsInfoService.DeleteBabNewsByIds(IDS, deletedBy); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateBabNews 更新BabNews
// @Tags BabNews
// @Summary 更新BabNews
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body babNews.BabNews true "更新BabNews"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /newsInfo/updateBabNews [put]
func (newsInfoApi *BabNewsApi) UpdateBabNews(c *gin.Context) {
	var newsInfo babNews.BabNews
	err := c.ShouldBind(&newsInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	newsInfo.UpdatedBy = utils.GetUserID(c)
	verify := utils.Rules{
		"Title":   {utils.NotEmpty()},
		"Content": {utils.NotEmpty()},
		"Image":   {utils.NotEmpty()},
	}
	if err := utils.Verify(newsInfo, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 图片处理
	oldNewsInfo, err := newsInfoService.GetBabNews(newsInfo.ID)
	newsInfo.Images = upload.ProcessUpdateImages(oldNewsInfo.Images, newsInfo.FormImages)

	if err := newsInfoService.UpdateBabNews(newsInfo); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindBabNews 用id查询BabNews
// @Tags BabNews
// @Summary 用id查询BabNews
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query babNews.BabNews true "用id查询BabNews"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /newsInfo/findBabNews [get]
func (newsInfoApi *BabNewsApi) FindBabNews(c *gin.Context) {
	var newsInfo babNews.BabNews
	err := c.ShouldBindQuery(&newsInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if renewsInfo, err := newsInfoService.GetBabNews(newsInfo.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"renewsInfo": renewsInfo}, c)
	}
}

// GetBabNewsList 分页获取BabNews列表
// @Tags BabNews
// @Summary 分页获取BabNews列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bab_newsReq.BabNewsSearch true "分页获取BabNews列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /newsInfo/getBabNewsList [get]
func (newsInfoApi *BabNewsApi) GetBabNewsList(c *gin.Context) {
	var pageInfo bab_newsReq.BabNewsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := newsInfoService.GetBabNewsInfoList(pageInfo); err != nil {
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
