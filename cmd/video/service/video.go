package service

import "sync"

var VideoServiceOnce sync.Once

var VideoServiceIns *VideoService

type VideoService struct {
}

func GetVideoService() *VideoService {
	VideoServiceOnce.Do(func() {
		VideoServiceIns = &VideoService{}
	})
	return VideoServiceIns
}
