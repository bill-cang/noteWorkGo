package timeTools

import (
	"time"
)

const (
	TimeLayout             = "2006-01-02 15:04:05"
	TimeLayoutFormatModel1 = "20060102150405000"
	TimeLayoutFormatModel2 = "01-02-2006"
)

// 获取实时时间字符串
func GetSimpleTimeFormat(tf TimeFormatSize, timeFormatLayout string) string {
	layout := getTimeFormatLayout(timeFormatLayout)
	switch tf {
	case Tf_complete:
		return time.Now().Format(layout)
	case Tf_year:
		str := time.Now().Format(layout)
		return str[:4]
	case Tf_mouth:
		str := time.Now().Format(layout)
		return str[:6]
	case Tf_date:
		str := time.Now().Format(layout)
		return str[:8]
	case Tf_hour:
		str := time.Now().Format(layout)
		return str[:10]
	case Tf_minute:
		str := time.Now().Format(layout)
		return str[:12]
	case Tf_second:
		str := time.Now().Format(layout)
		return str[:14]
	}
	return ""
}

//获取时间格式化样式
func getTimeFormatLayout(tl string) string {
	if tl == "" {
		tl = TimeLayout
	}
	return tl
}

//获取对应时区的零点时间戳
func GetAfterSomeDateZeroClockStamp(country string, dateNub int) (zeroStamp int) {
	af := time.Now().AddDate(0, 0, dateNub)
	stamp := time.Date(af.Year(), af.Month(), af.Day(), 0, 0, 0, 0, GetTimeLocation(country)).Unix()
	return int(stamp)
}

//获取对应时区Ｎ天后的H小时时间戳
func GetServerTimestamp(country string, dateNub int, hour int) (originStamp int64) {
	af := time.Now().AddDate(0, 0, dateNub)
	stamp := time.Date(af.Year(), af.Month(), af.Day(), hour, 0, 0, 0, GetTimeLocation(country)).Unix()
	return stamp
}

//以当前时间整点为基准获取前后整点时间戳
func GetOClickTimestamp(hours int64) int64 {
	now := time.Now()
	timestamp := now.Unix() - (int64(now.Second()) + int64((60 * now.Minute()))) + (3600 * hours)
	return timestamp
}

//根据时区获取时间戳对应时间(找不到时区以中国为准)
func GetFormatDateFormLocation(country string, timestamp int64) (formatDate string) {
	//时间戳转日期
	formatDate = time.Unix(timestamp, 0).In(GetTimeLocation(country)).Format(TimeLayout) //设置时间戳 使用模板格式化为日期字符串
	return
}

//根据国家名称获取时区
func GetTimeLocation(country string) (location *time.Location) {
	if s, ok := TimeZone[country]; ok {
		location, _ = time.LoadLocation(s)
	} else {
		location, _ = time.LoadLocation(TimeZone["China"])
	}
	return
}

//字符串时间转时间秒
func GetDateStrToStamp(country, dateStr string) (stamp int64, err error) {
	var ti time.Time
	ti, err = time.ParseInLocation(TimeLayout, dateStr, GetTimeLocation(country))
	if nil != err {
		return 0, err
	}
	return ti.Unix(), nil
}

/*==============================================================================*/
// 获取指定时间的字符串
func GetDateTimeStr(t *time.Time) string {
	if nil == t {
		return ""
	}
	return t.Format(TimeLayout)
}

// 根据时间字符串获取时间
func GetTimeFromDateTime(t string) (tt *time.Time, err error) {
	if "" == t {
		return
	}
	var ft time.Time
	ft, err = time.ParseInLocation(TimeLayout, t, time.Local)
	if err != nil {
		return
	}
	tt = &ft
	return
}

func GetTimeFromTimestamp(t int64) (tt *time.Time) {
	if t <= 0 {
		return
	}
	var ft time.Time
	ft = time.Unix(t, 0)
	tt = &ft
	return
}

// 获取指定时区的当天开始时间对应的北京时间 换算对应当地时区的今天开始时间(北京时间) subHour +- 时差 2019-11-30 05:00:00
func GetTimeZoneTodayDataTimeStr(subHour int) (ts string) {
	now := time.Now()
	day := now.Day()
	hour := 0 - subHour
	if now.Hour() < hour {
		day -= 1
	}

	tt := time.Date(now.Year(), now.Month(), day, hour, 0, 0, 0, time.Local)
	ts = tt.Format(TimeLayout)
	return
}

//时间秒数换算天、时、分
func GetEqualDay(seconds int64) (eq [4]int32) {
	day := seconds / 86400
	eq[0] = int32(day)
	hour := (seconds - day*86400) / 3600
	eq[1] = int32(hour)
	minute := (seconds - day*86400 - hour*3600) / 60
	eq[2] = int32(minute)
	second := seconds - day*86400 - hour*3600 - minute*60
	eq[3] = int32(second)
	return
}
