package datastruct

import "gorm.io/gorm"

type SubPaket struct {
	*gorm.Model

	IdSubPaket        string `json:"IdPaket" form:"IdPaket"`
	IdPaket           string `json:"IdPaket" form:"IdPaket"`
	NamaSubPaket      string `json:"NamaPaket" form:"NamaPaket"`
	DeskripsiSubPaket string `json:"DeskripsiPaket" form:"DeskripsiPaket"`
	Harga             int    `json:"Harga" form:"Harga"`
}
