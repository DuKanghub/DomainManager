package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

// InitCronJobRouter 初始化 CronJob 路由信息
func InitCronJobRouter(Router *gin.RouterGroup) {
	CronJobRouter := Router.Group("cronJob").Use(middleware.OperationRecord())
	{
		CronJobRouter.POST("createCronJob", v1.CreateCronJob)             // 新建CronJob
		CronJobRouter.DELETE("deleteCronJob", v1.DeleteCronJob)           // 删除CronJob
		CronJobRouter.DELETE("deleteCronJobByIds", v1.DeleteCronJobByIds) // 批量删除CronJob
		CronJobRouter.PUT("updateCronJob", v1.UpdateCronJob)              // 更新CronJob
		CronJobRouter.GET("findCronJob", v1.FindCronJob)                  // 根据ID获取CronJob
		CronJobRouter.GET("getCronJobList", v1.GetCronJobList)            // 获取CronJob列表
	}
}
