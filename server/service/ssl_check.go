package service

import (
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"gin-vue-admin/utils"
	"go.uber.org/zap"
	"time"
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

// 定时任务函数：检查所有证书是否30天内过期
func CheckAllSSL() {
	var records []model.SSLCheck
	duration, err := time.ParseDuration("720h")	// 30d
	if err != nil || duration < 0 {
		global.GVA_LOG.Error("ParseDuration失败!", zap.Any("err", err))
		return
	}

	db := global.GVA_DB.Where("expiredAt < ?", time.Now().Add(duration)).Find(&records)
	db.Order("expiredAt")
	for _, record := range records {
		remained := record.ExpiredAt.Sub(time.Now())
		// TODO: 这里以后要弄成自动续签证书或通知
		global.GVA_LOG.Warn("证书即将过期", zap.Any("证书", record.CertDomain), zap.Any("剩余时间", remained))
	}
}