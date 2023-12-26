package auth

import (
	"csw-golang/internal/domain/entity/dto"
	pw "csw-golang/internal/domain/helper/password"
	"errors"
)

func (ac *authUsecase) Register(user dto.RegisterRequest) error {
	hashedPassword, err := pw.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)

	err = ac.authRepo.Register(user)
	if err != nil {
		return err
	}

	return nil
}

func (ac *authUsecase) Login(user dto.LoginRequest) (dto.AuthResponse, error) {
	response, err := ac.authRepo.Login(user)
	if err != nil {
		//lint:ignore ST1005 Reason for ignoring this linter
		return dto.AuthResponse{}, errors.New("Email atau password salah")
	}

	err = pw.VerifyPassword(response.Password, user.Password)
	if err != nil {
		//lint:ignore ST1005 Reason for ignoring this linter
		return dto.AuthResponse{}, errors.New("Email atau password salah")
	}

	return response, nil
}
