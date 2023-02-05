package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/wanyuqin/etcd-ui/backend/internal/etcdv3"
	"github.com/wanyuqin/etcd-ui/backend/x/xgin"
)

func ClusterStatus(c *gin.Context) {
	s, err := etcdv3.Cli.ClusterStatus()
	xgin.Response(c, s, err)
}
