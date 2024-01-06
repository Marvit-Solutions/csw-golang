package datastruct

import (
	"time"

	"gorm.io/gorm"
)

type Subscription struct {
	ID          string `gorm:"type:text;primaryKey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	UserID      uint           `json:"UserID" form:"UserID"`
	SubPaketID  uint           `json:"SubPaketID" form:"SubPaketID"`
	Description string         `json:"Description" form:"Description"`
	DueDate     time.Time      `json:"DueDate" form:"DueDate"`
}
