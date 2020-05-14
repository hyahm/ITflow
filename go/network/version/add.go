package version

type Version struct {
	Id           int    `json:"id"`
	Name         string `json:'name'`
	Platform     string `json:'platform'`
	Version      string `json:'version'`
	Runenv       string `json:'runenv'`
	Iphoneurl    string `json:'iphoneurl'`
	Notiphoneurl string `json:'notiphoneurl'`
	Date         int64  `json:"date"`
}

type VersionList struct {
	VersionList []*Version `json:"versionlist"`
	Code        int        `json:"code"`
}
