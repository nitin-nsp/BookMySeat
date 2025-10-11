package logic

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"

	"ticket-booking-platform/proto/user"
	"ticket-booking-platform/services/user-service/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.LoginRequest) (*user.LoginResponse, error) {
	// Find user by email
	existingUser, err := l.svcCtx.UserModel.FindByEmail(in.Email)
	if err != nil {
		return &user.LoginResponse{
			Code:    401,
			Message: "Invalid credentials",
		}, nil
	}

	// Verify password
	err = bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(in.Password))
	if err != nil {
		return &user.LoginResponse{
			Code:    401,
			Message: "Invalid credentials",
		}, nil
	}

	// Generate JWT token
	token, err := generateToken(existingUser.Id, existingUser.Role, l.svcCtx.Config.JWT.Secret, l.svcCtx.Config.JWT.Expire)
	if err != nil {
		return &user.LoginResponse{
			Code:    500,
			Message: "Failed to generate token",
		}, err
	}

	return &user.LoginResponse{
		Code:    200,
		Message: "Login successful",
		Token:   token,
		User: &user.User{
			Id:        existingUser.Id,
			Email:     existingUser.Email,
			Phone:     existingUser.Phone,
			Name:      existingUser.Name,
			Role:      existingUser.Role,
			CreatedAt: existingUser.CreatedAt,
			UpdatedAt: existingUser.UpdatedAt,
		},
	}, nil
}

func generateToken(userId int64, role, secret string, expire int64) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userId,
		"role":    role,
		"exp":     time.Now().Unix() + expire,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}
