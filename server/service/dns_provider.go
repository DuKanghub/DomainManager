package service

import (
	"gin-vue-admin/model"
	"gin-vue-admin/service/ali"
	"gin-vue-admin/service/dnspod"
)

type DNSProvider interface {
	AddRecord(*model.Record) error
	DelRecord(*model.Record) error
	UpdateRecord(*model.Record) error
	FlushRecordsToDb(*model.Record) error
	FlushDomainsToDb(model.DnsUser) error
}

func NewDnsProvider(user model.DnsUser) DNSProvider {
	switch user.Platform {
	case 1:
		return &ali.AliDnsProvider{}
	case 2:
		return &dnspod.TencentDnsProvider{}
	default:
		return nil
	}
}
