package v1

import (
	"fmt"
	"gin-vue-admin/api/v1/webssh"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/response"
	"gin-vue-admin/utils"
	"github.com/gin-gonic/gin"
)



func WebTerminal(c *gin.Context) {
	hostId := c.Param("id")
	if hostId == "" {
		response.FailWithMessage("没有指定主机", c)
		return
	}
	fmt.Println("hostId", hostId)
	var host model.HostInfo
	err := global.GVA_DB.Preload("SSHUser").Where("id = ?", hostId).First(&host).Error
	if err != nil {
		response.FailWithMessage("查询数据库失败", c)
		return
	}
	fmt.Printf("host: %+v\n", host)
	passBytes, err := utils.DecodeStr(host.SSHUser.Password)
	if err != nil {
		response.FailWithMessage("无效的密码", c)
		return
	}
	fmt.Println("password:", string(passBytes))
	config := &webssh.WebSSHConfig{
		Record: false,
		RecPath: "./rec",
		RemoteAddr: fmt.Sprintf("%s:%d", host.IP, host.Port),
		User:       host.SSHUser.Username,
		Password:   string(passBytes),
		AuthModel:  webssh.PASSWORD,
	}
	handle := webssh.NewWebSSH(config)
	handle.ServeConn(c)
}


