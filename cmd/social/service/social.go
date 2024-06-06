package service

import "sync"

var SocialServiceOnce sync.Once

var SocialServiceIns *SocialService

type SocialService struct {
}

func GetSocialService() *SocialService {
	SocialServiceOnce.Do(func() {
		SocialServiceIns = &SocialService{}
	})
	return SocialServiceIns
}
