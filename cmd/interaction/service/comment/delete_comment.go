package service

import (
	"context"
	"github.com/pkg/errors"
	"strconv"
	"tiktok/cmd/interaction/dal/db"
)

func (s *CommentService) DeleteComment(ctx context.Context, cid, vid string) (err error) {
	commentID, _ := strconv.Atoi(cid)
	videoID, _ := strconv.Atoi(vid)
	commentDao := db.NewCommentDao(ctx)
	if err = commentDao.DeleteComment(commentID, videoID); err != nil {
		return errors.WithMessage(err, "service.Comment.DeleteComment failed")
	}
	return nil
}
