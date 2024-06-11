package rpc

import (
	"context"
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/pkg/errors"
	"tiktok/cmd/interaction/config"
	"tiktok/internal/errno"
	"tiktok/kitex_gen/video"
	"tiktok/kitex_gen/video/videoservice"
	"tiktok/pkg/constants"
)

func InitVideoRPC() {
	conf := config.Config
	r, err := etcd.NewEtcdResolver([]string{conf.EtcdHost + ":" + conf.EtcdPort})
	if err != nil {
		panic(err)
	}
	videoClient, err = videoservice.NewClient("video", client.WithResolver(r), client.WithHostPorts(constants.VideoServiceIP))
	if err != nil {
		panic(err)
	}
}

func GetVideoInfo(ctx context.Context, req *video.GetVideoInfoRequest) (v *video.Video, err error) {
	resp, err := videoClient.GetVideoInfo(ctx, req)
	if err != nil {
		return v, errors.WithMessage(err, "rpc.GetVideoInfo failed")
	}
	if resp.BaseResp.Code != errno.SuccessCode {
		return v, errno.NewErrNo(resp.BaseResp.Code, resp.BaseResp.Msg)
	}
	return resp.Video, nil
}

func WatchVideo(ctx context.Context, req *video.WatchVideoRequest) (v *video.Video, err error) {
	resp, err := videoClient.WatchVideo(ctx, req)
	if err != nil {
		return v, errors.WithMessage(err, "rpc.GetVideoInfo failed")
	}
	if resp.BaseResp.Code != errno.SuccessCode {
		return v, errno.NewErrNo(resp.BaseResp.Code, resp.BaseResp.Msg)
	}
	return resp.Video, nil
}
