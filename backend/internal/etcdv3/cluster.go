package etcdv3

import (
	"github.com/wanyuqin/etcd-ui/backend/internal/domain/model"
)

func (c EtcdCli) ClusterStatus() ([]model.ClusterStatus, error) {
	sm := make([]model.ClusterStatus, 0)
	eps := c.Client.Endpoints()
	for _, v := range eps {
		status, err := c.Client.Status(c.Ctx, v)
		if err != nil {
			return sm, err
		}
		cs := model.NewClusterStatus(status)
		cs.Endpoint = v
		sm = append(sm, cs)
	}

	return sm, nil
}

func (c EtcdCli) MemberList() {

}
