package service

import (
	"context"
	"github.com/pkg/errors"
	"tiktok/cmd/video/dal/cache"
	"tiktok/kitex_gen/video"
)

func (s *VideoService) GetRankList(ctx context.Context, pageNum, pageSize int64) (list []*video.Video, err error) {
	//从redis获取id
	start := (pageNum - 1) * pageSize
	end := start + pageSize - 1
	vids, err := cache.GetRankList(ctx, start, end)
	if err != nil {
		return nil, errors.WithMessage(err, "service.GetRankList failed")
	}
	//循环调用GetVideoInfo
	list = make([]*video.Video, len(vids))
	for i, v := range vids {
		videoInfo, err := s.GetVideoInfo(ctx, v)
		if err != nil {
			return nil, errors.WithMessage(err, "service.GetRankList failed")
		}

		list[i] = videoInfo
	}
	return list, nil
}
