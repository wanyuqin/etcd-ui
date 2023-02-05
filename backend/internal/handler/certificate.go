package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"

	"github.com/wanyuqin/etcd-ui/backend/internal/domain/model"
	"github.com/wanyuqin/etcd-ui/backend/internal/domain/service/certificate"
	"github.com/wanyuqin/etcd-ui/backend/internal/domain/valobj"
	"github.com/wanyuqin/etcd-ui/backend/internal/handler/dto"
	"github.com/wanyuqin/etcd-ui/backend/x/xgin"
)

func CreateCertificate(c *gin.Context) {
	cc := dto.Certificate{}
	err := c.ShouldBindJSON(&cc)
	if err != nil {
		xgin.Failed(c, err)
		return
	}

	mc := model.NewCertificate()
	copier.Copy(mc, cc)

	cs, err := certificate.NewCertificateService(certificate.WithMysqlCertificateRepository())
	if err != nil {
		xgin.Failed(c, err)
		return
	}

	err = cs.CreateCertificate(mc)

	xgin.Response(c, nil, err)

}

func UpdateCertificate(c *gin.Context) {
	id, err := xgin.ParamInt(c, "id")
	if err != nil {
		xgin.Failed(c, err)
		return
	}
	cc := dto.Certificate{}

	err = c.ShouldBindJSON(&cc)
	if err != nil {
		xgin.Failed(c, err)
		return
	}

	cc.ID = uint(id)

	sc, err := certificate.NewCertificateService(certificate.WithMysqlCertificateRepository())
	if err != nil {
		xgin.Failed(c, err)
		return
	}

	mc := model.NewCertificate()
	copier.Copy(mc, cc)

	err = sc.UpdateCertificate(mc)
	if err != nil {
		xgin.Failed(c, err)
		return
	}

	xgin.Response(c, nil, err)

}

func GetCertificate(c *gin.Context) {
	id, err := xgin.ParamInt64(c, "id")
	if err != nil {
		xgin.Failed(c, err)
		return
	}

	cs, err := certificate.NewCertificateService(certificate.WithMysqlCertificateRepository())
	if err != nil {
		xgin.Failed(c, err)
		return
	}

	cc, err := cs.GetCertificate(uint(id))
	if err != nil {
		xgin.Failed(c, err)
		return
	}

	xgin.Response(c, cc, err)
}

func DeleteCertificate(c *gin.Context) {
	id, err := xgin.ParamInt64(c, "id")
	if err != nil {
		xgin.Failed(c, err)
		return
	}

	cfgs := []certificate.CertificateConfiguration{
		certificate.WithMysqlCertificateRepository(),
		certificate.WithMysqlConnectionRepository(),
	}
	cs, err := certificate.NewCertificateService(cfgs...)
	if err != nil {
		xgin.Failed(c, err)
		return
	}

	err = cs.DeleteCertificate(uint(id))
	xgin.Response(c, nil, err)
}

func ListCertificate(c *gin.Context) {
	query := dto.PageInfo{}
	err := c.BindQuery(&query)
	if err != nil {
		xgin.Failed(c, err)
		return
	}

	name := c.Query("name")
	mc := model.NewCertificate()
	mc.Name = name

	p := valobj.NewPageInfo()
	copier.Copy(p, query)

	sc, err := certificate.NewCertificateService(certificate.WithMysqlCertificateRepository())
	if err != nil {
		xgin.Failed(c, err)
		return
	}

	mcs, count, err := sc.ListCertificate(mc, p)
	xgin.ResponsePage(c, mcs, count, err)

}
