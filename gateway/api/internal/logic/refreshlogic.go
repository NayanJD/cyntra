package logic

import (
	"context"

	"cyntra/gateway/api/internal/svc"
	"cyntra/gateway/api/internal/types"

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

func (l *RefreshLogic) Refresh(req *types.RegisterReq) (resp *types.RegisterReq, err error) {
	// resp, err := l.svcCtx.UserSvc.Refresh(l.ctx, &user.RefreshRequest{
	// 	RefreshToken: ,
	// })

	return
}
