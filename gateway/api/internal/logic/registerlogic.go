package logic

import (
	"context"

	"cyntra/gateway/api/internal/svc"
	"cyntra/gateway/api/internal/types"
	"cyntra/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterReq, err error) {
	// todo: add your logic here and delete this line
	resp, err := l.svcCtx.UserSvc.RegisterUser(l.ctx, &user.RegisterRequest{
		Username:  req.Username,
		FirstName: req.First_name,
	})
	if err != nil {
		return nil, err
	}

	return &types.ExpandResp{
		Url: resp.Url,
	}, nil
	return
}
