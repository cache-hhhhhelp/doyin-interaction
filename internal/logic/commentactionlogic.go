package logic

import (
	"context"

	"douyin-interaction/internal/svc"
	"douyin-interaction/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommentActionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCommentActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentActionLogic {
	return &CommentActionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}






func (l *CommentActionLogic) CommentAction(in *__.DouyinCommentActionRequest) (*__.DouyinCommentActionResponse, error) {
	// string token = 1; // 用户鉴权token
	// int64 video_id = 2; // 视频id
    // int32 action_type = 3; // 1-发布评论，2-删除评论
    // string comment_text = 4; // 用户填写的评论内容，在action_type=1的时候使用 optional 
    // int64 comment_id = 5; // 要删除的评论id，在action_type=2的时候使用
    
	user_id := in.token
	video_id := in.video_id
	action_type := in.action_type
	//user_id = 
	if action_type == 1 {
		comment_text := in.comment_text
		// comment_id      bigint auto_increment,
		// user_id		   bigint       not null,
		// content         varchar(255) not null,
		// created_at      bigint       not null,
		// primary key 	   (comment_id)
		comment := model.Comment{
			user_id:  user_id,
			content:     content_text,
			CreatedAt: time.Now().Unix(),
		}
		sqlResult, err := l.svcCtx.CommentModel.Insert(l.ctx, &comment)
		if err != nil {
			return &__.DouyinCommentActionResponse{status_code: -1, status_msg:"", comment:nil}, err
		}
		return &__.DouyinCommentActionResponse{status_code: 0, status_msg:"", comment:comment}, nil

	}

	comment_id := in.comment_id
	sqlResult, err := l.svcCtx.CommentModel.Delete(l.ctx, comment_id)
	if err != nil {
		return &__.DouyinCommentActionResponse{status_code: -1, status_msg:"", comment:nil}, err
	}
	return &__.DouyinCommentActionResponse{status_code: 0, status_msg:"", comment:nil}, nil
	
}
