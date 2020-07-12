package bug

import (
	"errors"
	"itflow/cache"
	"itflow/model"
)

func (reb *RespEditBug) ToResp(b *model.Bug) error {
	// 将获取的数据转为可以存表的数据

	reb.Id = b.ID

	reb.Content = b.Content
	reb.Title = b.Title

	var ok bool
	reb.Level = cache.LevelId(b.LevelId).Name()
	if reb.Level == "" {
		return errors.New("没有找到level key")
	}
	reb.Projectname = b.ProjectId.Name()
	if reb.Projectname == "" {
		return errors.New("没有找到project key")
	}
	//
	if reb.Envname, ok = cache.CacheEidName[b.EnvId]; !ok {
		return errors.New("没有找到env key")
	}
	//
	if reb.Important, ok = cache.CacheIidImportant[b.ImportanceId]; !ok {
		return errors.New("没有找到important key")
	}
	if reb.Version, ok = cache.CacheVidName[b.VersionId]; !ok {
		return errors.New("没有找到version key")
	}

	reb.Selectusers = b.OprateUsers.UsersIdToRealName()
	return nil
}

func RespEditBugData(id interface{}) []byte {
	reb := &RespEditBug{}
	bug := &model.Bug{}
	err := bug.NewBugById(id)
	if err != nil {
		reb.Code = 1
		reb.Msg = err.Error()
		return reb.Marshal()
	}
	err = reb.ToResp(bug)
	if err != nil {
		reb.Code = 1
		reb.Msg = err.Error()
		return reb.Marshal()
	}
	return reb.Marshal()
}
