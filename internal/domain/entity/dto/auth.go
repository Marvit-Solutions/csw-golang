package dto

import "time"

type AuthResponse struct {
	ID         string `json:"Id" form:"Id"`
	GoogleId   string `json:"GoogleId" form:"GoogleId"`
	FacebookId string `json:"FacebookId" form:"FacebookId"`
	Email      string `json:"Email" form:"Email" validate:"required,email"`
	Password   string `json:"-"`
	Username   string `json:"Username" form:"Username" validate:"required"`
	Nama       string `json:"Nama" form:"Nama"`
	Telepon    string `json:"Telepon" form:"Telepon" validate:"required,min=10,max=13"`
	Token      string `json:"Token" form:"Token"`
	Alamat     struct {
		Provinsi  string `json:"Provinsi" form:"Provinsi"`
		Kabupaten string `json:"Kabupaten" form:"Kabupaten"`
		Kecamatan string `json:"Kecamatan" form:"Kecamatan"`
	}
	TanggalLahir time.Time `json:"TanggalLahir" form:"TanggalLahir" validate:"required"`
	FotoProfil   string    `json:"FotoProfil" form:"FotoProfil" default:"assets/img/account.png"`
}

type RegisterRequest struct {
	Nama       string `json:"Nama" form:"Nama" validate:"required"`
	GoogleId   string `json:"GoogleId" form:"GoogleId"`
	FacebookId string `json:"FacebookId" form:"FacebookId"`
	Email      string `json:"Email" form:"Email" validate:"required,email"`
	Username   string `json:"Username" form:"Username" validate:"required"`
	Telepon    string `json:"Telepon" form:"Telepon" validate:"required,min=10,max=15,numeric"`
	Password   string `json:"Password" form:"Password" validate:"required,min=8"`
}

type LoginRequest struct {
	Email    string `json:"Email" form:"Email" validate:"required,email"`
	Password string `json:"Password" form:"Password" validate:"required"`
}
