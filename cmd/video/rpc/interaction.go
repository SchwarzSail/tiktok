package rpc

import (
	"context"
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/pkg/errors"
	"tiktok/cmd/video/config"
	"tiktok/internal/errno"
	"tiktok/kitex_gen/interaction"
	"tiktok/kitex_gen/interaction/interactionservice"
)

func InitInteractionRPC() {
	conf := config.Config
	r, err := etcd.NewEtcdResolver([]string{conf.EtcdHost + ":" + conf.EtcdPort})
	if err != nil {
		panic(err)
	}
	interactionClient, err = interactionservice.NewClient("interaction", client.WithResolver(r))
	if err != nil {
		panic(err)
	}
}

func GetVideoInfo(ctx context.Context, req *interaction.GetVideoInfoRequest) (likesCount, commentCount int64, err error) {
	resp, err := interactionClient.GetVideoInfo(ctx, req)
	if err != nil {
		return 0, 0, errors.Wrap(err, "rpc.GetVideoInfo failed")
	}
	if resp.BaseResp.Code != errno.SuccessCode {
		return 0, 0, errno.NewErrNo(resp.BaseResp.Code, resp.BaseResp.Msg)
	}
	return resp.LikeCount, resp.CommentCount, nil
}
