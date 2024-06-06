package service

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
	"tiktok/cmd/video/dal/es"
	"tiktok/config"
)

func TestFeed(t *testing.T) {
	config.InitConfig()
	es.LinkEs()
	l := GetVideoService()
	list, err := l.Feed(context.Background(), 1712391971879)
	// 使用 assert 断言函数进行断言
	assert.NoError(t, err, "Error should be nil")
	assert.NotNil(t, list)
}
