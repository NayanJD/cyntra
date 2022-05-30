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

type UpdateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserLogic) UpdateUser(in *user.UpdateRequest) (*user.RegisterResponse, error) {
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

	if in.FirstName != nil {
		res.FirstName = *in.FirstName
	}

	if in.LastName != nil {
		res.LastName = *in.LastName
	}

	if in.Gender != nil {
		res.Gender = *in.Gender
	}

	if in.Dob != nil {
		res.Dob = in.Dob.AsTime()
	}

	err = l.svcCtx.UserModel.Update(l.ctx, res)

	if err != nil {
		l.Logger.Error(err)
		return nil, errors.New("Unknown error")
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
