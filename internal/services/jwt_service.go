package services

import (
	"fmt"
	"time"

	"github.com/edaywalid/chat-app/configs"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type JwtService struct {
	config *configs.Config
}

func NewJwtService(config *configs.Config) *JwtService {
	return &JwtService{config}
}

type TokenPair struct {
	AccessToken  string
	RefreshToken string
}

func (s *JwtService) GenerateTokenPair(userID uuid.UUID) (*TokenPair, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Minute * 15).Unix(),
	})

	fmt.Println("accesstoken  : ", accessToken)

	fmt.Println("signed string : ", s.config.JWTSecret)
	accessTokenString, err := accessToken.SignedString([]byte(s.config.JWTSecret))
	if err != nil {
		return nil, err
	}

	fmt.Println("accessTokenString : ", accessTokenString)
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24 * 7).Unix(),
	})

	refreshTokenString, err := refreshToken.SignedString([]byte(s.config.JWTSecret))
	if err != nil {
		return nil, err
	}

	return &TokenPair{
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
	}, nil
}

func (s *JwtService) ValidateToken(tokenString string) (uint, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(s.config.JWTSecret), nil
	})
	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID := uint(claims["user_id"].(float64))
		return userID, nil
	}

	return 0, jwt.ErrSignatureInvalid
}
