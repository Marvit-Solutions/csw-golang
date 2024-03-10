package cmd

import (
	"csw-golang/internal/handler/auth"
	"csw-golang/internal/handler/mentor"
	"csw-golang/library/config"

	"gorm.io/gorm"
)

type InitialInjection struct {
	Auth   auth.AuthHandler
	Mentor mentor.MentorHandler
}

// NewInitialInjection initial value dependency injection for every handler
func NewInitialInjection(sQLMaster *gorm.DB, conf config.Config) InitialInjection {
	return InitialInjection{
		Auth:   auth.NewAuthHandler(sQLMaster),
		Mentor: mentor.NewMentorHandler(sQLMaster),
	}
}
