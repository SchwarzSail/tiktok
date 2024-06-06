package service

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
	"tiktok/cmd/interaction/dal/cache"
	"tiktok/cmd/interaction/dal/db"
	"tiktok/config"
)

func TestGetLikeList(t *testing.T) {
	config.InitConfig()
	db.InitMySQL()
	cache.LinkRedis()
	l := GetLikesService()
	ids, err := l.GetLikeList(context.Background(), "3", 0, 10)
	// 使用 assert 断言函数进行断言
	assert.NoError(t, err, "Error should be nil")
	assert.NotNil(t, ids)
}
