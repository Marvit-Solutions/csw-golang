package pretest

import (
	"csw-golang/internal/domain/entity/datastruct"
	"csw-golang/internal/domain/entity/request"

	"gorm.io/gorm"
)

type PretestRepo interface {
	GetAllPretests() (error, []datastruct.Modules)
	GetPretestById(pretestId string) (error, datastruct.SubjectTestTypeQuizzes)
	GetPretestReview(pretestId, status string) (error, datastruct.SubjectTestTypeQuizzes)
	SubmitPretest(id string, req request.PretestSubmitRequest) error
	GradingPretest(id string, req request.PretestSubmitRequest) error
}

type pretestRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) PretestRepo {
	return &pretestRepo{
		db,
	}
}
