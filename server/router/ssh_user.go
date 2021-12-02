package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

// InitSSHUserRouter 初始化 SSHUser 路由信息
func InitSSHUserRouter(Router *gin.RouterGroup) {
	SSHUserRouter := Router.Group("sshUser").Use(middleware.OperationRecord())
	{
		SSHUserRouter.POST("createSSHUser", v1.CreateSSHUser)   // 新建SSHUser
		SSHUserRouter.DELETE("deleteSSHUser", v1.DeleteSSHUser) // 删除SSHUser
		SSHUserRouter.DELETE("deleteSSHUserByIds", v1.DeleteSSHUserByIds) // 批量删除SSHUser
		SSHUserRouter.PUT("updateSSHUser", v1.UpdateSSHUser)    // 更新SSHUser
		SSHUserRouter.GET("findSSHUser", v1.FindSSHUser)        // 根据ID获取SSHUser
		SSHUserRouter.GET("getSSHUserList", v1.GetSSHUserList)  // 获取SSHUser列表
	}
}
