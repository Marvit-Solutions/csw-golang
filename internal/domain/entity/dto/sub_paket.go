package dto

type SubPaketResponse struct {
	Id                uint   `json:"IdSubPaket" form:"IdSubPaket"`
	NamaPaket         string `json:"NamaPaket" form:"NamaPaket"`
	NamaSubPaket      string `json:"NamaSubPaket" form:"NamaSubPaket"`
	DeskripsiSubPaket string `json:"DeskripsiSubPaket" form:"DeskripsiSubPaket"`
	Harga             int    `json:"Harga" form:"Harga"`
}
