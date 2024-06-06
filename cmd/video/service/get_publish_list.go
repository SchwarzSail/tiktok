package service

import (
	"context"
	"github.com/pkg/errors"
	"strconv"
	"tiktok/cmd/video/dal/db"
	"tiktok/kitex_gen/video"
)

func (s *VideoService) GetPublishList(ctx context.Context, uid string, pageNum, pageSize int64) (list []*video.Video, err error) {
	//从数据库中查询
	videoDao := db.NewVideoDao(ctx)
	userID, _ := strconv.Atoi(uid)
	videos, err := videoDao.GetListByUid(userID, pageNum, pageSize)

	if err != nil {
		return nil, errors.WithMessage(err, "service.GetPublishList failed")
	}

	list = make([]*video.Video, len(videos))
	for i, data := range videos {
		vid := strconv.Itoa(data.ID)
		videoInfo, err := s.GetVideoInfo(ctx, vid)
		if err != nil {
			return nil, errors.WithMessage(err, "service.GetPublishList failed")
		}

		list[i] = videoInfo
	}
	return list, nil
}
