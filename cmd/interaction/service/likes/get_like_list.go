package service

import (
	"context"
	"github.com/pkg/errors"
	"tiktok/cmd/interaction/dal/cache"
)

func (s *LikesService) GetLikeList(ctx context.Context, uid string, pageNum, pageSize int64) (ids []string, err error) {
	//先保证redis中有数据
	if err = s.UpdateRedisVideoLikesData(ctx, "0", uid); err != nil {
		return ids, errors.WithMessage(err, "service.GetLikesList failed")
	}
	//从redis获取
	ids, err = cache.GetUserLikesVideoList(ctx, uid, pageNum, pageSize)
	if err == nil {
		return
	}
	return
}
