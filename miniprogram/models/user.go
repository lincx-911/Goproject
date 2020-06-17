package models

import (
	"errors"
	"fmt"
	orm "miniprogram/database"
)

//User 用户
type User struct {
	UID       int    `json:"uid"`
	Uname     string `json:"uname"`
	Unickname string `json:"unickname"`
	Password  string `json:"password"`
	Uemail    string `json:"uemail"`
	Uschool   string `json:"uschool"`
}

//Uclloctre 个人收藏的招募信息
type Uclloctre struct {
	ID  int `json:"id"`
	UID int `json:"uid"`
	RID int `json:"rid"`
}

//Uclloctcon 个人收藏的赛事信息
type Uclloctcon struct {
	ID  int `json:"id"`
	UID int `json:"uid"`
	Cid int `json:"cid"`
}

//Urecruit 个人发布历史
type Urecruit struct {
	ID  int `json:"id"`
	UID int `json:"uid"`
	Rid int `json:"rid"`
}

//Urecord 参赛记录
type Urecord struct {
	ID  int `json:"id"`
	UID int `json:"uid"`
	Cid int `json:"cid"`
}

//InsertUser 添加用户
func InsertUser(user User) (id int, err error) {
	resutl := orm.DB.Create(&user)
	id = user.UID
	if resutl.Error != nil {
		err = resutl.Error
		return
	}
	return
}

//DeleteUser 删除user
func DeleteUser(id int) (result User, err error) {

	err = orm.DB.Where("uid=?", id).First(&result).Error
	fmt.Println("result:", result)
	if result.UID == 0 {
		return
	}
	if err = orm.DB.Where(User{UID: id}).Delete(User{}).Error; err != nil {
		return
	}
	err = orm.DB.Exec("ALTER TABLE users AUTO_INCREMENT = 1").Error
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

//SelectUserbynickName 通过名字查询用户
func (user *User) SelectUserbynickName() (result User, err error) {
	if err = orm.DB.Where("unickname=? AND password=?", user.Unickname, user.Password).Find(&result).Error; err != nil {
		return
	}
	return
}

//SelectUserbyemail 通过名字查询用户
func (user *User) SelectUserbyemail() (result User, err error) {
	if err = orm.DB.Where("uemail=? AND password=?", user.Uemail, user.Password).Find(&result).Error; err != nil {
		return
	}
	return
}

//SelectUserbyID 通过id查询用户
func SelectUserbyID(id int) (result User, err error) {
	err = orm.DB.Where("uid=?", id).First(&result).Error
	return
}

//GetallUsers 返回所有用户信息
func GetallUsers() (users []User, err error) {
	err = orm.DB.Find(&users).Error
	return
}

//GetallUclloctre 获取所有收藏的招募信息
func GetallUclloctre(uid int) (ucrs []Reinfo, err error) {
	eff := orm.DB.Table("reinfos").Select("*").Joins("left join urecruits on urecruits.rid=reinfos.rid").Where("urecruits.uid=?", uid).Scan(&ucrs).RowsAffected
	if eff == 0 {
		err = errors.New("sql error")
	}
	//err = orm.DB.Where("rid in ?",).Find(&ucrs).Error
	return

}

//InsertUclloctre 添加收藏招募信息
func InsertUclloctre(ucr Uclloctre) (result Uclloctre, err error) {
	resutl := orm.DB.Create(&ucr)
	if resutl.Error != nil {
		err = resutl.Error
		return
	}
	return
}

//DelUclloctre 删除收藏招募信息
func DelUclloctre(re Uclloctre) (result Uclloctre, err error) {
	err = orm.DB.Where("uid=? and rid=?", re.UID, re.RID).First(&result).Error
	fmt.Println("result:", result)
	if result.ID == 0 {
		return
	}
	if err = orm.DB.Delete(&result).Error; err != nil {
		return
	}
	err = orm.DB.Exec("ALTER TABLE uclloctres AUTO_INCREMENT = 1").Error
	return
}

//GetallUclloctcon 获取收藏所有赛事信息
func GetallUclloctcon(uid int) (rcs []Reinfo, err error) {

	eff := orm.DB.Table("contestinfos").Select("*").Joins("left join uclloctcons on uclloctcons.cid=contestinfos.cid").Where("uclloctcons.uid=?", uid).Scan(&rcs).RowsAffected
	if eff == 0 {
		err = errors.New("sql error")
	}
	//err = orm.DB.Where("rid in ?",).Find(&ucrs).Error
	return
}

//InsertUclloctcon 添加收藏赛事信息
func InsertUclloctcon(ucr Uclloctcon) (result Uclloctcon, err error) {
	err = orm.DB.Create(&ucr).Error
	if err != nil {
		//err = resutl.Error
		return
	}
	result = ucr
	return
}

//DelUclloctcon 删除收藏赛事信息
func DelUclloctcon(re Uclloctcon) (result Uclloctcon, err error) {
	err = orm.DB.Where("uid=? and cid=?", re.UID, re.Cid).First(&result).Error
	fmt.Println("result:", result)
	if result.ID == 0 {
		return
	}
	if err = orm.DB.Delete(&result).Error; err != nil {
		return
	}
	err = orm.DB.Exec("ALTER TABLE uclloctcons AUTO_INCREMENT = 1").Error
	return
}

//GetallUrecruit 获取所有个人发布的招募信息
func GetallUrecruit(uid int) (ucrs []Reinfo, err error) {
	eff := orm.DB.Table("reinfos").Select("*").Joins("left join urecruits on urecruits.rid=reinfos.rid").Where("urecruits.uid=?", uid).Scan(&ucrs).RowsAffected
	if eff == 0 {
		err = errors.New("sql error")
	}
	return
}

//InsertUrecruit 添加个人发布的招募信息
func InsertUrecruit(ucr Urecruit) (result Urecruit, err error) {
	resutl := orm.DB.Create(&ucr)
	if resutl.Error != nil {
		err = resutl.Error
		return
	}
	return
}

//DelUrecruit 删除个人发布的招募信息记录
func DelUrecruit(id int) (result Urecruit, err error) {
	err = orm.DB.Where("id=?", id).First(&result).Error
	fmt.Println("result:", result)
	if result.ID == 0 {
		return
	}
	if err = orm.DB.Where(Urecruit{ID: id}).Delete(Urecruit{}).Error; err != nil {
		return
	}
	err = orm.DB.Exec("ALTER TABLE users AUTO_INCREMENT = 1").Error
	return
}

//GetallUrecord 获取所有参赛记录
func GetallUrecord(uid int) (ucrs []Contestinfo, err error) {
	// err = orm.DB.Find(&ucrs).Error
	// return
	var r Contestinfo
	rows, err := orm.DB.Exec("select contestinfos.cid,contestinfos.cinfo,contestinfos.ctag,contestinfos.publicer,contestinfos.cimg,contestinfos.cdate contestinfos.cstatus from contestinfos,urecords where contestinfos.cid=urecords.cid and urecords.uid=?", uid).Rows()
	if err != nil {
		return
	}
	for rows.Next() {
		err = rows.Scan(&r)
		if err != nil {
			return
		}
		ucrs = append(ucrs, r)
	}
	return
}

//InsertUrecord 添加参赛记录
func InsertUrecord(ucr Urecord) (result Urecord, err error) {
	resutl := orm.DB.Create(&ucr)
	if resutl.Error != nil {
		err = resutl.Error
		return
	}
	return
}

//DelUrecord 删除参赛记录
func DelUrecord(id int) (result Urecord, err error) {
	err = orm.DB.Where("id=?", id).First(&result).Error
	fmt.Println("result:", result)
	if result.ID == 0 {
		return
	}
	if err = orm.DB.Where(Urecord{ID: id}).Delete(Urecord{}).Error; err != nil {
		return
	}
	err = orm.DB.Exec("ALTER TABLE users AUTO_INCREMENT = 1").Error
	return
}

//Uploaduserimage 上传头像
func Uploaduserimage(uid int,img string)(err error){
	err=orm.DB.Model(User{}).Where("uid=?",uid).Update("uimg",img).Error
	return
}