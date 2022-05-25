package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ScopeModel = (*customScopeModel)(nil)

type (
	// ScopeModel is an interface to be customized, add more methods here,
	// and implement the added methods in customScopeModel.
	ScopeModel interface {
		scopeModel
	}

	customScopeModel struct {
		*defaultScopeModel
	}
)

// NewScopeModel returns a model for the database table.
func NewScopeModel(conn sqlx.SqlConn, c cache.CacheConf) ScopeModel {
	return &customScopeModel{
		defaultScopeModel: newScopeModel(conn, c),
	}
}
