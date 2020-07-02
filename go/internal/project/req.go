package project

type Project struct {
	Id          int64  `json:"id"`
	ProjectName string `json:"projectname"`
}

type ProjectList struct {
	Plist []*Project `json:"projectlist"`
	Code  int        `json:"code"`
}
