package rpc

import (
	"context"
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/pkg/errors"
	"tiktok/config"
	"tiktok/internal/errno"
	"tiktok/kitex_gen/user"
	"tiktok/kitex_gen/user/userservice"
)

func InitUserRPC() {
	conf := config.Config
	r, err := etcd.NewEtcdResolver([]string{conf.EtcdHost + ":" + conf.EtcdPort})
	if err != nil {
		panic(err)
	}
	userClient, err = userservice.NewClient("user", client.WithResolver(r))
	if err != nil {
		panic(err)
	}
}

func GetUserInfo(ctx context.Context, req *user.InfoRequest) (u *user.User, err error) {
	resp, err := userClient.GetInfo(ctx, req)
	if err != nil {
		return nil, errors.WithMessage(err, "rpc.GetUserInfo failed")
	}
	if resp.BaseResp.Code != errno.SuccessCode {
		return nil, errno.NewErrNo(resp.BaseResp.Code, resp.BaseResp.Msg)
	}
	return resp.User, nil
}
