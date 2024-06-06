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

func (s *VideoService) Feed(ctx context.Context, timeStamp int64) (list []*video.Video, err error) {
	//从db中获取
	videoDao := db.NewVideoDao(ctx)
	videos, err := videoDao.Feed(timeStamp)
	if err != nil {
		return nil, errors.WithMessage(err, "service.Feed failed")
	}
	//循环调度rpc
	list = make([]*video.Video, len(videos))
	for i, v := range videos {
		//获取likeCount, commentCount
		likeCount, commentCount, err := rpc.GetVideoInfo(ctx, &interaction.GetVideoInfoRequest{VideoId: strconv.Itoa(v.ID)})
		if err != nil {
			return nil, errors.WithMessage(err, "service.Feed failed")
		}
		views, err := cache.GetVideoViews(ctx, strconv.Itoa(v.ID))
		if err != nil {
			return nil, errors.WithMessage(err, "service.Feed failed")
		}
		//将video信息加载到redis
		errCh := make(chan error, 1)
		go func() {
			videoJson, err := json.Marshal(v)
			if err != nil {
				errCh <- err
				return
			}
			err = cache.AddVideoInfo(ctx, strconv.Itoa(v.ID), videoJson)
			if err != nil {
				errCh <- err
				return
			}
			errCh <- nil
		}()
		//pack
		list[i] = pack.BuildVideo(v, views, likeCount, commentCount)
		if err := <-errCh; err != nil {
			return nil, errors.WithMessage(err, "service.Feed failed")
		}
	}
	return list, nil
}
