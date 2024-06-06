package service

import (
	"context"
	"github.com/pkg/errors"
	"tiktok/cmd/video/dal/es/document"
	"tiktok/kitex_gen/video"
)

func (s *VideoService) Filter(ctx context.Context, filter *video.SearchRequest) (list []*video.Video, err error) {
	//从es中获取
	videos, err := document.FilterVideo(ctx, filter)
	if err != nil {
		return nil, errors.WithMessage(err, "service.Filter failed")
	}
	list = make([]*video.Video, len(videos))
	for i, vid := range videos {
		videoInfo, err := s.GetVideoInfo(ctx, vid)
		if err != nil {
			return nil, errors.WithMessage(err, "service.Filter failed")
		}

		list[i] = videoInfo
	}
	return list, nil
}
