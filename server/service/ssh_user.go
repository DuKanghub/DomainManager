package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"gin-vue-admin/utils"
)

// 用户信息加密
func encodeUser(sshUser *model.SSHUser) error {
	if sshUser.Password != "" {
		src := sshUser.Password
		dst, err := utils.EncodeStr([]byte(src))
		if err != nil {
			return err
		}
		sshUser.Password = dst
	}
	if sshUser.PrivateKey != "" {
		src := sshUser.PrivateKey
		dst, err := utils.EncodeStr([]byte(src))
		if err != nil {
			return err
		}
		sshUser.PrivateKey = dst
	}
	return nil
}

// CreateSSHUser 创建SSHUser记录
// Author [piexlmax](https://github.com/piexlmax)
func CreateSSHUser(sshUser model.SSHUser) (err error) {
	// 将传入的密码和秘钥加密后再存入数据库
	err = encodeUser(&sshUser)
	if err != nil {
		return err
	}
	err = global.GVA_DB.Create(&sshUser).Error
	//if err != nil {
	//	return err
	//}
	////此时sshUser里的id就是数据库的id了。
	////每创建一个用户，将id和name写到SSHUser(id:13)字典详情
	//var status = new(bool)
	//*status = true
	//user := model.SysDictionaryDetail{
	//	Label: sshUser.Name,
	//	Value: int(sshUser.ID),
	//	Status: status,
	//	Sort: int(sshUser.ID),
	//	SysDictionaryID: 13,	//这个是SSHUser字典的id, 在`source/dictionary.go`里定义的。
	//}
	//err = global.GVA_DB.Create(&user).Error
	return err
}

// DeleteSSHUser 删除SSHUser记录
// Author [piexlmax](https://github.com/piexlmax)
func DeleteSSHUser(sshUser model.SSHUser) (err error) {
	err = global.GVA_DB.Delete(&sshUser).Error
	//if err != nil {
	//	return err
	//}
	////每删除一个用户，将删除SSHUser(id:13)字典详情
	//err = global.GVA_DB.Delete(&model.SysDictionaryDetail{}, "sys_dictionary_id = 13 and value = ?", int(sshUser.ID)).Error
	return err
}

// DeleteSSHUserByIds 批量删除SSHUser记录
// Author [piexlmax](https://github.com/piexlmax)
func DeleteSSHUserByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.SSHUser{}, "id in ?", ids.Ids).Error
	//if err != nil {
	//	return err
	//}
	//err = global.GVA_DB.Delete(&[]model.SysDictionaryDetail{}, "sys_dictionary_id = 13 and value in ?", ids.Ids).Error
	return err
}

// UpdateSSHUser 更新SSHUser记录
// Author [piexlmax](https://github.com/piexlmax)
func UpdateSSHUser(sshUser model.SSHUser) (err error) {
	// 将传入的密码和秘钥加密后再存入数据库
	err = encodeUser(&sshUser)
	if err != nil {
		return err
	}
	err = global.GVA_DB.Save(&sshUser).Error
	//if err != nil {
	//	return err
	//}
	////每更新一个用户，将更新SSHUser(id:13)字典详情
	//var user model.SysDictionaryDetail
	//err = global.GVA_DB.Where("sys_dictionary_id = 13 and value = ?", int(sshUser.ID)).Find(&user).Error
	//if err != nil {
	//	return err
	//}
	//if sshUser.Name != "" && sshUser.Name != user.Label {
	//	user.Label = sshUser.Name
	//	err = global.GVA_DB.Save(&user).Error
	//}
	return err
}

// GetSSHUser 根据id获取SSHUser记录
// Author [piexlmax](https://github.com/piexlmax)
func GetSSHUser(id uint) (err error, sshUser model.SSHUser) {
	err = global.GVA_DB.Where("id = ?", id).First(&sshUser).Error
	return
}

func GetSSHUserList() (err error, sshUsers []model.SSHUser) {
	err = global.GVA_DB.Select("id,name").Find(&sshUsers).Error
	// fmt.Println("users: ", sshUsers)
	return
}

// GetSSHUserInfoList 分页获取SSHUser记录
// Author [piexlmax](https://github.com/piexlmax)
func GetSSHUserInfoList(info request.SSHUserSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.SSHUser{})
	var sshUsers []model.SSHUser
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&sshUsers).Error
	return err, sshUsers, total
}
