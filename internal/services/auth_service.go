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

func (s *AuthService) Register(username, email, password string) (*TokenPair, error) {
	if !utils.ValidatePassword(password) {
		return nil, errors.New("password not secure")
	}
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Username: username,
		Email:    email,
		Password: string(hashedPassword),
	}

	err = s.userRepo.Create(user)
	if err != nil {
		return nil, err
	}

	tokenPair, err := s.jwtService.GenerateTokenPair(user.ID)
	if err != nil {
		return nil, err
	}

	return tokenPair, nil
}

func (s *AuthService) Login(username, password string) error {
	user, err := s.userRepo.FindByUsername(username)
	if err != nil {
		return err
	}

	err = utils.CheckPasswordHash(password, user.Password)
	if err != nil {
		return err
	}

	return nil
}
