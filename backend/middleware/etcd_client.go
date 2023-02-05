package middleware

import (
	"github.com/gin-gonic/gin"
)

func CheckEtcdClient(c *gin.Context) {
	// if etcdv3.Cli == nil {
	// 	err := etcdv3.NewDefaultClientV3()
	// 	if err != nil {
	// 		logger.Errorf("new etcd client failed: %v", err)
	// 		return
	// 	}
	// }
	c.Next()
}
