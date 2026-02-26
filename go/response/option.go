package response

type Option struct {
	Label string `json:"label"`
	Value int64  `json:"value"`
}

type Pager struct {
	PageIndex int `json:"page_index"`
	PageSize  int `json:"page_size"`
}
