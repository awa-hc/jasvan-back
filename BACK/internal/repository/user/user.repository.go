package user

import (
	"context"
	"jasvan/internal/domain/entities"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *entities.User) error
	GetUserByEmail(ctx context.Context, email string) (*entities.User, error)
	GetUserByUsername(ctx context.Context, username string) (*entities.User, error)
	GetAllUser(ctx context.Context) ([]entities.User, error)
}
