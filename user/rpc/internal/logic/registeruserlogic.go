package logic

import (
	"context"
	"errors"

	"cyntra/user/rpc/internal/svc"
	"cyntra/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterUserLogic {
	return &RegisterUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterUserLogic) RegisterUser(in *user.RegisterRequest) (*user.RegisterResponse, error) {
	// todo: add your logic here and delete this line

	res, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, in.Username)

	if err != nil {
		return nil, err
	}

	if res == nil {
		return nil, errors.New("User Already Exists")
	}

	

	return &user.RegisterResponse{}, nil
}
