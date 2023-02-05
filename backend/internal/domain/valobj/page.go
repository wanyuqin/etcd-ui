package valobj

type PageInfo struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

var (
	defaultPage     = 1
	defaultPageSize = 10
)

func NewPageInfo() *PageInfo {
	return &PageInfo{}
}

func (p *PageInfo) Validate() {
	if p.Page == 0 {
		p.Page = defaultPage
	}

	if p.PageSize == 0 {
		p.PageSize = defaultPageSize
	}

}
