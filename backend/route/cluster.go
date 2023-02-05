package Route

import (
	"github.com/gin-gonic/gin"

	"github.com/wanyuqin/etcd-ui/backend/internal/handler"
)

func ClusterRoute(r *gin.RouterGroup) {

	r.GET("/cluster:status", handler.ClusterStatus)

}
