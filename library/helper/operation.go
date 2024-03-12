package helper

import (
	"time"
)

func GetLimitOffset(page, size int) (limit int, offset int) {
	if page == 0 || size == 0 {
		// using -1 to disable gorm size and offset in case page and size not set
		size = -1
		offset = -1
		return size, offset
	}
	offset = (page - 1) * size
	return size, offset
}

func TimeLocJakarta() *time.Location {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	return loc
}

func CutString(current string, lenResult int) string {
	var result string
	if len(current) > lenResult {
		result = current[0 : lenResult-1]
	} else {
		result = current
	}

	return result
}
