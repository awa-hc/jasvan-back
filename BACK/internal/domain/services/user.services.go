package services

import (
	"context"
	"jasvan/internal/domain/entities"
	"jasvan/internal/repository/user"
)

type UserService struct {
	UserRepository user.UserRepository
}

func NewUserService(userRepository user.UserRepository) *UserService {
	return &UserService{
		UserRepository: userRepository,
	}
}

func (s *UserService) CreateUser(ctx context.Context, user *entities.User) error {

	if err := user.Validate(); err != nil {
		return err
	}
	if err := s.UserRepository.CreateUser(ctx, user); err != nil {
		return err
	}
	return nil

}

func (s *UserService) GetUserByEmail(ctx context.Context, email string) (*entities.User, error) {
	user, err := s.UserRepository.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) GetUserByUsername(ctx context.Context, username string) (*entities.User, error) {
	user, err := s.UserRepository.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) GetAllUser(ctx context.Context) ([]entities.User, error) {
	users, err := s.UserRepository.GetAllUser(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}
