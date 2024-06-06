package model

type Likes struct {
	VideoID    int    `json:"video_id"`
	CommentID  int    `json:"comment_id"`
	UserID     int    `json:"user_id"`
	ActionType string `json:"action_type"`
}
