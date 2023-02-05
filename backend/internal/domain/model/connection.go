package model

import (
	"errors"

	"github.com/wanyuqin/etcd-ui/backend/x/xslice"
)

type Connection struct {
	ID          uint             `json:"id"`
	Name        string           `json:"name" `   //  自定义名称
	Version     string           `json:"version"` //  etcd 版本
	Type        ConnectionType   `json:"type"`    //  连接类型 1 默认
	Endpoint    string           `json:"endpoint"`
	CreatedAt   string           `json:"created_at"`
	Active      ConnectionActive `json:"active"`
	Tls         ConnectionTls    `json:"tls"`
	Certificate Certificate      `json:"certificate"`
}

type ConnectionList []Connection

type ConnectionType int64

type ConnectionActive int

type ConnectionTls int

const (
	TlsClose ConnectionTls = iota + 1
	TlsOpen
)

const (
	Default ConnectionType = iota + 1
	KeyAuth
	JwtAuth
	BaseAuth
)

const (
	DeActive ConnectionActive = iota
	Active
)

func NewConnection() *Connection {
	return &Connection{}
}

func (c *Connection) Validate() error {
	if c.Endpoint == "" {
		return errors.New("endpoint cannot be null")
	}

	if c.Name == "" {
		return errors.New("name cannot be null")
	}

	if c.Tls == TlsOpen && c.Certificate.ID <= 0 {
		return errors.New("unknown certificate ")
	}

	return nil
}

func NewConnectionList() ConnectionList {
	return make([]Connection, 0)
}

func (c ConnectionList) GetCertificateIds() []uint {
	ids := make([]uint, 0, len(c))
	for _, v := range c {
		if v.Certificate.ID > 0 {
			ids = append(ids, v.Certificate.ID)
		}

	}

	return xslice.UniqueSlice(ids)
}

func (c ConnectionList) SetCertificate(m map[uint]Certificate) ConnectionList {
	for i, _ := range c {
		c[i].Certificate = m[c[i].Certificate.ID]
	}

	return c
}
