import service from '@/utils/requestMutiForm'

export const getCompanyList = () => {
  return service({
    url: '/private/getCompaniesList',
    method: 'get',
  })
}
export const getApartmentList = () => {
  return service({
    url: '/private/getApartmentList',
    method: 'get',
  })
}
export const getCityList = () => {
  return service({
    url: '/private/getCityList',
    method: 'get',
  })
}
// @Tags BabHouse
// @Summary 创建BabHouse
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BabHouse true "创建BabHouse"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /houseInfo/createBabHouse [post]
export const createBabHouse = (data) => {
  return service({
    url: '/houseInfo/createBabHouse',
    method: 'post',
    data
  })
}

// @Tags BabHouse
// @Summary 删除BabHouse
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BabHouse true "删除BabHouse"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /houseInfo/deleteBabHouse [delete]
export const deleteBabHouse = (data) => {
  return service({
    url: '/houseInfo/deleteBabHouse',
    method: 'delete',
    data
  })
}

// @Tags BabHouse
// @Summary 删除BabHouse
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除BabHouse"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /houseInfo/deleteBabHouse [delete]
export const deleteBabHouseByIds = (data) => {
  return service({
    url: '/houseInfo/deleteBabHouseByIds',
    method: 'delete',
    data
  })
}

// @Tags BabHouse
// @Summary 更新BabHouse
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BabHouse true "更新BabHouse"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /houseInfo/updateBabHouse [put]
export const updateBabHouse = (data) => {
  return service({
    url: '/houseInfo/updateBabHouse',
    method: 'put',
    data
  })
}

// @Tags BabHouse
// @Summary 用id查询BabHouse
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.BabHouse true "用id查询BabHouse"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /houseInfo/findBabHouse [get]
export const findBabHouse = (params) => {
  return service({
    url: '/houseInfo/findBabHouse',
    method: 'get',
    params
  })
}

// @Tags BabHouse
// @Summary 分页获取BabHouse列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取BabHouse列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /houseInfo/getBabHouseList [get]
export const getBabHouseList = (params) => {
  return service({
    url: '/houseInfo/getBabHouseList',
    method: 'get',
    params
  })
}
