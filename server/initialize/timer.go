package initialize

import (
	"fmt"
	"gin-vue-admin/config"
	"gin-vue-admin/global"
	"gin-vue-admin/service"
	"gin-vue-admin/utils"
)

func Timer() {
	if global.GVA_CONFIG.Timer.Start {
		for _, detail := range global.GVA_CONFIG.Timer.Detail {
			go func(detail config.Detail) {
				global.GVA_Timer.AddTaskByFunc("ClearDB", global.GVA_CONFIG.Timer.Spec, func() {
					err := utils.ClearTable(global.GVA_DB, detail.TableName, detail.CompareField, detail.Interval)
					if err != nil {
						fmt.Println("timer error:", err)
					}
				})
			}(detail)
		}
		go func() {
			global.GVA_Timer.AddTaskByFunc("CheckAllSSL", "@daily", func() {
				// 检查所有30天内过期的证书
				service.CheckAllSSL()
			})
		}()
	}
}
