package utils

import (
	"time"
)

const (
	TimeBarFormat          = "2006-01-02 15:04:05"
	TimeBarFormatPM        = "2006-01-02 15:04:05 PM"
	TimeFormatHMS          = "20060102150405"
	TimeUnderlineYearMonth = "2006_01"
	TimeBarYYMMDD          = "2006-01-02"
	TimeHHMMSS             = "15:04:05"
	TimeYYMMDD             = "20060102"
	TimeYYMMDH             = "2006010215"
	TimeYMDHM              = "200601021504"
	TimeBEIJINGFormat      = "2006-01-02 15:04:05 +08:00"
	TimeGDFormat           = "01/02/2006 15:04:05"
	TimeTFormat            = "2006-01-02T15:04:05"
	TimeTBjFormat          = "2006-01-02T15:04:05+08:00"

	Minute   = 60
	HourVal  = Minute * 60
	DayVal   = HourVal * 24
	MonthVal = DayVal * 30
	YearVal  = MonthVal * 365

	BeiJinAreaTime = "Asia/Shanghai"
)

func GetBjNowTime() time.Time {
	// 获取北京时间, 在 windows系统上 time.LoadLocation 会加载失败, 最好的办法是用 time.FixedZone
	var bjLoc *time.Location
	var err error
	bjLoc, err = time.LoadLocation(BeiJinAreaTime)
	if err != nil {
		bjLoc = time.FixedZone("CST", 8*3600)
	}

	return time.Now().In(bjLoc)
}

func StrToTime(value string) time.Time {
	if value == "" {
		return time.Time{}
	}
	layouts := []string{
		"2006-01-02 15:04:05",
		"2006-01-02 15:04:05 -0700 MST",
		"2006-01-02 15:04:05 -0700",
		"2006/01/02 15:04:05 -0700 MST",
		"2006/01/02 15:04:05 -0700",
		"2006/01/02 15:04:05",
		"2006-01-02 -0700 MST",
		"2006-01-02 -0700",
		"2006-01-02",
		"2006/01/02 -0700 MST",
		"2006/01/02 -0700",
		"2006/01/02",
		"2006-01-02 15:04:05 -0700 -0700",
		"2006/01/02 15:04:05 -0700 -0700",
		"2006-01-02 -0700 -0700",
		"2006/01/02 -0700 -0700",
		time.ANSIC,
		time.UnixDate,
		time.RubyDate,
		time.RFC822,
		time.RFC822Z,
		time.RFC850,
		time.RFC1123,
		time.RFC1123Z,
		time.RFC3339,
		time.RFC3339Nano,
		time.Kitchen,
		time.Stamp,
		time.StampMilli,
		time.StampMicro,
		time.StampNano,
	}

	var t time.Time
	var err error
	for i := 0; i < len(layouts); i++ {
		layout := layouts[i]
		t, err = time.ParseInLocation(layout, value, time.FixedZone("shanghai/beijing", 8*3600))
		if err == nil {
			return t
		}
	}

	return t
}

// BJNowTime 北京当前时间
func BJNowTime() time.Time {
	// 获取北京时间, 在 windows系统上 time.LoadLocation 会加载失败, 最好的办法是用 time.FixedZone, es 中的时间为: "2019-03-01T21:33:18+08:00"
	var beiJinLocation *time.Location
	var err error

	beiJinLocation, err = time.LoadLocation("Asia/Shanghai")
	if err != nil {
		beiJinLocation = time.FixedZone("CST", 8*3600)
	}

	nowTime := time.Now().In(beiJinLocation)

	return nowTime
}

func GetBjTimeLoc() *time.Location {
	// 获取北京时间, 在 windows系统上 time.LoadLocation 会加载失败, 最好的办法是用 time.FixedZone
	var bjLoc *time.Location
	var err error
	bjLoc, err = time.LoadLocation(BeiJinAreaTime)
	if err != nil {
		bjLoc = time.FixedZone("CST", 8*3600)
	}

	return bjLoc
}

func GetTime000(ti time.Time) time.Time {
	y, m, d := ti.Date()
	return time.Date(y, m, d, 0, 0, 0, 0, GetBjTimeLoc())
}
