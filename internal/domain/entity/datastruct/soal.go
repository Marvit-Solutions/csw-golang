package datastruct

import (
	"time"

	"gorm.io/gorm"
)

type Soal struct {
	ID            string         `gorm:"type:text;primaryKey"`
	CreatedAt     time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt     time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt     gorm.DeletedAt `gorm:"index"`
	LatihanSoalID string         `json:"LatihanSoalID" form:"LatihanSoalID"`
	Nomor         int            `json:"Nomor" form:"Nomor"`
	Status        string         `json:"Status" form:"Status"`
	Mark          int            `json:"Mark" form:"Mark"`
	Tanda         bool           `json:"Tanda" form:"Tanda"`
	Question      string         `json:"Question" form:"Question"`
	Answers       []Answer       `gorm:"foreignKey:SoalID"`
	StoredAnswers []StoredAnswer `gorm:"foreignKey:SoalID"`
}
