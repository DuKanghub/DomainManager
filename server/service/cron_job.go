package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// CreateCronJob 创建CronJob记录
// Author [piexlmax](https://github.com/piexlmax)
func CreateCronJob(cronJob model.CronJob) (err error) {
	err = global.GVA_DB.Create(&cronJob).Error
	return err
}

// DeleteCronJob 删除CronJob记录
// Author [piexlmax](https://github.com/piexlmax)
func DeleteCronJob(cronJob model.CronJob) (err error) {
	err = global.GVA_DB.Delete(&cronJob).Error
	return err
}

// DeleteCronJobByIds 批量删除CronJob记录
// Author [piexlmax](https://github.com/piexlmax)
func DeleteCronJobByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.CronJob{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateCronJob 更新CronJob记录
// Author [piexlmax](https://github.com/piexlmax)
func UpdateCronJob(cronJob model.CronJob) (err error) {
	err = global.GVA_DB.Save(&cronJob).Error
	return err
}

// GetCronJob 根据id获取CronJob记录
// Author [piexlmax](https://github.com/piexlmax)
func GetCronJob(id uint) (err error, cronJob model.CronJob) {
	err = global.GVA_DB.Where("id = ?", id).First(&cronJob).Error
	return
}

// GetCronJobInfoList 分页获取CronJob记录
// Author [piexlmax](https://github.com/piexlmax)
func GetCronJobInfoList(info request.CronJobSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.CronJob{})
	var cronJobs []model.CronJob
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.JobName != "" {
		db = db.Where("`job_name` LIKE ?", "%"+info.JobName+"%")
	}
	if info.Command != "" {
		db = db.Where("`command` LIKE ?", "%"+info.Command+"%")
	}
	if info.Creator != 0 {
		db = db.Where("`creator` = ?", info.Creator)
	}
	if info.ExecHost != "" {
		db = db.Where("`exec_host` LIKE ?", "%"+info.ExecHost+"%")
	}
	if info.Enabled != nil {
		db = db.Where("`enabled` = ?", info.Enabled)
	}
	if info.Comment != "" {
		db = db.Where("`comment` LIKE ?", "%"+info.Comment+"%")
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&cronJobs).Error
	return err, cronJobs, total
}
