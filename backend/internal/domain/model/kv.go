package model

import (
	"bytes"
	"encoding/json"

	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
	jsonserializer "k8s.io/apimachinery/pkg/runtime/serializer/json"
	"k8s.io/kubectl/pkg/scheme"

	"github.com/wanyuqin/etcd-ui/backend/logger"
)

type KeyTree struct {
	Id       string    `json:"id"`
	Label    string    `json:"label"`
	Children []KeyTree `json:"children"`
}

type KV struct {
	Key            string `json:"key"`
	Value          string `json:"value"`
	Lease          int64  `json:"lease"`
	Version        int64  `json:"version"`
	CreateRevision int64  `json:"create_revision"`
	ModRevision    int64  `json:"mod_revision"`
	Ttl            int64  `json:"ttl"`
}

var EventType = map[mvccpb.Event_EventType]string{
	mvccpb.PUT:    "PUT",
	mvccpb.DELETE: "DELETE",
}

type Event struct {
	EventType string `json:"event_type"`
	Key       string `json:"key"`
}

func NewEvent(e *clientv3.Event) *Event {
	return &Event{
		EventType: EventType[e.Type],
		Key:       string(e.Kv.Key),
	}
}

func NewKV(ekv *mvccpb.KeyValue) *KV {
	var (
		err    error
		buffer bytes.Buffer
	)
	kv := &KV{
		Key:            string(ekv.Key),
		Value:          string(ekv.Value),
		Lease:          ekv.Lease,
		Version:        ekv.Version,
		CreateRevision: ekv.CreateRevision,
		ModRevision:    ekv.ModRevision,
	}
	buffer = bytes.Buffer{}

	decoder := scheme.Codecs.UniversalDeserializer()
	encoder := jsonserializer.NewSerializer(jsonserializer.DefaultMetaFactory, scheme.Scheme, scheme.Scheme, true)
	obj, gvk, err := decoder.Decode(ekv.Value, nil, nil)
	if err != nil {
		logger.Warnf("WARN: unable to decode %s: %v\n", kv.Key, err)
		if json.Valid(ekv.Value) {
			kv.Value = string(ekv.Value)
		}
		return kv
	}
	logger.Debug(gvk)

	err = encoder.Encode(obj, &buffer)
	if err != nil {
		logger.Warnf("WARN: unable to encode %s: %v\n", kv.Key, err)
		if json.Valid(ekv.Value) {
			kv.Value = string(ekv.Value)
		}
		return kv

	}
	kv.Value = buffer.String()

	return kv
}

func NewKeyTree(ks []string) []KeyTree {
	kt := make([]KeyTree, len(ks))
	for i, v := range ks {
		kt[i] = KeyTree{
			Id:       v,
			Label:    v,
			Children: make([]KeyTree, 0),
		}
	}
	return kt
}
