package po

import (
	"github.com/jinzhu/copier"
	"gorm.io/gorm"

	"github.com/wanyuqin/etcd-ui/backend/internal/domain/model"
)

type Certificate struct {
	gorm.Model

	Name string `gorm:"column:name;NOT NULL"` // 自定义名称
	Key  string `gorm:"column:key"`           // key
	Ca   string `gorm:"column:ca"`
	Cert string `gorm:"column:cert"`
}

type CertificateList []Certificate

func (c *Certificate) TableName() string {
	return "certificate"
}

func NewCertificate() *Certificate {
	return &Certificate{}
}

func NewCertificateByModel(mc *model.Certificate) *Certificate {
	cc := NewCertificate()
	copier.Copy(cc, mc)

	return cc
}

func NewCertificateList() CertificateList {
	return make([]Certificate, 0)
}

func (c *Certificate) ToModel() (*model.Certificate, error) {
	mc := model.NewCertificate()
	copier.Copy(mc, c)

	mc.CreatedAt = c.CreatedAt.Format("2006-01-15")
	return mc, nil
}

func (c CertificateList) ToModel() ([]model.Certificate, error) {
	mcs := make([]model.Certificate, len(c))
	for i := range c {
		mc, err := c[i].ToModel()
		if err != nil {
			return mcs, err
		}
		mcs[i] = *mc
	}

	return mcs, nil
}
