package auth

import (
	"github.com/Marvit-Solutions/csw-golang/internal/app/repository/auth"
	"github.com/Marvit-Solutions/csw-golang/internal/domain/request"
	"github.com/Marvit-Solutions/csw-golang/internal/domain/response"

	"gorm.io/gorm"
)

type AuthUsecase interface {
	Register(req request.RegisterRequest) (*response.AuthResponse, error)
	Login(req request.LoginRequest) (*response.AuthResponse, error)
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
