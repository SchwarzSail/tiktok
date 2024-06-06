package service

import (
	"context"
	"github.com/pkg/errors"
	"strconv"
	"tiktok/cmd/interaction/dal/db"
	"tiktok/kitex_gen/interaction"
)

func (s *CommentService) CommentPublish(ctx context.Context, req *interaction.CommentPublishRequest) (err error) {
	commentDao := db.NewCommentDao(ctx)
	//对视频的评论
	if req.CommentId == "0" {
		videoID, _ := strconv.Atoi(req.VideoId)
		uid, _ := strconv.Atoi(req.UserId)
		comment := &db.Comment{
			ParentID: 0,
			VideoID:  videoID,
			UserID:   uid,
			Content:  req.Content,
		}
		if err = commentDao.CreateComment(comment); err != nil {
			return errors.WithMessage(err, "service.CommentPublish failed")
		}
		return nil
	}
	//对评论的回复
	parentID, _ := strconv.Atoi(req.CommentId)
	uid, _ := strconv.Atoi(req.UserId)
	comment := &db.Comment{
		ParentID: parentID,
		VideoID:  0,
		UserID:   uid,
		Content:  req.Content,
	}
	if err = commentDao.CreateComment(comment); err != nil {
		return errors.WithMessage(err, "service.CommentPublish failed")
	}
	return nil
}
