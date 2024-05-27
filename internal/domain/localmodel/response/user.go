package response

import "github.com/Marvit-Solutions/csw-golang/library/struct/model"

type User struct {
	UUID        string               `json:"uuid"`
	Email       string               `json:"email"`
	Name        string               `json:"name"`
	Media       *model.MultiResImage `json:"media"`
	Province    string               `json:"province"`
	Regency     string               `json:"regency"`
	District    string               `json:"district"`
	PhoneNumber string               `json:"phone_number"`
	Class       string               `json:"class"`
}
