package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

// InitHostInfoRouter 初始化 HostInfo 路由信息
func InitHostInfoRouter(Router *gin.RouterGroup) {
	HostInfoRouter := Router.Group("host").Use(middleware.OperationRecord())
	{
		HostInfoRouter.POST("createHostInfo", v1.CreateHostInfo)   // 新建HostInfo
		HostInfoRouter.DELETE("deleteHostInfo", v1.DeleteHostInfo) // 删除HostInfo
		HostInfoRouter.DELETE("deleteHostInfoByIds", v1.DeleteHostInfoByIds) // 批量删除HostInfo
		HostInfoRouter.PUT("updateHostInfo", v1.UpdateHostInfo)    // 更新HostInfo
		HostInfoRouter.GET("findHostInfo", v1.FindHostInfo)        // 根据ID获取HostInfo
		HostInfoRouter.GET("getHostInfoList", v1.GetHostInfoList)  // 获取HostInfo列表
	}
}
