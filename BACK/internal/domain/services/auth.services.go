package services

import (
	"context"
	"jasvan/internal/domain/entities"
	"jasvan/internal/repository/auth"
)

type AuthService struct {
	AuthRepository auth.AuthRepository
}

func NewAuthService(authRepository auth.AuthRepository) *AuthService {
	return &AuthService{
		AuthRepository: authRepository,
	}
}

func (s *AuthService) LoginWithEmail(ctx context.Context, email string, password string) (string, error) {
	token, err := s.AuthRepository.LoginWithEmail(ctx, email, password)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *AuthService) GetUserByContext(ctx context.Context) (*entities.User, error) {
	user, err := s.AuthRepository.GetUserByContext(ctx)
	if err != nil {
		return nil, err
	}
	return user, nil
}
