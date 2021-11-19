package utils

import (
	"crypto/tls"
	"errors"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"go.uber.org/zap"
	"mvdan.cc/xurls/v2"
	"strings"
)

//@param: domain string 可以是域名：dev.dukanghub.com，也可以是链接: https://dev.dukanghub.com
func GetCertInfo(domain string) (info model.SSLCheck, err error) {
	// 先打乱传入的文本，再提取域名
	tmp := strings.ReplaceAll(domain, "/", " ")
	rxRelaxed := xurls.Relaxed()
	str := rxRelaxed.FindString(tmp)
	if str == "" {
		global.GVA_LOG.Error("未从传入的参数中找到有效的域名", zap.Any("入参", domain))
		return model.SSLCheck{}, errors.New("传入域名格式不对")
	}
	// 将提取后的域名加上tls端口, 这段看自己需要，可以自己显式传入，也可能不是443端口
	str = strings.Join([]string{str, "443"}, ":")
	conn, err := tls.Dial("tcp", str, nil)
	if err != nil {
		global.GVA_LOG.Error("域名tcp连接失败!", zap.Any("domain", str), zap.Any("err", err))
		return model.SSLCheck{}, err
	}
	defer func() {
		err = conn.Close()
		if err != nil {
			global.GVA_LOG.Error("关闭conn失败!", zap.Any("err", err))
		}
	}()
	certs := conn.ConnectionState().PeerCertificates
	if len(certs) > 0 {
		certInfo := certs[0]
		expireTime := certInfo.NotAfter.Local() //加local是显示为当地时间
		subject := certInfo.Subject
		info.ExpiredAt = expireTime
		info.CertDomain = subject.CommonName
		info.Url = domain
	} else {
		err = errors.New("未找到证书信息")
	}
	return info, err
}
