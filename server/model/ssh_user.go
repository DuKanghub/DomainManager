// 自动生成模板SSHUser
package model

import (
	"gin-vue-admin/global"
)

// SSHUser 结构体
// 如果含有time.Time 请自行import time包
type SSHUser struct {
	global.GVA_MODEL
	Name       string `json:"name,omitempty" form:"name" gorm:"column:name;comment:易于识别的名称;type:varchar(128);"`
	Username   string `json:"username,omitempty" form:"username" gorm:"column:username;comment:远程用户名;type:varchar(128);"`
	Password   string `json:"password,omitempty" form:"password" gorm:"column:password;comment:远程密码;type:varchar(512);"`
	PrivateKey string `json:"privateKey,omitempty" form:"privateKey" gorm:"column:private_key;comment:私钥加密后文本;type:longtext;"`
	Become     *bool  `json:"become,omitempty" form:"become" gorm:"column:become;comment:启用sudo;type:tinyint"`
	KeyAuth    *bool  `json:"key_auth,omitempty" form:"key_auth" gorm:"column:key_auth;comment:秘钥认证;type:tinyint"`
	Comment    string `json:"comment,omitempty" form:"comment" gorm:"column:comment;comment:备注;type:longtext;"`
}

// TableName SSHUser 表名
func (SSHUser) TableName() string {
	return "ssh_user"
}
