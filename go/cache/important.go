package cache

import (
	"strconv"
	"strings"
)

var CacheIidImportant map[ImportantId]Important
var CacheImportantIid map[Important]ImportantId

type Important string
type ImportantId int64

type StoreImportantIds string

func (si StoreImportantIds) ToIds() []ImportantId {
	its := make([]ImportantId, 0)
	for _, v := range strings.Split(string(si), ",") {
		i64, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return nil
		}
		its = append(its, ImportantId(i64))
	}
	return its
}

func (si StoreImportantIds) ToNames() []Important {
	its := make([]Important, 0)
	for _, v := range strings.Split(string(si), ",") {
		i64, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return nil
		}
		its = append(its, ImportantId(i64).Name())
	}
	return its
}

var DefaultImportantId ImportantId = 0

func (si ImportantId) ToString() string {
	return strconv.FormatInt(int64(si), 10)
}

func (si ImportantId) ToInt64() int64 {
	return int64(si)
}

// 状态名组转状态id
func (si ImportantId) Name() Important {
	if v, ok := CacheIidImportant[si]; ok {
		return v
	} else {
		return ""
	}
}

func (s Important) Id() ImportantId {
	s = s.Trim()
	if v, ok := CacheImportantIid[s]; ok {
		return v
	} else {
		return DefaultImportantId
	}
}

func (s Important) ToString() string {
	return string(s)
}
func (s Important) Trim() Important {
	return Important(strings.Trim(string(s), " "))
}
