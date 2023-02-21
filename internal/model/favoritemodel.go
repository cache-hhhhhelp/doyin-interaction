package model

import (
	"context"
	"fmt"
	"database/sql"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
)

var _ FavoriteModel = (*customFavoriteModel)(nil)

type (
	// FavoriteModel is an interface to be customized, add more methods here,
	// and implement the added methods in customFavoriteModel.
	FavoriteModel interface {
		favoriteModel
		DeleteByVideoIdUserId(ctx context.Context, userId int64, videoId int64) (error)
		SearchByUserId(ctx context.Context, userId int64) ([]Favorite, error)
		FindByUserIdVideoId(ctx context.Context, userId int64, videoId int64) ([]Favorite, error) 
	}

	customFavoriteModel struct {
		*defaultFavoriteModel
	}
)

// NewFavoriteModel returns a model for the database table.
func NewFavoriteModel(conn sqlx.SqlConn, c cache.CacheConf) FavoriteModel {
	return &customFavoriteModel{
		defaultFavoriteModel: newFavoriteModel(conn, c),
	}
}


func (c customFavoriteModel) DeleteByVideoIdUserId(ctx context.Context, userId int64, videoId int64) (error) {
	
	_, err := c.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `user_id` = ? and `video_id` = ?", c.table)
		return conn.ExecCtx(ctx, query, userId, videoId)
	})
	return err	
}


func (c customFavoriteModel) SearchByUserId(ctx context.Context, userId int64) ([]Favorite, error) {
	var resp []Favorite
	query := fmt.Sprintf("select * from %s where `user_id` = ?", c.table)
	err := c.QueryRowsNoCacheCtx(ctx, &resp, query, userId)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}

}

func (c customFavoriteModel) FindByUserIdVideoId(ctx context.Context, userId int64, videoId int64) ([]Favorite, error) {
	var resp []Favorite
	query := fmt.Sprintf("select * from %s where `user_id` = ? and `video_id` = ?", c.table)
	err := c.QueryRowsNoCacheCtx(ctx, &resp, query, userId, videoId)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}

}