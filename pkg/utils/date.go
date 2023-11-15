package util

import (
	"time"
)

const (
	TimeLayout = "2006-01-02 15:04:05"
)

func ParseRecentDate(n int) []string {
	now := time.Now()
	res := make([]string, 0, n)
	for i := n - 1; i >= 0; i-- {
		dateStr := time.Date(now.Year(), now.Month(), now.Day()-i, 0, 0, 0, 0, time.Local).Format("2006-01-02")
		res = append(res, dateStr)
	}
	return res
}

func ThisDayBeginTime() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
}

// DiffDay相隔天数
func DiffDay(t1, t2 time.Time) int64 {
	t1 = time.Date(t1.Year(), t1.Month(), t1.Day(), 0, 0, 0, 0, time.Local)
	t2 = time.Date(t2.Year(), t2.Month(), t2.Day(), 0, 0, 0, 0, time.Local)
	days := int64(t2.Sub(t1) / time.Hour / 24)
	if days > 0 {
		return days
	}
	return -days
}
