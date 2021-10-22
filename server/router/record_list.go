package router

import (
	v1 "gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

// InitRecordListRouter 初始化 Record 路由信息
func InitRecordListRouter(Router *gin.RouterGroup) {
	RecordListRouter := Router.Group("recordList").Use(middleware.OperationRecord())
	{
		RecordListRouter.POST("createRecord", v1.CreateRecord)             // 新建Record
		RecordListRouter.DELETE("deleteRecord", v1.DeleteRecord)           // 删除Record
		RecordListRouter.DELETE("deleteRecordByIds", v1.DeleteRecordByIds) // 批量删除Record
		RecordListRouter.PUT("updateRecord", v1.UpdateRecord)              // 更新Record
		RecordListRouter.GET("findRecord", v1.FindRecordById)              // 根据ID获取Record
		RecordListRouter.GET("getRecordList", v1.GetRecordList)            // 获取RecordList列表
		RecordListRouter.PUT("updateRecords", v1.UpdateRecords)            // 批量更新Record
		RecordListRouter.GET("flushRecords", v1.FlushRecords)              // 强制刷新域名解析
	}
}