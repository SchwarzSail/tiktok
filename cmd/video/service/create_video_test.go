package service

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"tiktok/cmd/video/dal/db"
	"tiktok/cmd/video/dal/es"
	"tiktok/config"
)

func TestCreateVideo(t *testing.T) {
	config.InitConfig()
	db.InitMySQL()
	es.LinkEs()
	conf := config.Config
	videoPath := "video/" + uuid.Must(uuid.NewRandom()).String() + ".mp4"
	coverPath := "cover/" + uuid.Must(uuid.NewRandom()).String() + ".png"
	videoUrl := fmt.Sprintf("https://%s.%s/%s", conf.OssBucket, conf.OssEndPoint, videoPath)
	coverUrl := fmt.Sprintf("https://%s.%s/%s", conf.OssBucket, conf.OssEndPoint, coverPath)
	l := GetVideoService()
	err := l.CreateVideo(context.Background(), videoUrl, coverUrl, "test", "test", "3", "smallpig")

	// 使用 assert 断言函数进行断言
	assert.NoError(t, err, "Error should be nil")

}
