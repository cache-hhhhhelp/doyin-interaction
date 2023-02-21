package logic

import (
	"context"
	"time"
	//"fmt"
	"strconv"
	"douyin-interaction/internal/svc"
	"douyin-interaction/types"
	"douyin-interaction/internal/model"
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



func toUser(v model.User) *__.User {
	ret := &__.User{
		Id:   v.UserId,
		Name: v.Username,
		FollowCount: v.FollowCount,
		FollowerCount: -1,
		IsFollow: v.IsFollow,
		Avatar: v.Avatar,
		BackgroundImage: v.BackgroundImage,
		Signature: v.Signature,
		TotalFavorited: v.TotalFavorited,
		WorkCount: v.WorkCount,
		FavoriteCount: v.FavoriteCount,
	}
	return ret
}




func (l *CommentActionLogic) CommentAction(in *__.DouyinCommentActionRequest) (*__.DouyinCommentActionResponse, error) {
	userId, err := strconv.ParseInt(in.Token, 10, 64) //
	videoId := in.VideoId
	actionType := in.ActionType
	if actionType == 1 {

		userEntity, err := l.svcCtx.UserModel.FindOne(l.ctx, userId)
		if err != nil {
			return &__.DouyinCommentActionResponse{StatusCode: -1, StatusMsg:"User Id is not Exist.", Comment:nil}, err
		}

		commentText := in.CommentText
		comment := model.Comment{
			UserId:   userId,
			VideoId:   videoId,
			Content:   commentText,
			CreatedAt: time.Now().Unix(),
		}
		sqlResult, err := l.svcCtx.CommentModel.Insert(l.ctx, &comment)
		if err != nil{
			return &__.DouyinCommentActionResponse{StatusCode: -1, StatusMsg:"Comment Insert Error", Comment:nil}, err
		}
		commentId, err := sqlResult.LastInsertId()
		if err != nil {
			return &__.DouyinCommentActionResponse{StatusCode: -1, StatusMsg:"Get Key Error", Comment:nil}, err
		}

		return &__.DouyinCommentActionResponse{StatusCode: 0, StatusMsg:"Success Post a Comment", 
								Comment: &__.Comment{Id: commentId,
								User:        toUser(*userEntity),
								Content :	 commentText,
								CreateDate:  strconv.FormatInt(time.Now().Unix(), 10)}}, nil
	}

	commentId := in.CommentId

	err = l.svcCtx.CommentModel.DeleteByCommentUser(l.ctx, commentId, userId)
	if err != nil {
		return &__.DouyinCommentActionResponse{StatusCode: -1, StatusMsg:"Comment Item Not Exist", Comment:nil}, err
	}
	return &__.DouyinCommentActionResponse{StatusCode: 0, StatusMsg:"Sucess Delete Comment Item", Comment:nil}, nil
	
}
