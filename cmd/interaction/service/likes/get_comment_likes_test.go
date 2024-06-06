package service

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
	"tiktok/cmd/interaction/dal/cache"
	"tiktok/cmd/interaction/dal/db"
	"tiktok/config"
)

func TestGetCommentLikes(t *testing.T) {
	config.InitConfig()
	db.InitMySQL()
	cache.LinkRedis()
	l := GetLikesService()
	_, err := l.GetCommentLikes(context.Background(), "1")
	// 使用 assert 断言函数进行断言
	assert.NoError(t, err, "Error should be nil")

}
