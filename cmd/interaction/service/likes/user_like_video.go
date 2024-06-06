package service

import (
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"tiktok/cmd/interaction/dal/cache"
	"tiktok/cmd/interaction/dal/mq"
	"tiktok/cmd/interaction/dal/mq/model"

	"tiktok/internal/errno"
	"time"

	"strconv"
)

func (s *LikesService) UserLikeVideo(ctx context.Context, vid, uid, actionType string) (err error) {
	//确保redis有数据
	if err = s.UpdateRedisVideoLikesData(ctx, vid, uid); err != nil {
		return errors.WithMessage(err, "service.UserLikeVideo failed")
	}
	videoID, _ := strconv.Atoi(vid)
	userID, _ := strconv.Atoi(uid)
	//点赞操作
	if actionType == "1" {
		//先判断是否存在
		if ok := cache.IsExistUserLikesVideo(ctx, vid, uid); ok {
			return errno.LikesAlreadyExistErr
		}
		//更新redis
		timeStamp := float64(time.Now().Unix())
		if err = cache.AddUserLikesVideo(ctx, vid, uid, timeStamp); err != nil {
			return errors.WithMessage(err, "service.UserLikeVideo failed")
		}
		//更新另一个key
		if err = cache.UpdateVideoLikes(ctx, vid, 1); err != nil {
			return errors.WithMessage(err, "service.UserLikesVideo failed")
		}
	} else {
		//先判断是否存在
		if ok := cache.IsExistUserLikesVideo(ctx, vid, uid); !ok {
			return errno.LikesNotExistErr
		}
		//更新redis
		if err = cache.UserCancelLikesVideo(ctx, vid, uid); err != nil {
			return errors.WithMessage(err, "service.UserLikesVideo failed")
		}
		if err = cache.CancelVideoLikes(ctx, vid); err != nil {
			return errors.WithMessage(err, "service.UserLikesVideo failed")
		}
	}
	//将数据发布到mq
	//先创建结构体，将其序列化
	like := model.Likes{
		VideoID:    videoID,
		UserID:     userID,
		ActionType: actionType,
	}
	data, err := json.Marshal(like)
	if err != nil {
		return errors.Wrap(err, "json Marshal failed")
	}
	if err = mq.PublishVideoLikes(data, "mysql:likes"); err != nil {
		return errors.WithMessage(err, "service.UserLikesVideo failed")
	}
	return nil
}
