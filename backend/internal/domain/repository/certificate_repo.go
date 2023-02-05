package repository

import (
	"github.com/wanyuqin/etcd-ui/backend/internal/domain/model"
	"github.com/wanyuqin/etcd-ui/backend/internal/domain/valobj"
)

type CertificateRepo interface {
	CreateCertificate(mc *model.Certificate) error
	GetCertificateById(id uint) (*model.Certificate, error)
	GetCertificateByIds(ids []uint) ([]model.Certificate, error)
	DeleteCertificateById(id uint) error
	UpdateCertificate(mc *model.Certificate) error
	ListCertificate(mc *model.Certificate, p *valobj.PageInfo) ([]model.Certificate, int64, error)
}
