package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"ticket-booking-platform/proto/user"
	"ticket-booking-platform/services/user-service/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProfileLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProfileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProfileLogic {
	return &GetProfileLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetProfileLogic) GetProfile(in *user.GetProfileRequest) (*user.GetProfileResponse, error) {
	// Try cache first
	cacheKey := fmt.Sprintf("user:profile:%d", in.UserId)
	cached, err := l.svcCtx.Redis.Get(cacheKey)
	if err == nil && cached != "" {
		var cachedUser user.User
		if json.Unmarshal([]byte(cached), &cachedUser) == nil {
			return &user.GetProfileResponse{
				Code:    200,
				Message: "Success",
				User:    &cachedUser,
			}, nil
		}
	}

	// Get from database
	existingUser, err := l.svcCtx.UserModel.FindById(in.UserId)
	if err != nil {
		return &user.GetProfileResponse{
			Code:    404,
			Message: "User not found",
		}, nil
	}

	userProto := &user.User{
		Id:        existingUser.Id,
		Email:     existingUser.Email,
		Phone:     existingUser.Phone,
		Name:      existingUser.Name,
		Role:      existingUser.Role,
		CreatedAt: existingUser.CreatedAt,
		UpdatedAt: existingUser.UpdatedAt,
	}

	// Cache for 1 hour
	if data, err := json.Marshal(userProto); err == nil {
		l.svcCtx.Redis.Setex(cacheKey, string(data), 3600)
	}

	return &user.GetProfileResponse{
		Code:    200,
		Message: "Success",
		User:    userProto,
	}, nil
}

type UpdateProfileLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateProfileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProfileLogic {
	return &UpdateProfileLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateProfileLogic) UpdateProfile(in *user.UpdateProfileRequest) (*user.UpdateProfileResponse, error) {
	existingUser, err := l.svcCtx.UserModel.FindById(in.UserId)
	if err != nil {
		return &user.UpdateProfileResponse{
			Code:    404,
			Message: "User not found",
		}, nil
	}

	existingUser.Name = in.Name
	existingUser.Phone = in.Phone
	existingUser.UpdatedAt = time.Now().Unix()

	err = l.svcCtx.UserModel.Update(existingUser)
	if err != nil {
		return &user.UpdateProfileResponse{
			Code:    500,
			Message: "Failed to update profile",
		}, err
	}

	// Invalidate cache
	cacheKey := fmt.Sprintf("user:profile:%d", in.UserId)
	l.svcCtx.Redis.Del(cacheKey)

	return &user.UpdateProfileResponse{
		Code:    200,
		Message: "Profile updated successfully",
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
