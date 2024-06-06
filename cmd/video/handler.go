package main

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"tiktok/cmd/video/config"
	"tiktok/cmd/video/pack"
	"tiktok/cmd/video/service"
	"tiktok/internal/errno"
	video "tiktok/kitex_gen/video"
)

// VideoServiceImpl implements the last service interface defined in the IDL.
type VideoServiceImpl struct{}

// Feed implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) Feed(ctx context.Context, req *video.FeedRequest) (resp *video.FeedResponse, err error) {
	resp = video.NewFeedResponse()
	l := service.GetVideoService()
	list, err := l.Feed(ctx, req.TimeStamp)
	if err != nil {
		klog.Error(err)
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(nil)
	resp.Videos = list
	return
}

// Publish implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) Publish(ctx context.Context, req *video.PublishRequest) (resp *video.PublishResponse, err error) {
	resp = video.NewPublishResponse()
	conf := config.Config.Oss
	videoPath := "video/" + uuid.Must(uuid.NewRandom()).String() + ".mp4"
	coverPath := "cover/" + uuid.Must(uuid.NewRandom()).String() + ".png"
	//现将传过来的data和cover上传到oss
	l := service.GetVideoService()
	group := new(errgroup.Group)
	//视频上传
	group.Go(func() error {
		err = l.Upload(ctx, req.Data, videoPath)
		if err != nil {
			klog.Error(err)
			return errno.UploadErr
		}
		return nil
	})
	//封面上传
	group.Go(func() error {
		err = l.Upload(ctx, req.Cover, coverPath)
		if err != nil {
			klog.Error(err)
			return errno.UploadErr
		}
		return nil
	})
	//处理错误组
	if err := group.Wait(); err != nil {
		resp.BaseResp = pack.BuildBaseResp(errors.WithMessage(err, "videoService.Upload failed"))
		return resp, nil
	}
	//将数据更新到数据库
	videoUrl := fmt.Sprintf("https://%s.%s/%s", conf.OssBucket, conf.OssEndPoint, videoPath)
	coverUrl := fmt.Sprintf("https://%s.%s/%s", conf.OssBucket, conf.OssEndPoint, coverPath)
	err = l.CreateVideo(ctx, videoUrl, coverUrl, req.Title, req.Description, req.Uid, req.UserName)
	if err != nil {
		klog.Error(err)
		resp.BaseResp = pack.BuildBaseResp(errors.WithMessage(err, "videoService.CreateVideo failed"))
		return resp, nil
	}
	//pack
	resp.BaseResp = pack.BuildBaseResp(nil)
	return
}

// PublishList implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) PublishList(ctx context.Context, req *video.PublishListRequest) (resp *video.PublishListResponse, err error) {
	resp = video.NewPublishListResponse()
	l := service.GetVideoService()
	list, err := l.GetPublishList(ctx, req.Uid, req.PageNum, req.PageSize)
	if err != nil {
		klog.Error(err)
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResp(nil)
	resp.Videos = list
	return
}

// PopularList implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) PopularList(ctx context.Context, req *video.PopularListRequest) (resp *video.PopularListResponse, err error) {
	resp = video.NewPopularListResponse()
	l := service.GetVideoService()
	list, err := l.GetRankList(ctx, req.PageNum, req.PageSize)
	if err != nil {
		klog.Error(err)
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(nil)
	resp.Videos = list
	return
}

// Search implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) Search(ctx context.Context, req *video.SearchRequest) (resp *video.SearchResponse, err error) {
	resp = video.NewSearchResponse()
	l := service.GetVideoService()
	list, err := l.Filter(ctx, req)
	if err != nil {
		klog.Error(err)
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResp(nil)
	resp.Videos = list
	return
}

// GetVideoInfo implements the VideoServiceImpl interface.
// 这个handler用来加载redis和mysql, 并且返回Video.video
func (s *VideoServiceImpl) GetVideoInfo(ctx context.Context, req *video.GetVideoInfoRequest) (resp *video.GetVideoInfoResponse, err error) {
	resp = video.NewGetVideoInfoResponse()
	l := service.GetVideoService()
	v, err := l.GetVideoInfo(ctx, req.VideoId)
	if err != nil {
		klog.Error(err)
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}

	//pack
	resp.BaseResp = pack.BuildBaseResp(nil)
	resp.Video = v
	return
}

// WatchVideo 只是在获取视频信息的基础上往redis添加了观看量
// WatchVideo implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) WatchVideo(ctx context.Context, req *video.WatchVideoRequest) (resp *video.WatchVideoResponse, err error) {
	resp = video.NewWatchVideoResponse()
	l := service.GetVideoService()
	v, err := l.WatchVideo(ctx, req.VideoId)
	if err != nil {
		klog.Error(err)
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}

	//pack
	resp.BaseResp = pack.BuildBaseResp(nil)
	resp.Video = v
	return
}
