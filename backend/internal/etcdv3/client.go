package etcdv3

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"os"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"

	"github.com/wanyuqin/etcd-ui/backend/logger"
)

type EtcdCli struct {
	Client *clientv3.Client
	Ctx    context.Context
}

var Cli *EtcdCli

var (
	defaultEndpoints = []string{"127.0.0.1:2379"}

	dialTimeout = 5 * time.Second
)

type Client interface {
	NewClient() error
}

var keyFile = "./peer.key"
var certFile = "./peer.crt"
var caFile = "./ca.crt"

func NewClientV3(endpoints ...string) error {
	if len(endpoints) == 0 {
		endpoints = defaultEndpoints
	}
	tc, err := NewTlsConfig()
	if err != nil {
		return err
	}
	client, err := clientv3.New(
		clientv3.Config{
			Endpoints:   endpoints,
			DialTimeout: dialTimeout,
			TLS:         tc,
		})
	if err != nil {
		return err
	}

	Cli = &EtcdCli{
		Client: client,
		Ctx:    client.Ctx(),
	}
	return nil
}

func NewTlsConfig() (*tls.Config, error) {
	var config tls.Config
	cf, err := readFile(certFile)
	if err != nil {
		logger.Error("read cert file failed: %v", err)
		return nil, err
	}

	kf, err := readFile(keyFile)
	if err != nil {
		logger.Errorf("read key file failed:%v", err)
		return nil, err
	}

	cert, err := tls.X509KeyPair(cf, kf)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	ca, err := readFile(caFile)
	if err != nil {
		logger.Errorf("read ca file failed: %v", err)
		return nil, err
	}

	pool := x509.NewCertPool()
	pool.AppendCertsFromPEM(ca)
	config.Certificates = []tls.Certificate{cert}
	config.RootCAs = pool

	return &config, nil

}

func readFile(name string) ([]byte, error) {
	data, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func NewDefaultClientV3() error {
	client, err := clientv3.New(
		clientv3.Config{
			Endpoints:   []string{"127.0.0.1:2379"},
			DialTimeout: dialTimeout,
		})
	if err != nil {
		return err
	}

	Cli = &EtcdCli{
		Client: client,
		Ctx:    client.Ctx(),
	}
	return nil
}
