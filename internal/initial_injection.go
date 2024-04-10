package internal

import (
	"github.com/Marvit-Solutions/csw-golang/internal/app/handler/auth"
	"github.com/Marvit-Solutions/csw-golang/internal/app/handler/home"
	"github.com/Marvit-Solutions/csw-golang/internal/app/handler/location"
	"github.com/Marvit-Solutions/csw-golang/internal/app/handler/modul"
	"github.com/Marvit-Solutions/csw-golang/library/config"

	"gorm.io/gorm"
)

// InitialInjection represents initial dependency injection for each handler.
type InitialInjection struct {
	Auth     auth.Handler
	Home     home.Handler
	Location location.Handler
	Modul    modul.Handler
}

// NewInitialInjection initializes the dependencies for the handlers.
func NewInitialInjection(sQLMaster *gorm.DB, conf config.Config) InitialInjection {
	return InitialInjection{
		Auth:     auth.NewHandler(sQLMaster),
		Home:     home.NewHandler(sQLMaster),
		Location: location.NewHandler(sQLMaster, conf),
		Modul:    modul.NewHandler(sQLMaster),
	}
}
