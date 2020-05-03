package bug

type Status struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type ListStatus struct {
	StatusList []*Status `json:"statuslist,omitempty" type:"array" need:"否" information:"状态列表"`
	Code       int       `json:"code" type:"int" need:"是" information:"错误码"`
}
