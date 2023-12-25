package dto

type PaketResponse struct {
	IdPaket        string `json:"IdPaket" form:"IdPaket"`
	NamaPaket      string `json:"NamaPaket" form:"NamaPaket"`
	DeskripsiPaket string `json:"DeskripsiPaket" form:"DeskripsiPaket"`
}
