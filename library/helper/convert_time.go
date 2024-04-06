package helper

import (
	"fmt"
	"time"
)

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
	formattedTime := inputTime.Format("15:04")
	return formattedTime
}

func ConvertStrtoClock(inputString string) (time.Time, error) {
	formattedTime, err := time.Parse("15:04", inputString)
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
