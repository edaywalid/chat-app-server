package services

import (
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

func (s *JwtService) generateToken(userID uuid.UUID, live_time time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(live_time).Unix(),
	})

	tokenString, err := token.SignedString([]byte(s.config.JWTSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *JwtService) GenerateTokenPair(userID uuid.UUID) (*TokenPair, error) {
	accessToken, err := s.generateToken(userID, time.Minute*15)
	if err != nil {
		return nil, err
	}

	refreshToken, err := s.generateToken(userID, time.Hour*24*7)
	if err != nil {
		return nil, err
	}
	return &TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *JwtService) ValidateToken(tokenString string) (uuid.UUID, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.config.JWTSecret), nil
	})
	if err != nil {
		return uuid.Nil, err
	}

	if !token.Valid {
		return uuid.Nil, jwt.ValidationError{
			Errors: jwt.ValidationErrorExpired,
		}
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return uuid.Nil, jwt.ValidationError{
			Errors: jwt.ValidationErrorClaimsInvalid,
		}
	}

	userID, ok := claims["user_id"].(string)
	if !ok {
		return uuid.Nil, jwt.ValidationError{
			Errors: jwt.ValidationErrorClaimsInvalid,
		}
	}

	return uuid.Parse(userID)
}

func (s *JwtService) RefreshToken(refreshToken string) (string, error) {
	userID, err := s.ValidateToken(refreshToken)
	if err != nil {
		return "", err
	}

	token, err := s.generateToken(userID, time.Minute*15)
	if err != nil {
		return "", err
	}

	return token, nil
}
