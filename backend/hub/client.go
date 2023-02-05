package hub

import (
	"time"

	"github.com/gorilla/websocket"

	"github.com/wanyuqin/etcd-ui/backend/logger"
)

const (
	writeWait = 10 * time.Second

	pongWait = 60 * time.Second

	pingPeriod = (pongWait * 9) / 10
)

type Client struct {
	Hub *Hub

	//  websocket 连接
	Conn *websocket.Conn

	//
	Send chan []byte
}

func (c *Client) Notify() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.Conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.Send:

			if c.Conn == nil {
				return
			}

			if !ok {
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.Conn.NextWriter(websocket.TextMessage)
			if err != nil {
				logger.Errorf("create next writer failed: %v", err)
				return
			}

			_, err = w.Write(message)
			if err != nil {
				logger.Errorf("writer message failed: %v", err)
				return
			}
			n := len(c.Send)
			for i := 0; i < n; i++ {
				_, err = w.Write(<-c.Send)
				if err != nil {
					logger.Errorf("writer message failed: %v", err)
					return
				}
			}
		case <-ticker.C:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))

			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				logger.Errorf("write ping message failed: %v", err)
				return
			}

		}
	}
}
