package repository

import (
	"github.com/wanyuqin/etcd-ui/backend/internal/domain/model"
	"github.com/wanyuqin/etcd-ui/backend/internal/domain/valobj"
)

type ConnectionRepo interface {
	CreateConnection(conn model.Connection) error
	DeleteConnectionById(id uint) error
	GetConnectionById(id uint) (*model.Connection, error)
	GetConnection(conn *model.Connection) ([]model.Connection, error)
	ListConnection(conn *model.Connection, p *valobj.PageInfo) ([]model.Connection, int64, error)
	ActiveConnection(conn *model.Connection) error
}
