package dto

type PaketResponse struct {
	Id             uint   `json:"IdPaket" form:"IdPaket"`
	NamaPaket      string `json:"NamaPaket" form:"NamaPaket"`
	DeskripsiPaket string `json:"DeskripsiPaket" form:"DeskripsiPaket"`
}
