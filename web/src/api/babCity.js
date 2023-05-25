import service from '@/utils/requestMutiForm'

// @Tags BabCity
// @Summary 创建BabCity
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BabCity true "创建BabCity"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cityInfo/createBabCity [post]
export const createBabCity = (data) => {
  return service({
    url: '/cityInfo/createBabCity',
    method: 'post',
    data
  })
}

// @Tags BabCity
// @Summary 删除BabCity
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BabCity true "删除BabCity"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cityInfo/deleteBabCity [delete]
export const deleteBabCity = (data) => {
  return service({
    url: '/cityInfo/deleteBabCity',
    method: 'delete',
    data
  })
}

// @Tags BabCity
// @Summary 删除BabCity
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除BabCity"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cityInfo/deleteBabCity [delete]
export const deleteBabCityByIds = (data) => {
  return service({
    url: '/cityInfo/deleteBabCityByIds',
    method: 'delete',
    data
  })
}

// @Tags BabCity
// @Summary 更新BabCity
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BabCity true "更新BabCity"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /cityInfo/updateBabCity [put]
export const updateBabCity = (data) => {
  return service({
    url: '/cityInfo/updateBabCity',
    method: 'put',
    data
  })
}

// @Tags BabCity
// @Summary 用id查询BabCity
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.BabCity true "用id查询BabCity"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /cityInfo/findBabCity [get]
export const findBabCity = (params) => {
  return service({
    url: '/cityInfo/findBabCity',
    method: 'get',
    params
  })
}

// @Tags BabCity
// @Summary 分页获取BabCity列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取BabCity列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cityInfo/getBabCityList [get]
export const getBabCityList = (params) => {
  return service({
    url: '/cityInfo/getBabCityList',
    method: 'get',
    params
  })
}
