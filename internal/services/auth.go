package services

import (
	"context"
	"errors"

	"github.com/idkOybek/internal/models"
	"github.com/idkOybek/internal/repository"
	"github.com/idkOybek/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo *repository.UserRepository
}

func NewAuthService(userRepo *repository.UserRepository) *AuthService {
	return &AuthService{userRepo: userRepo}
}

func (s *AuthService) RegisterUser(ctx context.Context, user models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	return s.userRepo.Create(ctx, &user)
}

func (s *AuthService) AuthenticateUser(ctx context.Context, username, password string) (*models.User, string, error) {
	user, err := s.userRepo.GetByUsername(ctx, username)
	if err != nil {
		return nil, "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, "", errors.New("invalid password")
	}

	token, err := utils.GenerateJWT(user)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}
