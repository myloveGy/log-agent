package util

import (
	"time"
)

const (
	TimeFormat       = "2006-01-02 15:04:05"
	DateFormat       = "2006-01-02"
	TimeZoneShanghai = "Asia/Shanghai"
)

var loc, _ = time.LoadLocation(TimeZoneShanghai)

func Date() string {
	return time.Now().In(loc).Format(DateFormat)
}

func DateTime() string {
	return time.Now().In(loc).Format(TimeFormat)
}

func ParseDate(value string) (time.Time, error) {
	return time.ParseInLocation(DateFormat, value, loc)
}

func ParseDateTime(value string) (time.Time, error) {
	return time.ParseInLocation(TimeFormat, value, loc)
}
