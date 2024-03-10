package auth

import (
	"csw-golang/internal/domain/request"
	dto "csw-golang/internal/domain/response"
	"csw-golang/internal/repository/auth"

	"gorm.io/gorm"
)

type AuthUsecase interface {
	Register(req request.RegisterRequest) (*dto.AuthResponse, error)
	Login(req request.LoginRequest) (*dto.AuthResponse, error)
}

type authUsecase struct {
	authRepo auth.AuthRepository
}

func NewAuthUsecase(
	db *gorm.DB,
) AuthUsecase {
	return &authUsecase{
		authRepo: auth.NewAuthRepository(db),
	}
}
