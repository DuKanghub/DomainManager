package service

import (
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"gin-vue-admin/utils"
)

// CreateSSLCheck 创建SSLCheck记录
// Author [piexlmax](https://github.com/piexlmax)
func CreateSSLCheck(sslCheck model.SSLCheck) (err error) {
	url := sslCheck.Url
	fmt.Println(url)
	info, err := utils.GetCertInfo(url)
	if err != nil {
		return err
	}
	err = global.GVA_DB.Create(&info).Error
	return err
}

// DeleteSSLCheck 删除SSLCheck记录
// Author [piexlmax](https://github.com/piexlmax)
func DeleteSSLCheck(sslCheck model.SSLCheck) (err error) {
	err = global.GVA_DB.Delete(&sslCheck).Error
	return err
}

// DeleteSSLCheckByIds 批量删除SSLCheck记录
// Author [piexlmax](https://github.com/piexlmax)
func DeleteSSLCheckByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.SSLCheck{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateSSLCheck 更新SSLCheck记录
// Author [piexlmax](https://github.com/piexlmax)
func UpdateSSLCheck(sslCheck model.SSLCheck) (err error) {
	err = global.GVA_DB.Save(&sslCheck).Error
	return err
}

// GetSSLCheck 根据id获取SSLCheck记录
// Author [piexlmax](https://github.com/piexlmax)
func GetSSLCheck(id uint) (err error, sslCheck model.SSLCheck) {
	err = global.GVA_DB.Where("id = ?", id).First(&sslCheck).Error
	return
}

// GetSSLCheckInfoList 分页获取SSLCheck记录
// Author [piexlmax](https://github.com/piexlmax)
func GetSSLCheckInfoList(info request.SSLCheckSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.SSLCheck{})
	var sslChecks []model.SSLCheck
	// 如果有条件搜索 下方会自动创建搜索语句
	db = db.Order("expiredAt") //排序
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&sslChecks).Error
	return err, sslChecks, total
}
