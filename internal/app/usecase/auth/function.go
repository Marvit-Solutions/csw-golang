package auth

import (
	"fmt"

	"github.com/Marvit-Solutions/csw-golang/internal/domain/model/request"
	"github.com/Marvit-Solutions/csw-golang/internal/domain/model/response"
	"github.com/Marvit-Solutions/csw-golang/library/config"
	"github.com/Marvit-Solutions/csw-golang/library/helper"
	"github.com/Marvit-Solutions/csw-golang/library/middleware/auth"
	"github.com/Marvit-Solutions/csw-golang/library/struct/model"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (u *usecase) Register(req request.RegisterRequest) (*response.AuthResponse, error) {
	user, _ := u.userRepo.FindOneBy(map[string]interface{}{
		"email": req.Email,
	})
	if user != nil {
		return nil, fmt.Errorf("email already registered")
	}

	role, err := u.roleRepo.FindOneBy(map[string]interface{}{
		"name": "user",
	})
	if err != nil {
		return nil, fmt.Errorf("failed to finding role: %v", err)
	}

	class, err := u.classUserRepo.FindOneBy(map[string]interface{}{
		"name": req.Class,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to finding class: %v", err)
	}

	tx := u.db.Begin()
	defer tx.Rollback()

	bytePassword := []byte(req.Password)
	userPassword, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed hashing password: %v", err)
	}

	user = &model.User{
		ID:       uuid.NewString(),
		RoleID:   role.ID,
		Email:    req.Email,
		Password: string(userPassword),
	}

	user, err = u.userRepo.Create(user, tx)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %v", err)
	}

	userDetail := &model.UserDetail{
		ID:             uuid.NewString(),
		ClassUserID:    class.ID,
		UserID:         user.ID,
		Name:           req.Name,
		PhoneNumber:    req.Phone,
		ProfilePicture: "./assets/img/users/profiles/account.png",
	}

	_, err = u.userDetailRepo.Create(userDetail, tx)
	if err != nil {
		return nil, fmt.Errorf("failed to create user detail: %v", err)
	}

	conf := config.NewConfig()
	cswAuth := auth.NewCswAuth([]byte(conf.GetString("app.signature")))
	tokenStruct := auth.TokenStructure{
		UserID: user.ID,
		Email:  req.Email,
	}

	token, err := cswAuth.GenerateToken(tokenStruct)
	if err != nil {
		return nil, fmt.Errorf("failed generating token: %v", err)
	}

	tx.Commit()

	resp := &response.AuthResponse{
		AccessToken: token.AccessToken,
	}

	return resp, nil
}

func (u *usecase) Login(req request.LoginRequest) (*response.AuthResponse, error) {
	user, err := u.userRepo.FindOneBy(map[string]interface{}{
		"email": req.Email,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to finding user: %v", err)
	}

	if !helper.ComparePasswords(user.Password, req.Password) {
		return nil, fmt.Errorf("invalid email or password")
	}

	conf := config.NewConfig()
	cswAuth := auth.NewCswAuth([]byte(conf.GetString("app.signature")))
	tokenStruct := auth.TokenStructure{
		UserID: user.ID,
		Email:  user.Email,
	}

	token, err := cswAuth.GenerateToken(tokenStruct)
	if err != nil {
		return nil, fmt.Errorf("failed generating token: %v", err)
	}

	resp := &response.AuthResponse{
		AccessToken: token.AccessToken,
	}
	return resp, nil
}
