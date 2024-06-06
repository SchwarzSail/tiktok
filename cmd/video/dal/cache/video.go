package cache

import (
	"context"
	"github.com/pkg/errors"
	"time"
)

func IsExistsVideoInfo(ctx context.Context, vid string) bool {
	return RedisClient.Exists(ctx, VideoInfoKey(vid)).Val() == 1
}

func AddVideoInfo(ctx context.Context, vid string, data []byte) (err error) {
	err = RedisClient.Set(ctx, VideoInfoKey(vid), string(data), 5*time.Minute).Err()
	if err != nil {
		return errors.Wrap(err, "cache.AddVideoInfo failed")
	}
	return
}

func GetVideoInfo(ctx context.Context, vid string) (data string, err error) {
	data, err = RedisClient.Get(ctx, VideoInfoKey(vid)).Result()
	if err != nil {
		return "", errors.Wrap(err, "cache.GetVideoInfo failed")
	}
	return
}
