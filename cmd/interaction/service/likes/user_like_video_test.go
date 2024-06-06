package service

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
	"tiktok/cmd/interaction/dal/cache"
	"tiktok/cmd/interaction/dal/db"
	"tiktok/cmd/interaction/dal/mq"
	"tiktok/config"
)

func TestUserLikeVideo(t *testing.T) {
	config.InitConfig()
	db.InitMySQL()
	cache.LinkRedis()
	mq.LinkRabbitmq()
	l := GetLikesService()
	err := l.UserLikeVideo(context.Background(), "12", "3", "2")
	// 使用 assert 断言函数进行断言
	assert.NoError(t, err, "Error should be nil")

}
