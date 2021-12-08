package v1

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"gin-vue-admin/model/response"
	"gin-vue-admin/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// CreateCronJob 创建CronJob
// @Tags CronJob
// @Summary 创建CronJob
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CronJob true "创建CronJob"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cronJob/createCronJob [post]
func CreateCronJob(c *gin.Context) {
	var cronJob model.CronJob
	_ = c.ShouldBindJSON(&cronJob)
	if err := service.CreateCronJob(cronJob); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCronJob 删除CronJob
// @Tags CronJob
// @Summary 删除CronJob
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CronJob true "删除CronJob"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cronJob/deleteCronJob [delete]
func DeleteCronJob(c *gin.Context) {
	var cronJob model.CronJob
	_ = c.ShouldBindJSON(&cronJob)
	if err := service.DeleteCronJob(cronJob); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCronJobByIds 批量删除CronJob
// @Tags CronJob
// @Summary 批量删除CronJob
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除CronJob"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /cronJob/deleteCronJobByIds [delete]
func DeleteCronJobByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteCronJobByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCronJob 更新CronJob
// @Tags CronJob
// @Summary 更新CronJob
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CronJob true "更新CronJob"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /cronJob/updateCronJob [put]
func UpdateCronJob(c *gin.Context) {
	var cronJob model.CronJob
	_ = c.ShouldBindJSON(&cronJob)
	if err := service.UpdateCronJob(cronJob); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCronJob 用id查询CronJob
// @Tags CronJob
// @Summary 用id查询CronJob
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CronJob true "用id查询CronJob"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /cronJob/findCronJob [get]
func FindCronJob(c *gin.Context) {
	var cronJob model.CronJob
	_ = c.ShouldBindQuery(&cronJob)
	if err, recronJob := service.GetCronJob(cronJob.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recronJob": recronJob}, c)
	}
}

// GetCronJobList 分页获取CronJob列表
// @Tags CronJob
// @Summary 分页获取CronJob列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.CronJobSearch true "分页获取CronJob列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cronJob/getCronJobList [get]
func GetCronJobList(c *gin.Context) {
	var pageInfo request.CronJobSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetCronJobInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
