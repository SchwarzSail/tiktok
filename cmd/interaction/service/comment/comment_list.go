package service

import (
	"context"
	"github.com/pkg/errors"
	"strconv"
	"tiktok/cmd/interaction/dal/db"
	"tiktok/kitex_gen/interaction"
)

func (s *CommentService) GetCommentList(ctx context.Context, req *interaction.CommentListRequest) (list []db.Comment, err error) {
	commentDao := db.NewCommentDao(ctx)
	//获取视频评论表
	if req.CommentId == "0" {
		videoID, _ := strconv.Atoi(req.VideoId)
		list, err = commentDao.GetVideoCommentList(videoID, int(req.PageNum), int(req.PageSize))
		if err != nil {
			return nil, errors.WithMessage(err, "service.Comment.GetCommentList failed")
		}
		return list, nil
	}
	//回复列表
	parentID, _ := strconv.Atoi(req.CommentId)
	list, err = commentDao.GetCommentChildren(uint(parentID), int(req.PageNum), int(req.PageSize))
	if err != nil {
		return nil, errors.WithMessage(err, "service.Comment.GetCommentList failed")
	}
	return list, nil
}
