package datastruct

import (
	"time"

	"gorm.io/gorm"
)

type LatihanSoal struct {
	ID         string         `gorm:"type:text;primaryKey"`
	CreatedAt  time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt  time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
	MateriID   string         `json:"MateriID" form:"MateriID"`
	Nama       string         `json:"Nama" form:"Nama"`
	Deskripsi  string         `json:"Deskripsi" form:"Deskripsi"`
	Keterangan string         `json:"Keterangan" form:"Keterangan"`
	Status     string         `json:"Status" form:"Status"`
	Waktu      int            `json:"Waktu" form:"Waktu"`
	JumlahSoal int            `json:"JumlahSoal" form:"JumlahSoal"`
	Soal       []Soal         `gorm:"foreignKey:LatihanSoalID"`
	Hasil      []Hasil        `gorm:"foreignKey:LatihanSoalID"`
}
