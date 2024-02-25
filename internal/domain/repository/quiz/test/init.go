package tests

import (
	"csw-golang/internal/domain/entity/dto"

	"gorm.io/gorm"
)

type TestRepo interface {
	GetAllTests() ([]dto.QuizResponse, error)
}

type testRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) TestRepo {
	return &testRepo{
		db,
	}
}
