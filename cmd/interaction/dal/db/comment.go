package db

import (
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"time"
)

// 如果video_id==0 表示的是子评论
type Comment struct {
	ID        uint           `gorm:"column:id;primaryKey"`
	ParentID  int            `gorm:"column:parent_id"`
	VideoID   int            `gorm:"column:video_id"`
	UserID    int            `gorm:"column:user_id"`
	Content   string         `gorm:"column:content;size:255"`
	CreatedAt time.Time      `gorm:"column:created_at;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time      `gorm:"column:updated_at;default:CURRENT_TIMESTAMP"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index"`
}
type CommentDao struct {
	*gorm.DB
}

func NewDBClient(ctx context.Context) *gorm.DB {
	db := DB
	return db.WithContext(ctx)
}

func NewCommentDao(ctx context.Context) *CommentDao {
	if ctx == nil {
		ctx = context.Background()
	}
	return &CommentDao{NewDBClient(ctx)}
}

func (dao *CommentDao) CreateComment(comment *Comment) (err error) {
	err = dao.DB.Model(&Comment{}).Create(comment).Error
	if err != nil {
		return errors.Wrap(err, "db.CreateComment failed")
	}
	return
}

func (dao *CommentDao) GetCommentChildren(commentID uint, pageNum, pageSize int) (list []Comment, err error) {
	offset := (pageNum - 1) * pageSize
	err = dao.DB.Model(&Comment{}).Where("parent_id = ?", commentID).
		Order("created_at desc").
		Offset(offset).
		Limit(pageSize).
		Find(&list).Error
	if err != nil {
		return nil, errors.Wrap(err, "db.comment.GetCommentChildren failed")
	}
	return list, nil
}

func (dao *CommentDao) GetVideoCommentList(videoID int, pageNum, pageSize int) (list []Comment, err error) {
	offset := (pageNum - 1) * pageSize
	err = dao.DB.Model(&Comment{}).Where("video_id = ?", videoID).Order("created_at desc").
		Offset(offset).
		Limit(pageSize).
		Find(&list).Error
	if err != nil {
		return nil, errors.Wrap(err, "db.comment.GetVideoCommentList failed")
	}
	return list, nil
}

func (dao *CommentDao) GetVideoCommentCount(videoID int) (count int64, err error) {
	err = dao.DB.Model(&Comment{}).Where("video_id = ?", videoID).Count(&count).Error
	if err != nil {
		return 0, errors.Wrap(err, "db.GetVideoCommentCount failed")
	}
	return
}

func (dao *CommentDao) GetChildrenCount(commentID uint) (count int64, err error) {
	err = dao.DB.Model(&Comment{}).Where("parent_id = ?", commentID).Count(&count).Error
	if err != nil {
		return 0, errors.Wrap(err, "db.GetChildrenCount failed")
	}
	return
}

func (dao *CommentDao) DeleteComment(commentID, videoID int) (err error) {
	//删除视频的全部评论
	if commentID == 0 {
		if err = dao.DB.Where("video_id = ?", videoID).Delete(&Comment{}).Error; err != nil {
			return errors.Wrap(err, "db.DeleteComment failed")
		}
		return nil
	} else if videoID == 0 {
		//删除回复
		if err = dao.DB.Where("id = ?", commentID).Delete(&Comment{}).Error; err != nil {
			return errors.Wrap(err, "db.DeleteComment failed")
		}
		return nil
	}
	//删除视频中的指定评论
	if err = dao.DB.Where("id = ? AND video_id = ?", commentID, videoID).Delete(&Comment{}).Error; err != nil {
		return errors.Wrap(err, "db.DeleteComment failed")
	}
	return nil
}
