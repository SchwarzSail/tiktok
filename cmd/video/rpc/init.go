package rpc

import interaction "tiktok/kitex_gen/interaction/interactionservice"

var (
	interactionClient interaction.Client
)

func Init() {
	InitInteractionRPC()
}
