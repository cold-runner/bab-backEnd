import service from '@/utils/requestMutiForm'

// @Tags BabIdxImage
// @Summary 创建BabIdxImage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BabIdxImage true "创建BabIdxImage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /img/createBabIdxImage [post]
export const createBabIdxImage = (data) => {
  return service({
    url: '/img/createBabIdxImage',
    method: 'post',
    data
  })
}

// @Tags BabIdxImage
// @Summary 删除BabIdxImage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BabIdxImage true "删除BabIdxImage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /img/deleteBabIdxImage [delete]
export const deleteBabIdxImage = (data) => {
  return service({
    url: '/img/deleteBabIdxImage',
    method: 'delete',
    data
  })
}

// @Tags BabIdxImage
// @Summary 删除BabIdxImage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除BabIdxImage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /img/deleteBabIdxImage [delete]
export const deleteBabIdxImageByIds = (data) => {
  return service({
    url: '/img/deleteBabIdxImageByIds',
    method: 'delete',
    data
  })
}

// @Tags BabIdxImage
// @Summary 更新BabIdxImage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BabIdxImage true "更新BabIdxImage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /img/updateBabIdxImage [put]
export const updateBabIdxImage = (data) => {
  return service({
    url: '/img/updateBabIdxImage',
    method: 'put',
    data
  })
}

// @Tags BabIdxImage
// @Summary 用id查询BabIdxImage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.BabIdxImage true "用id查询BabIdxImage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /img/findBabIdxImage [get]
export const findBabIdxImage = (params) => {
  return service({
    url: '/img/findBabIdxImage',
    method: 'get',
    params
  })
}

// @Tags BabIdxImage
// @Summary 分页获取BabIdxImage列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取BabIdxImage列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /img/getBabIdxImageList [get]
export const getBabIdxImageList = (params) => {
  return service({
    url: '/img/getBabIdxImageList',
    method: 'get',
    params
  })
}
