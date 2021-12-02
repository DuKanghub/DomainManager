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

// CreateHostInfo 创建HostInfo
// @Tags HostInfo
// @Summary 创建HostInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HostInfo true "创建HostInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /host/createHostInfo [post]
func CreateHostInfo(c *gin.Context) {
	var host model.HostInfo
	_ = c.ShouldBindJSON(&host)
	if err := service.CreateHostInfo(host); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteHostInfo 删除HostInfo
// @Tags HostInfo
// @Summary 删除HostInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HostInfo true "删除HostInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /host/deleteHostInfo [delete]
func DeleteHostInfo(c *gin.Context) {
	var host model.HostInfo
	_ = c.ShouldBindJSON(&host)
	if err := service.DeleteHostInfo(host); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteHostInfoByIds 批量删除HostInfo
// @Tags HostInfo
// @Summary 批量删除HostInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除HostInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /host/deleteHostInfoByIds [delete]
func DeleteHostInfoByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteHostInfoByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateHostInfo 更新HostInfo
// @Tags HostInfo
// @Summary 更新HostInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HostInfo true "更新HostInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /host/updateHostInfo [put]
func UpdateHostInfo(c *gin.Context) {
	var host model.HostInfo
	_ = c.ShouldBindJSON(&host)
	if err := service.UpdateHostInfo(host); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindHostInfo 用id查询HostInfo
// @Tags HostInfo
// @Summary 用id查询HostInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HostInfo true "用id查询HostInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /host/findHostInfo [get]
func FindHostInfo(c *gin.Context) {
	var host model.HostInfo
	_ = c.ShouldBindQuery(&host)
	if err, rehost := service.GetHostInfo(host.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rehost": rehost}, c)
	}
}

// GetHostInfoList 分页获取HostInfo列表
// @Tags HostInfo
// @Summary 分页获取HostInfo列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.HostInfoSearch true "分页获取HostInfo列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /host/getHostInfoList [get]
func GetHostInfoList(c *gin.Context) {
	var pageInfo request.HostInfoSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetHostInfoInfoList(pageInfo); err != nil {
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
