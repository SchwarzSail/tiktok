package service

import "sync"

var CommentServiceOnce sync.Once

var CommentServiceIns *CommentService

type CommentService struct {
}

func GetCommentService() *CommentService {
	CommentServiceOnce.Do(func() {
		CommentServiceIns = &CommentService{}
	})
	return CommentServiceIns
}
