package auth

import (
	"csw-golang/internal/domain/entity/dto"

	"gorm.io/gorm"
)

type AuthRepo interface {
	Register(user dto.RegisterRequest) error
	Login(user dto.LoginRequest) (dto.AuthResponse, error)
}

type authRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) AuthRepo {
	return &authRepo{
		db,
	}
}
