package model

import (
	"errors"
)

type Certificate struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Key       string `json:"key"`
	Ca        string `json:"ca"`
	Cert      string `json:"cert"`
	CreatedAt string `json:"created_at"`
}

type CertificateList []Certificate

func NewCertificate() *Certificate {
	return &Certificate{}
}

func NewCertificateList() CertificateList {
	return make([]Certificate, 0)
}

// Validate TODO 校验 cert key
func (c *Certificate) Validate() error {
	if c.Name == "" {
		return errors.New("name cannot be null")
	}

	if c.Key == "" {
		return errors.New("key cannot be null")
	}

	if c.Cert == "" {
		return errors.New("cert cannot be null")
	}

	return nil
}

func (c *Certificate) ValidateCert() {
}

func (c CertificateList) Map() map[uint]Certificate {
	m := make(map[uint]Certificate, len(c))
	for _, v := range c {
		m[v.ID] = v
	}

	return m
}
