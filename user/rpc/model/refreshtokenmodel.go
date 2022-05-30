package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ RefreshTokenModel = (*customRefreshTokenModel)(nil)

type (
	// RefreshTokenModel is an interface to be customized, add more methods here,
	// and implement the added methods in customRefreshTokenModel.
	RefreshTokenModel interface {
		refreshTokenModel
	}

	customRefreshTokenModel struct {
		*defaultRefreshTokenModel
	}
)

// NewRefreshTokenModel returns a model for the database table.
func NewRefreshTokenModel(conn sqlx.SqlConn, c cache.CacheConf) RefreshTokenModel {
	return &customRefreshTokenModel{
		defaultRefreshTokenModel: newRefreshTokenModel(conn, c),
	}
}
