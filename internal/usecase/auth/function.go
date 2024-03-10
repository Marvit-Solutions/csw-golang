package auth

import (
	"csw-golang/internal/domain/request"
	dto "csw-golang/internal/domain/response"
)

func (ac *authUsecase) Register(req request.RegisterRequest) (*dto.AuthResponse, error) {
	response, err := ac.authRepo.Register(req)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (ac *authUsecase) Login(req request.LoginRequest) (*dto.AuthResponse, error) {
	response, err := ac.authRepo.Login(req)
	if err != nil {
		return nil, err
	}

	return response, nil
}
