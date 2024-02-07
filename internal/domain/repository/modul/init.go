package module

import (
	"csw-golang/internal/domain/entity/datastruct"

	"gorm.io/gorm"
)

type ModuleRepo interface {
	GetListModules() ([]datastruct.SubModules, error)
	GetListSubjectQuiz() ([]datastruct.SubjectTestTypeQuizzes, error)
	GetSubjectsBySubmoduleID(submoduleID string) ([]datastruct.Subjects, error)
	GetTestByTestTypeID(testTypeID string) (datastruct.SubjectTestTypeQuizzes, error)
}

type moduleRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) ModuleRepo {
	return &moduleRepo{
		db,
	}
}
