package cache

import (
	"strconv"
	"strings"
)

var CacheEidEnv map[EnvId]Env
var CacheEnvEid map[Env]EnvId

type Env string
type EnvId int64

type StoreEnvIds string

func (si StoreEnvIds) ToIds() []EnvId {
	its := make([]EnvId, 0)
	for _, v := range strings.Split(string(si), ",") {
		i64, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return nil
		}
		its = append(its, EnvId(i64))
	}
	return its
}

func (si StoreEnvIds) ToNames() []Env {
	its := make([]Env, 0)
	for _, v := range strings.Split(string(si), ",") {
		i64, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return nil
		}
		its = append(its, EnvId(i64).Name())
	}
	return its
}

var DefaultEnvId EnvId = 0

func (si EnvId) ToString() string {
	return strconv.FormatInt(int64(si), 10)
}

func (si EnvId) ToInt64() int64 {
	return int64(si)
}

// 状态名组转状态id
func (si EnvId) Name() Env {
	if v, ok := CacheEidEnv[si]; ok {
		return v
	} else {
		return ""
	}
}

func (s Env) Id() EnvId {
	s = s.Trim()
	if v, ok := CacheEnvEid[s]; ok {
		return v
	} else {
		return DefaultEnvId
	}
}

func (s Env) ToString() string {
	return string(s)
}
func (s Env) Trim() Env {
	return Env(strings.Trim(string(s), " "))
}
