package helper

import (
	"time"
)

func FormatTimeToStringPtr(t time.Time) *string {
	s := t.Format(DateFormat)
	return &s
}

func ParseTimeString(timeStr string) time.Duration {
	parsedTime, _ := time.Parse("15:04:05", timeStr)
	hours := parsedTime.Hour()
	minutes := parsedTime.Minute()
	seconds := parsedTime.Second()

	return time.Duration(hours)*time.Hour + time.Duration(minutes)*time.Minute + time.Duration(seconds)*time.Second
}
