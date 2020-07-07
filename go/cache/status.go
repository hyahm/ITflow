package cache

import (
	"strconv"
	"strings"
)

var CacheSidStatus map[StatusId]Status
var CacheStatusSid map[Status]StatusId

type Status string
type StatusId int64

type StatusList []Status
type StoreStatusId string

var DefaultStatusId StatusId = 0

func (si StatusId) ToString() string {
	return strconv.FormatInt(int64(si), 10)
}

func (si StatusId) ToInt64() int64 {
	return int64(si)
}

// 状态名组转状态id
func (si StatusId) Name() Status {
	if v, ok := CacheSidStatus[si]; ok {
		return v
	} else {
		return ""
	}
}

func (s Status) Id() StatusId {
	if v, ok := CacheStatusSid[s]; ok {
		return v
	} else {
		return DefaultStatusId
	}
}

func (s Status) ToString() string {
	return string(s)
}
func (s Status) Trim() Status {
	return Status(strings.Trim(string(s), " "))
}

func (sl StatusList) ToStore() StoreStatusId {
	tmp := make([]string, 0)
	for _, v := range sl {
		tmp = append(tmp, v.Id().ToString())
	}
	return StoreStatusId(strings.Join(tmp, ","))
}

func (ssi StoreStatusId) ToShow() []Status {
	sl := make([]Status, 0)
	for _, v := range strings.Split(string(ssi), ",") {
		i64, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			continue
		}
		sl = append(sl, StatusId(i64).Name())
	}
	return sl
}
