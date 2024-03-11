package auth

import (
	"github.com/Marvit-Solutions/csw-golang/internal/domain/request"
	"github.com/Marvit-Solutions/csw-golang/internal/domain/response"
)

func (ac *authUsecase) Register(req request.RegisterRequest) (*response.AuthResponse, error) {
	data, err := ac.authRepo.Register(req)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (ac *authUsecase) Login(req request.LoginRequest) (*response.AuthResponse, error) {
	data, err := ac.authRepo.Login(req)
	if err != nil {
		return nil, err
	}

	return data, nil
}
