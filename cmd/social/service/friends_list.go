package service

import (
	"context"
	"github.com/pkg/errors"
	"tiktok/cmd/social/dal/cache"
	"tiktok/cmd/social/rpc"
	"tiktok/internal/errno"
	"tiktok/kitex_gen/user"
)

func (s *SocialService) GetFriendsList(ctx context.Context, uid string, startIndex, endIndex int) (list []*user.User, err error) {
	//保证redis中有数据
	if err = s.UpdateRedisData(ctx, uid); err != nil {
		return nil, errors.WithMessage(err, "service.GetFriendsList failed")
	}
	uids, err := cache.GetFriendsList(ctx, uid)
	if err != nil {
		return nil, errors.WithMessage(err, "service.GetFriendsList failed")
	}
	if startIndex > len(uids) {
		err = errno.PageOutOfRange
		return nil, err
	}
	if endIndex > len(uids) {
		endIndex = len(uids)
	}
	uids = uids[startIndex:endIndex]
	list = make([]*user.User, len(uids))
	for i, uid := range uids {
		u, err := rpc.GetUserInfo(ctx, &user.InfoRequest{Uid: uid})
		if err != nil {
			return nil, errors.WithMessage(err, "service.FollowList rpc.GetUserInfo failed")
		}
		list[i] = u
	}
	return list, nil
}
