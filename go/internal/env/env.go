package env

import "itflow/cache"

type Envlist struct {
	Elist []*Env `json:"envlist"`
	Code  int    `json:"code"`
}
type Env struct {
	Id      cache.EnvId `json:"id"`
	EnvName cache.Env   `json:"envname"`
}
