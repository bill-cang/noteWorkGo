package timeTools

//截取粒度枚举
type TimeFormatSize int

const (
	Tf_complete TimeFormatSize = 0
	Tf_year     TimeFormatSize = 1
	Tf_mouth    TimeFormatSize = 2
	Tf_date     TimeFormatSize = 3
	Tf_hour     TimeFormatSize = 4
	Tf_minute   TimeFormatSize = 5
	Tf_second   TimeFormatSize = 6
)

var (
	TimeZone = map[string]string{
		"China":     "Asia/Shanghai",
		"Hong Kong": "Asia/Shanghai",
		"Kenya":     "Africa/Nairobi",
		"Nigeria":   "Africa/Lagos",
		"Ghana":     "Africa/Accra",
	}
)

