package svc

import (
	"cyntra/product/rpc/internal/config"
	"cyntra/product/rpc/model"

	_ "github.com/lib/pq"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config       config.Config
	ProductModel model.ProductModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:       c,
		ProductModel: model.NewProductModel(sqlx.NewSqlConn("postgres", c.DataSource)),
	}
}
