package time

import (
	"time"
)

func ConvertTimeFormat(inputTime time.Time) string {
	// Define the desired output format
	desiredLayout := "Monday, 02 January 2006, 15:04" // Adjust according to your needs

	// Format the time according to the desired layout
	outputTime := inputTime.Format(desiredLayout)

	return outputTime
}
