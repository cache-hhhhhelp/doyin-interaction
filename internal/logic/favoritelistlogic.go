package logic

import (
	"context"

	"douyin-interaction/internal/svc"
	"douyin-interaction/types"

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


func toVideoResp(v []model.Video) []*__.DouyinVideoResponse {
	ret := make([]*__.DouyinVideoListResponse, len(v))
	//    video_id             bigint auto_increment,
	//    user_id      bigint       not null,
	//    title          varchar(255) not null,
	//    created_at     bigint       not null,
	//    play_url		varchar(255) not null,
	//    cover_url 		varchar(255) not null,
	//    favorite_count	bigint       not null,
	//    comment_count	bigint       not null,
	//    is_favorite		boolean not null,

	for i := 0; i < len(v); i++ {
		//  required int64 id = 1; // 视频唯一标识
	 	//  required User author = 2; // 视频作者信息
	 	//  required string play_url = 3; // 视频播放地址
	 	//  required string cover_url = 4; // 视频封面地址
	 	//  required int64 favorite_count = 5; // 视频的点赞总数
	 	//  required int64 comment_count = 6; // 视频的评论总数
	 	//  required bool is_favorite = 7; // true-已点赞，false-未点赞
	 	//  required string title = 8; // 视频标题
		ret[i] = &__.DouyinVideoResponse{
			id:   v[i].video_id,
			author:   v[i].user_id,
			play_url: v[i].play_url,
			cover_url: v[i].cover_url,
			favorite_count: v[i].favorite_count,
			comment_count: v[i].comment_count,
			is_favorite: v[i].is_favorite,
			title: v[i].title
		}
	}
	return ret
}



func (l *FavoriteListLogic) FavoriteList(in *__.DouyinFavoriteListRequest) (*__.DouyinFavoriteListResponse, error) {
	// string token = 1; // 用户鉴权token
	// int64 video_id = 2; // 视频id
    // int32 action_type = 3; // 1-发布评论，2-删除评论
    // string comment_text = 4; // 用户填写的评论内容，在action_type=1的时候使用 optional 
    // int64 comment_id = 5; // 要删除的评论id，在action_type=2的时候使用
	user_id := in.token
	sqlResult, err := l.svcCtx.FavoriteModel.SearchByUserId(l.ctx, user_id)
	if err != nil {
		return &__.DouyinFavoriteListResponse{status_code: -1, status_msg:"", video_list:nil}, err
	}
	return &__.DouyinFavoriteListResponse{status_code: 0, status_msg:"", video_list:toVideoResp(sqlResult)}, nil
}
