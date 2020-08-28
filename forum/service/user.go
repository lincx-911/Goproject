package service

import (
	"forum/dao"
	"forum/model"
)

//GetUerByIDService 通过id获取user
func GetUerByIDService(id int) (*model.User,error) {
	return dao.GetUserByID(id)
}
//UpdateUserService 更新用户
func UpdateUserService(user *model.User) (bool,error) {
	return dao.UpdateUser(user)
}

//DeleteUserService 删除用户
func DeleteUserService(id int)(bool,error){
	return dao.DeleteUserByID(id)
}

//UpdateUserImgService 更新头像
func UpdateUserImgService(id int,img string)(bool,error){
	err:=dao.UpdateUserImg(id,img)
	if err!=nil{
		return false,err
	}
	return true,nil
}
//UpdatePassword 更新密码
func UpdatePassword (email string,psw string) (bool,error) {
	err:=dao.Updatepsw(email,psw)
	if err!=nil{
		return false,err
	}
	return true,nil
}