package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitWebTerminalRouter(Router *gin.RouterGroup) {
	WebTerminalRouter := Router.Group("ws").Use(middleware.OperationRecord())
	{
		WebTerminalRouter.GET(":id", v1.WebTerminal)             // 远程
	}
}
