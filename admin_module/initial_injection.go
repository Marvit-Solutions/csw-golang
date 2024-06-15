package admin_module

import (
	"github.com/Marvit-Solutions/csw-golang/library/config"
	"gorm.io/gorm"
)

type InitialInjection struct {
}

func NewInitialInjection(sQLMaster *gorm.DB, conf config.Config) InitialInjection {
	return InitialInjection{}
}
