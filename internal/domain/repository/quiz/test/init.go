package tests

import (
	"csw-golang/internal/domain/entity/dto"
	"csw-golang/internal/domain/entity/request"
	"csw-golang/internal/domain/helper/response"

	"gorm.io/gorm"
)

type TestRepo interface {
	GetAllTests(req request.QuizRequest) (*[]dto.QuizResponse, *response.Meta, error)
}

type testRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) TestRepo {
	return &testRepo{
		db,
	}
}
