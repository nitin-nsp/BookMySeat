package logic

import (
	"context"
	"errors"
	"ticket-booking-platform/services/user-service/model"
	"ticket-booking-platform/services/user-service/repository"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type LoginLogic struct {
	userRepo  repository.UserRepository
	jwtSecret string
}

func NewLoginLogic(userRepo repository.UserRepository, jwtSecret string) *LoginLogic {
	return &LoginLogic{
		userRepo:  userRepo,
		jwtSecret: jwtSecret,
	}
}

func (l *LoginLogic) Login(ctx context.Context, email, password string) (string, *model.User, error) {
	user, err := l.userRepo.FindUserByEmail(ctx, email)
	if err != nil {
		return "", nil, err
	}
	if user == nil {
		return "", nil, errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return "", nil, errors.New("invalid credentials")
	}

	token, err := l.generateToken(user.ID, user.Email)
	if err != nil {
		return "", nil, err
	}

	return token, user, nil
}

func (l *LoginLogic) generateToken(userID int64, email string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"email":   email,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(l.jwtSecret))
}
