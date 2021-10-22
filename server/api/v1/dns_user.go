package v1

import (
	"gin-vue-admin/global"
    "gin-vue-admin/model"
    "gin-vue-admin/model/request"
    "gin-vue-admin/model/response"
    "gin-vue-admin/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

// CreateDnsUser 创建DnsUser
// @Tags DnsUser
// @Summary 创建DnsUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DnsUser true "创建DnsUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dnsUser/createDnsUser [post]
func CreateDnsUser(c *gin.Context) {
	var dnsUser model.DnsUser
	_ = c.ShouldBindJSON(&dnsUser)
	if err := service.CreateDnsUser(dnsUser); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteDnsUser 删除DnsUser
// @Tags DnsUser
// @Summary 删除DnsUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DnsUser true "删除DnsUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dnsUser/deleteDnsUser [delete]
func DeleteDnsUser(c *gin.Context) {
	var dnsUser model.DnsUser
	_ = c.ShouldBindJSON(&dnsUser)
	if err := service.DeleteDnsUser(dnsUser); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteDnsUserByIds 批量删除DnsUser
// @Tags DnsUser
// @Summary 批量删除DnsUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除DnsUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /dnsUser/deleteDnsUserByIds [delete]
func DeleteDnsUserByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteDnsUserByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateDnsUser 更新DnsUser
// @Tags DnsUser
// @Summary 更新DnsUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DnsUser true "更新DnsUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /dnsUser/updateDnsUser [put]
func UpdateDnsUser(c *gin.Context) {
	var dnsUser model.DnsUser
	_ = c.ShouldBindJSON(&dnsUser)
	if err := service.UpdateDnsUser(dnsUser); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindDnsUserById 用id查询DnsUser
// @Tags DnsUser
// @Summary 用id查询DnsUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DnsUser true "用id查询DnsUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /dnsUser/findDnsUser [get]
func FindDnsUserById(c *gin.Context) {
	var dnsUser model.DnsUser
	_ = c.ShouldBindQuery(&dnsUser)
	if err, reDnsUser := service.GetDnsUserById(dnsUser.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reDnsUser": reDnsUser}, c)
	}
}

// GetDnsUserList 分页获取DnsUser列表
// @Tags DnsUser
// @Summary 分页获取DnsUser列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.DnsUserSearch true "分页获取DnsUser列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dnsUser/getDnsUserList [get]
func GetDnsUserList(c *gin.Context) {
	var pageInfo request.DnsUserSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetDnsUserList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}
