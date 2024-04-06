package helper

import "errors"

var (
	ErrDataNotFound = errors.New("data yang anda cari tidak ditemukan")
)

var (
	DateFormat = "2006-01-02T15:04:05"
	TimeFormat = "15:04"
)
