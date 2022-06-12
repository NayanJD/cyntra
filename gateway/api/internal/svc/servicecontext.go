package svc

import (
	"cyntra/gateway/api/internal/config"
	"cyntra/user/rpc/user"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	UserSvc user.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserSvc: user.NewUser(zrpc.MustNewClient(c.UserSvc)),
	}
}
