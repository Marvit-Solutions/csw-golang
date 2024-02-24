package tests

import (
	"csw-golang/internal/domain/entity/datastruct"

	"gorm.io/gorm"
)

type TestRepo interface {
	GetAllTests() ([]datastruct.Modules, error)
}

type testRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) TestRepo {
	return &testRepo{
		db,
	}
}
