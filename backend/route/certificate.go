package Route

import (
	"github.com/gin-gonic/gin"

	"github.com/wanyuqin/etcd-ui/backend/internal/handler"
)

func CertificateRoute(r *gin.RouterGroup) {
	v1 := r.Group("/certificates")
	{
		v1.POST("", handler.CreateCertificate)
		v1.GET("/:id", handler.GetCertificate)
		v1.GET("", handler.ListCertificate)
		v1.PUT("/:id", handler.UpdateCertificate)
		v1.DELETE("/:id", handler.DeleteCertificate)
	}
}
