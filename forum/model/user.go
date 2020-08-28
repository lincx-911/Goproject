package model

import (
	"database/sql/driver"
	"errors"
	"time"
)

//TimeFormat 时间格式
const TimeFormat = "2006-01-02 15:04:05"

//LocalTime 本地时间
type LocalTime time.Time

//User 用户
type User struct {
	ID       int       `json:"id" gorm:"primary_key"`
	Username string    `json:"username"`
	Password string    `json:"-"`
	Email    string    `json:"email"`
	Avatar   string    `json:"avatar"`
	Sign     string    `json:"sign"`
	Utag     []Utag    `json:"utag" gorm:"many2many:user_tag"`
	Regdate  LocalTime `json:"regdate"`
	Role     string    `json:"role"`
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
