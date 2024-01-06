package dto

type SubPaketResponse struct {
	ID                string `json:"IdSubPaket" form:"IdSubPaket"`
	NamaPaket         string `json:"NamaPaket" form:"NamaPaket"`
	NamaSubPaket      string `json:"NamaSubPaket" form:"NamaSubPaket"`
	DeskripsiSubPaket string `json:"DeskripsiSubPaket" form:"DeskripsiSubPaket"`
	Harga             int    `json:"Harga" form:"Harga"`
}

type SubPaketRequest struct {
	IDPaket           string `json:"IDPaket" form:"IDPaket"`
	NamaSubPaket      string `json:"NamaSubPaket" form:"NamaSubPaket"`
	DeskripsiSubPaket string `json:"DeskripsiSubPaket" form:"DeskripsiSubPaket"`
	Harga             int    `json:"Harga" form:"Harga"`
}

type TopSubPaketResponse struct {
	SubPaketID   string `json:"IdSubPaket" form:"IdSubPaket"`
	NamaSubPaket string `json:"NamaSubPaket" form:"NamaSubPaket"`
	Total        int    `json:"Total" form:"Total"`
}
