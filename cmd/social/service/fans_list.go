package service

import (
	"context"
	"github.com/pkg/errors"
	"tiktok/cmd/social/dal/cache"
)

func (s *SocialService) GetFansList(ctx context.Context, uid string) (uids []string, err error) {
	//保证redis中有数据
	if err = s.UpdateRedisData(ctx, uid); err != nil {
		return nil, errors.WithMessage(err, "service.GetFansList failed")
	}
	uids, err = cache.GetFansList(ctx, uid)
	if err != nil {
		return nil, errors.WithMessage(err, "service.GetFansList failed")
	}
	return uids, nil
}
