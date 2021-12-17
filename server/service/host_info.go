package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// CreateHostInfo 创建HostInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func CreateHostInfo(host model.HostInfo) (err error) {
	err = global.GVA_DB.Create(&host).Error
	return err
}

// DeleteHostInfo 删除HostInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func DeleteHostInfo(host model.HostInfo) (err error) {
	err = global.GVA_DB.Delete(&host).Error
	return err
}

// DeleteHostInfoByIds 批量删除HostInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func DeleteHostInfoByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.HostInfo{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateHostInfo 更新HostInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func UpdateHostInfo(host model.HostInfo) (err error) {
	err = global.GVA_DB.Save(&host).Error
	return err
}

// GetHostInfo 根据id获取HostInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func GetHostInfo(id uint) (err error, host model.HostInfo) {
	err = global.GVA_DB.Where("id = ?", id).First(&host).Error
	return
}

// GetHostInfoInfoList 分页获取HostInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func GetHostInfoInfoList(info request.HostInfoSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.HostInfo{})
	var hosts []model.HostRes
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.IP != "" {
		db = db.Where("`ip` LIKE ?", "%"+info.IP+"%")
	}
	if info.HostName != "" {
		db = db.Where("`hostname` LIKE ?", "%"+info.HostName+"%")
	}
	if info.Port != 0 {
		db = db.Where("`port` = ?", info.Port)
	}
	if info.Active != nil {
		db = db.Where("`active` = ?", info.Active)
	}
	err = db.Count(&total).Error
	sql := "select t.id,t.created_at as CreatedAt,t.ip,t.hostname as HostName,t.port,t.active,t.user_id as SSHUserID,t.group_id as GroupId,d.name from host as t,ssh_user as d where t.deleted_at is null and t.user_id = d.id;"
	err = db.Limit(limit).Offset(offset).Raw(sql).Scan(&hosts).Error
	//err = db.Preload("SSHUser", func(db *gorm.DB) *gorm.DB {
	//	return db.Select("id,name")
	//}).Limit(limit).Offset(offset).Find(&hosts).Error
	return err, hosts, total
}
