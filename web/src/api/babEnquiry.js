import service from '@/utils/request'

// @Tags BabEnquiry
// @Summary 创建BabEnquiry
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BabEnquiry true "创建BabEnquiry"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /enquiryInfo/createBabEnquiry [post]
export const createBabEnquiry = (data) => {
  return service({
    url: '/enquiryInfo/createBabEnquiry',
    method: 'post',
    data
  })
}

// @Tags BabEnquiry
// @Summary 删除BabEnquiry
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BabEnquiry true "删除BabEnquiry"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /enquiryInfo/deleteBabEnquiry [delete]
export const deleteBabEnquiry = (data) => {
  return service({
    url: '/enquiryInfo/deleteBabEnquiry',
    method: 'delete',
    data
  })
}

// @Tags BabEnquiry
// @Summary 删除BabEnquiry
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除BabEnquiry"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /enquiryInfo/deleteBabEnquiry [delete]
export const deleteBabEnquiryByIds = (data) => {
  return service({
    url: '/enquiryInfo/deleteBabEnquiryByIds',
    method: 'delete',
    data
  })
}

// @Tags BabEnquiry
// @Summary 更新BabEnquiry
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BabEnquiry true "更新BabEnquiry"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /enquiryInfo/updateBabEnquiry [put]
export const updateBabEnquiry = (data) => {
  return service({
    url: '/enquiryInfo/updateBabEnquiry',
    method: 'put',
    data
  })
}

// @Tags BabEnquiry
// @Summary 用id查询BabEnquiry
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.BabEnquiry true "用id查询BabEnquiry"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /enquiryInfo/findBabEnquiry [get]
export const findBabEnquiry = (params) => {
  return service({
    url: '/enquiryInfo/findBabEnquiry',
    method: 'get',
    params
  })
}

// @Tags BabEnquiry
// @Summary 分页获取BabEnquiry列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取BabEnquiry列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /enquiryInfo/getBabEnquiryList [get]
export const getBabEnquiryList = (params) => {
  return service({
    url: '/enquiryInfo/getBabEnquiryList',
    method: 'get',
    params
  })
}
