package time

import (
	"time"
)

func ConvertTimeFormat(inputTime time.Time) string {
	desiredLayout := "Monday, 02 January 2006, 15:04"
	outputTime := inputTime.Format(desiredLayout)

	return outputTime
}
