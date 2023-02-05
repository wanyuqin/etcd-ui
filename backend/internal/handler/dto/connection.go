package dto

import (
	"github.com/jinzhu/copier"

	"github.com/wanyuqin/etcd-ui/backend/internal/domain/model"
)

type Connection struct {
	Name          string `json:"name" `   //  自定义名称
	Version       string `json:"version"` //  etcd 版本
	Type          int64  `json:"type"`    //  连接类型 1 默认
	Endpoint      string `json:"endpoint"`
	Active        int    `json:"active"`
	Tls           int    `json:"tls"`
	CertificateId uint   `json:"certificate_id"`
}

type ConnectionResponse struct {
	ID              uint   `json:"id"`
	Name            string `json:"name" `   //  自定义名称
	Version         string `json:"version"` //  etcd 版本
	Type            int64  `json:"type"`    //  连接类型 1 默认
	Endpoint        string `json:"endpoint"`
	Active          int    `json:"active"`
	Tls             int    `json:"tls"`
	CertificateName string `json:"certificate_name"`
	CertificateId   uint   `json:"certificate_id"`
	CreatedAt       string `json:"created_at"`
}

type ConnectionResponseList []ConnectionResponse

func NewConnectionResponseList() ConnectionResponseList {
	return make([]ConnectionResponse, 0)
}

func NewConnectionResponseByModel(mc model.Connection) ConnectionResponse {
	cr := ConnectionResponse{}
	copier.Copy(&cr, mc)

	cr.CertificateName = mc.Certificate.Name
	cr.CertificateId = mc.Certificate.ID

	return cr
}

func NewConnectionResponseListByModel(mc []model.Connection) ConnectionResponseList {
	crl := make([]ConnectionResponse, len(mc))

	for i := range mc {
		crl[i] = NewConnectionResponseByModel(mc[i])
	}

	return crl
}
