package service

import (
	"bytes"
	"context"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/pkg/errors"
	"tiktok/config"
)

func (s *VideoService) Upload(ctx context.Context, data []byte, objectKey string) (err error) {
	conf := config.Config.Oss
	client, err := oss.New(conf.OssEndPoint, conf.OssAccessKeyId, conf.OssAccessKeySecret)
	if err != nil {
		return errors.Wrap(err, "service.Upload failed")
	}
	bucket, err := client.Bucket(conf.OssBucket)
	if err != nil {
		return errors.Wrap(err, "oss配置错误")
	}

	err = bucket.PutObject(objectKey, bytes.NewReader(data))
	if err != nil {
		return errors.Wrap(err, "oss上传失败")
	}
	return
}
