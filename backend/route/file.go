package Route

import (
	"github.com/gin-gonic/gin"

	"github.com/wanyuqin/etcd-ui/backend/internal/handler"
)

func FileRoute(r *gin.RouterGroup) {
	v1 := r.Group("files")
	{
		v1.POST("")
		v1.POST("/certificates", handler.UploadCertificate)
	}
}
