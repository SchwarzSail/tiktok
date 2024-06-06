package rpc

import (
	"context"
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/pkg/errors"
	"tiktok/config"
	"tiktok/internal/errno"
	"tiktok/kitex_gen/social"
	"tiktok/kitex_gen/social/socialservice"
	"tiktok/kitex_gen/user"
)

func InitSocialRPC() {
	conf := config.Config
	r, err := etcd.NewEtcdResolver([]string{conf.EtcdHost + ":" + conf.EtcdPort})
	if err != nil {
		panic(err)
	}
	socialClient, err = socialservice.NewClient("social", client.WithResolver(r))
	if err != nil {
		panic(err)
	}
}

func Follow(ctx context.Context, req *social.FollowRequest) (err error) {
	resp, err := socialClient.Follow(ctx, req)
	if err != nil {
		return errors.WithMessage(err, "api.rpc.social Follow failed")
	}
	if resp.BaseResp.Code != errno.SuccessCode {
		return errno.NewErrNo(resp.BaseResp.Code, resp.BaseResp.Msg)
	}
	return
}

func FollowList(ctx context.Context, req *social.FollowListRequest) (list []*user.User, err error) {
	resp, err := socialClient.FollowList(ctx, req)
	if err != nil {
		return nil, errors.WithMessage(err, "api.rpc.social FollowList failed")
	}
	if resp.BaseResp.Code != errno.SuccessCode {
		return nil, errno.NewErrNo(resp.BaseResp.Code, resp.BaseResp.Msg)
	}
	return resp.Users, nil
}

func FansList(ctx context.Context, req *social.FansListRequest) (list []*user.User, err error) {
	resp, err := socialClient.FansList(ctx, req)
	if err != nil {
		return nil, errors.WithMessage(err, "api.rpc.social FansList failed")
	}
	if resp.BaseResp.Code != errno.SuccessCode {
		return nil, errno.NewErrNo(resp.BaseResp.Code, resp.BaseResp.Msg)
	}
	return resp.Users, nil
}

func FriendsList(ctx context.Context, req *social.FriendsListRequest) (list []*user.User, err error) {
	resp, err := socialClient.FriendsList(ctx, req)
	if err != nil {
		return nil, errors.WithMessage(err, "api.rpc.social FriendsList failed")
	}
	if resp.BaseResp.Code != errno.SuccessCode {
		return nil, errno.NewErrNo(resp.BaseResp.Code, resp.BaseResp.Msg)
	}
	return resp.Users, nil
}
