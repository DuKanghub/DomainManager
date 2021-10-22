package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// CreateDnsUser 创建DnsUser记录
// Author [dukanghub](https://github.com/Dukanghub)
func CreateDnsUser(dnsUser model.DnsUser) (err error) {
	err = global.GVA_DB.Create(&dnsUser).Error
	return err
}

// DeleteDnsUser 删除DnsUser记录
// Author [dukanghub](https://github.com/Dukanghub)
func DeleteDnsUser(dnsUser model.DnsUser) (err error) {
	err = global.GVA_DB.Delete(&dnsUser).Error
	return err
}

// DeleteDnsUserByIds 批量删除DnsUser记录
// Author [dukanghub](https://github.com/Dukanghub)
func DeleteDnsUserByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.DnsUser{},"id in ?",ids.Ids).Error
	return err
}

// UpdateDnsUser 更新DnsUser记录
// Author [dukanghub](https://github.com/Dukanghub)
func UpdateDnsUser(dnsUser model.DnsUser) (err error) {
	err = global.GVA_DB.Save(&dnsUser).Error
	return err
}

// GetDnsUserById 根据id获取DnsUser记录
// Author [dukanghub](https://github.com/Dukanghub)
func GetDnsUserById(id uint) (err error, dnsUser model.DnsUser) {
	err = global.GVA_DB.Where("id = ?", id).First(&dnsUser).Error
	return
}

// GetDnsUserList 分页获取DnsUser记录
// Author [dukanghub](https://github.com/Dukanghub)
func GetDnsUserList(info request.DnsUserSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.DnsUser{})
    var dnsUsers []model.DnsUser
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&dnsUsers).Error
	return err, dnsUsers, total
}
