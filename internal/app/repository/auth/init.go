package auth

import (
	"github.com/Marvit-Solutions/csw-golang/internal/domain/request"
	"github.com/Marvit-Solutions/csw-golang/internal/domain/response"

	"gorm.io/gorm"
)

type AuthRepository interface {
	Register(req request.RegisterRequest) (*response.AuthResponse, error)
	Login(req request.LoginRequest) (*response.AuthResponse, error)
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{db}
}
