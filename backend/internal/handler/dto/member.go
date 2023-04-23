package dto

import (
	"github.com/jinzhu/copier"

	"github.com/wanyuqin/etcd-ui/backend/internal/domain/model"
)

type AddMemberRequest struct {
	PeerURLs []string `json:"peer_URLs"`
}

type Member struct {
	ID         string   `json:"id"`
	Name       string   `json:"name"`
	PeerURLs   []string `json:"peer_URLs"`
	ClientURLs []string `json:"client_URLs"`
	IsLearner  bool     `json:"is_learner"`
}

type MemberList []Member

func NewMemberListResponse(ml model.MemberList) MemberList {
	members := make([]Member, len(ml))
	for i := range ml {
		copier.Copy(&members[i], ml[i])
	}
	return members
}
