package svc

import (
	"cyntra/user/rpc/internal/config"
	"cyntra/user/rpc/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config
	UserModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		UserModel: model.NewUserModel(sqlx.NewSqlConn("pg", c.DataSource), c.Cache),
	}
}
