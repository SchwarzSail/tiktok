package db

import (
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"time"
)

type Follow struct {
	ID         uint           `gorm:"column:id;primaryKey"`
	UserID     int            `gorm:"column:user_id"`
	FollowerID int            `gorm:"column:follower_id"`
	CreatedAt  time.Time      `gorm:"column:created_at;default:CURRENT_TIMESTAMP"`
	UpdatedAt  time.Time      `gorm:"column:updated_at;default:CURRENT_TIMESTAMP"`
	DeletedAt  gorm.DeletedAt `gorm:"column:deleted_at;index"`
}

type FollowDao struct {
	*gorm.DB
}

func NewDBClient(ctx context.Context) *gorm.DB {
	db := DB
	return db.WithContext(ctx)
}

func NewFollowDao(ctx context.Context) *FollowDao {
	if ctx == nil {
		ctx = context.Background()
	}
	return &FollowDao{NewDBClient(ctx)}
}

func (dao *FollowDao) Create(comment *Follow) (err error) {
	err = dao.DB.Model(&Follow{}).Create(comment).Error
	if err != nil {
		return errors.Wrap(err, "db.Create failed")
	}
	return
}

func (dao *FollowDao) Cancel(uid, followerID int) (err error) {
	err = dao.DB.Where("user_id = ? AND follower_id = ?", uid, followerID).Delete(&Follow{}).Error
	if err != nil {
		return errors.Wrap(err, "db.Cancel failed")
	}
	return nil
}

func (dao *FollowDao) GetFollowers(uid int) (list []Follow, err error) {
	err = dao.DB.Where("user_id = ?", uid).Find(&list).Error
	if err != nil {
		return nil, errors.Wrap(err, "db.GetFollowers failed")
	}
	return list, nil
}

func (dao *FollowDao) GetFans(followerID int) (list []Follow, err error) {
	err = dao.DB.Where("follower_id = ?", followerID).Find(&list).Error
	if err != nil {
		return nil, errors.Wrap(err, "db.GetFans failed")
	}
	return list, nil
}
