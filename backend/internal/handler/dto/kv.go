package dto

type KV struct {
	Key            string `json:"key"`
	Value          string `json:"value"`
	Lease          int64  `json:"lease"`
	Version        int64  `json:"version"`
	CreateRevision int64  `json:"create_revision"`
	ModRevision    int64  `json:"mod_revision"`
	Ttl            int64  `json:"ttl"`
}
