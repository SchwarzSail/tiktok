package service

import (
	"context"
	"github.com/pkg/errors"
	"tiktok/cmd/video/dal/cache"
	"tiktok/kitex_gen/video"
)

func (s *VideoService) WatchVideo(ctx context.Context, vid string) (videoInfo *video.Video, err error) {
	videoInfo, err = s.GetVideoInfo(ctx, vid)
	if err != nil {
		return nil, errors.WithMessage(err, "service.WatchVideo failed")
	}
	err = cache.AddViews(ctx, vid)
	if err != nil {
		return nil, errors.WithMessage(err, "service.WatchVideo failed")
	}
	return videoInfo, nil
}
