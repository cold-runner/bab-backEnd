import service from '@/utils/request'

// @Tags BabCompany
// @Summary 创建BabCompany
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BabCompany true "创建BabCompany"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /companyInfo/createBabCompany [post]
export const createBabCompany = (data) => {
  return service({
    url: '/companyInfo/createBabCompany',
    method: 'post',
    data
  })
}

// @Tags BabCompany
// @Summary 删除BabCompany
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BabCompany true "删除BabCompany"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /companyInfo/deleteBabCompany [delete]
export const deleteBabCompany = (data) => {
  return service({
    url: '/companyInfo/deleteBabCompany',
    method: 'delete',
    data
  })
}

// @Tags BabCompany
// @Summary 删除BabCompany
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除BabCompany"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /companyInfo/deleteBabCompany [delete]
export const deleteBabCompanyByIds = (data) => {
  return service({
    url: '/companyInfo/deleteBabCompanyByIds',
    method: 'delete',
    data
  })
}

// @Tags BabCompany
// @Summary 更新BabCompany
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BabCompany true "更新BabCompany"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /companyInfo/updateBabCompany [put]
export const updateBabCompany = (data) => {
  return service({
    url: '/companyInfo/updateBabCompany',
    method: 'put',
    data
  })
}

// @Tags BabCompany
// @Summary 用id查询BabCompany
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.BabCompany true "用id查询BabCompany"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /companyInfo/findBabCompany [get]
export const findBabCompany = (params) => {
  return service({
    url: '/companyInfo/findBabCompany',
    method: 'get',
    params
  })
}

// @Tags BabCompany
// @Summary 分页获取BabCompany列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取BabCompany列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /companyInfo/getBabCompanyList [get]
export const getBabCompanyList = (params) => {
  return service({
    url: '/companyInfo/getBabCompanyList',
    method: 'get',
    params
  })
}
