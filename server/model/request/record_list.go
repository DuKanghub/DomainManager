package request

import "gin-vue-admin/model"

type RecordListSearch struct{
	model.Record
	PageInfo
}