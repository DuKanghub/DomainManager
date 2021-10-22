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

// CreateRecord 创建Record
// @Tags Record
// @Summary 创建Record
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Record true "创建Record"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /recordList/createRecord [post]
func CreateRecord(c *gin.Context) {
	var record model.Record
	_ = c.ShouldBindJSON(&record)
	if err := service.CreateRecord(record); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}


// DeleteRecord 删除解析
// @Tags Record
// @Summary 删除Record
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Record true "删除Record"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /recordList/deleteRecord [delete]
func DeleteRecord(c *gin.Context) {
	var record model.Record
	_ = c.ShouldBindJSON(&record)
	if err := service.DeleteRecord(record); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}


// DeleteRecordByIds 批量删除Record
// @Tags Record
// @Summary 批量删除Record
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Record"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /recordList/deleteRecordByIds [delete]
func DeleteRecordByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteRecordByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateRecord 更新Record
// @Tags Record
// @Summary 更新Record
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Record true "更新Record"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /recordList/updateRecord [put]
func UpdateRecord(c *gin.Context) {
	var record model.Record
	_ = c.ShouldBindJSON(&record)
	if err := service.UpdateRecord(record); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// UpdateRecords 批量更新Record
// @Tags Record
// @Summary 批量更新Record
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Records true "批量更新Record"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /recordList/UpdateRecords [put]
func UpdateRecords(c *gin.Context) {
	var records model.Records
	_ = c.ShouldBindJSON(&records)
	if err := service.UpdateRecordMulti(records); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindRecordById 用id查询Record
// @Tags Record
// @Summary 用id查询Record
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Record true "用id查询Record"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /recordList/findRecord [get]
func FindRecordById(c *gin.Context) {
	var rr model.Record
	_ = c.ShouldBindQuery(&rr)
	if err, record := service.GetRecord(rr.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rerecord": record}, c)
	}
}

// GetRecordList 分页获取Record列表
// @Tags Record
// @Summary 分页获取Record列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.RecordListSearch true "分页获取record列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /recordList/getRecordList [get]
func GetRecordList(c *gin.Context) {
	var pageInfo request.RecordListSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetRecordInfoList(pageInfo); err != nil {
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

// FlushRecords 强制刷新域名解析
// @Tags Record
// @Summary 强制刷新域名解析
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /recordList/flushRecords [get]
func FlushRecords(c *gin.Context)  {
	err := service.FlushRecordsToDb()
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithMessage("获取成功", c)
	}
}