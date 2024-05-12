package auth

import (
	"fmt"

	"github.com/Marvit-Solutions/csw-golang/internal/domain/localmodel/request"
	"github.com/Marvit-Solutions/csw-golang/internal/domain/localmodel/response"
	"github.com/Marvit-Solutions/csw-golang/library/config"
	"github.com/Marvit-Solutions/csw-golang/library/middleware/auth"
	"github.com/Marvit-Solutions/csw-golang/library/struct/model"
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
		"slug": "umum",
	})
	if err != nil {
		return nil, fmt.Errorf("failed to finding role: %v", err)
	}

	class, err := u.classUserRepo.FindOneBy(map[string]interface{}{
		"slug": req.Class,
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
		RoleID:   role.ID,
		Email:    req.Email,
		Password: string(userPassword),
	}

	user, err = u.userRepo.Create(user, tx)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %v", err)
	}

	userDetail := &model.UserDetail{
		ClassUserID: class.ID,
		UserID:      user.ID,
		Name:        req.Name,
		Province:    req.Province,
		Regency:     req.Regency,
		District:    req.District,
		PhoneNumber: req.Phone,
		// ProfilePicture: "account.png",
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

	result := &response.AuthResponse{
		AccessToken: token.AccessToken,
	}

	return result, nil
}
