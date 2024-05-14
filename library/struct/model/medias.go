package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameMedia = "medias"

// Media mapped from table <medias>
type Media struct {
	ID               int            `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	UUID             string         `gorm:"column:uuid;not null" json:"uuid"`
	Filename         string         `gorm:"column:filename;not null" json:"filename"`
	Mime             string         `gorm:"column:mime;not null" json:"mime"`
	OriginalFilename string         `gorm:"column:original_filename;not null" json:"original_filename"`
	Description      string         `gorm:"column:description;not null" json:"description"`
	CreatedBy        int            `gorm:"column:created_by;not null" json:"created_by"`
	UpdatedBy        int            `gorm:"column:updated_by;not null" json:"updated_by"`
	CreatedAt        time.Time      `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt        time.Time      `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName Media's table name
func (*Media) TableName() string {
	return TableNameMedia
}

type MultiResImage struct {
	MediaUUID string `json:"media_uuid"`
	Original  string `json:"original"`
	Thumbnail string `json:"thumbnail"`
	Desktop   string `json:"desktop"`
	Mobile    string `json:"mobile"`
}
