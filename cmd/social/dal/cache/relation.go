package cache

import (
	"context"
	"github.com/pkg/errors"
	"time"
)

//-----------------------Follow-----------------

func IsExistFollowKey(ctx context.Context, uid string) (ok bool) {
	return RedisClient.Exists(ctx, FollowerKey(uid)).Val() == 1
}

// 查询是否存在特定的关注人
func IsExistFollower(ctx context.Context, uid, followerID string) (ok bool) {
	return RedisClient.SIsMember(ctx, FollowerKey(uid), followerID).Val()
}

func AddFollower(ctx context.Context, uid, followerID string) (err error) {
	err = RedisClient.SAdd(ctx, FollowerKey(uid), followerID).Err()
	if err != nil {
		return errors.Wrap(err, "cache.AddFollower failed")
	}
	return nil
}

func CancelFollower(ctx context.Context, uid, followerID string) (err error) {
	err = RedisClient.SRem(ctx, FollowerKey(uid), followerID).Err()
	if err != nil {
		return errors.Wrap(err, "cache.CancelFollower failed")
	}
	return nil
}

// 从数据库中获取的数据加载到redis
func UpdateFollowers(ctx context.Context, uid string, followerIDs []string) (err error) {
	txn := RedisClient.TxPipeline()
	for _, data := range followerIDs {
		if err = txn.SAdd(ctx, FollowerKey(uid), data).Err(); err != nil {
			return errors.Wrap(err, "cache.UpdateFollowers failed")
		}
	}
	_, err = txn.Exec(ctx)
	if err != nil {
		return errors.Wrap(err, "cache.UpdateFollowers failed")
	}
	RedisClient.Expire(ctx, FollowerKey(uid), 10*time.Minute)
	return nil
}

func GetFollowList(ctx context.Context, uid string) (list []string, err error) {
	list, err = RedisClient.SMembers(ctx, FollowerKey(uid)).Result()
	if err != nil {
		return nil, errors.Wrap(err, "cache.GetFollowList failed")
	}
	return list, nil
}

//-------------------------Fan---------------------

func IsExistFansKey(ctx context.Context, uid string) (ok bool) {
	return RedisClient.Exists(ctx, FansKey(uid)).Val() == 1
}

// 查询是否存在特定的粉丝
func IsExistFans(ctx context.Context, uid, fansID string) (ok bool) {
	return RedisClient.SIsMember(ctx, FansKey(uid), fansID).Val()
}

func AddFan(ctx context.Context, uid, fanID string) (err error) {
	err = RedisClient.SAdd(ctx, FansKey(uid), fanID).Err()
	if err != nil {
		return errors.Wrap(err, "cache.AddFan failed")
	}
	return nil
}

func CancelFan(ctx context.Context, uid, fanID string) (err error) {
	err = RedisClient.SRem(ctx, FansKey(uid), fanID).Err()
	if err != nil {
		return errors.Wrap(err, "cache.CancelFan failed")
	}
	return nil
}

func UpdateFans(ctx context.Context, uid string, fansID []string) (err error) {
	txn := RedisClient.TxPipeline()
	for _, data := range fansID {
		if err = txn.SAdd(ctx, FansKey(uid), data).Err(); err != nil {
			return errors.Wrap(err, "cache.UpdateFans failed")
		}
	}
	_, err = txn.Exec(ctx)
	if err != nil {
		return errors.Wrap(err, "cache.UpdateFans failed")
	}
	RedisClient.Expire(ctx, FansKey(uid), 10*time.Minute)
	return nil
}

func GetFansList(ctx context.Context, uid string) (list []string, err error) {
	list, err = RedisClient.SMembers(ctx, FansKey(uid)).Result()
	if err != nil {
		return nil, errors.Wrap(err, "cache.GetFansList failed")
	}
	return list, nil
}

// ----------------Friend--------------
// 通过对关注列表和粉丝列表求交集获得好友列表
func GetFriendsList(ctx context.Context, uid string) (list []string, err error) {
	list, err = RedisClient.SInter(ctx, FollowerKey(uid), FansKey(uid)).Result()
	if err != nil {
		return nil, errors.Wrap(err, "cache.GetFriendsList failed")
	}
	return list, nil
}
