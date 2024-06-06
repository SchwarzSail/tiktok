package rpc

import (
	interaction "tiktok/kitex_gen/interaction/interactionservice"
	social "tiktok/kitex_gen/social/socialservice"
	"tiktok/kitex_gen/user/userservice"
	"tiktok/kitex_gen/video/videoservice"
)

var (
	userClient        userservice.Client
	videoClient       videoservice.Client
	interactionClient interaction.Client
	socialClient      social.Client
)

func Init() {
	InitUserRPC()
	InitVideoRPC()
	InitInteractionPRC()
	InitSocialRPC()
}
