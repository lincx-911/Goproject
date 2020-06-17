package models

import (
	"database/sql/driver"
	"errors"
	"fmt"
	orm "miniprogram/database"
	"time"
)

//TimeFormat 时间格式
const TimeFormat = "2006-01-02 15:04:05"

//LocalTime 本地时间
type LocalTime time.Time

//Info 组队招募信息
type Info struct {
	Rid    int       `json:"rid" gorm:"primary_key"`
	UID    string    `json:"uid"`
	Rtitle string    `json:"rtitle"`
	Info   string    `json:"info"`
	Tag    string    `json:"tag"`
	Ttot   string    `json:"ttot"`
	Status string    `json:"status"`
	Rdate  LocalTime `json:"rdate" db:"rdate"`
	Reid   string    `json:"reid"`
}

//UnmarshalJSON 解析时间json
func (t *LocalTime) UnmarshalJSON(data []byte) (err error) {
	// 空值不进行解析
	if len(data) == 2 {
		*t = LocalTime(time.Time{})
		return
	}
	// 指定解析的格式
	now, err := time.Parse(`"`+TimeFormat+`"`, string(data))
	*t = LocalTime(now)
	return
}

//MarshalJSON 时间转json
func (t LocalTime) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(TimeFormat)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, TimeFormat)
	b = append(b, '"')
	return b, nil
}

//Value 写入 mysql 时调用
func (t LocalTime) Value() (driver.Value, error) {
	// 0001-01-01 00:00:00 属于空值，遇到空值解析成 null 即可
	if t.String() == "0001-01-01 00:00:00" {
		return nil, nil
	}
	return []byte(time.Time(t).Format(TimeFormat)), nil
}

//Scan 检出 mysql 时调用
func (t *LocalTime) Scan(v interface{}) error {
	var timestring string
	switch vt := v.(type) {
	case []uint8:
		timestring = string(vt)
	case time.Time:
		timestring = v.(string)
	default:
		return errors.New("sql 类型错误")
	}
	tTime, _ := time.Parse("2006-01-02 15:04:05", string(timestring))
	*t = LocalTime(tTime)
	return nil
}

// 用于 fmt.Println 和后续验证场景
func (t LocalTime) String() string {
	return time.Time(t).Format(TimeFormat)
}

//InsertReinfo 添加招募信息
func InsertReinfo(info Info) (id int, err error) {
	err = orm.DB.Create(&info).Error
	id = info.Rid
	return
}

//DeleteReinfo 删除招募信息
func DeleteReinfo(rid int) (result Info, err error) {
	err = orm.DB.Where("rid=?", rid).First(&result).Error
	if err != nil {
		return
	}
	if result.Rid != rid {
		err = errors.New("infos do not exit")
		return
	}
	if err = orm.DB.Where(Info{Rid: rid}).Delete(Info{}).Error; err != nil {
		return
	}
	//err = orm.DB.Exec("ALTER TABLE users AUTO_INCREMENT = 1").Error

	return
}

//UpdateReinfo 更新招募信息
func UpdateReinfo(reinfo Info) (result Info, err error) {
	fmt.Println("reinfo:", reinfo)
	err = orm.DB.Model(Info{}).Updates(&reinfo).Error
	//err=orm.DB.Exec("UPDATE reinfos SET reinfo=?,rtag=?,rpublicer=?,rimg=?,rdate=?,rstatus=? WHERE rid=?",reinfo.Rinfo,reinfo.Rtag,reinfo.Rpublicer,reinfo.Rimg,reinfo.Rdate,reinfo.Rstatus,reinfo.Rid).Error
	if err == nil {
		result = reinfo
	}
	return

}

//SelectReinfoByTag 通过标签查询招募信息
func SelectReinfoByTag(tag string) (result []Info, err error) {
	if err = orm.DB.Where("tag=?", tag).Find(&result).Error; err != nil {
		return
	}
	return
}
//SelectReinfoByTtot 通过类别查询信息
func SelectReinfoByTtot(ttot string) (result []Info, err error) {
	if err = orm.DB.Where("ttot=?", ttot).Find(&result).Error; err != nil {
		return
	}
	return
}
//SelectReinfoBypublicer 通过发布组织查赛事
func SelectReinfoBypublicer(uid string) (result []Info, err error) {
	if err = orm.DB.Where("uid=?", uid).Find(&result).Error; err != nil {
		return
	}
	return
}

//SelectReinfobyID 通过id查询赛事
func SelectReinfobyID(id int) (result Info, err error) {
	fmt.Println("selectrid", id)
	err = orm.DB.First(&result, id).Error
	if err != nil {
		err = errors.New("info do not exit")
	}
	return
}

//GetallReinfos 返回所有招募信息
func GetallReinfos() (result []Info, err error) {
	err = orm.DB.Find(&result).Error
	return
}

//
