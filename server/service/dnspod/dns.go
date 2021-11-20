package dnspod

import (
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	ttdns "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dnspod/v20210323"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type TencentDnsProvider struct {
}
func GetAccount(req *model.Record) (client *ttdns.Client, err error) {
	var dnsUser model.DnsUser
	err = global.GVA_DB.Where("account = ?", req.Account).First(&dnsUser).Error
	if err != nil {
		return nil, err
	} else {
		credential := common.NewCredential(
			dnsUser.AccessKey,
			dnsUser.AccessSecret,
		)
		cpf := profile.NewClientProfile()
		cpf.HttpProfile.Endpoint = "dnspod.tencentcloudapi.com"
		client, err = ttdns.NewClient(credential, "", cpf)
		return
	}
}

func (t TencentDnsProvider) AddRecord(req *model.Record) (err error) {
	var cli *ttdns.Client
	cli, err = GetAccount(req)
	if err != nil {
		return err
	}
	request := ttdns.NewCreateRecordRequest()
	request.Domain = common.StringPtr(req.Domain)
	request.Value = common.StringPtr(req.Value)
	request.SubDomain = common.StringPtr(req.RR)
	var recordType model.SysDictionaryDetail
	err = global.GVA_DB.Where("`sys_dictionary_id` = 11 and `value` = ?", req.Type).First(&recordType).Error
	if err != nil {
		return err
	}
	request.RecordType = common.StringPtr(recordType.Label)
	// 线路这里采用简单写法，因为目前自己环境只有联通和默认线路
	if req.Line == 2 {
		request.RecordLine = common.StringPtr("联通")
	} else {
		request.RecordLine = common.StringPtr("默认")
	}
	var response *ttdns.CreateRecordResponse
	response, err = cli.CreateRecord(request)
	if err != nil {
		return err
	}
	req.RecordId = strconv.Itoa(int(*(response.Response.RecordId)))
	return nil
}

func (t TencentDnsProvider) DelRecord(req *model.Record) (err error) {
	var cli *ttdns.Client
	cli, err = GetAccount(req)
	if err != nil {
		return err
	}
	request := ttdns.NewDeleteRecordRequest()
	request.Domain = common.StringPtr(req.Domain)
	rid, _ := strconv.Atoi(req.RecordId)
	rd := uint64(rid)
	request.RecordId = common.Uint64Ptr(rd)
	_, err = cli.DeleteRecord(request)
	return err
}

func (t TencentDnsProvider) UpdateRecord(req *model.Record) (err error){
	var cli *ttdns.Client
	cli, err = GetAccount(req)
	if err != nil {
		return err
	}
	request := ttdns.NewModifyRecordRequest()
	request.Domain = common.StringPtr(req.Domain)
	request.Value = common.StringPtr(req.Value)
	rid, _ := strconv.Atoi(req.RecordId)
	rd := uint64(rid)
	request.RecordId = common.Uint64Ptr(rd)
	var recordType model.SysDictionaryDetail
	err = global.GVA_DB.Where("`sys_dictionary_id` = 11 and `value` = ?", req.Type).First(&recordType).Error
	if err != nil {
		return err
	}
	request.RecordType = common.StringPtr(recordType.Label)
	// 线路这里采用简单写法，因为目前自己环境只有联通和默认线路
	if req.Line == 2 {
		request.RecordLine = common.StringPtr("联通")
	} else {
		request.RecordLine = common.StringPtr("默认")
	}

	response, err := cli.ModifyRecord(request)
	if err != nil {
		return err
	}
	req.RecordId = strconv.Itoa(int(*(response.Response.RecordId)))
	return nil
}
func (t TencentDnsProvider) FlushRecordsToDb(req *model.Record) (err error) {
	var cli *ttdns.Client
	cli, err = GetAccount(req)
	if err != nil {
		return err
	}
	request := ttdns.NewDescribeRecordListRequest()
	request.Domain = common.StringPtr(req.Domain)
	resp, err := cli.DescribeRecordList(request)
	for _, v := range resp.Response.RecordList {
		r := new(model.Record)
		tmp := *v
		r.Domain = req.Domain
		r.RR = *(tmp.Name)
		rid := int(*(tmp.RecordId))
		r.RecordId = strconv.Itoa(rid)
		if *(tmp.Status) == "ENABLE" {
			r.Status = 1
		} else if *(tmp.Status) == "DISABLE" {
			r.Status = 2
		} else {
			r.Status = 0
		}
		r.Account = req.Account
		if *(tmp.Line) == "联通" {
			r.Line = 2
		} else if *(tmp.Line) == "默认" {
			r.Line = 1
		} else {
			r.Line = 0
		}
		var recordType model.SysDictionaryDetail
		err = global.GVA_DB.Where("`sys_dictionary_id` = 11 and `label` = ?", *(tmp.Type)).First(&recordType).Error
		if err != nil {
			return err
		}
		r.Type = recordType.Value
		r.Value = *(tmp.Value)
		results := global.GVA_DB.Where("record_id = ? and value = ?", r.RecordId, r.Value).FirstOrCreate(r)
		if results.Error != nil {
			return results.Error
		}

	}
	return err
}
func (t TencentDnsProvider) FlushDomainsToDb(req model.DnsUser) (err error) {
	var cli *ttdns.Client
	credential := common.NewCredential(
		req.AccessKey,
		req.AccessSecret,
	)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "dnspod.tencentcloudapi.com"
	cli, err = ttdns.NewClient(credential, "", cpf)
	request := ttdns.NewDescribeDomainListRequest()
	response, err := cli.DescribeDomainList(request)
	if err != nil {
		return err
	}
	//var od model.DomainInfo
	for _,v := range response.Response.DomainList {
		d := new(model.DomainInfo)
		d.RecordTotal = int(*(v.RecordCount))
		d.Domain = *(v.Name)
		d.DomainId = strconv.Itoa(int(*(v.DomainId)))
		d.Account = req.Account
		d.DnsProvider = req.Platform

		results := global.GVA_DB.Where("domain = ?", d.Domain).First(&model.DomainInfo{})
		if results.Error != nil {
			if results.Error == gorm.ErrRecordNotFound {
				d.RegisterTime = time.Now()
				results = global.GVA_DB.Create(&d)
				if results.Error != nil {
					return results.Error
				}
			}
		} else {
			results = global.GVA_DB.Model(&model.DomainInfo{}).Where("domain = ?", d.Domain).Updates(d)
			if results.Error != nil {
				return results.Error
			} else {
				fmt.Println(results)
			}
		}
		d.ID = 0
	}
	return err
}