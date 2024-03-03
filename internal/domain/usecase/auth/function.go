package auth

import (
	"csw-golang/internal/domain/entity/dto"
	"csw-golang/internal/domain/entity/request"
	pw "csw-golang/internal/domain/helper/password"
)

func (ac *authUsecase) Register(req request.RegisterRequest) error {
	hashedPassword, err := pw.HashPassword(req.Password)
	if err != nil {
		return err
	}

	req.Password = string(hashedPassword)

	err = ac.authRepo.Register(req)
	if err != nil {
		return err
	}

	return nil
}

func (ac *authUsecase) Login(req request.LoginRequest) (*dto.AuthResponse, error) {
	response, err := ac.authRepo.Login(req)
	if err != nil {
		return nil, err
	}

	return response, nil
}
