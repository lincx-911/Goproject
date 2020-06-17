package models

import (
	"errors"
	"fmt"
	orm "miniprogram/database"
)

//User 用户
type User struct {
	UID      string `json:"uid"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
	School   string `json:"school"`
	Faceimg  string `json:"faceimg"`
	Sex      int    `json:"sex"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Credit   int    `json:"credit"`
	Property int    `json:"property"`
}

//InsertUser 添加用户
func InsertUser(user User) (err error) {
	err = orm.DB.Create(&user).Error
	return
}

//DeleteUser 删除user
func DeleteUser(id string) (result User, err error) {

	err = orm.DB.Where("uid=?", id).First(&result).Error
	if err != nil {
		return
	}
	if result.UID != id {
		err = errors.New("user do not exit")
		return
	}
	if err = orm.DB.Where(User{UID: id}).Delete(User{}).Error; err != nil {
		return
	}
	//err = orm.DB.Exec("ALTER TABLE users AUTO_INCREMENT = 1").Error
	return
}

//UpdateUser 更新user
func UpdateUser(user User) (result User, err error) {
	err = orm.DB.Model(User{}).Where("uid=?", user.UID).Update(user).Error
	if err == nil {
		result = user
	}
	return

}

//SelectUserbyPhone 通过电话查询用户
func SelectUserbyPhone(phone string) (result User, err error) {
	if err = orm.DB.Where("phone=?", phone).Find(&result).Error; err != nil {
		return
	}
	return
}

//SelectUserbyEmail 通过邮箱查询用户
func SelectUserbyEmail(email string) (result User, err error) {
	if err = orm.DB.Where("email=?", email).Find(&result).Error; err != nil {
		return
	}
	return
}

//SelectUserbyID 通过id查询用户
func SelectUserbyID(id string) (result User, err error) {
	err = orm.DB.Where("uid=?", id).First(&result).Error
	return
}

//GetallUsers 返回所有用户信息
func GetallUsers() (users []User, err error) {
	err = orm.DB.Find(&users).Error
	return
}

//GetallInssue 获取所有个人发布的的招募信息
func GetallInssue(uid string) (ucrs []Info, err error) {
	err = orm.DB.Where("uid=?", uid).Find(&ucrs).Error
	return
}

//ComfireTask 确定任务
func ComfireTask(rid int, ueid string) (result Info, err error) {
	err = orm.DB.Where("rid=?", ueid, rid).First(&result).Error
	fmt.Println("result:", result)
	if result.Rid == 0 {
		err = errors.New("info do not exit")
		return
	}
	err = orm.DB.Model(&result).Update("ueid", ueid).Error
	return
}

//DelTask 取消正在进行的任务
func DelTask(rid int, ueid string) (result Info, err error) {
	err = orm.DB.Where("ueid=? and rid=?", ueid, rid).First(&result).Error
	fmt.Println("result:", result)
	if result.Rid == 0 {
		err = errors.New("info do not exit")
		return
	}
	err = orm.DB.Model(&result).Update("ueid", "").Error
	return
}

//Uploaduserimage 上传头像
func Uploaduserimage(uid string, img string) (err error) {
	err = orm.DB.Model(User{}).Where("uid=?", uid).Update("faceimg", img).Error
	return
}

//Updatepassword 修改密码
func Updatepassword(uid string, password string) (err error) {
	err = orm.DB.Model(User{}).Where("uid=?", uid).Update("password", password).Error
	return
}

// //GetallUclloctcon 获取收藏所有赛事信息
// func GetallUclloctcon(uid int) (rcs []Reinfo, err error) {

// 	eff := orm.DB.Table("contestinfos").Select("*").Joins("left join uclloctcons on uclloctcons.cid=contestinfos.cid").Where("uclloctcons.uid=?", uid).Scan(&rcs).RowsAffected
// 	if eff == 0 {
// 		err = errors.New("sql error")
// 	}
// 	//err = orm.DB.Where("rid in ?",).Find(&ucrs).Error
// 	return
// }

// //InsertUclloctcon 添加收藏赛事信息
// func InsertUclloctcon(ucr Uclloctcon) (result Uclloctcon, err error) {
// 	err = orm.DB.Create(&ucr).Error
// 	if err != nil {
// 		//err = resutl.Error
// 		return
// 	}
// 	result = ucr
// 	return
// }

// //DelUclloctcon 删除收藏赛事信息
// func DelUclloctcon(re Uclloctcon) (result Uclloctcon, err error) {
// 	err = orm.DB.Where("uid=? and cid=?", re.UID, re.Cid).First(&result).Error
// 	fmt.Println("result:", result)
// 	if result.ID == 0 {
// 		return
// 	}
// 	if err = orm.DB.Delete(&result).Error; err != nil {
// 		return
// 	}
// 	err = orm.DB.Exec("ALTER TABLE uclloctcons AUTO_INCREMENT = 1").Error
// 	return
// }

// //GetallUrecruit 获取所有个人发布的招募信息
// func GetallUrecruit(uid int) (ucrs []Reinfo, err error) {
// 	eff := orm.DB.Table("reinfos").Select("*").Joins("left join urecruits on urecruits.rid=reinfos.rid").Where("urecruits.uid=?", uid).Scan(&ucrs).RowsAffected
// 	if eff == 0 {
// 		err = errors.New("sql error")
// 	}
// 	return
// }

// //InsertUrecruit 添加个人发布的招募信息
// func InsertUrecruit(ucr Urecruit) (result Urecruit, err error) {
// 	resutl := orm.DB.Create(&ucr)
// 	if resutl.Error != nil {
// 		err = resutl.Error
// 		return
// 	}
// 	return
// }

// //DelUrecruit 删除个人发布的招募信息记录
// func DelUrecruit(id int) (result Urecruit, err error) {
// 	err = orm.DB.Where("id=?", id).First(&result).Error
// 	fmt.Println("result:", result)
// 	if result.ID == 0 {
// 		return
// 	}
// 	if err = orm.DB.Where(Urecruit{ID: id}).Delete(Urecruit{}).Error; err != nil {
// 		return
// 	}
// 	err = orm.DB.Exec("ALTER TABLE users AUTO_INCREMENT = 1").Error
// 	return
// }

// //GetallUrecord 获取所有参赛记录
// func GetallUrecord(uid int) (ucrs []Contestinfo, err error) {
// 	// err = orm.DB.Find(&ucrs).Error
// 	// return
// 	var r Contestinfo
// 	rows, err := orm.DB.Exec("select contestinfos.cid,contestinfos.cinfo,contestinfos.ctag,contestinfos.publicer,contestinfos.cimg,contestinfos.cdate contestinfos.cstatus from contestinfos,urecords where contestinfos.cid=urecords.cid and urecords.uid=?", uid).Rows()
// 	if err != nil {
// 		return
// 	}
// 	for rows.Next() {
// 		err = rows.Scan(&r)
// 		if err != nil {
// 			return
// 		}
// 		ucrs = append(ucrs, r)
// 	}
// 	return
// }

// //InsertUrecord 添加参赛记录
// func InsertUrecord(ucr Urecord) (result Urecord, err error) {
// 	resutl := orm.DB.Create(&ucr)
// 	if resutl.Error != nil {
// 		err = resutl.Error
// 		return
// 	}
// 	return
// }

// //DelUrecord 删除参赛记录
// func DelUrecord(id int) (result Urecord, err error) {
// 	err = orm.DB.Where("id=?", id).First(&result).Error
// 	fmt.Println("result:", result)
// 	if result.ID == 0 {
// 		return
// 	}
// 	if err = orm.DB.Where(Urecord{ID: id}).Delete(Urecord{}).Error; err != nil {
// 		return
// 	}
// 	err = orm.DB.Exec("ALTER TABLE users AUTO_INCREMENT = 1").Error
// 	return
// }
