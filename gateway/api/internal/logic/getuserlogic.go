package logic

import (
	"context"
	"time"

	"cyntra/gateway/api/internal/svc"
	"cyntra/gateway/api/internal/types"
	"cyntra/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserLogic) GetUser(req *types.GetUserReq) (*types.RegisterResp, error) {
	// todo: add your logic here and delete this line

	resp, err := l.svcCtx.UserSvc.GetUser(l.ctx, &user.IdRequest{Id: req.Id})

	if err != nil {
		return nil, err
	}

	return &types.RegisterResp{
		Id:         resp.Id,
		Username:   resp.Username,
		First_name: resp.FirstName,
		Last_name:  resp.LastName,
		Gender:     resp.Gender,
		Dob:        resp.Dob.AsTime().Format(time.RFC3339),
	}, nil
}
