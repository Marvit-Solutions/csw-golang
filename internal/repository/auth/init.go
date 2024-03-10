package auth

import (
	"csw-golang/internal/domain/request"
	dto "csw-golang/internal/domain/response"

	"gorm.io/gorm"
)

type AuthRepository interface {
	Register(req request.RegisterRequest) (*dto.AuthResponse, error)
	Login(req request.LoginRequest) (*dto.AuthResponse, error)
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{db}
}
