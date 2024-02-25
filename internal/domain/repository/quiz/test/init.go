package tests

import (
	"csw-golang/internal/domain/entity/dto"
	"csw-golang/internal/domain/entity/request"

	"gorm.io/gorm"
)

type TestRepo interface {
	GetAllTests(req request.QuizParamRequest) (*[]dto.QuizResponse, *dto.Meta, error)
}

type testRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) TestRepo {
	return &testRepo{
		db,
	}
}
