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

// CreateSSHUser 创建SSHUser
// @Tags SSHUser
// @Summary 创建SSHUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SSHUser true "创建SSHUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sshUser/createSSHUser [post]
func CreateSSHUser(c *gin.Context) {
	var sshUser model.SSHUser
	_ = c.ShouldBindJSON(&sshUser)
	if err := service.CreateSSHUser(sshUser); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteSSHUser 删除SSHUser
// @Tags SSHUser
// @Summary 删除SSHUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SSHUser true "删除SSHUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sshUser/deleteSSHUser [delete]
func DeleteSSHUser(c *gin.Context) {
	var sshUser model.SSHUser
	_ = c.ShouldBindJSON(&sshUser)
	if err := service.DeleteSSHUser(sshUser); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteSSHUserByIds 批量删除SSHUser
// @Tags SSHUser
// @Summary 批量删除SSHUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SSHUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /sshUser/deleteSSHUserByIds [delete]
func DeleteSSHUserByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteSSHUserByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateSSHUser 更新SSHUser
// @Tags SSHUser
// @Summary 更新SSHUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SSHUser true "更新SSHUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sshUser/updateSSHUser [put]
func UpdateSSHUser(c *gin.Context) {
	var sshUser model.SSHUser
	_ = c.ShouldBindJSON(&sshUser)
	if err := service.UpdateSSHUser(sshUser); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindSSHUser 用id查询SSHUser
// @Tags SSHUser
// @Summary 用id查询SSHUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SSHUser true "用id查询SSHUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sshUser/findSSHUser [get]
func FindSSHUser(c *gin.Context) {
	var sshUser model.SSHUser
	_ = c.ShouldBindQuery(&sshUser)
	if err, resshUser := service.GetSSHUser(sshUser.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"resshUser": resshUser}, c)
	}
}

func GetSSHUsers(c *gin.Context) {
	var (
		users []model.SSHUser
		err   error
	)
	err, users = service.GetSSHUserList()
	if err != nil {
		response.FailWithMessage("获取用户列表失败", c)
	} else {
		response.OkWithDetailed(users, "获取成功", c)
	}

}

// GetSSHUserList 分页获取SSHUser列表
// @Tags SSHUser
// @Summary 分页获取SSHUser列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SSHUserSearch true "分页获取SSHUser列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sshUser/getSSHUserList [get]
func GetSSHUserList(c *gin.Context) {
	var pageInfo request.SSHUserSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetSSHUserInfoList(pageInfo); err != nil {
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
