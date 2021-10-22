package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

// InitDnsUserRouter 初始化 DnsUser 路由信息
func InitDnsUserRouter(Router *gin.RouterGroup) {
	DnsUserRouter := Router.Group("dnsUser").Use(middleware.OperationRecord())
	{
		DnsUserRouter.POST("createDnsUser", v1.CreateDnsUser)             // 新建DnsUser
		DnsUserRouter.DELETE("deleteDnsUser", v1.DeleteDnsUser)           // 删除DnsUser
		DnsUserRouter.DELETE("deleteDnsUserByIds", v1.DeleteDnsUserByIds) // 批量删除DnsUser
		DnsUserRouter.PUT("updateDnsUser", v1.UpdateDnsUser)              // 更新DnsUser
		DnsUserRouter.GET("findDnsUser", v1.FindDnsUserById)              // 根据ID获取DnsUser
		DnsUserRouter.GET("getDnsUserList", v1.GetDnsUserList)            // 获取DnsUser列表
	}
}
