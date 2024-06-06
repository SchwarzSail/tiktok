package service

import (
	"sync"
)

var LikesServiceOnce sync.Once

var LikesServiceIns *LikesService

type LikesService struct {
}

func GetLikesService() *LikesService {
	LikesServiceOnce.Do(func() {
		LikesServiceIns = &LikesService{}
	})
	return LikesServiceIns
}
