package datastruct

import (
	"time"

	"gorm.io/gorm"
)

type Subscription struct {
	ID            string `gorm:"type:text;primaryKey"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
	TransactionID string         `json:"TransactionID" form:"TransactionID"`
	UserID        string         `json:"UserID" form:"UserID"`
	SubPaketID    string         `json:"SubPaketID" form:"SubPaketID"`
	Description   string         `json:"Description" form:"Description"`
	DueDate       time.Time      `json:"DueDate" form:"DueDate"`
}
