package auth

import (
	"context"
	"jasvan/internal/domain/entities"
)

type AuthRepository interface {
	LoginWithEmail(ctx context.Context, email string, password string) (string, error)
	GetUserByContext(ctx context.Context) (*entities.User, error)
}
