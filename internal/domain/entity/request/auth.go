package request

type RegisterRequest struct {
	Name            string `json:"Name" validate:"required"`
	GoogleID        string `json:"GoogleID"`
	FacebookID      string `json:"FacebookID"`
	Class           string `json:"Class" validate:"required"`
	District        string `json:"District" validate:"required"`
	Regency         string `json:"Regency" validate:"required"`
	Province        string `json:"Province"  validate:"required"`
	Phone           string `json:"Phone"  validate:"required,min=10,max=15,numeric"`
	Email           string `json:"Email" validate:"required,email"`
	Password        string `json:"Password" validate:"required,min=8"`
	ConfirmPassword string `json:"ConfirmPassword"  validate:"required,eqfield=Password"`
}

type LoginRequest struct {
	Email    string `json:"Email" validate:"required,email"`
	Password string `json:"Password" validate:"required"`
}
