// 自动生成模板DomainLists
package model

import (
	"gin-vue-admin/global"
      "time"
)

// DomainInfo 结构体
// 如果含有time.Time 请自行import time包
type DomainInfo struct {
      global.GVA_MODEL
      Domain  string `json:"domain" form:"domain" gorm:"column:domain;comment:域名;type:varchar(191);"`
      Company  int `json:"company" form:"company" gorm:"column:company;comment:公司;type:int;"`
      RecordTotal  int `json:"recordTotal" form:"recordTotal" gorm:"column:record_total;comment:解析数;type:bigint;size:19;"`
      DnsProvider  int `json:"dnsProvider" form:"dnsProvider" gorm:"column:dns_provider;comment:dns提供商;type:bigint;size:19;"`
      Account  int `json:"account" form:"account" gorm:"column:account;comment:所属账号;type:int;"`
      RegisterTime  time.Time `json:"registerTime" form:"registerTime" gorm:"column:register_time;comment:注册日期;type:datetime;"`
      DomainId string `json:"domainId" form:"domainId" gorm:"column:domain_id;comment:域名id"`
}
type Domains struct {
      Domains []DomainInfo `json:"domains" form:"domains"`
}

// TableName DomainInfo 表名
func (DomainInfo) TableName() string {
  return "domain_lists"
}

