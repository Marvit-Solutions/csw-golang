package datastruct

import (
	"time"

	"gorm.io/gorm"
)

type Subscriptions struct {
	ID            string `gorm:"type:text;primaryKey"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
	TransactionID string         `json:"TransactionID" form:"TransactionID"`
	UserID        string         `json:"UserID" form:"UserID"`
	SubPlanID     string         `json:"SubPlanID" form:"SubPlanID"`
	Description   string         `json:"Description" form:"Description"`
	DueDate       time.Time      `json:"DueDate" form:"DueDate"`
}
