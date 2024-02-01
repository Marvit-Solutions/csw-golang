package datastruct

import (
	"time"

	"gorm.io/gorm"
)

type SubModule struct {
	ID          string         `gorm:"type:text;primaryKey"`
	CreatedAt   time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	ModuleID    string         `json:"ModulID" form:"ModulID"`
	Name        string         `json:"Name" form:"Name"`
	Description string         `json:"Description" form:"Description"`
	Subject     []Subject      `gorm:"foreignKey:SubModuleID"`
}
