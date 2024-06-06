package db

import (
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"time"
)

type Likes struct {
	ID        uint           `gorm:"column:id;primaryKey"`
	VideoID   int            `gorm:"column:video_id"`
	UserID    int            `gorm:"column:user_id"`
	CommentID int            `gorm:"column:comment_id"`
	CreatedAt time.Time      `gorm:"column:created_at;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time      `gorm:"column:updated_at;default:CURRENT_TIMESTAMP"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index"`
}

type LikesDao struct {
	*gorm.DB
}

func NewLikesDao(ctx context.Context) *LikesDao {
	if ctx == nil {
		ctx = context.Background()
	}
	return &LikesDao{NewDBClient(ctx)}
}

func (dao *LikesDao) IsExistsLike(uid, vid, cid int) (ok bool) {
	var err error
	if vid != 0 {
		err = dao.DB.Model(&Likes{}).Where("user_id = ? AND video_id = ?", uid, vid).Error
	} else {
		err = dao.DB.Model(&Likes{}).Where("user_id = ? AND comment_id = ?", uid, cid).Error
	}
	return err == nil
}

func (dao *LikesDao) CreateLike(like *Likes) (err error) {
	err = dao.DB.Model(&Likes{}).Create(like).Error
	if err != nil {
		return errors.Wrap(err, "db.CreateLike failed")
	}
	return
}

func (dao *LikesDao) GetUserLikeList(uid int) (list []Likes, err error) {
	err = dao.DB.Model(&Likes{}).Where("user_id = ? AND comment_id = ?", uid, 0).
		Order("created_at desc").
		Find(&list).Error
	if err != nil {
		return list, errors.Wrap(err, "db.GetUserLikeList failed")
	}
	return
}

func (dao *LikesDao) GetVideoLikesCount(vid int) (count int64, err error) {
	err = dao.DB.Model(&Likes{}).Where("video_id = ?", vid).Count(&count).Error
	if err != nil {
		return 0, errors.Wrap(err, "db.GetVideoLikesCount failed")
	}
	return
}

func (dao *LikesDao) GetCommentLikesList(cid int) (list []Likes, err error) {
	err = dao.DB.Model(&Likes{}).Where("comment_id = ?", cid).Find(&list).Error
	if err != nil {
		return nil, errors.Wrap(err, "db.GetCommentLikesList failed")
	}
	return
}

func (dao *LikesDao) DeleteLikes(uid, vid, cid int) (err error) {
	if vid != 0 {
		err = dao.DB.Where("user_id = ? AND video_id = ?", uid, vid).Delete(&Likes{}).Error
		if err != nil {
			return errors.Wrap(err, "db.DeleteLikes failed")
		}
		return
	}
	err = dao.DB.Where("user_id = ? AND comment_id = ?", uid, cid).Delete(&Likes{}).Error
	if err != nil {
		return errors.Wrap(err, "db.DeleteLikes failed")
	}
	return
}
