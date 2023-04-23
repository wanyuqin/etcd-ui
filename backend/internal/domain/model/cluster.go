package model

import (
	"errors"
	"fmt"
	"strings"

	humanize "github.com/dustin/go-humanize"
	"github.com/jinzhu/copier"
	"go.etcd.io/etcd/api/v3/etcdserverpb"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type ClusterStatus struct {
	ID               string `json:"id"`
	DbSize           string `json:"db_size"`
	DbSizeInUse      string `json:"db_size_in_use"`
	Endpoint         string `json:"endpoint"`
	Version          string `json:"version"`
	Reversion        int64  `json:"reversion"`
	IsLeader         bool   `json:"is_leader"`
	IsLearner        bool   `json:"is_learner"`
	RaftIndex        uint64 `json:"raft_index"`
	RaftTerm         uint64 `json:"raft_term"`
	RaftAppliedIndex uint64 `json:"raft_applied_index"`
	Error            string `json:"error"`
}

func NewClusterStatus(status *clientv3.StatusResponse) ClusterStatus {
	cs := ClusterStatus{
		ID:               fmt.Sprintf("%x", status.Header.MemberId),
		DbSize:           humanize.Bytes(uint64(status.DbSize)),
		DbSizeInUse:      humanize.Bytes(uint64(status.DbSizeInUse)),
		IsLearner:        status.IsLearner,
		IsLeader:         status.Header.MemberId == status.Leader,
		Version:          status.Version,
		RaftIndex:        status.RaftIndex,
		RaftTerm:         status.RaftTerm,
		RaftAppliedIndex: status.RaftAppliedIndex,
		Error:            strings.Join(status.Errors, ","),
		Reversion:        status.Header.Revision,
	}

	return cs
}

type Member struct {
	ID         string   `json:"id"`
	Name       string   `json:"name"`
	PeerURLs   []string `json:"peer_URLs"` // 集群内部通信地址
	ClientURLs []string `json:"client_URLs"`
	IsLearner  bool     `json:"is_learner"`
}

type MemberList []Member

func NewMemberList(em []*etcdserverpb.Member) MemberList {
	ms := make([]Member, len(em))
	for i := range em {
		copier.Copy(&ms[i], *em[i])
		ms[i].ID = fmt.Sprintf("%x", em[i].ID)
	}
	return ms
}

func CheckPeerAddrs(peerAddrs []string) error {
	if len(peerAddrs) == 0 {
		return errors.New("peer address is null")
	}
	return nil
}
