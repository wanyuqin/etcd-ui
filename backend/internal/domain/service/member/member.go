package member

import (
	"github.com/wanyuqin/etcd-ui/backend/internal/domain/model"
	"github.com/wanyuqin/etcd-ui/backend/internal/etcdv3"
)

type MemberService struct {
}

func NewMemberService() *MemberService {
	return &MemberService{}
}

func (m *MemberService) CreateMember(peerAddrs []string) error {
	if err := model.CheckPeerAddrs(peerAddrs); err != nil {
		return err
	}

	return etcdv3.Cli.CreateMember(peerAddrs)
}

func (m *MemberService) MemberList() (model.MemberList, error) {
	return etcdv3.Cli.MemberList()
}
