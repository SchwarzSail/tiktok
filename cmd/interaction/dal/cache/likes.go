package cache

import (
	"context"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
	"strconv"
	"time"
)

func IsExistUserLikesVideoKey(ctx context.Context, uid string) (ok bool) {
	return RedisClient.Exists(ctx, UserLikesVideoKey(uid)).Val() == 1
}

// IsExistUserLikesVideo 查看有没有指定的点赞
func IsExistUserLikesVideo(ctx context.Context, vid, uid string) (ok bool) {
	_, err := RedisClient.ZScore(ctx, UserLikesVideoKey(uid), vid).Result()
	return err != redis.Nil
}

func AddUserLikesVideo(ctx context.Context, vid, uid string, timeStamp float64) (err error) {
	err = RedisClient.ZAdd(ctx, UserLikesVideoKey(uid), redis.Z{
		Score:  timeStamp,
		Member: vid,
	}).Err()
	if err != nil {
		return errors.Wrap(err, "cache.AddUserLikesVideo failed")
	}
	RedisClient.Expire(ctx, UserLikesVideoKey(uid), 5*time.Minute)
	return
}

func UserCancelLikesVideo(ctx context.Context, vid, uid string) (err error) {
	err = RedisClient.ZRem(ctx, UserLikesVideoKey(uid), vid).Err()
	if err != nil {
		return errors.Wrap(err, "cache.UserCancelLikesVideo failed")
	}
	return
}

func MSAddUserLikesVideo(ctx context.Context, uid string, vids []string, timeStamps []float64) (err error) {
	txn := RedisClient.TxPipeline()
	for i, v := range vids {
		err = txn.ZAdd(ctx, UserLikesVideoKey(uid), redis.Z{
			Score:  timeStamps[i],
			Member: v,
		}).Err()
		if err != nil {
			return errors.Wrap(err, "cache.MSAddUserLikesVideo failed")
		}
	}
	_, err = txn.Exec(ctx)
	if err != nil {
		return errors.Wrap(err, "cache.MSAddUserLikesVideo failed")
	}
	RedisClient.Expire(ctx, UserLikesVideoKey(uid), 5*time.Minute)
	return
}

func GetUserLikesVideoList(ctx context.Context, uid string, pageNum, pageSize int64) (list []string, err error) {
	startIndex := (pageNum - 1) * pageSize
	endIndex := startIndex + pageSize
	list, err = RedisClient.ZRevRange(ctx, UserLikesVideoKey(uid), startIndex, endIndex).Result()
	if err != nil {
		return list, errors.Wrap(err, "cache.GetUserLikesVideoList failed")
	}
	return
}

func IsExistsVideoLikes(ctx context.Context, vid string) (ok bool) {
	return RedisClient.Exists(ctx, VideoLikesKey(vid)).Val() == 1
}

// UpdateVideoLikes 更新视频点赞数
func UpdateVideoLikes(ctx context.Context, vid string, count int) (err error) {
	txn := RedisClient.TxPipeline()
	for i := 0; i < count; i++ {
		err = txn.Incr(ctx, VideoLikesKey(vid)).Err()
		if err != nil {
			return errors.Wrap(err, "cache.AddVideoLikes failed")
		}
	}
	_, err = txn.Exec(ctx)
	if err != nil {
		return errors.Wrap(err, "cache.AddVideoLikes failed")
	}
	RedisClient.Expire(ctx, VideoLikesKey(vid), 5*time.Minute)
	return
}

func CancelVideoLikes(ctx context.Context, vid string) (err error) {
	err = RedisClient.Decr(ctx, VideoLikesKey(vid)).Err()
	if err != nil {
		return errors.Wrap(err, "cache.CancelVideoLikes failed")
	}
	return
}

// GetVideoLikes 获取点赞量
func GetVideoLikes(ctx context.Context, vid string) (count int64, err error) {
	countStr, err := RedisClient.Get(ctx, VideoLikesKey(vid)).Result()
	if err == redis.Nil {
		return 0, nil
	}
	count, _ = strconv.ParseInt(countStr, 10, 64)
	return
}

// 使用集合类型
func IsExistCommentLikesKey(ctx context.Context, cid string) (ok bool) {
	return RedisClient.Exists(ctx, CommentLikesKey(cid)).Val() == 1
}

// 查看指定的点赞是否存在
func IsExistCommentLikes(ctx context.Context, cid, uid string) (ok bool) {
	return RedisClient.SIsMember(ctx, CommentLikesKey(cid), uid).Val()
}

func AddCommentLikes(ctx context.Context, cid, uid string) (err error) {
	err = RedisClient.SAdd(ctx, CommentLikesKey(cid), uid).Err()
	if err != nil {
		return errors.Wrap(err, "cache.AddCommentLikes failed")
	}
	return nil
}

// UpdateCommentLikes 更新评论点赞数
func UpdateCommentLikes(ctx context.Context, cid string, uids []string) (err error) {
	txn := RedisClient.TxPipeline()
	for _, uid := range uids {
		err = txn.SAdd(ctx, CommentLikesKey(cid), uid).Err()
		if err != nil {
			return errors.Wrap(err, "cache.UpdateCommentLikes failed")
		}
	}
	_, err = txn.Exec(ctx)
	if err != nil {
		return errors.Wrap(err, "cache.UpdateCommentLikes failed")
	}
	RedisClient.Expire(ctx, CommentLikesKey(cid), 5*time.Minute)
	return
}

func CancelCommentLikes(ctx context.Context, cid, uid string) (err error) {
	err = RedisClient.SRem(ctx, CommentLikesKey(cid), uid).Err()
	if err != nil {
		return errors.Wrap(err, "cache.CancelVideoLikes failed")
	}
	return
}

// GetCommentLikes 获取点赞量
func GetCommentLikes(ctx context.Context, cid string) (count int64, err error) {
	count, err = RedisClient.SCard(ctx, CommentLikesKey(cid)).Result()
	if err == redis.Nil {
		return 0, nil
	}
	if err != nil {
		return 0, errors.WithMessage(err, "cache.GetCommentLikes failed")
	}
	return
}
