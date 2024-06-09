package main

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"tiktok/cmd/social/pack"
	"tiktok/cmd/social/service"
	"tiktok/kitex_gen/social"
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
	startIndex := int((req.PageNum - 1) * req.PageSize)
	endIndex := startIndex + int(req.PageSize)
	list, err := l.GetFollowList(ctx, req.UserId, startIndex, endIndex)
	if err != nil {
		klog.Error(err)
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResp(nil)
	resp.Users = list
	return resp, nil
}

// FansList implements the SocialServiceImpl interface.
func (s *SocialServiceImpl) FansList(ctx context.Context, req *social.FansListRequest) (resp *social.FansListResponse, err error) {
	resp = social.NewFansListResponse()
	l := service.GetSocialService()
	startIndex := int((req.PageNum - 1) * req.PageSize)
	endIndex := startIndex + int(req.PageSize)
	list, err := l.GetFansList(ctx, req.Uid, startIndex, endIndex)
	if err != nil {
		klog.Error(err)
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(nil)
	resp.Users = list
	return resp, nil
}

// FriendsList implements the SocialServiceImpl interface.
func (s *SocialServiceImpl) FriendsList(ctx context.Context, req *social.FriendsListRequest) (resp *social.FriendsListResponse, err error) {
	resp = social.NewFriendsListResponse()
	l := service.GetSocialService()
	startIndex := int((req.PageNum - 1) * req.PageSize)
	endIndex := startIndex + int(req.PageSize)
	list, err := l.GetFriendsList(ctx, req.Uid, startIndex, endIndex)
	if err != nil {
		klog.Error(err)
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(nil)
	resp.Users = list
	return resp, nil
}
