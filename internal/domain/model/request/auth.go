package request

type RegisterRequest struct {
	Name            string `json:"name" form:"name" validate:"required"`
	GoogleID        string `json:"google_id" form:"google_id"`
	FacebookID      string `json:"facebook_id" form:"facebook_id"`
	Class           string `json:"class" form:"class" validate:"required"`
	District        string `json:"district" form:"district" validate:"required"`
	Regency         string `json:"regency" form:"regency" validate:"required"`
	Province        string `json:"province" form:"province" validate:"required"`
	Phone           string `json:"phone" form:"phone" validate:"required,min=10,max=15,numeric"`
	Email           string `json:"email" form:"email" validate:"required,email"`
	Password        string `json:"password" form:"password" validate:"required,min=8"`
	ConfirmPassword string `json:"confirm_password" form:"confirm_password" validate:"required,eqfield=Password"`
}

type LoginRequest struct {
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required"`
}
