package rpc

import video "tiktok/kitex_gen/video/videoservice"

var (
	videoClient video.Client
)

func Init() {
	InitVideoRPC()
}
