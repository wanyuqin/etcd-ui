package Route

import (
	"github.com/gin-gonic/gin"

	"github.com/wanyuqin/etcd-ui/backend/internal/handler"
)

func ConnectionRoute(r *gin.RouterGroup) {
	v1 := r.Group("/connections")
	{
		v1.POST("", handler.CreateConnection)
		v1.GET("", handler.ListConnection)
		v1.PUT("/:id", handler.UpdateConnection)
		v1.GET("/:id", handler.GetConnection)
		v1.DELETE("/:id", handler.DeleteConnection)
		v1.POST("/init", handler.InitConnection)
	}
}
