package Route

import (
	"github.com/gin-gonic/gin"

	"github.com/wanyuqin/etcd-ui/backend/internal/handler"
)

// MemberRoute add	Adds a member into the cluster
// list	Lists all members in the cluster
// promote	Promotes a non-voting member in the cluster
// remove	Removes a member from the cluster
// update	Updates a member in the cluster
func MemberRoute(r *gin.RouterGroup) {
	v1 := r.Group("/members")
	{
		v1.GET("", handler.ListMember)
		v1.GET("/:id", handler.GetMember)
		v1.POST("", handler.CreateMember)
		v1.PUT("/:id", handler.UpdateMember)
		v1.DELETE("/:id", handler.RemoveMember)

	}
}
