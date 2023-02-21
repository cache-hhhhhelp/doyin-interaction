package logic

import (
	"context"
	"strconv"
	"douyin-interaction/internal/svc"
	"douyin-interaction/types"
	"douyin-interaction/internal/model"
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


func (l *CommentListLogic) toCommentResp(v []model.Comment) ([]*__.Comment, error) {
	ret := make([]*__.Comment, len(v))
	// required int64 id = 1; // 视频评论id
 	//  	required User user =2; // 评论用户信息
 	//  	required string content = 3; // 评论内容
 	//  	required string create_date = 4; // 评论发布日期，格式 mm-dd

	for i := 0; i < len(v); i++ {
		userEntity, err := l.svcCtx.UserModel.FindOne(l.ctx, v[i].UserId)
		if err != nil {
			return nil, err
		}
		ret[i] = &__.Comment{
			Id:   v[i].CommentId,
			User:   toUser(*userEntity),
			Content: v[i].Content,
			CreateDate: strconv.FormatInt(v[i].CreatedAt, 10),
		}
	}
	return ret, nil
}



func (l *CommentListLogic) CommentList(in *__.DouyinCommentListRequest) (*__.DouyinCommentListResponse, error) {
	// userId, err := strconv.ParseInt(in.Token, 10, 64) //
	videoId := in.VideoId
	sqlResult, err := l.svcCtx.CommentModel.FindManyByVideoId(l.ctx, videoId)
	if err != nil {
		return &__.DouyinCommentListResponse{StatusCode: -1, StatusMsg:"Undefined Error", CommentList:nil}, err
	}
	commentList, err := l.toCommentResp(sqlResult)
	if err != nil {
		return &__.DouyinCommentListResponse{StatusCode: -1, StatusMsg:"User is not Exist", CommentList:nil}, err
	}
	return &__.DouyinCommentListResponse{StatusCode: 0, StatusMsg:"Success Find Out Comment List",
				CommentList:commentList}, nil
}
