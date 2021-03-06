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
	IdRequest           = user.IdRequest
	LoginRequest        = user.LoginRequest
	LoginResponse       = user.LoginResponse
	RefreshRequest      = user.RefreshRequest
	RegisterRequest     = user.RegisterRequest
	RegisterResponse    = user.RegisterResponse
	UpdateRequest       = user.UpdateRequest
	UserResponse        = user.UserResponse
	VerifyTokenRequest  = user.VerifyTokenRequest
	VerifyTokenResponse = user.VerifyTokenResponse

	User interface {
		GetUser(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
		RegisterUser(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
		UpdateUser(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
		Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
		Refresh(ctx context.Context, in *RefreshRequest, opts ...grpc.CallOption) (*LoginResponse, error)
		VerifyToken(ctx context.Context, in *VerifyTokenRequest, opts ...grpc.CallOption) (*VerifyTokenResponse, error)
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

func (m *defaultUser) GetUser(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
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

func (m *defaultUser) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.Login(ctx, in, opts...)
}

func (m *defaultUser) Refresh(ctx context.Context, in *RefreshRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.Refresh(ctx, in, opts...)
}

func (m *defaultUser) VerifyToken(ctx context.Context, in *VerifyTokenRequest, opts ...grpc.CallOption) (*VerifyTokenResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.VerifyToken(ctx, in, opts...)
}
