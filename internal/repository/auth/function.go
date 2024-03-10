package auth

import (
	"csw-golang/internal/domain/entity/datastruct"
	"csw-golang/internal/domain/entity/dto"
	"csw-golang/internal/domain/entity/request"
	"csw-golang/internal/domain/helper/password"
	"fmt"
	"os"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	jwt "csw-golang/internal/delivery/http/middleware/jwt"
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

	cswAuth := jwt.NewCswAuth([]byte(os.Getenv("SECRET_KEY")))
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

	tokenStruct := jwt.TokenStructure{
		UserID: newUser.ID,
		Email:  newUser.Email,
	}
	token, err := cswAuth.GenerateToken(tokenStruct)
	if err != nil {
		return nil, fmt.Errorf("failed to create token: %v", err)
	}

	response := &dto.AuthResponse{
		AccessToken: token.AccessToken,
	}

	return response, nil
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

	cswAuth := jwt.NewCswAuth([]byte(os.Getenv("SECRET_KEY")))

	tokenStruct := jwt.TokenStructure{
		UserID: user.ID,
		Email:  user.Email,
	}
	token, err := cswAuth.GenerateToken(tokenStruct)
	if err != nil {
		return nil, fmt.Errorf("failed to create token: %v", err)
	}

	response := &dto.AuthResponse{
		AccessToken: token.AccessToken,
	}

	return response, nil
}
