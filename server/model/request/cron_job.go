package request

import "gin-vue-admin/model"

type CronJobSearch struct {
	model.CronJob
	PageInfo
}
