package user

import (
	"context"
	"fmt"
	"jasvan/internal/domain/entities"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type gormUserRepository struct {
	db *gorm.DB
}

func NewUserrepository(db *gorm.DB) *gormUserRepository {
	return &gormUserRepository{db}
}

func (r *gormUserRepository) CreateUser(ctx context.Context, user *entities.User) error {
	if err := user.Validate(); err != nil {
		return err
	}

	if err := r.db.First(&user, "email = ?", user.Email).Error; err == nil {
		return fmt.Errorf("email already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return r.db.Create(user).Error

}

func (r *gormUserRepository) GetUserByEmail(ctx context.Context, email string) (*entities.User, error) {
	var user entities.User
	if err := r.db.First(&user, "email = ?", email).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *gormUserRepository) GetUserByUsername(ctx context.Context, username string) (*entities.User, error) {
	var user entities.User
	if err := r.db.First(&user, "username = ?", username).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *gormUserRepository) GetAllUser(ctx context.Context) ([]entities.User, error) {
	var users []entities.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
