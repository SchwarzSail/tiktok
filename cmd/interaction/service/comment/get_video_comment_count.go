package service

import (
	"context"
	"github.com/pkg/errors"
	"strconv"
	"tiktok/cmd/interaction/dal/db"
)

// 获取视频评论数
func (s *CommentService) GetVideoCommentCount(ctx context.Context, vid string) (count int64, err error) {
	videoID, _ := strconv.Atoi(vid)
	commentDao := db.NewCommentDao(ctx)
	count, err = commentDao.GetVideoCommentCount(videoID)
	if err != nil {
		return 0, errors.WithMessage(err, "service.comment.GetVideoCommentCount failed")
	}
	return count, nil
}
