// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	roleFieldNames          = builder.RawFieldNames(&Role{}, true)
	roleRows                = strings.Join(roleFieldNames, ",")
	roleRowsExpectAutoSet   = strings.Join(stringx.Remove(roleFieldNames, "create_time", "update_time"), ",")
	roleRowsWithPlaceHolder = builder.PostgreSqlJoin(stringx.Remove(roleFieldNames, "id", "create_time", "update_time"))

	cachePublicRoleIdPrefix   = "cache:public:role:id:"
	cachePublicRoleNamePrefix = "cache:public:role:name:"
)

type (
	roleModel interface {
		Insert(ctx context.Context, data *Role) (sql.Result, error)
		FindOne(ctx context.Context, id string) (*Role, error)
		FindOneByName(ctx context.Context, name string) (*Role, error)
		Update(ctx context.Context, data *Role) error
		Delete(ctx context.Context, id string) error
	}

	defaultRoleModel struct {
		sqlc.CachedConn
		table string
	}

	Role struct {
		Id         string    `db:"id"`
		Name       string    `db:"name"`
		CreatedAt  time.Time `db:"created_at"`
		UpdatedAt  time.Time `db:"updated_at"`
		ArchivedAt time.Time `db:"archived_at"`
	}
)

func newRoleModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultRoleModel {
	return &defaultRoleModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      `"public"."role"`,
	}
}

func (m *defaultRoleModel) Insert(ctx context.Context, data *Role) (sql.Result, error) {
	publicRoleIdKey := fmt.Sprintf("%s%v", cachePublicRoleIdPrefix, data.Id)
	publicRoleNameKey := fmt.Sprintf("%s%v", cachePublicRoleNamePrefix, data.Name)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values ($1, $2, $3, $4, $5)", m.table, roleRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Id, data.Name, data.CreatedAt, data.UpdatedAt, data.ArchivedAt)
	}, publicRoleIdKey, publicRoleNameKey)
	return ret, err
}

func (m *defaultRoleModel) FindOne(ctx context.Context, id string) (*Role, error) {
	publicRoleIdKey := fmt.Sprintf("%s%v", cachePublicRoleIdPrefix, id)
	var resp Role
	err := m.QueryRowCtx(ctx, &resp, publicRoleIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where id = $1 limit 1", roleRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultRoleModel) FindOneByName(ctx context.Context, name string) (*Role, error) {
	publicRoleNameKey := fmt.Sprintf("%s%v", cachePublicRoleNamePrefix, name)
	var resp Role
	err := m.QueryRowIndexCtx(ctx, &resp, publicRoleNameKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where name = $1 limit 1", roleRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, name); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultRoleModel) Update(ctx context.Context, data *Role) error {
	publicRoleIdKey := fmt.Sprintf("%s%v", cachePublicRoleIdPrefix, data.Id)
	publicRoleNameKey := fmt.Sprintf("%s%v", cachePublicRoleNamePrefix, data.Name)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where id = $1", m.table, roleRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.Id, data.Name, data.CreatedAt, data.UpdatedAt, data.ArchivedAt)
	}, publicRoleIdKey, publicRoleNameKey)
	return err
}

func (m *defaultRoleModel) Delete(ctx context.Context, id string) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	publicRoleIdKey := fmt.Sprintf("%s%v", cachePublicRoleIdPrefix, id)
	publicRoleNameKey := fmt.Sprintf("%s%v", cachePublicRoleNamePrefix, data.Name)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where id = $1", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, publicRoleIdKey, publicRoleNameKey)
	return err
}

func (m *defaultRoleModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cachePublicRoleIdPrefix, primary)
}

func (m *defaultRoleModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where id = $1 limit 1", roleRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultRoleModel) tableName() string {
	return m.table
}
