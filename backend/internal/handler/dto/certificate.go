package dto

type Certificate struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Key  string `json:"key"`
	Ca   string `json:"ca"`
	Cert string `json:"cert"`
}
