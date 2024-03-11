package auth

import (
	"fmt"
	"os"

	"github.com/Marvit-Solutions/csw-golang/internal/domain/datastruct"
	"github.com/Marvit-Solutions/csw-golang/internal/domain/request"
	dto "github.com/Marvit-Solutions/csw-golang/internal/domain/response"
	"github.com/Marvit-Solutions/csw-golang/library/helper/password"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"github.com/Marvit-Solutions/csw-golang/library/middleware/auth"
)

func (ar *authRepository) Register(req request.RegisterRequest) (*dto.AuthResponse, error) {
	var existingUser datastruct.Users
	if err := ar.db.Where("email = ?", req.Email).First(&existingUser).Error; err == nil {
		return nil, fmt.Errorf("email already exists")
	}

	var role datastruct.Roles
	if err := ar.db.Where("role = ?", "User").First(&role).Error; err != nil {
		return nil, fmt.Errorf("failed to get user role: %v", err)
	}

	cswAuth := auth.NewCswAuth([]byte(os.Getenv("SECRET_KEY")))
	bytePassword := []byte(req.Password)
	userPassword, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed hashing password: %v", err)
	}

	newUser := datastruct.Users{
		ID:         uuid.NewString(),
		RoleID:     role.ID,
		Email:      req.Email,
		Password:   string(userPassword),
		GoogleID:   req.GoogleID,
		FacebookID: req.FacebookID,
		UserDetails: datastruct.UserDetails{
			ID:          uuid.NewString(),
			Name:        req.Name,
			PhoneNumber: req.Phone,
			Class:       req.Class,
			Addresses: datastruct.Addresses{
				ID:          uuid.NewString(),
				Province:    req.Province,
				RegencyCity: req.Regency,
				SubDistrict: req.District,
			},
		},
	}

	tx := ar.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := ar.db.Create(&newUser).Error; err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("failed to create user: %v", err)
	}

	tx.Commit()

	tokenStruct := auth.TokenStructure{
		UserID: newUser.ID,
		Email:  newUser.Email,
	}
	token, err := cswAuth.GenerateToken(tokenStruct)
	if err != nil {
		return nil, fmt.Errorf("failed to create token: %v", err)
	}

	data := &dto.AuthResponse{
		AccessToken: token.AccessToken,
	}

	return data, nil
}

func (ar *authRepository) Login(req request.LoginRequest) (*dto.AuthResponse, error) {
	var user *datastruct.Users
	err := ar.db.Preload("UserDetails").Where("email = ?", req.Email).First(&user).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %v", err)
	}

	if !password.ComparePasswords(user.Password, req.Password) {
		return nil, fmt.Errorf("invalid email or password")
	}

	cswAuth := auth.NewCswAuth([]byte(os.Getenv("SECRET_KEY")))

	tokenStruct := auth.TokenStructure{
		UserID: user.ID,
		Email:  user.Email,
	}
	token, err := cswAuth.GenerateToken(tokenStruct)
	if err != nil {
		return nil, fmt.Errorf("failed to create token: %v", err)
	}

	data := &dto.AuthResponse{
		AccessToken: token.AccessToken,
	}

	return data, nil
}
