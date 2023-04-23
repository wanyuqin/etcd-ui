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

type ConnectionRepoImpl struct {
	po.BaseDao
}

func NewConnectionRepoImpl() *ConnectionRepoImpl {
	return &ConnectionRepoImpl{
		BaseDao: po.NewBaseDao(context.Background(), po.NewConnection().TableName()),
	}
}

var (
	_ repository.ConnectionRepo = &ConnectionRepoImpl{}

	ErrDuplicateName = errors.New("duplicate name")
)

func (c *ConnectionRepoImpl) CreateConnection(conn model.Connection) error {
	pc := po.NewConnection()
	tx := c.DB.Begin()
	err := tx.Where("name=?", conn.Name).First(pc).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		tx.Rollback()
		return err
	}
	if pc.ID > 0 {
		tx.Rollback()
		return ErrDuplicateName
	}

	copier.Copy(pc, conn)
	pc.CertificateID = conn.Certificate.ID

	err = tx.Save(pc).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return err
}

func (c *ConnectionRepoImpl) GetConnectionById(id uint) (*model.Connection, error) {
	pc := po.NewConnection()
	err := c.DB.Where("id = ?", id).Find(pc).Error
	if err != nil {
		return nil, err
	}
	return pc.ToModel()
}

func (c *ConnectionRepoImpl) DeleteConnectionById(id uint) error {
	tx := c.DB.Begin()
	err := tx.Where("id = ?", id).Delete(po.NewConnection()).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (c *ConnectionRepoImpl) ListConnection(conn *model.Connection, p *valobj.PageInfo) ([]model.Connection, int64, error) {
	var count int64
	tx := c.DB.Where("deleted_at is NULL")
	if conn.Name != "" {
		tx.Where("name like '%s'", conn.Name)
	}
	cl := po.NewConnectionList()
	tx.Order(po.DescCreatedAt).Limit(p.PageSize).Offset((p.Page) - 1*10).Count(&count).Find(&cl)
	cm, err := cl.ToModel()

	return cm, count, err
}

func (c *ConnectionRepoImpl) GetConnection(conn *model.Connection) ([]model.Connection, error) {
	pc := po.NewConnectionByModel(conn)
	tx := c.DB.Where(po.WhereTrue)

	if pc.Name != "" {
		tx.Where("name = ?", pc.Name)
	}

	if pc.Tls > 0 {
		tx.Where("tls = ?", pc.Tls)
	}

	if pc.CertificateID > 0 {
		tx.Where("certificate_id = ?", pc.CertificateID)
	}
	if pc.Active > 0 {
		tx.Where("active = ?", pc.Active)
	}

	list := po.NewConnectionList()
	tx.Find(&list)
	return list.ToModel()

}

func (c *ConnectionRepoImpl) ActiveConnection(conn *model.Connection) error {

	tx := c.DB.Begin()
	err := tx.Update("active", model.DeActive).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Where("id = ?", conn.ID).Update("active", model.Active).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
