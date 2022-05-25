package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ RoleScopeModel = (*customRoleScopeModel)(nil)

type (
	// RoleScopeModel is an interface to be customized, add more methods here,
	// and implement the added methods in customRoleScopeModel.
	RoleScopeModel interface {
		roleScopeModel
	}

	customRoleScopeModel struct {
		*defaultRoleScopeModel
	}
)

// NewRoleScopeModel returns a model for the database table.
func NewRoleScopeModel(conn sqlx.SqlConn, c cache.CacheConf) RoleScopeModel {
	return &customRoleScopeModel{
		defaultRoleScopeModel: newRoleScopeModel(conn, c),
	}
}
