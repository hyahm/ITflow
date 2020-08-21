package bug

import (
	"encoding/json"
	"itflow/cache"
	"itflow/internal/assist"
	"itflow/internal/comment"
)

type RespShowBug struct {
	Status      string                  `json:"status"`
	Title       string                  `json:"title"`
	Content     string                  `json:"content"`
	Id          int64                   `json:"id"`
	Selectusers assist.Names            `json:"selectuser"`
	Important   cache.Important         `json:"important"`
	Level       cache.Level             `json:"level"`
	Projectname cache.Project           `json:"projectname"`
	Envname     cache.Env               `json:"envname"`
	Version     string                  `json:"version"`
	Code        int                     `json:"code"`
	Msg         string                  `json:"msg,omitempty"`
	Comments    []*comment.Informations `json:"comments,omitempty"`
}

func (rsb *RespShowBug) Marshal() []byte {
	send, _ := json.Marshal(rsb)
	return send
}
