package logic

import (
	"context"
	"strconv"
	"douyin-interaction/internal/svc"
	"douyin-interaction/types"
	"douyin-interaction/internal/model"
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
	userId, err := strconv.ParseInt(in.Token, 10, 64)
	videoId := in.VideoId
	actionType := in.ActionType
	if actionType == 1 {
		favorite := model.Favorite{
			UserId:  userId,
			VideoId: videoId,
		}
		sqlResult, err := l.svcCtx.FavoriteModel.FindByUserIdVideoId(l.ctx, userId, videoId)
		if len(sqlResult) > 0 {
			return &__.DouyinFavoriteActionResponse{StatusCode: -1, StatusMsg:"Error Exist Favorite"}, err
		}
		_, err = l.svcCtx.FavoriteModel.Insert(l.ctx, &favorite)
		
		if err != nil {
			return &__.DouyinFavoriteActionResponse{StatusCode: -1, StatusMsg:"Insert Error"}, err
		}
		return &__.DouyinFavoriteActionResponse{StatusCode: 0, StatusMsg:"Success Favorite"}, nil
	}
	err = l.svcCtx.FavoriteModel.DeleteByVideoIdUserId(l.ctx, userId, videoId)
	if err != nil {
		return &__.DouyinFavoriteActionResponse{StatusCode: -1, StatusMsg:"Item not exist"}, err
	}
	return &__.DouyinFavoriteActionResponse{StatusCode: 0, StatusMsg:"Success Unlike"}, nil

}
