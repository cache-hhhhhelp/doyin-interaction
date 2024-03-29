// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	favoriteFieldNames          = builder.RawFieldNames(&Favorite{})
	favoriteRows                = strings.Join(favoriteFieldNames, ",")
	favoriteRowsExpectAutoSet   = strings.Join(stringx.Remove(favoriteFieldNames, "`favorite_id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	favoriteRowsWithPlaceHolder = strings.Join(stringx.Remove(favoriteFieldNames, "`favorite_id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheFavoriteFavoriteIdPrefix = "cache:favorite:favoriteId:"
)

type (
	favoriteModel interface {
		Insert(ctx context.Context, data *Favorite) (sql.Result, error)
		FindOne(ctx context.Context, favoriteId int64) (*Favorite, error)
		Update(ctx context.Context, data *Favorite) error
		Delete(ctx context.Context, favoriteId int64) error
	}

	defaultFavoriteModel struct {
		sqlc.CachedConn
		table string
	}

	Favorite struct {
		FavoriteId int64 `db:"favorite_id"`
		UserId     int64 `db:"user_id"`
		VideoId    int64 `db:"video_id"`
		AuthorId   int64 `db:"author_id"`
	}
)

func newFavoriteModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultFavoriteModel {
	return &defaultFavoriteModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`favorite`",
	}
}

func (m *defaultFavoriteModel) Delete(ctx context.Context, favoriteId int64) error {
	favoriteFavoriteIdKey := fmt.Sprintf("%s%v", cacheFavoriteFavoriteIdPrefix, favoriteId)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `favorite_id` = ?", m.table)
		return conn.ExecCtx(ctx, query, favoriteId)
	}, favoriteFavoriteIdKey)
	return err
}

func (m *defaultFavoriteModel) FindOne(ctx context.Context, favoriteId int64) (*Favorite, error) {
	favoriteFavoriteIdKey := fmt.Sprintf("%s%v", cacheFavoriteFavoriteIdPrefix, favoriteId)
	var resp Favorite
	err := m.QueryRowCtx(ctx, &resp, favoriteFavoriteIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `favorite_id` = ? limit 1", favoriteRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, favoriteId)
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

func (m *defaultFavoriteModel) Insert(ctx context.Context, data *Favorite) (sql.Result, error) {
	favoriteFavoriteIdKey := fmt.Sprintf("%s%v", cacheFavoriteFavoriteIdPrefix, data.FavoriteId)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, favoriteRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.UserId, data.VideoId, data.AuthorId)
	}, favoriteFavoriteIdKey)
	return ret, err
}

func (m *defaultFavoriteModel) Update(ctx context.Context, data *Favorite) error {
	favoriteFavoriteIdKey := fmt.Sprintf("%s%v", cacheFavoriteFavoriteIdPrefix, data.FavoriteId)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `favorite_id` = ?", m.table, favoriteRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.UserId, data.VideoId, data.AuthorId, data.FavoriteId)
	}, favoriteFavoriteIdKey)
	return err
}

func (m *defaultFavoriteModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheFavoriteFavoriteIdPrefix, primary)
}

func (m *defaultFavoriteModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `favorite_id` = ? limit 1", favoriteRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultFavoriteModel) tableName() string {
	return m.table
}
