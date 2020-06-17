package models

import (
	"fmt"
	orm "miniprogram/database"
)

//Reinfo 组队招募信息
type Reinfo struct {
	Rid       int       `json:"rid" gorm:"primary_key"`
	Rinfo     string    `json:"rinfo"`
	Rtag      string    `json:"rtag"`
	Rpublicer string    `json:"rpublicer"`
	Rimg      string    `json:"rimg"`
	Rdate     LocalTime `json:"rdate" db:"rdate"`
	Rstatus   int       `json:"rstatus"`
}
//InsertReinfo 添加招募信息
func InsertReinfo(reinfo Reinfo, uid int) (id int, err error) {
	tx := orm.DB.Begin()
	//tx.Exec("INSERT reinfos SET rinfo=?,rtag=?,rpublicer=?,rimg=?,rdate=?,rstatus=?",reinfo.Rinfo,reinfo.Rtag,reinfo.Rpublicer,reinfo.Rimg,reinfo.Rdate,reinfo.Rstatus)
	roweffect := tx.Create(&reinfo).RowsAffected
	fmt.Println("rid:", reinfo.Rid)
	if roweffect == 0 {
		fmt.Println("错了吗")
		err = tx.Error
		tx.Rollback()
		return
	}
	id = reinfo.Rid
	var urc Urecruit
	urc.Rid = reinfo.Rid
	urc.UID = uid
	roweffect = tx.Create(&urc).RowsAffected
	if roweffect == 0 {
		fmt.Println("错了吗")
		err = tx.Error
		tx.Rollback()
		return
	}
	tx.Commit()
	return
}

//DeleteReinfo 删除招募信息
func DeleteReinfo(id int) (result Reinfo, err error) {
	tx := orm.DB.Begin()
	// err=tx.Where("rid=?",id).First(&result).Error
	// if err != nil {
	// 	tx.Rollback()
	// 	return
	// }
	err = tx.Exec("DElETE FROM reinfos WHERE rid=?", id).Error
	// err = orm.DB.Exec("DElETE FROM reinfos WHERE rid=?", id).Error
	//fmt.Println("rid:",result.Rid)

	if err != nil {
		tx.Rollback()
		return
	}
	err = tx.Exec("DELETE FROM urecruits Where rid=?", id).Error
	// if err = orm.DB.Where(Reinfo{Rid: id}).Delete(Reinfo{}).Error; err != nil {
	// 	return
	// }
	if err != nil {
		tx.Rollback()
		return
	}
	err = tx.Exec("ALTER TABLE reinfos AUTO_INCREMENT = 1").Error
	if err != nil {
		tx.Rollback()
		return
	}
	tx.Commit()
	return
}

//UpdateReinfo 更新招募信息
func UpdateReinfo(reinfo Reinfo) (result Reinfo, err error) {
	fmt.Println("reinfo:", reinfo)
	err = orm.DB.Model(Reinfo{}).Updates(&reinfo).Error
	//err=orm.DB.Exec("UPDATE reinfos SET reinfo=?,rtag=?,rpublicer=?,rimg=?,rdate=?,rstatus=? WHERE rid=?",reinfo.Rinfo,reinfo.Rtag,reinfo.Rpublicer,reinfo.Rimg,reinfo.Rdate,reinfo.Rstatus,reinfo.Rid).Error
	if err == nil {
		result = reinfo
	}
	return

}

//SelectReinfoByTag 通过标签查询招募信息
func SelectReinfoByTag(tag string) (result []Reinfo, err error) {
	if err = orm.DB.Where("rtag=?", tag).Find(&result).Error; err != nil {
		return
	}
	return
}

//SelectReinfoBypublicer 通过发布组织查赛事
func SelectReinfoBypublicer(publicer string) (result []Reinfo, err error) {
	if err = orm.DB.Where("rpublicer=?", publicer).Find(&result).Error; err != nil {
		return
	}
	return
}

//SelectReinfobyID 通过id查询赛事
func SelectReinfobyID(id int) (result Reinfo, err error) {
	fmt.Println("selectid", id)
	err=orm.DB.First(&result, id).Error
	return
}

//GetallReinfos 返回所有招募信息
func GetallReinfos() (result []Reinfo, err error) {
	err = orm.DB.Find(&result).Error
	return
}
