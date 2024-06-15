package admin_module

import (
	quizAdmin "github.com/Marvit-Solutions/csw-golang/admin_module/app/handler/quiz"
	"github.com/Marvit-Solutions/csw-golang/library/config"
	"gorm.io/gorm"
)

type InitialInjection struct {
	QuizAdmin quizAdmin.Handler
}

func NewInitialInjection(sQLMaster *gorm.DB, conf config.Config) InitialInjection {
	return InitialInjection{
		QuizAdmin: quizAdmin.NewHandler(sQLMaster),
	}
}
