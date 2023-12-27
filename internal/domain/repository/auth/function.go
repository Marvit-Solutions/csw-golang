package auth

import (
	"csw-golang/internal/domain/entity/datastruct"
	"csw-golang/internal/domain/entity/dto"
	"csw-golang/internal/domain/helper/time"
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
			Nama:    user.Nama,
			Telepon: user.Telepon,
			Alamat:  datastruct.Address{},
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

	userAddress := &datastruct.Address{}
	err = ar.db.Where("user_detail_id = ?", existingUser.UserDetail.ID).First(&userAddress).Error
	if err != nil {
		//lint:ignore ST1005 Reason for ignoring this linter
		return dto.AuthResponse{}, errors.New("Record Not Found")
	}

	userRole := &datastruct.Role{}
	err = ar.db.Where("id = ?", existingUser.RoleId).First(&userRole).Error
	if err != nil {
		//lint:ignore ST1005 Reason for ignoring this linter
		return dto.AuthResponse{}, errors.New("Record Not Found")
	}

	token, err := md.CreateToken(existingUser.ID, existingUser.Email)
	if err != nil {
		return dto.AuthResponse{}, err
	}

	response := &dto.AuthResponse{
		ID:           existingUser.ID,
		Password:     existingUser.Password,
		GoogleId:     existingUser.GoogleId,
		FacebookId:   existingUser.FacebookId,
		Email:        existingUser.Email,
		Username:     existingUser.Username,
		Nama:         existingUser.UserDetail.Nama,
		Role:         userRole.Role,
		Telepon:      existingUser.UserDetail.Telepon,
		FotoProfil:   existingUser.UserDetail.FotoProfil,
		TanggalLahir: time.ConvertTimeFormat(existingUser.UserDetail.TanggalLahir),
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
