package logic

import (
	"context"
	"errors"

	"cyntra/user/rpc/internal/svc"
	"cyntra/user/rpc/types/user"
	"cyntra/user/rpc/utils"

	"github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RefreshLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRefreshLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RefreshLogic {
	return &RefreshLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RefreshLogic) Refresh(in *user.RefreshRequest) (*user.LoginResponse, error) {
	validatedToken, err := utils.ValidateToken(in.RefreshToken, l.svcCtx.Config.JwtSecret, &RefreshJwtClaims{})

	if err != nil {
		l.Logger.Error(err)
		return nil, status.Error(codes.InvalidArgument, "Invalid refresh token")
	}

	claims := validatedToken.Claims.(*RefreshJwtClaims)

	parsedUserId := claims.UserId
	parsedUuid := claims.RefreshUUID

	cacheUserId, err := l.svcCtx.RedisClient.Get(l.ctx, parsedUuid).Result()

	if err == redis.Nil {
		l.Logger.Error(err)
		return nil, status.Error(codes.InvalidArgument, "Invalid refresh token")
	} else if err != nil {
		l.Logger.Error(err)
		return nil, status.Error(codes.Unknown, codes.Unknown.String())
	}

	l.svcCtx.RedisClient.Del(l.ctx, parsedUuid)

	userId := parsedUserId

	if cacheUserId != userId {
		l.Logger.Error("Cache userId and refreshToken does not exists")
		return nil, status.Error(codes.InvalidArgument, "Invalid refresh token")
	}

	res, err := l.svcCtx.UserModel.FindOne(l.ctx, userId)

	if err != nil && !errors.Is(err, sqlc.ErrNotFound) {
		l.Logger.Error(err)
		return nil, status.Error(codes.Unknown, codes.Unknown.String())
	}

	if errors.Is(err, sqlc.ErrNotFound) {
		l.Logger.Error("User does not exists")

		return nil, status.Error(codes.InvalidArgument, "Invalid creds")
	}

	accessToken, refreshToken, err := generateAccessRefreshTokens(*l.svcCtx, l.ctx, res.Id)

	l.Logger.Error(err)

	return &user.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
