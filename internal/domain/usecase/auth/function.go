package auth

import (
	"csw-golang/internal/domain/entity/dto"
	"csw-golang/internal/domain/entity/request"
	pw "csw-golang/internal/domain/helper/password"
	"errors"
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

func (ac *authUsecase) Login(req request.LoginRequest) (dto.AuthResponse, error) {
	response, err := ac.authRepo.Login(req)
	if err != nil {
		//lint:ignore ST1005 Reason for ignoring this linter
		return dto.AuthResponse{}, errors.New("Email atau password salah")
	}

	err = pw.VerifyPassword(response.Password, req.Password)
	if err != nil {
		//lint:ignore ST1005 Reason for ignoring this linter
		return dto.AuthResponse{}, errors.New("Email atau password salah")
	}

	return response, nil
}
