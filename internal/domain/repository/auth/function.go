package auth

import (
	"csw-golang/internal/domain/entity/datastruct"
	"csw-golang/internal/domain/entity/dto"
	"csw-golang/internal/domain/entity/request"
	"fmt"

	"github.com/google/uuid"

	md "csw-golang/internal/delivery/http/middleware"
)

func (ar *authRepo) Register(req request.RegisterRequest) error {
	existingUser := datastruct.Users{}
	if err := ar.db.Where("email = ?", req.Email).First(&existingUser).Error; err == nil {
		return fmt.Errorf("email already exists")
	}

	var role datastruct.Roles
	if err := ar.db.Where("role = ?", "User").First(&role).Error; err != nil {
		return fmt.Errorf("failed to get user role: %v", err)
	}

	newUser := datastruct.Users{
		ID:         uuid.NewString(),
		RoleID:     role.ID,
		Email:      req.Email,
		Password:   req.Password,
		GoogleID:   req.GoogleID,
		FacebookID: req.FacebookID,
		UserDetails: datastruct.UserDetails{
			ID:          uuid.NewString(),
			Name:        req.Name,
			PhoneNumber: req.Phone,
			Addresses: datastruct.Addresses{
				Province:    req.Province,
				RegencyCity: req.Regency,
				SubDistrict: req.District,
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
		return fmt.Errorf("failed to create user: %v", err)
	}

	if err := ar.db.Create(&newAddress).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to create user address: %v", err)
	}

	tx.Commit()

	return nil
}

func (ar *authRepo) Login(req request.LoginRequest) (dto.AuthResponse, error) {
	existingUser := &datastruct.Users{}
	err := ar.db.Preload("UserDetails").Where("email = ?", req.Email).First(&existingUser).Error
	if err != nil {
		return dto.AuthResponse{}, fmt.Errorf("failed to get user: %v", err)
	}

	userAddress := &datastruct.Addresses{}
	err = ar.db.Where("user_detail_id = ?", existingUser.UserDetails.ID).First(&userAddress).Error
	if err != nil {
		return dto.AuthResponse{}, fmt.Errorf("failed to get user address: %v", err)
	}

	token, err := md.CreateToken(existingUser.ID, existingUser.Email)
	if err != nil {
		return dto.AuthResponse{}, fmt.Errorf("failed to create token: %v", err)
	}

	response := &dto.AuthResponse{
		ID:             existingUser.ID,
		Password:       existingUser.Password,
		GoogleID:       existingUser.GoogleID,
		FacebookID:     existingUser.FacebookID,
		Email:          existingUser.Email,
		Name:           existingUser.UserDetails.Name,
		ProfilePicture: existingUser.UserDetails.ProfilePicture,
		Token:          token,
	}

	return *response, nil
}
