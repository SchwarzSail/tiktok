package main

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"tiktok/cmd/social/pack"
	"tiktok/cmd/social/rpc"
	"tiktok/cmd/social/service"
	"tiktok/internal/errno"
	"tiktok/kitex_gen/social"
	"tiktok/kitex_gen/user"
)

// SocialServiceImpl implements the last service interface defined in the IDL.
type SocialServiceImpl struct{}

// Follow implements the SocialServiceImpl interface.
func (s *SocialServiceImpl) Follow(ctx context.Context, req *social.FollowRequest) (resp *social.FollowResponse, err error) {
	resp = social.NewFollowResponse()
	l := service.GetSocialService()
	if err = l.Follow(ctx, req.UserId, req.ToUserId, req.ActionType); err != nil {
		klog.Error(err)
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(nil)
	return resp, nil
}

// FollowList implements the SocialServiceImpl interface.
func (s *SocialServiceImpl) FollowList(ctx context.Context, req *social.FollowListRequest) (resp *social.FollowListResponse, err error) {
	resp = social.NewFollowListResponse()
	l := service.GetSocialService()
	list, err := l.GetFollowList(ctx, req.UserId)
	if err != nil {
		klog.Error(err)
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	startIndex := (req.PageNum - 1) * req.PageSize
	endIndex := startIndex + req.PageSize
	if startIndex >= int64(len(list)) {
		err = errno.PageOutOfRange
		klog.Error(err)
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	if endIndex > int64(len(list)) {
		endIndex = int64(len(list))
	}
	list = list[startIndex:endIndex]
	users := make([]*user.User, len(list))
	for i, data := range list {
		u, err := rpc.GetUserInfo(ctx, &user.InfoRequest{Uid: data})
		if err != nil {
			klog.Error(err)
			resp.BaseResp = pack.BuildBaseResp(err)
			return resp, nil
		}
		users[i] = u
	}
	resp.BaseResp = pack.BuildBaseResp(nil)
	resp.Users = users
	return resp, nil
}

// FansList implements the SocialServiceImpl interface.
func (s *SocialServiceImpl) FansList(ctx context.Context, req *social.FansListRequest) (resp *social.FansListResponse, err error) {
	resp = social.NewFansListResponse()
	l := service.GetSocialService()
	list, err := l.GetFansList(ctx, req.Uid)
	if err != nil {
		klog.Error(err)
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	startIndex := (req.PageNum - 1) * req.PageSize
	endIndex := startIndex + req.PageSize
	if startIndex >= int64(len(list)) {
		err = errno.PageOutOfRange
		klog.Error(err)
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	if endIndex > int64(len(list)) {
		endIndex = int64(len(list))
	}
	list = list[startIndex:endIndex]
	users := make([]*user.User, len(list))
	for i, data := range list {
		u, err := rpc.GetUserInfo(ctx, &user.InfoRequest{Uid: data})
		if err != nil {
			klog.Error(err)
			resp.BaseResp = pack.BuildBaseResp(err)
			return resp, nil
		}
		users[i] = u
	}
	resp.BaseResp = pack.BuildBaseResp(nil)
	resp.Users = users
	return resp, nil
}

// FriendsList implements the SocialServiceImpl interface.
func (s *SocialServiceImpl) FriendsList(ctx context.Context, req *social.FriendsListRequest) (resp *social.FriendsListResponse, err error) {
	resp = social.NewFriendsListResponse()
	l := service.GetSocialService()
	list, err := l.GetFriendsList(ctx, req.Uid)
	if err != nil {
		klog.Error(err)
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	startIndex := (req.PageNum - 1) * req.PageSize
	endIndex := startIndex + req.PageSize
	if startIndex >= int64(len(list)) {
		err = errno.PageOutOfRange
		klog.Error(err)
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	if endIndex > int64(len(list)) {
		endIndex = int64(len(list))
	}
	list = list[startIndex:endIndex]
	users := make([]*user.User, len(list))
	for i, data := range list {
		u, err := rpc.GetUserInfo(ctx, &user.InfoRequest{Uid: data})
		if err != nil {
			klog.Error(err)
			resp.BaseResp = pack.BuildBaseResp(err)
			return resp, nil
		}
		users[i] = u
	}
	resp.BaseResp = pack.BuildBaseResp(nil)
	resp.Users = users
	return resp, nil
}
