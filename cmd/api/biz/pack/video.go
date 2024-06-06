package pack

import (
	"tiktok/cmd/api/biz/model/api"
	"tiktok/kitex_gen/video"
)

func BuildVideo(v *video.Video) *api.Video {
	return &api.Video{
		ID:           v.Id,
		UserID:       v.UserId,
		VideoURL:     v.VideoUrl,
		CoverURL:     v.CoverUrl,
		Title:        v.Title,
		Description:  v.Description,
		VisitCount:   v.VisitCount,
		LikeCount:    v.LikeCount,
		CommentCount: v.CommentCount,
		CreatedAt:    v.CreatedAt,
		UpdatedAt:    v.UpdatedAt,
		DeletedAt:    v.DeletedAt,
	}
}

func BuildVideoList(videos []*video.Video) []*api.Video {
	resp := make([]*api.Video, 0)
	for _, data := range videos {
		resp = append(resp, BuildVideo(data))
	}
	return resp
}
