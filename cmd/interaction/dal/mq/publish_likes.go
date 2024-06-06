package mq

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/streadway/amqp"
)

func PublishVideoLikes(data []byte, queueName string) (err error) {
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

	// 发布消息
	if err = ch.Publish(
		"direct_exchange", // Exchange 名称
		queueName,         // 路由键
		false,             // 不等待服务器响应
		false,             // 不设置额外参数
		amqp.Publishing{
			ContentType: "application/json",
			Body:        data,
		},
	); err != nil {
		return errors.Wrap(err, "mq:PublishLikes failed")
	}
	return
}
