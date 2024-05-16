package helper

import (
	"fmt"
	"time"
)

// Pemetaan hari dalam bahasa Inggris ke bahasa Indonesia
var daysOfWeek = map[string]string{
	"Sunday":    "Minggu",
	"Monday":    "Senin",
	"Tuesday":   "Selasa",
	"Wednesday": "Rabu",
	"Thursday":  "Kamis",
	"Friday":    "Jumat",
	"Saturday":  "Sabtu",
}

// Pemetaan bulan dalam bahasa Inggris ke bahasa Indonesia
var months = map[string]string{
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

// output example : "Minggu, 12 Mei 2024, 00:43"
func FormatIndonesianDate(t time.Time) string {
	day := daysOfWeek[t.Weekday().String()]
	month := months[t.Month().String()]
	return fmt.Sprintf("%s, %02d %s %d, %02d:%02d", day, t.Day(), month, t.Year(), t.Hour(), t.Minute())
}

// ConvertTimeFormat converts time to Indonesian date format.
func ConvertDateFormat(t time.Time) string {
	// Using Indonesian date format
	months := [...]string{"Januari", "Februari", "Maret", "April", "Mei", "Juni", "Juli", "Agustus", "September", "Oktober", "November", "Desember"}
	month := months[t.Month()-1]

	// Format: "28 Desember 1998"
	formattedDate := fmt.Sprintf("%d %s %d", t.Day(), month, t.Year())

	return formattedDate
}

func ConvertTimeFormat(inputTime time.Time) string {
	formattedTime := inputTime.Format(TimeFormat)
	return formattedTime
}

func ConvertStrtoClock(inputString string) (time.Time, error) {
	formattedTime, err := time.Parse(TimeFormat, inputString)
	return formattedTime, err
}

func FormatTimeToStringPtr(t time.Time) *string {
	s := t.Format(DateFormat)
	return &s
}

// ConvertStringToTime converts a string representing time into time.Time format.
// inputString should be in the format "2024-04-04T06:0:000".
func ConvertStringToTime(inputString *string) (*time.Time, error) {
	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		return nil, err
	}

	formattedTime, err := time.ParseInLocation(DateFormat, *inputString, loc)
	if err != nil {
		return nil, err
	}

	return &formattedTime, nil
}
