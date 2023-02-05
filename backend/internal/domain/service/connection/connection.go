package connection

import (
	"errors"

	"github.com/jinzhu/copier"

	"github.com/wanyuqin/etcd-ui/backend/internal/domain/model"
	"github.com/wanyuqin/etcd-ui/backend/internal/domain/repository"
	"github.com/wanyuqin/etcd-ui/backend/internal/domain/valobj"
	"github.com/wanyuqin/etcd-ui/backend/internal/etcdv3"
	"github.com/wanyuqin/etcd-ui/backend/internal/infrastructure/repoimpl"
	"github.com/wanyuqin/etcd-ui/backend/logger"
)

type ConnectionService struct {
	connection  repository.ConnectionRepo
	certificate repository.CertificateRepo
}

type ConnectionConfiguration func(cs *ConnectionService) error

func NewConnectionService(cfgs ...ConnectionConfiguration) (*ConnectionService, error) {
	cs := &ConnectionService{}
	for _, cfg := range cfgs {
		if err := cfg(cs); err != nil {
			return nil, err
		}
	}
	return cs, nil
}

func WithConnectionRepository(cr repository.ConnectionRepo) ConnectionConfiguration {
	return func(cs *ConnectionService) error {
		cs.connection = cr
		return nil
	}
}

func WithMysqlConnectionRepository() ConnectionConfiguration {
	return func(cs *ConnectionService) error {
		cs.connection = repoimpl.NewConnectionRepoImpl()
		return nil
	}
}

func WithMysqlCertificateRepository() ConnectionConfiguration {
	return func(cs *ConnectionService) error {
		cs.certificate = repoimpl.NewCertificateRepoImpl()
		return nil
	}
}

func (c *ConnectionService) CreateConnection(conn model.Connection) error {
	if err := conn.Validate(); err != nil {
		return err
	}
	if conn.Tls == model.TlsOpen {
		cert, err := c.certificate.GetCertificateById(conn.Certificate.ID)
		if err != nil {
			logger.Errorf("GetCertificateById failed: %v", err)
			return err
		}

		if cert == nil || cert.ID <= 0 {
			return errors.New("invalid certificate")
		}
	}

	return c.connection.CreateConnection(conn)
}

func (c *ConnectionService) DeleteConnection(id uint) error {
	return c.connection.DeleteConnectionById(id)
}

func (c *ConnectionService) GetConnection(id uint) (*model.Connection, error) {
	return c.connection.GetConnectionById(id)
}

func (c *ConnectionService) ListConnection(conn *model.Connection, p *valobj.PageInfo) ([]model.Connection, int64, error) {
	p.Validate()
	cl, count, err := c.connection.ListConnection(conn, p)
	if err != nil {
		logger.Errorf("ListConnection failed: %v", err)
		return nil, count, err
	}

	mcl := model.ConnectionList(cl)
	cids := mcl.GetCertificateIds()
	ccs, err := c.certificate.GetCertificateByIds(cids)
	if err != nil {
		logger.Errorf("GetCertificateByIds failed: %v", err)
		return mcl, count, err
	}
	if len(ccs) > 0 {
		ccl := model.CertificateList(ccs)
		mcl.SetCertificate(ccl.Map())
	}

	return mcl, count, err

}

func (c *ConnectionService) ActiveConnection(id uint, conn *model.Connection) error {
	cb, err := c.connection.GetConnectionById(id)
	if err != nil {
		logger.Errorf("get connection from db failed: %v", err)
		return err
	}

	if cb.Tls == model.TlsOpen {
		certificate, err := c.certificate.GetCertificateById(cb.Certificate.ID)
		if err != nil {
			logger.Errorf("GetCertificateById failed: %v", err)
			return err
		}
		cb.Certificate = *certificate
	}

	switch cb.Type {
	case model.Default:
		dc := etcdv3.NewDefaultClient()
		copier.Copy(dc, cb)
		dc.Tls = int(cb.Tls)
		dc.Key = cb.Certificate.Key
		dc.Ca = cb.Certificate.Ca
		dc.Cert = cb.Certificate.Cert
		err = dc.NewClient()

		if err != nil {
			logger.Errorf("new default etcd client failed: %v", err)
			return err
		}
	}
	conn.ID = uint(id)
	// 数据更新
	return c.connection.ActiveConnection(conn)

}
