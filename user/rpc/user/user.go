// Code generated by goctl. DO NOT EDIT!
// Source: user.proto

package user

import (
	"context"

	"cyntra/user/rpc/types/user"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	IdRequest        = user.IdRequest
	RegisterRequest  = user.RegisterRequest
	RegisterResponse = user.RegisterResponse
	UpdateRequest    = user.UpdateRequest
	UserResponse     = user.UserResponse

	User interface {
		GetUser(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*UserResponse, error)
		RegisterUser(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
		UpdateUser(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	}

	defaultUser struct {
		cli zrpc.Client
	}
)

func NewUser(cli zrpc.Client) User {
	return &defaultUser{
		cli: cli,
	}
}

func (m *defaultUser) GetUser(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.GetUser(ctx, in, opts...)
}

func (m *defaultUser) RegisterUser(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.RegisterUser(ctx, in, opts...)
}

func (m *defaultUser) UpdateUser(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.UpdateUser(ctx, in, opts...)
}
