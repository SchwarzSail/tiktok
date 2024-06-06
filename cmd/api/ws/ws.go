package ws

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/websocket"
	"tiktok/cmd/api/biz/pack"
	"tiktok/cmd/api/dal/mq"
	"tiktok/internal/utils"
)

var upgrader = websocket.HertzUpgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(ctx *app.RequestContext) bool {
		return true
	},
}

func ServerWs(ctx context.Context, c *app.RequestContext) {
	//升级为websocket
	u, err := pack.GetUserInfo(ctx)
	if err != nil {
		utils.LogrusObj.Error(err)
		return
	}
	toUid := c.Query("to_uid")
	err = upgrader.Upgrade(c, func(conn *websocket.Conn) {
		client := &Client{
			Uid:     u.ID,
			SendID:  toUid,
			Message: make(chan []byte, 256),
			Conn:    conn,
		}
		Manager.Register <- client
		list, err := mq.Consume(u.ID)
		if err != nil {
			return
		}
		for _, v := range list {
			client.Message <- v
		}
		go client.WriteMsg()
		client.GetMsg()
	})
	if err != nil {
		utils.LogrusObj.Debug(err)
	}
}
