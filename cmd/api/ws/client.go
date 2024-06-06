package ws

import (
	"encoding/json"
	"github.com/hertz-contrib/websocket"
	"tiktok/internal/utils"
)

type Client struct {
	Uid     string          `json:"uid"`     //发送方
	SendID  string          `json:"send_id"` //接收方
	Message chan []byte     `json:"message"`
	Conn    *websocket.Conn `json:"conn"`
}

type Message struct {
	From    string `json:"from"`
	Content string `json:"content"`
}

func (c *Client) GetMsg() {
	defer func() {
		Manager.Unregister <- c
		_ = c.Conn.Close()
	}()
	for {
		msg := &Message{}
		_, msgBuf, err := c.Conn.ReadMessage()
		if err != nil {
			utils.LogrusObj.Error(err)
			Manager.Unregister <- c
			_ = c.Conn.Close()
			return
		}
		_ = json.Unmarshal(msgBuf, &msg)
		utils.LogrusObj.Info(msg.Content)
		Manager.Broadcast <- &Broadcast{
			Client:  c,
			Message: msgBuf,
		}
	}
}

func (c *Client) WriteMsg() {
	defer func() {
		_ = c.Conn.Close()
	}()

	for message := range c.Message {
		replyMsg := &Message{
			From:    c.Uid,
			Content: string(message),
		}
		msg, _ := json.Marshal(replyMsg)
		err := c.Conn.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			utils.LogrusObj.Error(err)
			Manager.Unregister <- c
			_ = c.Conn.Close()
			return
		}
	}
	// 如果 c.Message 通道关闭，关闭连接
	_ = c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
}
