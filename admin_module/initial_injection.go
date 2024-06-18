package admin_module

import (
	materialAdmin "github.com/Marvit-Solutions/csw-golang/admin_module/app/handler/material"
	quizAdmin "github.com/Marvit-Solutions/csw-golang/admin_module/app/handler/quiz"
	subjectAdmin "github.com/Marvit-Solutions/csw-golang/admin_module/app/handler/subject"
	"github.com/Marvit-Solutions/csw-golang/library/config"
	"gorm.io/gorm"
)

type InitialInjection struct {
	QuizAdmin     quizAdmin.Handler
	MaterialAdmin materialAdmin.Handler
	SubjectAdmin  subjectAdmin.Handler
}

func NewInitialInjection(sQLMaster *gorm.DB, conf config.Config) InitialInjection {
	return InitialInjection{
		QuizAdmin:     quizAdmin.NewHandler(sQLMaster),
		MaterialAdmin: materialAdmin.NewHandler(sQLMaster),
		SubjectAdmin:  subjectAdmin.NewHandler(sQLMaster),
	}
}
