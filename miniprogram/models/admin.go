package models

import (
	orm "miniprogram/database"
)

//Admin 管理员
type Admin struct {
	Aid      int    `json:"aid"`      //列名为aid
	Aname    string `json:"aname"`    //列名为aname
	Password string `json:"password"` //列名为password
}

//InsertAdmin 添加admin
func InsertAdmin(admin Admin) (id int, err error) {
	//orm.DB.Exec("INSERT admins SET aname=?,password=?",admin.Aname,admin)
	err = orm.DB.Create(&admin).Error
	if err == nil {
		id = admin.Aid
		return
	}
	return
}

//DeleteAdmin 删除admin
func DeleteAdmin(id int) (result Admin, err error) {
	
	err=orm.DB.Where("aid=?",id).First(&result).Error
	if result.Aid==0{
		return
	}
	if err = orm.DB.Where(Admin{Aid:id}).Delete(Admin{}).Error; err != nil {
		return
	}
	err=orm.DB.Exec("ALTER TABLE admins AUTO_INCREMENT = 1").Error
	return
}

//UpdateAdmin 更新admin
func UpdateAdmin(admin Admin) (updateAdmin Admin, err error) {
	err=orm.DB.Model(Admin{}).Where("aid=?",admin.Aid).Update(admin).Error
	if err==nil{
		updateAdmin = admin
	}
	return

}

//SelectAdminbyName 通过名字查询管理员
func (admin *Admin) SelectAdminbyName() (result Admin, err error) {
	if err = orm.DB.Where("aname=? AND password=?", admin.Aname, admin.Password).Find(&result).Error; err != nil {
		return
	}
	return
}
//SelectAdminbyID 通过id查询管理员
func SelectAdminbyID(id int)(result Admin,err error){
	err=orm.DB.Where("aid=?",id).First(&result).Error
	return
}
//GetallAdmins 返回
func GetallAdmins()  (admins []Admin,err error){
	err = orm.DB.Find(&admins).Error
	return
}
