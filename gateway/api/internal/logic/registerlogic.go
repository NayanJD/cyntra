package logic

import (
	"context"
	"time"

	"cyntra/gateway/api/internal/svc"
	"cyntra/gateway/api/internal/types"
	"cyntra/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/protobuf/types/known/timestamppb"
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

func (l *RegisterLogic) Register(req *types.RegisterReq) (*types.RegisterResp, error) {
	// todo: add your logic here and delete this line

	dobTime, err := time.Parse(time.RFC3339, req.Dob)

	if err != nil {
		return nil, err
	}

	resp, err := l.svcCtx.UserSvc.RegisterUser(l.ctx, &user.RegisterRequest{
		Username:  req.Username,
		FirstName: req.First_name,
		LastName:  req.Last_name,
		Gender:    req.Gender,
		Dob:       timestamppb.New(dobTime),
	})

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
