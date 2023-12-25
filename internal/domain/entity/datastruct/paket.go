package datastruct

import "gorm.io/gorm"

type Paket struct {
	*gorm.Model
	NamaPaket      string `json:"NamaPaket" form:"NamaPaket"`
	DeskripsiPaket string `json:"DeskripsiPaket" form:"DeskripsiPaket"`
}
