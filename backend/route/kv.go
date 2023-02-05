package Route

import (
	"github.com/gin-gonic/gin"

	"github.com/wanyuqin/etcd-ui/backend/internal/handler"
)

func KvRoute(r *gin.RouterGroup) {
	v1 := r.Group("/keys")
	{
		v1.GET("", handler.ListKeys)
		v1.GET("/value", handler.GetKV)
		v1.POST("", handler.PutKV)
		v1.DELETE("", handler.DeleteKV)
		v1.POST("/watch", handler.WatchKey)
	}
}
