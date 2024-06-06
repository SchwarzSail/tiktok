package service

import (
	"context"
	"github.com/pkg/errors"
	"strconv"
	"tiktok/cmd/interaction/dal/db"
)

// 获取子评论个数
func (s *CommentService) GetCommentChildrenCount(ctx context.Context, cid string) (count int64, err error) {
	commitID, _ := strconv.Atoi(cid)
	commentDao := db.NewCommentDao(ctx)
	count, err = commentDao.GetChildrenCount(uint(commitID))
	if err != nil {
		return 0, errors.WithMessage(err, "service.Comment.GetCommentChildrenCount failed")
	}
	return count, nil
}
