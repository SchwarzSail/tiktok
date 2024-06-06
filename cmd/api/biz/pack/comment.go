package pack

import (
	"tiktok/cmd/api/biz/model/api"
	"tiktok/kitex_gen/interaction"
)

func BuildComment(c *interaction.Comment) *api.Comment {
	return &api.Comment{
		ID:         c.Id,
		UserID:     c.UserId,
		VideoID:    c.VideoId,
		ParentID:   c.ParentId,
		LikeCount:  c.LikeCount,
		ChildCount: c.ChildCount,
		Content:    c.Content,
		CreatedAt:  c.CreatedAt,
		UpdatedAt:  c.UpdatedAt,
		DeletedAt:  c.DeletedAt,
	}
}

func BuildCommentList(list []*interaction.Comment) []*api.Comment {
	resp := make([]*api.Comment, 0)
	for _, data := range list {
		resp = append(resp, BuildComment(data))
	}
	return resp
}
