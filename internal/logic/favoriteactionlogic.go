package logic

import (
	"context"

	"douyin-interaction/internal/svc"
	"douyin-interaction/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavoriteActionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFavoriteActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteActionLogic {
	return &FavoriteActionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FavoriteActionLogic) FavoriteAction(in *__.DouyinFavoriteActionRequest) (*__.DouyinFavoriteActionResponse, error) {
	// todo: add your logic here and delete this line

	return &__.DouyinFavoriteActionResponse{}, nil
}
