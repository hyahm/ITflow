package version

// type RespVersion struct {
// 	Id      int64  `json:"id"`
// 	Name    string `json:"name"`
// 	Project string `json:"project"`
// 	RunEnv  string `json:"runenv"`
// 	Url     string `json:"url"`
// 	BakUrl  string `json:"bakurl"`
// 	Date    int64  `json:"date"`
// }

// type VersionList struct {
// 	VersionList []*RespVersion `json:"versionlist"`
// 	Code        int            `json:"code"`
// 	Msg         string         `json:"msg"`
// }

// func (vl *VersionList) Marshal() []byte {
// 	send, err := json.Marshal(vl)
// 	if err != nil {
// 		golog.Error(err)
// 	}
// 	return send
// }
// func (vl *VersionList) Error(msg string) []byte {
// 	vl.Code = 1
// 	vl.Msg = msg
// 	return vl.Marshal()
// }
// func (vl *VersionList) ErrorE(err error) []byte {
// 	return vl.Error(err.Error())
// }
