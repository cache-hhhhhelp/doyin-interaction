package logic

import (
	"context"
	"strconv"
	"douyin-interaction/internal/svc"
	"douyin-interaction/types"
	"douyin-interaction/internal/model"
	"github.com/zeromicro/go-zero/core/logx"
)

type FavoriteListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFavoriteListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteListLogic {
	return &FavoriteListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}


func (l *FavoriteListLogic) toVideoResp(v []model.Favorite) ([]*__.Video, error) {

	ret := make([]*__.Video, len(v))
	
	for i := 0; i < len(v); i++ {
		searchId := v[i].VideoId
		videoEntity, err := l.svcCtx.VideoModel.FindOne(l.ctx, searchId)
		if err != nil{
			return nil, err
		}
		userEntity, err := l.svcCtx.UserModel.FindOne(l.ctx, videoEntity.UserId)
		if err != nil{
			return nil, err
		}
		ret[i] = &__.Video{
			Id:   videoEntity.VideoId,
			Author:   toUser(*userEntity),
			PlayUrl: videoEntity.PlayUrl,
			CoverUrl: videoEntity.CoverUrl,
			FavoriteCount: videoEntity.FavoriteCount,
			CommentCount: videoEntity.CommentCount,
			IsFavorite: videoEntity.IsFavorite,
			Title: videoEntity.Title,
		}
	}
	return ret, nil
}




func (l *FavoriteListLogic) FavoriteList(in *__.DouyinFavoriteListRequest) (*__.DouyinFavoriteListResponse, error) {
	userId, err := strconv.ParseInt(in.Token, 10, 64)
	sqlResult, err := l.svcCtx.FavoriteModel.SearchByUserId(l.ctx, userId)
	if err != nil {
		return &__.DouyinFavoriteListResponse{StatusCode: -1, StatusMsg:"Not Find"}, err
	}
	videoResp, err := l.toVideoResp(sqlResult)
	if err != nil {
		return &__.DouyinFavoriteListResponse{StatusCode: -1, StatusMsg:"Error Happen when VideoResp transform"}, err
	}
	return &__.DouyinFavoriteListResponse{StatusCode: 0, StatusMsg:"Success Find", VideoList:videoResp}, nil
}
