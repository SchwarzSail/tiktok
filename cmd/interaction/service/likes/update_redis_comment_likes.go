package service

import (
	"context"
	"github.com/pkg/errors"
	"strconv"
	"tiktok/cmd/interaction/dal/cache"
	"tiktok/cmd/interaction/dal/db"
)

// 处理关于评论和回复的点赞
func (s *LikesService) UpdateRedisCommentLikes(ctx context.Context, cid string) (err error) {
	likesDao := db.NewLikesDao(ctx)
	//如果redis中没有缓存的数据，从数据库获取更新到redis中
	if ok := cache.IsExistCommentLikesKey(ctx, cid); !ok {
		commentID, _ := strconv.Atoi(cid)
		list, err := likesDao.GetCommentLikesList(commentID)
		if err != nil {
			return errors.WithMessage(err, "service.UpdateRedisCommentLikes failed")
		}
		//更新到redis
		uids := make([]string, 0)
		for _, data := range list {
			uids = append(uids, strconv.Itoa(data.CommentID))
		}
		if err = cache.UpdateCommentLikes(ctx, cid, uids); err != nil {
			return errors.WithMessage(err, "service.UpdateRedisCommentLikes failed")
		}
	}
	return nil
}
