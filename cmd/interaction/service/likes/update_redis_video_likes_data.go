package service

import (
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"strconv"
	"tiktok/cmd/interaction/dal/cache"
	"tiktok/cmd/interaction/dal/db"
)

// 处理关于视频点赞
// 如果redis中的数据过期进行更新
// 将数据库中的数据全部加载到redis
// 如果传入的对应参数为0，则表示更新没有意义，无需操作
func (s *LikesService) UpdateRedisVideoLikesData(ctx context.Context, vid, uid string) (err error) {
	likesDao := db.NewLikesDao(ctx)
	if ok := cache.IsExistUserLikesVideoKey(ctx, uid); !ok && uid != "0" {
		//从数据库中获取数据
		userID, _ := strconv.Atoi(uid)
		videos, err := likesDao.GetUserLikeList(userID)
		if err != nil {
			return errors.WithMessage(err, "service.UpdateRedisVideoLikesData failed")
		}
		//更新到redis
		vids := make([]string, 0)
		timeStamps := make([]float64, 0)
		for _, v := range videos {
			videoID := strconv.Itoa(v.VideoID)
			vids = append(vids, videoID)
			timeStamps = append(timeStamps, float64(v.CreatedAt.Unix()))
		}
		if err = cache.MSAddUserLikesVideo(ctx, uid, vids, timeStamps); err != nil {
			return errors.WithMessage(err, "service.UpdateRedisData failed")
		}
	}
	if ok := cache.IsExistsVideoLikes(ctx, vid); !ok && vid != "0" {
		videoID, _ := strconv.Atoi(vid)
		count, err := likesDao.GetVideoLikesCount(videoID)
		if count == 0 || err == gorm.ErrRecordNotFound {
			if err = cache.UpdateVideoLikes(ctx, vid, int(count)); err != nil {
				return errors.WithMessage(err, "service.UpdateRedisData failed")
			}
		}
		if err != nil {
			return errors.WithMessage(err, "service.UpdateRedisData failed")
		}
		if err = cache.UpdateVideoLikes(ctx, vid, int(count)); err != nil {
			return errors.WithMessage(err, "service.UpdateRedisData failed")
		}
	}
	return nil
}
