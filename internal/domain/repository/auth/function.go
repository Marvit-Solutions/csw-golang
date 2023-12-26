package auth

import (
	"csw-golang/internal/domain/entity/datastruct"
	"csw-golang/internal/domain/entity/dto"
	"errors"

	md "csw-golang/internal/delivery/http/middleware"
)

func (ar *authRepo) Register(user dto.RegisterRequest) error {
	existingUser := datastruct.User{}
	if err := ar.db.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		//lint:ignore ST1005 Reason for ignoring this linter
		return errors.New("Email already exists")
	}

	var role datastruct.Role
	if err := ar.db.Where("name = Admin").First(&role).Error; err != nil {
		return err
	}

	newUser := datastruct.User{
		RoleId:     role.ID,
		Email:      user.Email,
		Username:   user.Username,
		Password:   user.Password,
		GoogleId:   user.GoogleId,
		FacebookId: user.FacebookId,
		UserDetail: datastruct.UserDetail{
			Name:  user.Name,
			Phone: user.Phone,
		},
	}

	if err := ar.db.Create(&newUser).Error; err != nil {
		return err
	}

	return nil
}

func (ar *authRepo) Login(user dto.LoginRequest) (dto.AuthResponse, error) {
	existingUser := &datastruct.User{}
	err := ar.db.Preload("UserDetail").Where("email = ?", user.Email).First(&existingUser).Error
	if err != nil {
		//lint:ignore ST1005 Reason for ignoring this linter
		return dto.AuthResponse{}, errors.New("Record Not Found")
	}

	token, err := md.CreateToken(existingUser.ID, existingUser.Email)
	if err != nil {
		return dto.AuthResponse{}, err
	}

	response := &dto.AuthResponse{
		ID:         existingUser.ID,
		Password:   existingUser.Password,
		GoogleId:   existingUser.GoogleId,
		FacebookId: existingUser.FacebookId,
		Email:      existingUser.Email,
		Username:   existingUser.Username,
		Name:       existingUser.UserDetail.Name,
		Phone:      existingUser.UserDetail.Phone,
		Token:      token,
	}

	return *response, nil
}
