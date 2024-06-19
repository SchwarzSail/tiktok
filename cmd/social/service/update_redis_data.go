package service

import (
	"context"
	"github.com/pkg/errors"
	"strconv"
	"tiktok/cmd/social/dal/cache"
	"tiktok/cmd/social/dal/db"
)

// 将用户的关注列表和粉丝列表从mysql加载到redis
func (s *SocialService) UpdateRedisData(ctx context.Context, uid string) (err error) {
	if ok := cache.IsExistFollowKey(ctx, uid); !ok {
		//从数据库获取所有关注人的id
		followDao := db.NewFollowDao(ctx)
		userID, _ := strconv.Atoi(uid)
		list, err := followDao.GetFollowers(userID)
		if err != nil {
			return errors.WithMessage(err, "service.UpdateRedisData")
		}
		followerIDs := make([]string, len(list))
		for i, v := range list {
			followerIDs[i] = strconv.Itoa(v.FollowerID)
		}
		//加载到redis
		if err = cache.UpdateFollowers(ctx, uid, followerIDs); err != nil {
			return errors.WithMessage(err, "service.UpdateRedisData failed")
		}
	}
	if ok := cache.IsExistFansKey(ctx, uid); !ok {
		followDao := db.NewFollowDao(ctx)
		userID, _ := strconv.Atoi(uid)
		list, err := followDao.GetFans(userID)
		if err != nil {
			return errors.WithMessage(err, "service.UpdateRedisData")
		}
		fansIDs := make([]string, len(list))
		for i, v := range list {
			fansIDs[i] = strconv.Itoa(v.UserID)
		}
		if err = cache.UpdateFans(ctx, uid, fansIDs); err != nil {
			return errors.WithMessage(err, "service.UpdateRedisData failed")
		}
	}
	return nil
}
