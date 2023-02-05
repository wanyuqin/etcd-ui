package po

import (
	"github.com/jinzhu/copier"
	"gorm.io/gorm"

	"github.com/wanyuqin/etcd-ui/backend/internal/domain/model"
)

type Connection struct {
	gorm.Model

	Name          string `gorm:"column:name" db:"name" json:"name" form:"name"`                 //  自定义名称
	Endpoint      string `gorm:"column:endpoint" db:"endpoint" json:"endpoint" form:"endpoint"` //  连接端点
	Version       string `gorm:"column:version" db:"version" json:"version" form:"version"`     //  etcd 版本
	Type          int64  `gorm:"column:type" db:"type" json:"type" form:"type"`                 //  连接类型 1 默认
	Active        int    `gorm:"column:active" db:"active" json:"active" form:"active"`         // 是否激活
	CertificateID uint   `gorm:"column:certificate_id;default:0"`                               // 关联证书ID
	Tls           int    `gorm:"column:tls;default:1"`                                          // 是否开启tls 1 关闭 2 开启
}

type ConnectionList []Connection

func (c *Connection) TableName() string {
	return "connection"
}

func NewConnection() *Connection {
	return &Connection{}
}

func NewConnectionByModel(mc *model.Connection) *Connection {
	pc := NewConnection()
	copier.Copy(pc, mc)

	pc.CertificateID = mc.Certificate.ID
	pc.Tls = int(mc.Tls)
	return pc

}

func NewConnectionList() ConnectionList {
	return make([]Connection, 0)
}

func (c *Connection) ToModel() (*model.Connection, error) {
	conn := model.NewConnection()
	cc := model.NewCertificate()
	copier.Copy(conn, c)

	cc.ID = c.CertificateID
	conn.CreatedAt = c.CreatedAt.Format("2006-01-02")
	conn.Type = model.ConnectionType(c.Type)
	conn.Active = model.ConnectionActive(c.Active)
	conn.Tls = model.ConnectionTls(c.Tls)
	conn.Certificate = *cc
	return conn, nil
}

func (c ConnectionList) ToModel() ([]model.Connection, error) {
	cm := make([]model.Connection, len(c))

	for i := range c {
		m, err := c[i].ToModel()
		if err != nil {
			return cm, err
		}
		cm[i] = *m
	}
	return cm, nil
}
