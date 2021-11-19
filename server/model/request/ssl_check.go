package request

import "gin-vue-admin/model"

type SSLCheckSearch struct {
	model.SSLCheck
	PageInfo
}
