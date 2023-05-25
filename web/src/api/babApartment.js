import service from '@/utils/requestMutiForm'

// @Tags BabApartment
// @Summary 创建BabApartment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BabApartment true "创建BabApartment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /babInfo/createBabApartment [post]
export const createBabApartment = (data) => {
  return service({
    url: '/babInfo/createBabApartment',
    method: 'post',
    data
  })
}

// @Tags BabApartment
// @Summary 删除BabApartment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BabApartment true "删除BabApartment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /babInfo/deleteBabApartment [delete]
export const deleteBabApartment = (data) => {
  return service({
    url: '/babInfo/deleteBabApartment',
    method: 'delete',
    data
  })
}

// @Tags BabApartment
// @Summary 删除BabApartment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除BabApartment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /babInfo/deleteBabApartment [delete]
export const deleteBabApartmentByIds = (data) => {
  return service({
    url: '/babInfo/deleteBabApartmentByIds',
    method: 'delete',
    data
  })
}

// @Tags BabApartment
// @Summary 更新BabApartment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BabApartment true "更新BabApartment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /babInfo/updateBabApartment [put]
export const updateBabApartment = (data) => {
  return service({
    url: '/babInfo/updateBabApartment',
    method: 'put',
    data
  })
}

// @Tags BabApartment
// @Summary 用id查询BabApartment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.BabApartment true "用id查询BabApartment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /babInfo/findBabApartment [get]
export const findBabApartment = (params) => {
  return service({
    url: '/babInfo/findBabApartment',
    method: 'get',
    params
  })
}

// @Tags BabApartment
// @Summary 分页获取BabApartment列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取BabApartment列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /babInfo/getBabApartmentList [get]
export const getBabApartmentList = (params) => {
  return service({
    url: '/babInfo/getBabApartmentList',
    method: 'get',
    params
  })
}
