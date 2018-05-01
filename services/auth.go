package services

import (
	"encoding/base64"
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	cc "github.com/smithaitufe/courses/context"
	"github.com/smithaitufe/courses/models"
)

type AuthService struct {
	appName       string
	jwtSecret     string
	jwtExpiration int
}

func NewAuthService(config *cc.Config) *AuthService {
	return &AuthService{jwtSecret: config.TokenSigning.Secret, jwtExpiration: config.TokenSigning.Expiration, appName: config.App.Name}
}

func (a *AuthService) GenerateToken(user *models.User) (*string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":         base64.StdEncoding.EncodeToString([]byte(user.ID)),
		"first_name": user.FirstName,
		"last_name":  user.LastName,
		"created_at": time.Now(),
		"exp":        time.Now().Add(time.Second * time.Duration(a.jwtExpiration)).Unix(),
		"issuer":     a.appName,
	})

	signedToken, err := token.SignedString([]byte(a.jwtSecret))
	return &signedToken, err
}

func (a *AuthService) GenerateRefreshToken(user *models.User) (*string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":         base64.StdEncoding.EncodeToString([]byte(user.ID)),
		"created_at": time.Now(),
		"exp":        time.Now().Add(time.Second * time.Duration(a.jwtExpiration)).Unix(),
		"issuer":     a.appName,
	})

	signedToken, err := token.SignedString([]byte(a.jwtSecret))
	return &signedToken, err
}

func (a *AuthService) ValidateToken(tokenString *string) (*jwt.Token, error) {
	verifiedToken, err := jwt.Parse(*tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Failed to verify token. Reason: %v", token.Header["alg"])
		}
		return []byte(a.jwtSecret), nil
	})
	return verifiedToken, err
}
