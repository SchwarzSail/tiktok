package mq

import (
	"github.com/streadway/amqp"
	"tiktok/cmd/api/config"
)

var RabbitmqConn *amqp.Connection

func LinkRabbitmq() {
	conf := config.Config.Rabbitmq
	conn, err := amqp.Dial("amqp://" + conf.RabbitmqUser + ":" + conf.RabbitmqPassword + "@" + conf.RabbitmqHost + ":" + conf.RabbitmqPort + "/")
	if err != nil {
		panic(err)
	}
	RabbitmqConn = conn
}
