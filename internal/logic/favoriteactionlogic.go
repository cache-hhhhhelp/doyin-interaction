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
	// string token = 1; // 用户鉴权token
	// int64 video_id = 2; // 视频id
    // int32 action_type = 3; // 1-发布评论，2-删除评论
    // string comment_text = 4; // 用户填写的评论内容，在action_type=1的时候使用 optional 
    // int64 comment_id = 5; // 要删除的评论id，在action_type=2的时候使用
	user_id := in.token
	video_id := in.video_id
	action_type := in.action_type
	//user_id = getUserId(token)
	if action_type == 1 {
		favorite := model.Favorite{
			user_id:  user_id,
			video_id: video_id,
			CreatedAt: time.Now().Unix(),
		}
		sqlResult, err := l.svcCtx.FavoriteModel.Insert(l.ctx, &favorite)
		if err != nil {
			return &__.DouyinFavoriteActionResponse{status_code: -1, status_msg:""}, err
		}
		return &__.DouyinFavoriteActionResponse{status_code: 0, status_msg:""}, nil
	}
	comment_id := in.comment_id
	sqlResult, err := l.svcCtx.FavoriteModel.DeleteByVideoId(l.ctx, user_id, video_id)
	if err != nil {
		return &__.DouyinCommentActionResponse{status_code: -1, status_msg:""}, err
	}
	return &__.DouyinFavoriteActionResponse{status_code: 0, status_msg:""}, nil
}
