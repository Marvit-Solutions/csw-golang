package auth

import (
	"csw-golang/internal/domain/entity/dto"
	"csw-golang/internal/domain/entity/request"

	"gorm.io/gorm"
)

type AuthRepo interface {
	Register(req request.RegisterRequest) (*dto.AuthResponse, error)
	Login(req request.LoginRequest) (*dto.AuthResponse, error)
}

type authRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) AuthRepo {
	return &authRepo{
		db,
	}
}
