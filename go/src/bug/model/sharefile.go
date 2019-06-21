package model

type Table_sharefile struct {
	Id         int    `json:"id"`
	FilePath   string `json:"filepath"`
	Ruid       int    `json:"rgid"`
	Rgid       int    `json:"uid"`
	IsFile     bool   `json:"isfile"`
	OwnerId    int    `json:"ownerid"`
	Wuid       bool   `json:"wuid"`
	Wgid       bool   `json:"wgid"`
	Size       int64  `json:"size"`
	UpdateTime int    `json:"updatetime"`
	Name       string `json:"name"`
}

type Data_sharefile struct {
	Id         int    `json:"id"`
	FilePath   string `json:"filepath"`
	ReadUser   bool   `json:"readuser"`
	ReadName   string `json:"readname"`
	IsFile     bool   `json:"isfile"`
	IsOwner    bool   `json:"isowner"`
	WriteUser  bool   `json:"writeuser"`
	WriteName  string `json:"writename"`
	Size       int64  `json:"size"`
	UpdateTime int64  `json:"updatetime"`
	Name       string `json:"name"`
	Code       int    `json:"statuscode"`
	OldName    string `json:"oldname"`
}

type List_sharelist struct {
	FDList   []*Data_sharefile `json:"sharelist"`
	Code     int               `json:"statuscode"`
	RealName string            `json:"realname"`
}
