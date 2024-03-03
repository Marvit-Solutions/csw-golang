package auth

import (
	"csw-golang/internal/domain/entity/datastruct"
	"csw-golang/internal/domain/entity/dto"
	"errors"

	"github.com/google/uuid"

	md "csw-golang/internal/delivery/http/middleware"
)

func (ar *authRepo) Register(user dto.RegisterRequest) error {
	existingUser := datastruct.Users{}
	if err := ar.db.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		//lint:ignore ST1005 Reason for ignoring this linter
		return errors.New("Email already exists")
	}

	var role datastruct.Roles
	if err := ar.db.Where("role = ?", "User").First(&role).Error; err != nil {
		return err
	}

	newUser := datastruct.Users{
		ID:         uuid.NewString(),
		RoleID:     role.ID,
		Email:      user.Email,
		Password:   user.Password,
		GoogleID:   user.GoogleID,
		FacebookID: user.FacebookID,
		UserDetails: datastruct.UserDetails{
			ID:          uuid.NewString(),
			Name:        user.Name,
			PhoneNumber: user.Phone,
			Addresses: datastruct.Addresses{
				Province:    user.Province,
				RegencyCity: user.Regency,
				SubDistrict: user.District,
			},
		},
	}

	newAddress := datastruct.Addresses{
		ID:           uuid.NewString(),
		UserDetailID: newUser.UserDetails.ID,
		Province:     newUser.UserDetails.Addresses.Province,
		RegencyCity:  newUser.UserDetails.Addresses.RegencyCity,
		SubDistrict:  newUser.UserDetails.Addresses.SubDistrict,
	}

	tx := ar.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := ar.db.Create(&newUser).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := ar.db.Create(&newAddress).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}

func (ar *authRepo) Login(user dto.LoginRequest) (dto.AuthResponse, error) {
	existingUser := &datastruct.Users{}
	err := ar.db.Preload("UserDetails").Where("email = ?", user.Email).First(&existingUser).Error
	if err != nil {
		return dto.AuthResponse{}, err
	}

	userAddress := &datastruct.Addresses{}
	err = ar.db.Where("user_detail_id = ?", existingUser.UserDetails.ID).First(&userAddress).Error
	if err != nil {
		return dto.AuthResponse{}, err
	}

	userRole := &datastruct.Roles{}
	err = ar.db.Where("id = ?", existingUser.RoleID).First(&userRole).Error
	if err != nil {
		return dto.AuthResponse{}, err
	}

	token, err := md.CreateToken(existingUser.ID, existingUser.Email)
	if err != nil {
		return dto.AuthResponse{}, err
	}

	response := &dto.AuthResponse{
		ID:             existingUser.ID,
		Password:       existingUser.Password,
		GoogleID:       existingUser.GoogleID,
		FacebookID:     existingUser.FacebookID,
		Email:          existingUser.Email,
		Name:           existingUser.UserDetails.Name,
		Role:           userRole.Role,
		ProfilePicture: existingUser.UserDetails.ProfilePicture,
		Token:          token,
	}

	return *response, nil
}
