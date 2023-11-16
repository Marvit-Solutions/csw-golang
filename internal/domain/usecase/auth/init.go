package auth

import (
	"csw-golang/internal/domain/entity/dto"
	ar "csw-golang/internal/domain/repository/auth"
)

type AuthUsecase interface {
	Register(user dto.RegisterRequest) error
	Login(user dto.LoginRequest) (dto.AuthResponse, error)
}

type authUsecase struct {
	authRepo ar.AuthRepo
}

func New(authRepo ar.AuthRepo) *authUsecase {
	return &authUsecase{
		authRepo,
	}
}
