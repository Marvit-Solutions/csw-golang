package auth

import (
	"fmt"

	"github.com/Marvit-Solutions/csw-golang/internal/domain/model/request"
	"github.com/Marvit-Solutions/csw-golang/internal/domain/model/response"
	"github.com/Marvit-Solutions/csw-golang/library/config"
	"github.com/Marvit-Solutions/csw-golang/library/helper"
	"github.com/Marvit-Solutions/csw-golang/library/middleware/auth"
)

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

	result := &response.AuthResponse{
		AccessToken: token.AccessToken,
	}

	return result, nil
}
