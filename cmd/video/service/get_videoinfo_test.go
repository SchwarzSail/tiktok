package service

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
	"tiktok/cmd/video/dal/cache"
	"tiktok/cmd/video/dal/db"
	"tiktok/config"
)

func TestGetVideoInfo(t *testing.T) {
	config.InitConfig()
	db.InitMySQL()
	cache.LinkRedis()
	l := GetVideoService()
	v, err := l.GetVideoInfo(context.Background(), "10")
	// 使用 assert 断言函数进行断言
	assert.NoError(t, err, "Error should be nil")
	assert.NotNil(t, v)

}
