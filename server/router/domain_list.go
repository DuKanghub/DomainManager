package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

// InitDomainListRouter 初始化 DomainInfo 路由信息
func InitDomainListRouter(Router *gin.RouterGroup) {
	DomainListRouter := Router.Group("domainList").Use(middleware.OperationRecord())
	{
		DomainListRouter.POST("createDomain", v1.CreateDomain)             // 新建Domain
		DomainListRouter.DELETE("deleteDomain", v1.DeleteDomain)           // 删除Domain
		DomainListRouter.DELETE("deleteDomainByIds", v1.DeleteDomainByIds) // 批量删除Domain
		DomainListRouter.PUT("updateDomain", v1.UpdateDomain)              // 更新Domain
		DomainListRouter.GET("findDomain", v1.FindDomain)                  // 根据ID获取Domain
		DomainListRouter.GET("getDomainList", v1.GetDomainList)            // 获取Domain列表
		DomainListRouter.PUT("updateDomains", v1.UpdateDomains)            // 批量更新域名信息
		DomainListRouter.GET("flushDomains", v1.FlushDomainsToDb)          // 强制刷新域名缓存
	}
}
