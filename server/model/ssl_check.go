// 自动生成模板SSLCheck
package model

import (
	"gin-vue-admin/global"
	"time"
)

// SSLCheck 结构体
// 如果含有time.Time 请自行import time包
type SSLCheck struct {
	global.GVA_MODEL
	Url        string    `json:"url" form:"url" gorm:"column:url;comment:域名链接或域名;type:varchar(191);"`
	ExpiredAt  time.Time `json:"expiredAt" form:"expiredAt" gorm:"column:expiredAt;comment:证书过期时间;type:datetime;"`
	CertDomain string    `json:"certDomain" form:"certDomain" gorm:"column:cert_domain;comment:证书所属域名;type:varchar(191);"`
}

// TableName SSLCheck 表名
func (SSLCheck) TableName() string {
	return "ssl_check"
}
