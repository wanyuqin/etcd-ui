package server

import (
	"flag"
	"net/http"

	"github.com/gorilla/websocket"

	"github.com/wanyuqin/etcd-ui/backend/logger"
)

var addr = flag.String("addr", "localhost:8082", "http service address")

var upgrader = websocket.Upgrader{}

func RunWS() {

	flag.Parse()
	http.HandleFunc("/connect", connect)
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		logger.Errorf(err.Error())
	}
}

func connect(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logger.Errorf("upgrade failed: %v", err)
		return
	}

	defer c.Close()

	for {
		mt, msg, err := c.ReadMessage()
		if err != nil {
			logger.Errorf("read message failed: %v", err)
			break
		}

		err = c.WriteMessage(mt, msg)
		if err != nil {
			logger.Errorf("send message failed: %v", err)
			break
		}
	}
}
