package service

import (
	"context"
	"github.com/pkg/errors"
	"tiktok/cmd/interaction/dal/cache"
)

func (s *LikesService) GetCommentLikes(ctx context.Context, cid string) (count int64, err error) {
	if err = s.UpdateRedisCommentLikes(ctx, cid); err != nil {
		return 0, errors.WithMessage(err, "service.Likes.GetCommentLikes failed")
	}
	count, err = cache.GetCommentLikes(ctx, cid)
	if err != nil {
		return 0, errors.WithMessage(err, "service.Likes.GetCommentLikes failed")
	}
	return count, nil
}
