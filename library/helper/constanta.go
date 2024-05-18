package helper

import "errors"

var (
	ErrDataNotFound = errors.New("data yang anda cari tidak ditemukan")
)

var (
	DateFormat = "2006-01-02T15:04:05"
	TimeFormat = "15:04"
)

// Test Type ID
var (
	LatihanSoalTestTypeID = 1
	PretestTestTypeID     = 2
	PosttestTestTypeID    = 3
	PaketSoalTestTypeID   = 4
	TryOutTestTypeID      = 5
)

// Module ID
var (
	SKDModuleID        = 1
	MatematikaModuleID = 2
)
