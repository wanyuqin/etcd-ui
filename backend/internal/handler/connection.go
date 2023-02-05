package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"

	"github.com/wanyuqin/etcd-ui/backend/internal/domain/model"
	"github.com/wanyuqin/etcd-ui/backend/internal/domain/service/connection"
	"github.com/wanyuqin/etcd-ui/backend/internal/domain/valobj"
	"github.com/wanyuqin/etcd-ui/backend/internal/handler/dto"
	"github.com/wanyuqin/etcd-ui/backend/x/xgin"
)

func CreateConnection(c *gin.Context) {
	param := dto.Connection{}
	if err := c.ShouldBindJSON(&param); err != nil {
		xgin.Failed(c, err)
		return
	}

	conn := model.NewConnection()
	cc := model.NewCertificate()
	copier.Copy(conn, param)

	cc.ID = param.CertificateId
	conn.Certificate = *cc

	cfgs := []connection.ConnectionConfiguration{
		connection.WithMysqlCertificateRepository(),
		connection.WithMysqlConnectionRepository(),
	}
	cs, err := connection.NewConnectionService(cfgs...)
	if err != nil {
		xgin.Failed(c, err)
		return
	}

	err = cs.CreateConnection(*conn)

	xgin.Response(c, nil, err)
}

func UpdateConnection(c *gin.Context) {
	id, err := xgin.ParamInt64(c, "id")
	if err != nil {
		xgin.Failed(c, err)
		return
	}
	conn := dto.Connection{}
	err = c.ShouldBindJSON(&conn)
	if err != nil {
		xgin.Failed(c, err)
		return
	}
	mc := model.NewConnection()
	copier.Copy(mc, conn)

	cfgs := []connection.ConnectionConfiguration{
		connection.WithMysqlCertificateRepository(),
		connection.WithMysqlConnectionRepository(),
	}
	sc, err := connection.NewConnectionService(cfgs...)
	err = sc.ActiveConnection(uint(id), mc)
	xgin.Response(c, nil, err)
}

func GetConnection(c *gin.Context) {

	id, err := xgin.ParamInt(c, "id")
	if err != nil {
		xgin.Failed(c, err)
		return
	}
	cs, err := connection.NewConnectionService(connection.WithMysqlConnectionRepository())
	conn, err := cs.GetConnection(uint(id))
	xgin.Response(c, conn, err)
}

func DeleteConnection(c *gin.Context) {
	id, err := xgin.ParamInt(c, "id")
	if err != nil {
		xgin.Failed(c, err)
		return
	}

	cs, err := connection.NewConnectionService(connection.WithMysqlConnectionRepository())
	if err != nil {
		xgin.Failed(c, err)
		return
	}
	err = cs.DeleteConnection(uint(id))
	xgin.Response(c, nil, err)

}

func ListConnection(c *gin.Context) {
	query := dto.PageInfo{}
	err := c.BindQuery(&query)
	if err != nil {
		xgin.Failed(c, err)
		return
	}

	name := c.Query("name")
	conn := model.NewConnection()
	conn.Name = name

	p := valobj.NewPageInfo()
	copier.Copy(p, query)
	cfgs := []connection.ConnectionConfiguration{
		connection.WithMysqlConnectionRepository(),
		connection.WithMysqlCertificateRepository(),
	}
	sc, err := connection.NewConnectionService(cfgs...)
	if err != nil {
		xgin.Failed(c, err)
		return
	}

	cnn, count, err := sc.ListConnection(conn, p)

	cl := dto.NewConnectionResponseListByModel(cnn)
	xgin.ResponsePage(c, cl, count, err)

}
