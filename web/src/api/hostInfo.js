import service from '@/utils/request'

// @Tags HostInfo
// @Summary 创建HostInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HostInfo true "创建HostInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /host/createHostInfo [post]
export const createHostInfo = (data) => {
  return service({
    url: '/host/createHostInfo',
    method: 'post',
    data
  })
}

// @Tags HostInfo
// @Summary 删除HostInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HostInfo true "删除HostInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /host/deleteHostInfo [delete]
export const deleteHostInfo = (data) => {
  return service({
    url: '/host/deleteHostInfo',
    method: 'delete',
    data
  })
}

// @Tags HostInfo
// @Summary 删除HostInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HostInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /host/deleteHostInfo [delete]
export const deleteHostInfoByIds = (data) => {
  return service({
    url: '/host/deleteHostInfoByIds',
    method: 'delete',
    data
  })
}

// @Tags HostInfo
// @Summary 更新HostInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HostInfo true "更新HostInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /host/updateHostInfo [put]
export const updateHostInfo = (data) => {
  return service({
    url: '/host/updateHostInfo',
    method: 'put',
    data
  })
}

// @Tags HostInfo
// @Summary 用id查询HostInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HostInfo true "用id查询HostInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /host/findHostInfo [get]
export const findHostInfo = (params) => {
  return service({
    url: '/host/findHostInfo',
    method: 'get',
    params
  })
}

// @Tags HostInfo
// @Summary 分页获取HostInfo列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取HostInfo列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /host/getHostInfoList [get]
export const getHostInfoList = (params) => {
  return service({
    url: '/host/getHostInfoList',
    method: 'get',
    params
  })
}
