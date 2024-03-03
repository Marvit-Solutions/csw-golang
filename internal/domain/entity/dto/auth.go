package dto

type AuthResponse struct {
	AccessToken string  `json:"AccessToken"`
	ExpiredIn   float64 `json:"ExpiredIn"`
	ExpiredAt   int64   `json:"ExpiredAt"`
}
