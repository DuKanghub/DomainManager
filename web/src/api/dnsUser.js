import service from '@/utils/request'

// @Tags DnsUser
// @Summary 创建DnsUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DnsUser true "创建DnsUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dnsUser/createDnsUser [post]
export const createDnsUser = (data) => {
  return service({
    url: '/dnsUser/createDnsUser',
    method: 'post',
    data
  })
}

// @Tags DnsUser
// @Summary 删除DnsUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DnsUser true "删除DnsUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dnsUser/deleteDnsUser [delete]
export const deleteDnsUser = (data) => {
  return service({
    url: '/dnsUser/deleteDnsUser',
    method: 'delete',
    data
  })
}

// @Tags DnsUser
// @Summary 批量删除DnsUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除DnsUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /dnsUser/deleteDnsUserByIds [delete]
export const deleteDnsUserByIds = (data) => {
  return service({
    url: '/dnsUser/deleteDnsUserByIds',
    method: 'delete',
    data
  })
}

// @Tags DnsUser
// @Summary 更新DnsUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DnsUser true "更新DnsUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /dnsUser/updateDnsUser [put]
export const updateDnsUser = (data) => {
  return service({
    url: '/dnsUser/updateDnsUser',
    method: 'put',
    data
  })
}

// @Tags DnsUser
// @Summary 用id查询DnsUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DnsUser true "用id查询DnsUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /dnsUser/findDnsUser [get]
export const findDnsUser = (params) => {
  return service({
    url: '/dnsUser/findDnsUser',
    method: 'get',
    params
  })
}

// @Tags DnsUser
// @Summary 分页获取DnsUser列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.DnsUserSearch true "分页获取DnsUser列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dnsUser/getDnsUserList [get]
export const getDnsUserList = (params) => {
  return service({
    url: '/dnsUser/getDnsUserList',
    method: 'get',
    params
  })
}