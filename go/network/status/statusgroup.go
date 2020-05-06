package status

type StatusGroup struct {
	Id         int64    `json:"id" type:"int" need:"是" default:"0" information:"无效"`
	StatusList []string `json:"checklist, omitempty" type:"array" need:"是" default:"[]" information:"状态名列表"`
	Name       string   `json:"name" type:"int" need:"是" default:"0" information:"状态组名"`
}
