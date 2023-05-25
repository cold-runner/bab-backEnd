import service from '@/utils/requestMutiForm'

// @Tags BabAboutUs
// @Summary 创建BabAboutUs
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BabAboutUs true "创建BabAboutUs"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /abus/createBabAboutUs [post]
export const createBabAboutUs = (data) => {
  return service({
    url: '/abus/createBabAboutUs',
    method: 'post',
    data
  })
}

// @Tags BabAboutUs
// @Summary 删除BabAboutUs
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BabAboutUs true "删除BabAboutUs"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /abus/deleteBabAboutUs [delete]
export const deleteBabAboutUs = (data) => {
  return service({
    url: '/abus/deleteBabAboutUs',
    method: 'delete',
    data
  })
}

// @Tags BabAboutUs
// @Summary 删除BabAboutUs
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除BabAboutUs"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /abus/deleteBabAboutUs [delete]
export const deleteBabAboutUsByIds = (data) => {
  return service({
    url: '/abus/deleteBabAboutUsByIds',
    method: 'delete',
    data
  })
}

// @Tags BabAboutUs
// @Summary 更新BabAboutUs
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BabAboutUs true "更新BabAboutUs"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /abus/updateBabAboutUs [put]
export const updateBabAboutUs = (data) => {
  return service({
    url: '/abus/updateBabAboutUs',
    method: 'put',
    data
  })
}

// @Tags BabAboutUs
// @Summary 用id查询BabAboutUs
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.BabAboutUs true "用id查询BabAboutUs"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /abus/findBabAboutUs [get]
export const findBabAboutUs = (params) => {
  return service({
    url: '/abus/findBabAboutUs',
    method: 'get',
    params
  })
}

// @Tags BabAboutUs
// @Summary 分页获取BabAboutUs列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取BabAboutUs列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /abus/getBabAboutUsList [get]
export const getBabAboutUsList = (params) => {
  return service({
    url: '/abus/getBabAboutUsList',
    method: 'get',
    params
  })
}
