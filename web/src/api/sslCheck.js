import service from '@/utils/request'

// @Tags SSLCheck
// @Summary 创建SSLCheck
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SSLCheck true "创建SSLCheck"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sslCheck/createSSLCheck [post]
export const createSSLCheck = (data) => {
  return service({
    url: '/sslCheck/createSSLCheck',
    method: 'post',
    data
  })
}

// @Tags SSLCheck
// @Summary 删除SSLCheck
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SSLCheck true "删除SSLCheck"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sslCheck/deleteSSLCheck [delete]
export const deleteSSLCheck = (data) => {
  return service({
    url: '/sslCheck/deleteSSLCheck',
    method: 'delete',
    data
  })
}

// @Tags SSLCheck
// @Summary 删除SSLCheck
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SSLCheck"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sslCheck/deleteSSLCheck [delete]
export const deleteSSLCheckByIds = (data) => {
  return service({
    url: '/sslCheck/deleteSSLCheckByIds',
    method: 'delete',
    data
  })
}

// @Tags SSLCheck
// @Summary 更新SSLCheck
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SSLCheck true "更新SSLCheck"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sslCheck/updateSSLCheck [put]
export const updateSSLCheck = (data) => {
  return service({
    url: '/sslCheck/updateSSLCheck',
    method: 'put',
    data
  })
}

// @Tags SSLCheck
// @Summary 用id查询SSLCheck
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SSLCheck true "用id查询SSLCheck"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sslCheck/findSSLCheck [get]
export const findSSLCheck = (params) => {
  return service({
    url: '/sslCheck/findSSLCheck',
    method: 'get',
    params
  })
}

// @Tags SSLCheck
// @Summary 分页获取SSLCheck列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取SSLCheck列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sslCheck/getSSLCheckList [get]
export const getSSLCheckList = (params) => {
  return service({
    url: '/sslCheck/getSSLCheckList',
    method: 'get',
    params
  })
}
