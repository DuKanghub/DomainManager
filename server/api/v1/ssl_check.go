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

// CreateSSLCheck 创建SSLCheck
// @Tags SSLCheck
// @Summary 创建SSLCheck
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SSLCheck true "创建SSLCheck"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sslCheck/createSSLCheck [post]
func CreateSSLCheck(c *gin.Context) {
	var sslCheck model.SSLCheck
	_ = c.ShouldBindJSON(&sslCheck)
	if err := service.CreateSSLCheck(sslCheck); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteSSLCheck 删除SSLCheck
// @Tags SSLCheck
// @Summary 删除SSLCheck
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SSLCheck true "删除SSLCheck"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sslCheck/deleteSSLCheck [delete]
func DeleteSSLCheck(c *gin.Context) {
	var sslCheck model.SSLCheck
	_ = c.ShouldBindJSON(&sslCheck)
	if err := service.DeleteSSLCheck(sslCheck); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteSSLCheckByIds 批量删除SSLCheck
// @Tags SSLCheck
// @Summary 批量删除SSLCheck
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SSLCheck"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /sslCheck/deleteSSLCheckByIds [delete]
func DeleteSSLCheckByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteSSLCheckByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateSSLCheck 更新SSLCheck
// @Tags SSLCheck
// @Summary 更新SSLCheck
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SSLCheck true "更新SSLCheck"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sslCheck/updateSSLCheck [put]
func UpdateSSLCheck(c *gin.Context) {
	var sslCheck model.SSLCheck
	_ = c.ShouldBindJSON(&sslCheck)
	if err := service.UpdateSSLCheck(sslCheck); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindSSLCheck 用id查询SSLCheck
// @Tags SSLCheck
// @Summary 用id查询SSLCheck
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SSLCheck true "用id查询SSLCheck"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sslCheck/findSSLCheck [get]
func FindSSLCheck(c *gin.Context) {
	var sslCheck model.SSLCheck
	_ = c.ShouldBindQuery(&sslCheck)
	if err, resslCheck := service.GetSSLCheck(sslCheck.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"resslCheck": resslCheck}, c)
	}
}

// GetSSLCheckList 分页获取SSLCheck列表
// @Tags SSLCheck
// @Summary 分页获取SSLCheck列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SSLCheckSearch true "分页获取SSLCheck列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sslCheck/getSSLCheckList [get]
func GetSSLCheckList(c *gin.Context) {
	var pageInfo request.SSLCheckSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetSSLCheckInfoList(pageInfo); err != nil {
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
