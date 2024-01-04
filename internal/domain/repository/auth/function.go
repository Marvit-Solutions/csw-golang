package auth

import (
	"csw-golang/internal/domain/entity/datastruct"
	"csw-golang/internal/domain/entity/dto"
	"errors"

	"github.com/google/uuid"

	md "csw-golang/internal/delivery/http/middleware"
)

func (ar *authRepo) Register(user dto.RegisterRequest) error {
	existingUser := datastruct.User{}
	if err := ar.db.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		//lint:ignore ST1005 Reason for ignoring this linter
		return errors.New("Email already exists")
	}

	var role datastruct.Role
	if err := ar.db.Where("role = ?", "User").First(&role).Error; err != nil {
		return err
	}

	newUser := datastruct.User{
		ID:         uuid.NewString(),
		RoleID:     role.ID,
		Email:      user.Email,
		Password:   user.Password,
		GoogleID:   user.GoogleId,
		FacebookID: user.FacebookId,
		UserDetail: datastruct.UserDetail{
			ID:      uuid.NewString(),
			Nama:    user.Nama,
			Telepon: user.Telepon,
			Alamat: datastruct.Address{
				Provinsi:  user.Provinsi,
				Kabupaten: user.Kabupaten,
				Kecamatan: user.Kecamatan,
			},
		},
	}

	newAddress := datastruct.Address{
		ID:           uuid.NewString(),
		UserDetailID: newUser.UserDetail.ID,
		Provinsi:     newUser.UserDetail.Alamat.Provinsi,
		Kabupaten:    newUser.UserDetail.Alamat.Kabupaten,
		Kecamatan:    newUser.UserDetail.Alamat.Kecamatan,
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
	existingUser := &datastruct.User{}
	err := ar.db.Preload("UserDetail").Where("email = ?", user.Email).First(&existingUser).Error
	if err != nil {
		return dto.AuthResponse{}, err
	}

	userAddress := &datastruct.Address{}
	err = ar.db.Where("user_detail_id = ?", existingUser.UserDetail.ID).First(&userAddress).Error
	if err != nil {
		return dto.AuthResponse{}, err
	}

	userRole := &datastruct.Role{}
	err = ar.db.Where("id = ?", existingUser.RoleID).First(&userRole).Error
	if err != nil {
		return dto.AuthResponse{}, err
	}

	token, err := md.CreateToken(existingUser.ID, existingUser.Email)
	if err != nil {
		return dto.AuthResponse{}, err
	}

	response := &dto.AuthResponse{
		ID:         existingUser.ID,
		Password:   existingUser.Password,
		GoogleId:   existingUser.GoogleID,
		FacebookId: existingUser.FacebookID,
		Email:      existingUser.Email,
		Nama:       existingUser.UserDetail.Nama,
		Role:       userRole.Role,
		Telepon:    existingUser.UserDetail.Telepon,
		FotoProfil: existingUser.UserDetail.FotoProfil,
		Alamat: struct {
			Provinsi  string "json:\"Provinsi\" form:\"Provinsi\""
			Kabupaten string "json:\"Kabupaten\" form:\"Kabupaten\""
			Kecamatan string "json:\"Kecamatan\" form:\"Kecamatan\""
		}{
			Provinsi:  userAddress.Provinsi,
			Kabupaten: userAddress.Kabupaten,
			Kecamatan: userAddress.Kecamatan,
		},
		Token: token,
	}

	return *response, nil
}
