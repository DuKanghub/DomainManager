package ali

import (
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
	"gorm.io/gorm"
	"strconv"
	"time"
)
var (
	pageN = 0
	count int
)
func GetAccount(req *model.Record) (client *alidns.Client, err error) {
	var dnsUser model.DnsUser
	err = global.GVA_DB.Where("account = ?", req.Account).First(&dnsUser).Error
	fmt.Printf("%+v\n", dnsUser)
	if err != nil {
		return nil, err
	}
	client, err = alidns.NewClientWithAccessKey("cn-hangzhou", dnsUser.AccessKey, dnsUser.AccessSecret)
	return
}

func AddRecord(req *model.Record) (err error) {
	var cli *alidns.Client
	cli, err = GetAccount(req)
	if err != nil {
		return err
	}
	request := alidns.CreateAddDomainRecordRequest()
	request.Scheme = "https"
	request.Lang = "zh"
	request.Value = req.Value
	request.RR = req.RR
	request.DomainName = req.Domain
	var recordType model.SysDictionaryDetail
	err = global.GVA_DB.Where("`sys_dictionary_id` = 11 and `value` = ?", req.Type).First(&recordType).Error
	if err != nil {
		return err
	}
	request.Type = recordType.Label

	//var recordLine model.SysDictionaryDetail
	//err = global.GVA_DB.Where("`sys_dictionary_id` = 10 and `value` = ?", req.Line).First(&recordLine).Error
	//if err != nil {
	//	return err
	//}
	// 线路这里采用简单写法，因为目前自己环境只有联通和默认线路
	if req.Line == 2 {
		request.Line = "unicom"
	} else {
		request.Line = "default"
	}

	resp, err := cli.AddDomainRecord(request)
	if err != nil {
		return err
	}
	fmt.Printf("%+v\n", resp)
	req.RecordId = resp.RecordId
	return
}

func DelRecord(req *model.Record) (err error) {
	var cli *alidns.Client
	cli, err = GetAccount(req)
	if err != nil {
		return err
	}
	request := alidns.CreateDeleteDomainRecordRequest()
	request.Scheme = "https"
	request.Lang = "zh"
	request.RecordId = req.RecordId
	_, err = cli.DeleteDomainRecord(request)

	return err
}

func UpdateRecord(req *model.Record) (err error) {
	var cli *alidns.Client
	cli, err = GetAccount(req)
	if err != nil {
		return err
	}
	request := alidns.CreateUpdateDomainRecordRequest()
	request.Scheme = "https"
	request.Lang = "zh"
	request.RecordId = req.RecordId
	request.RR = req.RR
	request.Value = req.Value
	var recordType model.SysDictionaryDetail
	err = global.GVA_DB.Where("`sys_dictionary_id` = 11 and `value` = ?", req.Type).First(&recordType).Error
	if err != nil {
		return err
	}
	request.Type = recordType.Label

	//var recordLine model.SysDictionaryDetail
	//err = global.GVA_DB.Where("`sys_dictionary_id` = 10 and `value` = ?", req.Line).First(&recordLine).Error
	//if err != nil {
	//	return err
	//}
	// 线路这里采用简单写法，因为目前自己环境只有联通和默认线路
	if req.Line == 2 {
		request.Line = "unicom"
	} else {
		request.Line = "default"
	}
	resp, err := cli.UpdateDomainRecord(request)
	if err != nil {
		return err
	}
	req.RecordId = resp.RecordId
	return nil
}

func FlushRecordsToDb(req *model.Record) (err error) {
	var cli *alidns.Client
	cli, err = GetAccount(req)
	if err != nil {
		return err
	}
	request := alidns.CreateDescribeDomainRecordsRequest()
	request.Scheme = "https"
	request.Lang = "zh"
	request.PageSize = "50"
	if pageN == 0 {
		pageN = 1
	}
	request.PageNumber = requests.Integer(strconv.Itoa(pageN))
	request.DomainName = req.Domain
	resp, err := cli.DescribeDomainRecords(request)
	if err != nil {
		return err
	}
	//fmt.Printf("service: ali.FlushRecordsToDb: %+v\n", resp)
	var total = int(resp.TotalCount)
	var domain model.DomainInfo
	err = global.GVA_DB.Where("domain = ?", req.Domain).First(&domain).Error
	if err != nil {
		return err
	}
	//fmt.Printf("domain: %s, RecordTotal: %d, total: %d \n", req.DomainInfo, domain.RecordTotal, total)
	if domain.RecordTotal != total {
		domain.RecordTotal = total
		_ = global.GVA_DB.Where("domain = ?", req.Domain).Save(&domain).Error
	}

	for _,v := range resp.DomainRecords.Record {
		r := new(model.Record)
		r.Domain = v.DomainName
		r.RR = v.RR
		var recordType model.SysDictionaryDetail
		err = global.GVA_DB.Where("`sys_dictionary_id` = 11 and `label` = ?", v.Type).First(&recordType).Error
		if err != nil {
			return err
		}
		r.Type = recordType.Value
		// 线路这里采用简单写法，因为目前自己环境只有联通和默认线路
		if v.Line == "unicom" {
			r.Line = 2
		} else {
			r.Line = 1
		}
		r.Value = v.Value
		r.RecordId = v.RecordId
		// 这里也是简单写法，不读数据库了，ENABLE=1, DISABLE=2
		if v.Status == "DISABLE" {
			r.Status = 2
		} else {
			r.Status = 1
		}
		r.Account = req.Account
		results := global.GVA_DB.Where("record_id = ? and value = ?", v.RecordId, v.Value).FirstOrCreate(r)
		if results.Error != nil {
			return results.Error
		}
		//fmt.Printf("domain: %s, %+v\n", req.DomainInfo, r)
	}
	count += len(resp.DomainRecords.Record)
	if count < total {
		pageN += 1
		_ = FlushRecordsToDb(req)
	}
	return err
}
func FlushDomainsToDb(req model.DnsUser) (err error) {
	var cli *alidns.Client
	cli, err = alidns.NewClientWithAccessKey("cn-hangzhou", req.AccessKey, req.AccessSecret)
	request := alidns.CreateDescribeDomainsRequest()
	request.Scheme = "https"
	request.Lang = "zh"
	request.PageSize = "100"
	if pageN == 0 {
		pageN = 1
	}
	request.PageNumber = requests.Integer(strconv.Itoa(pageN))

	resp, err := cli.DescribeDomains(request)
	if err != nil {
		return err
	}
	var total  = int(resp.TotalCount)
	//var page  = int(resp.PageNumber)
	var t1 = len(resp.Domains.Domain)
	//var d model.DomainInfo
	for _, v := range resp.Domains.Domain {
		d := new(model.DomainInfo)
		d.RecordTotal = int(v.RecordCount)
		d.RegisterTime = time.Unix(v.CreateTimestamp/1000,0)
		d.DomainId = v.DomainId
		d.Account = req.Account
		d.Domain = v.DomainName
		req1 := alidns.CreateDescribeDomainNsRequest()
		req1.Scheme = "https"
		req1.DomainName = v.DomainName
		req1.Lang = "zh"
		resp1, err := cli.DescribeDomainNs(req1)
		if err != nil {
			return err
		} else {
			if resp1.AllAliDns {
				d.DnsProvider = 1
			} else {
				d.DnsProvider = 2
			}
		}
		if d.DnsProvider == 1 {
			//fmt.Printf("%+v\n", d)
			od := new(model.DomainInfo)
			results := global.GVA_DB.Where("domain = ?", d.Domain).First(od)
			if results.Error != nil {
				if results.Error == gorm.ErrRecordNotFound {
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
					//fmt.Println(results)
				}
			}

		}
		// 操作完后，重置id为0，否则第一条记录后会报ID重复，
		// 因为前面传入的d是指针，且第一次被自动赋值，后面又没被赋值，所以会继续使用这个ID，
		// 导致报INSERT: [1062] Duplicate entry '数字' for key 'PRIMARY'
		d.ID = 0
		count += t1
		if count < total {
			pageN += 1
			_ = FlushDomainsToDb(req)
		}
	}
	return err

}