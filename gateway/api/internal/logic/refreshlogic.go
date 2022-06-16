package logic

import (
	"context"

	"cyntra/gateway/api/internal/svc"
	"cyntra/gateway/api/internal/types"
	"cyntra/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type RefreshLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRefreshLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RefreshLogic {
	return &RefreshLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RefreshLogic) Refresh(req *types.RefreshReq) (*types.LoginResp, error) {
	resp, err := l.svcCtx.UserSvc.Refresh(l.ctx, &user.RefreshRequest{
		RefreshToken: req.RefreshToken,
	})

	if err != nil {
		return nil, err
	}

	return &types.LoginResp{
		AccessToken:  resp.AccessToken,
		RefreshToken: resp.RefreshToken,
	}, nil
}
