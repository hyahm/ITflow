package restful

type Resp struct {
	Headers []*Header `json:"header"`
	Resp    string    `json:"resp"`
	Url     string    `json:"url"`
	Method  string    `json:"method"`
}

type Header struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
