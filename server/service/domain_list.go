package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"gin-vue-admin/service/ali"
	"gin-vue-admin/service/dnspod"
	"strings"
)

// CreateDomain 创建Domain记录
// Author [dukanghub](https://github.com/Dukanghub)
func CreateDomain(domain model.DomainInfo) (err error) {
	err = global.GVA_DB.Create(&domain).Error
	return err
}

// DeleteDomain 删除Domain记录
// Author [dukanghub](https://github.com/Dukanghub)
func DeleteDomain(domain model.DomainInfo) (err error) {
	err = global.GVA_DB.Delete(&domain).Error
	return err
}

// DeleteDomainByIds 批量删除Domain记录
// Author [dukanghub](https://github.com/Dukanghub)
func DeleteDomainByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.DomainInfo{},"id in ?",ids.Ids).Error
	return err
}

// UpdateDomain 更新Domain记录
// Author [dukanghub](https://github.com/Dukanghub)
func UpdateDomain(domain model.DomainInfo) (err error) {
	err = global.GVA_DB.Save(&domain).Error
	return err
}

// GetDomainById 根据id获取Domain记录
// Author [dukanghub](https://github.com/Dukanghub)
func GetDomainById(id uint) (err error, domain model.DomainInfo) {
	err = global.GVA_DB.Where("id = ?", id).First(&domain).Error
	return
}

// GetDomainList 分页获取Domain记录
// Author [dukanghub](https://github.com/Dukanghub)
func GetDomainList(info request.DomainListSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	//fmt.Println("info.DnsProvider", info.DnsProvider)
    // 创建db
	db := global.GVA_DB.Model(&model.DomainInfo{})
    var domains []model.DomainInfo
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.RecordTotal != 0 {
        db = db.Where("`record_total` > ?",info.RecordTotal)
    }
    if info.DnsProvider != 0 {
        db = db.Where("`dns_provider` = ?",info.DnsProvider)
    }
	if info.Company != 0 {
		db = db.Where("`company` = ?",info.Company)
	}
    if info.Account != 0 {
        db = db.Where("`account` = ?",info.Account)
    }
	if strings.Contains(info.Domain, ",") {
		domainArr := strings.Split(info.Domain, ",")
		db = db.Where("`domain` in (?)", domainArr)
	} else if info.Domain != "" {
		db = db.Where("`domain` = ?",info.Domain)
	}
	db = db.Order("record_total desc").Order("updated_at desc").Order("company")
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&domains).Error
	return err, domains, total
}

// UpdateDomains 批量更新Domain记录
// Author [dukanghub](https://github.com/Dukanghub)
func UpdateDomains(domains model.Domains) (err error) {

	for _,v := range domains.Domains {

		results := global.GVA_DB.Where("domain_id = ?", v.DomainId).Updates(&v)
		if results.Error != nil {
			err = results.Error
		}
	}
	return err
}

// FlushDomainsToDb 更新账号下的域名列表到数据库
// Author [dukanghub](https://github.com/Dukanghub)
func FlushDomainsToDb() (err error) {
	var users []model.DnsUser
	err = global.GVA_DB.Find(&users).Error
	if err != nil {
		return err
	}
	for _,v := range users {
		if v.Platform == 1 {
			err = ali.FlushDomainsToDb(v)
			//go ali.FlushDomainsToDb(v)
			if err != nil {
				return err
			}
		} else if v.Platform == 2 {
			err = dnspod.FlushDomainsToDb(v)
			if err != nil {
				return err
			}
		}
	}
	return err
}