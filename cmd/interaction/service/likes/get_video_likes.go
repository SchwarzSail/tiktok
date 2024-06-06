package service

import (
	"context"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
	"tiktok/cmd/interaction/dal/cache"
)

// 获取关于video的likes
func (s *LikesService) GetVideoLikes(ctx context.Context, vid string) (likesCount int64, err error) {
	if err = s.UpdateRedisVideoLikesData(ctx, vid, "0"); err != nil {
		return 0, errors.WithMessage(err, "service.GetVideInfo failed")
	}
	//从redis中获取
	likesCount, err = cache.GetVideoLikes(ctx, vid)
	if err == redis.Nil {
		likesCount = 0
	}
	if err != nil {
		return 0, errors.WithMessage(err, "service.GetVideInfo failed")
	}
	return
}
