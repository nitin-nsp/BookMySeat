package logic

import (
	"context"
	"errors"
	"ticket-booking-platform/services/user-service/model"
	"ticket-booking-platform/services/user-service/repository"
)

type ProfileLogic struct {
	userRepo repository.UserRepository
}

func NewProfileLogic(userRepo repository.UserRepository) *ProfileLogic {
	return &ProfileLogic{userRepo: userRepo}
}

func (l *ProfileLogic) GetProfile(ctx context.Context, userID int64) (*model.User, error) {
	user, err := l.userRepo.FindUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (l *ProfileLogic) UpdateProfile(ctx context.Context, userID int64, fullName, phone string) (*model.User, error) {
	user, err := l.userRepo.FindUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}

	user.FullName = fullName
	user.Phone = phone

	if err := l.userRepo.UpdateUser(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}
