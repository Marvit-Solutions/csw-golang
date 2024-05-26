package helper

import (
	"fmt"
	"strings"
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

func ConvertToIndonesianFormat(t time.Time) string {
	formattedTime := t.Format("Monday, 02 January 2006, 15:04")

	dayMap := map[string]string{
		"Monday":    "Senin",
		"Tuesday":   "Selasa",
		"Wednesday": "Rabu",
		"Thursday":  "Kamis",
		"Friday":    "Jumat",
		"Saturday":  "Sabtu",
		"Sunday":    "Minggu",
	}

	monthMap := map[string]string{
		"January":   "Januari",
		"February":  "Februari",
		"March":     "Maret",
		"April":     "April",
		"May":       "Mei",
		"June":      "Juni",
		"July":      "Juli",
		"August":    "Agustus",
		"September": "September",
		"October":   "Oktober",
		"November":  "November",
		"December":  "Desember",
	}

	for en, id := range dayMap {
		formattedTime = strings.ReplaceAll(formattedTime, en, id)
	}

	for en, id := range monthMap {
		formattedTime = strings.ReplaceAll(formattedTime, en, id)
	}

	return formattedTime
}

func ConvertDurationToIndonesian(timeStr string) string {
	parsedTime, err := time.Parse("15:04:05", timeStr)
	if err != nil {
		return "Error parsing time"
	}

	hours := fmt.Sprintf("%02d", parsedTime.Hour())
	minutes := fmt.Sprintf("%02d", parsedTime.Minute())
	seconds := fmt.Sprintf("%02d", parsedTime.Second())

	return fmt.Sprintf("%s jam %s Menit %s Detik", hours, minutes, seconds)
}
