package models

import (
	"database/sql/driver"
	"errors"
	orm "miniprogram/database"
	"time"
)

//TimeFormat 时间格式
const TimeFormat = "2006-01-02 15:04:05"

//LocalTime 本地时间
type LocalTime time.Time

//Contestinfo 赛事信息
type Contestinfo struct {
	Cid      int       `json:"cid"`
	Oid      int       `json:"oid"`
	Ctinfo   string    `json:"ctinfo"`
	Ctag     string    `json:"ctag"`
	Publicer string    `json:"publicer"`
	Cimg     string    `json:"cimg"`
	Cdate    LocalTime `json:"cdate"`
	Cstatus  int       `json:"cstatus"`
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

//InsertContestinfo 添加赛事信息
func InsertContestinfo(contest Contestinfo) (result Contestinfo, err error) {
	err = orm.DB.Create(&contest).Error
	result = contest
	return

}

//DeleteContestinfo 删除Contestinfo
func DeleteContestinfo(id int) (result Contestinfo, err error) {

	err = orm.DB.Where("cid=?", id).First(&result).Error
	if result.Cid == 0 {
		return
	}
	if err = orm.DB.Where(Contestinfo{Cid: id}).Delete(Contestinfo{}).Error; err != nil {
		return
	}
	err = orm.DB.Exec("ALTER TABLE contestinfos AUTO_INCREMENT = 1").Error
	return
}

//UpdateContestinfo 更新Contestinfo
func UpdateContestinfo(contest Contestinfo) (result Contestinfo, err error) {
	err = orm.DB.Model(Contestinfo{}).Where("cid=?", contest.Cid).Update(contest).Error
	if err == nil {
		result = contest
	}
	return

}

//SelectConinfoByTag 通过标签查询赛事信息
func SelectConinfoByTag(tag string) (result []Contestinfo, err error) {
	if err = orm.DB.Where("ctag=?", tag).Find(&result).Error; err != nil {
		return
	}
	return
}

//SelectConinfoBypublicer 通过发布组织查赛事
func SelectConinfoBypublicer(publicer string) (result []Contestinfo, err error) {
	if err = orm.DB.Where("publicer=?", publicer).Find(&result).Error; err != nil {
		return
	}
	return
}

//SelectConinfobyID 通过id查询赛事
func SelectConinfobyID(id int) (result Contestinfo, err error) {
	err = orm.DB.Where("cid=?", id).First(&result).Error
	return
}

//GetallConinfos 返回
func GetallConinfos() (result []Contestinfo, err error) {
	err = orm.DB.Find(&result).Error
	return
}
