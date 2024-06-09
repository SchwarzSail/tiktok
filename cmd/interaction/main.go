package main

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/spf13/viper"
	"log"
	"net"
	"tiktok/cmd/interaction/config"
	"tiktok/cmd/interaction/dal/cache"
	"tiktok/cmd/interaction/dal/db"
	"tiktok/cmd/interaction/dal/mq"
	"tiktok/cmd/interaction/rpc"
	interaction "tiktok/kitex_gen/interaction/interactionservice"
	"time"
)

func init() {
	//初始化配置文件
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := viper.Unmarshal(&config.Config); err != nil {
		panic(err)
	}
	//dal的初始化
	db.InitMySQL()
	cache.LinkRedis()
	mq.LinkRabbitmq()

	//rpc
	rpc.Init()
	klog.SetLevel(klog.LevelDebug)
}

func main() {
	conf := config.Config
	r, err := etcd.NewEtcdRegistry([]string{conf.EtcdHost + ":" + conf.EtcdPort})
	if err != nil {
		klog.Fatal(err)
	}

	addr, err := net.ResolveTCPAddr("tcp", conf.System.Domain+":"+conf.System.Port)
	if err != nil {
		panic(err)
	}
	svr := interaction.NewServer(new(InteractionServiceImpl), server.WithServiceAddr(addr), server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "interaction"}), server.WithRegistry(r))
	//不知道这个监听的代码该放在哪
	//开启另一个线程，mysql消费
	go func() {
		for {
			err = mq.ConsumeLikes("mysql:video_likes")
			if err != nil {
				klog.Error(err)
			}
			time.Sleep(5 * time.Second)
		}
	}()

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
