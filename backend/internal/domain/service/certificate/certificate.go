package certificate

import (
	"errors"

	"github.com/wanyuqin/etcd-ui/backend/internal/domain/model"
	"github.com/wanyuqin/etcd-ui/backend/internal/domain/repository"
	"github.com/wanyuqin/etcd-ui/backend/internal/domain/valobj"
	"github.com/wanyuqin/etcd-ui/backend/internal/infrastructure/repoimpl"
)

type CertificateService struct {
	certificate repository.CertificateRepo
	connection  repository.ConnectionRepo
}

type CertificateConfiguration func(cs *CertificateService) error

func NewCertificateService(cfgs ...CertificateConfiguration) (*CertificateService, error) {
	cs := &CertificateService{}
	for _, cfg := range cfgs {
		if err := cfg(cs); err != nil {
			return nil, err
		}

	}
	return cs, nil
}

func WithCertificateRepository(cr repository.CertificateRepo) CertificateConfiguration {
	return func(cs *CertificateService) error {
		cs.certificate = cr
		return nil
	}
}

func WithMysqlCertificateRepository() CertificateConfiguration {
	return func(cs *CertificateService) error {
		cs.certificate = repoimpl.NewCertificateRepoImpl()
		return nil
	}
}

func WithMysqlConnectionRepository() CertificateConfiguration {
	return func(cs *CertificateService) error {
		cs.connection = repoimpl.NewConnectionRepoImpl()
		return nil
	}
}

func (c *CertificateService) CreateCertificate(mc *model.Certificate) error {

	if err := mc.Validate(); err != nil {
		return err
	}
	return c.certificate.CreateCertificate(mc)
}

func (c *CertificateService) UpdateCertificate(mc *model.Certificate) error {
	err := mc.Validate()
	if err != nil {
		return err
	}
	return c.certificate.UpdateCertificate(mc)
}

func (c *CertificateService) DeleteCertificate(id uint) error {
	// 查询是否被绑定
	mc := model.NewConnection()
	cc := model.NewCertificate()
	cc.ID = uint(id)

	mc.Certificate = *cc
	cs, err := c.connection.GetConnection(mc)
	if err != nil {
		return err
	}
	if len(cs) > 0 {
		return errors.New("certificate has been used")
	}
	return c.certificate.DeleteCertificateById(id)
}

func (c *CertificateService) GetCertificate(id uint) (*model.Certificate, error) {
	return c.certificate.GetCertificateById(id)
}

func (c *CertificateService) ListCertificate(mc *model.Certificate, p *valobj.PageInfo) ([]model.Certificate, int64, error) {
	p.Validate()
	return c.certificate.ListCertificate(mc, p)
}
