package etcdv3

import (
	"encoding/json"

	clientv3 "go.etcd.io/etcd/client/v3"

	"github.com/wanyuqin/etcd-ui/backend/hub"
	"github.com/wanyuqin/etcd-ui/backend/internal/domain/model"
	"github.com/wanyuqin/etcd-ui/backend/logger"
)

func (c EtcdCli) ListKeys() ([]model.KeyTree, error) {
	ks := make([]string, 0)
	res, err := c.Client.Get(c.Ctx, "", clientv3.WithKeysOnly(), clientv3.WithPrefix())
	if err != nil {
		logger.Errorf("etcd v3 get failed: %v", err)
		return nil, err
	}

	for _, v := range res.Kvs {
		ks = append(ks, string(v.Key))
	}

	return model.NewKeyTree(ks), nil
}

func (c EtcdCli) GetKey(name string) (*model.KV, error) {
	res, err := c.Client.Get(c.Ctx, name)
	if err != nil {
		logger.Errorf("etcd v3 get failed: %v", err)
		return nil, err
	}
	if len(res.Kvs) > 0 {
		kv := model.NewKV(res.Kvs[0])
		if kv.Lease != 0 {
			tl, err := c.Client.TimeToLive(c.Ctx, clientv3.LeaseID(kv.Lease))
			if err != nil {
				logger.Errorf("etcd v3 TimeToLive failed: %v", err)
				return nil, err
			}

			kv.Ttl = tl.TTL
		}
		return kv, nil
	}

	return nil, err
}

func (c EtcdCli) PutKey(kv model.KV) error {
	ops := make([]clientv3.OpOption, 0)
	if kv.Ttl > 0 {
		// 创建一个租约对象
		lease := clientv3.NewLease(c.Client)
		// 生成一个新的租约
		ls, err := lease.Grant(c.Ctx, kv.Ttl)
		if err != nil {
			logger.Errorf("grant lease failed: %v", err)
			return err
		}
		ops = append(ops, clientv3.WithLease(ls.ID))
	}

	_, err := c.Client.Put(c.Ctx, kv.Key, kv.Value, ops...)
	if err != nil {
		logger.Errorf("put key %s failed: %v", kv.Key, err)
		return err
	}
	return nil
}

func (c EtcdCli) DeleteKey(name string) (int64, error) {
	var deleted int64
	res, err := c.Client.Delete(c.Ctx, name)
	if err != nil {
		logger.Errorf("etcd v3 delete failed: %v", err)
		return deleted, err
	}
	deleted = res.Deleted
	return deleted, nil
}

func (c EtcdCli) WatchKey(name string) {
	logger.Debugf("staring to listing %s", name)
	go c.watchKey(name)
}

func (c EtcdCli) watchKey(name string) {
	ch := c.Client.Watch(c.Ctx, name)
	for {
		select {
		case msg := <-ch:
			logger.Debugf("watch message %v", msg)
			// 消息发送
			for _, v := range msg.Events {
				e := model.NewEvent(v)
				d, err := json.Marshal(e)
				if err != nil {
					logger.Errorf("json marshal failed: %v", err)
				}
				hub.H.Broadcast <- d
			}

		}
	}

}
