package server

import (
	"context"
	"ticket-booking-platform/proto/user"
	"ticket-booking-platform/services/user-service/internal/svc"
)

type UserServer struct {
	user.UnimplementedUserServiceServer
	svcCtx *svc.ServiceContext
}

func NewUserServer(svcCtx *svc.ServiceContext) *UserServer {
	return &UserServer{svcCtx: svcCtx}
}

func (s *UserServer) Register(ctx context.Context, req *user.RegisterRequest) (*user.RegisterResponse, error) {
	u, err := s.svcCtx.RegisterLogic.Register(ctx, req.Email, req.Password, req.FullName, req.Phone)
	if err != nil {
		return nil, err
	}

	return &user.RegisterResponse{
		UserId:   u.ID,
		Email:    u.Email,
		FullName: u.FullName,
		Message:  "User registered successfully",
	}, nil
}

func (s *UserServer) Login(ctx context.Context, req *user.LoginRequest) (*user.LoginResponse, error) {
	token, u, err := s.svcCtx.LoginLogic.Login(ctx, req.Email, req.Password)
	if err != nil {
		return nil, err
	}

	return &user.LoginResponse{
		Token:    token,
		UserId:   u.ID,
		Email:    u.Email,
		FullName: u.FullName,
	}, nil
}

func (s *UserServer) GetProfile(ctx context.Context, req *user.GetProfileRequest) (*user.GetProfileResponse, error) {
	u, err := s.svcCtx.ProfileLogic.GetProfile(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	return &user.GetProfileResponse{
		UserId:   u.ID,
		Email:    u.Email,
		FullName: u.FullName,
		Phone:    u.Phone,
	}, nil
}

func (s *UserServer) UpdateProfile(ctx context.Context, req *user.UpdateProfileRequest) (*user.UpdateProfileResponse, error) {
	u, err := s.svcCtx.ProfileLogic.UpdateProfile(ctx, req.UserId, req.FullName, req.Phone)
	if err != nil {
		return nil, err
	}

	return &user.UpdateProfileResponse{
		UserId:   u.ID,
		FullName: u.FullName,
		Phone:    u.Phone,
		Message:  "Profile updated successfully",
	}, nil
}
