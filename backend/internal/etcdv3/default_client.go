package etcdv3

import (
	"crypto/tls"
	"crypto/x509"
	"sync"

	clientv3 "go.etcd.io/etcd/client/v3"

	"github.com/wanyuqin/etcd-ui/backend/logger"
)

type DefaultClient struct {
	Endpoint string `json:"endpoint"`

	Ca   string `json:"ca"`
	Cert string `json:"cert"`
	Key  string `json:"key"`
	Tls  int    `json:"tls"`

	mux sync.Mutex
}

const (
	TlsClose int = iota + 1
	TlsOpen
)

func NewDefaultClient() *DefaultClient {
	return &DefaultClient{}
}

func (d *DefaultClient) NewClient() error {
	d.mux.Lock()
	defer d.mux.Unlock()

	cfg := clientv3.Config{
		Endpoints:   []string{d.Endpoint},
		DialTimeout: dialTimeout,
	}

	if d.Tls == TlsOpen {
		tc, err := d.NewTlsConfig()
		if err != nil {
			logger.Errorf("new tls config failed: %v", err)
			return err
		}
		cfg.TLS = tc
	}

	client, err := clientv3.New(cfg)
	if err != nil {
		logger.Errorf("create default etcd client failed: %v", err)
		return err
	}

	Cli = &EtcdCli{
		Client: client,
		Ctx:    client.Ctx(),
	}
	return nil
}

func (d *DefaultClient) NewTlsConfig() (*tls.Config, error) {
	var config tls.Config

	ct := []byte(d.Cert)
	kf := []byte(d.Key)
	ca := []byte(d.Ca)

	cert, err := tls.X509KeyPair(ct, kf)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	pool := x509.NewCertPool()
	pool.AppendCertsFromPEM(ca)
	config.Certificates = []tls.Certificate{cert}
	config.RootCAs = pool

	return &config, nil

}
