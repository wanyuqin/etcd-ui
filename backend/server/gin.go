package server

import (
	"github.com/gin-gonic/gin"

	"github.com/wanyuqin/etcd-ui/backend/logger"
	"github.com/wanyuqin/etcd-ui/backend/middleware"
	Route "github.com/wanyuqin/etcd-ui/backend/route"
)

var MaxMultipartMemory int64 = 8 << 20

func RunHttp() {
	gin.SetMode("debug")
	engine := gin.New()
	engine.MaxMultipartMemory = MaxMultipartMemory
	engine.Use(gin.LoggerWithFormatter(middleware.Log))
	engine.Use(middleware.Cros)
	engine.Use(middleware.CheckEtcdClient)

	r := engine.RouterGroup

	v1 := r.Group("/v1")

	Route.KvRoute(v1)
	Route.ClusterRoute(v1)
	Route.MemberRoute(v1)
	Route.WebSocket(v1)
	Route.ConnectionRoute(v1)
	Route.CertificateRoute(v1)
	Route.FileRoute(v1)
	err := engine.Run(":8081")
	if err != nil {
		logger.Error(err)
	}
}
