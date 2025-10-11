package logic

import (
	"context"

	"github.com/golang-jwt/jwt/v4"

	"ticket-booking-platform/proto/user"
	"ticket-booking-platform/services/user-service/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ValidateTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewValidateTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ValidateTokenLogic {
	return &ValidateTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ValidateTokenLogic) ValidateToken(in *user.ValidateTokenRequest) (*user.ValidateTokenResponse, error) {
	token, err := jwt.Parse(in.Token, func(token *jwt.Token) (interface{}, error) {
		return []byte(l.svcCtx.Config.JWT.Secret), nil
	})

	if err != nil || !token.Valid {
		return &user.ValidateTokenResponse{
			Code:    401,
			Message: "Invalid token",
		}, nil
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return &user.ValidateTokenResponse{
			Code:    401,
			Message: "Invalid token claims",
		}, nil
	}

	userId := int64(claims["user_id"].(float64))
	role := claims["role"].(string)

	return &user.ValidateTokenResponse{
		Code:    200,
		Message: "Token valid",
		UserId:  userId,
		Role:    role,
	}, nil
}
