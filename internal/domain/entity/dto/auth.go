package dto

type AuthResponse struct {
	ID         string `json:"Id" form:"Id"`
	GoogleId   string `json:"GoogleId" form:"GoogleId"`
	FacebookId string `json:"FacebookId" form:"FacebookId"`
	Email      string `json:"Email" form:"Email" validate:"required,email"`
	Password   string `json:"-"`
	Username   string `json:"Username" form:"Username" validate:"required"`
	Name       string `json:"Name" form:"Name"`
	Phone      string `json:"Phone" form:"Phone" validate:"required,min=10,max=13"`
	Token      string `json:"Token" form:"Token"`
}

type RegisterRequest struct {
	Name       string `json:"Name" form:"Name" validate:"required"`
	GoogleId   string `json:"GoogleId" form:"GoogleId"`
	FacebookId string `json:"FacebookId" form:"FacebookId"`
	Email      string `json:"Email" form:"Email" validate:"required,email"`
	Username   string `json:"Username" form:"Username" validate:"required"`
	Phone      string `json:"Phone" form:"Phone" validate:"required,min=10,max=15,numeric"`
	Password   string `json:"Password" form:"Password" validate:"required,min=8"`
}

type LoginRequest struct {
	Email    string `json:"Email" form:"Email" validate:"required,email"`
	Password string `json:"Password" form:"Password" validate:"required"`
}
