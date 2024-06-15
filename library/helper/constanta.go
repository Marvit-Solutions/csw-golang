package helper

import "errors"

// Error
var (
	ErrDataNotFound = errors.New("data not found")
	ErrAccessDenied = errors.New("access denied")
)

// Time Format
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

// Plan SKD & Matematika
const (
	_ = iota
	SKDBronze
	SKDSilver
	SKDGold
	SKDPlatinum
	MatematikaBronze
	MatematikaSilver
	MatematikaGold
	MatematikaPlatinum
)
