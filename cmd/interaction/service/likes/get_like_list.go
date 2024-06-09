package service

import (
	"context"
	"github.com/pkg/errors"
	"tiktok/cmd/interaction/dal/cache"
	"tiktok/cmd/interaction/rpc"
	"tiktok/kitex_gen/video"
)

func (s *LikesService) GetLikeList(ctx context.Context, uid string, pageNum, pageSize int64) (list []*video.Video, err error) {
	//先保证redis中有数据
	if err = s.UpdateRedisVideoLikesData(ctx, "0", uid); err != nil {
		return nil, errors.WithMessage(err, "service.GetLikesList failed")
	}
	//从redis获取到id字段
	ids, err := cache.GetUserLikesVideoList(ctx, uid, pageNum, pageSize)
	//根据vids通过rpc获取数据
	list = make([]*video.Video, 0)
	for _, vid := range ids {
		v, err := rpc.GetVideoInfo(ctx, &video.GetVideoInfoRequest{VideoId: vid})
		if err != nil {
			return nil, errors.WithMessage(err, "rpc.GetVideoInfo failed")
		}
		list = append(list, v)
	}
	return
}
