package pack

import (
	"strconv"
	"tiktok/cmd/video/dal/db"
	"tiktok/kitex_gen/video"
)

func BuildVideo(v db.Video, visitCount, likeCount, commentCount int64) *video.Video {
	var deletedAt string
	if v.DeletedAt.Valid {
		deletedAt = v.DeletedAt.Time.Format("2006-01-02 15:04:05")
	}
	return &video.Video{
		Id:           strconv.Itoa(v.ID),
		UserId:       strconv.Itoa(v.UserID),
		VideoUrl:     v.VideoURL,
		CoverUrl:     v.CoverURL,
		Title:        v.Title,
		Description:  v.Description,
		VisitCount:   visitCount,
		LikeCount:    likeCount,
		CommentCount: commentCount,
		CreatedAt:    v.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:    v.UpdatedAt.Format("2006-01-02 15:04:05"),
		DeletedAt:    deletedAt,
	}
}
