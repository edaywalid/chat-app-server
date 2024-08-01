package services

import (
	"errors"

	"github.com/edaywalid/chat-app/internal/models"
	"github.com/edaywalid/chat-app/internal/repositories"
	"github.com/edaywalid/chat-app/pkg/utils"
)

type AuthService struct {
	userRepo   *repositories.UserRepository
	jwtService *JwtService
}

func NewAuthService(
	userRepo *repositories.UserRepository,
	jwtService *JwtService,
) *AuthService {
	return &AuthService{
		userRepo,
		jwtService,
	}
}
