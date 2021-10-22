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

// CreateDomain 创建Domain
// @Tags DomainInfo
// @Summary 创建Domain
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DomainInfo true "创建Domain"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /domainList/createDomain [post]
func CreateDomain(c *gin.Context) {
	var domain model.DomainInfo
	_ = c.ShouldBindJSON(&domain)
	if err := service.CreateDomain(domain); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteDomain 删除Domain
// @Tags DomainInfo
// @Summary 删除Domain
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DomainInfo true "删除Domain"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /domainList/deleteDomain [delete]
func DeleteDomain(c *gin.Context) {
	var domain model.DomainInfo
	_ = c.ShouldBindJSON(&domain)
	if err := service.DeleteDomain(domain); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteDomainByIds 批量删除Domains
// @Tags DomainInfo
// @Summary 批量删除Domains
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Domains"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /domainList/deleteDomainByIds [delete]
func DeleteDomainByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteDomainByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateDomain 更新Domain
// @Tags DomainInfo
// @Summary 更新Domain
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DomainInfo true "更新Domain"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /domainList/updateDomain [put]
func UpdateDomain(c *gin.Context) {
	var domain model.DomainInfo
	_ = c.ShouldBindJSON(&domain)
	if err := service.UpdateDomain(domain); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// UpdateDomains 批量更新Domain
// @Tags DomainInfo
// @Summary 批量更新Domain
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body []model.DomainInfo true "批量更新Domain"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量更新成功"}"
// @Router /domainList/updateDomains [put]
func UpdateDomains(c *gin.Context) {
	var domains model.Domains
	_ = c.ShouldBindJSON(&domains)
	if err := service.UpdateDomains(domains); err != nil {
		global.GVA_LOG.Error("批量更新失败!", zap.Any("err", err))
		response.FailWithMessage("批量更新失败", c)
	} else {
		response.OkWithMessage("批量更新成功", c)
	}
}

// FindDomain 用id查询Domain
// @Tags DomainInfo
// @Summary 用id查询Domain
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DomainInfo true "用id查询Domain"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /domainList/findDomain [get]
func FindDomain(c *gin.Context) {
	var domain model.DomainInfo
	_ = c.ShouldBindQuery(&domain)
	if err, reDomain := service.GetDomainById(domain.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reDomain": reDomain}, c)
	}
}

// GetDomainList 分页获取Domain列表
// @Tags DomainInfo
// @Summary 分页获取Domain列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.DomainListSearch true "分页获取Domain列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /domainList/getDomainList [get]
func GetDomainList(c *gin.Context) {
	var pageInfo request.DomainListSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetDomainList(pageInfo); err != nil {
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

// FlushDomainsToDb 从线上获取最新域名信息并更新到数据库
// @Tags DomainInfo
// @Summary 从线上获取最新域名信息并更新到数据库
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /domainList/flushDomains [get]
func FlushDomainsToDb(c *gin.Context) {
	err := service.FlushDomainsToDb()
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithMessage("获取成功", c)
	}
}