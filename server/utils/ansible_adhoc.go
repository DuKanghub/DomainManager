package utils

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"github.com/apenella/go-ansible/pkg/adhoc"
	"github.com/apenella/go-ansible/pkg/execute"
	"github.com/apenella/go-ansible/pkg/options"
	"go.uber.org/zap"
	"io"
	"strconv"
	"strings"
)

type AnsibleAdHoc struct {
	Host  model.HostInfo
	SUser model.SSHUser
	adh *adhoc.AnsibleAdhocCmd
}

func NewCmdClient(a *AnsibleAdHoc) {
	ansibleConnectionOptions := &options.AnsibleConnectionOptions{
		User:         a.SUser.Username,
		AskPass:      false,
		SSHCommonArgs: "'-o StrictHostKeyChecking=no'",
		SSHExtraArgs: "'-p " + strconv.Itoa(a.Host.Port) + "'",
	}
	ansibleAdhocOptions := &adhoc.AnsibleAdhocOptions{
		//Inventory指定主机清单文件或逗号分隔的主机，默认为/etc/ansible/hosts。
		Inventory:   a.Host.IP + ",",
		SyntaxCheck: true,
	}
	var privilegeEscalationOptions *options.AnsiblePrivilegeEscalationOptions
	if a.SUser.Become != nil && *(a.SUser.Become) == true {
		privilegeEscalationOptions = &options.AnsiblePrivilegeEscalationOptions{
			Become:       true,
			BecomeMethod: "sudo",
			BecomeUser:   "root",
		}
	}
	a.adh = &adhoc.AnsibleAdhocCmd{
		Pattern:                    "all",
		Options:                    ansibleAdhocOptions,        //命令详情
		ConnectionOptions:          ansibleConnectionOptions,   //连接选项
		PrivilegeEscalationOptions: privilegeEscalationOptions, //特权升级选项
	}
}

func (a *AnsibleAdHoc) Ping() (stdout string, err error) {
	NewCmdClient(a)
	a.adh.Options.ModuleName = "ping"
	return Run(a)
}

func newExecutor(a *AnsibleAdHoc) *bytes.Buffer {
	buff := new(bytes.Buffer)
	exec := execute.NewDefaultExecute(
		execute.WithWrite(io.Writer(buff)),
	)
	a.adh.Exec = exec
	return buff
}

func Run(a *AnsibleAdHoc) (stdout string, err error) {
	buff := newExecutor(a)
	fmt.Println("执行的命令是：", a.adh.String())
	err = a.adh.Run(context.TODO())
	if err != nil {
		global.GVA_LOG.Error("adh.Run出错", zap.Any("err", err))
		return "", err
	}
	return buff.String(), err
}

func (a *AnsibleAdHoc) RunCmd(cmd string) (stdout string, err error) {
	NewCmdClient(a)
	a.adh.Options.ModuleName = "shell"
	a.adh.Options.Args = cmd
	return Run(a)
}
func parseCmd(a *AnsibleAdHoc, module,cmd string)  {
	a.adh.Options.ModuleName = module
	a.adh.Options.Args = cmd
}

func (a *AnsibleAdHoc) RunCmdTest(cmd string) error {
	NewCmdClient(a)
	a.adh.Options.Verbose = true
	parseCmd(a, "shell", cmd)
	fmt.Println("执行的命令是：", a.adh.String())
	err := a.adh.Run(context.TODO())
	if err != nil {
		global.GVA_LOG.Error("adh.Run出错", zap.Any("err", err))
		return err
	}
	return err
}

func (a *AnsibleAdHoc) DeployCronJob(c model.CronJob) (stdout string, err error) {
	NewCmdClient(a)
	a.adh.Options.ModuleName = "cron"
	arr := strings.Fields(c.Spec)
	if len(arr) != 5 {
		return "", errors.New("执行频率要求5个字段：分 时 日 月 周")
	}
	cmd := fmt.Sprintf("cron_file=/etc/crontab name=\"%s\" job=\"%s\" minute=%s hour=%s day=%s month=%s weekday=%s user=root", c.JobName, c.Command, arr[0], arr[1], arr[2], arr[3], arr[4])
	if c.Enabled != nil && *(c.Enabled) == false {
		cmd = fmt.Sprintf("%s disabled=true", cmd)
	}
	a.adh.Options.Args = cmd
	return Run(a)
}
