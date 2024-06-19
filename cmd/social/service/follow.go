package service

import (
	"context"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"strconv"
	"tiktok/cmd/social/dal/cache"
	"tiktok/cmd/social/dal/db"
	"tiktok/internal/errno"
)

func (s *SocialService) Follow(ctx context.Context, uid, followID string, actionType string) (err error) {
	//保证redis中包含用户的关注列表和粉丝列表
	if err = s.UpdateRedisData(ctx, uid); err != nil {
		return errors.WithMessage(err, "service.Follow failed")
	}
	if err = s.UpdateRedisData(ctx, followID); err != nil {
		return errors.WithMessage(err, "service.Follow failed")
	}
	//关注
	if actionType == "0" {
		//先判断是否已经关注过
		if ok := cache.IsExistFollower(ctx, uid, followID); ok {
			return errno.AlreadyFollowed
		}
		group := new(errgroup.Group)
		//更新redis
		group.Go(func() error {
			err = cache.AddFollower(ctx, uid, followID)
			return err
		})
		group.Go(func() error {
			err = cache.AddFan(ctx, followID, uid)
			return err
		})
		//更新mysql
		group.Go(func() error {
			followDao := db.NewFollowDao(ctx)
			userID, _ := strconv.Atoi(uid)
			followerID, _ := strconv.Atoi(followID)
			data := &db.Follow{
				UserID:     userID,
				FollowerID: followerID,
			}
			err = followDao.Create(data)
			return err
		})
		if err := group.Wait(); err != nil {
			return errors.WithMessage(err, "service.Follow failed")
		}
		return nil
	}
	//取关
	if ok := cache.IsExistFollower(ctx, uid, followID); !ok {
		return errno.NotFollowed
	}
	group := new(errgroup.Group)
	group.Go(func() error {
		err = cache.CancelFollower(ctx, uid, followID)
		return err
	})
	group.Go(func() error {
		err = cache.CancelFan(ctx, followID, uid)
		return err
	})
	group.Go(func() error {
		followDao := db.NewFollowDao(ctx)
		userID, _ := strconv.Atoi(uid)
		followerID, _ := strconv.Atoi(followID)
		err = followDao.Cancel(userID, followerID)
		return err
	})
	if err := group.Wait(); err != nil {
		return errors.WithMessage(err, "service.Follow failed")
	}
	return nil
}
