package request

import "gin-vue-admin/model"

type HostInfoSearch struct{
    model.HostInfo
    PageInfo
}