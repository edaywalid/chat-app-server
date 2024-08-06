package services

import (
	"errors"
	"time"

	"github.com/edaywalid/chat-app/internal/models"
	"github.com/edaywalid/chat-app/internal/repositories"
	"github.com/edaywalid/chat-app/pkg/utils"
)

type AuthService struct {
	userRepo     *repositories.UserRepository
	jwtService   *JwtService
	emailService *utils.EmailService
}

func NewAuthService(
	userRepo *repositories.UserRepository,
	jwtService *JwtService,
	emailService *utils.EmailService,
) *AuthService {
	return &AuthService{
		userRepo,
		jwtService,
		emailService,
	}
}

func (s *AuthService) Register(username, email, password string) error {
	if !utils.ValidatePassword(password) {
		return errors.New("password not secure")
	}
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return err
	}

	user := &models.User{
		Username:                    username,
		Email:                       email,
		Password:                    string(hashedPassword),
		EmailConfirmationCode:       utils.GenerateRandomCode(6),
		EmailConfirmationCodeExpiry: time.Now().Add(time.Minute * 15),
	}

	err = s.userRepo.Create(user)
	if err != nil {
		return err
	}

	err = s.emailService.SendEmail("Welcome to chat app", "your confirmation code is : "+user.EmailConfirmationCode, user.Email)
	if err != nil {
		return err
	}

	return nil
}

func (s *AuthService) Login(username, password string) (*TokenPair, error) {
	user, err := s.userRepo.FindByUsername(username)
	if err != nil {
		return nil, err
	}

	err = utils.CheckPasswordHash(password, user.Password)
	if err != nil {
		return nil, err
	}

	if !user.IsVerified {

		user.EmailConfirmationCode = utils.GenerateRandomCode(6)
		user.EmailConfirmationCodeExpiry = time.Now().Add(time.Minute * 15)

		err = s.userRepo.Update(user)
		if err != nil {
			return nil, err
		}

		err = s.emailService.SendEmail("Welcome to chat app", "your confirmation code is : "+user.EmailConfirmationCode, user.Email)
		if err != nil {
			return nil, err
		}

		return nil, errors.New("email not verified")
	}

	tokenPair, err := s.jwtService.GenerateTokenPair(user.ID)
	if err != nil {
		return nil, err
	}
	return tokenPair, nil
}

func (s *AuthService) ConfirmEmail(email, code string) error {
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return err
	}

	if user.IsVerified {
		return errors.New("email already verified")
	}

	if user.EmailConfirmationCode == code {
		if user.EmailConfirmationCodeExpiry.After(time.Now()) {
			user.IsVerified = true
			user.EmailConfirmationCode = ""
			user.EmailConfirmationCodeExpiry = time.Time{}

			err := s.userRepo.Update(user)
			if err != nil {
				return err
			}
			return nil
		} else {
			return errors.New("confirmation code Expired")
		}
	} else {
		return errors.New("confirmation code incorrect")
	}
}
