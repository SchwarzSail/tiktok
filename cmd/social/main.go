package main

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/spf13/viper"
	"log"
	"net"
	"tiktok/cmd/social/dal/cache"
	"tiktok/cmd/social/dal/db"
	"tiktok/cmd/social/rpc"
	"tiktok/config"
	social "tiktok/kitex_gen/social/socialservice"
)

func init() {
	//初始化配置文件
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../config")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := viper.Unmarshal(&config.Config); err != nil {
		panic(err)
	}
	//dal的初始化
	db.InitMySQL()
	cache.LinkRedis()

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

	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8884")
	if err != nil {
		panic(err)
	}
	svr := social.NewServer(new(SocialServiceImpl), server.WithServiceAddr(addr), server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "social"}), server.WithRegistry(r))

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
