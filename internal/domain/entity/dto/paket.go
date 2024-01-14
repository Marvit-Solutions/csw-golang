package dto

type SubPaketDetail struct {
	ID          string `json:"ID" form:"ID"`
	GrupPejuang bool   `json:"GrupPejuang" form:"GrupPejuang"`
	SoalLatihan uint   `json:"SoalLatihan" form:"SoalLatihan"`
	Akses       uint   `json:"Akses" form:"Akses"`
	Modul       bool   `json:"Modul" form:"Modul"`
	TryOut      uint   `json:"TryOut" form:"TryOut"`
	Zoom        bool   `json:"Zoom" form:"Zoom"`
}

type SubPaket struct {
	ID             string         `json:"ID" form:"ID"`
	NamaSubPaket   string         `json:"NamaSubPaket" form:"NamaSubPaket"`
	Harga          float64        `json:"Harga" form:"Harga"`
	SubPaketDetail SubPaketDetail `json:"SubPaketDetail" form:"SubPaketDetail"`
}
type PaketResponse struct {
	ID        string     `json:"ID" form:"ID"`
	NamaPaket string     `json:"NamaPaket" form:"NamaPaket"`
	SubPaket  []SubPaket `json:"SubPaket" form:"SubPaket"`
}

type PaketRequest struct {
	NamaPaket      string `json:"NamaPaket" form:"NamaPaket"`
	DeskripsiPaket string `json:"DeskripsiPaket" form:"DeskripsiPaket"`
}
