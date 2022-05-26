package svc

import (
	"cyntra/user/rpc/internal/config"
	"cyntra/user/rpc/model"

	_ "github.com/lib/pq"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config         config.Config
	UserModel      model.UserModel
	RoleModel      model.RoleModel
	ScopeModel     model.ScopeModel
	RoleScopeModel model.RoleScopeModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		UserModel:      model.NewUserModel(sqlx.NewSqlConn("postgres", c.DataSource), c.Cache),
		RoleModel:      model.NewRoleModel(sqlx.NewSqlConn("postgres", c.DataSource), c.Cache),
		ScopeModel:     model.NewScopeModel(sqlx.NewSqlConn("postgres", c.DataSource), c.Cache),
		RoleScopeModel: model.NewRoleScopeModel(sqlx.NewSqlConn("postgres", c.DataSource), c.Cache),
	}
}
