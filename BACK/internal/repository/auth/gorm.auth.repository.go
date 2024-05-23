package auth

import (
	"context"
	"fmt"
	"jasvan/internal/domain/entities"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type gormAuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *gormAuthRepository {
	return &gormAuthRepository{db: db}
}

func (r *gormAuthRepository) LoginWithEmail(ctx context.Context, email string, password string) (string, error) {
	var user entities.User
	if err := r.db.First(&user, "email = ?", email).Error; err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", fmt.Errorf("invalid password")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":    user.Email,
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (r *gormAuthRepository) GetUserByContext(ctx context.Context) (*entities.User, error) {
	user := ctx.Value("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	email := claims["email"].(string)
	var u entities.User
	if err := r.db.First(&u, "email = ?", email).Error; err != nil {
		return nil, err
	}
	return &u, nil
}
