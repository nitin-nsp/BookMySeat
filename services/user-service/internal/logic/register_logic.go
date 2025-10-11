package logic

import (
	"context"
	"fmt"
	"regexp"

	"golang.org/x/crypto/bcrypt"

	"ticket-booking-platform/proto/user"
	"ticket-booking-platform/services/user-service/internal/model"
	"ticket-booking-platform/services/user-service/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user.RegisterRequest) (*user.RegisterResponse, error) {
	// Validate email
	if !isValidEmail(in.Email) {
		return &user.RegisterResponse{
			Code:    400,
			Message: "Invalid email format",
		}, nil
	}

	// Check if user exists
	existingUser, _ := l.svcCtx.UserModel.FindByEmail(in.Email)
	if existingUser != nil {
		return &user.RegisterResponse{
			Code:    409,
			Message: "Email already registered",
		}, nil
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		return &user.RegisterResponse{
			Code:    500,
			Message: "Failed to process password",
		}, err
	}

	// Create user
	newUser := &model.User{
		Email:    in.Email,
		Password: string(hashedPassword),
		Phone:    in.Phone,
		Name:     in.Name,
		Role:     "user",
	}

	userId, err := l.svcCtx.UserModel.Insert(newUser)
	if err != nil {
		return &user.RegisterResponse{
			Code:    500,
			Message: "Failed to create user",
		}, err
	}

	return &user.RegisterResponse{
		Code:    200,
		Message: "User registered successfully",
		User: &user.User{
			Id:        userId,
			Email:     newUser.Email,
			Phone:     newUser.Phone,
			Name:      newUser.Name,
			Role:      newUser.Role,
			CreatedAt: newUser.CreatedAt,
			UpdatedAt: newUser.UpdatedAt,
		},
	}, nil
}

func isValidEmail(email string) bool {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	matched, _ := regexp.MatchString(pattern, email)
	return matched
}
