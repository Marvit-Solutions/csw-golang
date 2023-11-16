package auth

import (
	ac "csw-golang/internal/domain/usecase/auth"
)

type AuthHandler struct {
	authUsecase ac.AuthUsecase
}

func New(authUsecase ac.AuthUsecase) *AuthHandler {
	return &AuthHandler{
		authUsecase,
	}
}
