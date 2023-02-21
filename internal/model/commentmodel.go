package model

import (
	"context"
	"fmt"
	"database/sql"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ CommentModel = (*customCommentModel)(nil)

type (
	// CommentModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCommentModel.
	CommentModel interface {
		commentModel
		FindManyByVideoId(ctx context.Context, videoId int64) ([]Comment, error)
		DeleteByCommentUser(ctx context.Context, commentId int64, userId int64) error
	}

	customCommentModel struct {
		*defaultCommentModel
		
	}
)

// NewCommentModel returns a model for the database table.
func NewCommentModel(conn sqlx.SqlConn, c cache.CacheConf) CommentModel {
	return &customCommentModel{
		defaultCommentModel: newCommentModel(conn, c),
	}
}



func (c customCommentModel) FindManyByVideoId(ctx context.Context, videoId int64) ([]Comment, error) {
	var resp []Comment
	query := fmt.Sprintf("select * from %s where `video_Id` = ?", c.table)
	err := c.QueryRowsNoCacheCtx(ctx, &resp, query, videoId)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}

}


func (m *defaultCommentModel) DeleteByCommentUser(ctx context.Context, commentId int64, userId int64) error {
	commentCommentIdKey := fmt.Sprintf("%s%v", cacheCommentCommentIdPrefix, commentId)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `comment_id` = ? and `user_id` = ?", m.table)
		return conn.ExecCtx(ctx, query, commentId, userId)
	}, commentCommentIdKey)
	return err
}
