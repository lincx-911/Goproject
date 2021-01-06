package common

import (
	"time"
)
//TimeFormat 时间格式
const TimeFormat = "2006-01-02 15:04:05"

//UnixTime 时间类型
type UnixTime time.Time

//MarshalJson 时间转Json
func (t UnixTime) MarshalJSON() ([]byte, error)  {
	b := make([]byte, 0, len(TimeFormat)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, TimeFormat)
	b = append(b, '"')
	return b, nil
}

//UnmarshalJSON 解析时间json
func (t *UnixTime) UnmarshalJSON(data []byte) (err error) {
	// 空值不进行解析
	if len(data) == 2 {
		*t = UnixTime(time.Time{})
		return
	}
	// 指定解析的格式
	now, err := time.Parse(`"`+TimeFormat+`"`, string(data))
	*t = UnixTime(now)
	return
}