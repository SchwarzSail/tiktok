package main

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/spf13/viper"
	"log"
	"net"
	"net/http"
	_ "net/http/pprof"
	"tiktok/cmd/video/config"
	"tiktok/cmd/video/dal/cache"
	"tiktok/cmd/video/dal/db"
	"tiktok/cmd/video/dal/es"
	"tiktok/cmd/video/rpc"
	video "tiktok/kitex_gen/video/videoservice"
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
	//
	db.InitMySQL()
	cache.LinkRedis()
	es.LinkEs()

	rpc.Init()
	klog.SetLevel(klog.LevelDebug)
}
func main() {
	go func() {
		log.Println(http.ListenAndServe("0.0.0.0:6060", nil))
	}()
	conf := config.Config
	r, err := etcd.NewEtcdRegistry([]string{conf.EtcdHost + ":" + conf.EtcdPort})
	if err != nil {
		klog.Fatal(err)
	}

	addr, err := net.ResolveTCPAddr("tcp", conf.System.Host+":"+conf.System.Port)
	if err != nil {
		panic(err)
	}
	svr := video.NewServer(new(VideoServiceImpl), server.WithServiceAddr(addr), server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "video"}), server.WithRegistry(r))

	err = svr.Run()
	if err != nil {
		panic(err)
	}
}
