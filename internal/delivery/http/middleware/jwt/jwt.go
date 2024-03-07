package jwt

import (
	"fmt"
	"net/http"
	"os"
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
	GenerateToken(data TokenStructure) (*TokenResponse, error)
}

func NewCswAuth(signature []byte) CswAuth {
	return &cswAuth{signature}
}

const (
	EXPIRED_IN = time.Hour * (24 * 90) // 90 days
)

func NewMiddlewareConfig() error {
	InitJWTMiddlewareCustom([]byte(os.Getenv("SECRET_KEY")), jwt.SigningMethodHS512)
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
	token := jwt.New(jwt.SigningMethodHS512)
	claims := token.Claims.(jwt.MapClaims)

	expiredIn := EXPIRED_IN
	expiredAt := time.Now().Add(EXPIRED_IN)

	myCrypt, err := bcrypt.GenerateFromPassword([]byte(os.Getenv("SECRET_KEY")), 8)
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
		return nil, err
	}
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})
	if err != nil || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}
	claimValue, exists := claims[key]
	if !exists {
		return nil, fmt.Errorf("claim not found")
	}
	return claimValue, nil
}

func GetAuthenticatedUser(r *http.Request) (string, error) {
	userID, err := ExtractToken(r, "user_id")
	if err != nil {
		return "", err
	}
	return userID.(string), nil
}
