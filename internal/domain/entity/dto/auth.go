package dto

type AuthResponse struct {
	ID         string `json:"ID" form:"ID"`
	GoogleId   string `json:"GoogleID,omitempty" form:"GoogleID"`
	FacebookId string `json:"FacebookId,omitempty" form:"FacebookId"`
	Email      string `json:"Email" form:"Email" validate:"required,email"`
	Password   string `json:"-"`
	Nama       string `json:"Nama" form:"Nama"`
	Role       string `json:"Role" form:"Role" validate:"required"`
	Telepon    string `json:"Telepon" form:"Telepon" validate:"required,min=10,max=13"`
	Token      string `json:"Token" form:"Token"`
	Alamat     struct {
		Provinsi  string `json:"Provinsi" form:"Provinsi"`
		Kabupaten string `json:"Kabupaten" form:"Kabupaten"`
		Kecamatan string `json:"Kecamatan" form:"Kecamatan"`
	}
	FotoProfil string `json:"FotoProfil" form:"FotoProfil" default:"assets/img/account.png"`
}

type RegisterRequest struct {
	Nama            string `json:"Nama" form:"Nama" validate:"required"`
	GoogleId        string `json:"GoogleID" form:"GoogleID"`
	FacebookId      string `json:"FacebookID" form:"FacebookID"`
	Kelas           string `json:"Kelas" form:"Kelas" validate:"required"`
	Kecamatan       string `json:"Kecamatan" form:"Kecamatan" validate:"required"`
	Kabupaten       string `json:"Kabupaten" form:"Kabupaten" validate:"required"`
	Provinsi        string `json:"Provinsi" form:"Provinsi" validate:"required"`
	Telepon         string `json:"Telepon" form:"Telepon" validate:"required,min=10,max=15,numeric"`
	Email           string `json:"Email" form:"Email" validate:"required,email"`
	Password        string `json:"Password" form:"Password" validate:"required,min=8"`
	ConfirmPassword string `json:"ConfirmPassword" form:"ConfirmPassword" validate:"required,eqfield=Password"`
}

type LoginRequest struct {
	Email    string `json:"Email" form:"Email" validate:"required,email"`
	Password string `json:"Password" form:"Password" validate:"required"`
}
