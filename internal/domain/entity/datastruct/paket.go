package datastruct

import "gorm.io/gorm"

type Paket struct {
	*gorm.Model

	IdPaket        string `json:"IdPaket" form:"IdPaket"`
	NamaPaket      string `json:"NamaPaket" form:"NamaPaket"`
	DeskripsiPaket string `json:"DeskripsiPaket" form:"DeskripsiPaket"`
}
