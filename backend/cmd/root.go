package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"

	"github.com/wanyuqin/etcd-ui/backend/db"
	"github.com/wanyuqin/etcd-ui/backend/hub"
	"github.com/wanyuqin/etcd-ui/backend/internal/domain/service/connection"
	"github.com/wanyuqin/etcd-ui/backend/logger"
	"github.com/wanyuqin/etcd-ui/backend/server"
)

var rootCmd = &cobra.Command{
	Use:   "",
	Short: "Etcd-Ui is a ",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logger.Errorf("%v", err)
		os.Exit(1)
	}

	// 初始化数据库
	err := db.InitMysql()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	// 启动消息监听
	h := hub.NewHub()
	go h.Run()
	// 启动etcd客户端
	cs, err := connection.DefaultConnectionService()
	if err != nil {
		os.Exit(1)
	}
	err = cs.InitConnection()
	if err != nil {
		logger.Errorf("init connection failed: %v", err)
		os.Exit(1)
	}
	// 启动 http 服务
	server.RunHttp()

}
