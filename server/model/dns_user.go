// 自动生成模板DnsUsers
package model

import (
	"gin-vue-admin/global"
)

// DnsUser 结构体
// 如果含有time.Time 请自行import time包
type DnsUser struct {
      global.GVA_MODEL
      Name  string `json:"name" form:"name" gorm:"column:name;comment:dns子用户名;type:varchar(191);"`
      Platform  int `json:"platform" form:"platform" gorm:"column:platform;comment:平台"`
      AccessKey  string `json:"accessKey" form:"accessKey" gorm:"column:access_key;comment:AccessKey;type:varchar(191);"`
      AccessSecret  string `json:"accessSecret" form:"accessSecret" gorm:"column:access_secret;comment:AccessSecret;type:varchar(191);"`
      Account  int `json:"account" form:"account" gorm:"column:account;comment:所属账号"`
}

// TableName DnsUser 表名
func (DnsUser) TableName() string {
  return "dns_users"
}

