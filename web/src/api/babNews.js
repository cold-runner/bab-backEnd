import service from '@/utils/requestMutiForm'

// @Tags BabNews
// @Summary 创建BabNews
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BabNews true "创建BabNews"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /newsInfo/createBabNews [post]
export const createBabNews = (data) => {
  return service({
    url: '/newsInfo/createBabNews',
    method: 'post',
    data
  })
}

// @Tags BabNews
// @Summary 删除BabNews
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BabNews true "删除BabNews"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /newsInfo/deleteBabNews [delete]
export const deleteBabNews = (data) => {
  return service({
    url: '/newsInfo/deleteBabNews',
    method: 'delete',
    data
  })
}

// @Tags BabNews
// @Summary 删除BabNews
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除BabNews"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /newsInfo/deleteBabNews [delete]
export const deleteBabNewsByIds = (data) => {
  return service({
    url: '/newsInfo/deleteBabNewsByIds',
    method: 'delete',
    data
  })
}

// @Tags BabNews
// @Summary 更新BabNews
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BabNews true "更新BabNews"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /newsInfo/updateBabNews [put]
export const updateBabNews = (data) => {
  return service({
    url: '/newsInfo/updateBabNews',
    method: 'put',
    data
  })
}

// @Tags BabNews
// @Summary 用id查询BabNews
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.BabNews true "用id查询BabNews"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /newsInfo/findBabNews [get]
export const findBabNews = (params) => {
  return service({
    url: '/newsInfo/findBabNews',
    method: 'get',
    params
  })
}

// @Tags BabNews
// @Summary 分页获取BabNews列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取BabNews列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /newsInfo/getBabNewsList [get]
export const getBabNewsList = (params) => {
  return service({
    url: '/newsInfo/getBabNewsList',
    method: 'get',
    params
  })
}
