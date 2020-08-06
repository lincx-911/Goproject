package models

import(
	orm "miniprogram/database"
)


//Admin 管理员
type Admin struct{
	Aid int `json:"id" gorm:"primary_key"`//列名为aid
	Aname string `json:"aname"`//列名为aname
	Password string `json:"password"`//
}

//InsertAdmin 添加admin
func (admin Admin)InsertAdmin()(id int,err error)  {
	resutl:=orm.DB.Create(&admin)
	id = admin.Aid
	if resutl.Error!=nil{
		err = resutl.Error
		return
	}
	return
}
//DeleteAdmin 删除admin
func (admin *Admin)DeleteAdmin(id int)(result Admin,err error){
	if err = orm.DB.Select([]string{"id"}).First(&admin, id).Error; err != nil {
		return
	}
	if err = orm.DB.Delete(&admin).Error;err!=nil{
		return 
	}
	result = *admin
	return
}

//UpdateAdmin 更新admin
func (admin *Admin)UpdateAdmin(id int)(updateAdmin Admin,err error){
	if err = orm.DB.Select([]string{"id", "username"}).First(&updateAdmin, id).Error; err != nil {
		return
	}

	//参数1:是要修改的数据
	//参数2:是修改的数据
	if err = orm.DB.Model(&updateAdmin).Updates(&admin).Error; err != nil {
		return
	}
	return

}