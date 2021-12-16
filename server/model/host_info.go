// 自动生成模板HostInfo
package model

import (
	"gin-vue-admin/global"
)

// HostInfo 结构体
// 如果含有time.Time 请自行import time包
type HostInfo struct {
	global.GVA_MODEL
	IP        string `json:"ip" form:"ip" gorm:"column:ip;comment:;type:varchar(128);"`
	HostName  string `json:"hostname" form:"hostname" gorm:"column:hostname;comment:主机名;type:varchar(128);"`
	Port      int    `json:"port" form:"port" gorm:"column:port;comment:;type:int;size:11;"`
	Active    *bool  `json:"active" form:"active" gorm:"column:active;comment:是否启用;type:tinyint"`
	SSHUserID uint   `json:"user_id" form:"user_id" gorm:"column:user_id;comment:用户ID;type:int;size:11;"`
	SSHUser   SSHUser
	GroupId   int `json:"group_id" form:"group_id" gorm:"column:group_id;comment:主机分组ID;type:int;size:6;"`
}

// TableName HostInfo 表名
func (HostInfo) TableName() string {
	return "host"
}
