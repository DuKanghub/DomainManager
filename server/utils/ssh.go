package utils

import (
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"go.uber.org/zap"
	"golang.org/x/crypto/ssh"
)

type SSHClient struct {
	Host  model.HostInfo
	SUser model.SSHUser
}

func (s *SSHClient) RunCmd(cmd string) (stdout string, err error) {
	config := &ssh.ClientConfig{
		User:            s.SUser.Username,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), //这个可以, 但是不够安全
	}
	if s.SUser.KeyAuth != nil && *(s.SUser.KeyAuth) == true {
		if s.SUser.PrivateKey == "" {
			return "", err
		}
		key, err := DecodeStr(s.SUser.PrivateKey)
		if err != nil {
			return "", err
		}
		signer, err := ssh.ParsePrivateKey(key)
		if err != nil {
			return "", err
		}
		config.Auth = []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		}
	} else {
		key, err := DecodeStr(s.SUser.Password)
		if err != nil {
			return "", err
		}
		config.Auth = []ssh.AuthMethod{ssh.Password(string(key))}
	}
	//dial 获取ssh client
	addr := fmt.Sprintf("%s:%d", s.Host.IP, s.Host.Port)
	sshClient, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		global.GVA_LOG.Error("ssh连接错误!", zap.Any("err", err))
		return "", err
	}
	defer sshClient.Close()

	//创建ssh-session
	session, err := sshClient.NewSession()
	if err != nil {
		global.GVA_LOG.Error("创建ssh session 失败", zap.Any("err", err))
		return "", err
	}
	defer session.Close()
	//执行远程命令
	combo, err := session.CombinedOutput(cmd)
	if err != nil {
		global.GVA_LOG.Error("远程执行cmd 失败", zap.Any("err", err))
		return "", err
	}
	return string(combo), err
}
