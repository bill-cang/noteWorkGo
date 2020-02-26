package timeTools

import (
	"time"
)

const (
	TimeLayout       = "2006-01-02 15:04:05"
	TimeFormatLayout = "20060102150405000"
)

// 获取实时时间字符串
func GetRealTimeStr(tf TimeFormatSize, timeFormatLayout string) string {
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

func getTimeFormatLayout(tl string) string {
	if tl == "" {
		tl = TimeLayout
	}
	return tl
}

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

//获取N天后的凌晨）0点
func GetAfterSomeDateOriginStamp(dateNub int) (originStamp int) {
	af := time.Now().AddDate(0, 0, dateNub)
	stamp := time.Date(af.Year(), af.Month(), af.Day(), 0, 0, 0, 0, af.Location()).Unix()
	return int(stamp)
}

// 获取年龄 bir 为秒数
func GetAgeByBirth(bir int64) (age int) {
	yearBir := time.Unix(bir/1000, 0).Year()
	yearNow := time.Now().Year()
	age = yearNow - yearBir
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

// 获取当地时间的日期
func GetTimeZoneTodayDateStr(subHour int) (ts string) {
	now := time.Now()
	day := now.Day()
	hour := 0 - subHour
	if now.Hour() < hour {
		day -= 1
	}
	tt := time.Date(now.Year(), now.Month(), day, 0, 0, 0, 0, time.Local)
	ts = tt.Format(TimeFormatLayout)
	return
}

//获取服务器Ｎ天后的H小时时间戳
func GetServerTimestamp(dateNub int, hour int) (originStamp int64) {
	af := time.Now().AddDate(0, 0, dateNub)
	stamp := time.Date(af.Year(), af.Month(), af.Day(), hour, 0, 0, 0, af.Location()).Unix()
	return stamp
}

//以当前时间整点为基准获取前后整点时间戳
func GetOClickTimestamp(hours int64) int64 {
	now := time.Now()
	timestamp := now.Unix() - (int64(now.Second()) + int64((60 * now.Minute()))) + (3600 * hours)
	//fmt.Println(timestamp, time.Unix(timestamp, 0), now.Unix())
	return timestamp
}

//根据时区获取时间戳对应时间(获取失败以肯尼亚为准)
func GetFormatDateFormLocation(loc string, timestamp int64) (formatDate string, err error) {
	location := GetTimeLocation(loc)
	//时间戳转日期
	formatDate = time.Unix(timestamp, 0).In(location).Format(TimeLayout) //设置时间戳 使用模板格式化为日期字符串
	return
}

//根据国家名称获取时区
func GetTimeLocation(country string) (location *time.Location) {
	location, _ = time.LoadLocation(TimeZone["Kenya"])
	if s, ok := TimeZone[country]; ok {
		location, _ = time.LoadLocation(s)
	} else {
		//xlog.LevelLogfn(xlog.WARN, "[GetTimeLocation] country = %s is not record!", country)

	}
	return
}
