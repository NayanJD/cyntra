package logic

import (
	"context"
	"errors"
	"fmt"
	"time"

	"cyntra/user/rpc/internal/svc"
	"cyntra/user/rpc/types/user"
	"cyntra/user/rpc/utils"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
)

var jwtUserIdKey = "user_id"
var jwtRefreshUUIDKey = "refresh_uuid"

type RefreshJwtClaims struct {
	*jwt.StandardClaims
	UserId      string
	RefreshUUID string
}

type AccessJwtClaims struct {
	*jwt.StandardClaims
	UserId string
}

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
	// todo: add your logic here and delete this line

	// l.Logger.Info(l.svcCtx.Config.JwtSecret)

	res, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, in.Username)

	if err != nil && !errors.Is(err, sqlc.ErrNotFound) {
		l.Logger.Error(err)
		return nil, errors.New("Unknown error")
	}

	if errors.Is(err, sqlc.ErrNotFound) {
		l.Logger.Error("User does not exists")

		return nil, errors.New("Invalid creds")
	}

	if !CheckPasswordHash(in.Password, res.HashedPassword) {
		l.Logger.Error("password does not match")
		return nil, errors.New("Invalid creds")
	} else {
		accessToken, refreshToken, err := generateAccessRefreshTokens(*l.svcCtx, l.ctx, res.Id)

		l.Logger.Error(err)

		return &user.LoginResponse{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		}, nil
	}
}

func generateAccessRefreshTokens(svcCtx svc.ServiceContext, ctx context.Context, userId string) (accessToken, refreshToken string, err error) {
	payload := map[string]interface{}{jwtUserIdKey: userId}
	duration := time.Duration(svcCtx.Config.JwtExpiresSeconds * int64(10e9))
	refreshDuration := time.Duration(svcCtx.Config.JwtRefreshExpiresSeconds * int64(10e9))

	accessClaims := AccessJwtClaims{
		&jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(svcCtx.Config.JwtExpiresSeconds * int64(time.Second))).Unix(),
		},
		userId,
	}
	accessToken, err = utils.CreateToken(accessClaims, duration, svcCtx.Config.JwtSecret)

	if err != nil {
		return "", "", err
	}

	payload[jwtRefreshUUIDKey] = uuid.New().String()
	fmt.Println(payload[jwtRefreshUUIDKey])

	refreshClaims := RefreshJwtClaims{
		&jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(svcCtx.Config.JwtExpiresSeconds * int64(time.Second))).Unix(),
		},
		userId,
		payload[jwtRefreshUUIDKey].(string),
	}
	refreshToken, err = utils.CreateToken(refreshClaims, refreshDuration, svcCtx.Config.JwtSecret)

	err = svcCtx.RedisClient.Set(ctx, payload[jwtRefreshUUIDKey].(string), userId, refreshDuration).Err()

	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}
