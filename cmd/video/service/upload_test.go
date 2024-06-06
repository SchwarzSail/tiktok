package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"tiktok/cmd/video/dal/cache"
	"tiktok/cmd/video/dal/db"
	"tiktok/config"
)

func TestUpload(t *testing.T) {
	config.InitConfig()
	db.InitMySQL()
	cache.LinkRedis()

	coverPath := "cover/" + uuid.Must(uuid.NewRandom()).String() + ".png"
	l := GetVideoService()
	err := l.Upload(context.Background(), []byte(""), coverPath)
	// 使用 assert 断言函数进行断言
	assert.NoError(t, err, "Error should be nil")
}
