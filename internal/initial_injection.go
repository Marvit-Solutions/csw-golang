package internal

import (
	"github.com/Marvit-Solutions/csw-golang/internal/app/handler/auth"
	"github.com/Marvit-Solutions/csw-golang/library/config"

	"gorm.io/gorm"
)

// InitialInjection represents initial dependency injection for each handler.
type InitialInjection struct {
	Auth auth.Handler
}

// NewInitialInjection initializes the dependencies for the handlers.
func NewInitialInjection(sQLMaster *gorm.DB, conf config.Config) InitialInjection {
	return InitialInjection{
		Auth: auth.NewHandler(sQLMaster, conf),
	}
}
