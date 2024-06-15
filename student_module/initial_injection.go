package student_module

import (
	"github.com/Marvit-Solutions/csw-golang/student_module/app/handler/auth"
	"github.com/Marvit-Solutions/csw-golang/student_module/app/handler/dashboard"
	"github.com/Marvit-Solutions/csw-golang/student_module/app/handler/exercise"
	"github.com/Marvit-Solutions/csw-golang/student_module/app/handler/home"
	"github.com/Marvit-Solutions/csw-golang/student_module/app/handler/location"
	"github.com/Marvit-Solutions/csw-golang/student_module/app/handler/module"
	"github.com/Marvit-Solutions/csw-golang/student_module/app/handler/quiz"

	"github.com/Marvit-Solutions/csw-golang/library/config"
	"github.com/Marvit-Solutions/csw-golang/student_module/app/handler/user"

	"gorm.io/gorm"
)

// InitialInjection represents initial dependency injection for each handler.
type InitialInjection struct {
	Auth      auth.Handler
	Home      home.Handler
	Location  location.Handler
	Module    module.Handler
	Quiz      quiz.Handler
	Exercise  exercise.Handler
	Dashboard dashboard.Handler
	User      user.Handler
}

// NewInitialInjection initializes the dependencies for the handlers.
func NewInitialInjection(sQLMaster *gorm.DB, conf config.Config) InitialInjection {
	return InitialInjection{
		Auth:      auth.NewHandler(sQLMaster),
		Home:      home.NewHandler(sQLMaster),
		Location:  location.NewHandler(sQLMaster, conf),
		Module:    module.NewHandler(sQLMaster),
		Quiz:      quiz.NewHandler(sQLMaster),
		Dashboard: dashboard.NewHandler(sQLMaster),
		Exercise:  exercise.NewHandler(sQLMaster),
		User:      user.NewHandler(sQLMaster),
	}
}
