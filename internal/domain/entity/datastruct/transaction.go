package datastruct

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	ID                string `gorm:"type:text;primaryKey"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         gorm.DeletedAt `gorm:"index"`
	TransactionID     string         `json:"TransactionID" form:"TransactionID"`
	UserID            uint           `json:"UserID" form:"UserID"`
	OrderID           string         `json:"OrderID" form:"OrderID"`
	ReceiptNumber     string         `json:"ReceiptNumber" form:"ReceiptNumber"`
	StatusTransaction string         `json:"StatusTransaction" form:"StatusTransaction"`
	PaymentURL        string         `json:"PaymentURL" form:"PaymentURL"`
	TotalPrice        float64        `json:"TotalPrice" form:"TotalPrice"`
	PaymentMethod     string         `json:"PaymentMethod" form:"PaymentMethod"`
	PaymentStatus     string         `json:"PaymentStatus" form:"PaymentStatus"`
	CanceledReason    string         `json:"CanceledReason" form:"CanceledReason"`
}
