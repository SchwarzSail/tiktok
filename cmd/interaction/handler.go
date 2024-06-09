package main

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"strconv"
	"tiktok/cmd/interaction/pack"
	"tiktok/cmd/interaction/rpc"
	service2 "tiktok/cmd/interaction/service/comment"
	service "tiktok/cmd/interaction/service/likes"

	"tiktok/kitex_gen/interaction"
	"tiktok/kitex_gen/video"
)

// InteractionServiceImpl implements the last service interface defined in the IDL.
type InteractionServiceImpl struct{}

// Like implements the InteractionServiceImpl interface.
func (s *InteractionServiceImpl) Like(ctx context.Context, req *interaction.LikeRequest) (resp *interaction.LikeResponse, err error) {

	resp = interaction.NewLikeResponse()
	l := service.GetLikesService()
	//对视频的点赞
	if req.CommentId == "0" {
		//这里隐含一个操作，点赞默认是点击观看过视频的，不然没有接口更新视频的观看量
		if err = l.UserLikeVideo(ctx, req.VideoId, req.UserId, req.ActionType); err != nil {
			klog.Error(err)
			resp.BaseResp = pack.BuildBaseResp(err)
			return resp, nil
		}
		if req.ActionType == "1" {
			_, err = rpc.WatchVideo(ctx, &video.WatchVideoRequest{VideoId: req.VideoId})
			if err != nil {
				klog.Error(err)
				resp.BaseResp = pack.BuildBaseResp(err)
				return resp, nil
			}
		}
	} else {
		if err = l.LikeComment(ctx, req.CommentId, req.UserId, req.ActionType); err != nil {
			klog.Error(err)
			resp.BaseResp = pack.BuildBaseResp(err)
			return resp, nil
		}
	}
	resp.BaseResp = pack.BuildBaseResp(nil)
	return
}

// LikeList implements the InteractionServiceImpl interface.
func (s *InteractionServiceImpl) LikeList(ctx context.Context, req *interaction.LikeListRequest) (resp *interaction.LikeListResponse, err error) {
	resp = interaction.NewLikeListResponse()
	l := service.GetLikesService()
	list, err := l.GetLikeList(ctx, req.UserId, req.PageNum, req.PageSize)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		klog.Error(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(nil)
	resp.Videos = list
	return
}

// CommentPublish implements the InteractionServiceImpl interface.
func (s *InteractionServiceImpl) CommentPublish(ctx context.Context, req *interaction.CommentPublishRequest) (resp *interaction.CommentPublishResponse, err error) {
	resp = interaction.NewCommentPublishResponse()
	l := service2.GetCommentService()
	if err = l.CommentPublish(ctx, req); err != nil {
		klog.Error(err)
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(nil)
	return
}

// CommentList implements the InteractionServiceImpl interface.
func (s *InteractionServiceImpl) CommentList(ctx context.Context, req *interaction.CommentListRequest) (resp *interaction.CommentListResponse, err error) {
	resp = interaction.NewCommentListResponse()
	l1 := service.GetLikesService()
	l2 := service2.GetCommentService()
	list, err := l2.GetCommentList(ctx, req)
	if err != nil {
		klog.Error(err)
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}

	comments := make([]*interaction.Comment, 0)
	for _, data := range list {
		//获取点赞量和评论量
		likeCount, err := l1.GetCommentLikes(ctx, strconv.Itoa(int(data.ID)))
		if err != nil {
			klog.Error(err)
			resp.BaseResp = pack.BuildBaseResp(err)
			return resp, nil
		}
		commentCount, err := l2.GetCommentChildrenCount(ctx, strconv.Itoa(int(data.ID)))
		if err != nil {
			klog.Error(err)
			resp.BaseResp = pack.BuildBaseResp(err)
			return resp, nil
		}
		//pack
		comments = append(comments, pack.BuildComment(data, likeCount, commentCount))
	}
	resp.BaseResp = pack.BuildBaseResp(nil)
	resp.Comments = comments
	return
}

// DeleteComment implements the InteractionServiceImpl interface.
func (s *InteractionServiceImpl) DeleteComment(ctx context.Context, req *interaction.DeleteCommentRequest) (resp *interaction.DeleteCommentResponse, err error) {
	// TODO: Your code here...
	resp = interaction.NewDeleteCommentResponse()
	l := service2.GetCommentService()
	if err = l.DeleteComment(ctx, req.CommentId, req.VideoId); err != nil {
		klog.Error(err)
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(nil)
	return resp, nil
}

// GetVideoInfo implements the InteractionServiceImpl interface.
func (s *InteractionServiceImpl) GetVideoInfo(ctx context.Context, req *interaction.GetVideoInfoRequest) (resp *interaction.GetVideoInfoResponse, err error) {
	resp = interaction.NewGetVideoInfoResponse()
	l1 := service.GetLikesService()
	l2 := service2.GetCommentService()
	likeCount, err := l1.GetVideoLikes(ctx, req.VideoId)
	if err != nil {
		klog.Error(err)
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	commentCount, err := l2.GetVideoCommentCount(ctx, req.VideoId)
	if err != nil {
		klog.Error(err)
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.LikeCount = likeCount
	resp.CommentCount = commentCount
	resp.BaseResp = pack.BuildBaseResp(nil)
	return
}
