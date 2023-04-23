package Route

import (
	"github.com/gin-gonic/gin"

	"github.com/wanyuqin/etcd-ui/backend/internal/handler"
)

func WebSocket(r *gin.RouterGroup) {
	r.Any("/connection", handler.Connect)
}
