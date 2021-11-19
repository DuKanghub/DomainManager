import service from '@/utils/request'

// @Tags DomainInfo
// @Summary 创建Domain
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DomainInfo true "创建Domain"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /domainList/createDomain [post]
export const createDomain = (data) => {
  return service({
    url: '/domainList/createDomain',
    method: 'post',
    data
  })
}

// DeleteDomain 删除Domain
// @Tags DomainInfo
// @Summary 删除Domain
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DomainInfo true "删除Domain"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /domainList/deleteDomain [delete]
export const deleteDomain = (data) => {
  return service({
    url: '/domainList/deleteDomain',
    method: 'delete',
    data
  })
}

// @Tags DomainInfo
// @Summary 批量删除Domains
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Domains"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /domainList/deleteDomainByIds [delete]
export const deleteDomainByIds = (data) => {
  return service({
    url: '/domainList/deleteDomainByIds',
    method: 'delete',
    data
  })
}

// @Tags DomainInfo
// @Summary 更新Domain
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DomainInfo true "更新Domain"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /domainList/updateDomain [put]
export const updateDomain = (data) => {
  return service({
    url: '/domainList/updateDomain',
    method: 'put',
    data
  })
}

// @Tags DomainInfo
// @Summary 批量更新Domain
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body []model.DomainInfo true "批量更新Domain"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量更新成功"}"
// @Router /domainList/updateDomains [put]
export const updateDomains = (data) => {
  return service({
    url: '/domainList/updateDomains',
    method: 'put',
    data
  })
}

// @Tags DomainInfo
// @Summary 用id查询Domain
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DomainInfo true "用id查询Domain"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /domainList/findDomain [get]
export const findDomain = (params) => {
  return service({
    url: '/domainList/findDomain',
    method: 'get',
    params
  })
}

// @Tags DomainInfo
// @Summary 分页获取Domain列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.DomainListSearch true "分页获取Domain列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /domainList/getDomainList [get]
export const getDomainList = (params) => {
  return service({
    url: '/domainList/getDomainList',
    method: 'get',
    params
  })
}

// @Tags DomainInfo
// @Summary 从线上获取最新域名信息并更新到数据库
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /domainList/flushDomains [get]
export const flushDomains = (params) => {
  return service({
    url: '/domainList/flushDomains',
    method: 'get',
    params
  })
}
