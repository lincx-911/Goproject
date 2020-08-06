package models

import(
	orm "miniprogram/database"
)

//User 用户
type User struct{
	UID int `json:"uid" gorm:"primary_key"`
	Uname string `json:"uname"`
	Unickname string `json :"unickname"`
	Password string `json:"password"`
	Uemail string `json:"uemail"`
	Uschool string `json:"uemail"`
}
//InsertUser 添加用户
func (user User)InsertUser(id int,err error){
	resutl:=orm.DB.Create(&user)
	id = user.UID
	if resutl.Error!=nil{
		err = resutl.Error
		return
	}
	return
}

//DeleteUser 删除User
func (user *User)DeleteUser(id int)(result User,err error){
	if err = orm.DB.Select([]string{"id"}).First(&user, id).Error; err != nil {
		return
	}
	if err = orm.DB.Delete(&user).Error;err!=nil{
		return 
	}
	result = *user
	return
}

//UpdateUser 更新User
func (user *User)UpdateUser(id int)(updateUser User,err error){
	if err = orm.DB.Select([]string{"id", "username"}).First(&updateUser, id).Error; err != nil {
		return
	}

	//参数1:是要修改的数据
	//参数2:是修改的数据
	if err = orm.DB.Model(&updateUser).Updates(&user).Error; err != nil {
		return
	}
	return

}