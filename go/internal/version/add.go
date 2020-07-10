package version

type Version struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Platform string `json:"platform"`
	RunEnv   string `json:"runenv"`
	Url      string `json:"url"`
	BakUrl   string `json:"bakurl"`
	Date     int64  `json:"date"`
}

type VersionList struct {
	VersionList []*Version `json:"versionlist"`
	Code        int        `json:"code"`
	Msg         string     `json:"message"`
}
