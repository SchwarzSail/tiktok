package service

import (
	"context"
	"github.com/pkg/errors"
	"tiktok/cmd/social/dal/cache"
)

func (s *SocialService) GetFriendsList(ctx context.Context, uid string) (uids []string, err error) {
	//保证redis中有数据
	if err = s.UpdateRedisData(ctx, uid); err != nil {
		return nil, errors.WithMessage(err, "service.GetFriendsList failed")
	}
	uids, err = cache.GetFriendsList(ctx, uid)
	if err != nil {
		return nil, errors.WithMessage(err, "service.GetFriendsList failed")
	}
	return uids, nil
}
