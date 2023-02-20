package logic

import (
	"context"

	"douyin-interaction/internal/svc"
	"douyin-interaction/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommentListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentListLogic {
	return &CommentListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}


func toCommentResp(v []model.Comment) []*__.DouyinCommentResponse {
	ret := make([]*__.DouyinCommentListResponse, len(v))
	// required int64 id = 1; // 视频评论id
 	//  	required User user =2; // 评论用户信息
 	//  	required string content = 3; // 评论内容
 	//  	required string create_date = 4; // 评论发布日期，格式 mm-dd

	for i := 0; i < len(v); i++ {
		ret[i] = &__.DouyinCommentResponse{
			id:   v[i].comment_id,
			user:   v[i].user_id,
			content: v[i].content,
			create_date: v[i].create_at
		}
	}
	return ret
}



func (l *CommentListLogic) CommentList(in *__.DouyinCommentListRequest) (*__.DouyinCommentListResponse, error) {
	// required string token = 1; // 用户鉴权token
	// required int64 video_id = 2; // 视频id

	user_id := in.token
	video_id := in.video_id
	
	sqlResult, err := l.svcCtx.CommentModel.FindMany(l.ctx, video_id)
	if err != nil {
		return &__.DouyinCommentListResponse{status_code: -1, status_msg:"", comment_list:nil}, err
	}
	// 		required int32 status_code = 1; // 状态码，0-成功，其他值-失败
	//  	optional string status_msg = 2; // 返回状态描述
	//  	repeated Comment comment_list = 3; // 评论列表
	return &__.DouyinCommentListResponse{status_code: 0, status_msg:"", comment_list:toCommentResp(result)}, nil

}
