package db

import (
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"time"
)

type Video struct {
	ID          int            `gorm:"primaryKey"`
	UserID      int            `gorm:"column:user_id"`
	VideoURL    string         `gorm:"column:video_url"`
	CoverURL    string         `gorm:"column:cover_url"`
	Title       string         `gorm:"column:title"`
	Description string         `gorm:"column:description"`
	CreatedAt   time.Time      `gorm:"column:created_at;default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time      `gorm:"column:updated_at;default:CURRENT_TIMESTAMP"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at" `
}

// ToMap 将video转化成map结构存入redis
func (v *Video) ToMap() map[string]interface{} {
	var deletedAt string
	if v.DeletedAt.Valid {
		deletedAt = v.DeletedAt.Time.Format("2006-01-02 15:04:05")
	}
	return map[string]interface{}{
		"id":          v.ID,
		"uid":         v.UserID,
		"title":       v.Title,
		"description": v.Description,
		"cover_url":   v.CoverURL,
		"video_url":   v.VideoURL,
		"created_at":  v.CreatedAt.Format("2006-01-02 15:04:05"),
		"update_at":   v.UpdatedAt.Format("2006-01-02 15:04:05"),
		"delete_at":   deletedAt,
	}
}

type VideoDao struct {
	*gorm.DB
}

func NewDBClient(ctx context.Context) *gorm.DB {
	db := DB
	return db.WithContext(ctx)
}

func NewVideoDao(ctx context.Context) *VideoDao {
	if ctx == nil {
		ctx = context.Background()
	}
	return &VideoDao{NewDBClient(ctx)}
}

func (dao *VideoDao) Feed(timeStamp int64) (list []Video, err error) {
	err = dao.DB.Model(&Video{}).Where("created_at < ?", time.Unix(timeStamp/1000, 0)).Order("created_at DESC").Find(&list).Error
	if err != nil {
		return nil, errors.Wrap(err, "db.Feed failed")
	}
	return
}

func (dao *VideoDao) CreateVideo(video *Video) (err error) {
	err = dao.DB.Model(&Video{}).Create(video).Error
	if err != nil {
		return errors.Wrap(err, "db.CreateVideo failed")
	}
	return nil
}

func (dao *VideoDao) GetVideoByVid(videoID int) (v Video, err error) {
	err = dao.DB.First(&v, videoID).Error
	if err != nil {
		return v, errors.Wrap(err, "db.GetVideoInfo failed")
	}
	return
}

func (dao *VideoDao) GetListByUid(uid int, pageNum, pageSize int64) (list []Video, err error) {
	err = dao.DB.Where("user_id = ?", uid).Order("created_at desc").
		Limit(int(pageSize)).
		Offset(int((pageNum - 1) * pageSize)).
		Find(&list).Error
	if err != nil {
		return nil, errors.Wrap(err, "db.GetListByUid failed")
	}
	return
}
