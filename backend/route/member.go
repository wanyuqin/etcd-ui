package Route

import "github.com/gin-gonic/gin"

func MemberRoute(r *gin.RouterGroup) {
	v1 := r.Group("/members")
	{
		v1.GET("")
	}
}
