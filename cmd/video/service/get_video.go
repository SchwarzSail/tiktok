package service

import (
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"strconv"
	"tiktok/cmd/video/dal/cache"
	"tiktok/cmd/video/dal/db"
	"tiktok/cmd/video/pack"
	"tiktok/cmd/video/rpc"
	"tiktok/kitex_gen/interaction"
	"tiktok/kitex_gen/video"
)

// 加载数据库和redis
func (s *VideoService) GetVideoInfo(ctx context.Context, videoID string) (v *video.Video, err error) {
	//先从redis获取
	var video db.Video
	//如果有，直接返回
	if ok := cache.IsExistsVideoInfo(ctx, videoID); ok {
		data, err := cache.GetVideoInfo(ctx, videoID)
		if err != nil {
			return nil, errors.WithMessage(err, "service.GetVideoInfo failed")
		}
		err = json.Unmarshal([]byte(data), &video)
		if err != nil {
			return nil, errors.Wrap(err, "video反序列化失败")
		}
	} else { //如果没有，从数据库中获取
		videoDao := db.NewVideoDao(ctx)
		vid, _ := strconv.Atoi(videoID)
		video, err = videoDao.GetVideoByVid(vid)
		if err != nil {
			return nil, errors.WithMessage(err, "service.GetVideoInfo failed")
		}
	}
	//将新的数据写入redis,或者直接更新redis
	errCh := make(chan error, 1)
	go func() {
		videoJson, err := json.Marshal(video)
		if err != nil {
			errCh <- err
			return
		}
		err = cache.AddVideoInfo(ctx, videoID, videoJson)
		if err != nil {
			errCh <- err
			return
		}
		errCh <- nil
	}()
	views, err := cache.GetVideoViews(ctx, videoID)
	if err != nil {
		return nil, errors.WithMessage(err, "service:GetVideoViews failed")
	}
	//rpc
	likeCount, commentCount, err := rpc.GetVideoInfo(ctx, &interaction.GetVideoInfoRequest{VideoId: videoID})
	if err != nil {
		return nil, errors.WithMessage(err, "service:GetVideoInfo failed")
	}
	if err := <-errCh; err != nil {
		return nil, errors.WithMessage(err, "service:GetVideoViews failed")
	}
	v = pack.BuildVideo(video, views, likeCount, commentCount)
	return v, nil
}
