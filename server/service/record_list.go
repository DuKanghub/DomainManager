package service

import (
	"errors"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"strings"
)

// CreateRecord 创建Record记录
// Author [dukanghub](https://github.com/DuKanghub)
func CreateRecord(record model.Record) (err error) {
	if record.Domain == "" {
		err = errors.New("域名不能为空")
		return err
	}
	if record.RR == "" {
		err = errors.New("主机记录不能为空")
		return err
	}
	if record.Type == 0 {
		err = errors.New("记录类型不能为空")
		return err
	}
	if record.Line == 0 {
		record.Line = 1
	}
	if record.Value == "" {
		err = errors.New("记录值不能为空")
		return err
	}
	if record.Status == 0 {
		record.Status = 1
	}
	var domain model.DomainInfo
	err = global.GVA_DB.Where("domain = ?", record.Domain).First(&domain).Error
	if err != nil {
		return err
	}
	if record.Account == 0 {
		record.Account = domain.Account
	}
	provider := NewDnsProvider(model.DnsUser{Platform: domain.DnsProvider})
	if provider != nil {
		err = provider.AddRecord(&record)
	} else {
		return errors.New("未知的provider")
	}
	if err != nil {
		return err
	}
	err = global.GVA_DB.Create(&record).Error
	if err == nil {
		// 更新domain_lists里的域名解析总数
		domain.RecordTotal += 1
		err = global.GVA_DB.Updates(&domain).Error
	}
	return err
}

// DeleteRecord 删除Record记录
// Author [dukanghub](https://github.com/DuKanghub)
func DeleteRecord(record model.Record) (err error) {
	var domain model.DomainInfo
	var rec model.Record
	_ = global.GVA_DB.Where("id = ?", record.ID).First(&rec).Error
	err = global.GVA_DB.Where("domain = ?", rec.Domain).First(&domain).Error
	if err != nil {
		return err
	}
	provider := NewDnsProvider(model.DnsUser{Platform: domain.DnsProvider})
	if provider != nil {
		err = provider.DelRecord(&rec)
	} else {
		return errors.New("未知的provider")
	}
	if err != nil {
		return err
	}
	err = global.GVA_DB.Delete(&record).Error
	if err == nil {
		// 更新domain_lists里的域名解析总数
		domain.RecordTotal -= 1
		err = global.GVA_DB.Updates(&domain).Error
	}
	return err
}

// DeleteRecordByIds 批量删除Record记录
// Author [dukanghub](https://github.com/DuKanghub)
func DeleteRecordByIds(ids request.IdsReq) (err error) {
	//var domain model.DomainInfo
	var rec []model.Record
	_ = global.GVA_DB.Where("id in ?", ids.Ids).Find(&rec).Error
	for _,v := range rec {
		err = DeleteRecord(v)
		if err != nil {
			return err
		}
	}
	//err = global.GVA_DB.Delete(&[]model.Record{},"id in ?",ids.Ids).Error
	return err
}

// UpdateRecord 更新Record记录
// Author [dukanghub](https://github.com/DuKanghub)
func UpdateRecord(record model.Record) (err error) {
	var domain model.DomainInfo
	err = global.GVA_DB.Where("domain = ?", record.Domain).First(&domain).Error
	if err != nil {
		return err
	}
	provider := NewDnsProvider(model.DnsUser{Platform: domain.DnsProvider})
	if provider != nil {
		err = provider.UpdateRecord(&record)
	} else {
		return errors.New("未知的provider")
	}
	if err != nil {
		return err
	}
	err = global.GVA_DB.Updates(&record).Error
	return err
}

// UpdateRecordMulti 批量更新Record记录
// Author [dukanghub](https://github.com/DuKanghub)
func UpdateRecordMulti(records model.Records) (err error) {
	for _,v := range records.Records {
		err = UpdateRecord(v)
		if err != nil {
			return err
		}
	}
	return
}


// GetRecord 根据id获取Record记录
// Author [dukanghub](https://github.com/DuKanghub)
func GetRecord(id uint) (err error, record model.Record) {
	err = global.GVA_DB.Where("id = ?", id).First(&record).Error
	return
}

// FlushRecordsToDb 从线上获取最新域名解析并更新到数据库
// Author [dukanghub](https://github.com/DuKanghub)
func FlushRecordsToDb() (err error) {
	var domains []model.DomainInfo
	err = global.GVA_DB.Order("record_total desc").Find(&domains).Error
	if err != nil {
		return err
	}
	for _, v := range domains {
		d := model.Record{
			Domain: v.Domain,
			Account: v.Account,
		}
		provider := NewDnsProvider(model.DnsUser{Platform: v.DnsProvider})
		if provider != nil {
			err = provider.FlushRecordsToDb(&d)
		} else {
			global.GVA_LOG.Error("未匹配到DnsProvider")
		}
		if err != nil {
			return err
		}
	}
	return err
}

// GetRecordInfoList 分页获取Record记录
// Author [dukanghub](https://github.com/DuKanghub)
func GetRecordInfoList(info request.RecordListSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	//fmt.Println("info.DnsProvider", info.DnsProvider)
	// 创建db
	db := global.GVA_DB.Model(&model.Record{})
	var recordList []model.Record
	// 如果有条件搜索 下方会自动创建搜索语句

	if strings.Contains(info.Domain, ",") {
		domains := strings.Split(info.Domain, ",")
		db = db.Where("`domain` in (?)", domains)
	} else if info.Domain != "" {
		db = db.Where("`domain` = ?",info.Domain)
	}

	if info.Type != 0 {
		db = db.Where("`type` = ?",info.Type)
	}
	if info.Line != 0 {
		db = db.Where("`line` = ?",info.Line)
	}
	if strings.Contains(info.RR, ",") {
		rrs := strings.Split(info.RR, ",")
		db = db.Where("`rr` in (?)", rrs)
	} else if info.RR != "" {
		db = db.Where("`rr` = ?",info.RR)
	}
	if strings.Contains(info.Value, ",") {
		values := strings.Split(info.Value, ",")
		db = db.Where("`value` in (?)", values)
	} else if info.Value != "" {
		db = db.Where("`value` = ?",info.Value)
	}

	if info.Status != 0 {
		db = db.Where("`status` = ?",info.Status)
	}
	if info.Account != 0 {
		db = db.Where("`account` = ?",info.Account)
	}
	db = db.Order("updated_at desc").Order("domain desc").Order("type desc")

	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&recordList).Error
	return err, recordList, total
}
