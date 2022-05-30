package logic

import (
	"context"
	"errors"

	"cyntra/user/rpc/internal/svc"
	"cyntra/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserLogic) GetUser(in *user.IdRequest) (*user.RegisterResponse, error) {
	// todo: add your logic here and delete this line

	res, err := l.svcCtx.UserModel.FindOne(l.ctx, in.Id)

	if err != nil {
		if errors.Is(err, sqlc.ErrNotFound) {
			return nil, errors.New("User not found")
		} else {
			l.Logger.Error(err)
			return nil, errors.New("Unknown error")
		}
	}

	return &user.RegisterResponse{
		Id:        res.Id,
		Username:  res.Username,
		FirstName: res.FirstName,
		LastName:  res.LastName,
		Gender:    res.Gender,
		Dob:       timestamppb.New(res.Dob),
	}, nil
}
