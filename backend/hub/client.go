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
	ID  string
	Hub *Hub
	//  websocket 连接
	Conn   *websocket.Conn
	Send   chan []byte
	Ticker *time.Ticker
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
			logger.Debugf("online client %v", c.Hub.OnlineClientCount())
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
			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				logger.Errorf("write ping message failed: %v", err)
				return
			}
		}
	}
}

func (c *Client) ReadPump() {
	defer func() {
		c.Ticker.Stop()
		c.Hub.Unregister <- c
		c.Conn.Close()
	}()

	c.Conn.SetReadDeadline(time.Now().Add(pongWait))
	c.Conn.SetPongHandler(func(string) error {
		c.Conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})
	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			logger.Errorf("%v", err)
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				logger.Errorf("%v", err)
			}
			break
		}
		logger.Debugf(" %s client send msg %s , total client %v", c.ID, string(message), len(c.Hub.Clients))
		if string(message) == "ping" {
			c.Send <- []byte("pong")
		}

		c.Ticker.Reset(10 * time.Second)
	}
}

func (c *Client) HearBeat() {
	defer func() {
		c.Ticker.Stop()
		c.Hub.Unregister <- c
		c.Conn.Close()
	}()
	for {
		select {
		case <-c.Ticker.C:
			logger.Debugf("client %s offline ", c.ID)
			return
		}
	}
}
