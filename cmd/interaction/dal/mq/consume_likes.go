package mq

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"tiktok/cmd/interaction/dal/db"
	"tiktok/cmd/interaction/dal/mq/model"
)

func ConsumeLikes(queueName string) (err error) {
	//声明通道
	ch, err := RabbitmqConn.Channel()
	if err != nil {
		fmt.Println(err)
	}
	defer ch.Close()
	//声明交换机
	if err = ch.ExchangeDeclare(
		"direct_exchange", // Exchange 名称 TODO:考虑写入config.yaml
		"direct",          // Exchange 类型
		true,              // 持久化
		false,             // 不自动删除
		false,             // 不等待服务器响应
		false,
		nil, // 不设置额外参数
	); err != nil {
		return errors.Wrap(err, "mq:Declare Exchange failed")
	}
	// 声明队列
	queue, err := ch.QueueDeclare(
		queueName, // 队列名称
		false,     // 不持久化
		false,     // 不自动删除
		false,     // 不独占
		false,     // 不等待服务器响应
		nil,       // 不设置额外参数
	)
	if err != nil {
		return errors.Wrap(err, "mq:Declare queue failed")
	}
	// 交换机和队列绑定
	if err = ch.QueueBind(
		queue.Name,        // 队列名称
		queueName,         // 路由键，用于绑定 Exchange 和队列
		"direct_exchange", // Exchange 名称
		false,             // 不等待服务器响应
		nil,               // 不设置额外参数
	); err != nil {
		return errors.Wrap(err, "mq:QueueBind failed")
	}
	likesDao := db.NewLikesDao(context.Background())
	// 消费消息
	for {
		msg, ok, err := ch.Get(queueName, true)
		if err != nil {
			return errors.Wrap(err, "mq.consume likes failed")
		}
		if !ok {
			break
		}
		var like model.Likes
		if err = json.Unmarshal(msg.Body, &like); err != nil {
			return errors.Wrap(err, "unmarshal like failed")
		}

		if like.ActionType == "1" {
			likes := db.Likes{
				VideoID:   like.VideoID,
				UserID:    like.UserID,
				CommentID: like.CommentID,
			}
			if err = likesDao.CreateLike(&likes); err != nil {
				return errors.WithMessage(err, "mq:CreateLike failed")
			}
		} else {
			if err = likesDao.DeleteLikes(like.UserID, like.VideoID, 0); err != nil {
				return errors.WithMessage(err, "mq:DeleteLike failed")
			}
		}
	}
	return nil
}
