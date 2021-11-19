package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

// InitSSLCheckRouter 初始化 SSLCheck 路由信息
func InitSSLCheckRouter(Router *gin.RouterGroup) {
	SSLCheckRouter := Router.Group("sslCheck").Use(middleware.OperationRecord())
	{
		SSLCheckRouter.POST("createSSLCheck", v1.CreateSSLCheck)             // 新建SSLCheck
		SSLCheckRouter.DELETE("deleteSSLCheck", v1.DeleteSSLCheck)           // 删除SSLCheck
		SSLCheckRouter.DELETE("deleteSSLCheckByIds", v1.DeleteSSLCheckByIds) // 批量删除SSLCheck
		SSLCheckRouter.PUT("updateSSLCheck", v1.UpdateSSLCheck)              // 更新SSLCheck
		SSLCheckRouter.GET("findSSLCheck", v1.FindSSLCheck)                  // 根据ID获取SSLCheck
		SSLCheckRouter.GET("getSSLCheckList", v1.GetSSLCheckList)            // 获取SSLCheck列表
	}
}
