package dto

type AuthResponse struct {
	ID             string `json:"ID" form:"ID"`
	GoogleID       string `json:"GoogleID" form:"GoogleID"`
	FacebookID     string `json:"FacebookId" form:"FacebookId"`
	Email          string `json:"Email" form:"Email" validate:"required,email"`
	Password       string `json:"-"`
	Name           string `json:"Name" form:"Name"`
	Role           string `json:"Role" form:"Role" validate:"required"`
	Token          string `json:"Token" form:"Token"`
	ProfilePicture string `json:"FotoProfil" form:"FotoProfil" default:"assets/img/users/profile/account.png"`
}

type RegisterRequest struct {
	Name            string `json:"Name" form:"Name" validate:"required"`
	GoogleID        string `json:"GoogleID" form:"GoogleID"`
	FacebookID      string `json:"FacebookID" form:"FacebookID"`
	Class           string `json:"Class" form:"Class" validate:"required"`
	District        string `json:"District" form:"District" validate:"required"`
	Regency         string `json:"Regency" form:"Regency" validate:"required"`
	Province        string `json:"Province" form:"Province" validate:"required"`
	Phone           string `json:"Phone" form:"Phone" validate:"required,min=10,max=15,numeric"`
	Email           string `json:"Email" form:"Email" validate:"required,email"`
	Password        string `json:"Password" form:"Password" validate:"required,min=8"`
	ConfirmPassword string `json:"ConfirmPassword" form:"ConfirmPassword" validate:"required,eqfield=Password"`
}

type LoginRequest struct {
	Email    string `json:"Email" form:"Email" validate:"required,email"`
	Password string `json:"Password" form:"Password" validate:"required"`
}
