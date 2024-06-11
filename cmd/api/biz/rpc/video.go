package rpc

import (
	"context"
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/pkg/errors"
	"tiktok/cmd/api/config"
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

func Feed(ctx context.Context, req *video.FeedRequest) (videos []*video.Video, err error) {
	resp, err := videoClient.Feed(ctx, req)
	if err != nil {
		return nil, errors.WithMessage(err, "api.rpc.video feed failed")
	}
	if resp.BaseResp.Code != errno.SuccessCode {
		return nil, errno.NewErrNo(resp.BaseResp.Code, resp.BaseResp.Msg)
	}
	return resp.Videos, nil
}

func Publish(ctx context.Context, req *video.PublishRequest) (err error) {
	resp, err := videoClient.Publish(ctx, req)
	if err != nil {
		return errors.WithMessage(err, "api.rpc.video publish failed")
	}
	if resp.BaseResp.Code != errno.SuccessCode {
		return errno.NewErrNo(resp.BaseResp.Code, resp.BaseResp.Msg)
	}
	return
}

func PublishList(ctx context.Context, req *video.PublishListRequest) (videos []*video.Video, err error) {
	resp, err := videoClient.PublishList(ctx, req)
	if err != nil {
		return nil, errors.WithMessage(err, "api.rpc.video PublishList failed")
	}
	if resp.BaseResp.Code != errno.SuccessCode {
		return nil, errno.NewErrNo(resp.BaseResp.Code, resp.BaseResp.Msg)
	}
	return resp.Videos, nil
}

func PopularList(ctx context.Context, req *video.PopularListRequest) (videos []*video.Video, err error) {
	resp, err := videoClient.PopularList(ctx, req)
	if err != nil {
		return nil, errors.WithMessage(err, "api.rpc.video PopularList failed")
	}
	if resp.BaseResp.Code != errno.SuccessCode {
		return nil, errno.NewErrNo(resp.BaseResp.Code, resp.BaseResp.Msg)
	}
	return resp.Videos, nil
}

func Search(ctx context.Context, req *video.SearchRequest) (videos []*video.Video, err error) {
	resp, err := videoClient.Search(ctx, req)
	if err != nil {
		return nil, errors.WithMessage(err, "api.rpc.video Search failed")
	}
	if resp.BaseResp.Code != errno.SuccessCode {
		return nil, errno.NewErrNo(resp.BaseResp.Code, resp.BaseResp.Msg)
	}
	return resp.Videos, nil
}
