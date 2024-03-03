package auth

import (
	"csw-golang/internal/domain/entity/dto"
	"csw-golang/internal/domain/entity/request"
	ar "csw-golang/internal/domain/repository/auth"
)

type AuthUsecase interface {
	Register(req request.RegisterRequest) error
	Login(req request.LoginRequest) (*dto.AuthResponse, error)
}

type authUsecase struct {
	authRepo ar.AuthRepo
}

func New(authRepo ar.AuthRepo) *authUsecase {
	return &authUsecase{
		authRepo,
	}
}
