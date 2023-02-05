package repoimpl

import (
	"context"
	"errors"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"

	"github.com/wanyuqin/etcd-ui/backend/internal/domain/model"
	"github.com/wanyuqin/etcd-ui/backend/internal/domain/repository"
	"github.com/wanyuqin/etcd-ui/backend/internal/domain/valobj"
	"github.com/wanyuqin/etcd-ui/backend/internal/infrastructure/po"
)

type CertificateRepoImpl struct {
	po.BaseDao
}

var _ repository.CertificateRepo = &CertificateRepoImpl{}

func NewCertificateRepoImpl() *CertificateRepoImpl {
	return &CertificateRepoImpl{
		BaseDao: po.NewBaseDao(context.Background(), po.NewCertificate().TableName()),
	}
}

func (c *CertificateRepoImpl) CreateCertificate(mc *model.Certificate) error {
	tx := c.DB.Begin()
	pc := po.NewCertificate()
	copier.Copy(pc, mc)

	err := tx.Save(pc).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return err
}

func (c *CertificateRepoImpl) GetCertificateById(id uint) (*model.Certificate, error) {
	pc := po.NewCertificate()
	c.DB.Where("id = ?", id).Find(pc)
	return pc.ToModel()
}

func (c *CertificateRepoImpl) DeleteCertificateById(id uint) error {
	tx := c.DB.Begin()

	pc := po.NewCertificate()
	pc.ID = uint(id)

	err := tx.Delete(pc).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (c *CertificateRepoImpl) UpdateCertificate(mc *model.Certificate) error {
	tx := c.DB.Begin()
	cc := po.NewCertificateByModel(mc)

	if err := tx.Updates(cc).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (c *CertificateRepoImpl) ListCertificate(mc *model.Certificate, p *valobj.PageInfo) ([]model.Certificate, int64, error) {
	var count int64
	tx := c.DB.Where("deleted_at is NULL")

	if mc.Name != "" {
		tx.Where("name like '%s'", mc.Name)
	}

	cl := po.NewCertificateList()
	tx.Order(po.DescCreatedAt).Limit(p.PageSize).Offset((p.Page) - 1*10).Count(&count).Find(&cl)
	cm, err := cl.ToModel()
	return cm, count, err
}

func (c *CertificateRepoImpl) GetCertificateByIds(ids []uint) ([]model.Certificate, error) {
	cl := po.NewCertificateList()
	err := c.DB.Where("id in (?)", ids).Find(&cl).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return cl.ToModel()
}
