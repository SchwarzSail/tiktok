package cache

import "fmt"

// 关注列表
func FollowerKey(uid string) string {
	return fmt.Sprintf("follower:%s", uid)
}

func FansKey(uid string) string {
	return fmt.Sprintf("fans:%s", uid)
}
