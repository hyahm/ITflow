package bug

import (
	"encoding/json"
	"itflow/cache"
	"itflow/internal/assist"
	"itflow/internal/comment"

	"github.com/hyahm/golog"
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
	send, err := json.Marshal(rsb)
	if err != nil {
		golog.Error(err)
	}
	return send
}

func (rsb *RespShowBug) Error(msg string) []byte {
	rsb.Code = 1
	rsb.Msg = msg
	return rsb.Marshal()
}

func (rsb *RespShowBug) ErrorE(err error) []byte {
	return rsb.Error(err.Error())
}
