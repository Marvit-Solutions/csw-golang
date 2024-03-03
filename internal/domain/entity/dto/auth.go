package dto

type AuthResponse struct {
	ID             string `json:"ID" form:"ID"`
	GoogleID       string `json:"GoogleID" form:"GoogleID"`
	FacebookID     string `json:"FacebookId" form:"FacebookId"`
	Email          string `json:"Email" form:"Email" validate:"required,email"`
	Password       string `json:"-"`
	Name           string `json:"Name" form:"Name"`
	Token          string `json:"Token" form:"Token"`
	ProfilePicture string `json:"FotoProfil" form:"FotoProfil" default:"assets/img/users/profile/account.png"`
}
