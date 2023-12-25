package datastruct

import "gorm.io/gorm"

type SubPaket struct {
	*gorm.Model
	IdPaket           uint   `json:"IdPaket" form:"IdPaket"`
	NamaSubPaket      string `json:"NamaPaket" form:"NamaPaket"`
	DeskripsiSubPaket string `json:"DeskripsiPaket" form:"DeskripsiPaket"`
	Harga             int    `json:"Harga" form:"Harga"`
}
