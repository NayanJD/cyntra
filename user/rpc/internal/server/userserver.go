// Code generated by goctl. DO NOT EDIT!
// Source: user.proto

package server

import (
	"context"

	"cyntra/user/rpc/internal/logic"
	"cyntra/user/rpc/internal/svc"
	"cyntra/user/rpc/types/user"
)

type UserServer struct {
	svcCtx *svc.ServiceContext
	user.UnimplementedUserServer
}

func NewUserServer(svcCtx *svc.ServiceContext) *UserServer {
	return &UserServer{
		svcCtx: svcCtx,
	}
}

func (s *UserServer) GetUser(ctx context.Context, in *user.IdRequest) (*user.RegisterResponse, error) {
	l := logic.NewGetUserLogic(ctx, s.svcCtx)
	return l.GetUser(in)
}

func (s *UserServer) RegisterUser(ctx context.Context, in *user.RegisterRequest) (*user.RegisterResponse, error) {
	l := logic.NewRegisterUserLogic(ctx, s.svcCtx)
	return l.RegisterUser(in)
}

func (s *UserServer) UpdateUser(ctx context.Context, in *user.UpdateRequest) (*user.RegisterResponse, error) {
	l := logic.NewUpdateUserLogic(ctx, s.svcCtx)
	return l.UpdateUser(in)
}
