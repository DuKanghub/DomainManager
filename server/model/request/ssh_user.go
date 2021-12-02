package request

import "gin-vue-admin/model"

type SSHUserSearch struct{
    model.SSHUser
    PageInfo
}