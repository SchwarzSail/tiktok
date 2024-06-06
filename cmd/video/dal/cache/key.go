package cache

import "fmt"

func VideoInfoKey(vid string) string {
	return fmt.Sprintf("video:info:%s", vid)
}

// VideoViewsKey 视频观看数
func VideoViewsKey(vid string) string {
	return fmt.Sprintf("video:views:%s", vid)
}
