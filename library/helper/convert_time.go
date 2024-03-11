package helper

import (
	"fmt"
	"time"
)

// ConvertTimeFormat converts time to Indonesian date format.
func ConvertTimeFormat(t time.Time) string {
	// Using Indonesian date format
	months := [...]string{"Januari", "Februari", "Maret", "April", "Mei", "Juni", "Juli", "Agustus", "September", "Oktober", "November", "Desember"}
	month := months[t.Month()-1]

	// Format: "28 Desember 1998"
	formattedDate := fmt.Sprintf("%d %s %d", t.Day(), month, t.Year())

	return formattedDate
}
