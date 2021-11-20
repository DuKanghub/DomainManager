import service from '@/utils/request'

// @Tags RecordList
// @Summary 创建Record
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RecordList true "创建Record"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /recordList/createRecord [post]
export const createRecord = (data) => {
  return service({
    url: '/recordList/createRecord',
    method: 'post',
    data
  })
}

// @Tags RecordList
// @Summary 删除Record
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RecordList true "删除Record"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /recordList/deleteRecord [delete]
export const deleteRecord = (data) => {
  return service({
    url: '/recordList/deleteRecord',
    method: 'delete',
    data
  })
}

// @Tags RecordList
// @Summary 批量删除Record
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Record"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /recordList/deleteRecordByIds [delete]
export const deleteRecordByIds = (data) => {
  return service({
    url: '/recordList/deleteRecordByIds',
    method: 'delete',
    data
  })
}

// @Tags RecordList
// @Summary 更新Record
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RecordList true "更新Record"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /recordList/updateRecord [put]
export const updateRecord = (data) => {
  return service({
    url: '/recordList/updateRecord',
    method: 'put',
    data
  })
}

// @Tags RecordList
// @Summary 批量更新Record
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RecordList[] true "批量更新Record"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /recordList/UpdateRecords [put]
export const updateRecords = (data) => {
  return service({
    url: '/recordList/updateRecords',
    method: 'put',
    data
  })
}

// @Tags RecordList
// @Summary 用id查询Record
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RecordList true "用id查询Record"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /recordList/findRecord [get]
export const findRecord = (params) => {
  return service({
    url: '/recordList/findRecord',
    method: 'get',
    params
  })
}

// @Tags RecordList
// @Summary 分页获取RecordList列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取RecordList列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /recordList/getRecordList [get]
export const getRecordList = (params) => {
  return service({
    url: '/recordList/getRecordList',
    method: 'get',
    params
  })
}

// FlushRecords 强制刷新域名解析
// @Tags RecordList
// @Summary 强制刷新域名解析
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /recordList/flushRecords [get]
export const flushRecords = (params) => {
  return service({
    url: '/recordList/flushRecords',
    method: 'get',
    params
  })
}
