package bug

import "itflow/network/comment"

type ShowBug struct {
	Title      string                  `json:"title,omitempty" type:"string" need:"否" default:"" information:"标题"`
	Content    string                  `json:"content,omitempty" type:"string" need:"否" default:"" information:"内容"`
	Appversion string                  `json:"appversion,omitempty" type:"string" need:"否" default:"" information:"版本号"`
	Comment    []*comment.Informations `json:"comment,omitempty" type:"array" need:"否" default:"" information:"事件记录"`
	Status     string                  `json:"status,omitempty" type:"string" need:"否" default:"" information:"状态名"`
	Id         int                     `json:"id,omitempty" type:"int" need:"否" default:"0" information:"bug的id"`
	Code       int                     `json:"code" type:"int" need:"是" default:"0" information:"无效"`
}
