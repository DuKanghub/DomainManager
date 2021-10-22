package model

import "gin-vue-admin/global"

type Record struct {
	global.GVA_MODEL
	Domain string `json:"domain" form:"domain" gorm:"column:domain;comment:域名名称;type:varchar(100);"`
	RR string `json:"rr" form:"rr" gorm:"column:rr;comment:主机记录;type:varchar(50);"`
	Type int `json:"type" form:"type" gorm:"column:type;comment:记录类型;type:int;size:2;"`
	Line int `json:"line" form:"line" gorm:"column:line;comment:解析线路;type:int;size:3;"`
	Value string `json:"value" form:"value" gorm:"column:value;comment:记录值;type:varchar(191);"`
	RecordId string `json:"record_id" form:"record_id" gorm:"column:record_id;comment:解析记录id;type:varchar(191);"`
	Status int `json:"status" form:"status" gorm:"column:status;comment:解析状态;type:int;size:1;"`
	Account int `json:"account" form:"account" gorm:"column:account;comment:所属账号;type:int;size:10;"`
}

type Records struct {
	Records []Record `json:"records" form:"records"`
}

// TableName Record 表名
func (Record) TableName() string {
	return "record_list"
}
