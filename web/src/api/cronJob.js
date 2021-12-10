import service from '@/utils/request'

// @Tags CronJob
// @Summary 创建CronJob
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CronJob true "创建CronJob"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cronJob/createCronJob [post]
export const createCronJob = (data) => {
  return service({
    url: '/cronJob/createCronJob',
    method: 'post',
    data
  })
}

// @Tags CronJob
// @Summary 删除CronJob
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CronJob true "删除CronJob"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cronJob/deleteCronJob [delete]
export const deleteCronJob = (data) => {
  return service({
    url: '/cronJob/deleteCronJob',
    method: 'delete',
    data
  })
}

// @Tags CronJob
// @Summary 删除CronJob
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除CronJob"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cronJob/deleteCronJob [delete]
export const deleteCronJobByIds = (data) => {
  return service({
    url: '/cronJob/deleteCronJobByIds',
    method: 'delete',
    data
  })
}

// @Tags CronJob
// @Summary 更新CronJob
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CronJob true "更新CronJob"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /cronJob/updateCronJob [put]
export const updateCronJob = (data) => {
  return service({
    url: '/cronJob/updateCronJob',
    method: 'put',
    data
  })
}

// @Tags CronJob
// @Summary 用id查询CronJob
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CronJob true "用id查询CronJob"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /cronJob/findCronJob [get]
export const findCronJob = (params) => {
  return service({
    url: '/cronJob/findCronJob',
    method: 'get',
    params
  })
}

// @Tags CronJob
// @Summary 分页获取CronJob列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取CronJob列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cronJob/getCronJobList [get]
export const getCronJobList = (params) => {
  return service({
    url: '/cronJob/getCronJobList',
    method: 'get',
    params
  })
}

// @Tags CronJob
// @Summary 部署CronJob
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CronJob true "部署CronJob"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cronJob/deployCronJob [post]
export const deployCronJob = (data) => {
  return service({
    url: '/cronJob/deployCronJob',
    method: 'post',
    data
  })
}
