// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	userFieldNames          = builder.RawFieldNames(&User{}, true)
	userRows                = strings.Join(userFieldNames, ",")
	userRowsExpectAutoSet   = strings.Join(stringx.Remove(userFieldNames, "create_time", "update_time"), ",")
	userRowsWithPlaceHolder = builder.PostgreSqlJoin(stringx.Remove(userFieldNames, "id", "create_time", "update_time"))

	cachePublicUserIdPrefix           = "cache:public:user:id:"
	cachePublicUserIdWithScopesPrefix = "cache:public:user_with_scopes:id:"
	cachePublicUserUsernamePrefix     = "cache:public:user:username:"
)

type (
	userModel interface {
		Insert(ctx context.Context, data *User) (sql.Result, error)
		FindOne(ctx context.Context, id string) (*User, error)
		FindOneByUsername(ctx context.Context, username string) (*User, error)
		Update(ctx context.Context, data *User) error
		Delete(ctx context.Context, id string) error
		FindOneWithScopes(ctx context.Context, id string) (*UserWithScopes, error)
	}

	defaultUserModel struct {
		sqlc.CachedConn
		table string
	}

	User struct {
		Id             string    `db:"id"`
		Username       string    `db:"username"`
		HashedPassword string    `db:"hashed_password"`
		FirstName      string    `db:"first_name"`
		LastName       string    `db:"last_name"`
		Gender         string    `db:"gender"`
		Dob            time.Time `db:"dob"`
		RoleId         string    `db:"role_id"`
		CreatedAt      time.Time `db:"created_at"`
		UpdatedAt      time.Time `db:"updated_at"`
		ArchivedAt     time.Time `db:"archived_at"`
	}

	UserWithScopes struct {
		User
		Scopes string
	}
)

func newUserModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultUserModel {
	return &defaultUserModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      `"public"."user"`,
	}
}

func (m *defaultUserModel) Insert(ctx context.Context, data *User) (sql.Result, error) {
	publicUserIdKey := fmt.Sprintf("%s%v", cachePublicUserIdPrefix, data.Id)
	publicUserUsernameKey := fmt.Sprintf("%s%v", cachePublicUserUsernamePrefix, data.Username)
	data.Id = uuid.NewString()
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)", m.table, userRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Id, data.Username, data.HashedPassword, data.FirstName, data.LastName, data.Gender, data.Dob, data.RoleId, data.CreatedAt, data.UpdatedAt, data.ArchivedAt)
	}, publicUserIdKey, publicUserUsernameKey)
	return ret, err
}

func (m *defaultUserModel) FindOne(ctx context.Context, id string) (*User, error) {
	publicUserIdKey := fmt.Sprintf("%s%v", cachePublicUserIdPrefix, id)
	var resp User
	err := m.QueryRowCtx(ctx, &resp, publicUserIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where id = $1 limit 1", userRows, m.table)
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

func (m *defaultUserModel) FindOneByUsername(ctx context.Context, username string) (*User, error) {
	publicUserUsernameKey := fmt.Sprintf("%s%v", cachePublicUserUsernamePrefix, username)
	var resp User
	err := m.QueryRowIndexCtx(ctx, &resp, publicUserUsernameKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where username = $1 limit 1", userRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, username); err != nil {
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

func (m *defaultUserModel) FindOneWithScopes(ctx context.Context, id string) (*UserWithScopes, error) {
	publicUserUsernameKey := fmt.Sprintf("%s%v", cachePublicUserIdWithScopesPrefix, id)
	var resp UserWithScopes
	err := m.QueryRowCtx(ctx, &resp, publicUserUsernameKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) (e error) {
		query := fmt.Sprintf(`select %s, 
			(select string_agg(s.name, ',')  
			from public.user u 
			join role_scope rs on u.role_id = rs.role_id 
			join scope s on rs.scope_id = s.id 
			where u.id = $1) as scopes 
		from public.user u  where u.id = $1`, userRows)
		return conn.QueryRowCtx(ctx, &resp, query, id)
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

func (m *defaultUserModel) Update(ctx context.Context, data *User) error {
	publicUserIdKey := fmt.Sprintf("%s%v", cachePublicUserIdPrefix, data.Id)
	publicUserUsernameKey := fmt.Sprintf("%s%v", cachePublicUserUsernamePrefix, data.Username)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where id = $1", m.table, userRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.Id, data.Username, data.HashedPassword, data.FirstName, data.LastName, data.Gender, data.Dob, data.RoleId, data.CreatedAt, data.UpdatedAt, data.ArchivedAt)
	}, publicUserIdKey, publicUserUsernameKey)
	return err
}

func (m *defaultUserModel) Delete(ctx context.Context, id string) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	publicUserIdKey := fmt.Sprintf("%s%v", cachePublicUserIdPrefix, id)
	publicUserUsernameKey := fmt.Sprintf("%s%v", cachePublicUserUsernamePrefix, data.Username)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where id = $1", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, publicUserIdKey, publicUserUsernameKey)
	return err
}

func (m *defaultUserModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cachePublicUserIdPrefix, primary)
}

func (m *defaultUserModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where id = $1 limit 1", userRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultUserModel) tableName() string {
	return m.table
}
