package request

import "gin-vue-admin/model"

type DomainListSearch struct{
    model.DomainInfo
    PageInfo
}