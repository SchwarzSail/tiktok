package rpc

import (
	"context"
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/pkg/errors"
	"tiktok/config"
	"tiktok/internal/errno"
	"tiktok/kitex_gen/interaction"
	"tiktok/kitex_gen/interaction/interactionservice"
	"tiktok/kitex_gen/video"
)

func InitInteractionPRC() {
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

func Like(ctx context.Context, req *interaction.LikeRequest) (err error) {
	resp, err := interactionClient.Like(ctx, req)
	if err != nil {
		return errors.WithMessage(err, "rpc.Like failed")
	}
	if resp.BaseResp.Code != errno.SuccessCode {
		return errno.NewErrNo(resp.BaseResp.Code, resp.BaseResp.Msg)
	}
	return nil
}

func LikeList(ctx context.Context, req *interaction.LikeListRequest) (videos []*video.Video, err error) {
	resp, err := interactionClient.LikeList(ctx, req)
	if err != nil {
		return nil, errors.WithMessage(err, "rpc.LikeList failed")
	}
	if resp.BaseResp.Code != errno.SuccessCode {
		return nil, errno.NewErrNo(resp.BaseResp.Code, resp.BaseResp.Msg)
	}
	return resp.Videos, nil
}

func CommentPublish(ctx context.Context, req *interaction.CommentPublishRequest) (err error) {
	resp, err := interactionClient.CommentPublish(ctx, req)
	if err != nil {
		return errors.WithMessage(err, "rpc.CommentPublish failed")
	}
	if resp.BaseResp.Code != errno.SuccessCode {
		return errno.NewErrNo(resp.BaseResp.Code, resp.BaseResp.Msg)
	}
	return nil
}

func CommentList(ctx context.Context, req *interaction.CommentListRequest) (comments []*interaction.Comment, err error) {
	resp, err := interactionClient.CommentList(ctx, req)
	if err != nil {
		return nil, errors.WithMessage(err, "rpc.CommentList failed")
	}
	if resp.BaseResp.Code != errno.SuccessCode {
		return nil, errno.NewErrNo(resp.BaseResp.Code, resp.BaseResp.Msg)
	}
	return resp.Comments, nil
}

func DeleteComment(ctx context.Context, req *interaction.DeleteCommentRequest) (err error) {
	resp, err := interactionClient.DeleteComment(ctx, req)
	if err != nil {
		return errors.WithMessage(err, "rpc.DeleteComment failed")
	}
	if resp.BaseResp.Code != errno.SuccessCode {
		return errno.NewErrNo(resp.BaseResp.Code, resp.BaseResp.Msg)
	}
	return nil
}
