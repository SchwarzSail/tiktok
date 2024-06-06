package cache

import (
	"context"
	"github.com/pkg/errors"
	"strconv"
)

// AddViews 视频观看点击
func AddViews(ctx context.Context, vid string) (err error) {
	txn := RedisClient.TxPipeline()
	//观看量更新
	err = txn.Incr(ctx, VideoViewsKey(vid)).Err()
	if err != nil {
		return errors.Wrap(err, "cache.AddViews failed")
	}
	//排行榜
	err = txn.ZIncrBy(ctx, "Rank", 1, vid).Err()
	if err != nil {
		return errors.Wrap(err, "cache.AddViews update rank failed")
	}
	_, err = txn.Exec(ctx)
	if err != nil {
		return errors.Wrap(err, "cache.AddViews failed")
	}
	return
}

// 获取点击量
func GetVideoViews(ctx context.Context, vid string) (count int64, err error) {
	//预设key
	_ = RedisClient.SetNX(ctx, VideoViewsKey(vid), 0, 0)
	countStr, err := RedisClient.Get(ctx, VideoViewsKey(vid)).Result()
	if err != nil {
		return 0, errors.Wrap(err, "cache.GetVideoViews failed")
	}
	count, _ = strconv.ParseInt(countStr, 10, 64)
	return
}

func GetRankList(ctx context.Context, start, end int64) (list []string, err error) {
	list, err = RedisClient.ZRevRange(ctx, "Rank", start, end).Result()
	if err != nil {
		return nil, errors.Wrap(err, "cache.GetRank failed")
	}
	return list, nil
}
