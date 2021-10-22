package request

import "gin-vue-admin/model"

type DnsUserSearch struct{
    model.DnsUser
    PageInfo
}