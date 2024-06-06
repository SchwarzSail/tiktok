package service

import (
	"context"
	"github.com/pkg/errors"
	"strconv"
	"tiktok/cmd/video/dal/db"
	"tiktok/cmd/video/dal/es/document"
	"tiktok/cmd/video/dal/es/model"
	"time"
)

func (s *VideoService) CreateVideo(ctx context.Context, videoUrl, coverUrl, title, description, uid, userName string) (err error) {
	userId, _ := strconv.Atoi(uid)
	v := &db.Video{
		UserID:      userId,
		VideoURL:    videoUrl,
		CoverURL:    coverUrl,
		Title:       title,
		Description: description,
	}
	videoDao := db.NewVideoDao(ctx)
	err = videoDao.CreateVideo(v)
	if err != nil {
		return errors.WithMessage(err, "service.CreateVideo failed")
	}
	//更新es
	esModel := model.Video{
		Vid:         uint(v.ID),
		Uid:         uint(v.UserID),
		UserName:    userName,
		Title:       v.Title,
		Description: v.Description,
		CreatedAt:   time.Now().UnixNano() / int64(time.Millisecond),
	}

	if err = document.CreateVideo(ctx, esModel, strconv.Itoa(v.ID)); err != nil {
		return errors.WithMessage(err, "service.CreateVideo failed")
	}
	return nil
}
