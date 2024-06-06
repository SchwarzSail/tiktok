package pack

import (
	"strconv"
	"tiktok/cmd/interaction/dal/db"
	"tiktok/kitex_gen/interaction"
)

func BuildComment(comment db.Comment, likeCount, childrenCount int64) *interaction.Comment {
	var deletedAt string
	if comment.DeletedAt.Valid {
		deletedAt = comment.DeletedAt.Time.Format("2006-01-02 15:04:05")
	}
	return &interaction.Comment{
		Id:         strconv.Itoa(int(comment.ID)),
		UserId:     strconv.Itoa(comment.UserID),
		VideoId:    strconv.Itoa(comment.VideoID),
		ParentId:   strconv.Itoa(comment.ParentID),
		ChildCount: childrenCount,
		LikeCount:  likeCount,
		Content:    comment.Content,
		CreatedAt:  comment.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:  comment.UpdatedAt.Format("2006-01-02 15:04:05"),
		DeletedAt:  deletedAt,
	}
}
