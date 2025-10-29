package logic

import (
	"context"
	"errors"
	"regexp"
	"ticket-booking-platform/services/user-service/model"
	"ticket-booking-platform/services/user-service/repository"

	"golang.org/x/crypto/bcrypt"
)

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

type RegisterLogic struct {
	userRepo repository.UserRepository
}

func NewRegisterLogic(userRepo repository.UserRepository) *RegisterLogic {
	return &RegisterLogic{userRepo: userRepo}
}

func (l *RegisterLogic) Register(ctx context.Context, email, password, fullName, phone string) (*model.User, error) {
	if !emailRegex.MatchString(email) {
		return nil, errors.New("invalid email format")
	}

	if len(password) < 8 {
		return nil, errors.New("password must be at least 8 characters")
	}

	existing, err := l.userRepo.FindUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if existing != nil {
		return nil, errors.New("email already exists")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &model.User{
		Email:        email,
		PasswordHash: string(hash),
		FullName:     fullName,
		Phone:        phone,
	}

	if err := l.userRepo.CreateUser(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}
