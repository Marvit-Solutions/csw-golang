package dto

type AuthResponse struct {
	ID             string `json:"ID" form:"ID"`
	GoogleID       string `json:"GoogleID"`
	FacebookID     string `json:"FacebookID"`
	Email          string `json:"Email"`
	Password       string `json:"-"`
	Name           string `json:"Name"`
	Token          string `json:"Token"`
	ProfilePicture string `json:"ProfilePicture"`
}
