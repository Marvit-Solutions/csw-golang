package dto

type PaketResponse struct {
	Id             uint   `json:"IdPaket" form:"IdPaket"`
	NamaPaket      string `json:"NamaPaket" form:"NamaPaket"`
	DeskripsiPaket string `json:"DeskripsiPaket" form:"DeskripsiPaket"`
}

type PaketRequest struct {
	NamaPaket      string `json:"NamaPaket" form:"NamaPaket"`
	DeskripsiPaket string `json:"DeskripsiPaket" form:"DeskripsiPaket"`
}
