package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"

	"github.com/wanyuqin/etcd-ui/backend/internal/domain/model"
	"github.com/wanyuqin/etcd-ui/backend/internal/etcdv3"
	"github.com/wanyuqin/etcd-ui/backend/internal/handler/dto"
	"github.com/wanyuqin/etcd-ui/backend/x/xgin"
)

func ListKeys(c *gin.Context) {
	kt, err := etcdv3.Cli.ListKeys()
	xgin.Response(c, kt, err)
}

func GetKV(c *gin.Context) {
	name := c.Query("name")
	kv, err := etcdv3.Cli.GetKey(name)
	xgin.Response(c, kv, err)
}

func PutKV(c *gin.Context) {
	kv := dto.KV{}
	err := c.ShouldBindJSON(&kv)
	if err != nil {
		xgin.Failed(c, err)
		return
	}
	mkv := model.KV{}
	copier.Copy(&mkv, kv)
	err = etcdv3.Cli.PutKey(mkv)
	xgin.Response(c, nil, err)
}

func DeleteKV(c *gin.Context) {
	name := c.Query("name")
	deleted, err := etcdv3.Cli.DeleteKey(name)
	xgin.Response(c, deleted, err)
}

func WatchKey(c *gin.Context) {
	name := c.Query("name")
	etcdv3.Cli.WatchKey(name)
	xgin.Success(c)
}
