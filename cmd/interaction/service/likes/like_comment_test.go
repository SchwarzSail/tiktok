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

func TestLikeComment(t *testing.T) {
	config.InitConfig()
	db.InitMySQL()
	cache.LinkRedis()
	mq.LinkRabbitmq()
	l := GetLikesService()
	err := l.LikeComment(context.Background(), "1", "3", "1")
	// 使用 assert 断言函数进行断言
	assert.NoError(t, err, "Error should be nil")

}
