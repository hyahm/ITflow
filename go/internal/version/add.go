package version

import "itflow/cache"

type RespVersion struct {
	Id      int64         `json:"id"`
	Name    string        `json:"name"`
	Project cache.Project `json:"project"`
	RunEnv  string        `json:"runenv"`
	Url     string        `json:"url"`
	BakUrl  string        `json:"bakurl"`
	Date    int64         `json:"date"`
}

type VersionList struct {
	VersionList []*RespVersion `json:"versionlist"`
	Code        int            `json:"code"`
	Msg         string         `json:"message"`
}
