package dao

import (
	"forum/model"
	orm "forum/respository"

	"github.com/jinzhu/gorm"
)

//GetUserByID 通过id查找用户
func GetUserByID(id int) (*model.User,error) {
	user := &model.User{}
	find := orm.DB.First(&user, id)
	if find.RowsAffected == 0 {
		return user, nil
	}
	err:=orm.DB.Model(&user).Related(&user.Utag,"utag").Error
	return user, err
}
//UpdateUser 更新user
func UpdateUser(user *model.User)(bool,error){
	tx := orm.DB.Begin()
	var tags []model.Utag
	utags:= user.Utag
	for i := range utags{
		tag:=model.Utag{}
		if tx.Where("utagname=?",utags[i].Utagname).First(&tag).RecordNotFound(){
			tag = utags[i]
			if err:=tx.Create(&tag).Error;err!=nil{
				tx.Rollback()
				return false,err
			}
			
		}
		tags = append(tags, tag)
	}
	user.Utag=tags
	err:=tx.Model(&user).Association("utag").Replace(user.Utag).Error
	if err!=nil{
		tx.Rollback()
		return false,err
	}
	err =tx.Model(&user).Update(&user).Error
	if err!=nil{
		tx.Rollback()
		return false,err
	}
	tx.Commit()
	return true,nil
}
//DeleteUserByID 通过id删除用户
func DeleteUserByID(id int)(bool,error)  {
	tx:=orm.DB.Begin()
	var user model.User
	if tx.First(&user,id).RecordNotFound(){
		tx.Rollback()
		return false,gorm.ErrRecordNotFound
	}
	if err:=tx.Model(&user).Association("utag").Clear().Error;err!=nil{
		tx.Rollback()
		return false,err
	}
	if err:=tx.Model(&user).Delete(&user).Error;err!=nil{
		tx.Rollback()
		return false,err
	}
	tx.Commit()
	return true,nil
}

//UpdateUserImg 更新头像
func UpdateUserImg(id int,img string)(err error){
	err=orm.DB.Model(model.User{}).Where("id=?",id).Update("avatar",img).Error
	return
}
//updatepsw 更新密码
func Updatepsw(email string,psw string)(err error){
	err=orm.DB.Model(model.User{}).Where("email=?",email).Update("password",psw).Error
	return
}

