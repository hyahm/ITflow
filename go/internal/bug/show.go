package bug

import (
	"itflow/internal/assist"
	"itflow/internal/comment"
	"itflow/model"
)

type RespShowBug struct {
	Status      string                  `json:"status"`
	Title       string                  `json:"title"`
	Content     string                  `json:"content"`
	Id          int64                   `json:"id"`
	Selectusers assist.Names            `json:"selectuser"`
	Important   string                  `json:"important"`
	Level       string                  `json:"level"`
	Projectname string                  `json:"projectname"`
	Envname     string                  `json:"envname"`
	Version     string                  `json:"version"`
	Code        int                     `json:"code"`
	Msg         string                  `json:"msg,omitempty"`
	Comments    []*comment.Informations `json:"comments,omitempty"`
}

func (rsb *RespShowBug) ShowBug(id interface{}) ([]byte, error) {
	_, err := model.NewBugById(id)
	if err != nil {
		return nil, err
	}

	// art := &Article{

	// 	Title:       bug.Title,
	// 	Id:          bug.ID,
	// 	Important:   cache.CacheIidImportant[bug.ImportanceId],
	// 	Level:       cache.CacheLidLevel[bug.LevelId],
	// 	Projectname: cache.CachePidName[bug.ProjectId],
	// 	Envname:     cache.CacheEidName[bug.EnvId],
	// 	Version:     cache.CacheVidName[bug.VersionId],
	// }

	// art.Content = html.UnescapeString(bug.Content)
	// art.Selectusers = assist.FormatUserlistToShow(bug.OprateUsers)

	// send, _ := json.Marshal(art)
	return []byte(""), nil
}
