package main

import (
	"log"
	"os"
	"time"

	"github.com/wanyuqin/etcd-ui/backend/db"
	"github.com/wanyuqin/etcd-ui/backend/hub"
	"github.com/wanyuqin/etcd-ui/backend/server"
)

var requestTimeout = 10 * time.Second

func main() {
	err := db.InitMysql()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	h := hub.NewHub()
	go h.Run()

	server.RunHttp()

}
