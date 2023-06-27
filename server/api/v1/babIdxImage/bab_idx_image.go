package babIdxImage

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/babIdxImage"
	babIdxImageReq "github.com/flipped-aurora/gin-vue-admin/server/model/babIdxImage/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/uploadImage"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/upload"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type BabIdxImageApi struct {
}

var imgService = service.ServiceGroupApp.BabIdxImageServiceGroup.BabIdxImageService
var privateService = service.ServiceGroupApp.PrivateServiceGroup.PrivateService

// CreateBabIdxImage 创建BabIdxImage
// @Tags BabIdxImage
// @Summary 创建BabIdxImage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body babIdxImage.BabIdxImage true "创建BabIdxImage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /img/createBabIdxImage [post]
func (imgApi *BabIdxImageApi) CreateBabIdxImage(c *gin.Context) {
	var img babIdxImage.BabIdxImage
	err := c.ShouldBind(&img)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"HomePage": {utils.NotEmpty()},
		"Enquiry":  {utils.NotEmpty()},
		"News":     {utils.NotEmpty()},
	}

	if err := utils.Verify(img, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 图片处理
	for i := range img.HomePage {
		url, name, err := privateService.Upload(img.HomePage[i])
		if err != nil {
			response.FailWithDetailed(err, "文件"+name+"上传失败！", c)
			break
		}
		img.Hp = append(img.Hp, uploadImage.Image{name, url})
	}

	for i := range img.Enquiry {
		url, name, err := privateService.Upload(img.Enquiry[i])
		if err != nil {
			response.FailWithDetailed(err, "文件"+name+"上传失败！", c)
			break
		}
		img.Enq = append(img.Enq, uploadImage.Image{name, url})
	}

	for i := range img.News {
		url, name, err := privateService.Upload(img.News[i])
		if err != nil {
			response.FailWithDetailed(err, "文件"+name+"上传失败！", c)
			break
		}
		img.Ne = append(img.Ne, uploadImage.Image{name, url})
	}
	if err := imgService.CreateBabIdxImage(&img); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteBabIdxImage 删除BabIdxImage
// @Tags BabIdxImage
// @Summary 删除BabIdxImage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body babIdxImage.BabIdxImage true "删除BabIdxImage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /img/deleteBabIdxImage [delete]
func (imgApi *BabIdxImageApi) DeleteBabIdxImage(c *gin.Context) {
	var img babIdxImage.BabIdxImage
	err := c.ShouldBindJSON(&img)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := imgService.DeleteBabIdxImage(img); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteBabIdxImageByIds 批量删除BabIdxImage
// @Tags BabIdxImage
// @Summary 批量删除BabIdxImage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除BabIdxImage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /img/deleteBabIdxImageByIds [delete]
func (imgApi *BabIdxImageApi) DeleteBabIdxImageByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := imgService.DeleteBabIdxImageByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateBabIdxImage 更新BabIdxImage
// @Tags BabIdxImage
// @Summary 更新BabIdxImage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body babIdxImage.BabIdxImage true "更新BabIdxImage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /img/updateBabIdxImage [put]
func (imgApi *BabIdxImageApi) UpdateBabIdxImage(c *gin.Context) {
	var img babIdxImage.BabIdxImage
	err := c.ShouldBind(&img)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"HomePage": {utils.NotEmpty()},
		"Enquiry":  {utils.NotEmpty()},
		"News":     {utils.NotEmpty()},
	}
	if err := utils.Verify(img, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 图片处理
	oldImg, err := imgService.GetBabIdxImage(img.ID)
	img.Hp = upload.ProcessUpdateImages(oldImg.Hp, img.HomePage)
	img.Enq = upload.ProcessUpdateImages(oldImg.Enq, img.Enquiry)
	img.Ne = upload.ProcessUpdateImages(oldImg.Ne, img.News)

	if err := imgService.UpdateBabIdxImage(img); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindBabIdxImage 用id查询BabIdxImage
// @Tags BabIdxImage
// @Summary 用id查询BabIdxImage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query babIdxImage.BabIdxImage true "用id查询BabIdxImage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /img/findBabIdxImage [get]
func (imgApi *BabIdxImageApi) FindBabIdxImage(c *gin.Context) {
	var img babIdxImage.BabIdxImage
	err := c.ShouldBindQuery(&img)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reimg, err := imgService.GetBabIdxImage(img.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reimg": reimg}, c)
	}
}

// GetBabIdxImageList 分页获取BabIdxImage列表
// @Tags BabIdxImage
// @Summary 分页获取BabIdxImage列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query babIdxImageReq.BabIdxImageSearch true "分页获取BabIdxImage列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /img/getBabIdxImageList [get]
func (imgApi *BabIdxImageApi) GetBabIdxImageList(c *gin.Context) {
	var pageInfo babIdxImageReq.BabIdxImageSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := imgService.GetBabIdxImageInfoList(pageInfo); err != nil {
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
