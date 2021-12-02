import service from '@/utils/request'

// @Tags SSHUser
// @Summary 创建SSHUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SSHUser true "创建SSHUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sshUser/createSSHUser [post]
export const createSSHUser = (data) => {
  return service({
    url: '/sshUser/createSSHUser',
    method: 'post',
    data
  })
}

// @Tags SSHUser
// @Summary 删除SSHUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SSHUser true "删除SSHUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sshUser/deleteSSHUser [delete]
export const deleteSSHUser = (data) => {
  return service({
    url: '/sshUser/deleteSSHUser',
    method: 'delete',
    data
  })
}

// @Tags SSHUser
// @Summary 删除SSHUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SSHUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sshUser/deleteSSHUser [delete]
export const deleteSSHUserByIds = (data) => {
  return service({
    url: '/sshUser/deleteSSHUserByIds',
    method: 'delete',
    data
  })
}

// @Tags SSHUser
// @Summary 更新SSHUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SSHUser true "更新SSHUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sshUser/updateSSHUser [put]
export const updateSSHUser = (data) => {
  return service({
    url: '/sshUser/updateSSHUser',
    method: 'put',
    data
  })
}

// @Tags SSHUser
// @Summary 用id查询SSHUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SSHUser true "用id查询SSHUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sshUser/findSSHUser [get]
export const findSSHUser = (params) => {
  return service({
    url: '/sshUser/findSSHUser',
    method: 'get',
    params
  })
}

// @Tags SSHUser
// @Summary 分页获取SSHUser列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取SSHUser列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sshUser/getSSHUserList [get]
export const getSSHUserList = (params) => {
  return service({
    url: '/sshUser/getSSHUserList',
    method: 'get',
    params
  })
}
