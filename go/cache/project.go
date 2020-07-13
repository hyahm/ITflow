package cache

import (
	"strconv"
	"strings"
)

var CachePidProject map[ProjectId]Project
var CacheProjectPid map[Project]ProjectId

type Project string
type ProjectId int64

type StoreProjectIds string

func (si StoreProjectIds) ToIds() []ProjectId {
	its := make([]ProjectId, 0)
	for _, v := range strings.Split(string(si), ",") {
		i64, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return nil
		}
		its = append(its, ProjectId(i64))
	}
	return its
}

func (si StoreProjectIds) ToNames() []Project {
	its := make([]Project, 0)
	for _, v := range strings.Split(string(si), ",") {
		i64, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return nil
		}
		its = append(its, ProjectId(i64).Name())
	}
	return its
}

var DefaultProjectId ProjectId = 0

func (si ProjectId) ToString() string {
	return strconv.FormatInt(int64(si), 10)
}

func (si ProjectId) ToInt64() int64 {
	return int64(si)
}

// 状态名组转状态id
func (si ProjectId) Name() Project {
	if v, ok := CachePidProject[si]; ok {
		return v
	} else {
		return ""
	}
}

func (s Project) Id() ProjectId {
	s = s.Trim()
	if v, ok := CacheProjectPid[s]; ok {
		return v
	} else {
		return DefaultProjectId
	}
}

func (s Project) ToString() string {
	return string(s)
}
func (s Project) Trim() Project {
	return Project(strings.Trim(string(s), " "))
}
