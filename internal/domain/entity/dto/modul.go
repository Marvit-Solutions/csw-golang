package dto

type ModulResponse struct {
	ID        string `json:"ID" form:"ID"`
	Nama      string `json:"Nama" form:"Nama"`
	Deskripsi string `json:"Deskripsi" form:"Deskripsi"`
	Materi    struct {
		ID   string `json:"ID" form:"ID"`
		Nama string `json:"Nama" form:"Nama"`
	}
	LatihanSoal struct {
		ID   string `json:"ID" form:"ID"`
		Nama string `json:"Nama" form:"Nama"`
	}
}

type MateriResponse struct {
	ID        string      `json:"ID" form:"ID"`
	Nama      string      `json:"Nama" form:"Nama"`
	SubMateri []SubMateri `json:"SubMateri" form:"SubMateri"`
}

type SubMateri struct {
	ID     string `json:"ID" form:"ID"`
	Nama   string `json:"Nama" form:"Nama"`
	Konten string `json:"Konten" form:"Konten"`
}

type LatihanSoalResponse struct {
	ID         string `json:"ID" form:"ID"`
	Nama       string `json:"Nama" form:"Nama"`
	Deskripsi  string `json:"Deskripsi" form:"Deskripsi"`
	Keterangan string `json:"Keterangan" form:"Keterangan"`
	Status     string `json:"Status" form:"Status"`
	Waktu      int    `json:"Waktu" form:"Waktu"`
	JumlahSoal int    `json:"JumlahSoal" form:"JumlahSoal"`
	Soal       []Soal `json:"Soal" form:"Soal"`
}

type Soal struct {
	ID       string   `json:"ID" form:"ID"`
	Nomor    int      `json:"Nomor" form:"Nomor"`
	Status   string   `json:"Status" form:"Status"`
	Mark     int      `json:"Mark" form:"Mark"`
	Tanda    bool     `json:"Tanda" form:"Tanda"`
	Question string   `json:"Question" form:"Question"`
	Answers  []Answer `json:"Answers" form:"Answers"`
}

type Answer struct {
	ID      string  `json:"ID" form:"ID"`
	Jenis   string  `json:"Jenis" form:"Jenis"`
	Konten  string  `json:"Konten" form:"Konten"`
	Dipilih bool    `json:"Dipilih" form:"Dipilih"`
	Nilai   float64 `json:"Nilai" form:"Nilai"`
}
