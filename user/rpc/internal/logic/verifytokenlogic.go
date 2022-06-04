package logic

import (
	"context"
	"errors"
	"strings"

	"cyntra/user/rpc/internal/svc"
	"cyntra/user/rpc/types/user"
	"cyntra/user/rpc/utils"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type VerifyTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewVerifyTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VerifyTokenLogic {
	return &VerifyTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *VerifyTokenLogic) VerifyToken(in *user.VerifyTokenRequest) (*user.VerifyTokenResponse, error) {
	// todo: add your logic here and delete this line

	validatedToken, err := utils.ValidateToken(in.AccessToken, l.svcCtx.Config.JwtSecret, &AccessJwtClaims{})

	if err != nil {
		l.Logger.Error(err)
		return nil, errors.New("Invalid access token")
	}

	claims := validatedToken.Claims.(*AccessJwtClaims)

	parsedUserId := claims.UserId

	dbUsers, err := l.svcCtx.UserModel.FindOneWithScopes(l.ctx, parsedUserId)

	if err != nil && !errors.Is(err, sqlc.ErrNotFound) {
		l.Logger.Error(err)
		return nil, errors.New("Unknown error")
	}

	if errors.Is(err, sqlc.ErrNotFound) {
		l.Logger.Error("User does not exists")

		return nil, errors.New("Invalid creds")
	}

	// l.Logger.Error(fmt.Sprintf("Role Scopes: %v", dbUsers.Scopes))

	return &user.VerifyTokenResponse{
		Id:        dbUsers.Id,
		Username:  dbUsers.Username,
		FirstName: dbUsers.FirstName,
		LastName:  dbUsers.LastName,
		Gender:    dbUsers.Gender,
		Dob:       timestamppb.New(dbUsers.Dob),
		Scopes:    strings.Split(dbUsers.Scopes, ","),
	}, nil
}
