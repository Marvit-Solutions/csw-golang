package helper

import "errors"

var (
	ErrDataNotFound = errors.New("data yang anda cari tidak ditemukan")
	ErrAccessDenied = errors.New("access denied")
)

var (
	DateFormat    = "2006-01-02T15:04:05"
	TimeFormat    = "15:04"
	FormattedTime = "15:04:05"
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

// Sub Module ID
const (
	_ = iota
	TWK
	TIU
	TKP
	Matematika
)

// Role User
const (
	_ = iota
	Administrator
	Ceo
	Sestama
	ManajerStaffKeuangan
	ManajerStaffHumas
	KepalaDepartemenStaffSistemInformasi
	KepalaDepartemenOperasional
	ManajerAkademik
	KoordinatorTryOut
	KoordinatorBimbel
	StaffTryOut
	StaffBimbel
	Umum
	PembeliPaketBimbel
)

