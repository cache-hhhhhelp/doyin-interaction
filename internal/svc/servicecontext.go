package svc

import ("douyin-interaction/internal/config"
		"douyin-interaction/internal/model"
		"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config
	CommentModel model.CommentModel
	UserModel model.UserModel
	FavoriteModel model.FavoriteModel
	VideoModel model.VideoModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		CommentModel: model.NewCommentModel(sqlx.NewMysql(c.Mysql.Datasource), c.Cache),
		UserModel: model.NewUserModel(sqlx.NewMysql(c.Mysql.Datasource), c.Cache),
		FavoriteModel: model.NewFavoriteModel(sqlx.NewMysql(c.Mysql.Datasource), c.Cache),
		VideoModel: model.NewVideoModel(sqlx.NewMysql(c.Mysql.Datasource), c.Cache),

	}

}
