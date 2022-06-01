package svc

import (
	"cyntra/user/rpc/internal/config"
	"cyntra/user/rpc/model"
	"fmt"

	"github.com/go-redis/redis/v8"
	_ "github.com/lib/pq"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config            config.Config
	UserModel         model.UserModel
	RoleModel         model.RoleModel
	ScopeModel        model.ScopeModel
	RoleScopeModel    model.RoleScopeModel
	RefreshTokenModel model.RefreshTokenModel
	RedisClient       redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {

	options, err := redis.ParseURL(c.RedisUrl)

	if err != nil {
		fmt.Println("Could not read redis url")
		panic(err)
	}

	return &ServiceContext{
		Config:            c,
		UserModel:         model.NewUserModel(sqlx.NewSqlConn("postgres", c.DataSource), c.Cache),
		RoleModel:         model.NewRoleModel(sqlx.NewSqlConn("postgres", c.DataSource), c.Cache),
		ScopeModel:        model.NewScopeModel(sqlx.NewSqlConn("postgres", c.DataSource), c.Cache),
		RoleScopeModel:    model.NewRoleScopeModel(sqlx.NewSqlConn("postgres", c.DataSource), c.Cache),
		RefreshTokenModel: model.NewRefreshTokenModel(sqlx.NewSqlConn("postgres", c.DataSource), c.Cache),
		RedisClient:       *redis.NewClient(options),
	}
}
