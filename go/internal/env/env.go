package env

type Envlist struct {
	Elist []*Env `json:"envlist"`
	Code  int    `json:"code"`
}
type Env struct {
	Id      int64  `json:"id"`
	EnvName string `json:"envname"`
}
