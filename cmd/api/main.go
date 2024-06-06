// Code generated by hertz generator.

package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/spf13/viper"
	"tiktok/cmd/api/biz/rpc"
	"tiktok/cmd/api/dal/cache"
	"tiktok/cmd/api/dal/db"
	"tiktok/cmd/api/dal/es"
	"tiktok/cmd/api/dal/mq"
	"tiktok/cmd/api/ws"
	"tiktok/config"
	"tiktok/internal/utils"
)

func init() {
	//配置文件初始化
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../config")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := viper.Unmarshal(&config.Config); err != nil {
		panic(err)
	}

	//logger
	utils.InitLog()
	//dal, 测试
	mq.LinkRabbitmq()
	cache.LinkRedis()
	db.InitMySQL()
	es.LinkEs()
	rpc.Init()

}

func main() {
	go ws.Manager.Start()
	h := server.Default(
		server.WithHostPorts("0.0.0.0:10001"),
		server.WithMaxRequestBodySize(419430400),
	)

	register(h)
	h.Spin()
}