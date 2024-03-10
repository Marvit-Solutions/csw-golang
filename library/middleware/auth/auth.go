package auth

import (
	"csw-golang/library/config"
	"fmt"
	"net/http"
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/form3tech-oss/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

var jwtMiddleware *jwtmiddleware.JWTMiddleware
var signingKey []byte

type TokenStructure struct {
	UserID string
	Email  string
}

type TokenResponse struct {
	AccessToken string  `json:"AccessToken"`
	TokenType   string  `json:"TokenType"`
	ExpiredIn   float64 `json:"ExpiredIn"`
	ExpiredAt   int64   `json:"ExpiredAt"`
}

type cswAuth struct {
	signature []byte
}

type CswAuth interface {
	GenerateToken(data TokenStructure) (response *TokenResponse, err error)
}

func NewCswAuth(signature []byte) CswAuth {
	return &cswAuth{signature}
}

const (
	EXPIRED_IN = time.Hour * (24 * 90) // 90 days
)

func NewMiddlewareConfig(cfg config.Config) error {
	signature := cfg.GetString("app.signature")
	InitJWTMiddlewareCustom([]byte(signature), jwt.SigningMethodHS512)

	return nil
}

func InitJWTMiddlewareCustom(secret []byte, signingMethod jwt.SigningMethod) {
	signingKey = secret
	jwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return signingKey, nil
		},
		SigningMethod: signingMethod,
	})
}

func (cAuth *cswAuth) GenerateToken(data TokenStructure) (response *TokenResponse, err error) {
	cfg := config.NewConfig()

	token := jwt.New(jwt.SigningMethodHS512)
	claims := token.Claims.(jwt.MapClaims)

	expiredIn := EXPIRED_IN
	expiredAt := time.Now().Add(EXPIRED_IN)

	myCrypt, err := bcrypt.GenerateFromPassword([]byte(cfg.GetString("app.signature")), 8)
	if err != nil {
		return nil, fmt.Errorf("failed generating password: %v", err)
	}

	claims["user_id"] = data.UserID
	claims["email"] = data.Email
	claims["hash"] = string(myCrypt)
	claims["exp"] = expiredIn

	tokenString, err := token.SignedString(cAuth.signature)
	if err != nil {
		return nil, fmt.Errorf("failed signing string: %v", err)
	}

	response = new(TokenResponse)
	response.AccessToken = tokenString
	response.TokenType = "Bearer"
	response.ExpiredAt = expiredAt.Unix()
	response.ExpiredIn = expiredIn.Seconds()

	return response, nil
}

func ExtractToken(r *http.Request, key string) (interface{}, error) {
	tokenStr, err := jwtMiddleware.Options.Extractor(r)
	if err != nil {
		return "", err
	}
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims[key], nil
	} else {
		return "", nil
	}
}

func GetAuthenticatedUser(r *http.Request) (string, error) {
	userID, err := ExtractToken(r, "user_id")
	if err != nil {
		return "", err
	}
	return userID.(string), nil
}
