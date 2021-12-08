// 自动生成模板CronJob
package model

import (
	"gin-vue-admin/global"
)

// CronJob 结构体
// 如果含有time.Time 请自行import time包
type CronJob struct {
	global.GVA_MODEL
	JobName  string `json:"job_name" form:"job_name" gorm:"column:job_name;comment:任务名称;type:varchar(255);"`
	Spec     string `json:"spec" form:"spec" gorm:"column:spec;comment:执行频率;type:varchar(255);"`
	Command  string `json:"command" form:"command" gorm:"column:command;comment:命令;type:varchar(255);"`
	Creator  int    `json:"creator" form:"creator" gorm:"column:creator;comment:创建者;type:int;"`
	ExecHost string `json:"exec_host" form:"exec_host" gorm:"column:exec_host;comment:执行主机;type:varchar(255);"`
	Enabled  *bool  `json:"enabled" form:"enabled" gorm:"column:enabled;comment:启用;type:tinyint"`
	Comment  string `json:"comment" form:"comment" gorm:"column:comment;comment:备注;type:varchar(255);"`
}

// TableName CronJob 表名
func (CronJob) TableName() string {
	return "cron_jobs"
}
