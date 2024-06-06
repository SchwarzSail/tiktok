package service

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
	"tiktok/cmd/video/dal/cache"
	"tiktok/cmd/video/dal/db"
	"tiktok/cmd/video/dal/es"
	"tiktok/config"
	"tiktok/kitex_gen/video"
)

func TestFilter(t *testing.T) {
	config.InitConfig()
	db.InitMySQL()
	cache.LinkRedis()
	es.LinkEs()
	l := GetVideoService()
	list, err := l.Filter(context.Background(), &video.SearchRequest{
		Keyword:  "test",
		PageNum:  0,
		PageSize: 10,
		FromDate: nil,
		ToDate:   nil,
		Username: nil,
	})
	// 使用 assert 断言函数进行断言
	assert.NoError(t, err, "Error should be nil")
	assert.NotNil(t, list)
}
