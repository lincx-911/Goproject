package models

import (
	"fmt"
	orm "miniprogram/database"
)

//Organize 组织
type Organize struct {
	Oid      int    `json:"oid"`
	Oname    string `json:"oname"`
	Password string `json:"password"`
	Oschool  string `json:"oschool"`
	Oemail   string `json:"oemail"`
	Oimg     string `json:"oimg"`
}

//InsertOrganize 添加Organize
func InsertOrganize(org Organize) (id int, err error) {
	resutl := orm.DB.Create(&org)
	id = org.Oid
	if resutl.Error != nil {
		err = resutl.Error
		return
	}
	return
}

//DeleteOrganize 删除Organize
func DeleteOrganize(id int) (result Organize, err error) {

	err = orm.DB.Where("oid=?", id).First(&result).Error
	fmt.Println("result:", result)
	if result.Oid == 0 {
		return
	}
	if err = orm.DB.Where("oid=?",id).Delete(Organize{}).Error; err != nil {
		return
	}
	err = orm.DB.Exec("ALTER TABLE organizes AUTO_INCREMENT = 1").Error
	return
}

//UpdateOrganize 更新Organize
func UpdateOrganize(org Organize) (result Organize, err error) {
	err = orm.DB.Model(Organize{}).Where("oid=?", org.Oid).Update(org).Error
	if err == nil {
		result = org
	}
	return

}

//SelectorgbyName 通过名字查组织
func SelectorgbyName(name string) (result Organize, err error) {
	if err = orm.DB.Where("oname=?", name).Find(&result).Error; err != nil {
		return
	}
	return
}

//Selectorgbyemail 通过email查询组织
func Selectorgbyemail(email string) (result Organize, err error) {
	if err = orm.DB.Where("oemail=?", email).Find(&result).Error; err != nil {
		return
	}
	return
}

//SelectorgbyID 通过id查询管理员
func SelectorgbyID(id int) (result Organize, err error) {
	err = orm.DB.Where("oid=?", id).First(&result).Error
	return
}

//GetallOrganizes 返回全部组织
func GetallOrganizes() (results []Organize, err error) {
	err = orm.DB.Find(&results).Error
	return
}
//Uploadcontestfile 上传赛事证明
func Uploadcontestfile(cid int,img string)(err error){
	err=orm.DB.Model(Contestinfo{}).Where("cid=?",cid).Update("cimg",img).Error
	return
}