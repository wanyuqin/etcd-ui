package etcdv3

import (
	"github.com/wanyuqin/etcd-ui/backend/internal/domain/model"
)

func (e EtcdCli) ClusterStatus() ([]model.ClusterStatus, error) {
	sm := make([]model.ClusterStatus, 0)
	eps := e.Client.Endpoints()
	for _, v := range eps {
		status, err := e.Client.Status(e.Ctx, v)
		if err != nil {
			return sm, err
		}
		cs := model.NewClusterStatus(status)
		cs.Endpoint = v
		sm = append(sm, cs)
	}

	return sm, nil
}

func (e EtcdCli) MemberList() (model.MemberList, error) {
	ml, err := e.Client.MemberList(e.Ctx)
	if err != nil {
		return nil, err
	}

	mml := model.NewMemberList(ml.Members)

	return mml, nil
}

func (e EtcdCli) CreateMember(peerAddrs []string) error {
	_, err := e.Client.MemberAdd(e.Ctx, peerAddrs)
	if err != nil {
		return err
	}
	return nil
}
