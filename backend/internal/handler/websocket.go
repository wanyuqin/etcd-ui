package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"

	"github.com/wanyuqin/etcd-ui/backend/hub"
	"github.com/wanyuqin/etcd-ui/backend/logger"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {

		return true
	},
}

func Connect(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		logger.Errorf("upgrade failed: %v", err)
		return
	}
	uid := uuid.New()

	client := &hub.Client{
		ID:     uid.String(),
		Hub:    hub.H,
		Conn:   conn,
		Send:   make(chan []byte, 256),
		Ticker: time.NewTicker(10 * time.Second),
	}
	logger.Debugf("%s client register", client.ID)
	client.Hub.Register <- client
	go client.Notify()
	go client.ReadPump()
	go client.HearBeat()
}
