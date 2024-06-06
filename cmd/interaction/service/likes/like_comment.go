package service

import (
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"strconv"
	"tiktok/cmd/interaction/dal/cache"
	"tiktok/cmd/interaction/dal/mq"
	"tiktok/cmd/interaction/dal/mq/model"
	"tiktok/internal/errno"
)

// 对评论的点赞
func (s *LikesService) LikeComment(ctx context.Context, cid, uid string, actionType string) (err error) {
	//更新数据
	if err = s.UpdateRedisCommentLikes(ctx, cid); err != nil {
		return errors.WithMessage(err, "service.LikeComment failed")
	}
	//对redis进行操作
	if actionType == "1" {
		//判断是否已经点过赞了
		if ok := cache.IsExistCommentLikes(ctx, cid, uid); ok {
			return errno.LikesAlreadyExistErr
		}
		//点赞
		if err = cache.AddCommentLikes(ctx, cid, uid); err != nil {
			return errors.WithMessage(err, "service.LikeComment failed")
		}
	} else {
		//判断是否已经点过赞了
		if ok := cache.IsExistCommentLikes(ctx, cid, uid); !ok {
			return errno.LikesNotExistErr
		}
		//取消点赞
		if err = cache.CancelCommentLikes(ctx, cid, uid); err != nil {
			return errors.WithMessage(err, "service.LikeComment failed")
		}
	}
	//将数据更新到mq
	userID, _ := strconv.Atoi(uid)
	commentID, _ := strconv.Atoi(cid)
	like := model.Likes{
		VideoID:    0,
		CommentID:  commentID,
		UserID:     userID,
		ActionType: actionType,
	}
	data, err := json.Marshal(like)
	if err != nil {
		return errors.Wrap(err, "json Marshal failed")
	}
	if err = mq.PublishVideoLikes(data, "mysql:likes"); err != nil {
		return errors.WithMessage(err, "service.LikeComment failed")
	}
	return nil
}
