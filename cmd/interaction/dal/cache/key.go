package cache

import "fmt"

// UserLikesVideoKey 用户对视频点赞
func UserLikesVideoKey(uid string) string {
	return fmt.Sprintf("likes:user:%s:video", uid)
}

// VideoLikesKey 视频点赞数
func VideoLikesKey(vid string) string {
	return fmt.Sprintf("likes:video:%s", vid)
}

// VideoCommentKey 视频评论, key是视频id， value为cid ,采用有序集合
func VideoCommentKey(vid string) string {
	return fmt.Sprintf("comment:video:%s", vid)
}

// CommentLikesKey 统计comment的点赞次数
func CommentLikesKey(cid string) string {
	return fmt.Sprintf("likes:comment:%s", cid)
}

// CommentChildrenKey 评论的回复 key是评论id， value是子评论id 采用有序集合
func CommentChildrenKey(cid string) string {
	return fmt.Sprintf("comment:children:%s", cid)
}
